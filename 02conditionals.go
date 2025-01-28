package main

import (
  "fmt"
  "time"
)

func main() {
  a:=4

  if num :=a; num < 0 { // jakie to ma zastosowanie?
      fmt.Println(num, "is negative")
  } else if num < 10 {
      fmt.Println(num, "has 1 digit")
  } else {
      fmt.Println(num, "has multiple digits")
  }

  switch time.Now().Weekday() {
  case time.Saturday, time.Sunday:
    fmt.Println("It's the weekend")
  default:
    fmt.Println("It's a weekday")
  }

  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("It's before noon")
  default:
    fmt.Println("It's after noon")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
    case bool:
      fmt.Println("I'm a bool")
    case int:
      fmt.Println("I'm an int")
    default:
      fmt.Printf("Don't know type %T\n", t)
    }
  }
  whatAmI(true)
  whatAmI(1)
  whatAmI("hey")

  x:=213.7
  fmt.Printf("%T", x) // najszybszy i najprosszy sposób na sprawdzenie typu zmiennej
}
