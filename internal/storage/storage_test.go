package storage

import (
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
	"testing"
)

func TestInitDatabase_NilParam(t *testing.T) {
	err := InitDatabase(nil)
	assert.Nil(t, err)
}

func TestInitDatabase_InMemorySQLite(t *testing.T) {
	InitDatabase(&InitDatabaseOption{
		DatabaseType: InMemorySQLite,
	})
	db := DB()
	err := db.Ping()
	assert.Nil(t, err)
}

func TestOpenDatabase_SQLite(t *testing.T) {
	InitDatabase(&InitDatabaseOption{
		DatabaseType: SQLite,
		SqlitePath:   "/tmp/exif-web.sqlite",
	})
	db := DB()
	err := db.Ping()
	assert.Nil(t, err)
}

func Test_Store(t *testing.T) {
	InitDatabase(&InitDatabaseOption{
		DatabaseType: SQLite,
		SqlitePath:   "/tmp/exif-web.sqlite3",
	})
	var err error
	err = Store("key", "value")
	assert.Nil(t, err)
}

func Test_Get(t *testing.T) {
	InitDatabase(&InitDatabaseOption{
		DatabaseType: InMemorySQLite,
	})

	err := Store("key", "value")
	assert.Nil(t, err)
	value, err := Get("key")
	assert.True(t, value == "value")

	err = Store("key", "value-1")
	assert.Nil(t, err)
	value, err = Get("key")
	assert.True(t, value == "value-1")
}

func TestInitDatabase_InitSchemaError(t *testing.T) {
	err := InitDatabase(&InitDatabaseOption{
		DatabaseType: SQLite,
		SqlitePath:   "/not-exist-file-path/not-exist-file",
	})
	assert.NotNil(t, err)
}
