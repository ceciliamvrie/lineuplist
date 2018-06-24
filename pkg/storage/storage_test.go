package storage_test

import (
	"./memo"
	"../../../lineuplist"
)
var store *lineuplist.Storage

func init() {
	store = memo.New()
}
