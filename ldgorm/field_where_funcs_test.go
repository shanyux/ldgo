/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"google.golang.org/protobuf/proto"
)

func Test_Between(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("Between(nil, nil)", func() {
			cond := Between(nil, nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Between(int, nil)", func() {
			cond := Between(0, nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{0},
			})
		})
		convey.Convey("Between(string, nil)", func() {
			cond := Between("abc", nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Between(*int, nil)", func() {
			cond := Between(proto.Int32(0), nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Between(*int(nil), nil)", func() {
			cond := Between((*int)(nil), nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("Between(*string, nil)", func() {
			cond := Between(proto.String("abc"), nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Between(*string(nil), nil)", func() {
			cond := Between((*string)(nil), nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Between(nil, int)", func() {
			cond := Between(nil, 0)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{0},
			})
		})
		convey.Convey("Between(nil, string)", func() {
			cond := Between(nil, "abc")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Between(nil, *int)", func() {
			cond := Between(nil, proto.Int32(0))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Between(nil, *int(nil))", func() {
			cond := Between(nil, (*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("Between(nil, *string)", func() {
			cond := Between(nil, proto.String("abc"))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Between(nil, *string(nil))", func() {
			cond := Between(nil, (*string)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Between(int, int)", func() {
			cond := Between(0, 10)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " BETWEEN ? AND ?", Args: []interface{}{0, 10},
			})
		})
		convey.Convey("Between(int, int) && min == max", func() {
			cond := Between(3, 3)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " = ?", Args: []interface{}{3},
			})
		})
		convey.Convey("Between(*int, *int)", func() {
			cond := Between(proto.Int32(0), proto.Int32(10))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " BETWEEN ? AND ?", Args: []interface{}{int32(0), int32(10)},
			})
		})
		convey.Convey("Between(*int, *int) && min == max", func() {
			cond := Between(proto.Int32(3), proto.Int32(3))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " = ?", Args: []interface{}{int32(3)},
			})
		})
		convey.Convey("Between(*int(nil), *int)", func() {
			cond := Between((*int)(nil), proto.Int32(10))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{int32(10)},
			})
		})
		convey.Convey("Between(*int, *int(nil))", func() {
			cond := Between(proto.Int32(0), (*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Between(*int(nil), *int(nil))", func() {
			cond := Between((*int)(nil), (*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

	})
}

func Test_Gt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("Gt(nil)", func() {
			cond := Gt(nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Gt(int)", func() {
			cond := Gt(0)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " > ?", Args: []interface{}{0},
			})
		})
		convey.Convey("Gt(string)", func() {
			cond := Gt("abc")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " > ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Gt(*int)", func() {
			cond := Gt(proto.Int32(0))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " > ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Gt(*int(nil))", func() {
			cond := Gt((*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("Gt(*string)", func() {
			cond := Gt(proto.String("abc"))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " > ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Gt(*string(nil))", func() {
			cond := Gt((*string)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
	})
}

func Test_Lt(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("Lt(nil)", func() {
			cond := Lt(nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Lt(int)", func() {
			cond := Lt(0)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " < ?", Args: []interface{}{0},
			})
		})
		convey.Convey("Lt(string)", func() {
			cond := Lt("abc")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " < ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Lt(*int)", func() {
			cond := Lt(proto.Int32(0))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " < ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Lt(*int(nil))", func() {
			cond := Lt((*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("Lt(*string)", func() {
			cond := Lt(proto.String("abc"))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " < ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Lt(*string(nil))", func() {
			cond := Lt((*string)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
	})
}

func Test_Gte(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("Gte(nil)", func() {
			cond := Gte(nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Gte(int)", func() {
			cond := Gte(0)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{0},
			})
		})
		convey.Convey("Gte(string)", func() {
			cond := Gte("abc")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Gte(*int)", func() {
			cond := Gte(proto.Int32(0))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Gte(*int(nil))", func() {
			cond := Gte((*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("Gte(*string)", func() {
			cond := Gte(proto.String("abc"))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " >= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Gte(*string(nil))", func() {
			cond := Gte((*string)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
	})
}

func Test_Lte(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("Lte(nil)", func() {
			cond := Lte(nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Lte(int)", func() {
			cond := Lte(0)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{0},
			})
		})
		convey.Convey("Lte(string)", func() {
			cond := Lte("abc")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Lte(*int)", func() {
			cond := Lte(proto.Int32(0))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Lte(*int(nil))", func() {
			cond := Lte((*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("Lte(*string)", func() {
			cond := Lte(proto.String("abc"))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " <= ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Lte(*string(nil))", func() {
			cond := Lte((*string)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
	})
}

func Test_Equal(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("Equal(nil)", func() {
			cond := Equal(nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})

		convey.Convey("Equal(int)", func() {
			cond := Equal(0)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " = ?", Args: []interface{}{0},
			})
		})
		convey.Convey("Equal(string)", func() {
			cond := Equal("abc")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " = ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Equal(*int)", func() {
			cond := Equal(proto.Int32(0))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " = ?", Args: []interface{}{int32(0)},
			})
		})
		convey.Convey("Equal(*int(nil))", func() {
			cond := Equal((*int)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("Equal(*string)", func() {
			cond := Equal(proto.String("abc"))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " = ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Equal(*string(nil))", func() {
			cond := Equal((*string)(nil))
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
	})
}

func Test_In(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("In(nil)", func() {
			cond := In(nil)
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldBeZeroValue)
		})
		convey.Convey("In(int)", func() {
			convey.So(func() { In(0) }, convey.ShouldPanic)
		})
		convey.Convey("In([3]int)", func() {
			cond := In([3]int{310, 320, 330})
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " IN (?)", Args: []interface{}{[3]int{310, 320, 330}},
			})
		})
		convey.Convey("In([]int)", func() {
			cond := In([]int{310, 320, 330})
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " IN (?)", Args: []interface{}{[]int{310, 320, 330}},
			})
		})
	})
}

func Test_Like(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		const field = "field"

		convey.Convey("Like(abc)", func() {
			cond := Like("abc")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " LIKE ?", Args: []interface{}{"abc"},
			})
		})
		convey.Convey("Like(%abc%)", func() {
			cond := Like("%abc%")
			res := cond.buildWhere(field)
			convey.So(res, convey.ShouldResemble, whereResult{
				Query: field + " LIKE ?", Args: []interface{}{"%abc%"},
			})
		})
	})
}
