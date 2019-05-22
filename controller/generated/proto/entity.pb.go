//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EntityConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tomstoned            bool     `protobuf:"varint,2,opt,name=tomstoned,proto3" json:"tomstoned,omitempty"`
	Config               []byte   `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityConfig) Reset()         { *m = EntityConfig{} }
func (m *EntityConfig) String() string { return proto.CompactTextString(m) }
func (*EntityConfig) ProtoMessage()    {}
func (*EntityConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{0}
}

func (m *EntityConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityConfig.Unmarshal(m, b)
}
func (m *EntityConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityConfig.Marshal(b, m, deterministic)
}
func (m *EntityConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityConfig.Merge(m, src)
}
func (m *EntityConfig) XXX_Size() int {
	return xxx_messageInfo_EntityConfig.Size(m)
}
func (m *EntityConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityConfig.DiscardUnknown(m)
}

var xxx_messageInfo_EntityConfig proto.InternalMessageInfo

func (m *EntityConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntityConfig) GetTomstoned() bool {
	if m != nil {
		return m.Tomstoned
	}
	return false
}

func (m *EntityConfig) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

type EntityName struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tomstoned            bool     `protobuf:"varint,2,opt,name=tomstoned,proto3" json:"tomstoned,omitempty"`
	Incarnation          int32    `protobuf:"varint,3,opt,name=incarnation,proto3" json:"incarnation,omitempty"`
	LastUpdatedAt        int64    `protobuf:"varint,4,opt,name=last_updated_at,json=lastUpdatedAt,proto3" json:"last_updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityName) Reset()         { *m = EntityName{} }
func (m *EntityName) String() string { return proto.CompactTextString(m) }
func (*EntityName) ProtoMessage()    {}
func (*EntityName) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{1}
}

func (m *EntityName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityName.Unmarshal(m, b)
}
func (m *EntityName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityName.Marshal(b, m, deterministic)
}
func (m *EntityName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityName.Merge(m, src)
}
func (m *EntityName) XXX_Size() int {
	return xxx_messageInfo_EntityName.Size(m)
}
func (m *EntityName) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityName.DiscardUnknown(m)
}

var xxx_messageInfo_EntityName proto.InternalMessageInfo

func (m *EntityName) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntityName) GetTomstoned() bool {
	if m != nil {
		return m.Tomstoned
	}
	return false
}

func (m *EntityName) GetIncarnation() int32 {
	if m != nil {
		return m.Incarnation
	}
	return 0
}

func (m *EntityName) GetLastUpdatedAt() int64 {
	if m != nil {
		return m.LastUpdatedAt
	}
	return 0
}

type EntityList struct {
	LastUpdatedAt        int64         `protobuf:"varint,1,opt,name=last_updated_at,json=lastUpdatedAt,proto3" json:"last_updated_at,omitempty"`
	Entities             []*EntityName `protobuf:"bytes,2,rep,name=entities,proto3" json:"entities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EntityList) Reset()         { *m = EntityList{} }
func (m *EntityList) String() string { return proto.CompactTextString(m) }
func (*EntityList) ProtoMessage()    {}
func (*EntityList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{2}
}

func (m *EntityList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityList.Unmarshal(m, b)
}
func (m *EntityList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityList.Marshal(b, m, deterministic)
}
func (m *EntityList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityList.Merge(m, src)
}
func (m *EntityList) XXX_Size() int {
	return xxx_messageInfo_EntityList.Size(m)
}
func (m *EntityList) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityList.DiscardUnknown(m)
}

var xxx_messageInfo_EntityList proto.InternalMessageInfo

func (m *EntityList) GetLastUpdatedAt() int64 {
	if m != nil {
		return m.LastUpdatedAt
	}
	return 0
}

func (m *EntityList) GetEntities() []*EntityName {
	if m != nil {
		return m.Entities
	}
	return nil
}

func init() {
	proto.RegisterType((*EntityConfig)(nil), "models.EntityConfig")
	proto.RegisterType((*EntityName)(nil), "models.EntityName")
	proto.RegisterType((*EntityList)(nil), "models.EntityList")
}

func init() { proto.RegisterFile("entity.proto", fileDescriptor_cf50d946d740d100) }

var fileDescriptor_cf50d946d740d100 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0x86, 0xc9, 0xed, 0xb5, 0xdc, 0x7b, 0x5a, 0x11, 0xce, 0x20, 0x19, 0x1c, 0x42, 0x07, 0xc9,
	0x94, 0x41, 0x9f, 0x40, 0xc4, 0x4d, 0x1c, 0x02, 0x82, 0x5b, 0x89, 0x4d, 0x94, 0x40, 0x9b, 0x94,
	0xe6, 0x38, 0xf8, 0x06, 0x3e, 0xb6, 0x34, 0x55, 0xeb, 0xd0, 0xe5, 0x6e, 0xc9, 0xc7, 0xf9, 0x3f,
	0xce, 0xf9, 0xa1, 0x76, 0x81, 0x3c, 0x7d, 0xaa, 0x71, 0x8a, 0x14, 0xb1, 0x1c, 0xa2, 0x75, 0x7d,
	0x6a, 0x5e, 0xa0, 0x7e, 0xc8, 0xfc, 0x3e, 0x86, 0x37, 0xff, 0x8e, 0x08, 0xfb, 0x60, 0x06, 0xc7,
	0x99, 0x60, 0xf2, 0xa8, 0xf3, 0x1b, 0xaf, 0xe0, 0x48, 0x71, 0x48, 0x14, 0x83, 0xb3, 0x7c, 0x27,
	0x98, 0x3c, 0xe8, 0x15, 0xe0, 0x25, 0x94, 0x5d, 0xce, 0xf2, 0x42, 0x30, 0x59, 0xeb, 0x9f, 0x5f,
	0xf3, 0xc5, 0x00, 0x16, 0xf5, 0xd3, 0x2c, 0x39, 0x5d, 0x2c, 0xa0, 0xf2, 0xa1, 0x33, 0x53, 0x30,
	0xe4, 0x63, 0xc8, 0xf6, 0x33, 0xfd, 0x1f, 0xe1, 0x35, 0x5c, 0xf4, 0x26, 0x51, 0xfb, 0x31, 0x5a,
	0x43, 0xce, 0xb6, 0x86, 0xf8, 0x5e, 0x30, 0x59, 0xe8, 0xf3, 0x19, 0x3f, 0x2f, 0xf4, 0x8e, 0x1a,
	0xfb, 0xbb, 0xc9, 0xa3, 0x4f, 0xb4, 0x95, 0x62, 0x1b, 0x29, 0x54, 0x70, 0xc8, 0x95, 0x79, 0x97,
	0xf8, 0x4e, 0x14, 0xb2, 0xba, 0x41, 0xb5, 0xb4, 0xa6, 0xd6, 0xbb, 0xf4, 0xdf, 0xcc, 0x6b, 0x99,
	0x9b, 0xbd, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x68, 0x8e, 0x5b, 0x69, 0x01, 0x00, 0x00,
}
