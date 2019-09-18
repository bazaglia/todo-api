package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt hash and salt password. Ideally we would use a secret key for encrypting password, but
// in this case bcrypt is taking care of it
func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePasswords given hashed and plain password, returns either their match or not
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	return err == nil
}
