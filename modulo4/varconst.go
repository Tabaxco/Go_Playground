package main

import "fmt"

func main() {

	//var + nome variavel + tipo
	var nome string
	nome = "bento"
	fmt.Println(nome)

	var idade int
	idade = 4
	fmt.Printf("Meu nome é %s e tenho %d anos.\n", nome, idade)

	var a = 2.33
	fmt.Printf("%.2f = %T\n", a, a)

	var b, j, c, d, e, f, g int = 12, 12, 12, 12, 12, 12, 12

	fmt.Println(b, j, c, d, e, f, g)

	var verdade = true
	fmt.Println(verdade)

	maca := "apple"
	fmt.Print(maca)

	numeral := 2
	fmt.Println(numeral)

	//constantes
	const idadeBento = 4
	fmt.Println(idadeBento)

}
