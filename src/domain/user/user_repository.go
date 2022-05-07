package user

import (
	"context"
	"database/sql"
	"fmt"

	sqlc "go-docker-sample/sqlc/db"
)

type userRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepository{q: sqlc.New(db)}
}

func (u userRepository) Get(id int) (User, error) {
	ctx := context.Background()
	user, err := u.q.GetUser(ctx, int32(id))

	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	return User{
			Id:   int(user.ID),
			Name: user.Name},
		nil
}

func (u userRepository) Insert(user User) (int, error) {
	ctx := context.Background()
	userEntity, err := u.q.CreateUser(ctx, user.Name)
	if err != nil {
		return 0, err
	}

	return int(userEntity.ID), nil
}
