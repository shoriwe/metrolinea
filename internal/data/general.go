package data

func CheckUserExists(username string) (bool, error) {
	exists, checkError := checkUserExistsCallback(username)
	go LogUserExists(username, exists)
	return exists, checkError
}
