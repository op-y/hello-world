package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
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
	return nil, status.Errorf(codes.NotFound, "ticket does not exist.", in.Id)
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
