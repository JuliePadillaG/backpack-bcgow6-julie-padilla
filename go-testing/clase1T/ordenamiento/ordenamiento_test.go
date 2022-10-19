package ordenamiento

import (
	//"fmt"
	"testing"
	"github.com/stretchr/testify/assert" //se importa testify para los asserts
)

func TestOrdenamientoOK(t *testing.T) {
	//Arrange
	var1 := []int{1, 30, 4, 78, 40}
	esperado := []int{1, 4, 30, 40, 78}

	//Act
	resultado := Ordenar(var1)

	//Assert
	assert.Equal(t, esperado, resultado, "El resultado es distinto al esperado")
}

