# Guía fácil de Go: condicionales y operadores

Esta guía junta seis temas que se usan siempre juntos. Primero se explica cada uno y después cómo se conectan.

Puedes probar todo el código sin instalar nada en **go.dev/play**.

**Temas:**
1. `if`
2. `else`
3. `else if`
4. Operadores de comparación
5. Operadores aritméticos
6. Estructuras de control (visión general)

---

## 1. La sentencia `if`

`if` ejecuta un bloque de código **solo si** una condición es verdadera. Si es falsa, se lo salta.

```go
package main

import "fmt"

func main() {
    var edad uint8
    fmt.Println("¿Cuál es tu edad?")
    fmt.Scanln(&edad)

    if edad >= 18 {
        fmt.Println("Eres mayor de edad")
    }
}
```

**Cómo leerlo:**
- `var edad uint8` crea una variable para guardar la edad (número entero sin signo, de 0 a 255).
- `fmt.Scanln(&edad)` lee lo que escribe el usuario y lo guarda en `edad`. El `&` le dice a Go dónde guardarlo.
- `if edad >= 18 { ... }` revisa la condición. Si es verdadera, entra a las llaves.

**Reglas de Go:**
- La condición **no** lleva paréntesis (se escribe `if edad >= 18`, no `if (edad >= 18)`).
- Las llaves `{ }` son obligatorias, aunque haya una sola instrucción.
- La llave de apertura va en la **misma línea** que el `if`.

**Buena práctica:** mantén las condiciones simples y evita meter muchos `if` uno dentro de otro, porque el código se vuelve difícil de leer.

---

## 2. La sentencia `else`

`else` es el camino alternativo: se ejecuta **cuando la condición del `if` es falsa**.

```go
package main

import "fmt"

func main() {
    var edad uint8
    fmt.Println("¿Cuál es tu edad?")
    fmt.Scanln(&edad)

    if edad >= 18 {
        fmt.Println("Eres mayor de edad")
    } else {
        fmt.Println("Eres menor de edad")
    }
}
```

**Idea clave:** con `if / else` siempre se ejecuta **uno** de los dos bloques, nunca los dos y nunca ninguno.

**Regla de Go:** `else` va en la misma línea que la llave de cierre del `if`: `} else {`.

**Buena práctica:** usa `else` solo cuando de verdad necesites un camino alternativo.

---

## 3. La sentencia `else if`

Cuando hay **más de dos opciones**, encadenas condiciones con `else if`. Go las revisa **en orden** y ejecuta la primera que sea verdadera.

```go
package main

import "fmt"

func main() {
    var nota int
    fmt.Println("Escribe tu nota:")
    fmt.Scanln(&nota)

    if nota >= 90 {
        fmt.Println("Excelente")
    } else if nota >= 70 {
        fmt.Println("Aprobado")
    } else {
        fmt.Println("Reprobado")
    }
}
```

**Cómo funciona con la nota 95:**
1. ¿`95 >= 90`? Sí → imprime "Excelente" y **termina** (no revisa lo demás).

**Con la nota 75:**
1. ¿`75 >= 90`? No.
2. ¿`75 >= 70`? Sí → imprime "Aprobado".

**Con la nota 40:**
1. ¿`40 >= 90`? No.
2. ¿`40 >= 70`? No.
3. Ninguna se cumplió → entra al `else` → "Reprobado".

**Cuidado con el orden:** si pusieras primero `nota >= 70`, una nota de 95 entraría ahí y diría "Aprobado". Siempre empieza por la condición más específica.

---

## 4. Operadores de comparación

Comparan dos valores y siempre devuelven un booleano: `true` (verdadero) o `false` (falso). Son los que se usan dentro de los `if`.

| Operador | Significa | Ejemplo | Resultado |
|---|---|---|---|
| `==` | igual a | `2 == 1` | `false` |
| `!=` | diferente de | `2 != 1` | `true` |
| `<` | menor que | `2 < 1` | `false` |
| `<=` | menor o igual que | `2 <= 1` | `false` |
| `>` | mayor que | `2 > 1` | `true` |
| `>=` | mayor o igual que | `2 >= 1` | `true` |

**Error muy común:** confundir `=` con `==`.
- `=` **guarda** un valor en una variable.
- `==` **compara** dos valores.

**Reglas:**
- Los dos valores deben ser del **mismo tipo**. Comparar un número con un texto da error de compilación.
- Puedes comparar números entre sí y textos (strings) entre sí.
- Los slices y los mapas no se pueden comparar directamente con `==` (solo con `nil`).

---

## 5. Operadores aritméticos

Hacen cálculos matemáticos.

| Operador | Significa | Ejemplo | Resultado |
|---|---|---|---|
| `+` | suma | `5 + 3` | `8` |
| `-` | resta | `5 - 3` | `2` |
| `*` | multiplicación | `5 * 3` | `15` |
| `/` | división | `5 / 3` | `1` (con enteros) |
| `%` | residuo (módulo) | `5 % 3` | `2` |

**Lo más importante: la división entre enteros devuelve un entero.**
- `5 / 3` con enteros da `1` (se pierden los decimales).
- `5.0 / 3.0` con decimales (`float64`) da `1.6666666666666667`.

**Go es de tipado fuerte:** no puedes mezclar tipos sin convertirlos.

```go
package main

import "fmt"

func main() {
    var a int = 5
    var b int = 3

    fmt.Println(a / b)
    fmt.Println(float64(a) / float64(b))
    fmt.Println(a % b)
}
```

Salida:

```
1
1.6666666666666667
2
```

`float64(a)` convierte el entero `a` en decimal para que la división conserve los decimales.

**Truco útil:** para sacar un porcentaje con enteros, multiplica primero y divide después. Por ejemplo, el 10% de `total` es `total * 10 / 100`.

**Para qué sirve `%`:** para saber si un número es par (`n % 2 == 0`) o si una división es exacta (`a % b == 0`).

---

## 6. Estructuras de control (visión general)

Las estructuras de control deciden **qué parte del código se ejecuta y cuántas veces**. En Go las principales son:

| Estructura | Para qué sirve |
|---|---|
| `if` / `else` / `else if` | Ejecutar código según una condición |
| Operadores de comparación | Armar las condiciones (`==`, `<`, `>=`...) |
| Operadores aritméticos | Hacer cálculos dentro del programa y de las condiciones |
| Operadores lógicos (`&&`, `||`, `!`) | Combinar varias condiciones en una sola |
| `switch` | Elegir entre muchas opciones de forma ordenada |
| `for` | Repetir acciones (es el único bucle de Go) |
| `defer` | Dejar una función para el final, cuando termine la función que la contiene |

Los seis temas de esta guía son la base. Las demás estructuras las verás después.

---

## Cómo se conectan los temas

Los seis temas forman una cadena:

1. Los **operadores aritméticos** calculan valores (`total * 10 / 100`, `a % b`).
2. Los **operadores de comparación** comparan esos valores y dan `true` o `false` (`total >= 1000`, `a % b == 0`).
3. El **`if`** usa ese `true` o `false` para decidir si entra a su bloque.
4. El **`else`** cubre el caso contrario.
5. El **`else if`** agrega más caminos cuando hay más de dos casos.
6. Todo eso forma parte de las **estructuras de control**.

