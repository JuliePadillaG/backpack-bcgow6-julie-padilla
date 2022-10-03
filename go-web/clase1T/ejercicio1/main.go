package main

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"time"
	"os"
	//"fmt"
)

const jsonPath = "/Users/jpadillagonz/Desktop/github/backpack-bcgow6-julie-padilla/backpack-bcgow6-julie-padilla/go-web/ejercicio1/products.json"



type product struct {
	Id		 		int			`json:"id"`
	Name 	 		string		`json:"name"`
	Colour	 		string		`json:"colour"`
	Price    		float64		`json:"price"`
	Stock	 		int			`json:"stock"`
	Code	 		string		`json:"code"`
	Published		bool		`json:"published"`
	CreationDate	time.Time	`json:"creationDate"`
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

	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	ctx.JSON(200, products)
}

// Según la temática elegida, necesitamos agregarles filtros a nuestro endpoint, el mismo se tiene que poder filtrar por todos los campos.
// Dentro del handler del endpoint, recibí del contexto los valores a filtrar.
// Luego genera la lógica de filtrado de nuestro array.
// Devolver por el endpoint el array filtrado.

func productsFilter(ctx *gin.Context) {

	
}


func main() {

	router := gin.Default()

	router.GET("/products", getAll)

	router.GET("/products", productsFilter)

	router.Run(":8080")
}

