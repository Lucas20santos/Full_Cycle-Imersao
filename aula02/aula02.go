package main

import (
	"fmt"
)

// função principal
func main()  {
	// Pacote + "." + indentificador
	var numeroBytes, erros = fmt.Println("Aula 02")
	fmt.Println(numeroBytes, erros)
	/*
		A função Println retorna dois valores, que são:
			- números de bytes
			- e um erro, se aconteceu
		Se colocar apenas uma variável, o programa deve reclamar com um erro,
		já que ele espera duas variável, mas pode ser usado um undeleiner no 
		lugar da variável
	*/
	var	_, erro2 = fmt.Println("testando")
	fmt.Println(erro2)

	x := 23
	y := "strings"
	z := true
	fmt.Println(x, y, z)

	var bytes ,erro3 = fmt.Println("Teste 3")
	fmt.Println(bytes, erro3)
}
