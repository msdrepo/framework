package folcon

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}

	return err

}
