# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from musicservice.v1 import artists_pb2 as musicservice_dot_v1_dot_artists__pb2


class ArtistServiceStub(object):
    """
    Artist service for creating, listing, updating and deleting artists.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateArtist = channel.unary_unary(
                '/musicservice.ArtistService/CreateArtist',
                request_serializer=musicservice_dot_v1_dot_artists__pb2.CreateArtistRequest.SerializeToString,
                response_deserializer=musicservice_dot_v1_dot_artists__pb2.CreateArtistResponse.FromString,
                )
        self.GetArtists = channel.unary_unary(
                '/musicservice.ArtistService/GetArtists',
                request_serializer=musicservice_dot_v1_dot_artists__pb2.GetArtistsRequest.SerializeToString,
                response_deserializer=musicservice_dot_v1_dot_artists__pb2.GetArtistsResponse.FromString,
                )
        self.UpdateArtist = channel.unary_unary(
                '/musicservice.ArtistService/UpdateArtist',
                request_serializer=musicservice_dot_v1_dot_artists__pb2.UpdateArtistRequest.SerializeToString,
                response_deserializer=musicservice_dot_v1_dot_artists__pb2.UpdateArtistResponse.FromString,
                )
        self.DeleteArtist = channel.unary_unary(
                '/musicservice.ArtistService/DeleteArtist',
                request_serializer=musicservice_dot_v1_dot_artists__pb2.DeleteArtistRequest.SerializeToString,
                response_deserializer=musicservice_dot_v1_dot_artists__pb2.DeleteArtistResponse.FromString,
                )
        self.ListArtists = channel.unary_unary(
                '/musicservice.ArtistService/ListArtists',
                request_serializer=musicservice_dot_v1_dot_artists__pb2.ListArtistsRequest.SerializeToString,
                response_deserializer=musicservice_dot_v1_dot_artists__pb2.ListArtistsResponse.FromString,
                )


class ArtistServiceServicer(object):
    """
    Artist service for creating, listing, updating and deleting artists.
    """

    def CreateArtist(self, request, context):
        """
        Create a new Artist
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetArtists(self, request, context):
        """
        Batch gets multiple artists.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateArtist(self, request, context):
        """
        Updates specific fields of an Artist
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteArtist(self, request, context):
        """
        Deletes an artist from our system.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListArtists(self, request, context):
        """
        Finds and retrieves artists matching the particular criteria.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ArtistServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateArtist': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateArtist,
                    request_deserializer=musicservice_dot_v1_dot_artists__pb2.CreateArtistRequest.FromString,
                    response_serializer=musicservice_dot_v1_dot_artists__pb2.CreateArtistResponse.SerializeToString,
            ),
            'GetArtists': grpc.unary_unary_rpc_method_handler(
                    servicer.GetArtists,
                    request_deserializer=musicservice_dot_v1_dot_artists__pb2.GetArtistsRequest.FromString,
                    response_serializer=musicservice_dot_v1_dot_artists__pb2.GetArtistsResponse.SerializeToString,
            ),
            'UpdateArtist': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateArtist,
                    request_deserializer=musicservice_dot_v1_dot_artists__pb2.UpdateArtistRequest.FromString,
                    response_serializer=musicservice_dot_v1_dot_artists__pb2.UpdateArtistResponse.SerializeToString,
            ),
            'DeleteArtist': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteArtist,
                    request_deserializer=musicservice_dot_v1_dot_artists__pb2.DeleteArtistRequest.FromString,
                    response_serializer=musicservice_dot_v1_dot_artists__pb2.DeleteArtistResponse.SerializeToString,
            ),
            'ListArtists': grpc.unary_unary_rpc_method_handler(
                    servicer.ListArtists,
                    request_deserializer=musicservice_dot_v1_dot_artists__pb2.ListArtistsRequest.FromString,
                    response_serializer=musicservice_dot_v1_dot_artists__pb2.ListArtistsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'musicservice.ArtistService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ArtistService(object):
    """
    Artist service for creating, listing, updating and deleting artists.
    """

    @staticmethod
    def CreateArtist(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/musicservice.ArtistService/CreateArtist',
            musicservice_dot_v1_dot_artists__pb2.CreateArtistRequest.SerializeToString,
            musicservice_dot_v1_dot_artists__pb2.CreateArtistResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetArtists(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/musicservice.ArtistService/GetArtists',
            musicservice_dot_v1_dot_artists__pb2.GetArtistsRequest.SerializeToString,
            musicservice_dot_v1_dot_artists__pb2.GetArtistsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateArtist(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/musicservice.ArtistService/UpdateArtist',
            musicservice_dot_v1_dot_artists__pb2.UpdateArtistRequest.SerializeToString,
            musicservice_dot_v1_dot_artists__pb2.UpdateArtistResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteArtist(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/musicservice.ArtistService/DeleteArtist',
            musicservice_dot_v1_dot_artists__pb2.DeleteArtistRequest.SerializeToString,
            musicservice_dot_v1_dot_artists__pb2.DeleteArtistResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ListArtists(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/musicservice.ArtistService/ListArtists',
            musicservice_dot_v1_dot_artists__pb2.ListArtistsRequest.SerializeToString,
            musicservice_dot_v1_dot_artists__pb2.ListArtistsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)