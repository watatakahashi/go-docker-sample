package main

import (
	"database/sql"
	"fmt"
	Application "go-docker-sample/src/application"
	Controller "go-docker-sample/src/controller"
	User "go-docker-sample/src/domain/user"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("exec main")

	// ポート番号の指定
	server := http.Server{Addr: ":8080"}

	db := initDb()
	defer db.Close()

	router(db)

	// サーバー起動
	server.ListenAndServe()
}

/*
ルーティングを付与
*/
func router(db *sql.DB) {

	userRepository := User.NewUserRepository(db)
	userApplication := Application.NewUserApplication(userRepository)
	userController := Controller.NewUserController(userApplication)

	http.HandleFunc("/user", userController.GetUser)
}

/*
DBの初期化
*/
func initDb() *sql.DB {

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
