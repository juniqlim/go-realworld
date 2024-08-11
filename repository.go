package main

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func init() {
	// SQLite 메모리 DB 연결 설정
	var err error
	db, err = sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}

	// users 테이블 생성
	schema := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        email TEXT UNIQUE
    );`

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("테이블 생성 실패: %v", err)
	}
	log.Println("테이블 생성 성공!")
}

type UserDBRepository struct {
	db *sqlx.DB
}

func userDBRepository(db *sqlx.DB) UserRepository {
	return &UserDBRepository{
		db: db,
	}
}

// CreateUser 메서드 구현
func (r *UserDBRepository) CreateUser(user User) (User, error) {
	query := `INSERT INTO users (name, email) VALUES (:name, :email)`
	result, err := r.db.NamedExec(query, &user)
	if err != nil {
		return user, err
	}
	id, err := result.LastInsertId()
	user.ID = int(id)
	return user, err
}

// GetUsers 메서드 구현
func (r *UserDBRepository) GetUsers() ([]User, error) {
	var users []User
	err := r.db.Select(&users, "SELECT * FROM users")
	return users, err
}

// GetUserByID 메서드 구현
func (r *UserDBRepository) GetUserByID(id string) (User, error) {
	var user User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	return user, err
}
