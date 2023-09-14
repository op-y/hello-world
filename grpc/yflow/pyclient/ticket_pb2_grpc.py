# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import ticket_pb2 as ticket__pb2


class TicketInfoStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.createTicket = channel.unary_unary(
                '/proto.TicketInfo/createTicket',
                request_serializer=ticket__pb2.Ticket.SerializeToString,
                response_deserializer=ticket__pb2.TicketID.FromString,
                )
        self.getTicket = channel.unary_unary(
                '/proto.TicketInfo/getTicket',
                request_serializer=ticket__pb2.TicketID.SerializeToString,
                response_deserializer=ticket__pb2.Ticket.FromString,
                )
        self.getTickets = channel.unary_stream(
                '/proto.TicketInfo/getTickets',
                request_serializer=ticket__pb2.TicketRange.SerializeToString,
                response_deserializer=ticket__pb2.Ticket.FromString,
                )


class TicketInfoServicer(object):
    """Missing associated documentation comment in .proto file."""

    def createTicket(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getTicket(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def getTickets(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TicketInfoServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'createTicket': grpc.unary_unary_rpc_method_handler(
                    servicer.createTicket,
                    request_deserializer=ticket__pb2.Ticket.FromString,
                    response_serializer=ticket__pb2.TicketID.SerializeToString,
            ),
            'getTicket': grpc.unary_unary_rpc_method_handler(
                    servicer.getTicket,
                    request_deserializer=ticket__pb2.TicketID.FromString,
                    response_serializer=ticket__pb2.Ticket.SerializeToString,
            ),
            'getTickets': grpc.unary_stream_rpc_method_handler(
                    servicer.getTickets,
                    request_deserializer=ticket__pb2.TicketRange.FromString,
                    response_serializer=ticket__pb2.Ticket.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.TicketInfo', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TicketInfo(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def createTicket(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.TicketInfo/createTicket',
            ticket__pb2.Ticket.SerializeToString,
            ticket__pb2.TicketID.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getTicket(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.TicketInfo/getTicket',
            ticket__pb2.TicketID.SerializeToString,
            ticket__pb2.Ticket.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def getTickets(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/proto.TicketInfo/getTickets',
            ticket__pb2.TicketRange.SerializeToString,
            ticket__pb2.Ticket.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
