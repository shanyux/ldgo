/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"github.com/distroy/ldgo/ldlogger"
	"github.com/jinzhu/gorm"
)

type GormDb interface {
	// New clone a new db connection without search conditions
	New() GormDb
	Close() error

	Get() *gorm.DB
	Set(db *gorm.DB) GormDb

	RowsAffected() int64

	WithLogger(l ldlogger.Logger) GormDb

	Error() error
	RecordNotFound() bool

	Callback() *gorm.Callback
	Model(value interface{}) GormDb
	NewScope(value interface{}) *gorm.Scope

	Transaction(fc func(tx GormDb) error) error
	Begin() GormDb
	Commit() GormDb
	Rollback() GormDb
	RollbackUnlessCommitted() GormDb

	HasTable(value interface{}) bool
	CreateTable(models ...interface{}) GormDb
	DropTable(values ...interface{}) GormDb
	DropTableIfExists(values ...interface{}) GormDb

	Select(query interface{}, args ...interface{}) GormDb
	Group(query string) GormDb
	Having(query interface{}, values ...interface{}) GormDb
	Joins(query string, args ...interface{}) GormDb

	// Order specify order when retrieve records from database, set reorder to `true` to overwrite defined conditions
	//     db.Order("name DESC")
	//     db.Order("name DESC", true) // reorder
	//     db.Order(gorm.Expr("name = ? DESC", "first")) // sql expression
	Order(value interface{}, reorder ...bool) GormDb

	Where(query interface{}, args ...interface{}) GormDb
	Offset(offset interface{}) GormDb
	Limit(limit interface{}) GormDb

	Save(value interface{}) GormDb
	Create(value interface{}) GormDb
	Update(attrs ...interface{}) GormDb
	Updates(values interface{}, ignoreProtectedAttrs ...bool) GormDb
	Delete(value interface{}, where ...interface{}) GormDb

	First(out interface{}, where ...interface{}) GormDb
	FirstOrCreate(out interface{}, where ...interface{}) GormDb
	Find(out interface{}, where ...interface{}) GormDb
	Take(out interface{}, where ...interface{}) GormDb
	Last(out interface{}, where ...interface{}) GormDb
	Count(out interface{}) GormDb

	Exec(sql string, values ...interface{}) GormDb
	Raw(sql string, values ...interface{}) GormDb
	Scan(out interface{}) GormDb
}

type gormWapper struct {
	db      *gorm.DB
	txLevel int
	inSubTx bool
}

func NewGorm(db *gorm.DB) GormDb {
	return &gormWapper{db: db}
}

func (w *gormWapper) panicTxLevelLessZero() {
	panic("tx level must not be less than zero")
}

func (w *gormWapper) panicNotInSubTx() {
	panic("sub tx can not be committed or rollbacked again")
}

func (w *gormWapper) clone() *gormWapper {
	c := *w
	return &c
}

func (w *gormWapper) New() GormDb {
	w = w.clone()
	w.db = w.db.New()
	return w
}

func (w *gormWapper) Close() error {
	return w.db.Close()
}

func (w *gormWapper) Get() *gorm.DB {
	return w.db
}

func (w *gormWapper) Set(db *gorm.DB) GormDb {
	w = w.clone()
	w.db = db
	return w
}

func (w *gormWapper) WithLogger(l ldlogger.Logger) GormDb {
	w = w.clone()
	w.db = w.db.LogMode(true)
	w.db.SetLogger(l.Wrap())
	return w
}

func (w *gormWapper) Callback() *gorm.Callback {
	return w.db.Callback()
}

func (w *gormWapper) Error() error {
	return w.db.Error
}

func (w *gormWapper) RecordNotFound() bool {
	return w.db.RecordNotFound()
}

func (w *gormWapper) RowsAffected() int64 {
	return w.db.RowsAffected
}

func (w *gormWapper) NewScope(value interface{}) *gorm.Scope {
	return w.db.NewScope(value)
}

func (w *gormWapper) Model(value interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Model(value)
	return w
}

func (w *gormWapper) Transaction(fc func(tx GormDb) error) (err error) {
	if w.txLevel > 0 {
		return fc(w)
	}

	panicked := true
	tx := w.Begin()
	defer func() {
		// Make sure to rollback when panic, Block error or Commit error
		if panicked || err != nil {
			tx.Rollback()
		}
	}()

	err = fc(tx)

	if err == nil {
		err = tx.Commit().Error()
	}

	panicked = false
	return
}

func (w *gormWapper) Begin() GormDb {
	if w.txLevel < 0 {
		w.panicTxLevelLessZero()
	}

	w = w.clone()

	if w.txLevel > 0 {
		w.txLevel++
		return w
	}

	w.db = w.db.Begin()
	w.inSubTx = true
	w.txLevel++
	return w
}

func (w *gormWapper) Commit() GormDb {
	if !w.inSubTx {
		return w
	}

	w = w.clone()

	w.txLevel--
	if w.txLevel < 0 {
		w.panicTxLevelLessZero()
	}

	if w.txLevel == 0 {
		w.db = w.db.Rollback()
	}

	w.inSubTx = false
	return w
}

func (w *gormWapper) Rollback() GormDb {
	if !w.inSubTx {
		return w
	}

	w = w.clone()

	w.txLevel--
	if w.txLevel < 0 {
		w.panicTxLevelLessZero()
	}

	if w.txLevel == 0 {
		w.db = w.db.Rollback()
	}

	w.inSubTx = false
	return w
}

func (w *gormWapper) RollbackUnlessCommitted() GormDb {
	if !w.inSubTx {
		return w
	}

	w = w.clone()

	w.txLevel--
	if w.txLevel < 0 {
		w.panicTxLevelLessZero()
	}

	if w.txLevel == 0 {
		w.db = w.db.Rollback()
	}

	w.inSubTx = false
	return w
}

func (w *gormWapper) Select(query interface{}, args ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Select(query, args...)
	return w
}

func (w *gormWapper) Group(query string) GormDb {
	w = w.clone()
	w.db = w.db.Group(query)
	return w
}

func (w *gormWapper) Having(query interface{}, args ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Having(query, args...)
	return w
}

func (w *gormWapper) Joins(query string, args ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Joins(query, args...)
	return w
}

func (w *gormWapper) HasTable(value interface{}) bool {
	return w.db.HasTable(value)
}

func (w *gormWapper) CreateTable(models ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.CreateTable(models...)
	return w
}

func (w *gormWapper) DropTable(models ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.DropTable(models...)
	return w
}

func (w *gormWapper) DropTableIfExists(models ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.DropTableIfExists(models...)
	return w
}

func (w *gormWapper) Where(query interface{}, args ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Where(query, args...)
	return w
}

func (w *gormWapper) Order(value interface{}, reorder ...bool) GormDb {
	w = w.clone()
	w.db = w.db.Order(value, reorder...)
	return w
}

func (w *gormWapper) Limit(limit interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Limit(limit)
	return w
}

func (w *gormWapper) Offset(offset interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Offset(offset)
	return w
}

func (w *gormWapper) Save(value interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Save(value)
	return w
}

func (w *gormWapper) Create(value interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Create(value)
	return w
}

func (w *gormWapper) Update(attrs ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Update(attrs...)
	return w
}

func (w *gormWapper) Updates(values interface{}, ignoreProtectedAttrs ...bool) GormDb {
	w = w.clone()
	w.db = w.db.Updates(values, ignoreProtectedAttrs...)
	return w
}

func (w *gormWapper) FirstOrCreate(out interface{}, where ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.FirstOrCreate(out, where...)
	return w
}

func (w *gormWapper) Delete(value interface{}, where ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Delete(value, where...)
	return w
}

func (w *gormWapper) First(out interface{}, where ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.First(out, where...)
	return w
}

func (w *gormWapper) Find(out interface{}, where ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Find(out, where...)
	return w
}

func (w *gormWapper) Take(out interface{}, where ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Take(out, where...)
	return w
}

func (w *gormWapper) Last(out interface{}, where ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Last(out, where...)
	return w
}

func (w *gormWapper) Count(out interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Count(out)
	return w
}

func (w *gormWapper) Exec(sql string, values ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Exec(sql, values...)
	return w
}

func (w *gormWapper) Raw(sql string, values ...interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Raw(sql, values...)
	return w
}

func (w *gormWapper) Scan(out interface{}) GormDb {
	w = w.clone()
	w.db = w.db.Scan(out)
	return w
}
