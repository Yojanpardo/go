package main
import "fmt"

func main(){
  var names []string
  var name string
  friends := 0
  for { //así es un ciclo infinito, como un while True
    for friends == 0{ // esto es un for como un while
      fmt.Println("vamos a hacer una lista de amigos.\nCuantos amigos tienes?")
      fmt.Scanf("%d", &friends)
      if friends != 0{
        for i := 0; i<friends; i++{ //esto es un for normalito
          if i == 0{
            fmt.Println("Ingresa un nombre")
            fmt.Scanf("%s",&name)
            names = append(names,name)
          }else if i > 0 {
            fmt.Println("Ingresa otro nombre")
            fmt.Scanf("%s",&name)
            names = append(names,name)
          }
        }
      }else{
        fmt.Println("ah cabrón! has de tener al menos un amigo, no?")
      }
    }
    fmt.Println("tus amigos son: ", names)
    break
  }
}
