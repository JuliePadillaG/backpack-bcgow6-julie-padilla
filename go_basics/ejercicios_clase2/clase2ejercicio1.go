package ejercicios_clase2

var salario float64 = 160000
var descuento float64 = 0.0

func Impuesto()float64 {
	if salario > 50000 {
		descuento = 0.17
		salario = salario - (salario * descuento)
	}
	if salario > 150000 { 
		descuento = 0.27
		salario = salario - (salario * descuento)
	}
	return salario
}