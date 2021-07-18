/*
 * Copyright (C) distroy
 */

package ldhook

import (
	"github.com/agiledragon/gomonkey"
)

type Patches interface {
	Apply(hooks ...Hook)
	Applys(hooks []Hook)
	Reset()
}

func NewPatches() Patches {
	return &patches{}
}

type patches struct {
	patches []*gomonkey.Patches
}

func (p *patches) Reset() {
	for i := len(p.patches) - 1; i >= 0; i-- {
		v := p.patches[i]
		v.Reset()
		p.patches[i] = nil
	}
	p.patches = p.patches[:0]
}

func (p *patches) Apply(hooks ...Hook) {
	p.Applys(hooks)
}

func (p *patches) Applys(hooks []Hook) {
	if len(hooks) == 0 {
		return
	}

	patch := gomonkey.NewPatches()
	for _, h := range hooks {
		h.hook(patch)
	}
	p.patches = append(p.patches, patch)
}
