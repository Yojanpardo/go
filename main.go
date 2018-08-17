package main
import "fmt"

func main(){
  var name string
  fmt.Print("Ingresa tu nombre:")
  fmt.Scanf("%s", &name)
  fmt.Print("Hola, %s",name)
}
