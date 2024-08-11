package main

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var (
	repo UserRepository
)

// 테스트 환경을 위한 초기화
func init() {
	var err error
	testDB, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// users 테이블 생성
	schema := `CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE
	);`

	_, err = testDB.Exec(schema)
	if err != nil {
		panic(err)
	}

	repo = userDBRepository(testDB)
}

func TestCreateUserDB(t *testing.T) {
	user := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	result, _ := repo.CreateUser(user)
	assert.True(t, result.ID > 0)
}
