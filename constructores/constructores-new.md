# Crear structs con new en Go

Fecha: 6 de noviembre de 2026

## 1. Qué hace new

`new(Tipo)` reserva memoria para una variable de ese tipo y devuelve un **puntero** a ella. Todos los campos empiezan con su valor cero.

```go
colombia := new(Pais)
```

`colombia` no es un `Pais`, es un `*Pais` (puntero a Pais).

## 2. Ejemplo

```go
package main

import "fmt"

type Pais struct {
    Nombre    string
    Capital   string
    Idioma    string
    Poblacion int
}

func main() {
    colombia := new(Pais)
    fmt.Printf("%+v\n", colombia)
    colombia.Nombre = "Colombia"
    colombia.Capital = "Bogota"
    colombia.Idioma = "Español"
    colombia.Poblacion = 49e6
    fmt.Printf("%+v\n", colombia)
}
```

Primera salida (valores cero):

```
&{Nombre: Capital: Idioma: Poblacion:0}
```

Segunda salida (ya con valores):

```
&{Nombre:Colombia Capital:Bogota Idioma:Español Poblacion:49000000}
```

- El `&` al inicio de la salida indica que lo que se imprime es un puntero a un struct.

## 3. Asignar valores a los campos

Con el puntero se usa el mismo punto de siempre:

```go
colombia.Nombre = "Colombia"
```

Go entiende que `colombia` es un puntero y accede al struct por ti. No hace falta escribir `(*colombia).Nombre`.

## 4. new vs make

| new | make |
|---|---|
| Sirve para structs y otros tipos | Solo para slices, maps y canales |
| Devuelve un puntero (`*T`) | Devuelve el valor ya inicializado |
| Deja todo en valor cero | Deja el slice, map o canal listo para usar |

Para un struct se usa `new` o la inicialización literal.

## 5. Dos formas de obtener un puntero a un struct

```go
a := new(Pais)
b := &Pais{Nombre: "Colombia", Capital: "Bogota"}
```

- `new` deja los campos en cero y los llenas después.
- `&Pais{...}` crea el struct con valores desde el inicio y devuelve su puntero.

## 6. Ventajas

- La inicialización es clara.
- Los campos nunca quedan con basura, siempre empiezan en valor cero.
- Se puede pasar el puntero a una función para que modifique el struct original.

## Resumen

- `new(Pais)` devuelve `*Pais` con campos en cero.
- Se asignan los campos con `puntero.Campo = valor`.
- `make` es para slices, maps y canales; `new` y el literal son para structs.
