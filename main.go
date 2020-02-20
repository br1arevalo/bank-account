package main

import "fmt"

var saldo, nombre = 0.0, "luis pelonso"
var op, co int
var dep, ret float64
var p1 producto

type producto struct {
	name string
	cost float64
}

func main() {
	cliente()
}

func cliente() {
	fmt.Println("Crea tu cuenta \nNombre completo:")
	fmt.Scan(&nombre)
	menu()
}

func compras() {
	fmt.Println("Bienvenido al menu de compras\nSelecciona una opcion:\nDigita 1 para crear un nuevo producto\nDigita 2 para comprar un producto\nDigita 9 para regresar al menu anterior")
	if p1.cost > 0 {
		fmt.Println("Producto disponible:", p1)
	} else {
		fmt.Println("Actualmente no existe ningun producto para comprar")
	}
	fmt.Scan(&co)
	switch {
	case co == 1:
		crear()
	case co == 2:
		if p1.cost <= 0 {
			fmt.Println("No existe producto para comprar, cree un nuevo producto")
			compras()
		}
		comprar()
	case co == 9:
		menu()
	default:
		fmt.Println("Opcion incorrecta")
		compras()
	}
}

func crear() producto {

	var cr int
	fmt.Println("Estas creando un nuevo producto, para continuar ingresa cualquier caracter, para salir digita 9")
	fmt.Scan(&cr)
	if cr == 9 {
		compras()
	} else {
		fmt.Println("Nombra el nuevo producto:")
		fmt.Scan(&p1.name)
		fmt.Println("Ponle un precio:")
		fmt.Scan(&p1.cost)
		if p1.cost <= 0 {
			fmt.Println("Solo puede tener valores mayores a cero")
			crear()
		}
		compras()
	}
	return p1
}

func comprar() {
	var qw int
	fmt.Println("Recuerda que para comprar debes de tener saldo suficiente en tu cuenta\nTu saldo es:", saldo, "\nProducto disponible:", p1, "\nPresiona 1 para comprar\nPresiona 9 para regresar")
	fmt.Scan(&qw)
	switch {
	case qw == 9:
		compras()
	case qw == 1:
		if saldo >= p1.cost {
			saldo = saldo - p1.cost
			fmt.Println("Compraste un", p1.name, "de:", p1.cost, "pesos\nTe quedan:", saldo, "pesos")
			comprar()
		}
		fmt.Println("No tienes saldo suficiente, ve a depositar")
		deposito(dep)
	}
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
	case op == 9:
		compras()
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
