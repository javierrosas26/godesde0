package main

import (
	"fmt"

	"github.com/javierrosas26/godesde0/goroutines"
)

func main() {
	/*estado, texto := variables.ConviertoaTexto(169)
	fmt.Println(estado)
	fmt.Println(texto)
	if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es windows, es ", os)
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("esto es Darwin")
	default:
		fmt.Printf("%s \n", os)

	}

	numero, texto := ejercicios.Ejercicio1("190")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros()

	iteraciones.Iterar()

	fmt.Println(ejercicios.TabladeMultiplicar())

	//files.GrabaTabla()
	//files.SumaTabla()

	//files.LeoArchivo()

	Pedro := new(modelos.Hombre)
	Maria := new(modelos.Mujer)

	e.HumanosRespirando(Pedro)
	e.HumanosRespirando(Maria)

	d.EjemploPanic()*/

	go goroutines.MiNombreLentooo("Javier Rosas")

	fmt.Println("Estoy aquí")
	var x string
	fmt.Scanln(&x)

	fmt.Println(x)

}
