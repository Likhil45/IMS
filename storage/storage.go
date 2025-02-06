package storage

import (
	"fmt"
	"ims/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Inventory struct {
	products map[int]models.Product
}

func NewInventory() *Inventory {
	return &Inventory{products: make(map[int]models.Product)}
}

func (p *Inventory) AddProduct(c *gin.Context) {

	var prod models.Product
	err := c.BindJSON(&prod)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		fmt.Println("Unable to unmarshall the Data")
		return
	}
	if _, exists := p.products[prod.Id()]; exists {
		c.JSON(http.StatusConflict, gin.H{"msg": "Product already exists"})
		return
	}
	p.products[prod.Id()] = prod
	c.JSON(200, &prod)
	fmt.Println("Created product: ", prod)

	fmt.Println(p.products)

}

func (p *Inventory) GetAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, p.products)
	fmt.Println(p.products)
}

func (p *Inventory) UpdateProduct(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to Convert to Integer")
		return
	}
	var prod models.Product
	err1 := c.BindJSON(&prod)
	if err1 != nil {
		c.JSON(400, gin.H{"msg": err1.Error()})
		fmt.Println("Unable to unmarshall the Data")
		return
	}
	if _, exist := p.products[id]; !exist {
		c.JSON(http.StatusNotFound, gin.H{"err": "Product not Found "})
		return
	}
	p.products[id] = prod
	c.JSON(200, prod)

}
func (p *Inventory) DeleteProduct(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to Convert to Integer")
		return
	}
	if _, exist := p.products[id]; !exist {
		c.JSON(http.StatusNotFound, gin.H{"err": "Product not Found "})
		return
	}
	delete(p.products, id)

}

func (p *Inventory) FindProductByName(c *gin.Context) {
	name := c.Param("name")
	for _, val := range p.products {
		if val.Name == name {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, "No Product Found")

}

func (p *Inventory) ListByCategory(c *gin.Context) {
	ctg := c.Param("category")
	var prods []models.Product
	for _, val := range p.products {
		if val.Category == ctg {
			prods = append(prods, val)
		}
	}
	if len(prods) == 0 {
		c.IndentedJSON(http.StatusNotFound, "No products found")
	}
	c.IndentedJSON(http.StatusOK, prods)

}

func (p *Inventory) TotalValue(c *gin.Context) {
	var total float64
	for _, val := range p.products {
		total += val.Price
	}
	c.IndentedJSON(http.StatusOK, total)
}
func (p *Inventory) Sell(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to Convert to Integer")
		return
	}
	for _, val := range p.products {
		if val.ID == id {
			val.Sell(id)
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}

}
func (p *Inventory) Restock(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to Convert to Integer")
		return
	}
	for _, val := range p.products {
		if val.ID == id {
			val.Restock(id)
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}

}

func (p *Inventory) UpdatePrice(c *gin.Context) {
	var prod models.Product
	err := c.BindJSON(&prod)
	if err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		fmt.Println("Unable to unmarshall the Data")
		return
	}
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Unable to Convert to Integer")
		return
	}
	for _, val := range p.products {
		if val.ID == id {
			val.UpdatePrice(prod.Price)
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
}
