package calculadora

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert" //se importa testify para los asserts
)

func TestSumaOK(t *testing.T) { //puntero a la estructura de testing que le da el contexto a la funmcion para que go la interprete como test y ahí están los métodos de validación del test
	
	//Arrange: valores de entrada y qupe estoy esperando que devuelva
	num1 := 10
	num2:= 5
	esperado := 15

	//Act: Llamar a la función
	resultado, err := Suma(num1, num2)

	//Assert: lo que se espera
	assert.Equal(t, esperado, resultado, "El resultado es distinto al esperado")
	assert.Nil(t, err) //prueba un caso en que la función funcione correctamente
	// if resultado != esperado {
	// 	t.Errorf("El número resultado: %d es distinto del esperado: %d", resultado, esperado)
	// }
}

func TestSumaBad(t *testing.T) {
	//Arrange
	num1:= 0
	num2 := 5
	errorEsperado := fmt.Sprintf("a no puede ser: %d", num1)

	//Act
	_, err := Suma(num1, num2)

	//Assert
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errorEsperado) //se comprueba si falló como yo esperaba
}

//si se van a meter más casos de prueba en una sola función, se arma un slice de casos de uso y se va recorriendo para obtener el resultado

//EJERCICIO 1: Para el método Restar() visto en la clase, realizar el test unitario correspondiente.
func TestRestaOK(t * testing.T) {

	//Arrange
	num1 := 10
	num2 := 2
	esperado := 8

	//Act
	resultado := Resta(num1, num2)

	//Assert
	assert.Equal(t, esperado, resultado, "El número resultado: %d es distinto del esperado: %d", resultado, esperado)

}