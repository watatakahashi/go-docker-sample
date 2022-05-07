package tests

import (
	"database/sql"
	"fmt"
)

/*
DBに接続する
// TODO: テスト用に用意する
*/
func InitDb() *sql.DB {

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"root",
		"password",
		"sample-db",
		"5432",
		"sampledb")

	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println(err)
		panic(0)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	return db
}
