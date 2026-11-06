# Structs en Go: instanciación, acceso y structs anidados

Fecha: 5 de noviembre de 2026

## 1. Qué es un struct

Un struct agrupa datos de distintos tipos bajo un solo nombre. Sirve para organizar información compleja, como los datos de un país.

```go
type Pais struct {
    Nombre    string
    Capital   string
    Idioma    string
    Poblacion int
}
```

## 2. Formas de crear (instanciar) un struct

### Declarar la variable sin valores

Todos los campos se inicializan con su **valor cero**: `""` para string y `0` para int.

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
    var belice Pais
    fmt.Printf("%+v\n", belice)
}
```

Salida:

```
{Nombre: Capital: Idioma: Poblacion:0}
```

- `%+v` imprime el struct mostrando el nombre de cada campo.

### Literal con valores

Se asigna un valor a cada campo por su nombre.

```go
panama := Pais{
    Nombre:    "Panama",
    Capital:   "Ciudad de Panama",
    Idioma:    "Español",
    Poblacion: 4170607,
}
fmt.Println("Nombre: ", panama.Nombre)
fmt.Println("Capital: ", panama.Capital)
fmt.Println("Idioma: ", panama.Idioma)
fmt.Println("Población: ", panama.Poblacion)
```

## 3. Acceder y modificar campos

- Se accede con el punto: `panama.Capital`.
- Se modifica asignando un valor nuevo: `panama.Capital = "Panamá"`.
- Los valores de los campos se pueden cambiar, pero **el tipo de cada campo no cambia**.

### Mayúscula y minúscula

- Campo con **mayúscula** (`Nombre`): exportable, se puede usar desde otros paquetes.
- Campo con **minúscula** (`nombre`): privado, solo se usa dentro de su paquete.

## 4. Structs anidados

Un struct puede tener dentro otro struct. Así se organizan mejor los datos relacionados: un país tiene una capital, y la capital tiene sus propios datos.

```go
type Ciudad struct {
    Nombre    string
    Poblacion int
}

type Pais struct {
    Nombre    string
    Capital   Ciudad
    Idioma    string
    Poblacion int
}
```

### Inicializar en dos pasos

Primero se crea la ciudad y luego se usa dentro del país.

```go
ciudadDeMexico := Ciudad{
    Nombre:    "Ciudad de Mexico",
    Poblacion: 20e6,
}
mexico := Pais{
    Nombre:    "México",
    Capital:   ciudadDeMexico,
    Idioma:    "Español",
    Poblacion: 120e6,
}
```

### Inicializar todo junto

La ciudad se escribe directamente dentro del país. Es más corto.

```go
mexico := Pais{
    Nombre: "México",
    Capital: Ciudad{
        Nombre:    "Ciudad de Mexico",
        Poblacion: 20e6,
    },
    Idioma:    "Español",
    Poblacion: 120e6,
}
```

### Acceder a un campo anidado

Se encadenan los puntos:

```go
fmt.Println(mexico.Capital.Nombre)
fmt.Println(mexico.Capital.Poblacion)
```

### Ejemplo completo

```go
package main

import "fmt"

type Ciudad struct {
    Nombre    string
    Poblacion int
}

type Pais struct {
    Nombre    string
    Capital   Ciudad
    Idioma    string
    Poblacion int
}

func main() {
    ciudadDeMexico := Ciudad{
        Nombre:    "Ciudad de Mexico",
        Poblacion: 20e6,
    }
    mexico := Pais{
        Nombre:    "México",
        Capital:   ciudadDeMexico,
        Idioma:    "Español",
        Poblacion: 120e6,
    }
    fmt.Printf("%+v\n", mexico)
}
```

Salida:

```
{Nombre:México Capital:{Nombre:Ciudad de Mexico Poblacion:20000000} Idioma:Español Poblacion:120000000}
```

- `20e6` es notación científica y equivale a 20 000 000.

## 5. Struct simple vs struct anidado

| Struct simple | Struct anidado |
|---|---|
| Solo tiene campos básicos (string, int...) | Tiene campos que son otros structs |
| Sirve para datos sencillos | Sirve para modelar entidades más complejas |

## 6. Buenas prácticas

- Usar nombres descriptivos para cada struct.
- No anidar demasiados niveles, para que el código siga siendo fácil de leer.
- Reutilizar un struct cuando represente la misma cosa en varios lugares.
- Documentar los campos importantes.

## Resumen

- `var x Struct` crea el struct con valores cero.
- `Struct{Campo: valor}` lo crea con valores propios.
- Punto para leer y modificar: `x.Campo`.
- Mayúscula = público, minúscula = privado.
- Un struct puede contener otro struct, y se accede con `x.Campo.SubCampo`.
