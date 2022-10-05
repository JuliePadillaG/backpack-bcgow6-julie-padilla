package products

import "time"

// Paso 1. Se debe generar la interface Service con todos sus métodos.

// CREACIÓN PUT PASO 1: Agregar los métodos a utilizar en las interfaces de Repository y Service.
type Service interface {
	GetAll() ([]Product, error)
	CreateProduct(name, color string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error)
	UpdateProduct(id int, name, color string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error)
	Delete(id int) error
	UpdateNamePrice(id int, name string, price float64) (Product, error)
}

// Paso 2. Se debe generar la estructura service que contenga el repositorio.
type service struct {
	repository Repository
}

// Paso 3. Se debe generar una función que devuelva el Servicio.
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Paso 4. Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) CreateProduct(name, colour string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	product, err := s.repository.CreateProduct(lastID, name, colour, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

// CREACIÓN PUT PASO 3. Dentro del servicio se llama al repositorio para que proceda a actualizar el producto

func (s *service) UpdateProduct(id int, name, colour string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error) {
	return s.repository.UpdateProduct(id, name, colour, price, stock, code, published, creationDate)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNamePrice(id int, name string, price float64) (Product, error) {

	return s.repository.UpdateNamePrice(id, name, price)
 }
 