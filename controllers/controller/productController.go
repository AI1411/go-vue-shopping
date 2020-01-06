package controller

import (
	"github.com/gin-gonic/gin"
	"go-vue-spa/models/db"
	"go-vue-spa/models/entity"
	"strconv"
)

const (
	NotPurchased = 0

	Purchased = 1
)

func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	c.JSON(200, resultProducts)
}

func FindProduct(c *gin.Context) {
	productIdStr := c.Query("productID")
	productID, _ := strconv.Atoi(productIdStr)
	reslutProduct := db.FindProduct(productID)

	c.JSON(200, reslutProduct)
}

func AddProduct(c *gin.Context) {
	productName := c.PostForm("productName")
	productMemo := c.PostForm("productMemo")

	var product = entity.Product{
		Name:  productName,
		Memo:  productMemo,
		State: NotPurchased,
	}
	db.InsertProduct(&product)
}

func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")
	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}
	db.UpdateStateProduct(productID, changeState)
}

func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")

	productID, _ := strconv.Atoi(productIDStr)

	db.DeleteProduct(productID)
}
