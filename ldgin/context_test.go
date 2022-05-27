/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func TestGetContext(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		gin.SetMode(gin.TestMode)
		g, _ := gin.CreateTestContext(httptest.NewRecorder())

		c := GetContext(g)
		convey.So(c, convey.ShouldNotBeNil)
		convey.So(c, convey.ShouldEqual, GetContext(g))
	})
}

func TestGetGin(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		gin.SetMode(gin.TestMode)
		g, _ := gin.CreateTestContext(httptest.NewRecorder())

		c := GetContext(g)
		convey.So(c, convey.ShouldNotBeNil)
		convey.So(g, convey.ShouldEqual, GetGin(g))
		convey.So(g, convey.ShouldEqual, GetGin(c))
	})
}
