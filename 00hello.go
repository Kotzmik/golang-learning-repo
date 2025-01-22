package main

import "fmt"

func main() {
  fmt.Println(2.1+37)
  fmt.Println(2.1/3.7)
  var a="dupa"
  fmt.Println("a = "+a)

  b:=1337 //deklaracja przypisuje automatycznie typ zmiennej
  // b="1336" nie zadziała

  fmt.Println(b)
  // fmt.Println(a + b)
  // mieszanie typów nie działa w taki sposób
  var c int
  fmt.Println(c)
  
  const d=5e6
  fmt.Println(d)
  fmt.Println(int64(d))



  fmt.Println("#####")
}
