package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/julie-padilla/go-testing/clase2M/internal/products"
	"github.com/julie-padilla/go-testing/clase2M/pkg/web"
	//"time"
	"strconv"
	"os"
)

// Se debe generar la estructura request

type request struct {
	Name 	 		string		`json:"name" binding:"required"`
	Colour	 		string		`json:"colour" binding:"required"`
	Price    		float64		`json:"price" binding:"required"`
	Stock	 		int			`json:"stock" binding:"required"`
	Code	 		string		`json:"code" binding:"required"`
	Published		bool		`json:"published" binding:"required"`
	CreationDate	string		`json:"creationDate" binding:"required"`
}

type ProductRequestPatch struct {
	Nombre string `json:"nombre" binding:"required"`
	Price float64 `json:"price" binding:"required"`
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

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
         	return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
         	return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}


// CreateProducts godoc
// @Summary Create products
// @Tags Products
// @Description create products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to create"
// @Success 200 {object} web.Response
// @Router /products [post]
func (c *Product) CreateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
         	return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
         	return
		}
	
		p, err := c.service.CreateProduct(req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
         	return
		}
		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// CREACIÓN PUT PASO 4. Se agrega el controlador Update en el Handler de productos

// UpdateProduct godoc
// @Summary  Update product
// @Tags     Products
// @Accept   json
// @Produce  json
// @Param    id       path      int             true   "Id product"
// @Param    token    header    string          false  "Token"
// @Param    product  body      request  true   "Product to update"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  404      {object}  web.Response
// @Router   /products/{id} [PUT]
func (c *Product) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
         	return
		}

		id, err := strconv.ParseInt(ctx.Param("id"),10, 64)
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
         	return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
         	return
		}

		// Implementación de validaciones
		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}
		if req.Colour == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El color del producto es requerido"))
			return
		}
		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
			return
		}
		if req.Stock == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El stock del producto es requerido"))
			return
		}
		if req.Code == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El código del producto es requerido"))
			return
		}
		if req.Published == false {
			ctx.JSON(400, web.NewResponse(400, nil, "Es requerido informar si el producto está publicado o no"))
			return
		}
		if req.CreationDate == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "La fecha de creación del producto es requerida"))
			return
		}

		p, err := c.service.UpdateProduct(int(id), req.Name, req.Colour, req.Price, req.Stock, req.Code, req.Published, req.CreationDate)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
         	return

		}

		ctx.JSON(200, web.NewResponse(200, p, ""))
	}
}

// Delete Product
// @Summary  Delete product
// @Tags     Products
// @Param    id     path      int     true  "Product id"
// @Param    token  header    string  true  "Token"
// @Success  204
// @Failure  401    {object}  web.Response
// @Failure  400    {object}  web.Response
// @Failure  404    {object}  web.Response
// @Router   /products/{id} [DELETE]
func (c *Product) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
         	return
		}

		id, err := strconv.ParseInt(ctx.Param("id"),10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
         	return
		}
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
         	return
		}
		ctx.JSON(200, web.NewResponse(200, "El producto ha sido eliminado", ""))
	}
}

// Update NamePrice Product godoc
// @Summary      Update name price product
// @Tags         Products
// @Accept       json
// @Produce      json
// @Description  This endpoint update field name and file price from an product
// @Param        token  header    string               true  "Token header"
// @Param        id     path      int                  true  "Product Id"
// @Param        name   body      ProductRequestPatch  true  "Product name"
// @Param        price  body      ProductRequestPatch  true  "Product price"
// @Success      200    {object}  web.Response
// @Failure      401    {object}  web.Response
// @Failure      400    {object}  web.Response
// @Failure      404    {object}  web.Response
// @Router       /products/{id} [PATCH]
func (c *Product) UpdateNamePrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("token")

		if token != os.Getenv("TOKEN") {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
         	return
		}

		id, err := strconv.ParseInt(ctx.Param("id"),10, 64)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, "ID inválido"))
         	return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
         	return
		}

		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
         	return
		}
		if req.Price == 0.0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio del producto es requerido"))
         	return
		}
		p, err := c.service.UpdateNamePrice(int(id), req.Name, req.Price)
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
         	return
		}
		ctx.JSON(200, web.NewResponse(200, p, "")) 
	}
}