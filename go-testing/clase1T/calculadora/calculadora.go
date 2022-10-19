package calculadora

import "fmt"

func Suma(num1, num2 int) (int, error) {
	if num1 == 0 {
		return 0, fmt.Errorf("a no puede ser: %d", num1)
	}
	return num1 + num2, nil
}

func Resta(num1, num2 int) int {
	return num1 - num2
}