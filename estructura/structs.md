# Structs (estructuras) en Go

Una struct agrupa datos de distintos tipos bajo una sola entidad. Funciona como una plantilla para representar algo con varios atributos, por ejemplo un registro de base de datos.

## Definir una struct
```go
type Pais struct {
    Nombre      string
    Capital     string
    Idioma      string
    Poblacion   int
}
```

### Elementos de la sintaxis
- `type`: palabra clave para definir un nuevo tipo.
- `Pais`: el nombre de la struct.
- `struct`: indica el inicio de la definición.
- `{}`: delimitan el bloque.
- Cada campo con su tipo, por ejemplo `Nombre string`.

## Ejemplo completo
```go
package main

import "fmt"

type Pais struct {
    Nombre      string
    Capital     string
    Idioma      string
    Poblacion   int
}

func main() {
    irlanda := Pais{
        Nombre:     "Irlanda",
        Capital:    "Dublin",
        Idioma:     "Irlandes",
        Poblacion:  4857000,
    }
    fmt.Printf("%+v\n", irlanda)
}
```

Se agrupan varios tipos de datos en una sola variable. `%+v` imprime los nombres de los campos junto con sus valores.

## Inicialización y valor cero
Si un campo no se define al crear la struct, toma el valor cero de su tipo: 0 para enteros, "" para cadenas.

```go
Pais{
    Nombre:     "Irlanda",
    Capital:    "Dublin",
    Idioma:     "Irlandes",
    Poblacion:  4857000,
}
```

## Structs vs maps
- Struct: campos fijos, definidos en tiempo de compilación. Ideal para entidades con atributos conocidos.
- Map: claves dinámicas. Ideal para colecciones de pares clave-valor que cambian.

## Buenas prácticas
- Nombrar los campos de forma clara.
- Usar el tipo de dato apropiado para cada campo.

## Agregado mío (no venía en el texto original)
Leer un campo desde el teclado, de tu ejercicio del país:
```go
fmt.Print("Digite su Idioma aqui: ")
fmt.Scanln(&RepublicaDominicana.Idioma)
```

Un método sobre una struct:
```go
func (p Pais) detallesDeMiPais() {
    fmt.Println("el nombre de mi pais es:", p.Nombre)
}
```
