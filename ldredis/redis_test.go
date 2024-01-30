package ldredis

import (
	"testing"

	"github.com/distroy/ldgo/v2/ldhook"
	redis "github.com/redis/go-redis/v9"
	"github.com/smartystreets/goconvey/convey"
)

func Test_New(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		patches := ldhook.NewPatches()
		defer patches.Reset()

		convey.Convey("redis cluster", func() {
			patches.Applys([]ldhook.Hook{
				ldhook.FuncHook{
					Target: redis.NewClusterClient,
					Double: ldhook.Values{&redis.ClusterClient{}},
				},
			})
			cfg := &Config{
				Cluster:  true,
				Addrs:    []string{"1.1.1.1"},
				Password: "now password",
			}
			convey.So(NewByConfig(cfg), convey.ShouldNotBeNil)
		})

		convey.Convey("redis not cluster", func() {
			patches.Applys([]ldhook.Hook{
				ldhook.FuncHook{
					Target: redis.NewClient,
					Double: ldhook.Values{&redis.Client{}},
				},
			})
			cfg := &Config{
				Cluster:  false,
				Addrs:    []string{"1.1.1.1"},
				Addr:     "1.1.1.1",
				Password: "now password",
			}
			convey.So(NewByConfig(cfg), convey.ShouldNotBeNil)
		})
	})
}
