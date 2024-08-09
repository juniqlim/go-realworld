package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
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

// User 구조체 정의
type User struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

// createUser - 사용자 생성
func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO users (name, email) VALUES (:name, :email)`
	result, err := db.NamedExec(query, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	c.JSON(http.StatusCreated, user)
}

// getUsers - 모든 사용자 조회
func getUsers(c *gin.Context) {
	var users []User
	err := db.Select(&users, "SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// getUserByID - ID로 사용자 조회
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
