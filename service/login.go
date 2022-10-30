package service

func Login(user, password string) string {
	if user == "user-admin" || password == "password-admin" {
		return "admin"
	}

	if user == "user" || password == "password" {
		return "agency"
	}
	return ""
}
