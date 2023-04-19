/*
 * Copyright (C) distroy
 */

package gorm

import (
	"fmt"
	"hash/crc32"
	"strings"

	"github.com/distroy/ldgo/ldconv"
	"github.com/distroy/ldgo/ldrand"
	"github.com/jinzhu/gorm"
)

type gormDb = gorm.DB

type Logger interface {
	Print(v ...interface{})
}

var (
	// queryHintReplacer = strings.NewReplacer("/*", " ", "*/", " ")
	queryHintReplacer = strings.NewReplacer("/*", " ", "*/", " ")
)

type GormDb struct {
	*gormDb // it is currently used db, it is the master by default

	master  *gorm.DB
	slavers []*gorm.DB

	txLvl int
	inTx  bool
}

func (w *GormDb) panicTxLevelLessZero() {
	panic("tx level must not be less than zero")
}

func (w *GormDb) panicTxCommittedOrRollbacked() {
	panic("tx can not be committed or rollbacked again")
}

func (w *GormDb) clone() *GormDb {
	c := *w
	return &c
}

// New clone a new db connection without search conditions
func (w *GormDb) New() *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.New()
	return w
}

func (w *GormDb) Get() *gorm.DB {
	return w.gormDb
}

func (w *GormDb) Close() error {
	var err error

	db := w.master
	if e := db.Close(); e != nil {
		err = e
	}

	for _, db := range w.slavers {
		if e := db.Close(); err == nil && e != nil {
			err = e
		}
	}

	return err
}

// Set will set the db currently used.
// If master is not set, will set the master db also.
func (w *GormDb) Set(db *gorm.DB) *GormDb {
	w = w.clone()

	w.gormDb = db
	if w.master == nil {
		w.master = db
	}

	return w
}

func (w *GormDb) SetMaster(db *gorm.DB) *GormDb {
	w = w.clone()
	w.master = db
	return w
}

func (w *GormDb) AddSlaver(dbs ...*gorm.DB) *GormDb {
	if len(dbs) == 0 {
		return w
	}

	w = w.clone()
	slavers := make([]*gorm.DB, 0, len(w.slavers)+len(dbs))
	slavers = append(slavers, w.slavers...)
	slavers = append(slavers, dbs...)
	w.slavers = slavers
	return w
}

// UseMaster must be called before all Create/Update/Query/Delete methods
func (w *GormDb) UseMaster() *GormDb {
	if w.gormDb == w.master {
		return w
	}

	w = w.clone()
	w.gormDb = w.master
	return w
}

// UseSlaver must be called before all Query methods
//
// key must be int/int{8-64}/uint/uint{8-64}/uintptr/string/[]byte.
// if key is not set, will use rand slaver
func (w *GormDb) UseSlaver(key ...interface{}) *GormDb {
	n := len(w.slavers)
	switch n {
	case 0:
		return w

	case 1:
		w = w.clone()
		w.gormDb = w.slavers[0]
		return w
	}

	hash := w.getHashByKey(key)
	idx := hash % uint(n)

	w = w.clone()
	w.gormDb = w.slavers[idx]
	return w
}

func (w *GormDb) getHashByKey(keys []interface{}) uint {
	switch v := keys[0].(type) {
	case int:
		return uint(v)
	case int8:
		return uint(v)
	case int16:
		return uint(v)
	case int32:
		return uint(v)
	case int64:
		return uint(v)

	case uint:
		return uint(v)
	case uint8:
		return uint(v)
	case uint16:
		return uint(v)
	case uint32:
		return uint(v)
	case uint64:
		return uint(v)
	case uintptr:
		return uint(v)

	case string:
		return uint(crc32.ChecksumIEEE(ldconv.StrToBytesUnsafe(v)))

	case []byte:
		return uint(crc32.ChecksumIEEE(v))
	}

	return ldrand.Uint()
}

func (w *GormDb) withOption(opts ...func(db *gorm.DB) *gorm.DB) *GormDb {
	w = w.clone()

	apply := func(db *gorm.DB, opts []func(db *gorm.DB) *gorm.DB) *gorm.DB {
		for _, opt := range opts {
			db = opt(db)
		}
		return db
	}

	master := apply(w.master, opts)

	current := w.gormDb
	if current == w.master {
		current = master
	}

	slavers := make([]*gorm.DB, 0, len(w.slavers))
	for _, db0 := range w.slavers {
		db1 := apply(db0, opts)
		slavers = append(slavers, db1)
		if current == db0 {
			current = db1
		}
	}

	if current == w.gormDb {
		current = apply(current, opts)
	}

	w.master = master
	w.gormDb = current
	w.slavers = slavers
	return w
}

// WithLogger can be called before or after UseMaster/UseSlaver
func (w *GormDb) WithLogger(l Logger) *GormDb {
	return w.withOption(func(db *gorm.DB) *gorm.DB {
		db = db.LogMode(true)
		db.SetLogger(l)
		return db
	})
}

func (w *GormDb) Model(value interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Model(value)
	return w
}

func (w *GormDb) Table(table string) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Table(table)
	return w
}

func (w *GormDb) Transaction(fc func(tx *GormDb) error) (err error) {
	if w.txLvl > 0 {
		return fc(w)
	}

	paniced := true
	tx := w.Begin()
	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if paniced || err != nil {
			tx.Rollback()
		}
	}()

	err = fc(tx)

	if err == nil {
		err = tx.Commit().Error
	}

	paniced = false
	return
}

func (w *GormDb) Begin() *GormDb {
	if w.txLvl < 0 {
		w.panicTxLevelLessZero()
	}

	w = w.clone()

	if w.txLvl == 0 {
		w.gormDb = w.gormDb.Begin()
	}

	w.inTx = true
	w.txLvl++
	return w
}

func (w *GormDb) Commit() *GormDb {
	if !w.inTx {
		w.panicTxCommittedOrRollbacked()
	}

	w = w.clone()

	w.txLvl--
	if w.txLvl < 0 {
		w.panicTxLevelLessZero()
	}

	if w.txLvl == 0 {
		w.gormDb = w.gormDb.Commit()
	}

	w.inTx = false
	return w
}

func (w *GormDb) Rollback() *GormDb {
	if !w.inTx {
		w.panicTxCommittedOrRollbacked()
	}

	w = w.clone()

	w.txLvl--
	if w.txLvl < 0 {
		w.panicTxLevelLessZero()
	}

	if w.txLvl == 0 {
		w.gormDb = w.gormDb.Rollback()
	}

	w.inTx = false
	return w
}

func (w *GormDb) RollbackUnlessCommitted() *GormDb {
	if !w.inTx {
		return w
	}

	w = w.clone()

	w.txLvl--
	if w.txLvl < 0 {
		w.panicTxLevelLessZero()
	}

	if w.txLvl == 0 {
		w.gormDb = w.gormDb.RollbackUnlessCommitted()
	}

	w.inTx = false
	return w
}

func (w *GormDb) Select(query interface{}, args ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Select(query, args...)
	return w
}

func (w *GormDb) Group(query string) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Group(query)
	return w
}

func (w *GormDb) Having(query interface{}, args ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Having(query, args...)
	return w
}

func (w *GormDb) Joins(query string, args ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Joins(query, args...)
	return w
}

func (w *GormDb) HasTable(value interface{}) bool {
	return w.gormDb.HasTable(value)
}

func (w *GormDb) CreateTable(models ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.CreateTable(models...)
	return w
}

func (w *GormDb) DropTable(models ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.DropTable(models...)
	return w
}

func (w *GormDb) DropTableIfExists(models ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.DropTableIfExists(models...)
	return w
}

func (w *GormDb) Where(query interface{}, args ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Where(query, args...)
	return w
}

// Order specify order when retrieve records from database, set reorder to `true` to overwrite defined conditions
//
//	db.Order("name DESC")
//	db.Order("name DESC", true) // reorder
//	db.Order(gorm.Expr("name = ? DESC", "first")) // sql expression
func (w *GormDb) Order(value interface{}, reorder ...bool) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Order(value, reorder...)
	return w
}

func (w *GormDb) Limit(limit interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Limit(limit)
	return w
}

func (w *GormDb) Offset(offset interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Offset(offset)
	return w
}

func (w *GormDb) Save(value interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Save(value)
	return w
}

func (w *GormDb) Create(value interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Create(value)
	return w
}

func (w *GormDb) Update(attrs ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Update(attrs...)
	return w
}

func (w *GormDb) Updates(values interface{}, ignoreProtectedAttrs ...bool) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Updates(values, ignoreProtectedAttrs...)
	return w
}

func (w *GormDb) FirstOrCreate(out interface{}, where ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.FirstOrCreate(out, where...)
	return w
}

func (w *GormDb) Delete(value interface{}, where ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Delete(value, where...)
	return w
}

func (w *GormDb) First(out interface{}, where ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.First(out, where...)
	return w
}

func (w *GormDb) Find(out interface{}, where ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Find(out, where...)
	return w
}

func (w *GormDb) Take(out interface{}, where ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Take(out, where...)
	return w
}

func (w *GormDb) Last(out interface{}, where ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Last(out, where...)
	return w
}

func (w *GormDb) Count(out interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Count(out)
	return w
}

func (w *GormDb) Exec(sql string, values ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Exec(sql, values...)
	return w
}

func (w *GormDb) Raw(sql string, values ...interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Raw(sql, values...)
	return w
}

func (w *GormDb) Scan(out interface{}) *GormDb {
	w = w.clone()
	w.gormDb = w.gormDb.Scan(out)
	return w
}

// WithQueryHint add hint comment before sql
//
// WithQueryHint must be called after UseSlaver
func (w *GormDb) WithQueryHint(hint string) *GormDb {
	replacer := queryHintReplacer
	hint = replacer.Replace(hint)
	hint = fmt.Sprintf("/* %s */ ", hint)

	w = w.clone()
	w.gormDb = w.gormDb.Set("gorm:query_hint", hint)
	return w
}
