/*
 * Copyright (C) distroy
 */

package ldrand

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/distroy/ldgo/ldcontext"
	"github.com/distroy/ldgo/ldsort"
	"github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap"
)

/*
 * pkg: github.com/distroy/ldgo/ldrand
 * cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
 * BenchmarkRandGo
 * BenchmarkRandGo-8       10440508               112.6 ns/op
 * BenchmarkRand
 * BenchmarkRand-8         43819591                27.69 ns/op
 */
func BenchmarkRandGo(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Intn(100)
		}
	})
}

func BenchmarkRand(b *testing.B) {
	// rand := newRandom()
	r := rand.New(NewFastSource(rand.Int63()))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r.Intn(100)
		}
	})
}

type testFastSource struct {
	Mod, Scale, Diff int
}

func (t *testFastSource) Test() {
	mod := t.Mod
	scale := t.Scale
	diff := t.Diff
	name := fmt.Sprintf("mod=%d,scale=%d,diff=%d", mod, scale, diff)
	convey.Convey(name, func() {
		r := rand.New(NewFastSource(time.Now().UnixNano()))

		counts := make([]int, mod)
		for i := 0; i < mod*scale; i++ {
			// x := r.Int() % mod
			x := r.Intn(mod)
			counts[x]++
		}

		ldsort.SortInts(counts)

		ctx := ldcontext.Console()
		ctx.LogI("", zap.Int("minCount", counts[0]), zap.Int("maxCount", counts[mod-1]))
		convey.So(counts[mod-1]-counts[0], convey.ShouldBeLessThan, diff)
	})
}

func Test_fastSource(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		r := rand.New(NewFastSource(time.Now().UnixNano()))

		(&testFastSource{
			Mod:   100,
			Scale: 100000,
			Diff:  2000,
		}).Test()
		(&testFastSource{
			Mod:   16,
			Scale: 100000,
			Diff:  100,
		}).Test()
		(&testFastSource{
			Mod:   256,
			Scale: 100000,
			Diff:  100,
		}).Test()

		convey.Convey("check the probability of very 4 bits", func() {
			const scale = 100000
			const diff = 200

			countsPer4Bits := [16][16]int{}
			for i := 0; i < scale*16; i++ {
				v := r.Uint64()
				for i := range countsPer4Bits {
					countsPer4Bits[i][v&0xf]++
					v = v >> 4
				}
			}

			ctx := ldcontext.Console()
			for i, v := range countsPer4Bits {
				ldsort.SortInts(v[:])
				ctx.LogI("", zap.Int("postion", i), zap.Int("minCount", v[0]), zap.Int("maxCount", v[15]))
				convey.So(v[15]-v[0], convey.ShouldBeLessThan, diff)
			}
		})

		convey.Convey("check the probability of very byte", func() {
			const scale = 100000
			const diff = 200

			countsPer4Bits := [8][256]int{}
			for i := 0; i < scale*256; i++ {
				v := r.Uint64()
				for i := range countsPer4Bits {
					countsPer4Bits[i][v&0xff]++
					v = v >> 8
				}
			}

			ctx := ldcontext.Console()
			for i, v := range countsPer4Bits {
				ldsort.SortInts(v[:])
				ctx.LogI("", zap.Int("postion", i), zap.Int("minCount", v[0]), zap.Int("maxCount", v[255]))
				convey.So(v[255]-v[0], convey.ShouldBeLessThan, diff)
			}
		})

		convey.Convey("check result if repeated", func() {
			const times = 100 * 10000

			m := make(map[uint64]struct{}, times)
			for i := 0; i < times; i++ {
				x := r.Uint64()
				if _, ok := m[x]; ok {
					t.Fatalf("number repeated. %d", x)
				}
				m[x] = struct{}{}
			}
		})
	})
}
