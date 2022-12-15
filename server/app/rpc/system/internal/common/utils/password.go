package utils

import "golang.org/x/crypto/bcrypt"

func CheckPassword(passwd string, dbpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(passwd))
}

func GenPassword(pass string) (string, error) {
	p, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
	if err != nil {
		return "", err
	}
	return string(p), nil
}
