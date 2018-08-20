package main
import "fmt"

const HELLO_USER string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"

func main(){
  fmt.Print("Ingresa tu nombre:")
  name := get_text()
  fmt.Print("Ingresa tu apellido:")
  lastname := get_text()
  age, email,phone_number := get_personal_data()
  fmt.Printf(HELLO_USER, name, lastname)
  fmt.Printf("tu edad: %d.\ntu email: %s.\ntu numero de telefono: %s.\n",age,email,phone_number)
  get_hobbies()
  get_devices()
}

func get_text() string{
  var text string
  fmt.Scanf("%s", &text)
  return text
}

func get_personal_data() (int,string,string){
  var age int
  var email string
  var phone_number string
  fmt.Println("Que edad tienes?")
  fmt.Scanf("%d",&age)
  fmt.Println("cual es tu email?")
  fmt.Scanf("%s",&email)
  fmt.Println("Cual es tu numero de telefono?")
  fmt.Scanf("%s",&phone_number)
  return age,email,phone_number
}

func get_hobbies(){
  var arr1 [2]string
  fmt.Println("ingresa un hobbie")
  fmt.Scanf("%s",&arr1[0])
  fmt.Println("ingresa otro hobbie")
  fmt.Scanf("%s",&arr1[1])
  fmt.Printf("tus hobbies son %s y %s.\n",arr1[0],arr1[1])
}

func get_devices(){
  var devices []string
  fmt.Println(devices)
  devices = append(devices,"celular")
  fmt.Println(devices)
  devices = append(devices,"computador")
  fmt.Println(devices)
}
