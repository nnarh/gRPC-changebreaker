package server

import (
    cb "github.com/learning/gRPC-changebreaker/proto"
    "golang.org/x/net/context"

    //"fmt"
    "sort"
    "fmt"
)

type Server struct {}

func NewServer() *Server {
    return &Server{}
}

var _ cb.ChangebreakerServer = &Server{}



func (s *Server) Change(c context.Context, money *cb.ChangeReq) (*cb.ChangeResp, error) {
    resp := &cb.ChangeResp{
         Change: giveChange(money.Paid, money.Total),
     }
     return resp, nil
}

func giveChange(paid, total float32) float32 {

    switch {
    case paid-total < 0:
        fmt.Println("Give More Money Please!")
    case paid-total == 0:
        fmt.Println("Thanks for the exact change!")
    }
    return paid - total
}

func changeBreakdown(giveChange int) {
    var m = make(map[float64]int)
    m[20.00] = giveChange / 2000
    giveChange %= 2000
    m[10.00] = giveChange / 1000
    giveChange %= 1000
    m[5.00] = giveChange / 500
    giveChange %= 500
    m[1.00] = giveChange / 100
    giveChange %= 100
    m[0.25] = giveChange / 25
    giveChange %= 25
    m[0.10] = giveChange / 10
    giveChange %= 10
    m[0.05] = giveChange / 5
    giveChange %= 5
    m[0.01] = giveChange / 1
    giveChange %= 1

    var keys []float64
    for k := range m {
        keys = append(keys, k)
    }
    sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

    for _, k := range keys {
        if m[k] != 0 {
            fmt.Printf("$%.2f x %d\n", k, m[k])
            // fmt.Println("$Key:", k, "Value:", m[k])
        }
    }
}

