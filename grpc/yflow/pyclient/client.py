#!/usr/bin/env python

import grpc

import ticket_pb2
import ticket_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = ticket_pb2_grpc.TicketInfoStub(channel)
        # create 100 tickets
        for i in range(0, 100):
            response = stub.createTicket(ticket_pb2.Ticket(creator="felix", note=f"我考虑一{i}下...", variable=f"create_loop_{i}"))
            id = response.id
            print("ticket created: " + str(id))

        for ticket in stub.getTickets(ticket_pb2.TicketRange(low=0, high=49999999)):
            print(f"one more ticket received: {ticket}")

if __name__ == '__main__':
	run()
