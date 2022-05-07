package controller

import (
	"encoding/json"
	"fmt"
	User "go-docker-sample/src/application"
	"net/http"
)

type UserController interface {
	GetUser(http.ResponseWriter, *http.Request)
}

type userController struct {
	application User.UserApplication
}

func NewUserController(application User.UserApplication) UserController {
	return userController{application: application}
}

type UserDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c userController) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("exec UserController.GetUser")

	user, err := c.application.GetUser(1)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	dto := UserDto{Id: user.Id, Name: user.Name}
	userResponse, _ := json.MarshalIndent(dto, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(userResponse)

}
