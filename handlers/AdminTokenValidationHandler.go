package handlers

import (
	"database/sql"
	"errors"

	folcon "github.com/msdrepo/framework"
)

// admin token struct
type AdminTokens struct {
	AdminId     uint   `json:"admin_id"`
	AccessToken string `json:"access_token"`
}

// token verify
func VerifyAdminToken(c *folcon.Context, token string) (string, error) {

	var err error
	var adminToken AdminTokens

	query := `select admin_id, access_token from admin_tokens where access_token = ?`

	err = folcon.DB.QueryRow(query, token).Scan(
		&adminToken.AdminId,
		&adminToken.AccessToken,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("data not found")
		}
		return "", err
	}

	if adminToken.AdminId == 0 {

		return "", errors.New("token not valid")

	}

	return adminToken.AccessToken, nil
}

// token logout
// func AdminTokenLogout(c *folcon.Context, adminId uint) (string, error) {

// 	var t AdminTokens

// 	if err := conn.DB.Where("admin_id = ?", adminId).Unscoped().Delete(&t).Error; err != nil {

// 		response := res.ErrorResponse{
// 			Status:     "failure",
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "token not valid",
// 			Data:       gin.H{"error": err.Error()},
// 		}

// 		c.JSON(http.StatusBadRequest, response)

// 		return "", nil

// 	}

// 	return "", nil

// }
