/*
 * Copyright (C) distroy
 */

package ldgorm

import (
	"strings"
)

type tagMap map[string][]string

func newTagMap() tagMap {
	return make(tagMap)
}

func (m tagMap) Add(key, value string) {
	key = strings.ToLower(key)
	m[key] = append(m[key], value)
}

func (m tagMap) Set(key, value string) {
	key = strings.ToLower(key)
	m[key] = []string{value}
}

func (m tagMap) Has(key string) bool {
	key = strings.ToLower(key)

	_, ok := m[key]
	return ok
}

func (m tagMap) Values(key string) []string {
	key = strings.ToLower(key)
	return m[key]
}

func (m tagMap) Get(key string) string {
	key = strings.ToLower(key)

	v := m[key]
	if len(v) != 0 {
		return v[0]
	}

	return ""
}

func parseTagString(tag string) tagMap {
	tagList := strings.Split(tag, ";")
	m := newTagMap()
	for _, v := range tagList {
		if len(v) == 0 {
			continue
		}

		l := strings.SplitN(v, ":", 2)
		k := strings.TrimSpace(l[0])
		if len(k) == 0 {
			continue
		}

		v := k
		if len(l) >= 2 {
			v = l[1]
		}

		m.Add(k, v)
	}
	return m
}
