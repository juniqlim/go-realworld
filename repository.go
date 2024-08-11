package main

import (
	"database/sql"
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

func createUserDB(user User) (sql.Result, error) {
	query := `INSERT INTO users (name, email) VALUES (:name, :email)`
	return db.NamedExec(query, &user)
}

func getUsersDB() ([]User, error) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users")
	return users, err
}

func getUserByIDDB(id string) (User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	return user, err
}
