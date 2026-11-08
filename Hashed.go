package folcon

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// VerifyPassword compares a plain password with a hashed password
func VerifyPassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return errors.New("invalid password")
		}
		return err
	}
	return nil
}
