package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDD(t *testing.T) {
	user := User{Name: "John Doe", Email: "j@j.com"}
	fmt.Print(user)
	article := Article{title: "title", content: "content"}
	fmt.Print(article)
	assert.Equal(t, 1, 1)
}
