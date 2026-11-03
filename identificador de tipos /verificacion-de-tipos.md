# Verificación de tipos de datos en Go

En Go es común necesitar saber el tipo de una variable durante la ejecución, sobre todo al trabajar con interfaces vacías (`interface{}` / `any`) o datos dinámicos. Go ofrece varias formas de hacerlo.

## 1. Paquete `reflect`

`reflect.TypeOf` devuelve el tipo de una variable en tiempo de ejecución.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var alumno string = "Jose Luis"
    var edad int = 22
    var peso float64 = 85.5
    fmt.Println(reflect.TypeOf(alumno))
    fmt.Println(reflect.TypeOf(edad))
    fmt.Println(reflect.TypeOf(peso))
}
```

Salida:

```
string
int
float64
```

## 2. Formato `%T`

Ideal para imprimir o depurar.

```go
fmt.Printf("%T\n", edad)
```

## 3. Type switch (lo más idiomático)

Sirve para actuar según el tipo de un valor.

```go
func describir(v interface{}) {
    switch t := v.(type) {
    case string:
        fmt.Println("texto:", t)
    case int:
        fmt.Println("entero:", t)
    case float64:
        fmt.Println("decimal:", t)
    default:
        fmt.Printf("otro tipo: %T\n", t)
    }
}
```

## 4. Aserción de tipo

Para comprobar un tipo concreto.

```go
n, ok := v.(int)
```

## Cuándo usar cada una

| Herramienta | Uso |
|---|---|
| `%T` | Depuración e impresión rápida |
| Type switch | Ejecutar lógica distinta según el tipo |
| Aserción de tipo | Comprobar un solo tipo concreto |
| `reflect` | Inspección genérica de estructuras (campos, tags, tamaños), como hacen `encoding/json` o los ORMs |

## Notas

- `reflect` es más lento y menos seguro que un type switch; no usarlo para casos simples.
- Todo esto solo tiene sentido con `interface{}` o `any`: en una variable declarada como `int`, el tipo ya se conoce en compilación.
