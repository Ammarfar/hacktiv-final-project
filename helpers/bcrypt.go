package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 8)

	return string(hash), err
}

func ComparePass(h, p string) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
