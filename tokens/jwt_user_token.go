package tokens

// import (
// 	"fmt"

// 	res "gin_api_server/api/responses"
// 	"gin_api_server/api/utils/handler"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	conn "gin_api_server/api/connector"

// 	jwt "github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// )

// func GenerateUserToken(user_id uint) (string, error) {

// 	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

// 	if err != nil {
// 		return "", err
// 	}

// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["user_id"] = user_id
// 	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	newToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))

// 	// save tokens into db
// 	accessToken := handler.UserTokens{AccessToken: newToken, UserId: user_id}

// 	err = conn.DB.Create(&accessToken).Error
// 	if err != nil {
// 		return "", err
// 	}

// 	return newToken, err

// }

// func UserTokenValid(c *gin.Context) error {

// 	tokenString := ExtractUserToken(c)

// 	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("API_SECRET")), nil
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func ExtractUserToken(c *gin.Context) string {

// 	token := c.Query("token")

// 	if token != "" {
// 		return token
// 	}

// 	bearerToken := c.Request.Header.Get("Authorization")
// 	if len(strings.Split(bearerToken, " ")) == 2 {
// 		return strings.Split(bearerToken, " ")[1]
// 	}

// 	return ""
// }

// func ExtractUserTokenID(c *gin.Context) (uint, string, error) {

// 	tokenString1 := ExtractUserToken(c)

// 	tokenString, err := handler.VerifyUserToken(nil, tokenString1)

// 	if tokenString == "" {
// 		// return 0, "", errors.New("Token not valid")

// 		response := res.ErrorResponse{
// 			Status:     "failure",
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "token not valid",
// 			Data:       gin.H{"error": err.Error()},
// 		}

// 		c.JSON(http.StatusBadRequest, response)

// 		return 0, "", nil

// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("API_SECRET")), nil
// 	})
// 	if err != nil {
// 		return 0, "", err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {

// 		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
// 		if err != nil {
// 			return 0, "", err
// 		}

// 		// token time check and refesh token generate

// 		expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)

// 		remainingTime := time.Until(expirationTime)

// 		if remainingTime < 2*time.Minute {

// 			newToken, err := GenerateUserToken(uint(uid))
// 			if err != nil {
// 				return 0, "", err
// 			}

// 			return uint(uid), newToken, nil
// 		}

// 		return uint(uid), "", nil
// 	}
// 	return 0, "", nil
// }
