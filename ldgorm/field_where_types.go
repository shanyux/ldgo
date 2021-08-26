/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"fmt"
	"reflect"
)

type FieldWherer interface {
	isEmpty() bool

	And(b FieldWherer) FieldWherer
	Or(b FieldWherer) FieldWherer

	buildWhere(field string) whereResult
}

func getWhereValue(v interface{}) (interface{}, bool) {
	if v == nil {
		return nil, false
	}
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil, false
		}
		return val.Elem().Interface(), true
	}

	return v, true
}

type whereResult struct {
	Query string
	Args  []interface{}
}

func (r whereResult) IsValid() bool { return len(r.Query) != 0 }

type fieldWhereEmpty struct{}

func (_ fieldWhereEmpty) isEmpty() bool                 { return true }
func (_ fieldWhereEmpty) And(b FieldWherer) FieldWherer { return b }
func (_ fieldWhereEmpty) Or(b FieldWherer) FieldWherer  { return b }

func (_ fieldWhereEmpty) buildGorm(db GormDb, field string) GormDb { return db }
func (_ fieldWhereEmpty) buildWhere(field string) whereResult      { return whereResult{} }

type fieldWhereBase struct{}

func (_ fieldWhereBase) isEmpty() bool { return false }

type fieldWhereTreeNode struct {
	Or    bool        `json:"or"`
	Where FieldWherer `json:"where"`
}

type fieldWhereTree struct {
	fieldWhereBase

	Wheres []fieldWhereTreeNode `json:"wheres"`
}

func (that fieldWhereTree) buildWhere(field string) whereResult {
	res := that.Wheres[0].Where.buildWhere(field)
	if len(that.Wheres) == 1 {
		return res
	}

	res.Query = "(" + res.Query

	for _, w := range that.Wheres[1:] {
		tmp := w.Where.buildWhere(field)
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

func (that fieldWhereTree) And(b FieldWherer) FieldWherer {
	if b.isEmpty() {
		return that
	}

	that.Wheres = append(that.Wheres, fieldWhereTreeNode{
		Or:    false,
		Where: b,
	})
	return that
}

func (that fieldWhereTree) Or(b FieldWherer) FieldWherer {
	if b.isEmpty() {
		return that
	}

	that.Wheres = append(that.Wheres, fieldWhereTreeNode{
		Or:    true,
		Where: b,
	})
	return that
}

type fieldWhere struct {
	fieldWhereBase

	Query  string        `json:"query"`
	Values []interface{} `json:"values"`
}

func (that fieldWhere) buildWhere(field string) whereResult {
	return whereResult{
		Query: fmt.Sprintf("`%s`%s", field, that.Query),
		Args:  that.Values,
	}
}

func (that fieldWhere) toTree() fieldWhereTree {
	return fieldWhereTree{
		Wheres: []fieldWhereTreeNode{{
			Where: that,
		}},
	}
}

func (that fieldWhere) And(b FieldWherer) FieldWherer {
	return that.toTree().And(b)
}

func (that fieldWhere) Or(b FieldWherer) FieldWherer {
	return that.toTree().Or(b)
}

func newFieldWhere(query string, values ...interface{}) FieldWherer {
	return fieldWhere{
		Query:  query,
		Values: values,
	}
}

func newFieldWhereWithCheck(query string, value interface{}) FieldWherer {
	val, isSet := getWhereValue(value)
	if !isSet {
		return fieldWhereEmpty{}
	}

	return newFieldWhere(query, val)
}
