package encrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	// Generate a hashed version of our password
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	passwordHash := string(hashed)
	return passwordHash
}

func CheckPasswordHash(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}

	return true
}
