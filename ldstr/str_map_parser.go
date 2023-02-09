/*
 * Copyright (C) distroy
 */

package ldstr

import (
	"strings"
	"sync"

	"github.com/distroy/ldgo/lderr"
)

var (
	strMapFieldsPool = &sync.Pool{
		New: func() interface{} {
			return make([]strMapField, 0, 8)
		},
	}

	strMapParseContextPool = &sync.Pool{
		New: func() interface{} {
			return &strMapParseContext{
				Fields: make([]strMapFieldContext, 0, 8),
			}
		},
	}
)

type strMapField struct {
	Name  string
	After string
}

type strMapFieldContext struct {
	Prev  *strMapFieldContext
	Next  *strMapFieldContext
	Field *strMapField

	Text      string
	Value     string
	LastIndex int
	Duplicate bool
}

type strMapParseContext struct {
	Fields []strMapFieldContext
}

type strMapParser struct {
	template    string
	left, right string

	before string
	fields []strMapField
}

func (p *strMapParser) Done() {
	p.delFieldsBuf(p.fields)
	p.template = ""
	p.left = ""
	p.right = ""
	p.fields = nil
	p.before = ""
}

func (p *strMapParser) Init(tmpl string, l, r string) lderr.Error {
	p.delFieldsBuf(p.fields)
	p.fields = p.newFieldsBuf()

	for {
		lIdx := strings.Index(tmpl, l)
		if lIdx < 0 {
			break
		}
		rIdx := strings.Index(tmpl[lIdx+1:], r)
		if rIdx < 0 {
			break
		}

		lIdx = strings.LastIndex(tmpl[:lIdx+1+rIdx], l)

		before := tmpl[:lIdx]
		name := tmpl[lIdx+1 : lIdx+1+rIdx]
		if len(p.fields) == 0 {
			p.before = before

		} else if before == "" {
			return lderr.ErrInvalidTemplateSyntax

		} else {
			p.fields[len(p.fields)-1].After = before
		}

		p.fields = append(p.fields, strMapField{
			Name:  name,
			After: "",
		})

		tmpl = tmpl[lIdx+1+rIdx+1:]
	}

	if len(p.fields) == 0 {
		p.before = tmpl
	} else {
		p.fields[len(p.fields)-1].After = tmpl
	}

	p.template = tmpl
	p.left = l
	p.right = l
	return nil
}

func (p *strMapParser) Parse(text string) (map[string]string, lderr.Error) {
	res := make(map[string]string, len(p.fields))

	if len(p.fields) == 0 {
		if p.before != text {
			return nil, lderr.ErrInvalidTemplateSyntax
		}

		return res, nil
	}

	if !strings.HasPrefix(text, p.before) {
		return nil, lderr.ErrInvalidTemplateSyntax
	}

	text = text[len(p.before):]

	c := p.newContext(res)
	defer p.delContext(c)

	head := &c.Fields[0]
	head.LastIndex = -1
	head.Text = text

	for field := head; ; {
		newField, ok := p.parseField(field, res)
		field = newField
		if field == nil {
			if ok {
				break
			}
			return nil, lderr.ErrInvalidTemplateSyntax
		}
	}

	return res, nil
}

func (p *strMapParser) parseField(c *strMapFieldContext, res map[string]string) (*strMapFieldContext, bool) {
	after := c.Field.After
	name := c.Field.Name

	if v, ok := res[name]; ok {
		c.Duplicate = true
		c.Value = v
	}

	if c.Duplicate {
		idx := len(c.Value)
		if !strings.HasPrefix(c.Text, c.Value) {
			return p.prevField(c, res), false

		} else if !strings.HasPrefix(c.Text[idx:], after) {
			return p.prevField(c, res), false
		}

		c.LastIndex = idx

		return p.nextField(c, res), true
	}

	text := c.Text
	if c.LastIndex >= 0 {
		pos := c.LastIndex + len(after) - 1
		text = c.Text[:pos]
	}

	idx := strings.LastIndex(text, after)
	if idx < 0 {
		return p.prevField(c, res), false
	}

	c.Value = c.Text[:idx]
	c.LastIndex = idx

	return p.nextField(c, res), true
}

func (p *strMapParser) nextField(c *strMapFieldContext, res map[string]string) *strMapFieldContext {
	name := c.Field.Name
	if c.Field.Name != "" {
		res[name] = c.Value
	}

	next := c.Next
	if next != nil {
		nextIdx := c.LastIndex + len(c.Field.After)
		next.LastIndex = -1
		next.Text = c.Text[nextIdx:]
	}
	return next
}

func (p *strMapParser) prevField(c *strMapFieldContext, res map[string]string) *strMapFieldContext {
	name := c.Field.Name
	if name != "" && !c.Duplicate {
		delete(res, name)
	}

	return c.Prev
}

func (p *strMapParser) newFieldsBuf() []strMapField {
	buf := strMapFieldsPool.Get().([]strMapField)
	buf = buf[:0]
	return buf
}

func (p *strMapParser) delFieldsBuf(buf []strMapField) {
	if buf != nil {
		strMapFieldsPool.Put(buf)
	}
}

func (p *strMapParser) newContext(res map[string]string) *strMapParseContext {
	c := strMapParseContextPool.Get().(*strMapParseContext)

	c.Fields = c.Fields[:0]

	for i := range p.fields {
		field := &p.fields[i]
		c.Fields = append(c.Fields, strMapFieldContext{
			Field: field,
		})
	}

	for i := 1; i < len(c.Fields); i++ {
		prev := &c.Fields[i-1]
		next := &c.Fields[i]
		prev.Next = next
		next.Prev = prev
	}

	return c
}
func (p *strMapParser) delContext(c *strMapParseContext) {
	c.Fields = c.Fields[:0]
	strMapParseContextPool.Put(c)
}
