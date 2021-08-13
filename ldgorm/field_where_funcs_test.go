/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"testing"

	"github.com/distroy/ldgo/ldhook"
	"github.com/jinzhu/gorm"
	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/proto"
)

func Test_Between(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Between(nil, nil)", func() {
			cond := Between(nil, nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Between(int, nil)", func() {
			cond := Between(0, nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Between(string, nil)", func() {
			cond := Between("abc", nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Between(*int, nil)", func() {
			cond := Between(proto.Int32(0), nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Between(*int(nil), nil)", func() {
			cond := Between((*int)(nil), nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Between(*string, nil)", func() {
			cond := Between(proto.String("abc"), nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Between(*string(nil), nil)", func() {
			cond := Between((*string)(nil), nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Between(nil, int)", func() {
			cond := Between(nil, 0)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Between(nil, string)", func() {
			cond := Between(nil, "abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Between(nil, *int)", func() {
			cond := Between(nil, proto.Int32(0))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Between(nil, *int(nil))", func() {
			cond := Between(nil, (*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Between(nil, *string)", func() {
			cond := Between(nil, proto.String("abc"))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Between(nil, *string(nil))", func() {
			cond := Between(nil, (*string)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Between(int, int)", func() {
			cond := Between(0, 10)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " BETWEEN ? AND ?", Args: []interface{}{0, 10}},
			})
		})
		convey.Convey("Between(int, int) && min == max", func() {
			cond := Between(3, 3)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{3}},
			})
		})
		convey.Convey("Between(*int, *int)", func() {
			cond := Between(proto.Int32(0), proto.Int32(10))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " BETWEEN ? AND ?", Args: []interface{}{int32(0), int32(10)}},
			})
		})
		convey.Convey("Between(*int, *int) && min == max", func() {
			cond := Between(proto.Int32(3), proto.Int32(3))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{int32(3)}},
			})
		})
		convey.Convey("Between(*int(nil), *int)", func() {
			cond := Between((*int)(nil), proto.Int32(10))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{int32(10)}},
			})
		})
		convey.Convey("Between(*int, *int(nil))", func() {
			cond := Between(proto.Int32(0), (*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Between(*int(nil), *int(nil))", func() {
			cond := Between((*int)(nil), (*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

	})
}

func Test_Gt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Gt(nil)", func() {
			cond := Gt(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Gt(int)", func() {
			cond := Gt(0)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " > ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Gt(string)", func() {
			cond := Gt("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " > ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Gt(*int)", func() {
			cond := Gt(proto.Int32(0))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " > ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Gt(*int(nil))", func() {
			cond := Gt((*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Gt(*string)", func() {
			cond := Gt(proto.String("abc"))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " > ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Gt(*string(nil))", func() {
			cond := Gt((*string)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
	})
}

func Test_Lt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Lt(nil)", func() {
			cond := Lt(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Lt(int)", func() {
			cond := Lt(0)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " < ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Lt(string)", func() {
			cond := Lt("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " < ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Lt(*int)", func() {
			cond := Lt(proto.Int32(0))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " < ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Lt(*int(nil))", func() {
			cond := Lt((*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Lt(*string)", func() {
			cond := Lt(proto.String("abc"))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " < ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Lt(*string(nil))", func() {
			cond := Lt((*string)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
	})
}

func Test_Gte(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Gte(nil)", func() {
			cond := Gte(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Gte(int)", func() {
			cond := Gte(0)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Gte(string)", func() {
			cond := Gte("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Gte(*int)", func() {
			cond := Gte(proto.Int32(0))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Gte(*int(nil))", func() {
			cond := Gte((*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Gte(*string)", func() {
			cond := Gte(proto.String("abc"))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " >= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Gte(*string(nil))", func() {
			cond := Gte((*string)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
	})
}

func Test_Lte(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Lte(nil)", func() {
			cond := Lte(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Lte(int)", func() {
			cond := Lte(0)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Lte(string)", func() {
			cond := Lte("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Lte(*int)", func() {
			cond := Lte(proto.Int32(0))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Lte(*int(nil))", func() {
			cond := Lte((*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Lte(*string)", func() {
			cond := Lte(proto.String("abc"))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " <= ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Lte(*string(nil))", func() {
			cond := Lte((*string)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
	})
}

func Test_Equal(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Equal(nil)", func() {
			cond := Equal(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})

		convey.Convey("Equal(int)", func() {
			cond := Equal(0)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{0}},
			})
		})
		convey.Convey("Equal(string)", func() {
			cond := Equal("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Equal(*int)", func() {
			cond := Equal(proto.Int32(0))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{int32(0)}},
			})
		})
		convey.Convey("Equal(*int(nil))", func() {
			cond := Equal((*int)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("Equal(*string)", func() {
			cond := Equal(proto.String("abc"))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " = ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Equal(*string(nil))", func() {
			cond := Equal((*string)(nil))
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
	})
}

func Test_In(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("In(nil)", func() {
			cond := In(nil)
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldBeNil)
		})
		convey.Convey("In(int)", func() {
			convey.So(func() { In(0) }, convey.ShouldPanic)
		})
		convey.Convey("In([3]int)", func() {
			cond := In([3]int{310, 320, 330})
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " IN (?)", Args: []interface{}{[3]int{310, 320, 330}}},
			})
		})
		convey.Convey("In([]int)", func() {
			cond := In([]int{310, 320, 330})
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " IN (?)", Args: []interface{}{[]int{310, 320, 330}}},
			})
		})
	})
}

func Test_Like(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		patches := ldhook.NewPatches()
		defer patches.Reset()

		var wheres []testGormWhere
		patches.Applys([]ldhook.Hook{
			ldhook.FuncHook{
				Target: (*gorm.DB).Where,
				Double: func(db *gorm.DB, query interface{}, args ...interface{}) *gorm.DB {
					wheres = append(wheres, testGormWhere{
						Query: query,
						Args:  args,
					})
					return db
				},
			},
		})

		convey.Convey("Like(abc)", func() {
			cond := Like("abc")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " LIKE ?", Args: []interface{}{"abc"}},
			})
		})
		convey.Convey("Like(%abc%)", func() {
			cond := Like("%abc%")
			cond.buildGorm(nil, field)
			convey.So(wheres, convey.ShouldResemble, []testGormWhere{
				{Query: field + " LIKE ?", Args: []interface{}{"%abc%"}},
			})
		})
	})
}
