package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"time"
	"os"
	"net/http"

	"fmt"
)

type product struct {
	Id		 		int			`json:"id"`
	Name 	 		string		`json:"name" binding:"required"`
	Colour	 		string		`json:"colour" binding:"required"`
	Price    		float64		`json:"price" binding:"required"`
	Stock	 		int			`json:"stock" binding:"required"`
	Code	 		string		`json:"code" binding:"required"`
	Published		bool		`json:"published" binding:"required"`
	CreationDate	time.Time	`json:"creationDate" binding:"required"`
}

const jsonPath = "/Users/jpadillagonz/Desktop/github/backpack-bcgow6-julie-padilla/backpack-bcgow6-julie-padilla/go-web/clase1M/ejercicio1/products.json"


var products []product

func main() {
	r := gin.Default()
	
	pr := r.Group("products")
	
	pr.POST("/create", Create())
	pr.GET("", getAll)
	r.Run()
}

func (p product) getFromFile(filePath string) (*[]product, error) {
	var products []product

	rawData, err := os.ReadFile(filePath)

	if err != nil {
		return &products, err
	}

	err = json.Unmarshal(rawData, &products)

	if err != nil {
		return &products, err
	}

	return &products, nil
}

func getAll(ctx *gin.Context) {
	products, err := product{}.getFromFile(jsonPath)

	fmt.Println(jsonPath)
	
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusCreated, products)
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var productRequest product

		if err := c.ShouldBindJSON(&productRequest); err != nil { //GET ALL THE STRUCT USER BY BODY
			if productRequest.Name == "" || productRequest.Colour == "" || productRequest.Price == 0.0 || productRequest.Stock == 0 || productRequest.Code == "" || productRequest.Published == false || productRequest.CreationDate.IsZero() {
				if productRequest.Name == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Field name is required"})
				}
				return
			}
		}

		var req product
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//fmt.Println(product{}.getFromFile(jsonPath))// esto me trae el array del JSON
		fmt.Println("Productos antes:", products)

		data, err := product{}.getFromFile(jsonPath)
		if err != nil {
			panic(err)
		}

		req.Id = len(*data) + 1

		products = append(*data, req)
		fmt.Println("Productos despues", products)

		convert, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}

		err1 := os.WriteFile("./products.json", convert, 0644)
		if err != nil {
		panic(err1)
		}

		c.JSON(http.StatusOK, req)

	}
}