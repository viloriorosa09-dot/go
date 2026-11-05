# ¿Qué son los tipos numéricos?

*En el lenguaje de  programación Go, los tipos numéricos son fundamentales para el desarrollo de aplicaciones robustas y eficientes. Permiten trabajar con valores enteros y flotantes, facilitando la realización de cálculos  matemáticos precisos y el manejo de datos cuantitativos. **Cómo declarar variables numéricas en Go** es una de las primeras habilidades que todo programador debe dominar, ya que la correcta elección del tipo de dato impacta directamente en el rendimiento y la seguridad de los programas.*

Go ofrece una variedad de tipos numéricos, cada uno diseñado para cubrir necesidades específicas en cuanto a rango, precisión y uso de memoria. En este artículo, exploraremos en profundidad los tipos numéricos en Go, centrándonos en los enteros y los flotantes, sus diferencias, límites y mejores prácticas para su utilización. Cursosde mates

## Enteros

Los enteros son números sin parte decimal y se dividen en diferentes tamaños y signos en Go. Comprender las **diferencias entre int y uint en Golang** es esencial para evitar errores y aprovechar al máximo los recursos del sistema.

## int

El tipo `int` se adapta al tamaño de la arquitectura subyacente, pudiendo ser de 32 o 64 bits dependiendo del sistema operativo y el hardware. Esto significa que, en sistemas modernos de 64 bits, un `int` puede almacenar valores mucho mayores que en sistemas de 32 bits. AprenderCódigo

`var edad int = 33`

El uso de `int` es común cuando no se requiere un tamaño específico y se busca portabilidad entre diferentes plataformas. Sin embargo, es importante tener en cuenta que el rango de valores puede variar, por lo que en aplicaciones críticas es recomendable especificar el tamaño exacto del entero.

## int8, int16, int32, int64

Estos tipos enteros tienen un tamaño fijo de 8, 16, 32 y 64 bits, respectivamente. El número que sigue a `int`indica el número de bits utilizados para almacenar el valor. Por ejemplo, `int32` puede almacenar valores desde -2,147,483,648 hasta 2,147,483,647. Cursosde mates

`var temperatura int16 = -15
var poblacion int64 = 7800000000`

La elección del tipo adecuado depende del rango de valores que se espera manejar y del uso eficiente de la memoria.

## uint8, uint16, uint32, uint64

Los tipos `uint` son enteros sin signo, es decir, solo pueden representar valores positivos. Esto permite duplicar el rango de valores positivos en comparación con sus equivalentes con signo. Por ejemplo, `uint8` puede almacenar valores de 0 a 255. Cursosde mates

`var edad uint8 = 25`

Utilizar tipos sin signo es útil cuando se sabe que los valores nunca serán negativos, como en el caso de conteos, índices o tamaños de estructuras. Sin embargo, es fundamental tener cuidado al realizar operaciones entre tipos con y sin signo para evitar errores de conversión.

 *"Además de estos tipos enteros básicos, Go también proporciona alias de tipo para facilitar la legibilidad del código, como byte (equivalente a uint8) y rune (equivalente a int32). Estos alias son especialmente útiles al trabajar con cadenas de texto y caracteres Unicode. Cursosde mates"*

## Flotantes

Los números flotantes representan valores con parte decimal y son esenciales para cálculos científicos, financieros y cualquier aplicación que requiera precisión decimal. En Go, existen dos tipos principales de flotantes: `float32` y `float64`. Saber **cuándo usar float32 o float64 en Go** es clave para balancear precisión y uso de memoria.

`const PI float32 = 3.1416`

## float32

Este tipo flotante de 32 bits puede representar números en coma flotante con una precisión de aproximadamente 7 dígitos decimales. Es adecuado para aplicaciones donde la memoria es limitada y la precisión no es crítica. AprenderCódigo

## float64

El tipo float64 utiliza 64 bits y ofrece una mayor precisión, permitiendo representar números con hasta 15 dígitos decimales. Es el tipo por defecto para operaciones  matemáticas en Go y se recomienda para la mayoría de los casos donde la precisión es importante.

`var gravedad float64 = 9.80665`

 *"Es importante tener en cuenta que los números flotantes no pueden representar con precisión todos los valores reales debido a las limitaciones de la representación binaria. Por lo tanto, es posible que se produzcan pequeños errores de redondeo en operaciones aritméticas con números flotantes. Cursosde mates"*

## Ejemplos de tipos numéricos

La siguiente tabla describe los tipos de valores numéricos posibles y sus rangos. Conocer los **límites de los tipos numéricos en Go** es fundamental para evitar desbordamientos y errores inesperados.

`tipo	    tamaño	    rango
int8	    8 bits	    -128 a 127
int16	    16 bits	    -2^15 a 2^15 -1
int32   	32 bits	    -2^31 a 2^31 -1
int64	    64 bits	    -2^63 a 2^63 -1
uint8	    8 bits	    0 a 255
uint16	    16 bits 	0 a 2^16 -1
uint32	    32 bits	    0 a 2^32 -1
uint64	    64 bits	    0 a 2^64 -1`


## Ejemplos de valores decimales expresados como binarios

La representación binaria es la base del almacenamiento de datos en computadoras. A continuación, algunos ejemplos de cómo se representan los valores decimales en binario: Cursosde mates

`Binario 	Decimal
0000	    0
0001	    1
0010	    2
0011	    3
0100	    4
1111	    16`

## Acerca de los valores numéricos

Todos estos constituyen posibles tipos numéricos, sin embargo, a nivel de la computadora todos se guardan como valores binarios (0 y 1).` Un bit equivale ya sea a un 0 o un 1`. Para un campo de 4 bits, por ejemplo, se pueden almacenar 16 números diferentes, lo que ilustra la importancia de comprender la relación entre el tamaño del tipo y la cantidad de valores posibles. Cursosde mates

También es importante hacer notar que los valores enteros “signed” incluyen números positivos y negativos, mientras que los unsigned solo números positivos. Esta diferencia es crucial al diseñar estructuras de datos y algoritmos eficientes.

El número máximo que puede ser utilizado por una computadora depende de la arquitectura de esta. En los equipos de 64 bits de hoy en día, esto representa el entero **9223372036854775807** para un `int64`.

## Valores máximos y mínimos

En Go, podemos **obtener los valores máximos y mínimos que pueden ser representados por los tipos numéricos** utilizando las constantes predefinidas en el paquete `math`. Por ejemplo, podemos acceder al valor máximo de un entero de 32 bits utilizando` math.MaxInt32` y al valor mínimo de un flotante de 64 bits utilizando `math.SmallestNonzeroFloat64`.

`import "math"

maxEntero32 := math.MaxInt32
minFlotante64 := math.SmallestNonzeroFloat64`


Estas constantes son útiles para validar datos de entrada y evitar errores de desbordamiento. Además, permiten escribir código más seguro y robusto, especialmente en aplicaciones críticas.

 *"Los tipos numéricos en Go nos brindan la capacidad de representar y manipular valores enteros y flotantes en nuestros programas. Con una variedad de operaciones y funciones matemáticas disponibles, podemos realizar cálculos precisos y tomar decisiones basadas en valores numéricos. Cursosde mates"*

## Operaciones matemáticas con tipos numéricos en Go

Go proporciona un conjunto completo de operadores y funciones matemáticas para trabajar con tipos numéricos. Entre las operaciones más comunes se encuentran la suma, resta, multiplicación, división y módulo. Además, el paquete `math` ofrece funciones avanzadas como potencias, raíces, trigonometría y logaritmos.

**Operaciones matemáticas con tipos numéricos en Go** se realizan de manera eficiente y segura, siempre que se respeten los tipos de datos involucrados. Es importante recordar que las operaciones entre diferentes tipos requieren conversión explícita para evitar errores de compilación. Cursosde mates

`a := 10
b := 3
resultado := a / b // División entera, resultado: 3

x := 10.0
y := 3.0
resultadoFlotante := x / y // División flotante, resultado: 3.333...`

La conversión entre tipos se realiza utilizando la sintaxis:

`var entero int = 42
var flotante float64 = float64(entero)`

El uso adecuado de los tipos numéricos y las conversiones garantiza la precisión y el correcto funcionamiento de los algoritmos.

## Conclusión
Los tipos numéricos en Go son una herramienta esencial para el desarrollo de aplicaciones eficientes y seguras. Comprender sus características, límites y diferencias permite elegir el tipo adecuado para cada situación, optimizando el uso de recursos y evitando errores comunes. La correcta declaración y manipulación de variables numéricas es clave para escribir código robusto y mantenible en Go.