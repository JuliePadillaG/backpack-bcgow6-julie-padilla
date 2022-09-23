package ejercicios

import "fmt"

var nombre string = ""
var direccion string = ""

func Nombre() {
	nombre = "Julie"
	direccion = "Bogotá"
	fmt.Printf("nombre: \n%v\n", nombre)
	fmt.Printf("direccion: \n%v\n",direccion)

	Meteorologia()
}



