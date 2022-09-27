package main

import (
	// "encoding/csv"
	// "os"
	// "fmt"
)
// import "github.com/JuliePadillaG/backpack-bcgow6-julie-padilla/ejercicios_clase2"

// type Students struct {
// 	Name string
// 	LastName string
// 	DNI string
// 	Date string
// }

// func (s Students) Detail() Students{
// 	return s
// }

// type Product struct {
// 	ID string
// 	Precio string
// 	Cantidad string
// }

func main() {

	// student := Students{
	// Name: "Julie",
	// LastName: "Padilla",
	// DNI: "52998776",
	// Date: "26/09/2022",
	// }

	// fmt.Println(student.Detail())

	//ejercicios.Nombre()
	//ejercicios.Meteorologia()
	//ejercicios.Deletrear()
	//ejercicios.Prestamos()
	//ejercicios.Meses()
	//ejercicios.MapEmpleado()
	//fmt.Println(ejercicios_clase2.Impuesto())

	// promedio, err := ejercicios_clase2.Promedio(3.5, 0.9)

	// if err != nil {
	// 	fmt.Println("El valor de la calificación no puede ser menor que cero")
	// } else {
	// 	fmt.Println("El promedio de notas es: ", promedio)
	// }

	// Una empresa que se encarga de vender productos de limpieza necesita: 
	// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
	// Debe tener el id del producto, precio y la cantidad.
	// Estos valores pueden ser hardcodeados o escritos en duro en una variable.
	// ID                            Precio  Cantidad
	// 111223                      30012.00         1
	// 444321                    1000000.00         4
	// 434321                         50.50         1
	//                           4030062.50

	// products := []Product{
	// 	{"111223", "30012.00", "1"},
	// 	{"444321", "1000000.00", "4"},
	// 	{"434321", "50.50", "1"},
	// }

	// file, err := os.Create("./ejercicios_clase3/products.csv")
	// defer file.Close()
	// if err != nil {
	// 	panic (err)
	// }

	// w := csv.NewWriter(file)
	// defer w.Flush()

	// for _, product := range products {
	// 	row := []string{product.ID, product.Precio, product.Cantidad}
	// 	if err := w.Write(row); err != nil {
	// 		panic(err)
	// 	}
	// }

	// La misma empresa necesita leer el archivo almacenado, para ello requiere que: se imprima por pantalla mostrando los valores tabulados, con un título (tabulado a la izquierda para el ID y a la derecha para el Precio y Cantidad), el precio, la cantidad y abajo del precio se debe visualizar el total (Sumando precio por cantidad)

	// data, err := os.ReadFile("./ejercicios_clase3/products.csv")

	// fmt.Printf("%s", data)

	// newFile, err := os.Open("./ejercicios_clase3/products.csv")
	// if err != nil {
	// 	panic (err)
	// }

	// reader := csv.NewReader(newFile)
	// products1,_ := reader.ReadAll()

	// fmt.Println(products1)



}

