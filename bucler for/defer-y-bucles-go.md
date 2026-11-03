# Go: `defer` y bucles (`while` y `do while`)

Apuntes de práctica para Go.

---

## 1. ¿Qué es `defer`?

`defer` aplaza la ejecución de una función hasta que la función que la contiene termina (por `return` o por un `panic`).

No libera memoria RAM por sí mismo. Sirve para **limpiar recursos**: cerrar archivos, conexiones, soltar un candado (mutex), etc. Lo que logra es que el cierre quede escrito justo al lado de la apertura, para no olvidarlo.

### Cómo funciona paso a paso

```go
package main

import "fmt"

func main() {
    defer fmt.Println("al final")
    fmt.Println("primero")
    fmt.Println("segundo")
}
```

1. Go llega a `defer` y **no lo ejecuta**: lo anota en una lista.
2. Sigue normal: imprime "primero".
3. Imprime "segundo".
4. `main` termina y entonces ejecuta lo anotado: imprime "al final".

Salida:

```
primero
segundo
al final
```

Analogía: es como dejarte una nota "antes de salir de casa, apagar la luz". La escribes al entrar, pero la haces al salir.

### Reglas importantes

- Varios `defer` se ejecutan en **orden inverso** (el último en declararse corre primero).
- Los argumentos de la función diferida se **evalúan al escribir el `defer`**, no cuando se ejecuta.
- Se ejecutan **al terminar la función**, no al terminar el bloque ni el bucle.

### Orden inverso (LIFO)

```go
package main

import "fmt"

func main() {
    fmt.Println("abro archivo")
    defer fmt.Println("cierro archivo")

    fmt.Println("leo datos")

    defer fmt.Println("defer A")
    defer fmt.Println("defer B")

    fmt.Println("fin del trabajo")
}
```

Salida:

```
abro archivo
leo datos
fin del trabajo
defer B
defer A
cierro archivo
```

Se anotaron en este orden: "cierro archivo", A, B. Al terminar `main` salen al revés: B, A, "cierro archivo".

---

## 2. Ejemplo real con archivos y funciones

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func escribirArchivo(nombre string) error {
	f, err := os.Create(nombre)
	if err != nil {
		return err
	}
	defer f.Close()
	defer fmt.Println("fin de escribirArchivo")

	fmt.Println("escribiendo...")
	fmt.Fprintln(f, "linea 1: hola Go")
	fmt.Fprintln(f, "linea 2: aprendiendo defer")
	return nil
}

func leerArchivo(nombre string) error {
	f, err := os.Open(nombre)
	if err != nil {
		return err
	}
	defer f.Close()
	defer fmt.Println("fin de leerArchivo")

	fmt.Println("leyendo...")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

func main() {
	nombre := "datos.txt"

	if err := escribirArchivo(nombre); err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := leerArchivo(nombre); err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("listo")
}
```

Salida esperada:

```
escribiendo...
fin de escribirArchivo
leyendo...
linea 1: hola Go
linea 2: aprendiendo defer
fin de leerArchivo
listo
```

`defer f.Close()` se escribe justo después de abrir el archivo. Se anota primero, así que es lo último que se ejecuta, después de todo lo que hagas con el archivo.

---

## 3. Bucles en Go: no existe `while` ni `do while`

Go solo tiene la palabra `for`, pero con ella se hacen los tres tipos de bucle.

| Otros lenguajes | En Go |
|---|---|
| `for` clásico | `for i := 0; i < n; i++ { }` |
| `while (condicion)` | `for condicion { }` |
| `do { } while (condicion)` | `for { ...; if !condicion { break } }` |
| bucle infinito | `for { }` |

### `for` clásico

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

Salida: `0 1 2 3 4` (uno por línea).

### `while`: se hace con `for` y solo la condición

La condición se revisa **antes** de cada vuelta. Si es falsa desde el inicio, nunca entra.

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

Salida: `0 1 2 3 4`.

Caso donde no entra ni una vez:

```go
package main

import "fmt"

func main() {
    n := 10
    for n < 5 {
        fmt.Println("esto no se imprime")
        n++
    }
    fmt.Println("terminó")
}
```

Salida: `terminó`.

### `do while`: se hace con `for {}` y `break` al final

El `do while` ejecuta el cuerpo **al menos una vez** y revisa la condición **después**. En Go se escribe un `for` infinito y se sale con `break` cuando la condición de continuar deja de cumplirse.

```go
package main

import "fmt"

func main() {
    n := 10
    for {
        fmt.Println("n =", n)
        n++
        if !(n < 5) {
            break
        }
    }
}
```

Salida: `n = 10`. Aunque `n < 5` es falso desde el inicio, el cuerpo corrió una vez.

### Diferencia entre `while` y `do while` (mismo valor inicial)

```go
package main

import "fmt"

func main() {
    x := 10

    fmt.Println("while:")
    for x < 5 {
        fmt.Println("x =", x)
        x++
    }

    x = 10
    fmt.Println("do while:")
    for {
        fmt.Println("x =", x)
        x++
        if !(x < 5) {
            break
        }
    }
}
```

Salida:

```
while:
do while:
x = 10
```

### Bucle infinito con salida

```go
package main

import "fmt"

func main() {
    contador := 0
    for {
        contador++
        if contador == 3 {
            continue
        }
        if contador > 5 {
            break
        }
        fmt.Println(contador)
    }
}
```

Salida: `1 2 4 5`. `continue` salta a la siguiente vuelta y `break` sale del bucle.

### Ejemplo práctico de `do while`: pedir un dato hasta que sea válido

```go
package main

import "fmt"

func main() {
    var edad int
    for {
        fmt.Print("Escribe tu edad (1-120): ")
        fmt.Scan(&edad)
        if edad >= 1 && edad <= 120 {
            break
        }
        fmt.Println("Edad inválida, intenta de nuevo")
    }
    fmt.Println("Edad guardada:", edad)
}
```

Pregunta al menos una vez y repite hasta que el valor sea válido: es el uso típico de un `do while`.

---

## 4. `defer` dentro de un bucle

Cuidado: `defer` no corre al final de cada vuelta, corre al final de la **función**.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Println("defer", i)
        fmt.Println("vuelta", i)
    }
    fmt.Println("fin de main")
}
```

Salida:

```
vuelta 0
vuelta 1
vuelta 2
fin de main
defer 2
defer 1
defer 0
```

Se anotan tres `defer` y salen al revés al terminar `main`.

### Si quieres cerrar algo en cada vuelta: mete el trabajo en una función

```go
package main

import "fmt"

func procesar(i int) {
    defer fmt.Println("cierro recurso", i)
    fmt.Println("uso recurso", i)
}

func main() {
    for i := 0; i < 3; i++ {
        procesar(i)
    }
}
```

Salida:

```
uso recurso 0
cierro recurso 0
uso recurso 1
cierro recurso 1
uso recurso 2
cierro recurso 2
```

Así el `defer` corre al terminar cada llamada a `procesar`, o sea, en cada vuelta. Es el patrón correcto al abrir archivos dentro de un bucle.

---

## 5. Ejercicios para practicar

1. Escribe un programa con tres `defer` y tres `Println` normales. Predice la salida antes de correrlo.
2. Haz un `while` que cuente de 10 a 1 hacia atrás.
3. Haz un `do while` que pida una contraseña hasta que el usuario escriba "go123".
4. Crea un archivo con `os.Create`, escribe 3 líneas y ciérralo con `defer f.Close()`. Luego léelo en otra función.
5. Haz un bucle que procese 3 archivos (`a.txt`, `b.txt`, `c.txt`) usando una función con `defer` para cerrar cada uno.
