package main

import (
    "fmt"
    "sync"
)

var wg = &sync.WaitGroup{}

var a string

func f() {
    fmt.Println(a)
    wg.Done()
}

//func hello() {
func main() {
    wg.Add(1)

    a = "hello, world"
    go f()

    wg.Wait()
}
