package handlers

import (
	"fmt"
	"go-api/config"
	"go-api/internal/auth"
	"go-api/internal/models"
	"go-api/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// token struct
type Token struct {
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	ActiveUntil TimeJSON `json:"activeUntil"`
	Token       string   `json:"token"`
}

// Custom time type that formats JSON output
type TimeJSON struct {
	time.Time
}

// Custom JSON marshaling to format time
func (t TimeJSON) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// LoginHandler authenticates users and generates JWT
func LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	// hashing the password
	if err := user.HashPassword(loginRequest.Password); err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to hash password", nil)
		return
	}

	if err := config.DB.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		utils.Response(c, http.StatusUnauthorized, false, "Invalid email", nil)
		return
	}

	if !user.CheckPassword(loginRequest.Password) {
		utils.Response(c, http.StatusUnauthorized, false, "Invalid password", nil)
		return
	}

	// Generate JWT
	token, err := auth.GenerateToken(loginRequest.Email, user.Role)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to generate token", nil)
		return
	}

	// build the reponse
	dateTime := TimeJSON{Time: time.Now().Add(1 * time.Hour)}
	dataResponse := Token{
		Username:    user.Name,
		Email:       user.Email,
		ActiveUntil: dateTime,
		Token:       token,
	}
	utils.Response(c, http.StatusOK, true, "Success to generate token", dataResponse)
}
