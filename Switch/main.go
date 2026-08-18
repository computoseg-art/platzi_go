package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Escribe ", i, " como")

	switch i {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Es fin de semana")
	default:
		fmt.Println("Es un día de semana")
	}

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Lunes")
	case time.Tuesday:
		fmt.Println("Martes")
	case time.Wednesday:
		fmt.Println("Miercoles")
	case time.Thursday:
		fmt.Println("Jueves")
	case time.Friday:
		fmt.Println("Viernes")
	case time.Saturday:
		fmt.Println("Sabado")
	default:
		fmt.Println("Domingo")
	}
}