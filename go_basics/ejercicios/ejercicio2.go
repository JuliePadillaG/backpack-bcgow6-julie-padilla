package ejercicios

import "fmt"

var temperatura int = 0
var humedad int = 0
var presion int = 0

func Meteorologia() {
	temperatura = 23
	humedad = 10
	presion = 220
	fmt.Printf("temperatura: \n%v\n", temperatura)
	fmt.Printf("humedad: \n%v\n", humedad)
	fmt.Printf("presion: \n%v\n", presion)
}