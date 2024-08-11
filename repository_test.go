package main

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

var testDB *sqlx.DB

// 테스트 환경을 위한 초기화
func init() {
	var err error
	testDB, err = sqlx.Open("sqlite3", ":memory:")
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

	// 함수를 테스트하기 전에 실제 DB를 사용하지 않도록, 테스트용 DB를 사용하도록 설정합니다.
	originalDB := db
	db = testDB
	defer func() { db = originalDB }() // 테스트 후 원래 DB로 복원
}

func TestCreateUserDB(t *testing.T) {
	user := User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	result, _ := createUserDB(user)
	lastInsertID, _ := result.LastInsertId()
	assert.True(t, lastInsertID > 0)
	rowsAffected, _ := result.RowsAffected()
	assert.Equal(t, int64(1), rowsAffected)
}
