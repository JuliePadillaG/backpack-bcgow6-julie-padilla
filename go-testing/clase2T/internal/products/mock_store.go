package products

// Paso 1. Crear un mock de Store, dicho mock debe contener en su data un producto con las especificaciones que desee.

type MockStorage struct {
	dataMock Product
	//errOnWrite error
	errOnRead error
}

// Paso 2. El método Read del Mock, debe contener una lógica que permita comprobar que dicho método fue invocado. 
func (m *MockStorage) Read(data interface{}) (err error) {

	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*Product)
	*castedData = m.dataMock
	return nil
}