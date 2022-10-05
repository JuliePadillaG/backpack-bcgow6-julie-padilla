package products

import "time"

// Paso 1. Se debe generar la interface Service con todos sus métodos.
type Service interface {
	GetAll() ([]Product, error)
	CreateProduct(name, color string, price float64, stock int, code string, published bool, creationDate time.Time) (Product, error)
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