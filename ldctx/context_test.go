/*
 * Copyright (C) distroy
 */

package ldctx

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.Convey("nil", func(c convey.C) {
			ctx := New(nil)
			c.So(ctx, convey.ShouldResemble, Default())
			// c.So(ctx, convey.ShouldEqual, Default())
		})

		c.Convey("background", func(c convey.C) {
			ctx := New(context.Background())
			c.So(ctx, convey.ShouldResemble, Default())
			// c.So(ctx, convey.ShouldEqual, Default())
		})

		c.Convey("default", func(c convey.C) {
			ctx := New(Default())
			c.So(ctx, convey.ShouldResemble, Default())
			// c.So(ctx, convey.ShouldEqual, Default())
		})

		c.Convey("console", func(c convey.C) {
			ctx := New(Console())
			c.So(ctx, convey.ShouldResemble, Console())
			// c.So(ctx, convey.ShouldEqual, Console())
		})

		c.Convey("background with fields", func(c convey.C) {
			tmp := context.Background()
			tmp = context.WithValue(tmp, 1, 1)
			cc := New(tmp, zap.String("abc", "abc"))

			c.So(cc, convey.ShouldNotResemble, Default())
			c.So(cc, convey.ShouldHaveSameTypeAs, ctx{})
			c.So(cc.(ctx).Context, convey.ShouldHaveSameTypeAs, logCtx{})
			c.So(cc.(ctx).Context.(logCtx).Context, convey.ShouldNotResemble, context.Background())
			c.So(cc.Value(1), convey.ShouldResemble, 1)
		})

		c.Convey("default with fields", func(c convey.C) {
			tmp := Default()
			tmp = WithValue(tmp, 1, 1)
			cc := New(tmp, zap.String("abc", "abc"))

			c.So(cc, convey.ShouldNotResemble, Default())
			c.So(cc, convey.ShouldHaveSameTypeAs, ctx{})
			c.So(cc.(ctx).Context, convey.ShouldHaveSameTypeAs, logCtx{})
			c.So(cc.(ctx).Context.(logCtx).Context, convey.ShouldNotResemble, context.Background())
			c.So(cc.Value(1), convey.ShouldResemble, 1)
		})
	})
}
