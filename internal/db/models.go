// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import ()

type Account struct {
	ID       int64
	Name     string
	Email    string
	Username string
	Password string
}