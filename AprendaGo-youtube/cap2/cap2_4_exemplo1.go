package main

import (
	"fmt"
)

// a variavel abaixo é declarada ao nivel package level scope
var y = 11

func main () {


// a marmota so pode ser usada dentro de um code block
	a := 10
	imprimir(a)

}

func imprimir(x int) {
	fmt.Println(x)
	fmt.Println(y)
}
