# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greeter.proto
# Protobuf Python Version: 4.25.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rgreeter.proto\x12\x07greeter\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"!\n\rHelloResponse\x12\x10\n\x08greeting\x18\x01 \x01(\t2D\n\x07Greeter\x12\x39\n\x08SayHello\x12\x15.greeter.HelloRequest\x1a\x16.greeter.HelloResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'greeter_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_HELLOREQUEST']._serialized_start=26
  _globals['_HELLOREQUEST']._serialized_end=54
  _globals['_HELLORESPONSE']._serialized_start=56
  _globals['_HELLORESPONSE']._serialized_end=89
  _globals['_GREETER']._serialized_start=91
  _globals['_GREETER']._serialized_end=159
# @@protoc_insertion_point(module_scope)