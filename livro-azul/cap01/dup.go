package main

import (
	"bufio"
	"fmt"
	"os"
)

func main4() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// para validar entrada de dados
	if err := input.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro na leitura:", err)
	}

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
