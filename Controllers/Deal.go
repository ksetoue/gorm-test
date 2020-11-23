package Controllers

import (
	"fmt"
	"gorm-test/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetDeals ... Get all deals
func GetDeals(c *gin.Context) {
	var deal []Models.Deal
	err := Models.GetAllDeals(&deal)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, deal)
	}
}

//CreateDeal ... Create Deal
func CreateDeal(c *gin.Context) {
	var deal Models.Deal
	c.BindJSON(&deal)
	err := Models.CreateDeal(&deal)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, deal)
	}
}

//GetDealByID ... Get the Deal by id
func GetDealByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var deal Models.Deal
	err := Models.GetDealByID(&deal, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, deal)
	}
}

//UpdateDeal ... Update the Deal information
func UpdateDeal(c *gin.Context) {
	var deal Models.Deal
	id := c.Params.ByName("id")
	err := Models.GetDealByID(&deal, id)
	if err != nil {
		c.JSON(http.StatusNotFound, deal)
	}
	c.BindJSON(&deal)
	err = Models.UpdateDeal(&deal, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, deal)
	}
}

//DeleteDeal ... Delete the Deal
func DeleteDeal(c *gin.Context) {
	var deal Models.Deal
	id := c.Params.ByName("id")
	err := Models.DeleteDeal(&deal, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
