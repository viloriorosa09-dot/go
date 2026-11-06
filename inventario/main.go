package main

import (
	"fmt"
)

type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func valorTotal(inventario map[string]Productos) float64 {

	var total float64
	for _, v := range inventario {
		suma := float64(v.Cantidad) * v.Precio
		total += suma
	}

	fmt.Println("El valor total del inventario es: ", total)
	return total
}

func bajoStock(inventario map[string]Productos, minimo int) []string {
	stock := make([]string, 0)

	for _, cantidad := range inventario {
		if cantidad.Cantidad < minimo {
			stock = append(stock, cantidad.Nombre)

		}

	}
	fmt.Print("Productos con bajo stock: ", stock)
	return stock
}

func main() {
	monitor := Productos{Nombre: "Monitor", Precio: 1000, Cantidad: 4}
	teclado := Productos{Nombre: "Teclado", Precio: 300, Cantidad: 5}
	procesador := Productos{Nombre: "Procesador", Precio: 5000, Cantidad: 1}
	discoDeAlmacenamiento := Productos{Nombre: "Disco de almacenamiento", Precio: 800, Cantidad: 2}

	var productos = map[string]Productos{

		"A1": monitor,
		"B2": teclado,
		"C3": procesador,
		"D4": discoDeAlmacenamiento,
	}
	valorTotal(productos)
	bajoStock(productos, 5)
}
