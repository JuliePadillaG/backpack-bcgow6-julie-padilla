package ordenamiento

import (
	"sort"
)

func Ordenar(var1 []int) []int {
	
	sort.Slice(var1, func(i, j int) bool {
		return var1[i] < var1[j]
	})

	return var1
}