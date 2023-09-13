package main

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateTicket(ctx, &pb.Ticket{Creator: "fanxu.felix", Note: "I wanna send messages.", Variable: "Go!"})
	if err != nil {
		log.Fatalf("could not create ticket: %v", err)
	}
	log.Printf("ticket ID: %d added successfully", r.Id)

	ticket, err := c.GetTicket(ctx, &pb.TicketID{Id: r.Id})
	if err != nil {
		log.Fatalf("could not get ticket: %v", err)
	}
	log.Printf("ticket: %v", ticket.String())
}
