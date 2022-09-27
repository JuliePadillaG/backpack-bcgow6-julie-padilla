package ejercicios_clase3

// Una empresa que se encarga de vender productos de limpieza necesita: 
// Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
// Debe tener el id del producto, precio y la cantidad.
// Estos valores pueden ser hardcodeados o escritos en duro en una variable.
// ID                            Precio  Cantidad
// 111223                      30012.00         1
// 444321                    1000000.00         4
// 434321                         50.50         1
//                           4030062.50


var idProduct string = "11111"


product := []byte(idProduct)

err := os.WriteFile("./ejercicios_clase3/products.csv", product, 0644)
