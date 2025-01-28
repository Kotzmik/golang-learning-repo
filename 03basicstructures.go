package main

import (
  "fmt"
  // "slices" // przydatna biblioteka do sliców
  // "maps" to samo co wyżej 
)

func main() {

  // ARRAYS - zwykła lista zmiennych, wymagany typ, ilosć zaalokowanej pamięci jest stała i przypisana podczas kompilacji (?)
  var a[5]int
  fmt.Println("deklaracja pustej tabeli:", a)
  a[4]=10
  fmt.Println(a)
  fmt.Println(a[4])
  fmt.Println("len:",len(a))

  b:=[5]int{1,2,3,4,5}
  fmt.Println("deklaracja wypełnionej tabeli(ustalona liczba indexów):", b)
  b=[...]int{2,3,4,5,6}
  fmt.Println("deklaracja wypełnionej tabeli(kompilator liczy ilość elementów):", b)
  b = [...]int{100, 3: 400, 500}
  fmt.Println("idx:", b)
  var twoD [2][3]int
  for i := 0; i < 2; i++ {
    for j := 0; j < 3; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d:", twoD)


  // SLICES - tablica dynamiczna, nie ma przypisanej stałej długości, jedynie typ zmiennych 
  var c []string
  fmt.Println("uninit:", c, c == nil, len(c) == 0)
  c=make([]string, 3)
  fmt.Println("emp:", c, "len:", len(c), "cap:", cap(c))
  c[1]="dupa"
  c[0]="siku"
  c[2]="i"
  fmt.Println(c)
  cc:=append(c, "kupa")
  fmt.Println(c, cc)

  d:= make([]string, len(cc))
  copy(d, cc)
  fmt.Println(d)
  fmt.Println(cc[2:3])
  fmt.Println(cc[2:])
  fmt.Println(cc[:4])


  // MAPS - mapa klucz-wartość
  m := make(map[string]int)
  m["k1"]=21
  m["k2"]=37
  fmt.Println("map:", m)
  v1:=m["k1"]
  fmt.Println(v1)
  v3 := m["k3"]
  fmt.Println("nieistniejący klucz-wartość:", v3)

  clear(m)
  fmt.Println(m)

  _, prs := m["k2"]
  fmt.Println("prs:", prs)

  m["k2"]=69
  _, prs = m["k2"]
  fmt.Println("prs:", prs)
  
  n := map[string]int{"foo": 1, "bar": 2}
  fmt.Println("map:", n)
}
