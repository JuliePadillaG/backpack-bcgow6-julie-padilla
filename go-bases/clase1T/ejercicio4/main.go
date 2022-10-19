package main

import "fmt"

func main() {
	employees := map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	// Ayuda a imprimir la edad de Benjamin.
	fmt.Println("La edad de Benjamín es : ", employees["Benjamin"])

	// Saber cuántos de sus empleados son mayores de 21 años.
	var cantMayores21 int
	for _, value := range employees {
		if value > 21 {
			cantMayores21++
		}
	}
	fmt.Printf("Cantidad de empleados mayores de 21 = %d\n", cantMayores21)

	// Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	employees["Federico"] = 25
	fmt.Println(employees)

	// Eliminar a Pedro del mapa.
	delete(employees, "Pedro")
	fmt.Println(employees)
}
