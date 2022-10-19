package main

import "fmt"

/* Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.
Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17% del sueldo y
si gana más de $150.000 se le descontará además un 10%.
*/

func calcularImpuesto(salario float64) float64 {
	switch {
	case salario > 150000:
			return salario * 0.27
	case salario > 50000:
			return salario * 0.17
	default:
		return 0
	}
}

func main() {

	var salario float64 = 150000

	fmt.Printf("El impuesto sobre su suelto es de: %.2f\n", calcularImpuesto(salario))
}
