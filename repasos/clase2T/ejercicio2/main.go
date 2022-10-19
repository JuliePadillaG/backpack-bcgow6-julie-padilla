package main

import "fmt"

// Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

// Si es categoría C, su salario es de $1.000 por hora
// Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
// Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

// Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.

func calculoSalario(minutos int, categoria string)float64 {
	horasTrabajadas := minutos/60
	switch categoria {
	case "A":
		salario := float64(horasTrabajadas * 3000) + ((float64(horasTrabajadas * 3000)) * 0.5)
		return salario
	case "B":
		salario := float64(horasTrabajadas * 1500) + ((float64(horasTrabajadas * 1500)) * 0.2)
		return salario
	case "C":
		salario := float64(horasTrabajadas * 1000)
		return salario
	default:
		return 0
	}
}

func main() {

	fmt.Println("El salario del empleado es:" , calculoSalario(480, "A"))
}
