/*
 * Copyright (C) distroy
 */

package ldconsistenthash

type (
	Hash   = func(data []byte) uint32
	Option = func(m *Map)
)

func WithHash(hash Hash) Option {
	return func(m *Map) {
		m.hash = hash
	}
}

func WithReplicas(replicas int) Option {
	return func(m *Map) {
		m.replicas = replicas
	}
}
