# Ejercicio: Inventario de tienda

Fecha: 5 de noviembre de 2026

## Qué pedía el ejercicio

Hacer un programa que guarde productos en un `map`, calcule el valor total del inventario y diga cuáles productos tienen poco stock.

## Qué hice

1. Definí un struct `Productos` con `Nombre`, `Precio` y `Cantidad`.
2. Creé 4 productos y los guardé en un `map[string]Productos`, usando el código del producto como clave.
3. Escribí la función `valorTotal`, que suma `Precio * Cantidad` de cada producto.
4. Escribí la función `bajoStock`, que devuelve un slice con los nombres de los productos que tienen menos del mínimo.
5. En `main` llamé a las dos funciones.

## Código completo

```go
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
```

## Explicación detallada

### El struct

```go
type Productos struct {
	Nombre   string
	Precio   float64
	Cantidad int
}
```

- Agrupa los tres datos de un producto en un solo tipo.
- `Precio` es `float64` porque puede tener decimales.
- `Cantidad` es `int` porque se cuenta en unidades enteras.

### La función valorTotal

```go
func valorTotal(inventario map[string]Productos) float64 {
	var total float64
	for _, v := range inventario {
		suma := float64(v.Cantidad) * v.Precio
		total += suma
	}
	fmt.Println("El valor total del inventario es: ", total)
	return total
}
```

- Recibe el map completo y devuelve un `float64`.
- `var total float64` crea el acumulador en cero.
- `for _, v := range inventario` recorre el map. Uso `_` para ignorar la clave porque solo necesito el producto (`v`).
- `float64(v.Cantidad)` convierte la cantidad de `int` a `float64`. Go no permite multiplicar tipos distintos sin convertir.
- `total += suma` va sumando el valor de cada producto.
- Imprime el resultado y también lo devuelve con `return`.

### La función bajoStock

```go
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
```

- Recibe el map y el número mínimo, y devuelve un slice de strings.
- `make([]string, 0)` crea un slice vacío listo para usar.
- Recorro el map. En cada vuelta pregunto si la cantidad del producto es menor al mínimo.
- Si es menor, agrego el nombre con `append` y guardo el resultado en `stock`.
- Al final imprimo el slice y lo devuelvo.

### La función main

- Creé los 4 productos con la forma `Campo: valor`.
- Los puse en el map con las claves `"A1"`, `"B2"`, `"C3"` y `"D4"`.
- Llamé a `valorTotal(productos)` y a `bajoStock(productos, 5)`.

## Cómo calculé el resultado a mano

| Producto | Precio | Cantidad | Valor |
|---|---|---|---|
| Monitor | 1000 | 4 | 4000 |
| Teclado | 300 | 5 | 1500 |
| Procesador | 5000 | 1 | 5000 |
| Disco de almacenamiento | 800 | 2 | 1600 |
| **Total** | | | **12100** |

Con mínimo 5, los de bajo stock son los que tienen menos de 5 unidades: Monitor (4), Procesador (1) y Disco de almacenamiento (2). El Teclado tiene exactamente 5, así que no entra porque la condición es `< 5`.

## Salida del programa

```
El valor total del inventario es:  12100
Productos con bajo stock: [Monitor Procesador Disco de almacenamiento]
```

El orden de los nombres en el slice puede cambiar entre ejecuciones, porque un map en Go no tiene orden fijo.

## Lo que aprendí

- Un struct agrupa datos relacionados.
- Un map con structs como valores sirve para guardar un inventario con claves.
- `range` recorre un map y `_` ignora lo que no necesito.
- Hay que convertir tipos con `float64(...)` antes de operar con `int` y `float64`.
- `append` devuelve un slice nuevo y hay que guardarlo.
- Una función puede imprimir y devolver un valor a la vez.

## Cosas que puedo mejorar

Ya corregí los errores de escritura del nombre `bajoStock` y del mensaje "Productos con bajo stock". Lo que queda por mejorar:


- Cambiar el nombre del struct a `Producto` (singular), porque representa un solo producto.
- En `bajoStock`, la variable `cantidad` guarda un producto completo, así que un nombre mejor sería `producto`.
- Dejar que `main` imprima los resultados y que las funciones solo devuelvan valores. Así se pueden reutilizar.
- Hacer el bonus: borrar un producto con `delete(productos, "A1")` y volver a calcular el total.
