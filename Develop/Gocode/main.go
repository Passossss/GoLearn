package main

import "fmt"

type Luis struct {
	Nome   string
	Idade  int
	Altura float64
	QI     int
}

func main() {
	fmt.Println("Bah guri start Goolang")
	Puc("luis gay")
	l := Luis{
		Nome:   "Luis",
		Idade:  25,
		Altura: 1.75,
		QI:     10,
	}
	fmt.Println(l)
}

// var materia string
// let materia = GEAD
//
// func (materia) teste
// {
//
// }
func Puc(material string) {
	fmt.Println("professor é maluco")
}
