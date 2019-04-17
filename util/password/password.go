package password

import "golang.org/x/crypto/bcrypt"

//VerifyPassword ..
func VerifyPassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}

//HashPassword ..
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 3)
}
