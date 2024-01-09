package controllers

import (
	"gin_gorm/database"
	"gin_gorm/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Views(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func Create(ctx *gin.Context) {
	var createData model.Data

	if err := ctx.ShouldBind(&createData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	database.DB.Create(&createData)
	ctx.Redirect(http.StatusSeeOther, "/dashboard")
}

func Dashboard(ctx *gin.Context) {
	var allData []model.Data

	database.DB.Find(&allData)
	ctx.HTML(http.StatusOK, "dashboard.html", gin.H{"message": allData})
}

func Edit(ctx *gin.Context) {
	idQuery := ctx.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var data []model.Data

	database.DB.Find(&data, id)
	ctx.HTML(http.StatusOK, "update.html", gin.H{"message": data})
}

func UpdateData(ctx *gin.Context) {
	var newData model.Data

	if err := ctx.ShouldBind(&newData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot bind existing id"})
		return
	}

	id := newData.ID
	if err := database.DB.Save(&newData).Where("id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no data"})
		return
	}

	// ctx.JSON(http.StatusOK, newData) code to check

	ctx.Redirect(http.StatusSeeOther, "/dashboard")
}

func Delete(ctx *gin.Context) {
	idQuery := ctx.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	var deleteData []model.Data

	database.DB.Delete(&deleteData, id)
	ctx.Redirect(http.StatusSeeOther, "/dashboard")
}
