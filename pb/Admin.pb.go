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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: Admin.proto

// This file contains protocol buffers that are used for Admin service.

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRegionInfoResponse_CompactionState int32

const (
	GetRegionInfoResponse_NONE            GetRegionInfoResponse_CompactionState = 0
	GetRegionInfoResponse_MINOR           GetRegionInfoResponse_CompactionState = 1
	GetRegionInfoResponse_MAJOR           GetRegionInfoResponse_CompactionState = 2
	GetRegionInfoResponse_MAJOR_AND_MINOR GetRegionInfoResponse_CompactionState = 3
)

// Enum value maps for GetRegionInfoResponse_CompactionState.
var (
	GetRegionInfoResponse_CompactionState_name = map[int32]string{
		0: "NONE",
		1: "MINOR",
		2: "MAJOR",
		3: "MAJOR_AND_MINOR",
	}
	GetRegionInfoResponse_CompactionState_value = map[string]int32{
		"NONE":            0,
		"MINOR":           1,
		"MAJOR":           2,
		"MAJOR_AND_MINOR": 3,
	}
)

func (x GetRegionInfoResponse_CompactionState) Enum() *GetRegionInfoResponse_CompactionState {
	p := new(GetRegionInfoResponse_CompactionState)
	*p = x
	return p
}

func (x GetRegionInfoResponse_CompactionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetRegionInfoResponse_CompactionState) Descriptor() protoreflect.EnumDescriptor {
	return file_Admin_proto_enumTypes[0].Descriptor()
}

func (GetRegionInfoResponse_CompactionState) Type() protoreflect.EnumType {
	return &file_Admin_proto_enumTypes[0]
}

func (x GetRegionInfoResponse_CompactionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GetRegionInfoResponse_CompactionState) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GetRegionInfoResponse_CompactionState(num)
	return nil
}

// Deprecated: Use GetRegionInfoResponse_CompactionState.Descriptor instead.
func (GetRegionInfoResponse_CompactionState) EnumDescriptor() ([]byte, []int) {
	return file_Admin_proto_rawDescGZIP(), []int{1, 0}
}

type GetRegionInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region          *RegionSpecifier `protobuf:"bytes,1,req,name=region" json:"region,omitempty"`
	CompactionState *bool            `protobuf:"varint,2,opt,name=compaction_state,json=compactionState" json:"compaction_state,omitempty"`
}

func (x *GetRegionInfoRequest) Reset() {
	*x = GetRegionInfoRequest{}
	mi := &file_Admin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRegionInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionInfoRequest) ProtoMessage() {}

func (x *GetRegionInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionInfoRequest.ProtoReflect.Descriptor instead.
func (*GetRegionInfoRequest) Descriptor() ([]byte, []int) {
	return file_Admin_proto_rawDescGZIP(), []int{0}
}

func (x *GetRegionInfoRequest) GetRegion() *RegionSpecifier {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *GetRegionInfoRequest) GetCompactionState() bool {
	if x != nil && x.CompactionState != nil {
		return *x.CompactionState
	}
	return false
}

type GetRegionInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegionInfo      *RegionInfo                            `protobuf:"bytes,1,req,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	CompactionState *GetRegionInfoResponse_CompactionState `protobuf:"varint,2,opt,name=compaction_state,json=compactionState,enum=pb.GetRegionInfoResponse_CompactionState" json:"compaction_state,omitempty"`
	IsRecovering    *bool                                  `protobuf:"varint,3,opt,name=isRecovering" json:"isRecovering,omitempty"`
}

func (x *GetRegionInfoResponse) Reset() {
	*x = GetRegionInfoResponse{}
	mi := &file_Admin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRegionInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRegionInfoResponse) ProtoMessage() {}

func (x *GetRegionInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRegionInfoResponse.ProtoReflect.Descriptor instead.
func (*GetRegionInfoResponse) Descriptor() ([]byte, []int) {
	return file_Admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetRegionInfoResponse) GetRegionInfo() *RegionInfo {
	if x != nil {
		return x.RegionInfo
	}
	return nil
}

func (x *GetRegionInfoResponse) GetCompactionState() GetRegionInfoResponse_CompactionState {
	if x != nil && x.CompactionState != nil {
		return *x.CompactionState
	}
	return GetRegionInfoResponse_NONE
}

func (x *GetRegionInfoResponse) GetIsRecovering() bool {
	if x != nil && x.IsRecovering != nil {
		return *x.IsRecovering
	}
	return false
}

// *
// Compacts the specified region.  Performs a major compaction if specified.
// <p>
// This method is asynchronous.
type CompactRegionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region *RegionSpecifier `protobuf:"bytes,1,req,name=region" json:"region,omitempty"`
	Major  *bool            `protobuf:"varint,2,opt,name=major" json:"major,omitempty"`
	Family []byte           `protobuf:"bytes,3,opt,name=family" json:"family,omitempty"`
}

func (x *CompactRegionRequest) Reset() {
	*x = CompactRegionRequest{}
	mi := &file_Admin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompactRegionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactRegionRequest) ProtoMessage() {}

func (x *CompactRegionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactRegionRequest.ProtoReflect.Descriptor instead.
func (*CompactRegionRequest) Descriptor() ([]byte, []int) {
	return file_Admin_proto_rawDescGZIP(), []int{2}
}

func (x *CompactRegionRequest) GetRegion() *RegionSpecifier {
	if x != nil {
		return x.Region
	}
	return nil
}

func (x *CompactRegionRequest) GetMajor() bool {
	if x != nil && x.Major != nil {
		return *x.Major
	}
	return false
}

func (x *CompactRegionRequest) GetFamily() []byte {
	if x != nil {
		return x.Family
	}
	return nil
}

type CompactRegionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CompactRegionResponse) Reset() {
	*x = CompactRegionResponse{}
	mi := &file_Admin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompactRegionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompactRegionResponse) ProtoMessage() {}

func (x *CompactRegionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Admin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompactRegionResponse.ProtoReflect.Descriptor instead.
func (*CompactRegionResponse) Descriptor() ([]byte, []int) {
	return file_Admin_proto_rawDescGZIP(), []int{3}
}

var File_Admin_proto protoreflect.FileDescriptor

var file_Admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x0b, 0x48, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x8a,
	0x02, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x54, 0x0a, 0x10, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x22, 0x46, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x4d, 0x49, 0x4e, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4d,
	0x41, 0x4a, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x41, 0x4a, 0x4f, 0x52, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x4d, 0x49, 0x4e, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x71, 0x0a, 0x14, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x22, 0x17,
	0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9a, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x48, 0x0a, 0x2a, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63,
	0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x0b, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48,
	0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_Admin_proto_rawDescOnce sync.Once
	file_Admin_proto_rawDescData = file_Admin_proto_rawDesc
)

func file_Admin_proto_rawDescGZIP() []byte {
	file_Admin_proto_rawDescOnce.Do(func() {
		file_Admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_Admin_proto_rawDescData)
	})
	return file_Admin_proto_rawDescData
}

var file_Admin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Admin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Admin_proto_goTypes = []any{
	(GetRegionInfoResponse_CompactionState)(0), // 0: pb.GetRegionInfoResponse.CompactionState
	(*GetRegionInfoRequest)(nil),               // 1: pb.GetRegionInfoRequest
	(*GetRegionInfoResponse)(nil),              // 2: pb.GetRegionInfoResponse
	(*CompactRegionRequest)(nil),               // 3: pb.CompactRegionRequest
	(*CompactRegionResponse)(nil),              // 4: pb.CompactRegionResponse
	(*RegionSpecifier)(nil),                    // 5: pb.RegionSpecifier
	(*RegionInfo)(nil),                         // 6: pb.RegionInfo
}
var file_Admin_proto_depIdxs = []int32{
	5, // 0: pb.GetRegionInfoRequest.region:type_name -> pb.RegionSpecifier
	6, // 1: pb.GetRegionInfoResponse.region_info:type_name -> pb.RegionInfo
	0, // 2: pb.GetRegionInfoResponse.compaction_state:type_name -> pb.GetRegionInfoResponse.CompactionState
	5, // 3: pb.CompactRegionRequest.region:type_name -> pb.RegionSpecifier
	3, // 4: pb.AdminService.CompactRegion:input_type -> pb.CompactRegionRequest
	1, // 5: pb.AdminService.GetRegionInfo:input_type -> pb.GetRegionInfoRequest
	4, // 6: pb.AdminService.CompactRegion:output_type -> pb.CompactRegionResponse
	2, // 7: pb.AdminService.GetRegionInfo:output_type -> pb.GetRegionInfoResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_Admin_proto_init() }
func file_Admin_proto_init() {
	if File_Admin_proto != nil {
		return
	}
	file_HBase_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Admin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Admin_proto_goTypes,
		DependencyIndexes: file_Admin_proto_depIdxs,
		EnumInfos:         file_Admin_proto_enumTypes,
		MessageInfos:      file_Admin_proto_msgTypes,
	}.Build()
	File_Admin_proto = out.File
	file_Admin_proto_rawDesc = nil
	file_Admin_proto_goTypes = nil
	file_Admin_proto_depIdxs = nil
}