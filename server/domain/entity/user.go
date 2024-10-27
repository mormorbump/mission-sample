package entity

type User struct {
	ID UserID
}

type Users []*User

type UserID int64

type UserPK struct {
	UserID UserID
}

type UserPKs []*UserPK
