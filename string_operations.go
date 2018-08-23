package main
import (
  "fmt"
  "strings"
)

func main(){
  var text string
  fmt.Println("ingresa un texto y lo vamos a operar de diferentes maneras\nsi vas a ingresar mas de una palabra entonce serparalas por guiones")
  fmt.Scanf("%s",&text)
  fmt.Println(strings.ToUpper(text))//el texto que se pasa se vuelve todo mayusculas
  fmt.Println(strings.ToLower(text))//todo se vuelve  inusculas
  fmt.Println(strings.Replace(text, "hello","hola",-1))//el primer parametro es la variable a cambiar, lo que vamos a buscar para cambias, por lo que amos a cambiar y cuantos
  fmt.Println(strings.Split(text,"-"))
}
