package handlers

import (
	"go-api/config"
	"go-api/internal/models"
	"go-api/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DataResponse formats success messages consistently
type DataResponse struct {
	CurrentPage     int         `json:"currentPage"`
	TotalPages      int         `json:"totalPages"`
	TotalItems      int64       `json:"totalItems"`
	Limit           int         `json:"limit"`
	HasNextPage     bool        `json:"hasNextPage"`
	HasPreviousPage bool        `json:"hasPreviousPage"`
	Items           interface{} `json:"items"`
}

// Get All Users
func GetUsers(c *gin.Context) {
	// define the table
	var users []models.User

	// define the query params
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	// get total items
	var totalItems int64
	config.DB.Model(&models.User{}).Count(&totalItems)

	// set the limit and page
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}
	totalPages := int(totalItems) / limit
	if int(totalItems)%limit != 0 {
		totalPages++
	}
	if page > totalPages {
		page = totalPages
	}
	offset := (page - 1) * limit

	// check if there is a next page
	var nextPage = false
	var previousPage = false
	if page >= totalPages {
		nextPage = false
	} else {
		nextPage = true
	}

	if page <= 1 {
		previousPage = false
	} else {
		previousPage = true
	}

	// fetching data from database
	if err := config.DB.Limit(limit).Offset(offset).Order("id desc").Find(&users).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	// build the reponse
	dataResponse := DataResponse{
		CurrentPage:     page,
		TotalPages:      totalPages,
		TotalItems:      totalItems,
		Limit:           limit,
		HasNextPage:     nextPage,
		HasPreviousPage: previousPage,
		Items:           users,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching users data", dataResponse)
}

// Get User by ID
func GetUserByID(c *gin.Context) {
	id := c.Query("id")
	var user models.User

	// error handling
	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	// build the reponse
	dataResponse := DataResponse{
		CurrentPage:     1,
		TotalPages:      1,
		TotalItems:      1,
		Limit:           1,
		HasNextPage:     false,
		HasPreviousPage: false,
		Items:           user,
	}
	utils.Response(c, http.StatusOK, true, "Succes fetching user data", dataResponse)
}

// Create User
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, false, "Wrong JSON format", nil)
		return
	}

	// Hash the password before saving
	if err := user.HashPassword(user.Password); err != nil {
		utils.Response(c, http.StatusInternalServerError, false, "Failed to hash password", nil)
		return
	}

	user.Password = ""

	if err := config.DB.Create(&user).Error; err != nil {
		utils.Response(c, http.StatusConflict, false, "Email already taken", nil)
		return
	}

	utils.Response(c, http.StatusCreated, true, "User created", user)
}

// Update User
func UpdateUser(c *gin.Context) {
	id := c.Query("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.Response(c, http.StatusBadRequest, false, "Wrong JSON format", nil)
		return
	}

	config.DB.Save(&user)
	utils.Response(c, http.StatusOK, true, "User updated", user)
}

// Delete User
func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		utils.Response(c, http.StatusNotFound, false, "User not found", nil)
		return
	}

	config.DB.Delete(&user)
	utils.Response(c, http.StatusOK, true, "User deleted", user)
}
