# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: message.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rmessage.proto\"q\n\x0eMessageRequest\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06userId\x18\x02 \x01(\x05\x12\x13\n\x0bmessageType\x18\x03 \x01(\x05\x12\x0f\n\x07subject\x18\x04 \x01(\t\x12\x0f\n\x07message\x18\x05 \x01(\t\x12\x0c\n\x04\x66ile\x18\x06 \x01(\t\"r\n\x0fMessageResponse\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0e\n\x06userId\x18\x02 \x01(\x05\x12\x13\n\x0bmessageType\x18\x03 \x01(\x05\x12\x0f\n\x07subject\x18\x04 \x01(\t\x12\x0f\n\x07message\x18\x05 \x01(\t\x12\x0c\n\x04\x66ile\x18\x06 \x01(\t\"D\n\x13MessageListResponse\x12\r\n\x05total\x18\x01 \x01(\x05\x12\x1e\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x10.MessageResponse2s\n\x07Message\x12\x34\n\x0bMessageList\x12\x0f.MessageRequest\x1a\x14.MessageListResponse\x12\x32\n\rCreateMessage\x12\x0f.MessageRequest\x1a\x10.MessageResponseB\tZ\x07.;protob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'message_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\007.;proto'
  _MESSAGEREQUEST._serialized_start=17
  _MESSAGEREQUEST._serialized_end=130
  _MESSAGERESPONSE._serialized_start=132
  _MESSAGERESPONSE._serialized_end=246
  _MESSAGELISTRESPONSE._serialized_start=248
  _MESSAGELISTRESPONSE._serialized_end=316
  _MESSAGE._serialized_start=318
  _MESSAGE._serialized_end=433
# @@protoc_insertion_point(module_scope)
