# ¿Qué es la sentencia `switch` en Go?

## Definición

`switch` es una estructura de control que permite **elegir entre varias rutas de ejecución** de forma ordenada. Evalúa una expresión y ejecuta el bloque del **primer `case` que coincida**.

Frente a una cadena larga de `if-else`, el `switch` es más legible y fácil de mantener.

---

## Ejemplo práctico

```go
package main

import "fmt"

func main() {
    var juguete string
    fmt.Println("Elija qué tipo de juguete agregar: persona, animal o cosa")
    fmt.Scanln(&juguete)
    switch juguete {
    case "persona":
        fmt.Println("El juguete es una figura de acción")
    case "cosa":
        fmt.Println("El juguete es una cosa")
    case "animal":
        fmt.Println("El juguete es una mascota")
    default:
        fmt.Println("El juguete es otra categoría")
    }
}
```

### Cómo funciona

1. El programa pide un texto y lo guarda en `juguete`.
2. `switch juguete` compara ese valor con cada `case`, de arriba hacia abajo.
3. Al primer `case` que coincide, ejecuta su bloque y **sale del `switch`**.
4. Si ninguno coincide, se ejecuta `default`.

Ejemplo de salida si escribes `animal`:

```
El juguete es una mascota
```

---

## Diferencias entre `switch` e `if` en Go

- `switch` compara **un mismo valor** contra muchas opciones de forma clara y compacta.
- Con `if-else` hay que escribir una condición por cada caso, y se vuelve largo.
- `switch` agrupa todos los casos posibles en un solo lugar.

---

## Características propias de Go

### No necesita `break`

En otros lenguajes (C, C++, Java) hay que escribir `break` al final de cada `case` o se siguen ejecutando los siguientes. **En Go no**: cada `case` termina solo.

### Varios valores en un mismo `case`

```go
package main

import "fmt"

func main() {
    dia := "sábado"
    switch dia {
    case "sábado", "domingo":
        fmt.Println("Fin de semana")
    default:
        fmt.Println("Día laboral")
    }
}
```

Salida: `Fin de semana`.

### `switch` sin expresión (reemplaza un `if-else if`)

```go
package main

import "fmt"

func main() {
    nota := 85
    switch {
    case nota >= 90:
        fmt.Println("A")
    case nota >= 80:
        fmt.Println("B")
    case nota >= 70:
        fmt.Println("C")
    default:
        fmt.Println("Reprobado")
    }
}
```

Salida: `B`. Cada `case` lleva una condición; gana la primera que sea verdadera.

### `fallthrough`: continuar al siguiente `case`

Si quieres que Go siga ejecutando el siguiente `case` (sin revisar su condición), usas `fallthrough`.

```go
package main

import "fmt"

func main() {
    n := 1
    switch n {
    case 1:
        fmt.Println("uno")
        fallthrough
    case 2:
        fmt.Println("dos")
    case 3:
        fmt.Println("tres")
    }
}
```

Salida:

```
uno
dos
```

### `switch` con inicialización

```go
package main

import "fmt"

func main() {
    switch x := 5 * 2; x {
    case 10:
        fmt.Println("es diez")
    default:
        fmt.Println("otro valor")
    }
}
```

Salida: `es diez`. La variable `x` solo existe dentro del `switch`.

---

## Buenas prácticas

- Úsalo cuando varias condiciones dependen de **un mismo valor**.
- Incluye siempre un `default` para manejar lo que no esperabas y que el programa se comporte de forma predecible.
- Usa `switch` sin expresión en lugar de cadenas largas de `if-else if`.

---

## Manejo de múltiples casos

Cada `case` es una posible coincidencia. Si ninguna se cumple, actúa `default`. Así se controla el flujo del programa de forma eficiente y clara.

---

## Conclusión

`switch` simplifica la toma de decisiones múltiples en Go: mejora la legibilidad, facilita el mantenimiento y permite escribir programas más claros y estructurados.

---

## Ejercicios para practicar

1. Pide un número del 1 al 7 y muestra el nombre del día de la semana con `switch`.
2. Pide una nota (0-100) y muestra A, B, C o Reprobado con `switch` sin expresión.
3. Haz una mini calculadora: pide dos números y un operador (`+`, `-`, `*`, `/`) y usa `switch` para calcular.
4. Pide un mes y usa un `case` con varios valores para decir cuántos días tiene.
