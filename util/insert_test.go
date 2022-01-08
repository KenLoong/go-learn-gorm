package util

import (
	"fmt"
	"learn-gorm/model"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func InitTest() error {
	err := Connect("warson_test", "root", "dc")
	return err
}

// MustTruncateTable truncates given tables, fatal on error.
func MustTruncateTable(tb testing.TB, db *gorm.DB, tables ...string) {
	for _, table := range tables {
		err := db.Exec(fmt.Sprintf("truncate table `%s`", table)).Error
		require.NoError(tb, err)
	}
}

func TestCreate2(t *testing.T) {
	InitTest()
	MustTruncateTable(t, Db, "user")
	Create2()
}

func TestBatchInset(t *testing.T) {
	err := InitTest()
	require.NoError(t, err)

	stmt := &gorm.Statement{DB: Db}
	stmt.Parse(&model.User{})
	tableName := stmt.Schema.Table
	MustTruncateTable(t, Db, tableName)
	err = BatchInset()
	require.NoError(t, err)
}
