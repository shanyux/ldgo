/*
 * Copyright (C) distroy
 */

package ldgo

import (
	"math/rand"
	"testing"
	"time"
)

func testTopkInt(t testing.TB, n, k int) {
	rand.Seed(time.Now().UnixNano())
	origin := make([]int, 0, n)
	topk := make([]int, 0, k)

	for i := 0; i < n; i++ {
		x := rand.Intn(100)
		origin = append(origin, x)
		topk, _ = TopkIntsAdd(topk, k, x)
	}

	SortInts(origin)
	SortInts(topk)
	size := MinInt(n, k)
	for i := 0; i < size; i++ {
		if origin[i] != topk[i] {
			origin = origin[:size]
			t.Errorf("%s fail. n:%d, k:%d, output:%v, expected:%v", t.Name(), n, k, topk, origin)
			return
		}
	}
	t.Logf("%s succ. n:%d, k:%d, output:%v", t.Name(), n, k, topk)
}

func TestTopkInt(t *testing.T) {
	testTopkInt(t, 10, 20)
	testTopkInt(t, 100, 5)
	testTopkInt(t, 100, 10)
}
