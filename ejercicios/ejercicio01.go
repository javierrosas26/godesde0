package ejercicios

import (
	"strconv"
)

func Ejercicio1(numeroTxt string) (int, string) {
	numero, _ := strconv.Atoi(numeroTxt)
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}

}
