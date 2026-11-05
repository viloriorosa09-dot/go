# Tipos de datos en Go.

*Go es un lenguaje fuertemente tipado y estático, lo que significa que cada variable y constante tiene un tipo de dato definido en tiempo de compilación. Comprender los **tipos de datos primitivos en Golang** es fundamental para escribir código eficiente y seguro. Los tipos de datos más comunes en Go incluyen:*

- **int:** Números enteros, pueden ser positivos o negativos.
- **float32, float64:** Números decimales de precisión simple y doble.
- **string:** Cadenas de texto.
- **bool:** Valores booleanos (true o false).
- **byte:** Alias para uint8, útil para manipulación de datos binarios.
- **rune:** Alias para int32, representa un punto de código Unicode.
- **array:** Colección de elementos del mismo tipo y tamaño fijo.
- **slice:** Colección dinámica de elementos del mismo tipo.
- **map:** Colección de pares clave-valor.
- **struct:** Estructura personalizada que agrupa diferentes tipos de datos.

## Cómo declarar variables en Go correctamente

Para declarar variables en Go, se utiliza la palabra clave `var`, seguida del nombre de la variable y su
tipo. Por ejemplo:
 
` var edad int 
  var nombre string
  var activo bool `

* También es posible declarar e inicializar variables en una sola línea: *

` var ciudad string = "Punta Cana"
  var temperatura float64 = 23.5 `

* Gracias a la inferencia de tipos, se puede omitir el tipo si se asigna un valor inicial: *

`pais := "Republica Dominicana"
 poblacion := 11000000`


## Diferencias entre tipos de datos en Go

Una de las **diferencias entre tipos de datos en Go** y otros lenguajes es la claridad con la que se deben definir los tipos y la imposibilidad de mezclar tipos en operaciones sin conversión explícita. Por ejemplo, no se puede sumar un entero y un flotante directamente sin convertir uno de los valores:

` var a int = 5
  var b float64 = 6.2
  // var c = a + b // Error: mismatched types int and float64
  var c float64(a) + b // correcto `

  **Esto obliga al programador a ser consciente de los tipos que utiliza, lo que reduce la probabilidad de errores en tiempo de ejecución.**

 ## Errores comunes con tipos de datos en Go

Uno de los **errores comunes con tipos de datos en Go** es intentar asignar un valor de un tipo a una variable de otro tipo sin conversión explícita. Por ejemplo:

` var edad int = 30
  // edad = "treinta" // Error: cannot use "treinta" (type string) as type int `

Otro error frecuente es olvidar el tamaño de los arrays, ya que en Go el tamaño es parte del tipo:

var numeros [3]int
// numeros = [1, 2, 3, 4] // Error: cannot use [1, 2, 3, 4] (type [4]int) as type [3]int

Para evitar estos problemas, es recomendable siempre verificar los tipos y utilizar conversiones explícitas cuando sea necesario.

## Ventajas de los tipos estáticos en Go

Las **ventajas de los tipos estáticos en Go** incluyen la detección temprana de errores, mayor rendimiento y facilidad de mantenimiento. Al saber exactamente qué tipo de dato se espera en cada variable y función, el compilador puede optimizar el código y advertir sobre posibles inconsistencias antes de ejecutar el programa.

Además, el tipado estático facilita el trabajo en equipo, ya que cada miembro puede entender rápidamente qué tipo de datos se manejan en cada parte del código, reduciendo la posibilidad de errores y mejorando la calidad general del software.

## Ejemplo de una función en Go utilizando un argumento

En el presente ejemplo se ha creado una función `hola`, la cual recibe un parámetro `s` de tipo string. Esta función a su vez retorna un valor tipo string.

` package main
 
  import "fmt 

  func hola(s string) string {
    return "hola " + s
  }
  
  fund main (){
    fmt.println(hola("wilber"))
  }`

  Este ejemplo ilustra cómo se define y utiliza una función con parámetros tipados en Go, reforzando la importancia de especificar el tipo de cada argumento para evitar errores y mejorar la legibilidad del código.
## Función en Go que utiliza dos parámetros enteros

En el siguiente ejemplo tenemos una función que suma 2 parámetros enteros recibidos, y que retorna también un valor entero.

` package main 

  import "fmt"

  func sumar(a int , b int ) int {
    return a + b 
  }

  func main() {

    fmt.println(sumar( 5, 6))
  }

  Si por ejemplo utilizamos un string en lugar de un entero como argumento, el compilador lanzará un error.

 `003_tipado/main.go:10:23: cannot use "a" (type string) as type int in argument to sumar`

Este tipo de error es común cuando no se presta atención a los tipos de datos esperados por las funciones. Por eso, es fundamental comprender **cómo declarar variables en Go correctamente** y las **ventajas de los tipos estáticos en Go** para escribir código robusto y eficiente.

# Buenas prácticas y recomendaciones

- Utiliza la inferencia de tipos (`:=`) para simplificar la declaración de variables cuando el tipo es evidente.

- Prefiere tipos de datos específicos (por ejemplo, `int32`, `float64`) cuando sea necesario controlar el    tamaño y la precisión.

- Realiza conversiones explícitas entre tipos para evitar errores de compilación.
    
- ocumenta las funciones y estructuras, especificando claramente los tipos de datos utilizados.
    
- Aprovecha las herramientas de análisis estático de Go para detectar posibles inconsistencias y mejorar la   calidad del código.

# Conclusión

Comprender y utilizar correctamente los tipos de datos en Go es esencial para desarrollar aplicaciones seguras, eficientes y fáciles de mantener. El tipado estático de Go ayuda a detectar errores antes de ejecutar el programa y promueve buenas prácticas de  programación. Dominar la declaración de variables, la conversión de tipos y el uso adecuado de los tipos primitivos y compuestos permitirá aprovechar al máximo las capacidades del lenguaje y evitar problemas comunes en el desarrollo de software.