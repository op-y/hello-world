package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "yflow/service/proto"
)

const (
	port = ":50051"
)

type server struct {
	ticketMap map[int64]*pb.Ticket
}

func (s *server) CreateTicket(ctx context.Context,
	in *pb.Ticket) (*pb.TicketID, error) {

	// mock deley
	//time.Sleep(time.Second * 5)

	if ctx.Err() == context.DeadlineExceeded {
		log.Println("client deadline exceeded...")
		return nil, ctx.Err()
	}

	if in.Creator == "expert" {
		errStatus := status.New(codes.InvalidArgument, "creator expert not allowed")
		ds, err := errStatus.WithDetails(
			&errdetails.BadRequest_FieldViolation{
				Field:       "Creator",
				Description: "我们不接受专家创建工单！",
			},
		)
		if err != nil {
			return nil, errStatus.Err()
		}
		return nil, ds.Err()
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		for k, v := range md {
			log.Printf("metadata %s:%v", k, v)
		}
	}

	in.Id = rand.Int63n(99999999)
	if s.ticketMap == nil {
		s.ticketMap = make(map[int64]*pb.Ticket)
	}
	s.ticketMap[in.Id] = in
	log.Printf("ticket %d : %s - created by %s", in.Id, in.Note, in.Creator)
	return &pb.TicketID{Id: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetTicket(ctx context.Context, in *pb.TicketID) (*pb.Ticket, error) {
	ticket, exists := s.ticketMap[in.Id]
	if exists && ticket != nil {
		log.Printf("ticket %d : creator %s - retrieved.", ticket.Id, ticket.Creator)
		return ticket, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "ticket %d does not exist.", in.Id)
}

func (s *server) GetTickets(tr *pb.TicketRange, stream pb.TicketInfo_GetTicketsServer) error {
	for id, ticket := range s.ticketMap {
		if id > tr.Low && id < tr.High {
			if err := stream.Send(ticket); err != nil {
				log.Printf("ticket %d error sending message to stream: %s", ticket.Id, err.Error())
			}
			log.Printf("found ticket %d", ticket.Id)
		}
	}
	return nil
}

func (s *server) UpdateTicket(stream pb.TicketInfo_UpdateTicketServer) error {
	ticket_ids := []int64{}
	for {
		ticket, err := stream.Recv()
		if err == io.EOF {
			result := "tickets updated: "
			for _, id := range ticket_ids {
				result = fmt.Sprintf("%s %d", result, id)
			}
			return stream.SendAndClose(
				&wrappers.StringValue{Value: result})
		}
		s.ticketMap[ticket.Id] = ticket
		ticket_ids = append(ticket_ids, ticket.Id)
		log.Printf("ticket %d updated", ticket.Id)
	}
}

func (s *server) GetTicketOneByOne(stream pb.TicketInfo_GetTicketOneByOneServer) error {
	idx := 0
	indice := []int64{}
	for key := range s.ticketMap {
		indice = append(indice, key)
	}

	for {
		cmd, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if cmd.Value != "next" {
			return nil
		}
		if idx < len(indice) {
			stream.Send(s.ticketMap[indice[idx]])
			idx++
		} else {
			return nil
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTicketInfoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
