package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
    ch4 := make(chan int)

    go prt(ch1, ch2)
    go prt(ch2, ch3)
    go prt(ch3, ch4)
    go prt(ch4, ch1)

    ch1 <- 1

    select {}
}

func prt(in, out chan int) {
    for {
        id := <-in
        fmt.Println(id)
        time.Sleep(time.Second)
        if id % 4 == 0 {
            id = 1
        } else {
            id++
        }
        out <- id
    }
}


//type Token struct{}
//
//func newWorker(id int, ch chan Token, nextCh chan Token) {
//    for {
//        token := <-ch         // 取得令牌
//        fmt.Println((id + 1)) // id从1开始
//        time.Sleep(time.Second)
//        nextCh <- token
//    }
//}
//
//func main() {
//    chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}
//
//    // 创建4个worker
//    for i := 0; i < 4; i++ {
//        go newWorker(i, chs[i], chs[(i+1)%4])
//    }
//
//    //首先把令牌交给第一个worker
//    chs[0] <- struct{}{}
//    select {}
//}
