package main

import "fmt"

func plus(a int, b int) int {return a + b}
func plusPlus(a, b, c int) int { return a + b + c }
func vals() (int, int) {return 6,9} // funckja zwracająca więcej ninż 1 wartość
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total:=0

    for _, num :=range nums { // pierwsza zmienna to index slajsa
        fmt.Println("num:",num)
        total += num
    }
    // for i:=0; i<len(nums); i++ {
    //     total+=nums[i]
    // }
    fmt.Println(total)
}
func licznikhehe() func() int { //domknięcie
    i:=0
    return func() int {
        i++
        return i
    }
}
func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

    a, b :=vals()
    fmt. Println(a, b)
    sum(1, 2, 3, 5)
    numerki:=[]int{9,9,7}
    sum(numerki...)
    // fmt.Printf("%T", numerki...) sprawdzanie czy to slice, tak nie działa
    fmt.Println(numerki)
    
    nextInt:=licznikhehe()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    newInt:=licznikhehe()
    fmt.Println(newInt())

    fmt.Println("XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
    var fib func(n int) int
    fib=func(n int) int {
        if n<2 {
            return n
        }
        return fib(n-1)+fib(n-2)
    }
    for i:=2; i<9; i++ {
    fmt.Println(fib(i))
    }
}
