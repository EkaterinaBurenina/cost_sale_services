/*подключиться к серверу по адресу 127.0.0.1:8001 с помощью grps
вызвать что-то (get-sale), передать в него число и должно вернуться о,5*/
package main

import (
        "context"
        "log"
        "os"
        "time"

        "google.golang.org/grpc"
        pb "src/home/ekaterina/cost_sale_services"
)

const (
        address     = "127.0.0.1:8001"
)

func main() {
        // Set up a connection to the server.
        conn, err := grpc.Dial(address, grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        c := pb.NewGreeterClient(conn)

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := c.GetSale(ctx, &pb.Cost{cost: 100})
        if err != nil {
                log.Fatalf("could not response: %v", err)
        }
        log.Printf("Response: %s", r.Message)
        r, err = c.SayHelloAgain(ctx, &pb.HelloRequest{Name: name})
        if err != nil {
                log.Fatalf("could not response: %v", err)
        }
        log.Printf("Response: %s", r.Message)
}