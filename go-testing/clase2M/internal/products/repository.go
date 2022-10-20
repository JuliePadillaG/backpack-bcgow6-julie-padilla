package products

import (
	//"time"
	"fmt"
	"github.com/julie-padilla/go-testing/clase2M/pkg/store"
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
	CreationDate	string	`json:"creationDate" binding:"required"`
}

// Paso 6. Se deben crear las variables globales donde guardar las entidades

var ps []Product
var lastID int

// Paso 7. Se debe generar la interface Repository con todos sus métodos

// CREACIÓN PUT PASO 1: Agregar los métodos a utilizar en las interfaces de Repository y Service.
type Repository interface {
	GetAll() ([]Product, error)
	LastID() (int, error)
	CreateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	UpdateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate string) (Product, error)
	Delete(id int) error
	UpdateNamePrice(id int, name string, price float64) (Product, error)
}

// Paso 8. Se debe generar la estructura repository
type repository struct {
	db store.Store
}

// Paso 9. Se debe generar una función que devuelva el Repositorio
func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// Paso 10. Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
func (r *repository) GetAll() (products []Product, err error){
	err = r.db.Read(&products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	err := r.db.Read(ps)
	if err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps) + 1].ID, nil
}

func (r *repository) CreateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	var ps []Product

	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}
	
	p := Product{id, name, colour, price, stock, code, published, creationDate}
	ps = append(ps, p)

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	return p, nil
}

//CREACIÓN PUT PASO 2: Implementar la funcionalidad

func (r *repository) UpdateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate string) (Product, error) {
	
	var ps []Product

	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}
	
	p := Product{Name: name, Colour: colour, Price: price, Stock: stock, Code: code, Published: published, CreationDate: creationDate}
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

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps []Product

	err := r.db.Read(&ps)
	if err != nil {
		return err
	}

	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	
	if !deleted {
		fmt.Println(index, id)
		return fmt.Errorf("Producto %d no encontrado", id)
	}
	
	ps = append(ps[:index], ps[index+1:]...)

	if err := r.db.Write(ps); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateNamePrice(id int, name string, price float64) (Product, error) {
	
	var ps []Product

	err := r.db.Read(&ps)
	if err != nil {
		return Product{}, err
	}

	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			ps[i].Price = price
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("Producto %d no encontrado", id)
	}

	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}

	return p, nil
 }
 
