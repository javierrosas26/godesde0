package ejercicios

import (
	"strconv"
)

func Ejercicio1(numeroTxt string) (int, string) {
	numero, err := strconv.Atoi(numeroTxt)
	if err != nil {
		return 0, "error in conversion" + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}

}
