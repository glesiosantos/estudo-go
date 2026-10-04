// Echo02 exibe seu argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main2() {
	var s, sep string

	for i, arg := range os.Args[1:] {
		fmt.Printf("Índice: %d | Argumento: %s\n", i, arg)
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
