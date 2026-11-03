# Scope (alcance) de variables en Go

## 1. Definición

El **scope** (o **alcance**) de una variable determina **en qué partes del código esa variable existe y se puede usar**.

En Go, el alcance se define por el **bloque** donde se declara la variable. Un bloque es el código entre llaves `{}`. Entenderlo evita errores como `undefined: variable` y ayuda a escribir código limpio y fácil de mantener.

Regla principal:

> Una variable es visible desde el punto donde se declara hasta el final del bloque donde fue declarada, **incluyendo los bloques anidados hacia adentro**, pero **no hacia afuera**.

---

## 2. Niveles de scope en Go

Go tiene varios niveles de bloques, de mayor a menor:

| Nivel | Dónde se declara | Dónde se puede usar |
|---|---|---|
| Universo | Nombres incorporados del lenguaje (`len`, `int`, `true`, `nil`, `println`...) | En todo el programa |
| Paquete (global) | Fuera de cualquier función | En todos los archivos del mismo paquete |
| Archivo | Los `import` | Solo en ese archivo |
| Función | Parámetros y variables dentro de una función | Solo dentro de esa función |
| Bloque | Dentro de cualquier `{}` (`if`, `for`, `switch`, o llaves sueltas) | Solo dentro de ese bloque |

Cuando se habla de "scope global y local en Go":
- **Global** = a nivel de paquete (fuera de las funciones).
- **Local** = dentro de una función o de un bloque.

---

## 3. Ejemplo de scope con error

```go
package main

import "fmt"

var var1 = "Este es el nivel 1"

func main() {
    var var2 = "Este es el nivel 2"
    {
        var var3 = "Este es el nivel 3"
    }
    fmt.Println(var1)
    fmt.Println(var2)
    fmt.Println(var3)
}
```

Este código **no compila**. El error principal es:

```
undefined: var3
```

### Por qué falla

Hay tres niveles de scope en este programa:

```
nivel 1: var1   -> declarada fuera de las funciones (paquete)
nivel 2: var2   -> declarada dentro de main (función)
nivel 3: var3   -> declarada dentro de las llaves sueltas { } (bloque)
```

- `var1` se ve en todas partes.
- `var2` se ve en todo `main`, incluyendo el bloque interno.
- `var3` solo existe **dentro de sus llaves**. Al salir del bloque desaparece, y el `Println(var3)` de afuera no la encuentra.

Además, Go prohíbe **variables locales sin usar**. En este código `var3` se declara pero nunca se usa dentro de su bloque, así que Go también puede quejarse con un error del tipo `declared and not used`. Con el arreglo de abajo se resuelven ambos problemas.

---

## 4. Versión corregida

```go
package main

import "fmt"

var var1 = "Este es el nivel 1"

func main() {
    var var2 = "Este es el nivel 2"
    {
        var var3 = "Este es el nivel 3"
        fmt.Println(var3)
    }
    fmt.Println(var1)
    fmt.Println(var2)
}
```

Salida:

```
Este es el nivel 3
Este es el nivel 1
Este es el nivel 2
```

Ahora `var3` se usa **dentro** de su bloque, donde sí existe.

### Qué ve cada nivel

| Variable | Dentro del bloque interno | En `main`, fuera del bloque | En otras funciones |
|---|---|---|---|
| `var1` (paquete) | Sí | Sí | Sí |
| `var2` (función) | Sí | Sí | No |
| `var3` (bloque) | Sí | No | No |

---

## 5. Una variable existe desde donde se declara, no desde el inicio del bloque

Dentro de una función, no puedes usar una variable **antes** de declararla:

```go
package main

import "fmt"

func main() {
    fmt.Println(x)
    x := 5
    fmt.Println(x)
}
```

Esto **no compila** (`undefined: x`), porque en la primera línea `x` todavía no existe.

### Excepción: nivel de paquete

Las variables globales **sí** pueden declararse en cualquier orden, porque Go resuelve sus dependencias:

```go
package main

import "fmt"

var a = b + 1
var b = 5

func main() {
    fmt.Println(a)
}
```

Salida:

```
6
```

---

## 6. Variables globales (nivel de paquete)

Una variable declarada fuera de las funciones es visible para **todas las funciones del paquete**:

```go
package main

import "fmt"

var contador = 0

func incrementar() {
    contador++
}

func main() {
    incrementar()
    incrementar()
    fmt.Println(contador)
}
```

Salida:

```
2
```

### Reglas de las globales

- Se declaran con `var` (**no** se puede usar `:=` fuera de una función).
- Pueden quedar sin usar (el error de "variable sin usar" aplica solo a locales).
- Si tienen nombre con **mayúscula inicial** (`Version`), son **exportadas** y otros paquetes pueden usarlas. Con minúscula (`version`), solo se ven dentro de su paquete.

```go
package config

var Version = "1.0"
var secreto = "abc"
```

Desde otro paquete se puede usar `config.Version`, pero no `config.secreto`.

### Por qué conviene usar pocas globales

- Cualquier función puede modificarlas, así que es difícil saber quién cambió qué.
- Dificultan las pruebas y el mantenimiento.
- Mejor: declara la variable **en el scope más pequeño posible** y pásala como parámetro.

---

## 7. Parámetros y variables de función

Los parámetros de una función son locales a ella:

```go
package main

import "fmt"

func duplicar(n int) int {
    resultado := n * 2
    return resultado
}

func main() {
    fmt.Println(duplicar(4))
}
```

Salida:

```
8
```

`n` y `resultado` solo existen dentro de `duplicar`. En `main` no se pueden usar.

---

## 8. Variables dentro de `if`, `for` y `switch`

Estas estructuras crean su propio bloque. Las variables declaradas en su encabezado **viven solo ahí**.

### En `if`

```go
package main

import "fmt"

func main() {
    if n := 5; n > 3 {
        fmt.Println("n es mayor que 3:", n)
    } else {
        fmt.Println("n no es mayor que 3:", n)
    }
}
```

Salida:

```
n es mayor que 3: 5
```

La variable `n` se ve en el `if` y en su `else`, pero **no después**. Si escribes `fmt.Println(n)` debajo, da `undefined: n`.

### En `for`

```go
package main

import "fmt"

func main() {
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
}
```

Salida:

```
0
1
2
```

`i` solo existe dentro del `for`. Fuera de él, `i` no está definida.

### Si necesitas el valor después del bucle

Declara la variable **antes**, en el scope exterior:

```go
package main

import "fmt"

func main() {
    var i int
    for i = 0; i < 3; i++ {
    }
    fmt.Println(i)
}
```

Salida:

```
3
```

---

## 9. Shadowing (sombreado de variables)

Si en un bloque interno declaras una variable con **el mismo nombre** que una externa, la interna **oculta** a la externa mientras dura ese bloque. La externa no se modifica.

```go
package main

import "fmt"

func main() {
    x := 10
    {
        x := 20
        fmt.Println("dentro:", x)
    }
    fmt.Println("fuera:", x)
}
```

Salida:

```
dentro: 20
fuera: 10
```

Son **dos variables distintas** que se llaman igual.

### Con una variable global

```go
package main

import "fmt"

var nombre = "global"

func mostrar() {
    fmt.Println(nombre)
}

func main() {
    nombre := "local"
    fmt.Println(nombre)
    mostrar()
}
```

Salida:

```
local
global
```

`main` crea su propia variable `nombre`, y `mostrar` sigue viendo la global.

### El error clásico: `:=` en un bloque interno

Esta es la causa de muchos errores difíciles de ver:

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

Quien lo escribió probablemente quería cambiar `resultado` a 2, pero `:=` **crea una variable nueva** en el bloque del `if`. La solución es usar `=` (asignación):

```go
package main

import "fmt"

func main() {
    resultado := 1
    if true {
        resultado = 2
        fmt.Println("dentro:", resultado)
    }
    fmt.Println("fuera:", resultado)
}
```

Salida:

```
dentro: 2
fuera: 2
```

### Resumen de `:=`, `=` y `var`

| Forma | Qué hace |
|---|---|
| `x := 5` | **Declara** una variable nueva en el scope actual (solo dentro de funciones) |
| `x = 5` | **Asigna** a una variable que ya existe |
| `var x = 5` | Declara una variable (sirve dentro y fuera de funciones) |

Si usas `:=` con una variable que ya existe **en el mismo scope** y no hay ninguna nueva en el lado izquierdo, Go da el error `no new variables on left side of :=`.

### Sombrear nombres del lenguaje

Los nombres del universo (`len`, `int`, `true`...) también se pueden sombrear. Compila, pero es mala idea:

```go
package main

import "fmt"

func main() {
    len := 3
    fmt.Println(len)
}
```

Salida: `3`. Pero a partir de ahí, en esa función, `len(...)` ya no sirve como función. Evita usar esos nombres.

---

## 10. Scope y closures

Una función definida dentro de otra puede **ver las variables de la función que la rodea** (esto se llama closure):

```go
package main

import "fmt"

func main() {
    mensaje := "hola"
    mostrar := func() {
        fmt.Println(mensaje)
    }
    mostrar()
}
```

Salida:

```
hola
```

La función anónima `mostrar` ve `mensaje` porque está dentro de su alcance, siguiendo la misma regla de "se ve hacia adentro".

---

## 11. Scope vs. ciclo de vida

Son dos ideas distintas:

| | Qué significa |
|---|---|
| **Scope** | **Dónde** en el código se puede usar el nombre de la variable |
| **Ciclo de vida** | **Cuánto tiempo** existe la variable en memoria |

Normalmente una variable local deja de existir cuando termina su bloque. Pero con **closures** puede seguir viva más allá de su scope:

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
    c := crearContador()
    fmt.Println(c())
    fmt.Println(c())
}
```

Salida:

```
1
2
```

`cuenta` está fuera de scope para `main` (no puede escribir `cuenta` directamente), pero **sigue viva** porque la función devuelta la usa.

---

## 12. Errores comunes con el scope

1. **Usar una variable fuera de su bloque**: `undefined: x`.
2. **Declarar una variable local y no usarla**: `declared and not used`.
3. **Usar `:=` por error en un bloque interno** y crear una variable nueva sin querer (shadowing).
4. **Usar una variable antes de declararla** dentro de una función.
5. **Usar `:=` fuera de una función**: error `non-declaration statement outside function body`. A nivel de paquete hay que usar `var`.
6. **Abusar de variables globales**: cualquier parte del código puede cambiarlas.
7. **Esperar que la variable de un `for` o `if` siga existiendo después** de la estructura.
8. **Usar un nombre con minúscula esperando que otro paquete lo vea**: solo las exportadas (mayúscula) se ven fuera.

---

## 13. Buenas prácticas

- Declara cada variable en el **scope más pequeño posible**.
- Evita las variables globales; pásalas como parámetros.
- Evita el shadowing, o úsalo con intención. Cuando quieras modificar una variable externa, usa `=`, no `:=`.
- Usa la **indentación** para ver los niveles de scope con claridad.
- Declara variables de uso temporal en el encabezado de `if`, `for` o `switch`.
- No uses nombres como `len`, `string` o `int` para tus variables.
- Revisa los mensajes del compilador: `undefined` casi siempre es un problema de scope.

---

## 14. Resumen

| Concepto | Idea clave |
|---|---|
| Bloque | Código entre `{}` |
| Visibilidad | Hacia adentro sí, hacia afuera no |
| Global | Fuera de funciones; se ve en todo el paquete |
| Local | Dentro de una función o bloque |
| Exportada | Nombre con mayúscula inicial; visible desde otros paquetes |
| Shadowing | Variable interna con el mismo nombre oculta la externa |
| `:=` | Declara una variable **nueva** |
| `=` | Asigna a una variable existente |

---

## 15. Conclusión

El alcance de las variables en Go se controla con los bloques `{}`. Entenderlo permite controlar la visibilidad y el ciclo de vida de las variables, evitar errores como `undefined`, y mantener el código mejor organizado y más fácil de mantener.

---

## 16. Ejercicios para practicar

1. Escribe un programa con tres niveles de scope (global, función y bloque) e imprime cada variable en el lugar donde sí se puede ver. Luego intenta imprimir la del bloque desde afuera y lee el error.
2. Predice la salida de un programa que tenga `x := 1`, un bloque interno con `x := 2` y otro `Println` afuera. Luego córrelo.
3. Cambia el `:=` del bloque interno por `=` y compara la salida.
4. Declara una variable global y modifícala desde dos funciones distintas. Después refactoriza para usar parámetros y valores de retorno en lugar de la global.
5. Escribe un `for` con `i := 0` en el encabezado e intenta usar `i` después del bucle. Corrígelo declarando `i` antes.
6. Haz un `if` con una variable en el encabezado (`if n := ...; n > 0 {}`) y úsala también en el `else`.
7. Crea un closure que use una variable de la función que lo rodea, y comprueba que sigue funcionando después de que esa función terminó.
8. Crea un paquete con una variable exportada y una no exportada, y prueba cuál se puede usar desde `main`.
