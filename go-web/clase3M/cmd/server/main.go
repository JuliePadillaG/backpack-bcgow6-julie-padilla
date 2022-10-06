package main

// Se debe importar e inyectar el repositorio, servicio y handler
import (
	"github.com/joho/godotenv"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/julie-padilla/go-web/clase3M/cmd/server/handler"
	"github.com/julie-padilla/go-web/clase3M/internal/products"
	"github.com/julie-padilla/go-web/clase3M/pkg/store"
)

// Se debe implementar el router para los diferentes endpoints
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

	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.CreateProduct())

	//CREACIÓN PUT PASO 5. Definir un servicio web mediante el método PUT, el cual tendrá como path "products/:id".
	pr.PUT("/:id", p.UpdateProduct())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateNamePrice())

	r.Run()
 }