# Declaración corta de variables en Go (`:=`)

## 1. Definición

La **declaración corta de variables** usa el operador `:=` para **declarar y asignar una variable en un solo paso**, sin escribir la palabra `var` ni el tipo. El compilador **infiere el tipo** a partir del valor.

```go
nombre := "Benjamin Button"
```

Equivale a:

```go
var nombre string = "Benjamin Button"
```

Es la forma más usada de declarar variables **dentro de funciones**, porque el código queda más corto y limpio.

---

## 2. Sintaxis

```
nombreVariable := valor
```

- Va el nombre, luego `:=`, luego el valor.
- El tipo **no se escribe**: lo deduce Go.
- Declara una variable **nueva** en el scope actual.

---

## 3. Ejemplo básico

```go
package main

import "fmt"

func main() {
    nombre := "Benjamin Button"
    edad := 99
    peso := 80
    estudiante := false
    fmt.Println("Nombre: ", nombre)
    fmt.Println("Edad: ", edad)
    fmt.Println("Peso: ", peso)
    fmt.Println("Estudiante: ", estudiante)
}
```

Salida:

```
Nombre:  Benjamin Button
Edad:  99
Peso:  80
Estudiante:  false
```

Hay dos espacios después de los dos puntos: uno está dentro del texto (`"Nombre: "`) y otro lo agrega `Println` al separar los argumentos.

---

## 4. Inferencia de tipos: qué tipo le pone Go

El tipo depende del **valor literal** que escribas:

```go
package main

import "fmt"

func main() {
    entero := 42
    decimal := 3.14
    texto := "hola"
    verdadero := true
    letra := 'a'
    complejo := 2 + 3i

    fmt.Printf("%T\n", entero)
    fmt.Printf("%T\n", decimal)
    fmt.Printf("%T\n", texto)
    fmt.Printf("%T\n", verdadero)
    fmt.Printf("%T\n", letra)
    fmt.Printf("%T\n", complejo)
}
```

Salida:

```
int
float64
string
bool
int32
complex128
```

| Valor | Tipo inferido |
|---|---|
| `42` | `int` |
| `3.14` | `float64` |
| `"hola"` | `string` |
| `true` / `false` | `bool` |
| `'a'` (comillas simples) | `int32` (también llamado `rune`) |
| `2 + 3i` | `complex128` |

`%T` en `Printf` muestra el tipo de una variable. Es muy útil para comprobar qué infirió Go.

### Ojo: `peso := 80` es `int`

Si necesitas decimales, tienes que escribir el decimal: `peso := 80.5` o `peso := 80.0`. Con `80` a secas, Go lo toma como entero.

---

## 5. Reglas del operador `:=`

### Regla 1: solo dentro de funciones

`:=` **no se puede usar fuera de una función** (a nivel de paquete). Para variables globales hay que usar `var`.

```go
package main

import "fmt"

profesion := "Deportista"

func main() {
    fmt.Println(profesion)
}
```

Esto **no compila**:

```
syntax error: non-declaration statement outside function body
```

La forma correcta:

```go
package main

import "fmt"

var profesion = "Deportista"

func main() {
    sueldo := 1000000
    fmt.Println("Profesion: ", profesion)
    fmt.Println("Sueldo: ", sueldo)
}
```

Salida:

```
Profesion:  Deportista
Sueldo:  1000000
```

Aquí `profesion` es **global** (declarada con `var`, fuera de funciones) y `sueldo` es **local** (declarada con `:=`, dentro de `main`).

### Regla 2: crea una variable nueva

`:=` **declara**, no solo asigna. Si la variable ya existe en el mismo scope, da error:

```go
package main

import "fmt"

func main() {
    a := 1
    a := 2
    fmt.Println(a)
}
```

Error:

```
no new variables on left side of :=
```

Para cambiar el valor de una variable existente se usa `=`:

```go
package main

import "fmt"

func main() {
    a := 1
    a = 2
    fmt.Println(a)
}
```

Salida: `2`.

### Regla 3: el tipo queda fijo

Aunque no lo escribas, la variable **sí tiene tipo** y no se puede cambiar después:

```go
package main

func main() {
    x := 5
    x = "hola"
}
```

Esto **no compila**: no se puede asignar un `string` a una variable `int`. Go sigue siendo un lenguaje de tipado estático, solo que deduce el tipo por ti.

### Regla 4: toda variable local declarada debe usarse

Si declaras una variable local y no la usas, Go da error:

```go
package main

func main() {
    x := 5
}
```

Error: `declared and not used: x`.

---

## 6. Declarar varias variables a la vez

```go
package main

import "fmt"

func main() {
    a, b := 1, 2
    nombre, edad := "Ana", 30
    fmt.Println(a, b)
    fmt.Println(nombre, edad)
}
```

Salida:

```
1 2
Ana 30
```

También sirve para intercambiar valores (con `=`):

```go
package main

import "fmt"

func main() {
    a, b := 1, 2
    a, b = b, a
    fmt.Println(a, b)
}
```

Salida: `2 1`.

### Con funciones que devuelven varios valores

Este es el uso más común. Muchas funciones devuelven un resultado y un error:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    numero, err := strconv.Atoi("123")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(numero)
}
```

Salida: `123`.

---

## 7. Redeclaración parcial (el punto más delicado)

Con `:=` y varias variables, Go permite que **algunas ya existan**, siempre que **al menos una sea nueva**. Las que ya existen **se reutilizan** (se les asigna valor), y solo la nueva se crea.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    a, err := strconv.Atoi("1")
    b, err := strconv.Atoi("2")
    fmt.Println(a, b, err)
}
```

Salida:

```
1 2 <nil>
```

- En la primera línea se crean `a` y `err`.
- En la segunda, `b` es nueva y `err` ya existía: Go **reutiliza** `err` y solo crea `b`.
- Si ninguna fuera nueva, daría `no new variables on left side of :=`.

Esta regla solo aplica dentro del **mismo scope**. En un bloque interno, `:=` crea **todas** las variables nuevas (ver sección siguiente).

---

## 8. `:=` y el scope (shadowing)

Dentro de un bloque interno (`if`, `for`, llaves sueltas), `:=` **crea una variable nueva** aunque exista una con el mismo nombre afuera. Esto oculta la externa y es causa frecuente de errores:

```go
package main

import "fmt"

func main() {
    resultado := 1
    if true {
        resultado := 2
        fmt.Println("dentro:", resultado)
    }
    fmt.Println("fuera:", resultado)
}
```

Salida:

```
dentro: 2
fuera: 1
```

Si quieres modificar la variable de afuera, usa `=`, no `:=`.

### Uso intencional en `if`, `for` y `switch`

Declarar variables en el encabezado limita su vida a esa estructura:

```go
package main

import "fmt"

func main() {
    edades := map[string]int{"Ana": 30}
    if e, ok := edades["Ana"]; ok {
        fmt.Println("Ana tiene", e)
    }

    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
}
```

Salida:

```
Ana tiene 30
0
1
2
```

`e`, `ok` e `i` solo existen dentro de su estructura.

---

## 9. Diferencias entre `var` y `:=`

| | `var` | `:=` |
|---|---|---|
| Dónde se puede usar | Dentro y **fuera** de funciones | **Solo dentro** de funciones |
| Tipo | Puede escribirse o inferirse | Siempre inferido |
| Valor inicial | Opcional (usa el valor cero) | **Obligatorio** |
| Redeclaración | No permitida en el mismo scope | Permitida si hay al menos una variable nueva |
| Uso típico | Globales, valores cero, tipo explícito | Locales con valor conocido |

### Formas de `var`

```go
package main

import "fmt"

var global = "visible en todo el paquete"

var (
    ancho  int     = 10
    alto   float64 = 2.5
    titulo         = "rectángulo"
)

func main() {
    var n int
    var s string
    var b bool
    var x, y int = 1, 2
    fmt.Printf("%d %q %t\n", n, s, b)
    fmt.Println(x, y)
    fmt.Println(global)
    fmt.Println(ancho, alto, titulo)
}
```

Salida:

```
0 "" false
1 2
visible en todo el paquete
10 2.5 rectángulo
```

### Valor cero

Con `var` sin valor, Go asigna el **valor cero** del tipo:

| Tipo | Valor cero |
|---|---|
| `int`, `float64` | `0` |
| `string` | `""` (vacío) |
| `bool` | `false` |
| punteros, slices, mapas | `nil` |

Con `:=` no existe esa opción: siempre debes dar un valor.

---

## 10. Cuándo usar `var` en lugar de `:=`

1. **Variables globales** (fuera de funciones).
2. Cuando quieres el **valor cero** sin asignar nada: `var contador int`.
3. Cuando necesitas un **tipo distinto al que infiere Go**.

### El tercer caso, con ejemplo

```go
package main

import "fmt"

func main() {
    precio1 := 10
    precio2 := 10.0
    var precio3 float64 = 10

    fmt.Printf("%T %T %T\n", precio1, precio2, precio3)
    fmt.Println(precio1 / 4)
    fmt.Println(precio3 / 4)
}
```

Salida:

```
int float64 float64
2
2.5
```

`precio1 := 10` se infiere como `int`, y `10 / 4` da `2` (división entera). Con `var precio3 float64 = 10` se fuerza el tipo decimal y da `2.5`.

---

## 11. Usos típicos de `:=`

- Variables locales con valor conocido: `total := 0`.
- Resultados de funciones: `datos, err := leer()`.
- Bucles: `for i := 0; i < n; i++`, `for _, x := range lista`.
- Consultar un mapa: `valor, ok := mapa[clave]`.
- Crear colecciones: `nums := []int{1, 2, 3}`, `edades := map[string]int{}`.
- Funciones anónimas: `suma := func(a, b int) int { return a + b }`.

### El identificador en blanco `_`

Cuando una función devuelve un valor que no necesitas, lo descartas con `_`:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    _, err := strconv.Atoi("abc")
    fmt.Println(err)
}
```

Salida:

```
strconv.Atoi: parsing "abc": invalid syntax
```

`_` no cuenta como variable usada ni sin usar: simplemente descarta el valor.

---

## 12. Errores comunes

1. **Usar `:=` fuera de una función**: `non-declaration statement outside function body`.
2. **Repetir `:=` con la misma variable en el mismo scope**: `no new variables on left side of :=`.
3. **Querer modificar una variable externa con `:=`** dentro de un bloque: crea una nueva y la externa no cambia.
4. **Declarar una variable local y no usarla**: `declared and not used`.
5. **Esperar decimales de un número entero**: `x := 10` es `int`; para decimal, `x := 10.0`.
6. **Intentar cambiar el tipo después**: `x := 5` y luego `x = "hola"`.
7. **Usar `:=` cuando se necesita valor cero**: ahí corresponde `var`.
8. **Confundir `'a'` con `"a"`**: comillas simples es un `rune` (`int32`), comillas dobles es un `string`.

---

## 13. Mejores prácticas

- Usa `:=` dentro de funciones cuando conoces el valor inicial.
- Usa `var` para globales, para valor cero o cuando necesitas un tipo explícito.
- Revisa el tipo inferido cuando importe (con `%T`), sobre todo en números.
- Cuidado con `:=` en bloques internos: puede ocultar variables sin que lo notes.
- Elige nombres claros: `edad`, `totalVentas`, no `x1`, `v2`.
- Aprovecha `_` para descartar valores que no usas.

---

## 14. Resumen

| Concepto | Idea clave |
|---|---|
| `:=` | Declara y asigna, con tipo inferido |
| Dónde | Solo dentro de funciones |
| Globales | Siempre con `var` |
| `=` | Solo asigna a una variable que ya existe |
| Tipo | Se deduce del valor y luego queda fijo |
| Redeclaración | Válida si hay al menos una variable nueva (mismo scope) |
| Bloque interno | `:=` crea variables nuevas (shadowing) |
| Variable local sin usar | Error de compilación |

---

## 15. Conclusión

La declaración corta con `:=` permite escribir código más limpio y conciso dentro de funciones, porque Go infiere el tipo por ti. Para usarla bien hay que recordar sus reglas: solo en funciones, siempre con valor inicial, crea variables nuevas (ojo con el shadowing) y el tipo queda fijo. Para globales, valor cero o tipos explícitos, `var` sigue siendo la herramienta correcta.

---

## 16. Ejercicios para practicar

1. Declara con `:=` un entero, un decimal, un texto y un booleano, e imprime el tipo de cada uno con `%T`.
2. Escribe `x := 10` y `y := 4`, y compara `x / y` con `float64(x) / float64(y)`.
3. Declara una variable global con `var` y una local con `:=`, e imprime ambas.
4. Intenta usar `:=` fuera de una función y lee el mensaje de error.
5. Escribe `a := 1` y luego `a := 2` en el mismo scope, lee el error y corrígelo.
6. Usa `a, err := strconv.Atoi("5")` y luego `b, err := strconv.Atoi("x")`. Explica qué variable se crea y cuál se reutiliza.
7. Crea un bloque `if` con `resultado := 2` y comprueba que no cambia el `resultado` de afuera. Luego arréglalo con `=`.
8. Declara `var contador int`, `var nombre string` y `var activo bool`, e imprime sus valores cero.
9. Usa `_` para descartar el número en `_, err := strconv.Atoi("abc")` y muestra el error.
