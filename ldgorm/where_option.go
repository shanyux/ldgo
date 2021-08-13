/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

type WhereOption interface {
	Option

	And(o WhereOption) WhereOption
	Or(o WhereOption) WhereOption
}

func Where(cond interface{}) WhereOption {
	if cond == nil {
		return nil
	}

	val := reflect.ValueOf(cond)
	w := getWhereInfo(val.Type())

	return &whereOption{
		Value: val,
		Where: w,
	}
}

type whereOption struct {
	Value reflect.Value
	Where *whereReflect
}

func (that *whereOption) buildGorm(db *gorm.DB) *gorm.DB {
	root := db
	return that.Where.buildWhere(root, that.Value)
}

func (that *whereOption) toTree() whereOptionTree {
	return whereOptionTree{
		Wheres: []whereOptionTreeNode{
			{Where: that},
		},
	}
}

func (that *whereOption) And(o WhereOption) WhereOption {
	return that.toTree().And(o)
}

func (that *whereOption) Or(o WhereOption) WhereOption {
	return that.toTree().Or(o)
}

type whereOptionTreeNode struct {
	Or    bool        `json:"or"`
	Where WhereOption `json:"where"`
}

type whereOptionTree struct {
	Wheres []whereOptionTreeNode `json:"wheres"`
}

func (that whereOptionTree) buildGorm(db *gorm.DB) *gorm.DB {
	root := db

	db = that.Wheres[0].Where.buildGorm(root)
	for _, w := range that.Wheres[1:] {
		tmp := w.Where.buildGorm(root)
		if w.Or {
			db = db.Or(tmp)
		} else {
			db = db.Where(tmp)
		}
	}

	return db
}

func (that whereOptionTree) And(o WhereOption) WhereOption {
	return that.append(false, o)
}

func (that whereOptionTree) Or(o WhereOption) WhereOption {
	return that.append(true, o)
}

func (that whereOptionTree) append(or bool, o WhereOption) WhereOption {
	that.Wheres = append(that.Wheres, whereOptionTreeNode{
		Or:    or,
		Where: o,
	})

	return that
}
