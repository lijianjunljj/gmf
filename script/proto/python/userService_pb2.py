# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: userService.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import userModels_pb2 as userModels__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11userService.proto\x12\x08services\x1a\x10userModels.proto\"J\n\x0bUserRequest\x12\x10\n\x08UserName\x18\x01 \x01(\t\x12\x10\n\x08Password\x18\x02 \x01(\t\x12\x17\n\x0fPasswordConfirm\x18\x03 \x01(\t\"K\n\x12UserDetailResponse\x12\'\n\nUserDetail\x18\x01 \x01(\x0b\x32\x13.services.UserModel\x12\x0c\n\x04\x43ode\x18\x02 \x01(\r2\x94\x01\n\x0bUserService\x12@\n\tUserLogin\x12\x15.services.UserRequest\x1a\x1c.services.UserDetailResponse\x12\x43\n\x0cUserRegister\x12\x15.services.UserRequest\x1a\x1c.services.UserDetailResponseB\x0bZ\t./;protosb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'userService_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\t./;protos'
  _USERREQUEST._serialized_start=49
  _USERREQUEST._serialized_end=123
  _USERDETAILRESPONSE._serialized_start=125
  _USERDETAILRESPONSE._serialized_end=200
  _USERSERVICE._serialized_start=203
  _USERSERVICE._serialized_end=351
# @@protoc_insertion_point(module_scope)
