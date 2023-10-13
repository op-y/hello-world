package main

var ch = make(chan struct{})
var s string

func f() {
    s = "hello, world"
    <-ch
}

func main() {
    go f()
    ch <- struct{}{}
    print(s)
}
