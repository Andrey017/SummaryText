# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: protos/summary.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14protos/summary.proto\x12\x07summary\"\xb6\x01\n\x07Summary\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x10\n\x08nameFile\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x0c\n\x04note\x18\x04 \x01(\t\x12\x11\n\tarticleId\x18\x05 \x01(\x03\x12\x11\n\tcountWord\x18\x06 \x01(\x03\x12\x15\n\rcountSentense\x18\x07 \x01(\x03\x12\x18\n\x10minCountSentense\x18\x08 \x01(\x03\x12\x18\n\x10maxCountSentense\x18\t \x01(\x03\"4\n\x12SummaryTextRequest\x12\x1e\n\x04info\x18\x01 \x01(\x0b\x32\x10.summary.Summary\"!\n\x13SummaryTextResponse\x12\n\n\x02id\x18\x01 \x01(\x03\"3\n\x11SummaryLSARequest\x12\x1e\n\x04info\x18\x01 \x01(\x0b\x32\x10.summary.Summary\" \n\x12SummaryLSAResponse\x12\n\n\x02id\x18\x01 \x01(\x03\"2\n\x10SummaryT5Request\x12\x1e\n\x04info\x18\x01 \x01(\x0b\x32\x10.summary.Summary\"\x1f\n\x11SummaryT5Response\x12\n\n\x02id\x18\x01 \x01(\x03\x32\xf3\x01\n\x12SummaryTextService\x12N\n\x0fSummaryTextRank\x12\x1b.summary.SummaryTextRequest\x1a\x1c.summary.SummaryTextResponse\"\x00\x12G\n\nSummaryLSA\x12\x1a.summary.SummaryLSARequest\x1a\x1b.summary.SummaryLSAResponse\"\x00\x12\x44\n\tSummaryT5\x12\x19.summary.SummaryT5Request\x1a\x1a.summary.SummaryT5Response\"\x00\x62\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'protos.summary_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _SUMMARY._serialized_start=34
  _SUMMARY._serialized_end=216
  _SUMMARYTEXTREQUEST._serialized_start=218
  _SUMMARYTEXTREQUEST._serialized_end=270
  _SUMMARYTEXTRESPONSE._serialized_start=272
  _SUMMARYTEXTRESPONSE._serialized_end=305
  _SUMMARYLSAREQUEST._serialized_start=307
  _SUMMARYLSAREQUEST._serialized_end=358
  _SUMMARYLSARESPONSE._serialized_start=360
  _SUMMARYLSARESPONSE._serialized_end=392
  _SUMMARYT5REQUEST._serialized_start=394
  _SUMMARYT5REQUEST._serialized_end=444
  _SUMMARYT5RESPONSE._serialized_start=446
  _SUMMARYT5RESPONSE._serialized_end=477
  _SUMMARYTEXTSERVICE._serialized_start=480
  _SUMMARYTEXTSERVICE._serialized_end=723
# @@protoc_insertion_point(module_scope)
