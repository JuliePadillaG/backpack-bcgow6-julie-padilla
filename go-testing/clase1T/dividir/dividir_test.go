package dividir

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDividirBad(t *testing.T) {
	//Arrange
	num1:= 10
	num2 := 0
	errorEsperado := fmt.Sprintf("el denominador no puede ser: %d", num2)

	//Act
	_, err := Dividir(num1, num2)

	//Assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado) //se comprueba si falló como yo esperaba
}
	