package main

import "fmt"

func main() {
  var i int

  fmt.Println("loop 0")
  for i<6 {
    fmt.Println(i)
    i++
  }

  fmt.Println("loop 1")
  for j:=0; j<6; j++ {
    fmt.Print(j,"\n") //Print nie dodaje newlinów
  }

  fmt.Println("loop 2")
  for i:=range 6 {
    fmt.Println("range",i)
  }

  fmt.Println("loop 3")
  for {
    fmt.Print("while true")
    break
  }

}
