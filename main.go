package main
import "fmt"

const HELLO_USER string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"

func main(){
  var name string
  lastname := "apellido"
  var (
    a = 1
    b = "dos"
    c = false
  )
  fmt.Print("Ingresa tu nombre:")
  fmt.Scanf("%s", &name)
  fmt.Print("Ingresa tu apellido:")
  fmt.Scanf("%s", &lastname)
  fmt.Printf(HELLO_USER, name, lastname)
  fmt.Printf("%b %s %t\n",a,b,c)
}
