package main

import (
	"context"
	//"fmt"
	"io"
	"log"
	"time"

	//"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

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

	// ddl
	//ddl := time.Now().Add(time.Duration(time.Second * 2))
	//ctx, cancel := context.WithDeadline(context.Background(), ddl)
	//defer cancel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// metadata
	md := metadata.Pairs(
		"biz", "edr",
		"code", "abc123",
	)
	ctxWithMetaData := metadata.NewOutgoingContext(ctx, md)

	// create one ticket
	r, err := c.CreateTicket(
		ctxWithMetaData,
		&pb.Ticket{Creator: "felix", Note: "I wanna send messages.", Variable: "Go!"},
		grpc.UseCompressor(gzip.Name),
	)
	if err != nil {
		got := status.Code(err)
		if got == codes.InvalidArgument {
			log.Printf("Invalid Argument Error: %s", got.String())
			errStatus := status.Convert(err)
			for _, d := range errStatus.Details() {
				switch info := d.(type) {
				case *errdetails.BadRequest_FieldViolation:
					log.Printf("Request Field Invalid: %s", info)
				default:
					log.Printf("Unexpected error type: %s", info)
				}
			}
		} else {
			log.Fatalf("error occured when create ticket: %v", got)
		}
	} else {
		log.Printf("ticket ID: %d added successfully", r.Id)
	}

	// get a single ticket
	//ticket, err := c.GetTicket(ctx, &pb.TicketID{Id: r.Id})
	//if err != nil {
	//	log.Fatalf("could not get ticket: %v", err)
	//}
	//log.Printf("ticket: %v", ticket.String())

	// create 100 tickets
	//for i := 0; i < 100; i++ {
	//	r, err := c.CreateTicket(ctx, &pb.Ticket{Creator: "fanxu.felix", Note: "I wanna send messages.", Variable: fmt.Sprintf("var_loop_%d", i)})
	//	if err != nil {
	//		log.Fatalf("could not create ticket: %v", err)
	//	}
	//	log.Printf("ticket ID: %d create successfully", r.Id)
	//}

	// get tickets in a range
	//stream, err := c.GetTickets(ctx, &pb.TicketRange{Low: 0, High: 49999999})
	//if err != nil {
	//	log.Fatalf("failed to get the stream: %s", err.Error())
	//}
	//for {
	//	ticket, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	log.Printf("ticket in range: %s", ticket.String())
	//}

	// update ticket, I don't care weither the ticket exists or not
	//stream, err := c.UpdateTicket(ctx)
	//if err != nil {
	//	log.Fatalf("%v.UpdateTicket(_) = _, %s", c, err.Error())
	//}
	//ticket := &pb.Ticket{Id: 66666666, Creator: "lsj", Note: "updated one", Variable: "OK"}
	//if err := stream.Send(ticket); err != nil {
	//	log.Fatalf("%v.Send(%v) = %v", stream, ticket, err.Error())
	//}
	//ticket = &pb.Ticket{Id: 88888888, Creator: "lsj", Note: "updated two", Variable: "OK"}
	//if err := stream.Send(ticket); err != nil {
	//	log.Fatalf("%v.Send(%v) = %v", stream, ticket, err.Error())
	//}
	//result, err := stream.CloseAndRecv()
	//if err != nil {
	//	log.Fatalf("%v.CloseAndRecv() got error %s, want %v", stream, err.Error(), nil)
	//}
	//log.Printf("ticket updated: %s", result)

	// get ticket one by one
	//stream, err := c.GetTicketOneByOne(ctx)
	//if err != nil {
	//	log.Fatalf("%v.GetTicketOneByOne(_) = _, %s", c, err.Error())
	//}
	//for i := 0; i < 10; i++ {
	//	if err := stream.Send(&wrappers.StringValue{Value: "next"}); err != nil {
	//		log.Fatalf("%v.Send(next) = %s", c, err.Error())
	//	}
	//}
	//channel := make(chan struct{})
	//go asyncPrintTicket(stream, channel)
	//time.Sleep(time.Millisecond * 100)

	//cancel()
	//log.Printf("RPC status: %s", ctx.Err().Error())

	//if err := stream.Send(&wrappers.StringValue{Value: "next"}); err != nil {
	//	log.Fatalf("%v.Send(next) = %s", c, err.Error())
	//}
	//if err := stream.CloseSend(); err != nil {
	//	log.Fatalf("%v.CloseSend() = %s", c, err.Error())
	//}
	//channel <- struct{}{}
}

func asyncPrintTicket(stream pb.TicketInfo_GetTicketOneByOneClient, c chan struct{}) {
	for {
		ticket, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if stream.Context().Err() == context.Canceled {
			log.Printf("request cancelled when receive message %s", err.Error())
			break
		}
		if err != nil {
			log.Printf("error occured when receive message %s", err.Error())
		}
		log.Printf("receive a ticket: %s", ticket.String())
	}
	<-c
	log.Println("quit async print")
}
