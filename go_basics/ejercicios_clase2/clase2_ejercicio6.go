package ejercicios_clase2

// import "fmt"

type Students struct {
	Name string
	LastName string
	DNI string
	Date string
}

func (s Students) Detail() Students{
	return s
}