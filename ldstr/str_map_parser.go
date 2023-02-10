/*
 * Copyright (C) distroy
 */

package ldstr

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/distroy/ldgo/lderr"
)

func mustMarshalJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

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
	Key       string // key
	After     string // after text
	DupIndex  int    // duplicate key index
	DupFirst  bool   // is first duplicate key
	PrevIndex int    // prev index
	NextIndex int    // next index
	Skip      bool
}

type strMapFieldContext struct {
	Field     *strMapField
	Prev      *strMapFieldContext `json:"-"`
	Next      *strMapFieldContext `json:"-"`
	Duplicate *strMapFieldContext `json:"-"`

	Text      string
	Value     string
	LastIndex int
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

	p.template = tmpl
	p.left = l
	p.right = r

	if err := p.initFieldsByTemplate(); err != nil {
		return err
	}

	if err := p.initFieldsDuplicate(); err != nil {
		return err
	}

	if err := p.initFieldsPrevIndex(); err != nil {
		return err
	}
	return nil
}

func (p *strMapParser) initFieldsByTemplate() lderr.Error {
	tmpl := p.template
	l, r := p.left, p.right

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
		key := tmpl[lIdx+1 : lIdx+1+rIdx]
		if len(p.fields) == 0 {
			p.before = before

		} else {
			p.fields[len(p.fields)-1].After = before
		}

		p.fields = append(p.fields, strMapField{
			Key:       key,
			After:     "",
			PrevIndex: -1,
			DupIndex:  -1,
		})

		tmpl = tmpl[lIdx+1+rIdx+1:]
	}

	if len(p.fields) == 0 {
		p.before = tmpl
	} else {
		p.fields[len(p.fields)-1].After = tmpl
	}

	return nil
}

func (p *strMapParser) initFieldsDuplicate() lderr.Error {
	if len(p.fields) <= 32 {
		return p.initFieldsDuplicateSmall()
	}

	fieldKeys := make(map[string]int)
	for i := range p.fields {
		v := &p.fields[i]
		v.DupIndex = -1

		if v.Key == "" {
			continue
		}

		if idx, ok := fieldKeys[v.Key]; ok {
			v.DupIndex = idx
			continue
		}

		fieldKeys[v.Key] = i
	}

	return nil
}

func (p *strMapParser) initFieldsDuplicateSmall() lderr.Error {
	for i := 1; i < len(p.fields); i++ {
		vi := &p.fields[i]
		vi.DupIndex = -1

		if vi.Key == "" {
			continue
		}

		for j := 0; j < i; j++ {
			vj := &p.fields[j]
			if vi.Key == vj.Key {
				vi.DupIndex = j
				break
			}
		}
	}

	return nil
}

func (p *strMapParser) initFieldsPrevIndex() lderr.Error {
	for i := 1; i < len(p.fields); i++ {
		curr := &p.fields[i]
		prev := &p.fields[i-1]
		curr.PrevIndex = i - 1
		prev.NextIndex = i

		if curr.After != "" || curr.DupIndex >= 0 || i == len(p.fields)-1 {
			continue
		}

		j := i + 1
		for ; j < len(p.fields) && p.fields[j].After == ""; j++ {
			next := &p.fields[j]
			if next.After == "" && next.DupIndex < 0 {
				next.Skip = true
				continue
			}

			curr.After = next.After
			curr.NextIndex = j
			next.PrevIndex = i
			break
		}
		i = j
	}
	return nil
}

func (p *strMapParser) Parse(text string) (map[string]string, lderr.Error) {
	if len(p.fields) == 0 {
		if p.before != text {
			return nil, lderr.ErrInvalidTemplateSyntax
		}

		return make(map[string]string), nil
	}

	if !strings.HasPrefix(text, p.before) {
		return nil, lderr.ErrInvalidTemplateSyntax
	}

	text = text[len(p.before):]

	c := p.newContext()
	defer p.delContext(c)

	head := &c.Fields[0]
	head.LastIndex = -1
	head.Text = text

	for field := head; ; {
		newField, ok := p.parseField(field)
		// log.Printf("=== %#v", c.Fields)
		// log.Printf("=== %s", mustMarshalJson(c.Fields))

		if newField == nil {
			if ok {
				break
			}
			return nil, lderr.ErrInvalidTemplateSyntax
		}

		field = newField
	}

	res := make(map[string]string, len(p.fields))
	for i := range c.Fields {
		v := &c.Fields[i]
		if v.Field.Key == "" {
			continue
		}

		res[v.Field.Key] = v.Value
	}
	return res, nil
}

func (p *strMapParser) parseField(c *strMapFieldContext) (*strMapFieldContext, bool) {
	after := c.Field.After

	if c.Duplicate != nil {
		if c.LastIndex >= 0 {
			return p.nextField(c)
		}

		c.Value = c.Duplicate.Value
		idx := len(c.Value)
		if !strings.HasPrefix(c.Text, c.Value) {
			return p.prevField(c)

		} else if !strings.HasPrefix(c.Text[idx:], after) {
			return p.prevField(c)
		}

		c.LastIndex = idx

		return p.nextField(c)
	}

	text := c.Text
	if c.LastIndex >= 0 {
		pos := c.LastIndex + len(after) - 1
		text = c.Text[:pos]
	}

	idx := strings.LastIndex(text, after)
	if idx < 0 {
		return p.prevField(c)
	}

	c.Value = c.Text[:idx]
	c.LastIndex = idx

	return p.nextField(c)
}

func (p *strMapParser) nextField(c *strMapFieldContext) (*strMapFieldContext, bool) {
	nextIdx := c.LastIndex + len(c.Field.After)
	nextText := c.Text[nextIdx:]

	next := c.Next
	if next != nil {
		next.LastIndex = -1
		next.Text = nextText
		return next, true
	}

	if nextText == "" {
		return nil, true
	}

	return c, false
}

func (p *strMapParser) prevField(c *strMapFieldContext) (*strMapFieldContext, bool) {
	return c.Prev, false
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

func (p *strMapParser) newContext() *strMapParseContext {
	c := strMapParseContextPool.Get().(*strMapParseContext)

	c.Fields = c.Fields[:0]

	for i := range p.fields {
		v := &p.fields[i]
		c.Fields = append(c.Fields, strMapFieldContext{
			Field:     v,
			Prev:      nil,
			Next:      nil,
			Duplicate: nil,
		})
	}

	for i := 1; i < len(c.Fields); i++ {
		curr := &c.Fields[i]
		if idx := curr.Field.PrevIndex; idx >= 0 {
			prev := &c.Fields[idx]
			prev.Next = curr
			curr.Prev = prev
		}

		if idx := curr.Field.DupIndex; idx >= 0 {
			curr.Duplicate = &c.Fields[idx]
		}
	}

	return c
}
func (p *strMapParser) delContext(c *strMapParseContext) {
	c.Fields = c.Fields[:0]
	strMapParseContextPool.Put(c)
}
