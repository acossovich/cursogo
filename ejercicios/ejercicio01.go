package ejercicios

import "strconv"

func ArrayToInteger(cadena string) (int, string) {
	integer, ok := strconv.Atoi(cadena)

	if ok != nil {
		return 0, "Error"
	}

	if integer > 100 {
		return integer, "Es mayor a 100"
	} else {
		return integer, "Es menor a 100"
	}
}
