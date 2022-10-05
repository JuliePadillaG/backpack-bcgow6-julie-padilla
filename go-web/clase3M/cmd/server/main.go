package main

// Se debe importar e inyectar el repositorio, servicio y handler
import (
	"github.com/gin-gonic/gin"
	"github.com/julie-padilla/go-web/clase3M/cmd/server/handler"
	"github.com/julie-padilla/go-web/clase3M/internal/products"
)

// Se debe implementar el router para los diferentes endpoints
 func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p:= handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.CreateProduct())

	//CREACIÓN PUT PASO 5. Definir un servicio web mediante el método PUT, el cual tendrá como path "products/:id".
	pr.PUT("/:id", p.UpdateProduct())

	r.Run()
 }