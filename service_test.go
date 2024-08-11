package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type UserFakeRepository struct {
	users  map[int]User
	nextID int
}

func userFakeRepository() UserRepository {
	return &UserFakeRepository{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (r *UserFakeRepository) CreateUser(user User) (User, error) {
	user.ID = r.nextID
	r.users[r.nextID] = user
	r.nextID++
	return user, nil
}

func (r *UserFakeRepository) GetUsers() ([]User, error) {
	users := make([]User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}

func (r *UserFakeRepository) GetUserByID(id string) (User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return User{}, errors.New("invalid user ID")
	}

	user, exists := r.users[userID]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

func TestCreateUser(t *testing.T) {
	fakeRepo := userFakeRepository()
	createdUser, err := createUser(User{Name: "John Doe", Email: "john@example.com"}, fakeRepo)
	assert.NoError(t, err)
	assert.Equal(t, 1, createdUser.ID)
	assert.Equal(t, "John Doe", createdUser.Name)

	createdUser, err = createUser(User{Name: "Park Jisoo", Email: "jisoo@example.com"}, fakeRepo)
	assert.NoError(t, err)
	assert.Equal(t, 2, createdUser.ID)
	assert.Equal(t, "Park Jisoo", createdUser.Name)
}

func TestGetUsers(t *testing.T) {
	fakeRepo := userFakeRepository()
	createUser(User{Name: "John Doe", Email: "john@example.com"}, fakeRepo)
	users, err := getUsers(fakeRepo)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, "John Doe", users[0].Name)
}

func TestGetUserByID(t *testing.T) {
	fakeRepo := userFakeRepository()
	createUser(User{Name: "John Doe", Email: "john@example.com"}, fakeRepo)
	user, err := getUserByID("1", fakeRepo)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.Name)
}
