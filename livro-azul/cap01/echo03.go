package main

import (
	"fmt"
	"os"
	"strings"
)

func main3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
