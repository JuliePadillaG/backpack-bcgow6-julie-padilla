package main

import "fmt"

/*Ejercicio 1 - Imprimí tu nombre

Crea una aplicación donde tengas como variable tu nombre y dirección.
Imprime en consola el valor de cada una de las variables.
*/

func main() {

	var (
		nombre    = "Julie Padilla"
		direccion = "Bogotá D.C."
	)

	// Ejemplos para representar distinto tipos de datos en un string:
	// %s string
	// %d interger
	// %f float

	fmt.Printf("Mi nombre es %s y vivo en %s\n", nombre, direccion)
	fmt.Printf("Mi nombre es %v y vivo en %v\n", nombre, direccion)
}
