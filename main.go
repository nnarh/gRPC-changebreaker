// When I type in "go run main.go --total=5.32 --paid=20.00"
// And I press enter
// Then the program prints out:
//   You are owed:
//   $10.00 x 1
//   $1.00 x 4
//   $0.25 x 2
//   $0.10 x 1
//   $0.05 x 1
//   $0.01 x 3

package main

import (
    "os"

    "github.com/urfave/cli"
    cb "github.com/learning/gRPC-changebreaker/proto"
    "google.golang.org/grpc"
    "github.com/pkg/errors"
    "net"
    "github.com/learning/gRPC-changebreaker/pkg"
    "github.com/Sirupsen/logrus"
)

var amountGiven = 30
var subtotal = 30

// var amount = big.NewRat(30).SetFloat64(.2)
// var subtotal = big.(22.45).SetFloat64(1.2)

// var denom = map[string]int{
// 	"20.00": 2000,
// 	"10.00": 1000,.
// 	"5.00":  500,
// 	"1.00":  100,
// 	"0.25":  25,
// 	"0.10":  10,
// 	"0.05":  5,
// 	"0.01":  1,
// }
//
// var keys []string

func main() {
    app := cli.NewApp()
    app.Name = "changebreaker"
    app.Usage = "To break change!"
    app.Commands = []cli.Command{cliValues()}

    //sort keys: https://blog.golang.org/go-maps-in-action

    // for k, _ := range denom {
    // 	keys = append(keys, k)
    //
    // }
    // sort.Strings(keys)
    // fmt.Println(keys)
    // for _, k := range keys {
    // 	fmt.Println("Key:", k, "Value:", denom[k])
    // }

    // fmt.Println(giveChange(amountGiven, subtotal))
    // fmt.Println(changeBreakdown(giveChange(amountGiven, subtotal)))

    if err := app.Run(os.Args); err != nil {
        logrus.WithError(err).Fatal("could not run application")
    }

}

func cliValues() cli.Command {
    return cli.Command{
        Name:   "server",
        Usage:  "starts a grpc server",
        Action: configureServer,
        Flags: []cli.Flag{
            cli.Float64Flag{
                Name:  "total",
                Usage: "Total amount due",
                Value:  32.50,
            },
            cli.Float64Flag{
                Name:  "paid",
                Usage: "Amount Paid",
                Value:  50.00,
            },
            cli.StringFlag{
                Name:   "addr",
                EnvVar: "ADDR",
                Value:  "0.0.0.0:50051",
                Usage:  "address for the grpc server to listen on",
            },
        },

        //app.Action = func(c *cli.Context) error {
        //    total := int(c.Float64("total") * 100)
        //    paid := int(c.Float64("paid") * 100)
        //
        //    co := giveChange(paid, total)
        //
        //    change := float64(giveChange(paid, total))
        //    fmt.Printf("You are owed: $%.2f\n", change/100)
        //    fmt.Println("Here is your change breakdown:")
        //    changeBreakdown(co)
        //
        //    return nil
        //}
        //
        //app.Run(os.Args)
    }
}

func configureServer(c *cli.Context) error {
    addr := c.String("addr")
    gs := grpc.NewServer()
    s := server.NewServer()
    cb.RegisterChangebreakerServer(gs, s)

    lis, err := net.Listen("tcp", addr)
    if err != nil {
        return errors.Wrap(err, "could not setup tcp listener")
    }

    logrus.WithField("addr", addr).Info("starting server")

    return gs.Serve(lis)
}