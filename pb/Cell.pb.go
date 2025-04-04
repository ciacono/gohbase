//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Cell and KeyValue protos

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.28.3
// source: Cell.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// The type of the key in a Cell
type CellType int32

const (
	CellType_MINIMUM       CellType = 0
	CellType_PUT           CellType = 4
	CellType_DELETE        CellType = 8
	CellType_DELETE_COLUMN CellType = 12
	CellType_DELETE_FAMILY CellType = 14
	// MAXIMUM is used when searching; you look from maximum on down.
	CellType_MAXIMUM CellType = 255
)

// Enum value maps for CellType.
var (
	CellType_name = map[int32]string{
		0:   "MINIMUM",
		4:   "PUT",
		8:   "DELETE",
		12:  "DELETE_COLUMN",
		14:  "DELETE_FAMILY",
		255: "MAXIMUM",
	}
	CellType_value = map[string]int32{
		"MINIMUM":       0,
		"PUT":           4,
		"DELETE":        8,
		"DELETE_COLUMN": 12,
		"DELETE_FAMILY": 14,
		"MAXIMUM":       255,
	}
)

func (x CellType) Enum() *CellType {
	p := new(CellType)
	*p = x
	return p
}

func (x CellType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CellType) Descriptor() protoreflect.EnumDescriptor {
	return file_Cell_proto_enumTypes[0].Descriptor()
}

func (CellType) Type() protoreflect.EnumType {
	return &file_Cell_proto_enumTypes[0]
}

func (x CellType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CellType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CellType(num)
	return nil
}

// Deprecated: Use CellType.Descriptor instead.
func (CellType) EnumDescriptor() ([]byte, []int) {
	return file_Cell_proto_rawDescGZIP(), []int{0}
}

// *
// Protocol buffer version of Cell.
type Cell struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Row           []byte                 `protobuf:"bytes,1,opt,name=row" json:"row,omitempty"`
	Family        []byte                 `protobuf:"bytes,2,opt,name=family" json:"family,omitempty"`
	Qualifier     []byte                 `protobuf:"bytes,3,opt,name=qualifier" json:"qualifier,omitempty"`
	Timestamp     *uint64                `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	CellType      *CellType              `protobuf:"varint,5,opt,name=cell_type,json=cellType,enum=pb.CellType" json:"cell_type,omitempty"`
	Value         []byte                 `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`
	Tags          []byte                 `protobuf:"bytes,7,opt,name=tags" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Cell) Reset() {
	*x = Cell{}
	mi := &file_Cell_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cell) ProtoMessage() {}

func (x *Cell) ProtoReflect() protoreflect.Message {
	mi := &file_Cell_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cell.ProtoReflect.Descriptor instead.
func (*Cell) Descriptor() ([]byte, []int) {
	return file_Cell_proto_rawDescGZIP(), []int{0}
}

func (x *Cell) GetRow() []byte {
	if x != nil {
		return x.Row
	}
	return nil
}

func (x *Cell) GetFamily() []byte {
	if x != nil {
		return x.Family
	}
	return nil
}

func (x *Cell) GetQualifier() []byte {
	if x != nil {
		return x.Qualifier
	}
	return nil
}

func (x *Cell) GetTimestamp() uint64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *Cell) GetCellType() CellType {
	if x != nil && x.CellType != nil {
		return *x.CellType
	}
	return CellType_MINIMUM
}

func (x *Cell) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Cell) GetTags() []byte {
	if x != nil {
		return x.Tags
	}
	return nil
}

// *
// Protocol buffer version of KeyValue.
// It doesn't have those transient parameters
type KeyValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Row           []byte                 `protobuf:"bytes,1,req,name=row" json:"row,omitempty"`
	Family        []byte                 `protobuf:"bytes,2,req,name=family" json:"family,omitempty"`
	Qualifier     []byte                 `protobuf:"bytes,3,req,name=qualifier" json:"qualifier,omitempty"`
	Timestamp     *uint64                `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	KeyType       *CellType              `protobuf:"varint,5,opt,name=key_type,json=keyType,enum=pb.CellType" json:"key_type,omitempty"`
	Value         []byte                 `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`
	Tags          []byte                 `protobuf:"bytes,7,opt,name=tags" json:"tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	mi := &file_Cell_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_Cell_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_Cell_proto_rawDescGZIP(), []int{1}
}

func (x *KeyValue) GetRow() []byte {
	if x != nil {
		return x.Row
	}
	return nil
}

func (x *KeyValue) GetFamily() []byte {
	if x != nil {
		return x.Family
	}
	return nil
}

func (x *KeyValue) GetQualifier() []byte {
	if x != nil {
		return x.Qualifier
	}
	return nil
}

func (x *KeyValue) GetTimestamp() uint64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *KeyValue) GetKeyType() CellType {
	if x != nil && x.KeyType != nil {
		return *x.KeyType
	}
	return CellType_MINIMUM
}

func (x *KeyValue) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *KeyValue) GetTags() []byte {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_Cell_proto protoreflect.FileDescriptor

var file_Cell_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x43, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xc1, 0x01, 0x0a, 0x04, 0x43, 0x65, 0x6c, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x29, 0x0a, 0x09, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x22, 0xc3, 0x01, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6f, 0x77, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x03,
	0x72, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x09,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x65, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x2a, 0x60, 0x0a, 0x08, 0x43, 0x65,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x49, 0x4e, 0x49, 0x4d, 0x55,
	0x4d, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4c, 0x55, 0x4d, 0x4e, 0x10, 0x0c, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x10, 0x0e, 0x12, 0x0c,
	0x0a, 0x07, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x55, 0x4d, 0x10, 0xff, 0x01, 0x42, 0x44, 0x0a, 0x2a,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f,
	0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0a, 0x43, 0x65, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0xa0,
	0x01, 0x01,
})

var (
	file_Cell_proto_rawDescOnce sync.Once
	file_Cell_proto_rawDescData []byte
)

func file_Cell_proto_rawDescGZIP() []byte {
	file_Cell_proto_rawDescOnce.Do(func() {
		file_Cell_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_Cell_proto_rawDesc), len(file_Cell_proto_rawDesc)))
	})
	return file_Cell_proto_rawDescData
}

var file_Cell_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Cell_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Cell_proto_goTypes = []any{
	(CellType)(0),    // 0: pb.CellType
	(*Cell)(nil),     // 1: pb.Cell
	(*KeyValue)(nil), // 2: pb.KeyValue
}
var file_Cell_proto_depIdxs = []int32{
	0, // 0: pb.Cell.cell_type:type_name -> pb.CellType
	0, // 1: pb.KeyValue.key_type:type_name -> pb.CellType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_Cell_proto_init() }
func file_Cell_proto_init() {
	if File_Cell_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_Cell_proto_rawDesc), len(file_Cell_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Cell_proto_goTypes,
		DependencyIndexes: file_Cell_proto_depIdxs,
		EnumInfos:         file_Cell_proto_enumTypes,
		MessageInfos:      file_Cell_proto_msgTypes,
	}.Build()
	File_Cell_proto = out.File
	file_Cell_proto_goTypes = nil
	file_Cell_proto_depIdxs = nil
}
