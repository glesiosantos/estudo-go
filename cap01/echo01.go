// Echo01 exibe seu argumentos de linha de comando
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	// fmt.Println(s)
	fmt.Println(strings.Join(s[1:0]))
}
