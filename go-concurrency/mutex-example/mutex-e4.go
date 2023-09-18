package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    CounterType int
    Name string

    mu sync.Mutex
    count uint64
}

func (c *Counter) Incr() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}

func (c *Counter) Count() uint64 {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}

func main() {
    var counter Counter

    var wg sync.WaitGroup
    wg.Add(10)

    for i := 0; i < 10; i++ {
        go func() {
            defer wg.Done()
            for j := 0 ; j < 100000; j++ {
                counter.Incr()
            }
        } ()
    }

    wg.Wait()
    fmt.Println(counter.Count())
}
