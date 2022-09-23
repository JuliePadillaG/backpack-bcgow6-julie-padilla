package ejercicios

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func MapEmpleado() {
	fmt.Printf("Edad de Benjamín: \n%v\n" , employees["Benjamin"])

	//Saber cuántos de sus empleados son mayores de 21 años.
	count := 0
	for key, element := range employees {
		if element >= 21 {
			count += 1
		}
	fmt.Println("Key:", key, "=>", "Element:", element)
	}
	fmt.Println(count, "empleados son mayores de 21 años")

	//Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	employees["Federico"] = 25
	fmt.Println(employees)

	//Eliminar a Pedro del mapa.
	delete(employees, "Pedro")
	fmt.Println(employees)
}

