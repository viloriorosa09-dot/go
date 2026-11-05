# Maps en Go

## 1. Creación

Un map guarda pares clave-valor. También se le llama diccionario, hash o array asociativo. A diferencia de arrays y slices, que se acceden por índice numérico, un map se accede con una clave personalizada.

### Ejemplo: días de la semana con su número
```go
package main

import "fmt"

func main() {
    diasDeLaSemana := make(map[string]int)
    diasDeLaSemana["lunes"] = 1
    diasDeLaSemana["martes"] = 2
    diasDeLaSemana["miércoles"] = 3
    diasDeLaSemana["jueves"] = 4
    diasDeLaSemana["viernes"] = 5
    diasDeLaSemana["sábado"] = 6
    diasDeLaSemana["domingo"] = 7
    fmt.Println(diasDeLaSemana)
    fmt.Println(diasDeLaSemana["lunes"])
}
```

### Explicación
- `make(map[string]int)` crea un map con claves `string` y valores `int`.
- No hace falta definir un tamaño.
- Se agrega con `mapa[clave] = valor`.
- Se lee con `mapa[clave]`.
- Diferencia con el slice: el map se accede por clave, el slice por índice numérico.

## 2. Eliminación

Se usa la función `delete(mapa, clave)`. Quita la clave y su valor. Si la clave no existe, no da error.

### Ejemplo: días en inglés
```go
package main

import "fmt"

func main() {
    diasDeLaSemanaEnIngles := map[string]string{
        "lunes":     "Monday",
        "martes":    "Tuesday",
        "miercoles": "Wednesday",
        "jueves":    "Thursday",
        "sabado":    "Saturday",
        "domingo":   "Sunday",
    }
    fmt.Println(diasDeLaSemanaEnIngles)
    delete(diasDeLaSemanaEnIngles, "domingo")
    fmt.Println(diasDeLaSemanaEnIngles)
}
```

### Explicación
- Se crea el map con valores iniciales.
- `delete(diasDeLaSemanaEnIngles, "domingo")` elimina la clave "domingo" y su valor "Sunday".
- Se imprime antes y después para ver el cambio.

### Buenas prácticas
- Verifica si la clave existe antes de operar con su valor.
- Usa `delete` para no dejar datos obsoletos en el map.
- Eliminar una clave que no existe no genera error.

## Agregado mío (no venía en el texto original)
Para saber si una clave existe:
```go
valor, ok := diasDeLaSemana["viernes"]
if ok {
    fmt.Println(valor)
} else {
    fmt.Println("no encontrado")
}
```

- Leer una clave que no existe devuelve el valor cero del tipo (0 en int, "" en string).
- `fmt.Println` imprime los maps ordenados por clave, no en el orden en que los escribiste.
