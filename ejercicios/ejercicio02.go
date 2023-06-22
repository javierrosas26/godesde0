package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TabladeMultiplicar() string {
	escanear(*bufio.NewScanner(os.Stdin))
	fmt.Println("la tabla del 1 al 10 del numero", numero, "es: ")
	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i*numero)
	}
	return texto
}

func escanear(scanner bufio.Scanner) {
	fmt.Println("Ingrese el numero ")
	for scanner.Scan() {
		fmt.Println("Ingrese el numero ")
		numero, err = strconv.Atoi(scanner.Text())
		if err == nil {
			break
		}
	}
}
