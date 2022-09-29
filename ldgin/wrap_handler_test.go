/*
 * Copyright (C) distroy
 */

package ldgin

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/distroy/ldgo/lderr"
	"github.com/distroy/ldgo/ldio"
	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

func TestWrapHandler(t *testing.T) {

	log.SetOutput(ldio.Discard())

	body := `{ "Where": "abc", "channel_id": 123 }`

	convey.Convey(t.Name(), t, func() {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		g, _ := gin.CreateTestContext(w)

		g.Params = append(g.Params, gin.Param{Key: "project_id", Value: "101"})
		g.Params = append(g.Params, gin.Param{Key: "channel_id", Value: "201"})
		g.Request = httptest.NewRequest("GET", "http://github.com/?page=301", strings.NewReader(body))

		var request interface{}
		response := &CommResponse{}

		convey.Convey("func (): panic", func() {
			convey.So(func() { WrapHandler(func() {}) }, convey.ShouldPanic)
		})

		convey.Convey("func (Request): panic", func() {
			convey.So(func() { WrapHandler(func(*testRequest) {}) }, convey.ShouldPanic)
		})

		convey.Convey("func (*gin.Context) Response", func() {
			convey.So(func() { WrapHandler(func(*gin.Context) *testResponse { return nil }) }, convey.ShouldPanic)
		})

		// convey.Convey("Request --> ParseValidator", func() {})
		// convey.Convey("Request --> Parser", func() {})
		// convey.Convey("Request --> Validator", func() {})
		// convey.Convey("Request --> GinParseValidator", func() {})
		// convey.Convey("Request --> GinParser", func() {})
		// convey.Convey("Request --> GinValidator", func() {})
		// convey.Convey("Request --> interface{}", func() {})

		convey.Convey("func (*gin.Context)", func() {
			handler := WrapHandler(func(g *gin.Context) {
				c := GetContext(g)
				c.AbortWithError(lderr.ErrUnauthorized)
			})

			handler(g)
			convey.So(GetError(g), convey.ShouldEqual, lderr.ErrUnauthorized)
			convey.So(w.Code, convey.ShouldEqual, lderr.ErrUnauthorized.Status())

			rsp := response
			convey.So(json.Unmarshal(w.Body.Bytes(), rsp), convey.ShouldBeNil)
			convey.So(rsp.ErrCode, convey.ShouldEqual, lderr.ErrUnauthorized.Code())
		})

		convey.Convey("func (*gin.Context) error", func() {
			handler := WrapHandler(func(g *gin.Context) error {
				return lderr.ErrHttpInvalidStatus
			})

			handler(g)

			convey.So(w.Code, convey.ShouldEqual, lderr.ErrHttpInvalidStatus.Status())
			convey.So(GetError(g), convey.ShouldEqual, lderr.ErrHttpInvalidStatus)

			// rsp := GetResponse(g)
			// convey.So(rsp, convey.ShouldBeNil)
		})

		convey.Convey("func (*gin.Context, Request) Error", func() {
			convey.Convey("Request --> interface{}", func() {
				handler := WrapHandler(func(g *gin.Context, req *testRequest) Error {
					request = req
					return nil
				})

				handler(g)

				req := request
				convey.So(req, convey.ShouldResemble, &testRequest{
					ProjectId: 101,
					ChannelId: 201,
					Page:      301,
					Where:     "abc",
				})

				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(GetError(g), convey.ShouldBeNil)

				rsp := GetResponse(g)
				convey.So(rsp, convey.ShouldNotBeNil)
				convey.So(rsp.ErrCode, convey.ShouldEqual, 0)
			})
		})

		convey.Convey("func (*gin.Context) (Response, Error)", func() {
			convey.Convey("Response --> Renderer", func() {
				handler := WrapHandler(func(g *gin.Context) (*testRenderer, Error) {
					return &testRenderer{str: "succ"}, nil
				})

				handler(g)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(GetError(g), convey.ShouldBeNil)

				rsp := GetResponse(g)
				convey.So(rsp, convey.ShouldBeNil)
				convey.So(w.Body.String(), convey.ShouldEqual, "succ")

				convey.So(GetRenderer(g), convey.ShouldResemble, &testRenderer{str: "succ"})
			})
		})

		convey.Convey("func (*gin.Context, Request) (Response, Error)", func() {
			convey.Convey("Request --> GinParser: fail", func() {
				convey.Convey("Response --> interface{}", func() {
					handler := WrapHandler(func(g *gin.Context, req *testGinParserFail) (*testResponse, Error) {
						return &testResponse{
							UserId: 301,
							ShopId: 401,
						}, nil
					})

					handler(g)

					convey.So(w.Code, convey.ShouldEqual, lderr.ErrHttpReadBody.Status())
					convey.So(GetError(g), convey.ShouldEqual, lderr.ErrHttpReadBody)
					// convey.So(GetResponse(g), convey.ShouldBeNil)
				})
				convey.Convey("Request --> GinParser: succ", func() {
					convey.Convey("Response --> interface{}", func() {
						handler := WrapHandler(func(g *gin.Context, req *testGinParser) (*testResponse, Error) {
							request = req
							return &testResponse{
								UserId: 301,
								ShopId: 401,
							}, nil
						})

						handler(g)

						req := request
						convey.So(req, convey.ShouldResemble, &testGinParser{
							ProjectId: 101,
							ChannelId: 201,
							Page:      301,
							Where:     "abc",
						})

						convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
						convey.So(GetError(g), convey.ShouldBeNil)

						rsp := GetResponse(g)
						convey.So(rsp, convey.ShouldNotBeNil)
						convey.So(rsp.ErrCode, convey.ShouldEqual, 0)
						convey.So(rsp.Data, convey.ShouldResemble, &testResponse{
							UserId: 301,
							ShopId: 401,
						})
					})
				})
			})
		})
		convey.Convey("func (*Context): panic", func() {
			handler := WrapHandler(func(c *Context) {
				panic(lderr.ErrServicePanic)
			})

			handler(g)
			convey.So(GetError(g), convey.ShouldEqual, lderr.ErrServicePanic)
			convey.So(w.Code, convey.ShouldEqual, lderr.ErrServicePanic.Status())

			rsp := response
			convey.So(json.Unmarshal(w.Body.Bytes(), rsp), convey.ShouldBeNil)
			convey.So(rsp.ErrCode, convey.ShouldEqual, lderr.ErrServicePanic.Code())
		})

		convey.Convey("func (*Context) Error", func() {
			handler := WrapHandler(func(c *Context) Error {
				return lderr.ErrHttpInvalidStatus
			})

			handler(g)

			convey.So(w.Code, convey.ShouldEqual, lderr.ErrHttpInvalidStatus.Status())
			convey.So(GetError(g), convey.ShouldEqual, lderr.ErrHttpInvalidStatus)

			// rsp := GetResponse(g)
			// convey.So(rsp, convey.ShouldBeNil)
		})

		convey.Convey("func (*Context, Request) Error", func() {
			convey.Convey("Request --> GinValidator: fail", func() {
				handler := WrapHandler(func(g *gin.Context, req *testGinValidatorFail) Error {
					return nil
				})

				handler(g)

				convey.So(w.Code, convey.ShouldEqual, lderr.ErrHttpReadBody.Status())
				convey.So(GetError(g), convey.ShouldEqual, lderr.ErrHttpReadBody)

				// convey.So(GetResponse(g), convey.ShouldBeNil)
			})
			convey.Convey("Request --> GinValidator: succ", func() {
				handler := WrapHandler(func(g *gin.Context, req *testGinValidator) Error {
					request = req
					return nil
				})

				handler(g)

				req := request
				convey.So(req, convey.ShouldResemble, &testGinValidator{
					ProjectId: 101,
					ChannelId: 201,
					Page:      301,
					Where:     "abc",
				})

				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(GetError(g), convey.ShouldBeNil)
			})
		})
		convey.Convey("func (*Context) (Response, Error)", func() {
			convey.Convey("Response --> GinRenderer", func() {
				handler := WrapHandler(func(c *Context) (*testGinRenderer, Error) {
					return &testGinRenderer{str: "succ"}, nil
				})

				handler(g)
				convey.So(w.Code, convey.ShouldEqual, http.StatusOK)
				convey.So(GetError(g), convey.ShouldBeNil)

				rsp := GetResponse(g)
				convey.So(rsp, convey.ShouldBeNil)
				convey.So(w.Body.String(), convey.ShouldEqual, "succ")

				convey.So(GetRenderer(g), convey.ShouldResemble, &testGinRenderer{str: "succ"})
			})
		})
		convey.Convey("func (*Context, Request) (Response, Error)", func() {
			convey.Convey("Request --> Parser: fail", func() {
				convey.Convey("Response --> interface{}", func() {
					handler := WrapHandler(func(c *Context, req *testParserFail) (*testResponse, Error) {
						return &testResponse{
							UserId: 301,
							ShopId: 401,
						}, nil
					})

					handler(g)

					convey.So(w.Code, convey.ShouldEqual, lderr.ErrHttpReadBody.Status())
					convey.So(GetError(g), convey.ShouldEqual, lderr.ErrHttpReadBody)
				})
			})
			convey.Convey("Request --> ParseValidator: succ", func() {
				convey.Convey("Response --> interface{}", func() {
					handler := WrapHandler(func(c *Context, req *testParseValidator) (*testResponse, Error) {
						request = req
						return &testResponse{
							UserId: 301,
							ShopId: 401,
						}, nil
					})

					handler(g)

					req := request
					convey.So(req, convey.ShouldResemble, &testParseValidator{
						ProjectId: 101,
						ChannelId: 201,
						Page:      301,
						Where:     "abc",
					})

					rsp := GetResponse(g)
					convey.So(rsp, convey.ShouldNotBeNil)
					convey.So(rsp.ErrCode, convey.ShouldEqual, 0)
					convey.So(rsp.Data, convey.ShouldResemble, &testResponse{
						UserId: 301,
						ShopId: 401,
					})
				})
			})
		})
	})
}
