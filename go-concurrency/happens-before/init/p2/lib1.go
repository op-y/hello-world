package p2

import (
    "fmt"

    "demo-init/p3"
    "demo-init/trace"
)

var V1_p2 = trace.Trace("init v1_p2", 2)
var V2_p2 = trace.Trace("init v2_p2", p3.V2_p3)

func init() {
    fmt.Println("init func in p2")
    V1_p2 = 200
}
