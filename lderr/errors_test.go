/*
 * Copyright (C) distroy
 */

package lderr

import (
	"context"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	convey.Convey(t.Name(), t, func(c convey.C) {
		c.So(Is(ErrCtxCanceled, context.Canceled), convey.ShouldEqual, true)
		c.So(Is(ErrCtxDeadlineExceeded, context.DeadlineExceeded), convey.ShouldEqual, true)
	})
	tests := []struct {
		name   string
		err    error
		target error
		want   bool
	}{
		{
			name:   "context canceled",
			err:    ErrCtxCanceled,
			target: context.Canceled,
			want:   true,
		},
		{
			name:   "context deadline exceeded",
			err:    ErrCtxDeadlineExceeded,
			target: context.DeadlineExceeded,
			want:   true,
		},
		{
			name:   "unknown - overwrite",
			err:    ErrUnkown,
			target: Override(ErrUnkown, "abc"),
			want:   true,
		},
		{
			name:   "unknown - string error",
			err:    ErrUnkown,
			target: strError(errMessageUnkonw),
			want:   true,
		},
		{
			name:   "unknown - not is",
			err:    ErrUnkown,
			target: ErrCacheMarshal,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Is(tt.err, tt.target); got != tt.want {
				t.Errorf("Is() = %v, want %v", got, tt.want)
			}
		})
	}
}
