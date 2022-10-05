package products

import (
	"time"
	"fmt"
)

// Paso 5. Se debe crear la estructura de la entidad
type Product struct {
	ID		 		int			`json:"id"`
	Name 	 		string		`json:"name" binding:"required"`
	Colour	 		string		`json:"colour" binding:"required"`
	Price    		float64		`json:"price" binding:"required"`
	Stock	 		int			`json:"stock" binding:"required"`
	Code	 		string		`json:"code" binding:"required"`
	Published		bool		`json:"published" binding:"required"`
	CreationDate	time.Time	`json:"creationDate" binding:"required"`
}

// Paso 6. Se deben crear las variables globales donde guardar las entidades

var ps []Product
var lastID int

// Paso 7. Se debe generar la interface Repository con todos sus métodos

// CREACIÓN PUT PASO 1: Agregar los métodos a utilizar en las interfaces de Repository y Service.
type Repository interface {
	GetAll() ([]Product, error)
	LastID() (int, error)
	CreateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error)
	UpdateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error)
}

// Paso 8. Se debe generar la estructura repository
type repository struct {

}

// Paso 9. Se debe generar una función que devuelva el Repositorio
func NewRepository() Repository {
	return &repository{}
}

// Paso 10. Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
func (r *repository) GetAll() ([]Product, error){
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) CreateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error) {
	p := Product{id, name, colour, price, stock, code, published, creationDate}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}

//CREACIÓN PUT PASO 2: Implementar la funcionalidad

func (r *repository) UpdateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error) {
	p := Product{id, name, colour, price, stock, code, published, creationDate}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}

