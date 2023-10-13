package trace

import (
    "fmt"
)

func init() {
	fmt.Println("init trace")
}

func Trace(t string, v int) int {
    fmt.Println(t, ":", v)
    return v
}
