# Conversión de tipos en Go

## 1. Por qué existe este tema

Go es un lenguaje de **tipado estático y estricto**: cada variable tiene un tipo fijo, y **Go no convierte tipos automáticamente**. No puedes sumar un `int` con un `float64`, ni un texto con un número, sin convertir primero **de forma explícita**.

Esto es muy frecuente al trabajar con datos que llegan de formularios, archivos, APIs o la consola, porque casi siempre llegan como **texto** (`string`) y hay que pasarlos a número o booleano.

Hay dos herramientas principales:

| Herramienta | Para qué sirve |
|---|---|
| Conversión directa `Tipo(valor)` | Entre tipos compatibles: `int` ↔ `float64`, entre enteros, `string` ↔ `[]byte` / `[]rune` |
| Paquete `strconv` | Entre **texto y números/booleanos**: `string` ↔ `int`, `float`, `bool` |

Regla práctica: si hay un `string` de por medio (leer o escribir un número/booleano en texto), usa `strconv`.

---

## 2. `string` → `bool`: `strconv.ParseBool`

Convierte un texto en un booleano real.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var texto string = "true"
    booleano, err := strconv.ParseBool(texto)
    if err != nil {
        fmt.Println("Error al convertir:", err)
    } else {
        fmt.Println(booleano)
    }
}
```

Salida:

```
true
```

### Valores que acepta `ParseBool`

| Se interpreta como `true` | Se interpreta como `false` |
|---|---|
| `1`, `t`, `T`, `true`, `True`, `TRUE` | `0`, `f`, `F`, `false`, `False`, `FALSE` |

Cualquier otro texto (como `"yes"`, `"si"` o `"abc"`) devuelve un error:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    _, err := strconv.ParseBool("si")
    fmt.Println(err)
}
```

Salida:

```
strconv.ParseBool: parsing "si": invalid syntax
```

---

## 3. `bool` → `string`: `strconv.FormatBool`

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var valorBool bool = false
    texto := strconv.FormatBool(valorBool)
    fmt.Println(texto)
}
```

Salida:

```
false
```

---

## 4. `int` → `string`: `strconv.Itoa`

`Itoa` significa "integer to ASCII".

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var numero int = 42
    texto := strconv.Itoa(numero)
    fmt.Println(texto)
}
```

Salida:

```
42
```

### Trampa clásica: `string(65)` NO da `"65"`

Si conviertes un entero con `string(n)`, Go lo interpreta como un **código de carácter** (un *rune* Unicode), no como dígitos. `go vet` avisa de esto porque casi siempre es un error.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println(string(rune(65)))
    fmt.Println(strconv.Itoa(65))
}
```

Salida:

```
A
65
```

Para convertir un número en sus dígitos como texto, usa siempre `strconv.Itoa` (o `fmt.Sprint`).

---

## 5. `string` → `int`: `strconv.Atoi`

`Atoi` significa "ASCII to integer". Devuelve **dos valores**: el número y un error.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var texto string = "123"
    numero, err := strconv.Atoi(texto)
    if err != nil {
        fmt.Println("Error al convertir:", err)
    } else {
        fmt.Println(numero)
    }
}
```

Salida:

```
123
```

### Qué pasa con entradas inválidas

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    entradas := []string{"25", "abc", "12.5", " 7", "99999999999999999999"}
    for _, e := range entradas {
        n, err := strconv.Atoi(e)
        if err != nil {
            fmt.Printf("%q -> error: %v\n", e, err)
            continue
        }
        fmt.Printf("%q -> %d\n", e, n)
    }
}
```

Salida:

```
"25" -> 25
"abc" -> error: strconv.Atoi: parsing "abc": invalid syntax
"12.5" -> error: strconv.Atoi: parsing "12.5": invalid syntax
" 7" -> error: strconv.Atoi: parsing " 7": invalid syntax
"99999999999999999999" -> error: strconv.Atoi: parsing "99999999999999999999": value out of range
```

Observa:
- `"abc"` no es un número: error de sintaxis (`invalid syntax`).
- `"12.5"` tiene punto decimal: `Atoi` solo acepta enteros.
- `" 7"` tiene un espacio: `Atoi` **no** quita espacios. Si el dato puede traer espacios, límpialo antes con `strings.TrimSpace`.
- Un número demasiado grande da `value out of range`.

---

## 6. `strconv.ParseInt` y `FormatInt`: otras bases y tamaños

`ParseInt(texto, base, bits)` es la versión completa: permite elegir la **base** (2, 8, 10, 16...) y el **tamaño** del entero (8, 16, 32, 64 bits). Devuelve siempre un `int64`.

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    hex, _ := strconv.ParseInt("ff", 16, 64)
    bin, _ := strconv.ParseInt("101", 2, 64)
    fmt.Println(hex)
    fmt.Println(bin)
    fmt.Println(strconv.FormatInt(255, 2))
}
```

Salida:

```
255
5
11111111
```

(Aquí se ignoró el error con `_` solo para que el ejemplo sea corto. En código real, siempre revísalo.)

---

## 7. `string` ↔ `float64`: `ParseFloat` y `FormatFloat`

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    f, err := strconv.ParseFloat("3.14", 64)
    if err != nil {
        fmt.Println("Error al convertir:", err)
        return
    }
    fmt.Println(f * 2)
    fmt.Println(strconv.FormatFloat(3.14159, 'f', 2, 64))
}
```

Salida:

```
6.28
3.14
```

- `ParseFloat(texto, 64)`: el segundo argumento indica la precisión (`64` para `float64`).
- `FormatFloat(valor, 'f', 2, 64)`: formato decimal (`'f'`), con 2 decimales.

---

## 8. `float` → `int`: conversión directa (¡trunca!)

```go
package main

import (
    "fmt"
)

func main() {
    var decimal float64 = 9.99
    entero := int(decimal)
    fmt.Println(entero)
}
```

Salida:

```
9
```

### Truncar no es redondear

`int(...)` **corta** la parte decimal, no redondea. `9.99` da `9`, no `10`. Con negativos, también corta hacia cero: `-9.99` da `-9`.

Si quieres redondear, usa el paquete `math`:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    x := 9.5
    fmt.Println(int(x))
    fmt.Println(math.Round(x))
    fmt.Println(math.Floor(x))
    fmt.Println(math.Ceil(x))
}
```

Salida:

```
9
10
9
10
```

| Función | Qué hace | Con `9.5` |
|---|---|---|
| `int(x)` | Corta los decimales | `9` |
| `math.Round(x)` | Redondea al más cercano | `10` |
| `math.Floor(x)` | Hacia abajo | `9` |
| `math.Ceil(x)` | Hacia arriba | `10` |

Nota: `math.Round`, `Floor` y `Ceil` devuelven `float64`. Si necesitas un entero, envuélvelos: `int(math.Round(x))`.

### Constante vs. variable

`int(9.99)` escrito directamente **no compila**, porque Go ve que la constante perdería sus decimales. Con una variable (`decimal`) sí compila y trunca en tiempo de ejecución.

---

## 9. `int` → `float64`

```go
package main

import "fmt"

func main() {
    a, b := 7, 2
    fmt.Println(a / b)
    fmt.Println(float64(a) / float64(b))
    fmt.Println(float64(a / b))
}
```

Salida:

```
3
3.5
3
```

- `7 / 2` con enteros da `3` (división entera).
- Para obtener `3.5` hay que convertir **antes** de dividir: `float64(a) / float64(b)`.
- `float64(a / b)` convierte **después** de dividir, así que ya se perdió el decimal.

---

## 10. Entre tipos enteros: cuidado con el desbordamiento

Convertir a un tipo más pequeño puede **cambiar el valor sin avisar**:

```go
package main

import "fmt"

func main() {
    var grande int = 300
    fmt.Println(int8(grande))

    var negativo int = -1
    fmt.Println(uint8(negativo))
}
```

Salida:

```
44
255
```

- `int8` guarda de -128 a 127. El 300 "da la vuelta" y queda en `44` (300 - 256).
- `uint8` guarda de 0 a 255. El -1 da la vuelta y queda en `255`.

Go no genera error ni aviso en estos casos. Antes de convertir a un tipo más pequeño, asegúrate de que el valor cabe.

---

## 11. `string` ↔ `[]byte` y `[]rune`

Un `string` en Go es una secuencia de **bytes** en UTF-8. Los caracteres como `ñ` o `á` ocupan **más de un byte**.

```go
package main

import "fmt"

func main() {
    s := "añ"
    fmt.Println(len(s))
    fmt.Println(len([]rune(s)))

    b := []byte("hola")
    fmt.Println(b)
    fmt.Println(string(b))

    fmt.Println([]rune(s))
}
```

Salida:

```
3
2
[104 111 108 97]
hola
[97 241]
```

- `len(s)` cuenta **bytes**: `a` (1) + `ñ` (2) = 3.
- `[]rune(s)` separa por **caracteres**: 2.
- `[]byte("hola")` da los códigos de cada byte.
- `string(b)` reconstruye el texto.

Si vas a trabajar con letras con tilde o emojis, usa `[]rune`.

---

## 12. Otra forma de pasar a `string`: `fmt.Sprint` y `fmt.Sprintf`

Cualquier valor se puede convertir a texto con el paquete `fmt`:

```go
package main

import "fmt"

func main() {
    s1 := fmt.Sprint(42)
    s2 := fmt.Sprintf("Edad: %d", 20)
    s3 := fmt.Sprintf("%.2f", 3.14159)
    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(s3)
}
```

Salida:

```
42
Edad: 20
3.14
```

`strconv` es más rápido y específico; `fmt.Sprintf` es más cómodo cuando quieres dar formato.

---

## 13. Resumen de funciones

| De → A | Función |
|---|---|
| `string` → `bool` | `strconv.ParseBool(s)` |
| `bool` → `string` | `strconv.FormatBool(b)` |
| `string` → `int` | `strconv.Atoi(s)` |
| `int` → `string` | `strconv.Itoa(n)` |
| `string` → `int64` (otra base) | `strconv.ParseInt(s, base, bits)` |
| `int64` → `string` (otra base) | `strconv.FormatInt(n, base)` |
| `string` → `float64` | `strconv.ParseFloat(s, 64)` |
| `float64` → `string` | `strconv.FormatFloat(f, 'f', dec, 64)` |
| `float64` → `int` | `int(f)` (trunca) |
| `int` → `float64` | `float64(n)` |
| `string` → `[]byte` / `[]rune` | `[]byte(s)` / `[]rune(s)` |
| `[]byte` / `[]rune` → `string` | `string(b)` / `string(r)` |

Las funciones `Parse...` y `Atoi` devuelven **valor y error**. Las `Format...` e `Itoa` devuelven solo el texto.

---

## 14. Nota: conversión de tipos vs. aserción de tipos

Lo anterior es **conversión** entre tipos concretos. No lo confundas con la **aserción de tipos** (`valor.(Tipo)`) o el `type switch`, que se usan para extraer el tipo real que hay dentro de una `interface{}` o `any`. Son temas distintos.

---

## 15. Mejores prácticas

- Usa `strconv` para convertir entre texto y números/booleanos.
- **Siempre revisa el error** de `Atoi`, `ParseBool`, `ParseFloat` y `ParseInt`.
- Limpia el texto de entrada con `strings.TrimSpace` si puede traer espacios o saltos de línea.
- Recuerda que de `float` a `int` se **pierde** la parte decimal. Si necesitas redondear, usa `math.Round`.
- Convierte a `float64` **antes** de dividir si quieres un resultado decimal.
- Antes de pasar a un tipo entero más pequeño, comprueba que el valor cabe.
- Lee la documentación oficial del paquete `strconv` para ver todas las funciones disponibles.

---

## 16. Errores comunes

1. **No comprobar el error**: usar `n, _ := strconv.Atoi(s)` y seguir con `n = 0` sin darte cuenta de que falló.
2. **Convertir texto inválido**: por ejemplo `"abc"` a `int`.
3. **Esperar redondeo**: `int(9.99)` da `9`, no `10`.
4. **Usar `string(n)` en vez de `strconv.Itoa(n)`**: da un carácter, no los dígitos.
5. **Dividir enteros y esperar decimales**: `7 / 2` da `3`.
6. **Desbordamiento silencioso** al convertir a `int8`, `uint8`, etc.
7. **Espacios en el texto**: `" 7"` falla con `Atoi`.
8. **Contar caracteres con `len`** en textos con tildes: cuenta bytes, no letras.

---

## 17. Conclusión

La conversión de tipos es una habilidad esencial al trabajar con datos que vienen de fuera del programa. Con `strconv` para texto ↔ números/booleanos, la conversión directa para tipos numéricos y siempre revisando los errores, puedes manejar los datos de forma segura y evitar fallos inesperados.

---

## 18. Ejercicios para practicar

1. Pide un número por consola con `fmt.Scanln`, léelo como `string`, conviértelo con `Atoi` y muestra su doble. Si el usuario escribe algo inválido, muestra un mensaje de error.
2. Convierte las cadenas `"1"`, `"t"`, `"FALSE"` y `"si"` con `ParseBool` y muestra qué pasa con cada una.
3. Convierte `3.99`, `-3.99` y `3.5` a entero de tres formas: `int`, `math.Round` y `math.Floor`. Predice los resultados antes de correr el programa.
4. Calcula el promedio de dos notas enteras (por ejemplo 85 y 90) y muéstralo con decimales.
5. Convierte `255` a binario y a hexadecimal con `FormatInt`, y regresa a decimal con `ParseInt`.
6. Escribe una función `aEntero(s string) (int, error)` que limpie espacios con `strings.TrimSpace` y use `Atoi`.
7. Cuenta cuántas letras tiene `"canción"` con `len` y con `[]rune`, y explica la diferencia.
8. Haz una calculadora que reciba dos números como texto y un operador, los convierta a `float64` y muestre el resultado.
