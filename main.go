package main

import "fmt"

var saldo, nombre = 0.0, "luis pelonso"
var op int
var dep, ret float64

func main() {
	cliente()
}

func cliente() {
	fmt.Println("Crea tu cuenta \nNombre completo:")
	fmt.Scan(&nombre)
	menu()
}

func deposito(float64) {
	fmt.Println("Haz un deposito a tu cuenta\nMonto a ingresar:")
	fmt.Scan(&dep)
	if dep < 0 {
		var try int
		fmt.Println("Debes depositar, para retirar dinero debes de ir al menu de operaciones y elegir hacer un retiro")
		fmt.Println("Que deseas hacer?\nPara volver a intentarlo presiona 1\nPara ir al menu de operaciones presiona 9")
		fmt.Scan(&try)
		switch {
		case try == 1:
			deposito(dep)
		case try == 9:
			menu()
		default:
			fmt.Println("Eleccion incorrecta, por pendejo te voy a llevar al menu")
			menu()
		}
	} else {
		saldo = dep + saldo
		fmt.Println("\n", nombre, "has depositado en tu cuenta:", dep, "\nAhora tienes:", saldo, "pesos en tu cuenta")
		volver()
	}
}

func menu() {
	fmt.Println("\nBienvenido", nombre, "estas en tu perfil bancario\nPara ir de compras digita 9\n\nOPERACIONES\n\nTu saldo actual es", saldo, "pesos\nDigite 1 para retirar\nDigite 2 para depositar\nDigite 3 para salir\nQue operacion desea realizar?")
	fmt.Scan(&op)
	switch {
	case op == 1:
		retirar(ret)
	case op == 2:
		deposito(saldo)
	case op == 3:

	default:
		fmt.Println("Operacion incorrecta, redireccionando al menu anterior")
		menu()
	}
}

func retirar(float64) {
	fmt.Println("Cuanto desea retirar?")
	fmt.Scan(&ret)
	if ret < 0 {
		tr := 0
		fmt.Println("Estas intentando retirar un numero negativo, es una doble negacion\nLo cual se convierte en un deposito\nPara depositar dinero debes de ir al menu de operaciones y elegir hacer un deposito")
		fmt.Println("Que deseas hacer?\nPara volver a intentarlo presiona 1\nPara ir al menu de operaciones presiona 9")
		fmt.Scan(&tr)
		switch {
		case tr == 1:
			retirar(ret)
		case tr == 9:
			menu()
		default:
			fmt.Println("Eleccion incorrecta, por pendejo te voy a llevar al menu")
			menu()
		}
	} else {
		if saldo-ret < 0 {
			fmt.Println("No tienes fondos suficientes, vuelve a intentar")
			retirar(ret)
		} else {
			saldo = saldo - ret
			fmt.Println("\n", nombre, "has retirado de tu cuenta:", ret, "\nAhora tienes:", saldo, "pesos en tu cuenta")
			volver()
		}
	}
}

func volver() {
	var vo string
	fmt.Println("\nPresiona cualquier tecla para volver al menu")
	fmt.Scan(&vo)
	menu()

}
