package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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

func main() {
	r := gin.Default()

	// 기본 라우트 설정
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserByID)

	// 서버 실행
	r.Run(":8080") // 포트 8080에서 실행
}
