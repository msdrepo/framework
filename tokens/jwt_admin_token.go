package tokens

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"

	folcon "github.com/msdrepo/framework"
	"github.com/msdrepo/framework/handlers"
)

func GenerateAdminToken(admin_id uint) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["admin_id"] = admin_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	newToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))

	// save tokens into db

	query := `insert into admin_tokens (admin_id, access_token) VALUES (?, ?)`
	_, err = folcon.DB.Exec(query, admin_id, newToken)

	if err != nil {
		return "", errors.New("failed to create token")
	}

	return newToken, err

}

func AdminTokenValid(c *folcon.Context) error {

	tokenString := ExtractAdminToken(c)

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractAdminToken(c *folcon.Context) string {

	token := c.Query("token")

	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractAdminTokenID(c *folcon.Context) (uint, string, error) {

	tokenString1 := ExtractAdminToken(c)

	tokenString, err := handlers.VerifyAdminToken(nil, tokenString1)

	if tokenString == "" {

		return 0, "", nil

	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["admin_id"]), 10, 32)
		if err != nil {
			return 0, "", err
		}

		// token time check and refesh token generate

		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

		remainingTime := time.Until(expirationTime)

		if remainingTime < 2*time.Minute {

			newToken, err := GenerateAdminToken(uint(uid))
			if err != nil {
				return 0, "", err
			}

			return uint(uid), newToken, nil
		}

		return uint(uid), "", nil
	}
	return 0, "", nil
}
