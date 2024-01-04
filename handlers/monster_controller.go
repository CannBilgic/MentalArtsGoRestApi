package controller

import (
	"mentalartsgo/config"
	"mentalartsgo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListMonster godoc
// @Summary      ListMonster
// @Description  get string monster
// @Tags         monster
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Monster
// @Failure 400
// @Failure 500
// @Router       /monster [get]
func ListMonster(ctx *gin.Context) {
	monsters := []models.Monster{}
	config.DB.Find(&monsters)
	ctx.JSON(http.StatusOK, &monsters)

}

// CreateMonster godoc
// @Summary      create an monster
// @Description  create by monster
// @Tags         monster
// @Accept       json
// @Produce      json
// @Param user body models.Monster true "Monster bilgileri"
// @Success      200  {object}  models.Monster
// @Failure 400
// @Failure 500
// @Router       /monster [post]
func CreateMonster(ctx *gin.Context) {
	var monster models.Monster
	ctx.BindJSON(&monster)
	if err := config.DB.Create(&monster).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusCreated, &monster)
	}

}

// DeleteMonster godoc
// @Summary      DeleteMonster an user
// @Description  delete by id monster
// @Tags         monster
// @Accept       json
// @Produce      json
// @Param id path int true "Monster ID"
// @Success      200  {object}  models.Monster
// @Failure 400
// @Failure 500
// @Router       /monster/{id} [delete]
func DeleteMonster(ctx *gin.Context) {
	// Parametre olarak gelen user ID'sini alınır
	monsterID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// User'ı veritabanından silme işlemi
	if err := config.DB.Where("id = ?", monsterID).Delete(&models.Monster{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Monster deleted successfully"})
}
