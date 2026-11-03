# El valor cero en Go

El **valor cero** es el valor que Go asigna automáticamente a una variable cuando no se le proporciona un valor inicial. A diferencia de otros lenguajes donde podríamos encontrar `nil` o `undefined`, en Go cada tipo de dato tiene un valor predeterminado, lo que elimina la posibilidad de variables sin inicializar.

---

## Ejemplo práctico

```go
package main

import "fmt"

func main() {
    var nombre string
    var edad int
    var peso float64
    var estudiante bool
    fmt.Println("Nombre:", nombre)
    fmt.Println("Edad:", edad)
    fmt.Println("Peso:", peso)
    fmt.Println("Estudiante:", estudiante)
}
```

Salida del programa:

```
Nombre:
Edad: 0
Peso: 0
Estudiante: false
```

### Valores cero por tipo

| Tipo               | Valor cero        |
|--------------------|-------------------|
| `string`           | `""` (cadena vacía) |
| `int`              | `0`               |
| `float64`          | `0`               |
| `bool`             | `false`           |

---

## Importancia de conocer los valores cero

- Al declarar una variable sin asignarle un valor, toma automáticamente su valor cero.
- Es fundamental al hacer validaciones: una cadena no se compara con `nil`, sino con `""` (cadena vacía).
- Las diferencias entre valor cero y `nil` pueden causar errores sutiles si no se tienen en cuenta.

---

## Conclusión

El valor cero garantiza que todas las variables tengan un estado definido desde su declaración, lo que aporta seguridad y previsibilidad al código. Entender este comportamiento es esencial para evitar errores comunes y escribir programas más robustos.
