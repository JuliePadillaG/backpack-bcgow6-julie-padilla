package products

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

type MockStore struct {}

func (d MockStore) GetAllBySeller(sellerID string) ([]Product, error) {
	products := []Product{
		{
		ID: "123",
		SellerID: "f4g5t5", //el data en la función del service devuelve el producto con este SellerID ¿?
		Description: "producto xx",
		Price: 80.00,
		},
	}

	//sellerID = "h3h21"

	return products, nil
}

//Test

func TestGetAllBySellerOK(t *testing.T) {
	// Arrange

	sellerID := "f4g5t5"

	myMock := MockStore{}
	service := NewService(myMock)
	esperado := []Product{
		{
		ID: "123",
		SellerID: "f4g5t5", //el data en la función del service devuelve el producto con este SellerID
		Description: "producto xx",
		Price: 80.00,
		},
	}

	// Act
	resultado, err := service.GetAllBySeller(sellerID)
	if err != nil {
		fmt.Println(err)
	}

	// Assert
	assert.Equal(t, esperado, resultado)
	assert.Nil(t, err)
}



