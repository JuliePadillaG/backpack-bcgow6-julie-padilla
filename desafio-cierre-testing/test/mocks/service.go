package mocks

import (
	"fmt"
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
)

type MockService struct {
	DataMock []model.Product
	Error string
}

func (d *MockService) GetAllBySeller(sellerID string) ([]model.Product, error) {
	// products := []Product{
	// 	{
	// 	ID: "123",
	// 	SellerID: "f4g5t5", //el data en la función del service devuelve el producto con este SellerID ¿?
	// 	Description: "producto xx",
	// 	Price: 80.00,
	// 	},
	// }

	if d.Error != "" {
		return nil, fmt.Errorf(d.Error)
	}

	return d.DataMock, nil
}