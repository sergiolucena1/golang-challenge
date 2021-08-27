package controllers

import (
	"github.com/gin-gonic/gin"
	"golang-challenge/database"
	"golang-challenge/models"
	"log"
	"strconv"
)

func ShowProduct(c *gin.Context){
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		log.Println(err)
	}

	db := database.GetDatabase()

	var product models.Product
	err = db.First(&product, newid).Error
	if err != nil{
		c.JSON(400, gin.H{
			"error": "cannot fing book" + err.Error(),
		})
		return
	}

	c.JSON(200, product)
}
