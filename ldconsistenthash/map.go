/*
 * Copyright (C) distroy
 */

package ldconsistenthash

import (
	"crypto/md5"
	"sort"
	"strconv"

	"github.com/distroy/ldgo/ldconv"
)

var (
	_defaultHash     = md5Hash
	_defaultReplicas = 1
)

type Map struct {
	hash     Hash
	replicas int
	keys     []keyInfo // Sorted
	// keyHashes map[string][]uint32
}

func New(opts ...Option) *Map {
	m := &Map{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// IsEmpty returns true if there are no items available.
func (m *Map) IsEmpty() bool { return len(m.keys) == 0 }

// Adds adds some keys to the hash.
func (m *Map) Adds(keys ...string) {
	m.init()
	m.checkCapicity(len(keys))

	for _, key := range keys {
		m.add(&Key{Key: key, Weight: 1})
	}

	m.sort()
}

// Add adds one key (with weight) to the hash.
func (m *Map) Add(key string, weight ...int) {
	m.init()

	w := 1
	if len(weight) > 0 && weight[0] > 0 {
		w = weight[0]
	}
	m.checkCapicity(w)

	m.add(&Key{Key: key, Weight: w})
	m.sort()
}

func (m *Map) AddKeys(keys ...Key) {
	m.init()

	newCount := 0
	for i := range keys {
		key := &keys[i]
		if key.Weight <= 0 {
			key.Weight = 1
		}
		newCount += key.Weight
	}

	m.checkCapicity(newCount)
	for i := range keys {
		m.add(&keys[i])
	}

	m.sort()
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}

	hash := m.getHash(key)

	// Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i].Hash >= hash
	})

	// Means we have cycled back to the first replica.
	if idx == len(m.keys) {
		idx = 0
	}

	return m.keys[idx].Key
}

func (m *Map) add(key *Key) {
	replicas := m.replicas * key.Weight
	if m.hash == nil {
		replicas = replicas + 4 - (replicas % 4)
	}

	if m.hash != nil {
		for i := 0; i < replicas; i++ {
			hash := m.getHash(key.Key + strconv.Itoa(i))
			m.keys = append(m.keys, keyInfo{Key: key.Key, Hash: hash})
			// m.hashMap[hash] = key.Key
		}
		return
	}

	for i, n := 0, replicas/4; i < n; i++ {
		h := md5.Sum(ldconv.StrToBytesUnsafe(key.Key + strconv.Itoa(i)))
		for j := 0; j < len(h); j += 4 {
			s := h[j : j+4]
			x := ((uint32(s[3]) << 24) | (uint32(s[2]) << 16) | (uint32(s[1]) << 8) | uint32(s[0]))
			m.keys = append(m.keys, keyInfo{Key: key.Key, Hash: x})
			// m.hashMap[x] = key.Key
		}
	}
}

// func (m *Map) del(key string, hashes []uint32) {
// 	// hashes := m.keyHashes[key]
//
// 	sort.Sort(sortUint32s(hashes))
//
// 	lastEnd := 0
// 	for _, hash := range hashes {
// 		idx := sort.Search(len(m.keys), func(i int) bool {
// 			k := &m.keys[i]
// 		})
// 	}
// }

func (m *Map) init() {
	// if m.keyHashes == nil {
	// 	m.keyHashes = make(map[string][]uint32)
	// }
	if m.replicas <= 0 {
		m.replicas = _defaultReplicas
	}
}

func (m *Map) checkCapicity(newCount int) {
	replicas := m.replicas
	if m.hash == nil {
		replicas = replicas + 4 - (replicas % 4)
	}

	newLen := (newCount * replicas) + len(m.keys)
	if newLen <= cap(m.keys) {
		return
	}

	newKeys := make([]keyInfo, 0, newLen)
	newKeys = append(newKeys, m.keys...)
	m.keys = newKeys
}

func (m *Map) sort() {
	sort.Sort(sortKeyInfos(m.keys))
}

func (m *Map) getHash(s string) uint32 {
	fnHash := m.hash
	if m.hash == nil {
		fnHash = _defaultHash
	}
	hash := fnHash(ldconv.StrToBytesUnsafe(s))
	// hash = (hash >> 16) ^ (hash & 0xffff)
	return hash
}
