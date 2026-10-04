// Echo02 exibe seu argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for _, arg := r	 {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
