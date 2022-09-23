package ejercicios

import "fmt"

var palabra string = "afc"
var sum int = 0

func Deletrear() {
	fmt.Printf("Longitud de la palabra: %d\n", len(palabra))
	for i := 0; i < len(palabra); i++ {
		fmt.Printf("letra: %c\n", palabra[i])
	}
}