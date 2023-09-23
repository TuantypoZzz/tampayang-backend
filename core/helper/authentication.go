package helper

import "golang.org/x/crypto/bcrypt"


func HashingPassword(password string) string {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic("helper - HashingPassword, bcrypt.GenerateFromPassword " + err.Error())
	}

	return string(hashByte)
}

func ValidatePassword(hashedPassword, password  string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}