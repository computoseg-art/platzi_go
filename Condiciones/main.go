package main

import "fmt"

func main() {
	nombre := "Admin"
	edad := 40
	
	if nombre == "Admin" && edad >= 18 {
		fmt.Println("Acceso permitido")
	} else {
		fmt.Println("Acceso denegado")
	}

	if 8%2 == 0 {
		fmt.Println("El número 8 es par")
	}else{
		fmt.Println("El número 8 es impar")
	}

	if numero :=9; numero > 0 {
		fmt.Println("El número es positivo")
	}else if numero < 0 {
		fmt.Println("El número es negativo")
	}else{
		fmt.Println("El número es cero")
	}
}