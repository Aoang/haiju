package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Check 检查密码哈希和密码是否对应
func Check(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil {
		return true
	}
	return false
}

// Check 根据密码生成哈希
func Generate(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		return string(hashedPassword), nil
	}
	return "", err
}
