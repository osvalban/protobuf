// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fieldnames/fieldnames.proto

package fieldnames

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

// Assorted edge cases in field name conflict resolution.
//
// Not all (or possibly any) of these behave in an easily-understood fashion.
// This exists to demonstrate the current behavior and catch unintended
// changes in it.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Various CamelCase conversions.
	FieldOne   *string `protobuf:"bytes,1,opt,name=field_one,json=fieldOne" json:"field_one,omitempty"`
	FieldTwo   *string `protobuf:"bytes,2,opt,name=FieldTwo" json:"FieldTwo,omitempty"`
	FieldThree *string `protobuf:"bytes,3,opt,name=fieldThree" json:"fieldThree,omitempty"`
	Field_Four *string `protobuf:"bytes,4,opt,name=field__four,json=fieldFour" json:"field__four,omitempty"`
	// Field names that conflict with standard methods on the message struct.
	Descriptor_   *string `protobuf:"bytes,10,opt,name=descriptor" json:"descriptor,omitempty"`
	Marshal_      *string `protobuf:"bytes,11,opt,name=marshal" json:"marshal,omitempty"`
	Unmarshal_    *string `protobuf:"bytes,12,opt,name=unmarshal" json:"unmarshal,omitempty"`
	ProtoMessage_ *string `protobuf:"bytes,13,opt,name=proto_message,json=protoMessage" json:"proto_message,omitempty"`
	// Field names that conflict with each other after CamelCasing.
	CamelCase    *string `protobuf:"bytes,20,opt,name=CamelCase" json:"CamelCase,omitempty"`
	CamelCase_   *string `protobuf:"bytes,21,opt,name=CamelCase_,json=CamelCase" json:"CamelCase_,omitempty"`
	CamelCase__  *string `protobuf:"bytes,22,opt,name=camel_case,json=camelCase" json:"camel_case,omitempty"`   // conflicts with 20, 21
	CamelCase___ *string `protobuf:"bytes,23,opt,name=CamelCase__,json=CamelCase" json:"CamelCase__,omitempty"` // conflicts with 21, 21, renamed 22
	// Field with a getter that conflicts with another field.
	GetName *string `protobuf:"bytes,30,opt,name=get_name,json=getName" json:"get_name,omitempty"`
	Name_   *string `protobuf:"bytes,31,opt,name=name" json:"name,omitempty"`
	// Oneof that conflicts with its first field: The oneof is renamed.
	//
	// Types that are assignable to OneofConflictA_:
	//	*Message_OneofConflictA
	OneofConflictA_ isMessage_OneofConflictA_ `protobuf_oneof:"oneof_conflict_a"`
	// Oneof that conflicts with its second field: The field is renamed.
	//
	// Types that are assignable to OneofConflictB:
	//	*Message_OneofNoConflict
	//	*Message_OneofConflictB_
	OneofConflictB isMessage_OneofConflictB `protobuf_oneof:"oneof_conflict_b"`
	// Oneof with a field name that conflicts with a nested message.
	//
	// Types that are assignable to OneofConflictC:
	//	*Message_OneofMessageConflict_
	OneofConflictC isMessage_OneofConflictC `protobuf_oneof:"oneof_conflict_c"`
}

func (x *Message) Reset() {
	*x = Message{}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_fieldnames_fieldnames_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_fieldnames_fieldnames_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetFieldOne() string {
	if x != nil && x.FieldOne != nil {
		return *x.FieldOne
	}
	return ""
}

func (x *Message) GetFieldTwo() string {
	if x != nil && x.FieldTwo != nil {
		return *x.FieldTwo
	}
	return ""
}

func (x *Message) GetFieldThree() string {
	if x != nil && x.FieldThree != nil {
		return *x.FieldThree
	}
	return ""
}

func (x *Message) GetField_Four() string {
	if x != nil && x.Field_Four != nil {
		return *x.Field_Four
	}
	return ""
}

func (x *Message) GetDescriptor_() string {
	if x != nil && x.Descriptor_ != nil {
		return *x.Descriptor_
	}
	return ""
}

func (x *Message) GetMarshal_() string {
	if x != nil && x.Marshal_ != nil {
		return *x.Marshal_
	}
	return ""
}

func (x *Message) GetUnmarshal_() string {
	if x != nil && x.Unmarshal_ != nil {
		return *x.Unmarshal_
	}
	return ""
}

func (x *Message) GetProtoMessage_() string {
	if x != nil && x.ProtoMessage_ != nil {
		return *x.ProtoMessage_
	}
	return ""
}

func (x *Message) GetCamelCase() string {
	if x != nil && x.CamelCase != nil {
		return *x.CamelCase
	}
	return ""
}

func (x *Message) GetCamelCase_() string {
	if x != nil && x.CamelCase_ != nil {
		return *x.CamelCase_
	}
	return ""
}

func (x *Message) GetCamelCase__() string {
	if x != nil && x.CamelCase__ != nil {
		return *x.CamelCase__
	}
	return ""
}

func (x *Message) GetCamelCase___() string {
	if x != nil && x.CamelCase___ != nil {
		return *x.CamelCase___
	}
	return ""
}

func (x *Message) GetGetName() string {
	if x != nil && x.GetName != nil {
		return *x.GetName
	}
	return ""
}

func (x *Message) GetName_() string {
	if x != nil && x.Name_ != nil {
		return *x.Name_
	}
	return ""
}

func (m *Message) GetOneofConflictA_() isMessage_OneofConflictA_ {
	if m != nil {
		return m.OneofConflictA_
	}
	return nil
}

func (x *Message) GetOneofConflictA() string {
	if x, ok := x.GetOneofConflictA_().(*Message_OneofConflictA); ok {
		return x.OneofConflictA
	}
	return ""
}

func (m *Message) GetOneofConflictB() isMessage_OneofConflictB {
	if m != nil {
		return m.OneofConflictB
	}
	return nil
}

func (x *Message) GetOneofNoConflict() string {
	if x, ok := x.GetOneofConflictB().(*Message_OneofNoConflict); ok {
		return x.OneofNoConflict
	}
	return ""
}

func (x *Message) GetOneofConflictB_() string {
	if x, ok := x.GetOneofConflictB().(*Message_OneofConflictB_); ok {
		return x.OneofConflictB_
	}
	return ""
}

func (m *Message) GetOneofConflictC() isMessage_OneofConflictC {
	if m != nil {
		return m.OneofConflictC
	}
	return nil
}

func (x *Message) GetOneofMessageConflict() string {
	if x, ok := x.GetOneofConflictC().(*Message_OneofMessageConflict_); ok {
		return x.OneofMessageConflict
	}
	return ""
}

type isMessage_OneofConflictA_ interface {
	isMessage_OneofConflictA_()
}

type Message_OneofConflictA struct {
	OneofConflictA string `protobuf:"bytes,40,opt,name=OneofConflictA,oneof"`
}

func (*Message_OneofConflictA) isMessage_OneofConflictA_() {}

type isMessage_OneofConflictB interface {
	isMessage_OneofConflictB()
}

type Message_OneofNoConflict struct {
	OneofNoConflict string `protobuf:"bytes,50,opt,name=oneof_no_conflict,json=oneofNoConflict,oneof"`
}

type Message_OneofConflictB_ struct {
	OneofConflictB_ string `protobuf:"bytes,51,opt,name=OneofConflictB,oneof"`
}

func (*Message_OneofNoConflict) isMessage_OneofConflictB() {}

func (*Message_OneofConflictB_) isMessage_OneofConflictB() {}

type isMessage_OneofConflictC interface {
	isMessage_OneofConflictC()
}

type Message_OneofMessageConflict_ struct {
	OneofMessageConflict string `protobuf:"bytes,60,opt,name=oneof_message_conflict,json=oneofMessageConflict,oneof"`
}

func (*Message_OneofMessageConflict_) isMessage_OneofConflictC() {}

type Message_OneofMessageConflict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message_OneofMessageConflict) Reset() {
	*x = Message_OneofMessageConflict{}
}

func (x *Message_OneofMessageConflict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_OneofMessageConflict) ProtoMessage() {}

func (x *Message_OneofMessageConflict) ProtoReflect() protoreflect.Message {
	mi := &file_fieldnames_fieldnames_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_OneofMessageConflict.ProtoReflect.Descriptor instead.
func (*Message_OneofMessageConflict) Descriptor() ([]byte, []int) {
	return file_fieldnames_fieldnames_proto_rawDescGZIP(), []int{0, 0}
}

var File_fieldnames_fieldnames_proto protoreflect.FileDescriptor

var file_fieldnames_fieldnames_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xb8, 0x05, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x6e,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x6e,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x77, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x77, 0x6f, 0x12, 0x1e, 0x0a,
	0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x68, 0x72, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x54, 0x68, 0x72, 0x65, 0x65, 0x12, 0x1e, 0x0a,
	0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x5f, 0x66, 0x6f, 0x75, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x46, 0x6f, 0x75, 0x72, 0x12, 0x1e, 0x0a,
	0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x6e, 0x6d, 0x61, 0x72,
	0x73, 0x68, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x6e, 0x6d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x61,
	0x6d, 0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43,
	0x61, 0x6d, 0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x43, 0x61, 0x6d, 0x65,
	0x6c, 0x43, 0x61, 0x73, 0x65, 0x5f, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x61,
	0x6d, 0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x6d, 0x65, 0x6c,
	0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x6d,
	0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x43, 0x61, 0x6d, 0x65, 0x6c, 0x43,
	0x61, 0x73, 0x65, 0x5f, 0x5f, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x61, 0x6d,
	0x65, 0x6c, 0x43, 0x61, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f,
	0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x41, 0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x41, 0x12,
	0x2c, 0x0a, 0x11, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6e, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x6c, 0x69, 0x63, 0x74, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0f, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x4e, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x12, 0x28, 0x0a,
	0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x42, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x43, 0x6f,
	0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x42, 0x12, 0x36, 0x0a, 0x16, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63,
	0x74, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x14, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x1a,
	0x16, 0x0a, 0x14, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x42, 0x12, 0x0a, 0x10, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x5f, 0x61, 0x42, 0x12, 0x0a, 0x10, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x5f, 0x62, 0x42,
	0x12, 0x0a, 0x10, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63,
	0x74, 0x5f, 0x63, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x73,
}

var (
	file_fieldnames_fieldnames_proto_rawDescOnce sync.Once
	file_fieldnames_fieldnames_proto_rawDescData = file_fieldnames_fieldnames_proto_rawDesc
)

func file_fieldnames_fieldnames_proto_rawDescGZIP() []byte {
	file_fieldnames_fieldnames_proto_rawDescOnce.Do(func() {
		file_fieldnames_fieldnames_proto_rawDescData = protoimpl.X.CompressGZIP(file_fieldnames_fieldnames_proto_rawDescData)
	})
	return file_fieldnames_fieldnames_proto_rawDescData
}

var file_fieldnames_fieldnames_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_fieldnames_fieldnames_proto_goTypes = []interface{}{
	(*Message)(nil),                      // 0: goproto.protoc.fieldnames.Message
	(*Message_OneofMessageConflict)(nil), // 1: goproto.protoc.fieldnames.Message.OneofMessageConflict
}
var file_fieldnames_fieldnames_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_fieldnames_fieldnames_proto_init() }
func file_fieldnames_fieldnames_proto_init() {
	if File_fieldnames_fieldnames_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fieldnames_fieldnames_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fieldnames_fieldnames_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_OneofMessageConflict); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_fieldnames_fieldnames_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Message_OneofConflictA)(nil),
		(*Message_OneofNoConflict)(nil),
		(*Message_OneofConflictB_)(nil),
		(*Message_OneofMessageConflict_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fieldnames_fieldnames_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fieldnames_fieldnames_proto_goTypes,
		DependencyIndexes: file_fieldnames_fieldnames_proto_depIdxs,
		MessageInfos:      file_fieldnames_fieldnames_proto_msgTypes,
	}.Build()
	File_fieldnames_fieldnames_proto = out.File
	file_fieldnames_fieldnames_proto_rawDesc = nil
	file_fieldnames_fieldnames_proto_goTypes = nil
	file_fieldnames_fieldnames_proto_depIdxs = nil
}
