/*
 * Copyright (C) distroy
 */

package ldsort

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestSearch(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		arr := []int{7, 11, 20, 25, 38, 50, 55, 76, 76, 84}
		convey.So(Search(len(arr), func(i int) bool { return arr[i] >= 19 }), convey.ShouldEqual, 2)
		convey.So(Search(len(arr), func(i int) bool { return arr[i] >= 20 }), convey.ShouldEqual, 2)
		convey.So(Search(len(arr), func(i int) bool { return arr[i] >= 21 }), convey.ShouldEqual, 3)
	})
}

func TestIndex(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		arr := []int{7, 11, 20, 25, 38, 50, 55, 76, 76, 84}
		convey.So(Index(len(arr), func(i int) int { return arr[i] - 19 }), convey.ShouldEqual, -1)
		convey.So(Index(len(arr), func(i int) int { return arr[i] - 20 }), convey.ShouldEqual, 2)
		convey.So(Index(len(arr), func(i int) int { return arr[i] - 21 }), convey.ShouldEqual, -1)
	})
}
