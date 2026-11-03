# ¿Qué es una función recursiva en Go?

## 1. Definición

Una **función recursiva** es una función que **se llama a sí misma** durante su ejecución.

En lugar de repetir instrucciones con un bucle, la recursividad resuelve un problema **dividiéndolo en versiones más pequeñas del mismo problema**, hasta llegar a una versión tan simple que se resuelve directamente.

Es un tema que confunde al principio, así que este documento va despacio, con trazas paso a paso.

---

## 2. Las dos partes obligatorias

Toda función recursiva correcta tiene:

1. **Caso base** (condición de parada): el caso tan simple que se responde sin llamarse otra vez. Es lo que detiene la recursión.
2. **Caso recursivo**: la función se llama a sí misma con un problema **más pequeño** (parámetros que se acercan al caso base).

Si falta el caso base, o los parámetros nunca se acercan a él, la función se llama infinitamente y el programa se cae por **desbordamiento de pila** (*stack overflow*).

Receta mental para escribir una:

1. ¿Cuál es el caso más simple y su respuesta? (caso base)
2. ¿Cómo reduzco el problema un paso? (caso recursivo)
3. ¿Cómo combino el resultado de la llamada más pequeña con lo que me toca hacer a mí?

---

## 3. La pila de llamadas (el concepto clave)

Cuando una función llama a otra, Go guarda el estado de la primera (sus variables y por dónde iba) en una zona de memoria llamada **pila** (*stack*) y se pone a ejecutar la segunda. Cuando la segunda termina, Go **regresa** a la primera exactamente donde se quedó.

En la recursividad, cada llamada es una "copia nueva" de la función con **sus propios parámetros y variables**. Las llamadas se van apilando hacia abajo (ida) y luego se resuelven de abajo hacia arriba (vuelta).

Por eso cada llamada recursiva consume memoria: mientras más profunda la recursión, más pila se usa.

---

## 4. Ejemplo 1: el eco de la montaña

Repite un mensaje N veces numerado, usando recursividad.

```go
package main

import "fmt"

func ecoDeLaMontana(mensaje string, iteraciones uint) {
    if iteraciones > 1 {
        ecoDeLaMontana(mensaje, iteraciones-1)
    }
    fmt.Println(mensaje, iteraciones)
}

func main() {
    ecoDeLaMontana("yodelayheehoo", 5)
}
```

Salida:

```
yodelayheehoo 1
yodelayheehoo 2
yodelayheehoo 3
yodelayheehoo 4
yodelayheehoo 5
```

### ¿Por qué imprime de 1 a 5 si empezó en 5?

Esta es la parte que más confunde. Fíjate **dónde** está el `Println`: está **después** de la llamada recursiva. Entonces la función primero se mete hasta el fondo y solo imprime cuando va **de regreso**.

Traza paso a paso:

```
ecoDeLaMontana(5)   iteraciones > 1 es true  -> llama a (4), el Println de 5 queda en espera
  ecoDeLaMontana(4) iteraciones > 1 es true  -> llama a (3), el Println de 4 queda en espera
    ecoDeLaMontana(3) true -> llama a (2), el Println de 3 queda en espera
      ecoDeLaMontana(2) true -> llama a (1), el Println de 2 queda en espera
        ecoDeLaMontana(1) false -> NO se llama más (caso base)
                          imprime "yodelayheehoo 1" y termina
      vuelve a (2): imprime "yodelayheehoo 2"
    vuelve a (3): imprime "yodelayheehoo 3"
  vuelve a (4): imprime "yodelayheehoo 4"
vuelve a (5): imprime "yodelayheehoo 5"
```

Tabla de simulación:

| Profundidad | `iteraciones` | `iteraciones > 1` | ¿Qué hace? |
|---|---|---|---|
| 1 | 5 | true | se llama con 4 |
| 2 | 4 | true | se llama con 3 |
| 3 | 3 | true | se llama con 2 |
| 4 | 2 | true | se llama con 1 |
| 5 | 1 | false | caso base: imprime 1 y regresa |

Mueve el `Println` **antes** del `if` y la salida cambia a 5, 4, 3, 2, 1. Pruébalo, es el mejor ejercicio para entender el orden.

### Elementos del ejemplo

- La función recursiva `ecoDeLaMontana`.
- El parámetro `iteraciones`, que va bajando y asegura que se llegue al caso base.
- El `mensaje` que se muestra.

---

## 5. Ejemplo 2: factorial

El factorial de `n` (`n!`) es `n × (n-1) × ... × 1`, y por definición `0! = 1`.

- Caso base: `factorial(0) = 1`.
- Caso recursivo: `factorial(n) = n * factorial(n-1)`.

```go
package main

import "fmt"

func factorial(n uint) uint {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5))
}
```

Salida:

```
120
```

### Traza de `factorial(4)`

Fase de **ida** (se apilan las llamadas, nada se calcula todavía):

```
factorial(4) = 4 * factorial(3)
factorial(3) = 3 * factorial(2)
factorial(2) = 2 * factorial(1)
factorial(1) = 1 * factorial(0)
factorial(0) = 1     <- caso base
```

Fase de **vuelta** (se resuelve de abajo hacia arriba):

```
factorial(1) = 1 * 1 = 1
factorial(2) = 2 * 1 = 2
factorial(3) = 3 * 2 = 6
factorial(4) = 4 * 6 = 24
```

### Ver la recursividad en acción

Este programa imprime cuándo entra y cuándo sale cada llamada, con sangría según la profundidad:

```go
package main

import (
    "fmt"
    "strings"
)

func factorialVisual(n uint, nivel int) uint {
    sangria := strings.Repeat("  ", nivel)
    fmt.Printf("%sentra factorial(%d)\n", sangria, n)
    if n == 0 {
        fmt.Printf("%scaso base: devuelve 1\n", sangria)
        return 1
    }
    r := n * factorialVisual(n-1, nivel+1)
    fmt.Printf("%ssale factorial(%d) = %d\n", sangria, n, r)
    return r
}

func main() {
    factorialVisual(3, 0)
}
```

Salida:

```
entra factorial(3)
  entra factorial(2)
    entra factorial(1)
      entra factorial(0)
      caso base: devuelve 1
    sale factorial(1) = 1
  sale factorial(2) = 2
sale factorial(3) = 6
```

### Cuidado: desbordamiento de números

`uint` en una máquina de 64 bits guarda hasta unos 1,8 × 10^19. `factorial(20)` cabe, pero `factorial(21)` **se desborda** y Go no avisa: el resultado da la vuelta y sale un número incorrecto. Para factoriales grandes se usa el paquete `math/big`.

---

## 6. Ejemplo 3: serie de Fibonacci

Cada número es la suma de los dos anteriores: `0, 1, 1, 2, 3, 5, 8, 13...`

- Casos base: `fib(0) = 0` y `fib(1) = 1` (aquí son **dos** casos base).
- Caso recursivo: `fib(n) = fib(n-1) + fib(n-2)` (aquí hay **dos** llamadas recursivas).

```go
package main

import "fmt"

func fibonacci(n uint) uint {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println(fibonacci(6))
}
```

Salida:

```
8
```

### Por qué esta versión es lenta

Cada llamada genera dos más, y se repiten cálculos una y otra vez:

```
                    fib(4)
                  /        \
             fib(3)        fib(2)
            /     \        /    \
        fib(2)  fib(1)  fib(1)  fib(0)
        /    \
    fib(1)  fib(0)
```

`fib(2)` se calcula 2 veces, `fib(1)` 3 veces, etc. Para `fibonacci(6)` son 25 llamadas; para `fibonacci(40)` son más de 300 millones. El tiempo crece de forma exponencial.

### Solución: memoización

Guardar en un mapa los resultados ya calculados para no repetirlos:

```go
package main

import "fmt"

func fibMemo(n uint, memo map[uint]uint) uint {
    if n <= 1 {
        return n
    }
    if v, ok := memo[n]; ok {
        return v
    }
    memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
    return memo[n]
}

func main() {
    memo := make(map[uint]uint)
    fmt.Println(fibMemo(50, memo))
}
```

Salida:

```
12586269025
```

Ahora cada valor se calcula una sola vez, y `fibMemo(50)` es instantáneo.

### Versión con bucle (la más eficiente)

```go
package main

import "fmt"

func fibIterativo(n uint) uint {
    a, b := uint(0), uint(1)
    for i := uint(0); i < n; i++ {
        a, b = b, a+b
    }
    return a
}

func main() {
    fmt.Println(fibIterativo(50))
}
```

---

## 7. Ejemplo 4: más funciones recursivas simples

Suma de 1 a n:

```go
package main

import "fmt"

func sumaHasta(n int) int {
    if n == 0 {
        return 0
    }
    return n + sumaHasta(n-1)
}

func main() {
    fmt.Println(sumaHasta(5))
}
```

Salida: `15` (5 + 4 + 3 + 2 + 1 + 0).

Potencia (`base` elevada a `exp`):

```go
package main

import "fmt"

func potencia(base, exp int) int {
    if exp == 0 {
        return 1
    }
    return base * potencia(base, exp-1)
}

func main() {
    fmt.Println(potencia(2, 10))
}
```

Salida: `1024`.

Cuenta regresiva:

```go
package main

import "fmt"

func cuentaRegresiva(n int) {
    if n < 0 {
        return
    }
    fmt.Println(n)
    cuentaRegresiva(n - 1)
}

func main() {
    cuentaRegresiva(3)
}
```

Salida: `3 2 1 0` (uno por línea).

---

## 8. Ejemplo 5: recorrido de un árbol binario

Los árboles son el caso donde la recursividad brilla: un árbol es un nodo con dos subárboles, que a su vez son árboles.

```go
package main

import "fmt"

type Nodo struct {
    Valor int
    Izq   *Nodo
    Der   *Nodo
}

func preorden(n *Nodo) {
    if n == nil {
        return
    }
    fmt.Println(n.Valor)
    preorden(n.Izq)
    preorden(n.Der)
}

func inorden(n *Nodo) {
    if n == nil {
        return
    }
    inorden(n.Izq)
    fmt.Println(n.Valor)
    inorden(n.Der)
}

func postorden(n *Nodo) {
    if n == nil {
        return
    }
    postorden(n.Izq)
    postorden(n.Der)
    fmt.Println(n.Valor)
}

func main() {
    arbol := &Nodo{
        Valor: 1,
        Izq: &Nodo{
            Valor: 2,
            Izq:   &Nodo{Valor: 4},
            Der:   &Nodo{Valor: 5},
        },
        Der: &Nodo{Valor: 3},
    }

    fmt.Println("preorden:")
    preorden(arbol)
    fmt.Println("inorden:")
    inorden(arbol)
    fmt.Println("postorden:")
    postorden(arbol)
}
```

El árbol que se construye:

```
        1
       / \
      2   3
     / \
    4   5
```

Resultados:

| Recorrido | Orden de visita | Resultado |
|---|---|---|
| preorden | nodo, izquierda, derecha | 1 2 4 5 3 |
| inorden | izquierda, nodo, derecha | 4 2 5 1 3 |
| postorden | izquierda, derecha, nodo | 4 5 2 3 1 |

- **Caso base**: `n == nil` (un nodo vacío, no hay nada que recorrer).
- **Caso recursivo**: procesar el nodo y llamar a la función con cada hijo.
- Lo único que cambia entre los tres recorridos es **dónde** se pone el `Println` respecto a las llamadas recursivas, igual que en el eco de la montaña.

---

## 9. Recursividad vs. bucles

| | Bucle (`for`) | Recursividad |
|---|---|---|
| Cómo repite | Ejecuta las mismas instrucciones hasta que la condición falla | La función se llama a sí misma con un problema más pequeño |
| Memoria | Poca y constante | Una capa de pila por cada llamada |
| Velocidad | Normalmente más rápido en tareas simples | Un poco más lenta por el costo de las llamadas |
| Legibilidad | Mejor en tareas lineales | Mejor en problemas con estructura repetitiva o jerárquica |
| Riesgo | Bucle infinito | Recursión infinita y *stack overflow* |

Misma tarea, las dos formas (sumar de 1 a n):

```go
package main

import "fmt"

func sumaBucle(n int) int {
    total := 0
    for i := 1; i <= n; i++ {
        total += i
    }
    return total
}

func sumaRecursiva(n int) int {
    if n == 0 {
        return 0
    }
    return n + sumaRecursiva(n-1)
}

func main() {
    fmt.Println(sumaBucle(5))
    fmt.Println(sumaRecursiva(5))
}
```

Ambas dan `15`. Para algo así el bucle es mejor; la recursividad se justifica cuando el problema es naturalmente recursivo.

### Cuándo conviene recursividad

- Recorrer árboles y grafos.
- Backtracking (laberintos, sudoku, N reinas).
- Generar combinaciones y permutaciones.
- Algoritmos "divide y vencerás": quicksort, mergesort, búsqueda binaria.
- Estructuras anidadas, como carpetas dentro de carpetas.

### Cuándo conviene bucle

- Recorrer listas o grandes volúmenes de datos de forma lineal.
- Cuando la profundidad puede ser enorme (miles o millones de pasos).
- Cuando el rendimiento es crítico y el problema no es jerárquico.

---

## 10. Desbordamiento de pila (*stack overflow*)

Este es el error clásico (código **con error a propósito**, no lo uses):

```go
package main

import "fmt"

func sinFin(n int) {
    fmt.Println(n)
    sinFin(n + 1)
}

func main() {
    sinFin(1)
}
```

No tiene caso base, así que nunca para. Cada llamada ocupa pila hasta que se agota y el programa termina con un error fatal `stack overflow`.

Detalles de Go:
- La pila de cada goroutine crece dinámicamente, pero tiene un máximo (alrededor de 1 GB en sistemas de 64 bits). Al pasarlo, el programa muere.
- Este error es **fatal**: no se puede atrapar con `recover`.

---

## 11. Errores comunes

1. **Olvidar el caso base**: recursión infinita y caída del programa.
2. **Caso base inalcanzable**: por ejemplo, llamar con `n+1` cuando se esperaba llegar a 0.
3. **Parámetros mal modificados**: no se acercan al caso base en cada llamada.
4. **Usar recursividad cuando un bucle basta**: más lento y más memoria sin ganar claridad.
5. **Cálculos repetidos** (como Fibonacci simple): se arregla con memoización.
6. **Desbordamiento numérico**: en factoriales y potencias, los números crecen rápido y Go no avisa.
7. **Confundir el orden**: lo que va antes o después de la llamada recursiva cambia el resultado (como en el eco).

---

## 12. Mejores prácticas

- Escribe primero el **caso base**, luego el caso recursivo.
- Verifica que **cada llamada se acerque al caso base**.
- Prueba con valores pequeños (0, 1, 2) antes de valores grandes.
- Usa **memoización** cuando se repiten subproblemas.
- Si la profundidad puede ser muy grande, usa un bucle.
- Documenta con claridad cuál es el caso base y cuál el recursivo.
- **Recursión de cola** (*tail recursion*): es cuando la llamada recursiva es lo último que hace la función. En otros lenguajes se optimiza para no gastar pila, pero **Go no la optimiza automáticamente**, así que no te salva del *stack overflow*.

---

## 13. Casos de uso reales

- Recorridos de árboles y grafos.
- Búsqueda y ordenamiento (quicksort, mergesort, búsqueda binaria).
- Laberintos y problemas de backtracking.
- Generación de combinaciones y permutaciones.
- Recorrer el sistema de archivos (carpetas dentro de carpetas).

---

## 14. Conclusión

La recursividad resuelve un problema grande reduciéndolo a versiones más pequeñas de sí mismo. Lo esencial es un **caso base claro** y llamadas que **se acerquen a él**. Es muy elegante para problemas jerárquicos, pero consume más memoria y puede ser más lenta que un bucle, así que se usa donde aporta claridad.

---

## 15. Ejercicios para practicar

1. Mueve el `Println` del eco de la montaña antes del `if` y predice la salida antes de correrlo.
2. Haz `factorialVisual` con `n = 5` y comprueba que entiendes cada línea de la traza.
3. Escribe una función recursiva que cuente los dígitos de un número (por ejemplo, 12345 tiene 5). Pista: caso base cuando `n < 10`; si no, `1 + cuentaDigitos(n / 10)`.
4. Escribe una función recursiva que sume los dígitos de un número (12345 da 15). Pista: `n % 10 + sumaDigitos(n / 10)`.
5. Invierte una cadena con recursividad.
6. Implementa `fibonacci` simple y `fibMemo`, y compara el tiempo con `n = 40`.
7. Dibuja a mano el árbol de llamadas de `fibonacci(5)`.
8. Agrega un nodo más al árbol binario y predice los tres recorridos.
