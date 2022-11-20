package auth

import "golang.org/x/crypto/bcrypt"

func EncryptionPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func DecryptionPassword(hashed, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}
