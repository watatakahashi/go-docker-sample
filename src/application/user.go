package application

import (
	User "go-docker-sample/src/domain/user"
)

/*
ユーザーに関するアプリケーション
*/
type UserApplication interface {
	/*
		ユーザーを取得する
		id: ユーザーID
	*/
	GetUser(id int) (User.User, error)
}

type userApplication struct {
	repository User.UserRepository
}

func NewUserApplication(repository User.UserRepository) userApplication {
	return userApplication{repository: repository}
}

func (u userApplication) GetUser(id int) (User.User, error) {
	user, err := u.repository.Get(id)
	return user, err
}
