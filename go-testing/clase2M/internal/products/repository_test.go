package products

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

// Ejercicio 1 - Test Unitario GetAll()

type StubStore struct{}

func (d StubStore) Read(data interface{}) error {
	products := []Product{
		{
            ID: 1,
            Name: "pantalon",
            Colour: "cafe",
            Price: 50000,
            Stock: 40,
            Code: "a1b2c3",
            Published: true,
            CreationDate: "2022-10-03T13:40:19.087Z",
        },
        {
            ID: 2,
            Name: "camiseta",
            Colour: "blanca",
            Price: 20000,
            Stock: 40,
            Code: "d1e2f3",
            Published: true,
            CreationDate: "2022-10-03T13:40:19.087Z",
        },
	}

	a := data.(*[]Product)
	*a = append(*a, products...)

	//otra opción:
	//json marshal recibe interfaz y pasa a bytes
	//json unmarshal pasa al tipo de dato que uno quiera
	
	return nil
}

func (d StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	//arrange
	myStubDB := StubStore{}
	repository := NewRepository(myStubDB)

	productoEsperado := []Product{
		{
            ID: 1,
            Name: "pantalon",
            Colour: "cafe",
            Price: 50000,
            Stock: 40,
            Code: "a1b2c3",
            Published: true,
            CreationDate: "2022-10-03T13:40:19.087Z",
        },
        {
            ID: 2,
            Name: "camiseta",
            Colour: "blanca",
            Price: 20000,
            Stock: 40,
            Code: "d1e2f3",
            Published: true,
            CreationDate: "2022-10-03T13:40:19.087Z",
        },
	}

	//act
	productoResultado, err := repository.GetAll()
	if err != nil {
		fmt.Println("error")
	}

	//assert
	assert.Equal(t, productoEsperado, productoResultado)
    assert.Nil(t, err)
}

// Ejercicio 2 - Test Unitario UpdateName()

// Se crea un struct para el mock con las variables
type MockStore struct {
    ReadInvoked bool // Variable bandera, se crea con su valor por defecto de false
    Data        []Product
}

func (d *MockStore) Read(data interface{}) error {
    d.ReadInvoked = true // Si se ingresa a la función la bandera cambia a true
    a := data.(*[]Product) // "Parsea" el data a un puntero de slice de productos
	*a = d.Data // Se le asigna el slice
    return nil
}

func (d *MockStore) Write(data interface{}) error {
    return nil
}

// Tests Update
func TestUpdateProduct(t *testing.T) {

    // Arrange

    id, name, colour, price, stock, code, published, creationDate := 1, "Name After Update", "blue", 30.0, 10, "a1b2c3", true, "2022-10-03T13:40:19.087Z"

    productBeforeUpdate := []Product{
		{
            ID: 1,
            Name: "Name Before Update",
            Colour: "green",
            Price: 50.0,
            Stock: 40,
            Code: "a1b2c3",
            Published: true,
            CreationDate: "2022-10-03T13:40:19.087Z",
        },
    }

    
    mock := MockStore{Data: productBeforeUpdate}

    // Act

    repository := NewRepository(&mock)
    updatedProduct, err := repository.UpdateProduct(id, name, colour, price, stock, code, published, creationDate)

    // Assert

    assert.Equal(t, id, updatedProduct.ID)
    assert.Equal(t, name, updatedProduct.Name)
    assert.Equal(t, colour, updatedProduct.Colour)
    assert.Equal(t, price, updatedProduct.Price)
    assert.Equal(t, stock, updatedProduct.Stock)
    assert.Equal(t, code, updatedProduct.Code)
    assert.Equal(t, published, updatedProduct.Published)
    assert.Equal(t, creationDate, updatedProduct.CreationDate)
   
    assert.Nil(t, err)
    assert.True(t, true, mock.ReadInvoked)
}
