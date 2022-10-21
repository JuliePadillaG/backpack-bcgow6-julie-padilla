package products

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

type StubDB struct{}

func (d StubDB) Read(data interface{}) error {
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

    //solución con punteros
	a := data.(*[]Product)
	*a = append(*a, products...)

	//otra opción:
	//json marshal recibe interfaz y pasa a bytes
	//json unmarshal pasa al tipo de dato que uno quiera
	
	return nil
}

func (d StubDB) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {

	//arrange
	myStubDB := StubDB{}
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
    product, err := repository.GetAll()
	if err != nil {
		fmt.Println("error")
	}

	//assert
	assert.Equal(t, productoEsperado, product)
}

func TestUpdate(t *testing.T) {
    
}