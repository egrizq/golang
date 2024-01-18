package controllers

import (
	"gin_gorm/database"
	"gin_gorm/model"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Views(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func Create(ctx *gin.Context) {
	var createData model.Data
	var tempData struct {
		Name        string `form:"name"`
		Produsen    string `form:"produsen"`
		Description string `form:"description"`
		Quantity    int    `form:"quantity"`
		Area        string `form:"area"`
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot bind file"})
		return
	}

	filename := filepath.Join("public", file.Filename)
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot save the file"})
		return
	}

	if err := ctx.ShouldBind(&tempData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot bind"})
		return
	}

	createData.File = file.Filename

	createData.Name = tempData.Name
	createData.Produsen = tempData.Produsen
	createData.Description = tempData.Description
	createData.Area = tempData.Area
	createData.Quantity = tempData.Quantity

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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id := newData.ID
	if err := database.DB.Save(&newData).Where("id = ?", id).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
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
