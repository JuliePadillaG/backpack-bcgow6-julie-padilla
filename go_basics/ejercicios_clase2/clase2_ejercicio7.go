package ejercicios_clase2

import "fmt"

// Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una estructura que represente una matriz de datos.
// Para ello requieren una estructura Matrix que tenga los métodos:
// Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
// Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
// La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, si es cuadrática y cuál es el valor máximo.

type Matrix struct {
	valores []int
	dimensionAlto int
	dimensionAncho int
	cuadratica bool
	valorMaximo int
}

func (m Matrix) Set(){
	fmt.Println(m)
}

func (m Matrix) Print(){
	fmt.Println(m)
}