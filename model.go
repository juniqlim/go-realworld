package main

// User 구조체 정의
type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUsers() ([]User, error)
	GetUserByID(id string) (User, error)
}
