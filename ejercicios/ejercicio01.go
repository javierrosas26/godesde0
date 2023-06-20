package ejercicios

import (
	"strconv"
)

func Ejercicio1(numeroTxt string) (int, string) {
	numero, err := strconv.Atoi(numeroTxt)
	if err != nil {
		return 0, "error during conversion"
	}
	if numero > 100 {
		return int(numero), "Es mayor a 100"
	} else {
		return int(numero), "Es menor a 100"
	}

}
