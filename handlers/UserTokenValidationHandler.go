package handlers

// import (
// 	conn "gin_api_server/api/connector"
// 	// "gin_api_server/api/response"

// 	res "gin_api_server/api/responses"

// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// )

// // table create
// type UserTokens struct {
// 	gorm.Model
// 	UserId      uint   `gorm:"type:text null;" json:"userid"`
// 	AccessToken string `gorm:"type:text null;" json:"accesstoken"`
// }

// // token verify
// func VerifyUserToken(c *gin.Context, token string) (string, error) {

// 	var t UserTokens

// 	if err := conn.DB.Where("access_token = ?", token).First(&t).Error; err != nil {

// 		response := res.ErrorResponse{
// 			Status:     "failure",
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "token not valid",
// 			Data:       gin.H{"error": err.Error()},
// 		}

// 		c.JSON(http.StatusBadRequest, response)

// 		return "", nil

// 		// return t.AccessToken, errors.New("Token not found")
// 	}

// 	if t.ID == 0 {

// 		response := res.ErrorResponse{
// 			Status:     "failure",
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "token not valid",
// 			Data:       nil,
// 		}

// 		c.JSON(http.StatusBadRequest, response)

// 		return "", nil

// 		// return t.AccessToken, errors.New("Token not found")
// 	}

// 	return t.AccessToken, nil

// }

// // token logout
// func UserTokenLogout(c *gin.Context, userId uint) (string, error) {

// 	var t UserTokens

// 	if err := conn.DB.Where("user_id = ?", userId).Unscoped().Delete(&t).Error; err != nil {

// 		response := res.ErrorResponse{
// 			Status:     "failure",
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "token not valid",
// 			Data:       gin.H{"error": err.Error()},
// 		}

// 		c.JSON(http.StatusBadRequest, response)

// 		return "", nil

// 		// return "", errors.New("Token not found")
// 	}

// 	return "", nil

// }
