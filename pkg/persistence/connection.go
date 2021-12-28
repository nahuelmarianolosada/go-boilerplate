package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err)
	}

    if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("unable to reach database: %v", err))
	}

	if errMigration := makeMigrations(); errMigration != nil {
		panic(errMigration)
	}
}

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}
	return db
}

func makeMigrations() error {
	db := GetConnection()
	q := `CREATE TABLE IF NOT EXISTS USER (
            ID INTEGER PRIMARY KEY AUTOINCREMENT,
            USERNAME VARCHAR(100) NULL,
            PASSWORD VARCHAR(200) NULL,
            LAST_UPDATED TIMESTAMP DEFAULT DATETIME
         );`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}

	qMsg := `CREATE TABLE IF NOT EXISTS MESSAGE (
        ID INTEGER PRIMARY KEY AUTOINCREMENT,
        SENDER INTEGER NOT NULL,
        RECIPIENT INTEGER NOT NULL,
        TYPE VARCHAR(5) NULL,
        TEXT VARCHAR(500) NULL,
        LAST_UPDATED TIMESTAMP DEFAULT DATETIME
     );`
	_, errMsg := db.Exec(qMsg)
	if errMsg != nil {
		return errMsg
	}

	// fmt.Println(result)
	// if res, err := db.Query("SELECT sql FROM sqlite_schema WHERE name = 'MESSAGE';"); err != nil {
	// 	return err
	// } else {
	// 	var table string

	// 	for res.Next() {
	// 		res.Scan(&table)
	// 		fmt.Println(table)
	// 	}
	// }

	return nil
}
