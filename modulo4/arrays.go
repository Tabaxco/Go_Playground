package main

import "fmt"

// 1 - Arrays e Slices: Homogêneos
// todos os elementos tem o mesmo tipo

//Array:
//tamanho fixo
//acessamos os valores pelos indices
//função embutida len() retorna o tamanho do array
//não é muito usado devido ao tamanho fixo

func main() {

	/*var nomeAlunos [2]string
	nomeAlunos[0] = "Arnaldo"
	nomeAlunos[1] = "Geraldo"

	var testeTipo = nomeAlunos

	fmt.Println(nomeAlunos[0], nomeAlunos[1])
	fmt.Printf("%T, %v\n", testeTipo, testeTipo)

	numPrimos := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[1:])*/

	//nomesCachorro := make([]string, 5)
	var nomesCachorro []string

	//nomesCachorro[0] = "Belinha"
	//nomesCachorro[1] = "Rex"

	nomesCachorro = append(nomesCachorro, "Belinha")
	nomesCachorro = append(nomesCachorro, "Rex")
	nomesCachorro = append(nomesCachorro, "Bolt")

	fmt.Println(nomesCachorro)
	//nomesCachorro = append(nomesCachorro, "Bolt")
	//fmt.Println(nomesCachorro[5])

	numPares := []int{2, 3, 4, 5, 6, 7, 8}
	fmt.Println(numPares)
	numPares = append(numPares, 9, 10, 12, 16, 18)

}

//Slice:
//Não tem tamanho fixo
//acessa os valores pelos indices
//função len()
//função appenda para adicionar valores
//muito semelhante a lists em Python

// 2 - Maps: Heterogêneos
// pode misturar tipos
