package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/julie-padilla/go-web/clase3M/internal/products"
	"net/http"
	"time"
	"strconv"
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

// CREACIÓN PUT PASO 4. Se agrega el controlador Update en el Handler de productos
func (c *Product) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "789" {
			ctx.JSON(401, gin.H{ "error": "token inválido" })
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"),10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{ "error": "invalid ID"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{ "error": err.Error() })
			return
		}
		if req.Name == "" {
			ctx.JSON(400, gin.H{ "error": "El nombre del producto es requerido"})
			return
		}
		if req.Colour == "" {
			ctx.JSON(400, gin.H{ "error": "El color del producto es requerido"})
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, gin.H{ "error": "El precio del producto es requerido"})
			return
		}
		if req.Stock == 0 {
			ctx.JSON(400, gin.H{ "error": "El stock del producto es requerido"})
			return
		}
		if req.Code == "" {
			ctx.JSON(400, gin.H{ "error": "El código del producto es requerido"})
			return
		}
		if req.Published == false {
			ctx.JSON(400, gin.H{ "error": "El estado de publicado del producto es requerido"})
			return
		}
		if req.CreationDate.IsZero() {
			ctx.JSON(400, gin.H{ "error": "La fecha de creación del producto es requerido"})
			return
		}
		p, err := c.service.UpdateProduct(int(id), req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(404, gin.H{ "error": err.Error() })
			return
		}
		ctx.JSON(200, p)
	}
}