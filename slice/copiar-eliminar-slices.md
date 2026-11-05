# Copiar y eliminar elementos de un slice en Go

## Copiar con copy
Asignar un slice a otra variable no crea una copia, solo otra referencia al mismo array. Para tener una copia independiente se usa `copy`.

```go
figuras := []string{"circulo", "cuadrado", "triangulo", "rombo", "trapecio", "heptagono"}
figuras2 := make([]string, len(figuras))
copy(figuras2, figuras)
```

## Eliminar un elemento
Se elimina uniendo lo que está antes y después del índice:

```go
figuras = append(figuras[:1], figuras[2:]...)
```

Esto quita "cuadrado" (índice 1).

## Ejemplo completo
```go
package main

import "fmt"

func main() {
    figuras := []string{"circulo", "cuadrado", "triangulo", "rombo", "trapecio", "heptagono"}
    figuras2 := make([]string, len(figuras))
    copy(figuras2, figuras)

    figuras = append(figuras[:1], figuras[2:]...)

    fmt.Println(figuras2)
    fmt.Println(figuras)
}
```

Salida:
```
[circulo cuadrado triangulo rombo trapecio heptagono]
[circulo triangulo rombo trapecio heptagono]
```

## Notas
- `copy(destino, origen)` copia solo hasta el tamaño del destino, por eso se crea con `make([]T, len(origen))`.
- Si no usas `copy` y eliminas con `append`, el slice original comparte el array y se modifica por debajo.
- `copy` devuelve la cantidad de elementos copiados.
