# Funciones como parámetros en Go (y closures)

## 1. Idea central

En Go las funciones son **valores**, igual que un número o un texto. Eso significa que puedes:

1. **Guardar** una función en una variable.
2. **Pasarla** como argumento a otra función.
3. **Devolverla** desde otra función.

Esto permite escribir código más flexible y reutilizable: en vez de fijar *qué hacer*, la función recibe *qué hacer* como parámetro.

---

## 2. El tipo función

Una función tiene un **tipo**, que depende de sus parámetros y de lo que devuelve (su *firma*). Se escribe con `func`:

| Función | Su tipo |
|---|---|
| `func saludar(nombre string) string` | `func(string) string` |
| `func sumar(a, b int) int` | `func(int, int) int` |
| `func imprimir()` | `func()` |

Dos funciones tienen el mismo tipo si reciben y devuelven lo mismo. El nombre de los parámetros no importa.

---

## 3. Paso 1: guardar una función en una variable

```go
package main

import "fmt"

func saludar(nombre string) string {
    return "Hola, " + nombre
}

func main() {
    f := saludar
    fmt.Println(f("Wilber"))

    var g func(string) string
    g = saludar
    fmt.Println(g("Go"))
}
```

Salida:

```
Hola, Wilber
Hola, Go
```

Detalle que confunde mucho:
- `saludar` (sin paréntesis) es **la función como valor**.
- `saludar("x")` (con paréntesis) es **llamarla** y obtener su resultado.

---

## 4. Paso 2: pasar una función como argumento (callback)

```go
package main

import "fmt"

func sumar(a, b int) int       { return a + b }
func restar(a, b int) int      { return a - b }
func multiplicar(a, b int) int { return a * b }

func aplicar(a, b int, operacion func(int, int) int) int {
    return operacion(a, b)
}

func main() {
    fmt.Println(aplicar(6, 3, sumar))
    fmt.Println(aplicar(6, 3, restar))
    fmt.Println(aplicar(6, 3, multiplicar))
    fmt.Println(aplicar(6, 3, func(x, y int) int { return x*x + y }))
}
```

Salida:

```
9
3
18
39
```

### Cómo leerlo

- `aplicar` recibe tres cosas: `a`, `b` y `operacion`.
- `operacion` es un parámetro de tipo `func(int, int) int`: una función que recibe dos enteros y devuelve uno.
- Dentro de `aplicar`, `operacion(a, b)` **llama** a la función que le pasaron.
- Cada vez que cambias la función que pasas, `aplicar` hace algo distinto, sin tocar su código.
- En el último caso se pasó una **función anónima** (sin nombre) escrita en el mismo lugar.

### Darle un nombre al tipo

Cuando la firma es larga, se le da un nombre con `type`:

```go
package main

import "fmt"

type Operacion func(int, int) int

func aplicar(a, b int, op Operacion) int {
    return op(a, b)
}

func main() {
    var suma Operacion = func(a, b int) int { return a + b }
    fmt.Println(aplicar(2, 3, suma))
}
```

Salida: `5`. Es solo una forma de escribirlo más corto y legible.

---

## 5. Funciones anónimas

Una **función anónima** (o literal de función) es una función sin nombre, escrita donde se necesita.

Se puede guardar en una variable, pasarla como argumento o **ejecutarla al instante** poniendo `()` al final:

```go
package main

import "fmt"

func main() {
    resultado := func(a, b int) int {
        return a + b
    }(3, 4)
    fmt.Println(resultado)
}
```

Salida: `7`. Se definió y se llamó de inmediato con los argumentos `(3, 4)`.

---

## 6. Paso 3: devolver una función

Una función puede devolver otra función:

```go
package main

import "fmt"

func crearMultiplicador(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}

func main() {
    doble := crearMultiplicador(2)
    triple := crearMultiplicador(3)
    fmt.Println(doble(5))
    fmt.Println(triple(5))
    fmt.Println(crearMultiplicador(10)(7))
}
```

Salida:

```
10
15
70
```

- `crearMultiplicador(2)` devuelve una función que multiplica por 2. Esa función **recuerda** que `factor` es 2.
- `crearMultiplicador(10)(7)` crea la función y la llama de inmediato con 7.

Aquí aparece el concepto de la siguiente sección.

---

## 7. Closures (el punto más delicado)

Un **closure** es una función que **captura variables del entorno donde fue creada** y puede seguir usándolas aunque la función que las creó ya haya terminado.

En `crearMultiplicador`, la función interna usa `factor`, que no es suyo: pertenece a `crearMultiplicador`. Aunque `crearMultiplicador` ya terminó, `factor` sigue vivo porque el closure lo tiene "agarrado".

### El closure captura la variable, no una copia del valor

Esto significa que si la variable cambia, el closure ve el cambio. Ejemplo con un contador:

```go
package main

import "fmt"

func crearContador() func() int {
    cuenta := 0
    return func() int {
        cuenta++
        return cuenta
    }
}

func main() {
    a := crearContador()
    b := crearContador()
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(b())
    fmt.Println(a())
}
```

Salida:

```
1
2
3
1
4
```

Cómo entenderlo:
- Cada llamada a `crearContador()` crea una variable `cuenta` **nueva**, propia de ese closure.
- `a` y `b` son contadores **independientes**: `b()` empieza en 1 aunque `a` ya iba en 3.
- La variable `cuenta` **conserva su valor entre llamadas**, porque el closure la mantiene viva.

---

## 8. El ejemplo del círculo (devolver más de una función)

```go
package main

import "fmt"

func circulo(radio float64) (area func() float64, perimetro func() float64) {
    area = func() float64 {
        return 3.1416 * radio * radio
    }

    perimetro = func() float64 {
        return 2 * 3.1416 * radio
    }

    return
}

func main() {
    area, perimetro := circulo(10)
    fmt.Println("El area del circulo es", area())
    fmt.Println("El perimetro del circulo es", perimetro())
}
```

### Qué pasa

- `circulo(10)` devuelve **dos funciones**: `area` y `perimetro`.
- Las dos son closures: ambas capturan `radio` (que vale 10).
- Aunque `circulo` ya terminó, `area()` y `perimetro()` siguen recordando el radio.
- `return` sin valores funciona porque los resultados tienen nombre (*named returns*): `area` y `perimetro` ya están asignados.

### Salida correcta

```
El area del circulo es 314.16
El perimetro del circulo es 62.832
```

Los valores salen aproximados: al ser números decimales (`float64`), puede aparecer algún decimal de más.

> El texto original de donde viene este ejemplo muestra un área de 201.0624, y eso es un error: con radio 10, `3.1416 × 10 × 10 ≈ 314.16`. El valor 201.0624 corresponde a un radio de 8. El perímetro de 62.832 sí es correcto.

### Matiz importante

Este ejemplo **devuelve** funciones, no las **pasa como argumento**. Son dos usos distintos de las funciones como valores:

| Uso | Ejemplo |
|---|---|
| Pasar como argumento | `aplicar(6, 3, sumar)` |
| Devolver como resultado | `crearMultiplicador(2)` |
| Ambos aparecen en closures | los dos ejemplos de arriba |

---

## 9. Patrones comunes con funciones como argumentos

### `map`, `filter` y `reduce` hechos a mano

```go
package main

import "fmt"

func mapear(nums []int, f func(int) int) []int {
    resultado := make([]int, 0, len(nums))
    for _, n := range nums {
        resultado = append(resultado, f(n))
    }
    return resultado
}

func filtrar(nums []int, cond func(int) bool) []int {
    resultado := []int{}
    for _, n := range nums {
        if cond(n) {
            resultado = append(resultado, n)
        }
    }
    return resultado
}

func reducir(nums []int, inicial int, f func(int, int) int) int {
    acumulado := inicial
    for _, n := range nums {
        acumulado = f(acumulado, n)
    }
    return acumulado
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6}
    cuadrados := mapear(nums, func(n int) int { return n * n })
    pares := filtrar(nums, func(n int) bool { return n%2 == 0 })
    suma := reducir(nums, 0, func(acc, n int) int { return acc + n })
    fmt.Println(cuadrados)
    fmt.Println(pares)
    fmt.Println(suma)
}
```

Salida:

```
[1 4 9 16 25 36]
[2 4 6]
21
```

- `mapear`: aplica una función a cada elemento y devuelve la lista transformada.
- `filtrar`: se queda solo con los elementos que cumplen una condición.
- `reducir`: combina todos los elementos en un solo valor.

La lógica del bucle se escribe **una sola vez**, y lo que cambia (qué hacer con cada elemento) se pasa como función.

### Ordenar con una función de comparación

La biblioteca estándar usa este patrón. `sort.Slice` recibe una función que dice cuándo un elemento va antes que otro:

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    nombres := []string{"Carlos", "Ana", "Beatriz", "Li"}
    sort.Slice(nombres, func(i, j int) bool {
        return len(nombres[i]) < len(nombres[j])
    })
    fmt.Println(nombres)
}
```

Salida:

```
[Li Ana Carlos Beatriz]
```

### Ejecutar código "antes y después" de otra tarea

```go
package main

import "fmt"

func medirTiempo(nombre string, f func()) {
    fmt.Println("inicio:", nombre)
    f()
    fmt.Println("fin:", nombre)
}

func main() {
    medirTiempo("tarea 1", func() {
        fmt.Println("trabajando...")
    })
}
```

Salida:

```
inicio: tarea 1
trabajando...
fin: tarea 1
```

Este es el patrón básico detrás de los *middlewares* y *decoradores*: una función que envuelve a otra.

---

## 10. Conexión con `defer`

`defer` también recibe una llamada a función, y muchas veces se usa con una función anónima. Aquí se nota la diferencia entre un argumento evaluado al instante y un closure que captura la variable:

```go
package main

import "fmt"

func main() {
    x := 1
    defer fmt.Println("directo:", x)
    defer func() {
        fmt.Println("closure:", x)
    }()
    x = 2
}
```

Salida:

```
closure: 2
directo: 1
```

- `defer fmt.Println("directo:", x)`: el valor de `x` se **evalúa al escribir el `defer`**, o sea 1.
- El closure lee `x` cuando **se ejecuta**, o sea al final, cuando ya vale 2.
- Salen en orden inverso: el último `defer` escrito corre primero.

---

## 11. Closures dentro de bucles

```go
package main

import "fmt"

func main() {
    var funciones []func()
    for i := 0; i < 3; i++ {
        funciones = append(funciones, func() {
            fmt.Println(i)
        })
    }
    for _, f := range funciones {
        f()
    }
}
```

- Desde **Go 1.22**, cada vuelta del `for` tiene su propia variable `i`, y la salida es `0 1 2`.
- En versiones **anteriores a 1.22**, todos los closures compartían la misma `i`, y la salida era `3 3 3`.

Si usas una versión vieja de Go, la solución era copiar la variable dentro del bucle (`i := i`). Revisa tu versión con `go version`.

---

## 12. Errores comunes

1. **Confundir `f` con `f()`**: `f` es la función como valor; `f()` la ejecuta.
2. **Firma que no coincide**: pasar una función de tipo `func(int) int` donde se espera `func(int, int) int` da error de compilación.
3. **Llamar una función `nil`**: una variable de tipo función sin asignar vale `nil`, y llamarla provoca un `panic` (`invalid memory address or nil pointer dereference`).

```go
package main

import "fmt"

func main() {
    var f func(int) int
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("error:", r)
        }
    }()
    fmt.Println(f(1))
}
```

Salida:

```
error: runtime error: invalid memory address or nil pointer dereference
```

4. **No esperar que el closure cambie**: como captura la variable (no una copia), si la variable cambia después, el closure ve el nuevo valor.
5. **Closures en bucles** en versiones de Go anteriores a 1.22 (sección 11).
6. **Abusar de funciones anónimas largas**: dificultan la lectura; si pasa de unas cuantas líneas, mejor ponle nombre.

---

## 13. Ventajas

- **Flexibilidad**: el comportamiento se adapta según la función que pasas.
- **Modularidad**: separas responsabilidades en piezas pequeñas.
- **Patrones funcionales**: callbacks, filtros, mapeos, reducciones.
- **Menos código duplicado**: el bucle o la estructura se escribe una vez.
- **Mejor testabilidad**: puedes pasar versiones falsas (*mocks*) en las pruebas.

## 14. Desventajas

- **Más complejidad**: el abuso hace el código difícil de leer, sobre todo para quien no conoce el estilo funcional.
- **Depuración más difícil**: seguir qué función se ejecuta y cuándo cuesta más.
- **Carga mental**: hay que entender la función principal y las que se le pasan.
- **Rendimiento**: una llamada a través de una variable de función tiene un pequeño costo extra. Casi nunca importa, salvo en casos extremos.
- **Closures y memoria**: las variables capturadas siguen vivas mientras el closure exista.

---

## 15. Resumen rápido

| Concepto | Qué es |
|---|---|
| Función como valor | Se puede guardar en variable, pasar y devolver |
| Tipo función | `func(parámetros) retorno` |
| Callback | Función pasada como argumento para que otra la llame |
| Función anónima | Función sin nombre, escrita donde se usa |
| Closure | Función que captura variables de donde fue creada |

---

## 16. Ejercicios para practicar

1. Escribe `aplicar` con una operación `dividir` y maneja la división entre cero.
2. Crea `crearSumador(n int) func(int) int` y úsalo para hacer `sumaCinco` y `sumaDiez`.
3. Crea un contador con `crearContador()` que además permita reiniciarse. Pista: devuelve **dos** funciones, `siguiente` y `reiniciar`.
4. Con `filtrar`, saca de una lista solo los números mayores que 10.
5. Con `mapear`, convierte una lista de enteros en una lista con cada número multiplicado por 3.
6. Con `reducir`, calcula el producto de todos los elementos de una lista.
7. Con `sort.Slice`, ordena una lista de números de mayor a menor.
8. Escribe `medirTiempo` usando el paquete `time` para mostrar cuánto tarda la función que recibe.
9. Predice la salida de la sección 10 si cambias `x = 2` por `x = 5`.
