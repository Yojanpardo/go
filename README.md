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
