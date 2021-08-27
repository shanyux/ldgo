/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"reflect"
	"strings"
)

type WhereOption interface {
	Option

	And(where interface{}) WhereOption
	Or(where interface{}) WhereOption

	buildWhere() whereResult
}

func Where(where interface{}) WhereOption {
	if where == nil {
		panic("the where type must not be nil")
	}

	if w, ok := where.(WhereOption); ok {
		return w
	}

	val := reflect.ValueOf(where)
	w := getWhereReflect(val.Type())

	return &whereOption{
		Value: val,
		Where: w,
	}
}

type whereOption struct {
	Value reflect.Value
	Where *whereReflect
}

func (that *whereOption) buildGorm(db *GormDb) *GormDb {
	return that.Where.buildGorm(db, that.Value)
}

func (that *whereOption) buildWhere() whereResult {
	return that.Where.buildWhere(that.Value)
}

func (that *whereOption) toTree() whereOptionTree {
	return whereOptionTree{
		Wheres: []whereOptionTreeNode{
			{Where: that},
		},
	}
}

func (that *whereOption) And(o interface{}) WhereOption {
	return that.toTree().And(o)
}

func (that *whereOption) Or(o interface{}) WhereOption {
	return that.toTree().Or(o)
}

type whereOptionTreeNode struct {
	Or    bool        `json:"or"`
	Where WhereOption `json:"where"`
}

type whereOptionTree struct {
	Wheres []whereOptionTreeNode `json:"wheres"`
}

func (that whereOptionTree) buildWhere() whereResult {
	res := that.Wheres[0].Where.buildWhere()
	if len(that.Wheres) == 1 {
		return res
	}

	res.Query = "(" + res.Query

	for _, w := range that.Wheres[1:] {
		tmp := w.Where.buildWhere()
		symbol := " AND "
		if w.Or {
			symbol = " OR "
		}

		res.Query = res.Query + symbol + tmp.Query
		res.Args = append(res.Args, tmp.Args...)
	}

	res.Query = res.Query + ")"
	return res
}

func (that whereOptionTree) buildGorm(db *GormDb) *GormDb {
	res := that.buildWhere()
	if strings.HasPrefix(res.Query, "(") && strings.HasSuffix(res.Query, ")") {
		res.Query = res.Query[1 : len(res.Query)-1]
	}

	if res.IsValid() {
		db = db.Where(res.Query, res.Args...)
	}

	return db
}

func (that whereOptionTree) And(where interface{}) WhereOption {
	return that.append(false, where)
}

func (that whereOptionTree) Or(where interface{}) WhereOption {
	return that.append(true, where)
}

func (that whereOptionTree) append(or bool, where interface{}) WhereOption {
	that.Wheres = append(that.Wheres, whereOptionTreeNode{
		Or:    or,
		Where: Where(where),
	})

	return that
}
