// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import ()

type Article struct {
	ID     int32
	Name   string
	UserID int32
}

type User struct {
	ID   int32
	Name string
}
