package main

// createUser - 사용자 생성
func createUser(user User, repo UserRepository) (User, error) {
	return repo.CreateUser(user)
}

// getUsers - 모든 사용자 조회
func getUsers(repo UserRepository) ([]User, error) {
	return repo.GetUsers()
}

// getUserByID - ID로 사용자 조회
func getUserByID(id string, repo UserRepository) (User, error) {
	return repo.GetUserByID(id)
}
