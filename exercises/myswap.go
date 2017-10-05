package main

import "fmt"

func main() {
	var nome, cognome string
	nome = "Rachid"
	cognome = "Ellya"
	fmt.Printf("nome= %v, cognome= %v", nome, cognome)
	nome, cognome = swap(nome, cognome)
	fmt.Printf("nome= %v, cognome= %v", nome, cognome)

}

func swap(nome, cognome string) (string, string) {

	return cognome, nome
}
