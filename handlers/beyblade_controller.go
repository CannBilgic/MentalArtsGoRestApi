package controller

import (
	"mentalartsgo/config"
	"mentalartsgo/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListBeyblade godoc
// @Summary      ListBeyblade
// @Description  get string beyblade
// @Tags         beyblade
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.Beyblade
// @Failure 400
// @Failure 500
// @Router       /beyblade [get]
func ListBeyblade(ctx *gin.Context) {
	beyblades := []models.Beyblade{}
	config.DB.Preload("Monster").First(&beyblades)
	config.DB.Preload("User").Find(&beyblades)
	ctx.JSON(http.StatusOK, &beyblades)
}

// CreateBeyblade godoc
// @Summary      create an beyblade
// @Description  create by beyblade
// @Tags         beyblade
// @Accept       json
// @Produce      json
// @Param beyblade body models.Beyblade true "Beyblade bilgileri"
// @Success      200  {object}  models.Beyblade
// @Failure 400
// @Failure 500
// @Router       /beyblade [post]
func CreateBeyblade(ctx *gin.Context) {
	var beyblade models.Beyblade
	ctx.BindJSON(&beyblade)

	if err := config.DB.Create(&beyblade).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		config.DB.Preload("Monster").First(&beyblade)
		config.DB.Preload("User").First(&beyblade)
		ctx.JSON(http.StatusCreated, &beyblade)
	}

}

// DeleteBeyblade godoc
// @Summary      DeleteBeyblade an beyblade
// @Description  delete by id beyblade
// @Tags         beyblade
// @Accept       json
// @Produce      json
// @Param id path int true "Beyblade ID"
// @Success      200  {object}  models.Beyblade
// @Failure 400
// @Failure 500
// @Router       /beyblade/{id} [delete]
func DeleteBeyblade(ctx *gin.Context) {
	// Parametre olarak gelen beyblade ID'sini alınır
	beybladeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Beyblade'i veritabanından silme işlemi
	if err := config.DB.Where("id = ?", beybladeID).Delete(&models.Beyblade{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Beyblade deleted successfully"})
}
