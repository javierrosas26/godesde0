package main

import (
	"fmt"

	"github.com/javierrosas26/godesde0/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(169)
	fmt.Println(estado)
	fmt.Println(texto)
}
