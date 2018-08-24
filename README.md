# Desarrollo con Go
Go es el nuevo lenguaje de programación de google, este repo contiene todos los archivos y documentos que voy a ir creando a medida que avanzo en el curso profesional de desarrollo con Go de Platzi.
* go es un lenguaje multiplataforma
* utiliza el paradigma imperativo y compilado
* la sintaxis es muy similar a C
* Tipado estático.
* tiene algunas caracteristicas de lenguajes interpretados.

## por que go?
* Es facil de aprender
* la velocidad de ejecución es bastante rapida
* Concurrencia (permite el manejo de multiples hilos de forma sencilla)
* Librerias estandar muy completas
* linter definido desde el principio == codigo legible
* Facil despligue (gracias a que es compilado)

## Instalacion de go en linux
Primero debemos agregar el repositorio y luego hacemos la instalacion de la siguiente manera.
~~~sh
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
~~~

## organizando el entorno de trabajo
Una peculiaridad de Go es que no se puede crear una carpeta en cualquier lugar y empezar a desarrollar en ella, es necesario crear un GOPATH y por lo general va a la carpeta /home/usuario/go/ pero se puede poner en lo documentos o donde sea. go tiene una forma peculiar para funcionar y es que debe estar todo bajo una gerarquia de carpetas y para ello debemos hacer lo siguiente
~~~sh
export GOPATH=/home/user/Documents/go
cd home/user/Documents/go
mkdir src
mkdir src/github.com
mkdir src/github.com/tu_usuario_de_github
#ahora podemos crear todos nuestros proyectos aquí
~~~

## Creando nuestro primer hola mundo
Dentro del proyecto hay un archivo llamado main.go y contiene mi primer hola mundo.
en la consola podemos utilizar un comando para la ejecucion y otro para la compilacion.
~~~go
package main
import "fmt"

func main(){
  fmt.Print("hola mundo")
}
~~~
~~~sh
go run main.go #para compilar y ejectuar
go build main.go #para compilar y genera un archivo con el mismo nombre pero sin extensión
./main #ejecuta el archivo creado
~~~

## Variables y constantes
Dentro de Go podemos utilizar variables y contantes, y tienen sus respectivas formas para ser declaradas, una contante por lo general va a ser global. de la siguiente manera se declaran variables y constantes:
~~~go
package main
import "fmt"

const HELLO_USER string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"//así se declara una constante

func main(){
  var name string //esta es la declaración de una variable de tipo string
  lastname := "apellido" //así tambien se puede declarar una variable y el compilador detecta de que tipo es
  var (
    a = 1
    b = "dos"
    c = false
  )//Aqui estoy declarando varias variables al tiempo y el compilador detecta automaticamente de que tipo son
  fmt.Print("Ingresa tu nombre:")
  fmt.Scanf("%s", &name)
  fmt.Print("Ingresa tu apellido:")
  fmt.Scanf("%s", &lastname)
  fmt.Printf(HELLO_USER, name, lastname)
  fmt.Printf("%b %s %t\n",a,b,c)//aqui las formateo para imprimirlas
}
~~~

## funciones
Podemos utilizarlas para no repetir codigo mientras escribimos nuestros programas

## Tipos de datos numericos
En go podemos tener en cuenta los usuarios que utilizan computadoras de 32 y 64 bits, la diferencia es que los numeros de 64 bits son mas precisos ya que pueden tener muchos mas valores.

## Cadenas de texto
Go tiene soporte nativo para utf-8, asi que no es necesario hacer algo especial para que funcionen nuestros programas en español. Podemos tambien recorrer un estring utilizando los indices, obtener la longitud de una palabra con len()

## arrays y slices
Son dos tipos de vectores, los arrays son fijos y los slices son dinamicos.

## switch case
Dentro de Go tambien podemos encontrar el condicional switch, el cual nos permite hacer uso de diferentes opciones y se utiliza de la siguiente manera
~~~go
switch variable {
case caso_con_la_variable:
    //aqui va la logica
  default:
    //acciones por default  
}
//tambien se puede hacer con condiciones
switch {
case <condition>:
  //TODO code here
default:
  //logica aquí
}
~~~

## Manejo de paquetes en go
Al igual que en otros lenguajes de programacion, en go se puede hacer uso de manejo de paquetes, un paquete es una carpeta que contiene archivos para ejecutar.

## Map
Dentro de go existe un tipo de archivo de llave valor que se llama Map, su uso es similar al de los diccionarios en python pero su declaracion es algo diferente, su forma de declararlo es asi:
~~~go
//map[<tipo de dato de la llave>]<tipo de dato del valor>
amigos_telefono := make(map[string]int)
amigos_telefono["llave1"]=1
//tambien se puede asignar valores rapidamente
mapTest := [string]int{
  "llave1": 1,
  "llave2": 2,
}
~~~

## structs
Son un tipo de dato especial que usa go y es parecido a crear un objeto ya que puede contener varios datos y se pueden crear algo como objetos a partir de ellos. su forma de usarlos es asi:
~~~go
type Person struct{
  name string
  id int
  age int
}
func main(){
  person1 := Person{name:"laura",id:2554558,age:21}
  person2 := new(Person)
  person2.name = "yojan"
  person2.id = 1532586664
  person2.age = 21
}
~~~
para hacerlo de forma mas interactiva podemos llamar el struct a un metodo y lo pasamos como un parametro para poder utilizarlo.
~~~go
type Course struct{
  name string
  id int
  url string
}

func main(){
  course := Course{Name: "Curso profesional de Go", id:1, url:"/clases/go-basico"}
  course.Subscribe("yojan")
}

func (c course) Subscribe(name string){
  fmt.Printf("La persona %s fue registrada correctamente al curso de ", name, c.name)
}
~~~
