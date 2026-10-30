# ¿Cuántas formas existen para pasar variables a una función en Go?

En el lenguaje Go, existen dos métodos principales para enviar información a una función:**diferencia entre paso de variables por valor y por referencia en Go**. Es fundamental comprender ambos enfoques para escribir código eficiente y seguro.

- Enviar una copia del valor de la variable. 

- un puntero, es decir, la dirección de memoria de la variable.

## Ejemplo de paso de argumentos por valor en Go

Un **ejemplo de paso por valor en Go** se observa cuando una función recibe una copia de la variable original. Cualquier modificación dentro de la función no afecta el valor fuera de ella.

`func equivalenciaEnPies(altura float32) float32 {
    altura = altura * 3.28
    return altura
}

var altura float32 = 1.70

fmt.Println("La altura es:", altura, "mts")
fmt.Println("La altura es:", equivalenciaEnPies(altura), " pies")
fmt.Println("La nueva altura es:", altura, "mts")`

`La altura es: 1.7 mts
La altura es: 5.576  pies
La nueva altura es: 1.7 mts`

 *"En este caso, la función convierte una altura de metros a pies. El valor original de la variable altura permanece sin cambios, ya que solo se trabaja con una copia."*

## Ejemplo de paso de argumentos por referencia en Go

Cuando se utiliza el paso por referencia, se envía la dirección de memoria de la variable. Esto permite modificar el valor original desde la función, lo que ilustra claramente las **ventajas y desventajas de 
pasar variables por referencia en Go.**

`func conversionEnPies(altura *float32) float32 {
    *altura = *altura - 0.10
    return *altura
}

var altura float32 = 1.70

fmt.Println("La altura es:", altura, "mts")
fmt.Println("Al envejecer:", conversionEnPies(&altura), "mts")
fmt.Println("Despues de envejecer:", altura, "mts")`

`La altura es: 1.7 mts
Al envejecer: 1.6 mts
Despues de envejecer: 1.6 mts`

 *"En este ejemplo, la función modifica directamente el valor de la variable original, ya que trabaja con su puntero."*

## ¿Cuándo usar punteros en Go?

Saber **cuándo usar punteros en Go** es clave para optimizar el rendimiento y la gestión de memoria. Los punteros son útiles cuando se necesita modificar el valor original de una variable dentro de una función, o cuando se desea evitar la copia de grandes estructuras de datos.

## Mejores prácticas con punteros en Go

Al trabajar con punteros, es importante seguir las **mejores prácticas con punteros en Go**: evitar el uso innecesario de punteros en tipos primitivos, asegurarse de que los punteros no apunten a memoria liberada y documentar claramente las funciones que modifican valores a través de referencias.

## Conclusión

Comprender la diferencia entre pasar variables por valor y por referencia en Go es esencial para escribir código robusto y eficiente. Elegir el método adecuado depende del contexto y de los objetivos de tu función, permitiendo un mejor control sobre el comportamiento y la memoria de tus programas.
