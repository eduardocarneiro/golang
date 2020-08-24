package main 

import (
	"fmt"
)

// declaracao
var x int

func main() {
	//inicializacao e atribuicao
	x = 10
	fmt.Printf("%v, %T\n", x, x)
	//atribuicao
	x = 20
	fmt.Printf("%v, %T\n", x, x)

	// operador curto
	y := 30
	fmt.Printf("%v, %T\n", y, y)
}
