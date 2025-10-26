package sqlite

import (
	"database/sql"
	"go.slink.ws/util/files"
	"path"
)

const (
	releaser = "releaser"

	perfTune = `
pragma journal_mode = WAL;
pragma synchronous = normal;
pragma temp_store = memory;
pragma mmap_size = 30000000;`
)

func OpenDatabase(dbFilePath string) *sql.DB {
	files.EnsureDir(path.Dir(dbFilePath))
	db, err := sql.Open("sqlite", dbFilePath)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(perfTune)
	if err != nil {
		panic(err)
	}
	return db
}
