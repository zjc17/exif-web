package storage

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

type (
	DatabaseType       int
	InitDatabaseOption struct {
		SqlitePath   string
		DatabaseType DatabaseType
	}
)

var (
	db                        *sql.DB
	defaultInitDatabaseOption = &InitDatabaseOption{
		SqlitePath:   "/tmp/exif-web.sqlite3",
		DatabaseType: SQLite,
	}
)

const (
	InMemorySQLite DatabaseType = iota
	SQLite
)

func init() {
	InitDatabase(nil)
}

func InitDatabase(option *InitDatabaseOption) {
	if option == nil {
		option = defaultInitDatabaseOption
	}
	var database *sql.DB
	var err error
	switch option.DatabaseType {
	case InMemorySQLite:
		database, err = sql.Open("sqlite", ":memory:")
	case SQLite:
		database, err = sql.Open("sqlite", option.SqlitePath)
	}
	if err != nil || database == nil {
		panic("failed to open database")
	} else {
		fmt.Println("open database success")
		db = database
	}
	initSchema()
}

func initSchema() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS kv (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key VARCHAR(512) NOT NULL UNIQUE,
			value TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		panic("failed to init schema")
	}
}

func Store(key, value string) (err error) {
	_, err = db.Exec(`INSERT OR REPLACE INTO kv (key, value) VALUES (?, ?)`, key, value)
	return
}

func Get(key string) (value string, err error) {
	err = db.QueryRow(`SELECT value FROM kv WHERE key = ? LIMIT 1`, key).Scan(&value)
	return
}

func DB() *sql.DB {
	return db
}
