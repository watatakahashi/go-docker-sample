package user

/*
ユーザー
*/
type User struct {
	Id   int
	Name string
}

type UserRepository interface {
	Get(id int) (User, error) // 取得
	Insert(user User) error   //追加
}
