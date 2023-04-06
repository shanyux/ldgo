/*
 * Copyright (C) distroy
 */

package gorm

import (
	"fmt"
	"sync/atomic"

	"github.com/distroy/ldgo/ldhook"
	"github.com/jinzhu/gorm"
)

var (
	testErrTransaction = fmt.Errorf("transaction error")
)

type transactionMockData struct {
	Begin                   int32
	Commit                  int32
	Rollback                int32
	RollbackUnlessCommitted int32
}

func mockTransaction(patches ldhook.Patches) *transactionMockData {
	data := &transactionMockData{}

	patches.Applys([]ldhook.Hook{
		ldhook.FuncHook{
			Target: (*gorm.DB).Begin,
			Double: func(db *gorm.DB) *gorm.DB {
				atomic.AddInt32(&data.Begin, 1)
				return db
			},
		},
		ldhook.FuncHook{
			Target: (*gorm.DB).Commit,
			Double: func(db *gorm.DB) *gorm.DB {
				atomic.AddInt32(&data.Commit, 1)
				return db
			},
		},
		ldhook.FuncHook{
			Target: (*gorm.DB).Rollback,
			Double: func(db *gorm.DB) *gorm.DB {
				atomic.AddInt32(&data.Rollback, 1)
				return db
			},
		},
		ldhook.FuncHook{
			Target: (*gorm.DB).RollbackUnlessCommitted,
			Double: func(db *gorm.DB) *gorm.DB {
				atomic.AddInt32(&data.RollbackUnlessCommitted, 1)
				return db
			},
		},
	})

	return data
}
