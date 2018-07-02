package storage_test

import (
	"flag"
	"os"

	"../../../lineuplist"
	"./memo"
	"./postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

var storage *lineuplist.Storage
var dsn string

func init() {
	testMemo := flag.Bool("memo", false, "tests the memo package")
	flag.Parse()

	if *testMemo {
		storage = memo.New()
	} else {
		dsn = os.Getenv("PG_TEST_DSN")
		postgres.MigrateDown("file://postgres/migrations/", dsn)
		postgres.MigrateUp("file://postgres/migrations/", dsn)

		storage = postgres.New(dsn)
	}
}
