package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func PedirNumero() {
	escanear(*bufio.NewScanner(os.Stdin))
	fmt.Println("la tabla del 1 al 10 del numero", numero, "es: ")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d \n", numero, i, i*numero)
	}
}

func escanear(scanner bufio.Scanner) {
	for {
		fmt.Println("Ingrese el numero ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err == nil {
				break
			}
		}

	}
}
