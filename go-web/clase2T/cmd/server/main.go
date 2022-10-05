package main

// Se debe importar e inyectar el repositorio, servicio y handler
import (
	"github.com/gin-gonic/gin"
	"github.com/julie-padilla/go-web/clase2T/cmd/server/handler"
	"github.com/julie-padilla/go-web/clase2T/internal/products"
)

// Se debe implementar el router para los diferentes endpoints
 func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p:= handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/create", p.CreateProduct())

	r.Run()
 }