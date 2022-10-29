package service

func CheckUser(user string) bool {
	if user == "user" {
		return true
	}

	return false
}

func CheckPassword(password string) bool {
	if password == "password" {
		return true
	}

	return false
}

func Login(user, password string) bool {
	if CheckPassword(password) == true && CheckUser(user) {
		return true
	}

	return false
}
