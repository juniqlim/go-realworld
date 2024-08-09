package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
