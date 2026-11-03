# ¿Qué es la sentencia `defer` en Go?

## Definición

`defer` permite **aplazar la ejecución de una función hasta que la función que la contiene termina**. Es clave para gestionar recursos y hacer tareas de limpieza: cerrar conexiones, borrar archivos temporales, soltar candados (mutex), etc.

Es útil porque garantiza que ciertas acciones se ejecuten **sin importar el camino que siga el programa**. Por ejemplo, si abres una conexión a una base de datos, puedes asegurar que se cierre al final, incluso si ocurre un error a mitad del proceso.

> Nota: `defer` se ejecuta al terminar la **función** (no el bloque `if` o `for` donde está escrito), y no libera memoria RAM por sí mismo. Lo que hace es ejecutar tu código de limpieza (por ejemplo `Close()`) en el momento correcto.

---

## Ejemplo práctico

```go
package main

import "fmt"

func desconectar() {
    fmt.Println("Se ha desconectado de la base de datos")
}

func leer() {
    fmt.Println("Se han leído los registros de la base de datos")
}

func actualizar() {
    fmt.Println("Se han actualizado registros de la base de datos")
}

func conectar() {
    fmt.Println("Se ha conectado a la base de datos")
}

func main() {
    conectar()
    defer desconectar()
    leer()
    actualizar()
}
```

### Explicación

Se definen cuatro funciones: `conectar`, `leer`, `actualizar` y `desconectar`. El flujo en `main` es:

1. `conectar()`
2. `defer desconectar()`
3. `leer()`
4. `actualizar()`

Gracias a `defer`, `desconectar()` se ejecuta automáticamente **al final de `main`**, después de las demás operaciones. Así se garantiza liberar el recurso sin importar cómo termine la función.

### Salida

```
Se ha conectado a la base de datos
Se han leído los registros de la base de datos
Se han actualizado registros de la base de datos
Se ha desconectado de la base de datos
```

---

## Ventajas de usar `defer`

- Código más limpio y seguro: evita olvidar liberar recursos.
- La limpieza queda junto a la apertura del recurso, en un solo lugar.
- Mejora la legibilidad y el mantenimiento.

---

## Diferencia entre `defer` y `return`

| | `return` | `defer` |
|---|---|---|
| Qué hace | Termina la función y devuelve un valor | Pospone una función hasta que la función actual termina |
| Cuándo actúa | En ese momento | Justo antes de que la función salga |
| Retorno anticipado | Sale de la función | Se ejecuta igualmente, aunque haya un `return` temprano |

---

## ¿Cuándo usar `defer`?

Úsalo cuando necesites garantizar la liberación de recursos:

- Cerrar archivos (`defer f.Close()`).
- Cerrar conexiones de red o de base de datos.
- Soltar bloqueos (`defer mu.Unlock()`).

---

## Varios `defer` en una función

Se pueden usar varios `defer` en la misma función. Se ejecutan en **orden inverso** al que se declararon: el último `defer` es el primero en ejecutarse.

```go
package main

import "fmt"

func main() {
    defer fmt.Println("uno")
    defer fmt.Println("dos")
    defer fmt.Println("tres")
    fmt.Println("trabajo")
}
```

Salida:

```
trabajo
tres
dos
uno
```

---

## Conclusión

`defer` es una herramienta poderosa para gestionar recursos y mantener el código limpio. Usarlo bien hace las aplicaciones en Go más robustas y fáciles de mantener.
