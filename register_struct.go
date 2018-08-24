package main

import "fmt"

type Course struct{
  name string
  id int
  url string
}

func (c Course) Subscribe(name string){
  fmt.Printf("La persona %s fue registrada correctamente al curso de %s\n", name, c.name)
}

func main(){
  course := Course{name: "Curso profesional de Go", id:1, url:"/clases/go-basico"}
  course.Subscribe("yojan")
}
