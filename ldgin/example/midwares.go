/*
 * Copyright (C) distroy
 */

package main

import (
	"net/http"

	"github.com/distroy/ldgo/v2/lderr"
	"github.com/distroy/ldgo/v2/ldgin"
)

func midware1(c *ldgin.Context) {
	c.LogI("midware1")
}

func midware2(c *ldgin.Context) ldgin.Error {
	c.LogI("midware2")
	return nil
}

func midware3(c *ldgin.Context) ldgin.Error {
	c.LogI("midware3")
	return lderr.New(http.StatusOK, 120, "midware error")
}
