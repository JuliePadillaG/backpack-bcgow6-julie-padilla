package dividir

import "fmt"

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("el denominador no puede ser: %d", den)
	}
	return num / den, nil
}