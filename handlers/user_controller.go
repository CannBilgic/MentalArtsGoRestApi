package controller

import (
	"mentalartsgo/config"
	"mentalartsgo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListUsers godoc
// @Summary      ListUsers
// @Description  get string user
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.User
// @Failure 400
// @Failure 500
// @Router       /user [get]
func ListUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(http.StatusOK, &users)

}

// CreateUser godoc
// @Summary      create an user
// @Description  create by user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param user body models.User true "User bilgileri"
// @Success      200  {object}  models.User
// @Failure 400
// @Failure 500
// @Router       /user [post]
func CreateUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	if err := config.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusCreated, &user)
	}

}

// DeleteUser godoc
// @Summary      DeleteUser an user
// @Description  delete by id user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param id path int true "User ID"
// @Success      200  {object}  models.User
// @Failure 400
// @Failure 500
// @Router       /user/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	// Parametre olarak gelen user ID'sini alınır
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// User'ı veritabanından silme işlemi
	if err := config.DB.Where("id = ?", userID).Delete(&models.User{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
