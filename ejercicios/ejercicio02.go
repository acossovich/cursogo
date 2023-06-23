package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func MakeNumericTable() {
	scanner := bufio.NewScanner(os.Stdin)

	var number int
	var err error

	for {
		fmt.Println("Ingrese un numero: ")
		if scanner.Scan() {
			number, err = strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Error, debe ser un numero entero")
				continue
			} else {
				break
			}

		} else {
			fmt.Println("Error")
			continue
		}
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(number, " x ", i, "= ", number*i)
	}

}
