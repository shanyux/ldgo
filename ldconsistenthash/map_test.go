/*
 * Copyright (C) distroy
 */

package ldconsistenthash

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/distroy/ldgo/ldmath"
	"github.com/distroy/ldgo/ldptr"
	"github.com/distroy/ldgo/ldrand"
)

/*
 * cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
 * BenchmarkMap
 * BenchmarkMap/init-nodes:200-replicas:200
 * BenchmarkMap/init-nodes:200-replicas:200-8                   386           2897380 ns/op
 * BenchmarkMap/run-nodes:200-replicas:200
 * BenchmarkMap/run-nodes:200-replicas:200-8                1631971               737.2 ns/op
 */
func BenchmarkMap(b *testing.B) {
	nodeCount := 200
	replicas := 200
	benchmarkMap(b, nodeCount, replicas)
}

func benchmarkMap(b *testing.B, nodeCount, replicas int) {
	nodes := make([]string, 0, nodeCount)
	for i := 0; i < nodeCount; i++ {
		nodes = append(nodes, ldrand.String(32))
	}

	b.Run(fmt.Sprintf("init-nodes:%d-replicas:%d", nodeCount, replicas), func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				chash := New(WithReplicas(replicas))
				chash.Adds(nodes...)
			}
		})
	})

	begin := time.Now()
	chash := New(WithReplicas(replicas))
	chash.Adds(nodes...)
	cost := time.Since(begin)

	totalHits := int32(0)
	nodeHits := make(map[string]*int32, nodeCount)
	for _, node := range nodes {
		nodeHits[node] = ldptr.NewInt32(0)
	}

	b.Run(fmt.Sprintf("run-nodes:%d-replicas:%d", nodeCount, replicas), func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				key := ldrand.String(32)
				node := chash.Get(key)
				atomic.AddInt32(nodeHits[node], 1)
				atomic.AddInt32(&totalHits, 1)
			}
		})
	})

	min, max := testMinAndMax(nodeHits)
	b.Logf("hits array: %v", testHits(nodeHits))
	b.Logf("node count:%d, replicas:%d, init duration:%s", nodeCount, replicas, cost.String())
	b.Logf("min hits:%d, max hits:%d, arg hits:%d", min, max, int(totalHits)/nodeCount)
}

func testMinAndMax(m map[string]*int32) (int32, int32) {
	a := testHits(m)
	return ldmath.MinInt32(a[0], a...), ldmath.MaxInt32(a[0], a...)
}

func testHits(m map[string]*int32) []int32 {
	a := make([]int32, 0, len(m))
	for _, v := range m {
		a = append(a, *v)
	}
	return a
}
