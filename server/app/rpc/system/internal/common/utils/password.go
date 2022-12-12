package utils

import "golang.org/x/crypto/bcrypt"

func CheckPassword(passwd string, dbpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(passwd))
}
