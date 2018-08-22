package main
import "fmt"

func main(){
  fmt.Println("ingresa un numero y verificaremos si es par e impar")
  var number int
  fmt.Scanf("%d", &number)
  if number%2 == 0{
    fmt.Printf("el numero %d es par\n",number)
  }else{
    fmt.Printf("el numero %d es impar\n",number)
  }
}
