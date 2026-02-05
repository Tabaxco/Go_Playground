package main

import "fmt"

//Structs
//Forma de criar a sua própria estrutura de dados

// type <nome estrutura> struct { <campos> }

type Pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	anoNascimento int
}

type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Guilherme", "ML", 18, 2007})
	fmt.Println(Pessoa{Nome: "Jonas", Sobrenome: "Alex", Idade: 19, anoNascimento: 2006})
	fmt.Println(Pessoa{Idade: 20, Sobrenome: "Alexandro"})

	p1 := Pessoa{Nome: "Elias"}
	fmt.Println(p1)
	fmt.Println(p1.Nome)
	p1.Idade = 13
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Patrick", Idade: 2}
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	p3 := Pessoa{Nome: "Pablo"}
	pessoas = append(pessoas, p3)
	fmt.Println(pessoas)

	alunos := map[string][]Pessoa{}
	alunos["programação"] = pessoas
	alunos["historia"] = pessoas
	fmt.Println(alunos)
	fmt.Println(alunos["historia"])

	var aluninhos = map[string][]Pessoa{
		"Programação": {
			{Nome: "Guilherme", Idade: 18},
			{Nome: "Nycollas", Idade: 20},
		},
		"Engenharia": {
			{Nome: "Alice", Idade: 13},
		},
	}
	fmt.Println(aluninhos)

	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Tipo)
}
