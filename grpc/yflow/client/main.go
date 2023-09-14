package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "yflow/client/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTicketInfoClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// create one ticket
	//r, err := c.CreateTicket(ctx, &pb.Ticket{Creator: "fanxu.felix", Note: "I wanna send messages.", Variable: "Go!"})
	//if err != nil {
	//	log.Fatalf("could not create ticket: %v", err)
	//}
	//log.Printf("ticket ID: %d added successfully", r.Id)

	// get a single ticket
	//ticket, err := c.GetTicket(ctx, &pb.TicketID{Id: r.Id})
	//if err != nil {
	//	log.Fatalf("could not get ticket: %v", err)
	//}
	//log.Printf("ticket: %v", ticket.String())

	//
	for i := 0; i < 100; i++ {
		r, err := c.CreateTicket(ctx, &pb.Ticket{Creator: "fanxu.felix", Note: "I wanna send messages.", Variable: fmt.Sprintf("var_loop_%d", i)})
		if err != nil {
			log.Fatalf("could not create ticket: %v", err)
		}
		log.Printf("ticket ID: %d create successfully", r.Id)
	}

	// get tickets in a range
	stream, err := c.GetTickets(ctx, &pb.TicketRange{Low: 0, High: 49999999})
	if err != nil {
		log.Fatalf("failed to get the stream: %s", err.Error())
	}
	for {
		ticket, err := stream.Recv()
		if err == io.EOF {
			break
		}
		log.Printf("ticket in range: %s", ticket.String())
	}
}
