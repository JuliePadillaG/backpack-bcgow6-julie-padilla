package main

import "fmt"

/* Una empresa de meteorología quiere tener una aplicación donde pueda tener la temperatura y humedad y
presión atmosférica de distintos lugares.
	1. Declara 3 variables especificando el tipo de dato, como valor deben tener la temperatura,
	humedad y presión de donde te encuentres.
	2. Imprime los valores de las variables en consola.
	3. ¿Qué tipo de dato le asignarías a las variables? */

func main() {
	var temperatura, humedad, presion float32 = 16, 80, 1000.2

	fmt.Println("Reporte del clima: ")
	fmt.Printf("Temperatura: %f C°\n", temperatura)
	fmt.Printf("Humedad: %.0f %%\n", humedad)                
	fmt.Printf("Presión Atmosferica: %.2f hPa\n", presion)
}
