package main

func isValidPassword(password string) bool {
	// ?
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	hasUppercase := false
	hasNumber := false
	for _, c := range password {
		if c >= 65 && c <= 90 {
			hasUppercase = true
		}
		if c >= 48 && c <= 57 {
			hasNumber = true
		}
	}
	return hasNumber && hasUppercase
}
