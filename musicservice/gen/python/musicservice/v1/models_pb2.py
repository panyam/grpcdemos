# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: musicservice/v1/models.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1cmusicservice/v1/models.proto\x12\x0cmusicservice\x1a\x1fgoogle/protobuf/timestamp.proto\"\x82\x02\n\x04Song\x12\x39\n\ncreated_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x0e\n\x02id\x18\x03 \x01(\tR\x02id\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\x12\x1f\n\x0b\x63omposer_id\x18\x05 \x01(\tR\ncomposerId\x12\x1a\n\x08homepage\x18\x06 \x01(\tR\x08homepage\x12#\n\rperformer_ids\x18\x07 \x03(\tR\x0cperformerIds\"\xfc\x01\n\x06\x41rtist\x12\x39\n\ncreated_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x0e\n\x02id\x18\x03 \x01(\tR\x02id\x12\x12\n\x04name\x18\x05 \x01(\tR\x04name\x12>\n\rdate_of_birth\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0b\x64\x61teOfBirth\x12\x18\n\x07\x63ountry\x18\x07 \x01(\tR\x07\x63ountry\"\x91\x02\n\x05\x41lbum\x12\x39\n\ncreated_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x0e\n\x02id\x18\x03 \x01(\tR\x02id\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\x12\x19\n\x08song_ids\x18\x05 \x03(\tR\x07songIds\x12\x14\n\x05venue\x18\x06 \x01(\tR\x05venue\x12=\n\x0crelease_date\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0breleaseDate\"\xe8\x01\n\x05Label\x12\x39\n\ncreated_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x0e\n\x02id\x18\x03 \x01(\tR\x02id\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\x12\x45\n\x10\x65stablished_date\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0f\x65stablishedDateB\x8e\x01\n\x10\x63om.musicserviceB\x0bModelsProtoP\x01Z\x1dmusic.com/musicservice/protos\xa2\x02\x03MXX\xaa\x02\x0cMusicservice\xca\x02\x0cMusicservice\xe2\x02\x18Musicservice\\GPBMetadata\xea\x02\x0cMusicserviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'musicservice.v1.models_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\020com.musicserviceB\013ModelsProtoP\001Z\035music.com/musicservice/protos\242\002\003MXX\252\002\014Musicservice\312\002\014Musicservice\342\002\030Musicservice\\GPBMetadata\352\002\014Musicservice'
  _globals['_SONG']._serialized_start=80
  _globals['_SONG']._serialized_end=338
  _globals['_ARTIST']._serialized_start=341
  _globals['_ARTIST']._serialized_end=593
  _globals['_ALBUM']._serialized_start=596
  _globals['_ALBUM']._serialized_end=869
  _globals['_LABEL']._serialized_start=872
  _globals['_LABEL']._serialized_end=1104
# @@protoc_insertion_point(module_scope)