package controller

import (
	"fmt"
	"ims/storage"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	inv := storage.NewInventory()
	r := gin.Default()
	r.GET("/product", Home)
	r.POST("/product/create", inv.AddProduct)
	r.PUT("/product/:id", inv.UpdateProduct)
	r.GET("/product/all", inv.GetAllProducts)
	r.DELETE("/product/:id", inv.DeleteProduct)
	r.GET("/product/:name", inv.FindProductByName)
	r.GET("/product/ctg/:category", inv.ListByCategory)
	r.GET("/product/total", inv.TotalValue)

	fmt.Println("Starting Server...")
	r.Run()

}
func Home(c *gin.Context) {
	c.JSON(200, "Welcome Home")

}
