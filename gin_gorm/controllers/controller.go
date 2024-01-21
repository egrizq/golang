package controllers

import (
	"errors"
	"fmt"
	"gin_gorm/database"
	"gin_gorm/model"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))

func Views(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func LoginForm(ctx *gin.Context) {
	var name struct {
		Name string `form:"name"`
	}

	var data model.Data

	if err := ctx.ShouldBind(&name); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	result := database.DB.Where("name = ?", name.Name).First(&data).Error

	// check if the query is not 0
	if errors.Is(result, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "name is not found")
		return
	}

	// storing the session
	session, err := store.Get(ctx.Request, "session-key")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	session.Values["username"] = name.Name
	session.Save(ctx.Request, ctx.Writer)

	ctx.Redirect(http.StatusSeeOther, "/dashboard")
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

	// parse the img from HTML
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot bind file"})
		return
	}

	// saving the file into folder public
	filename := filepath.Join("public", file.Filename)
	if err := ctx.SaveUploadedFile(file, filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot save the file"})
		return
	}

	if err := ctx.ShouldBind(&tempData); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot bind"})
		return
	}

	// saving img name
	createData.File = file.Filename

	// storing the data into db
	createData.Name = tempData.Name
	createData.Produsen = tempData.Produsen
	createData.Description = tempData.Description
	createData.Area = tempData.Area
	createData.Quantity = tempData.Quantity

	database.DB.Create(&createData)
	ctx.Redirect(http.StatusSeeOther, "/dashboard")
}

func Dashboard(ctx *gin.Context) {
	session, err := store.Get(ctx.Request, "session-key")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	// get the session
	username, ok := session.Values["username"].(string)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var allData []model.Data

	database.DB.Where("name = ?", username).Find(&allData)

	log.Println(username)
	ctx.HTML(http.StatusOK, "dashboard.html", gin.H{"message": allData, "username": username})
}

func Edit(ctx *gin.Context) {
	idQuery := ctx.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot bind"})
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

	ctx.Redirect(http.StatusSeeOther, "/dashboard")
}

func Delete(ctx *gin.Context) {
	idQuery := ctx.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var deleteData model.Data

	database.DB.Find(&deleteData, id)
	imagePath := fmt.Sprintf("public/%v", deleteData.File)
	if err := os.Remove(imagePath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	database.DB.Delete(&deleteData, id)
	ctx.Redirect(http.StatusSeeOther, "/dashboard")
}

func ClearSession(ctx *gin.Context) {
	session, err := store.Get(ctx.Request, "session-key")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	delete(session.Values, "username")
	session.Save(ctx.Request, ctx.Writer)
	ctx.Redirect(http.StatusSeeOther, "/")
}
