/*
 * Copyright (C) distroy
 */

package ldflag

import (
	"reflect"
	"strings"
	"testing"

	"github.com/distroy/ldgo/v2/ldptr"
	"github.com/smartystreets/goconvey/convey"
)

const (
	testCommandName = "unittest"
)

type testIncludes []string

func (p *testIncludes) Default() string { return strings.Join(nil, "\n") }
func (p *testIncludes) String() string  { return mustMarshalJson(*p) }
func (p *testIncludes) Set(s string) error {
	*p = append(*p, s)
	return nil
}

type testExcludes []string

func (p *testExcludes) Default() string { return strings.Join([]string{"a", "b"}, "\n") }
func (p *testExcludes) String() string  { return mustMarshalJson(*p) }
func (p *testExcludes) Set(s string) error {
	*p = append(*p, s)
	return nil
}

func newTestFlagSet() *FlagSet {
	s := NewFlagSet()
	s.init()
	s.name = testCommandName
	return s
}

func TestFlagSet_wrigeUsageHeader(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		convey.Convey(`name == "" && no args`, func() {
			s := &FlagSet{
				name: "",
				args: nil,
			}
			b := &strings.Builder{}
			s.writeUsageHeader(b)
			convey.So(b.String(), convey.ShouldEqual, "Usage of <command>:\nFlags:\n")
		})
		convey.Convey(`name == "abc" && args.meta == ""`, func() {
			s := &FlagSet{
				name: "abc",
				args: &Flag{Meta: ""},
			}
			b := &strings.Builder{}
			s.writeUsageHeader(b)
			convey.So(b.String(), convey.ShouldEqual, "Usage: abc [<flags>] [<arg>...]\n\nFlags:\n")
		})
	})
}

func TestFlagSet_wrigeUsage(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		b := &strings.Builder{}
		s := newTestFlagSet()

		convey.Convey("normal value", func() {
			type Flags struct {
				Top      int      `flag:"name:top; meta:N; usage:show the top <N>"`
				Avg      bool     `flag:"usage:show the average complexity"`
				DebugLog bool     `flag:"usage:print debug log; bool"`
				Rate     float64  `flag:"default:0.65; usage:"`
				Branch   string   `flag:"meta:branch; usage:git branch name"`
				Includes []string `flag:"name:include; meta:regexp; usage:include file regexps"`
				Excludes []string `flag:"name:exclude; meta:regexp; usage:exclude file regexps"`
				Pathes   []string `flag:"args; meta:path; default:."`
			}

			flags := &Flags{
				Excludes: []string{
					`(^|/)vendor/`,
					`\.pb\.go$`,
				},
			}
			s.Model(flags)
			s.writeUsage(b)

			convey.So(b.String(), convey.ShouldEqual, `Usage: unittest [<flags>] [path...]

Flags:
        -top <N>
                show the top <N>
        -avg <bool>
                show the average complexity
        -debug-log
                print debug log
        -rate <float> (default: 0.65)
        -branch <branch>
                git branch name
        -include <regexp>
                include file regexps
        -exclude <regexp>
                exclude file regexps
                default:
                        "(^|/)vendor/"
                        "\\.pb\\.go$"
`)
		})

		convey.Convey("pointer value", func() {
			type Flags struct {
				Top      *int     `flag:"name:top; meta:N; usage:show the top <N>"`
				Avg      *bool    `flag:"usage:show the average complexity"`
				DebugLog *bool    `flag:"usage:print debug log; bool"`
				Rate     *float64 `flag:"default:0.65; usage:"`
				Branch   *string  `flag:"meta:branch; usage:git branch name"`
				Includes []string `flag:"name:include; meta:regexp; usage:include file regexps"`
				Excludes []string `flag:"name:exclude; meta:regexp; usage:exclude file regexps"`
				Pathes   []string `flag:"args; meta:path; default:."`
			}

			flags := &Flags{
				Excludes: []string{
					`(^|/)vendor/`,
					`\.pb\.go$`,
				},
			}
			s.Model(flags)
			s.writeUsage(b)

			convey.So(b.String(), convey.ShouldEqual, `Usage: unittest [<flags>] [path...]

Flags:
        -top <N>
                show the top <N>
        -avg <bool>
                show the average complexity
        -debug-log
                print debug log
        -rate <float> (default: 0.65)
        -branch <branch>
                git branch name
        -include <regexp>
                include file regexps
        -exclude <regexp>
                exclude file regexps
                default:
                        "(^|/)vendor/"
                        "\\.pb\\.go$"
`)
		})
		convey.Convey("value with default", func() {
			type Flags struct {
				Includes testIncludes `flag:"name:include; meta:regexp; usage:include file regexps"`
				Excludes testExcludes `flag:"name:exclude; meta:regexp; usage:exclude file regexps"`
			}

			flags := &Flags{}
			s.Model(flags)
			s.writeUsage(b)

			convey.So(b.String(), convey.ShouldEqual, `Usage of unittest:
Flags:
        -include <regexp>
                include file regexps
        -exclude <regexp>
                exclude file regexps
                default:
                        a
                        b
`)
		})

		c.Convey("short flag && no usage", func(c convey.C) {
			type Flags struct {
				Csv bool `flag:"bool"`
			}

			flags := &Flags{}

			s.Model(flags)
			s.writeUsage(b)

			convey.So(b.String(), convey.ShouldEqual, `Usage of unittest:
Flags:
        -csv
`)
		})
	})
}

func TestFlagSet_Model(t *testing.T) {
	testClearFlagSet := func(s *FlagSet) {
		for _, v := range s.flagSlice {
			v.tags = nil
			v.val = reflect.Value{}
			v.Value = nil
		}
		if args := s.args; args != nil {
			args.tags = nil
			args.val = reflect.Value{}
			args.Value = nil
		}
	}

	type Flags struct {
		Top      int      `flag:"name:top; meta:N; usage:show the top <N>"`
		Avg      bool     `flag:"usage:show the average complexity"`
		DebugLog bool     `flag:"usage:print debug log; bool"`
		Rate     float64  `flag:"default:0.65; usage:"`
		Branch   string   `flag:"meta:branch; usage:git branch name"`
		Pathes   []string `flag:"args; meta:path; default:."`
	}

	convey.Convey(t.Name(), t, func() {
		convey.Convey("", func() {
			flags := &Flags{}

			s := NewFlagSet()
			s.Model(flags)
			testClearFlagSet(s)

			convey.So(s.flagSlice, convey.ShouldResemble, []*Flag{
				{
					lvl:     0,
					Name:    "top",
					Meta:    "N",
					Default: "0",
					Usage:   "show the top <N>",
					IsArgs:  false,
					Bool:    false,
				},
				{
					lvl:     0,
					Name:    "avg",
					Meta:    "",
					Default: "false",
					Usage:   "show the average complexity",
					IsArgs:  false,
					Bool:    false,
				},
				{
					lvl:     0,
					Name:    "debug-log",
					Meta:    "",
					Default: "false",
					Usage:   "print debug log",
					IsArgs:  false,
					Bool:    true,
				},
				{
					lvl:     0,
					Name:    "rate",
					Meta:    "",
					Default: "0.65",
					Usage:   "",
					IsArgs:  false,
					Bool:    false,
				},
				{
					lvl:     0,
					Name:    "branch",
					Meta:    "branch",
					Default: "",
					Usage:   "git branch name",
					IsArgs:  false,
					Bool:    false,
				},
			})
			convey.So(s.args, convey.ShouldResemble, &Flag{
				lvl:     0,
				Name:    "pathes",
				Meta:    "path",
				Default: ".",
				Usage:   "",
				IsArgs:  true,
				Bool:    false,
			})
		})
	})
}

func TestFlagSet_Parse(t *testing.T) {
	convey.Convey(t.Name(), t, func() {
		s := newTestFlagSet()

		convey.Convey("normal value", func() {
			type Flags struct {
				Top      int      `flag:"name:top; meta:N; usage:show the top <N>"`
				Avg      bool     `flag:"usage:show the average complexity"`
				DebugLog bool     `flag:"usage:print debug log; bool"`
				Rate     float64  `flag:"default:0.65; usage:"`
				Branch   string   `flag:"meta:branch; usage:git branch name"`
				Pathes   []string `flag:"args; meta:path; default:."`
			}

			convey.Convey("no set default", func() {
				flags := &Flags{}
				s.EnableDefault(false)
				s.Model(flags)

				err := s.Parse([]string{
					"-top", "5",
					"-avg", "1",
					"-debug-log",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(flags, convey.ShouldResemble, &Flags{
					Top:      5,
					Avg:      true,
					DebugLog: true,
					Rate:     0,
					Branch:   "",
					Pathes:   []string{},
				})
			})

			convey.Convey("set default", func() {
				flags := &Flags{}
				s.EnableDefault(true)
				s.Model(flags)

				err := s.Parse([]string{
					"-top", "5",
					"-avg", "1",
					"-debug-log",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(flags, convey.ShouldResemble, &Flags{
					Top:      5,
					Avg:      true,
					DebugLog: true,
					Rate:     0.65,
					Branch:   "",
					Pathes:   []string{"."},
				})
			})
		})

		convey.Convey("ptr value", func() {
			type Flags struct {
				Over     *int     `flag:"name:over; meta:N; default:15; usage:show functions with complexity <N>"`
				Top      *int     `flag:"name:top; meta:N; usage:show the top <N>"`
				Avg      *bool    `flag:"usage:show the average complexity"`
				DebugLog *bool    `flag:"usage:print debug log; bool"`
				Rate     *float64 `flag:"default:0.65; usage:"`
				Branch   *string  `flag:"meta:branch; usage:git branch name"`
				Pathes   []string `flag:"args; meta:path; default:."`
			}

			convey.Convey("no set default", func() {
				flags := &Flags{}
				s.EnableDefault(false)
				s.Model(flags)

				err := s.Parse([]string{
					"-top", "5",
					"-avg", "1",
					"-debug-log",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(flags, convey.ShouldResemble, &Flags{
					Over:     nil,
					Top:      ldptr.NewInt(5),
					Avg:      ldptr.NewBool(true),
					DebugLog: ldptr.NewBool(true),
					Rate:     nil,
					Branch:   nil,
					Pathes:   []string{},
				})
			})

			convey.Convey("set default", func() {
				flags := &Flags{}
				s.EnableDefault(true)
				s.Model(flags)

				err := s.Parse([]string{
					"-top", "5",
					"-avg", "1",
					"-debug-log",
				})
				convey.So(err, convey.ShouldBeNil)
				convey.So(flags, convey.ShouldResemble, &Flags{
					Over:     ldptr.NewInt(15),
					Top:      ldptr.NewInt(5),
					Avg:      ldptr.NewBool(true),
					DebugLog: ldptr.NewBool(true),
					Rate:     ldptr.NewFloat64(0.65),
					Branch:   nil,
					Pathes:   []string{"."},
				})
			})
		})
	})
}
