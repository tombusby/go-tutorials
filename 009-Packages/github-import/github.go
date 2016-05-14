package main

import (
	"database/sql"
	"fmt"

	// The underscore means we're importing solely for side-effects:
	// the package itself is never referenced by name, the _ prevents
	// this causing a compiler error or being removed from source by
	// goimports.
	// Also, this will not automatically download, you must use `go get`
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	fmt.Println(db, err)
}
