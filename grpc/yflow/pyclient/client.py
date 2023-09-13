#!/usr/bin/env python

import grpc

import ticket_pb2
import ticket_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = ticket_pb2_grpc.TicketInfoStub(channel)
        response = stub.createTicket(ticket_pb2.Ticket(creator="xld", note="run了", variable="hello"))
        id = response.id
        print("ticket created: " + str(id))
        response = stub.getTicket(ticket_pb2.TicketID(id=id))
        print(f"ticket received: {response}")

if __name__ == '__main__':
	run()
