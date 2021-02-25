package main

import (
    "fmt"
    "reflect"
)

func main() {
    var a = "string"
    fmt.Println(a)
    fmt.Println(reflect.TypeOf(a))
    fmt.Println("-----------------------")
    
    
    var b,c int=1,2
    fmt.Println(b,c)
    fmt.Println(reflect.TypeOf(b),reflect.TypeOf(c))
    fmt.Println("-----------------------")
    
    var d = true
    fmt.Println(d)
    fmt.Println(reflect.TypeOf(d))
    fmt.Println("-----------------------")
    
    var e int
    fmt.Println(e)
    fmt.Println(reflect.TypeOf(e))
    fmt.Println("-----------------------")
    
    f:= "apple"
    fmt.Println(f)
    fmt.Println(reflect.TypeOf(f))
    fmt.Println("-----------------------")
    
    var g,h = 2,3
    fmt.Println(g,h)
    fmt.Println(reflect.TypeOf(g),reflect.TypeOf(h))
    fmt.Println("-----------------------")
    
    var i string="amol"
    fmt.Println(i)
    fmt.Println(reflect.TypeOf(i))
    fmt.Println("-----------------------")
    
    var j = 4.56
    fmt.Println(j)
    fmt.Println(reflect.TypeOf(j))
    fmt.Println("-----------------------")
    
    var k float64 = 7.89
    fmt.Println(k)
    fmt.Println(reflect.TypeOf(k))
    fmt.Println("-----------------------")
    
}

