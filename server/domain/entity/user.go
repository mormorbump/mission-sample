package entity

type User struct {
	UserID UserID
}

type Users []*User

type UserID string
