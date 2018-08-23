package main
import "fmt"

func main(){
  var names []string
  fmt.Println("vamos a hacer una lista de amigos.\nCuantos amigos tienes?")
  var friends int
  fmt.Scanf("%d"&friends)
  for i := 0; i<friends; i++{
    if i == 0{
      fmt.Println("Ingresa un nombre")
      fmt.Scanf("%s",&names[i])
    }else if i > 0 {
      fmt.Println("Ingresa otro nombre")
      fmt.Scanf("%s",&names[i])      
    }
  }
}
