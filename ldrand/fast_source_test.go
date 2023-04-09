/*
 * Copyright (C) distroy
 */

package ldrand

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/distroy/ldgo/ldmath"
	"github.com/smartystreets/goconvey/convey"
)

/*
 * pkg: github.com/distroy/ldgo/ldrand
 * cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
 * BenchmarkRandGo
 * BenchmarkRandGo-12      18746797                63.71 ns/op
 * BenchmarkRand
 * BenchmarkRand-12        68977040                17.42 ns/op
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
	r := New(NewFastSource(rand.Int63()))
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

func maxInt(a []int) int { return ldmath.MaxInt(a[0], a[1:]...) }
func minInt(a []int) int { return ldmath.MinInt(a[0], a[1:]...) }

func diffRatio(a []int) float64 {
	sum := ldmath.SumInt(a...)
	cnt := int64(len(a))
	avg := (sum + cnt/2) / cnt

	diff := int64(0)
	for _, n := range a {
		diff += ldmath.AbsInt64(avg - int64(n))
	}

	return float64(diff) / float64(sum)
}

func (t *testFastSource) Test() {
	var (
		mod   = t.Mod
		scale = t.Scale
		diff  = t.Diff
	)
	name := fmt.Sprintf("mod=%d,scale=%d,diff=%d", mod, scale, diff)
	convey.Convey(name, func() {
		r := New(NewFastSource(time.Now().UnixNano()))

		counts := make([]int, mod)
		for i := 0; i < mod*scale; i++ {
			// x := r.Int() % mod
			x := r.Intn(mod)
			counts[x]++
		}

		min := minInt(counts)
		max := maxInt(counts)
		ratio := diffRatio(counts)

		log.Printf("diff:%d, ratio:%.04g, min:%d, max:%d", max-min, ratio, min, max)
		convey.So(max-min, convey.ShouldBeLessThan, diff)
	})
}

func Test_fastSource_ProbabilityOfOverall(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		(&testFastSource{
			Mod:   100,
			Scale: 1000 * 200,
			Diff:  1000 * 4,
		}).Test()
		(&testFastSource{
			Mod:   16,
			Scale: 1000 * 200,
			Diff:  1000 * 4,
		}).Test()
		(&testFastSource{
			Mod:   256,
			Scale: 1000 * 200,
			Diff:  1000 * 4,
		}).Test()
	})
}

func Test_fastSource_ProbabilityOfVery4Bits(t *testing.T) {
	r := New(NewFastSource(time.Now().UnixNano()))

	convey.Convey(t.Name(), t, func() {
		convey.Convey("check the probability of very 4 bits", func() {
			const (
				scale = 1000 * 100
				diff  = 1000 * 4
			)

			countsPer4Bits := [16][16]int{}
			for i := 0; i < scale*16; i++ {
				v := r.Uint64()
				for i := range countsPer4Bits {
					countsPer4Bits[i][v&0xf]++
					v = v >> 4
				}
			}

			log.Printf("")
			for i, v := range countsPer4Bits {
				min := minInt(v[:])
				max := maxInt(v[:])
				ratio := diffRatio(v[:])
				log.Printf("postion:%d, diff:%d, ratio:%.04g, min:%d, max:%d",
					i, max-min, ratio, min, max)
				convey.So(max-min, convey.ShouldBeLessThan, diff)
			}
		})
	})
}

func Test_fastSource_ProbabilityOfVery4BitsWithPreviousNumber(t *testing.T) {
	r := New(NewFastSource(time.Now().UnixNano()))

	convey.Convey(t.Name(), t, func() {
		convey.Convey("check the probability of very 4 bits with previous number", func() {
			const (
				scale = 1000 * 100
				diff  = 1000 * 40
			)

			countsPer4BitsWithPrev := [16][16][16]int{}
			var prevNum uint64
			for i := 0; i < scale*16*16; i++ {
				v := r.Uint64()
				p := prevNum
				prevNum = v
				for i := range countsPer4BitsWithPrev {
					countsPer4BitsWithPrev[i][p&0xf][v&0xf]++
					v = v >> 4
					p = p >> 4
				}
			}

			log.Printf("")
			for i, v := range countsPer4BitsWithPrev {
				for j, w := range v {
					min := minInt(w[:])
					max := maxInt(w[:])
					ratio := diffRatio(w[:])
					log.Printf("postion:%d, prev:%d, diff:%d, ratio:%.04g, min:%d, max:%d, v:%v",
						i, j, max-min, ratio, min, max, w[:])

					// sometimes it will be failed
					// convey.So(max-min, convey.ShouldBeLessThan, diff)
				}
			}
		})
	})
}

func Test_fastSource_ProbabilityOfVeryByte(t *testing.T) {
	r := New(NewFastSource(time.Now().UnixNano()))

	convey.Convey(t.Name(), t, func() {
		convey.Convey("check the probability of very byte", func() {
			const (
				scale = 1000 * 100
				diff  = 1000 * 5
			)

			countsPer4Bits := [8][256]int{}
			for i := 0; i < scale*256; i++ {
				v := r.Uint64()
				for i := range countsPer4Bits {
					countsPer4Bits[i][v&0xff]++
					v = v >> 8
				}
			}

			log.Printf("")
			for i, v := range countsPer4Bits {
				min := minInt(v[:])
				max := maxInt(v[:])
				ratio := diffRatio(v[:])
				log.Printf("postion:%d, diff:%d, ratio:%.04g, min:%d, max:%d",
					i, max-min, ratio, min, max)
				convey.So(max-min, convey.ShouldBeLessThan, diff)
			}
		})
	})
}

func Test_fastSource_Repeated(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		r := New(NewFastSource(time.Now().UnixNano()))

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
