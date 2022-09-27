package ejercicios_clase2

import "errors"

func Promedio(calificaciones ...float64) (promedio float64, err error) {
	var suma float64
	for _, calificacion := range calificaciones {
		if calificacion < 0 {
			return 0, errors.New("El valor de la calificación no puede ser menor a cero")
		}
		suma += calificacion
		promedio = suma / float64(len(calificaciones))
	}
	return promedio, nil
}