package main

// Se debe importar e inyectar el repositorio, servicio y handler
import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/julie-padilla/go-testing/clase2T/cmd/server/handler"
	"github.com/julie-padilla/go-testing/clase2T/internal/products"
	"github.com/julie-padilla/go-testing/clase2T/pkg/store"
	"github.com/julie-padilla/go-testing/clase2T/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"os"
)

// Se debe implementar el router para los diferentes endpoints

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
 func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	// usuario := os.Getenv("MY_USER")
	// password := os.Getenv("MY_PASS")
	 
	db := store.New(store.FileType, "products.json")

	repo := products.NewRepository(db)
	service := products.NewService(repo)

	p:= handler.NewProduct(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
   	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.CreateProduct())

	//CREACIÓN PUT PASO 5. Definir un servicio web mediante el método PUT, el cual tendrá como path "products/:id".
	pr.PUT("/:id", p.UpdateProduct())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateNamePrice())

	err = r.Run()
	if err!= nil {
		fmt.Println("error")
	}
 }