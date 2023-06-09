package storage

import (
	"database/sql"
	"fmt"
	"github.com/zjc17/exif-web/internal/util"
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
		SqlitePath:   util.GetEnvOrDefault("EXIF_WEB_SQLITE_PATH", "/tmp/exif-web.sqlite3"),
		DatabaseType: SQLite,
	}
)

const (
	InMemorySQLite DatabaseType = iota
	SQLite
)

func InitDatabase(option *InitDatabaseOption) (err error) {
	if option == nil {
		option = defaultInitDatabaseOption
	}
	var database *sql.DB
	switch option.DatabaseType {
	case InMemorySQLite:
		database, err = sql.Open("sqlite", ":memory:")
	case SQLite:
		fmt.Println("sqlite path:", option.SqlitePath)
		database, err = sql.Open("sqlite", option.SqlitePath)
	}
	if err != nil {
		return
	}
	fmt.Println("open database success")
	db = database
	err = initSchema()
	return
}

func initSchema() (err error) {
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS kv (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			key VARCHAR(512) NOT NULL UNIQUE,
			value TEXT NOT NULL,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		return
	}
	return
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
