package main

import (
	"encoding/json"
	"net/http"
	"rest_api/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
TODO EXPLANATIONS

todo HTTP Methods in RESTful APIs:

1. GET: Used to retrieve data from a specified resource.
2. POST: Used to submit data to be processed to a specified resource.
3. PUT: Used to update a resource or create a new resource if it does not exist.
4. DELETE: Used to delete a resource.

todo explain

the ctx = the context of an HTTP request and response. The ctx.JSON function is used to send a JSON response.

gin.H{"product": product} == map[string]interface{"product": product} have the same meaning


*/

func main() {
	// todo starter gin framework
	router := gin.Default()

	// todo connect to database
	database.Connect()

	// todo show all the row in table
	router.GET("/product", func(ctx *gin.Context) {
		var product []database.Product

		// select * from product
		database.DB.Find(&product)

		// sending an HTTP response with a status code of 200 (OK) along with a JSON payload containing information about a product
		ctx.JSON(http.StatusOK, gin.H{"product": product})
	})

	// todo show selected id
	router.GET("/product/:id", func(ctx *gin.Context) {
		// create a variable to store the id
		var product database.Product

		// retrieve a parameter from the URL. Specifically, it's used to extract values from named parameters in the path of a URL.
		id := ctx.Param("id")

		// select * from product where "id" = id
		if err := database.DB.First(&product, id).Error; err != nil {
			// create an option for error
			switch err {
			// if record not found in the table
			case gorm.ErrRecordNotFound:
				// return an error in json
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data not found"})
				return
			// if the server encountered an unexpected condition
			default:
				// return the error in json
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}

		// sending an HTTP response with a status code of 200 (OK) along with a JSON payload containing information about a product
		ctx.JSON(http.StatusOK, gin.H{"product": product})
	})

	// todo create new data
	router.POST("/product", func(ctx *gin.Context) {
		var product database.Product

		// used to parse and bind JSON data from the request body to a Go struct
		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		database.DB.Create(&product)
		ctx.JSON(http.StatusOK, gin.H{"product": product})
	})

	// todo update the data
	router.PUT("/product/:id", func(ctx *gin.Context) {
		var product database.Product
		id := ctx.Param("id")

		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if database.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot update the data"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"product": "success update the data"})
	})

	// todo delete data from selected id
	router.DELETE("/product/delete/:id", func(ctx *gin.Context) {
		var product database.Product

		var input struct {
			Id json.Number
		}

		// used to parse and bind JSON data from the request body to a Go struct
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		id, _ := input.Id.Int64()
		if database.DB.Delete(&product, id).RowsAffected == 0 {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "cannot delete the data"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success delete data"})
	})

	// todo run gin
	router.Run()
}
