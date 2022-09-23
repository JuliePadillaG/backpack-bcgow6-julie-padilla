package ejercicios

import "fmt"

var edad6 int = 23
var empleado bool = true
var antiguedad int = 2
var sueldo2 int = 102000

func Prestamos() {

	if edad6 > 22 && empleado && antiguedad > 1 && sueldo2 > 100000 {
		fmt.Printf("Se le puede prestar dinero a este cliente sin interés")
	} else if edad6 < 22 || !empleado || antiguedad < 1 {
		fmt.Printf("No se le puede prestar dinero a este cliente")
	} else {
		fmt.Printf("Debe ingresar los datos")
	}
}