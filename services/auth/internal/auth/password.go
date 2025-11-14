package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes a plaintext password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPassword compares plain text with hashed
func CheckPassword(hashedPassword, plainPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
    return err == nil
}
