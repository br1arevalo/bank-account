package main

import "fmt"

var saldo, nombre, op = 0.0, "luis pelonso", 0

func main() {
	cliente()
}

//Crear un cliente que contenga nombre y saldo, el saldo por defecto de un cliente es de 0.0
func cliente() {
	fmt.Println("Crea tu cuenta \nNombre completo:")
	fmt.Scan(&nombre)
	fmt.Println("Bienvenido", nombre, "\nQue operacion quiere realizar?")
	fmt.Scan(&op)
	deposito(saldo)
}

func deposito(float64) {
	fmt.Println("Haz un deposito a tu cuenta\nMonto a ingresar:")
	fmt.Scan(&saldo)
	if saldo < 0 {
		fmt.Println("Debes depositar, no retirar dinero")
		deposito(saldo)
	} else {
		fmt.Println(nombre, "has depositado en tu cuenta:", saldo)
	}
}
