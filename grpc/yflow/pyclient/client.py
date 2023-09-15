#!/usr/bin/env python

import grpc

import ticket_pb2
import ticket_pb2_grpc

from  google.protobuf.wrappers_pb2 import StringValue

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = ticket_pb2_grpc.TicketInfoStub(channel)
        # create 100 tickets
        for i in range(0, 100):
            response = stub.createTicket(ticket_pb2.Ticket(creator="felix", note=f"我考虑一{i}下...", variable=f"create_loop_{i}"))
            id = response.id
            print("ticket created: " + str(id))

        # get tickets by respone-streamin
        #for ticket in stub.getTickets(ticket_pb2.TicketRange(low=0, high=49999999)):
        #    print(f"one more ticket received: {ticket}")

        # update tickets by request-streaming
        #ticket1 = ticket_pb2.Ticket(id=66666666, creator="lsj", note="update one", variable="OK")
        #ticket2 = ticket_pb2.Ticket(id=88888888, creator="lsj", note="update two", variable="OK")
        #tickets = [ticket1, ticket2]

        #def ticket_generator(tickets):
        #    for ticket in tickets:
        #        yield ticket

        #result = stub.updateTicket(ticket_generator(tickets))
        #print(f"tickets updated: {result}")

        # get ticket one by one
        cmds = [
            StringValue(value="next"),
            StringValue(value="next"),
            StringValue(value="next"),
            StringValue(value="next"),
            StringValue(value="next"),
            StringValue(value="stop"),
        ]
        def cmd_generator(cmds):
            for cmd in cmds:
                yield cmd
        response =  stub.getTicketOneByOne(cmd_generator(cmds))
        for ticket in response:
            print(f"receive a ticket: {ticket}")


if __name__ == '__main__':
	run()
