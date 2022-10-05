package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/julie-padilla/go-web/clase2T/internal/products"
	"net/http"
	"time"
)

// Se debe generar la estructura request

type request struct {
	Name 	 		string		`json:"name" binding:"required"`
	Colour	 		string		`json:"colour" binding:"required"`
	Price    		float64		`json:"price" binding:"required"`
	Stock	 		int			`json:"stock" binding:"required"`
	Code	 		string		`json:"code" binding:"required"`
	Published		bool		`json:"published" binding:"required"`
	CreationDate	time.Time	`json:"creationDate" binding:"required"`
}

// Se debe generar la estructura del controlador que tenga como campo el servicio
type Product struct {
	service products.Service
}

// Se debe generar la función que retorne el controlador
func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}

// Se deben generar todos los métodos correspondientes a los endpoints
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "789" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

func (c *Product) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "789" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "error en la request",
			})
			return
		}
	
		p, err := c.service.CreateProduct(req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "error al crear el producto"})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}