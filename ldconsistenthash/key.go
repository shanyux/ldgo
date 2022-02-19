/*
 * Copyright (C) distroy
 */

package ldconsistenthash

type Key struct {
	Key    string
	Weight int
}

type keyInfo struct {
	Key  string
	Hash uint32
}

type sortKeyInfos []keyInfo

func (s sortKeyInfos) Len() int      { return len(s) }
func (s sortKeyInfos) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortKeyInfos) Less(i, j int) bool {
	a, b := &s[i], &s[j]
	if a.Hash != b.Hash {
		return a.Hash < b.Hash
	}
	return a.Key <= b.Key
}

// type sortUint32s []uint32
//
// func (s sortUint32s) Len() int           { return len(s) }
// func (s sortUint32s) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
// func (s sortUint32s) Less(i, j int) bool { return s[i] <= s[j] }
