/*
 * Copyright (C) distroy
 */

package ldrbtree

import (
	"fmt"
	"testing"

	"github.com/distroy/ldgo/ldrand"
)

// BenchmarkRBTree_Insert/insert-size:100-8                  137914              7937 ns/op
// BenchmarkRBTree_Insert/insert-size:1000-8                  13520             86605 ns/op
// BenchmarkRBTree_Insert/insert-size:10000-8                   798           1265990 ns/op
// BenchmarkRBTree_Search/search-size:100-8                26130121                39.60 ns/op
// BenchmarkRBTree_Search/search-size:1000-8               19865358                52.53 ns/op
// BenchmarkRBTree_Search/search-size:10000-8              15433614                73.95 ns/op

func benchmarkRBTree_Insert(b *testing.B, size int) {
	b.Run(fmt.Sprintf("insert-size:%d", size), func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				rbtree := &RBTree{
					Compare: testCompareInt,
				}
				for i := 0; i < size; i++ {
					rbtree.Insert(ldrand.Int())
				}
				rbtree.Clear()
			}
		})
	})
}

func BenchmarkRBTree_Insert(b *testing.B) {
	benchmarkRBTree_Insert(b, 100)
	benchmarkRBTree_Insert(b, 1000)
	benchmarkRBTree_Insert(b, 10000)
}

func benchmarkRBTree_Search(b *testing.B, size int) {
	b.Run(fmt.Sprintf("search-size:%d", size), func(b *testing.B) {
		rbtree := &RBTree{
			Compare: testCompareInt,
		}
		for i := 0; i < size; i++ {
			rbtree.Insert(ldrand.Int())
		}
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				rbtree.Search(ldrand.Int())
			}
		})
		rbtree.Clear()
	})
}

func BenchmarkRBTree_Search(b *testing.B) {
	benchmarkRBTree_Search(b, 100)
	benchmarkRBTree_Search(b, 1000)
	benchmarkRBTree_Search(b, 10000)
}
