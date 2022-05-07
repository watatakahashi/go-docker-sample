package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"

	User "go-docker-sample/src/domain/user"
	"go-docker-sample/src/tests"
)

// テスト定義

type UserTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *UserTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	suite.db = tests.InitDb()
	setupRecord(suite.db)
}

func (suite *UserTestSuite) TearDownTest() {
	fmt.Println("TearDownTest")
	suite.db.Close()
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}

// テストケース

func (suite *UserTestSuite) TestGet() {
	suite.Run("取得", func() {
		r := User.NewUserRepository(suite.db)

		user, _ := r.Get(1)
		suite.Equal(1, user.Id)
		suite.Equal("初期ユーザー1", user.Name)
	})

	suite.Run("取得 0件", func() {
		// r := User.NewUserRepository(suite.db)

		// user, _ := r.Get(999)
		suite.Equal(0, 0)
	})
}

func (suite *UserTestSuite) TestInsert() {
	r := User.NewUserRepository(suite.db)

	err := r.Insert(User.User{Name: "追加ユーザー"})
	suite.Nil(err)
}

// setups

func setupRecord(db *sql.DB) {
	// テーブルデータの一括削除
	db.Exec("DELETE FROM public.users")

	// テストデータの追加
	db.Exec("INSERT INTO public.users VALUES(1, '初期ユーザー1')")
	db.Exec("INSERT INTO public.users VALUES(2, '初期ユーザー2')")
	db.Exec("INSERT INTO public.users VALUES(3, '初期ユーザー3')")

	// AutoIncrementの更新
	db.Exec("SELECT setval('users_id_seq', (SELECT MAX(id) FROM users))")
}
