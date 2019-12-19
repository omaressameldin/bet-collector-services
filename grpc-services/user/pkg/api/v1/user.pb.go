// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package grpc_user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string               `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string               `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UserUpdate struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               *wrappers.StringValue `protobuf:"bytes,4,opt,name=avatar,proto3" json:"avatar,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UserUpdate) Reset()         { *m = UserUpdate{} }
func (m *UserUpdate) String() string { return proto.CompactTextString(m) }
func (*UserUpdate) ProtoMessage()    {}
func (*UserUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UserUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserUpdate.Unmarshal(m, b)
}
func (m *UserUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserUpdate.Marshal(b, m, deterministic)
}
func (m *UserUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserUpdate.Merge(m, src)
}
func (m *UserUpdate) XXX_Size() int {
	return xxx_messageInfo_UserUpdate.Size(m)
}
func (m *UserUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_UserUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_UserUpdate proto.InternalMessageInfo

func (m *UserUpdate) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserUpdate) GetAvatar() *wrappers.StringValue {
	if m != nil {
		return m.Avatar
	}
	return nil
}

func (m *UserUpdate) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type LoginRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	AuthId               string   `protobuf:"bytes,2,opt,name=auth_id,json=authId,proto3" json:"auth_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LoginRequest) GetAuthId() string {
	if m != nil {
		return m.AuthId
	}
	return ""
}

type LoginResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *LoginResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type ReadRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReadRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ReadResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type DoesUserExistRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoesUserExistRequest) Reset()         { *m = DoesUserExistRequest{} }
func (m *DoesUserExistRequest) String() string { return proto.CompactTextString(m) }
func (*DoesUserExistRequest) ProtoMessage()    {}
func (*DoesUserExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *DoesUserExistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoesUserExistRequest.Unmarshal(m, b)
}
func (m *DoesUserExistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoesUserExistRequest.Marshal(b, m, deterministic)
}
func (m *DoesUserExistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoesUserExistRequest.Merge(m, src)
}
func (m *DoesUserExistRequest) XXX_Size() int {
	return xxx_messageInfo_DoesUserExistRequest.Size(m)
}
func (m *DoesUserExistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DoesUserExistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DoesUserExistRequest proto.InternalMessageInfo

func (m *DoesUserExistRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DoesUserExistRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DoesUserExistResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Exists               bool     `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DoesUserExistResponse) Reset()         { *m = DoesUserExistResponse{} }
func (m *DoesUserExistResponse) String() string { return proto.CompactTextString(m) }
func (*DoesUserExistResponse) ProtoMessage()    {}
func (*DoesUserExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *DoesUserExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoesUserExistResponse.Unmarshal(m, b)
}
func (m *DoesUserExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoesUserExistResponse.Marshal(b, m, deterministic)
}
func (m *DoesUserExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoesUserExistResponse.Merge(m, src)
}
func (m *DoesUserExistResponse) XXX_Size() int {
	return xxx_messageInfo_DoesUserExistResponse.Size(m)
}
func (m *DoesUserExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DoesUserExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DoesUserExistResponse proto.InternalMessageInfo

func (m *DoesUserExistResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DoesUserExistResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type UpdateRequest struct {
	Api                  string      `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	User                 *UserUpdate `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Token                string      `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *UpdateRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRequest) GetUser() *UserUpdate {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UpdateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type UpdateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type DeleteRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DeleteResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "grpc.user.User")
	proto.RegisterType((*UserUpdate)(nil), "grpc.user.UserUpdate")
	proto.RegisterType((*LoginRequest)(nil), "grpc.user.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "grpc.user.LoginResponse")
	proto.RegisterType((*ReadRequest)(nil), "grpc.user.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "grpc.user.ReadResponse")
	proto.RegisterType((*DoesUserExistRequest)(nil), "grpc.user.DoesUserExistRequest")
	proto.RegisterType((*DoesUserExistResponse)(nil), "grpc.user.DoesUserExistResponse")
	proto.RegisterType((*UpdateRequest)(nil), "grpc.user.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "grpc.user.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "grpc.user.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "grpc.user.DeleteResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x53, 0xc7, 0x25, 0x93, 0x26, 0xa0, 0x55, 0x9b, 0x18, 0x0b, 0xd1, 0x68, 0xb9, 0x94,
	0x8b, 0x8b, 0x02, 0x12, 0x14, 0x89, 0x43, 0x50, 0x03, 0x42, 0xe2, 0x80, 0x5c, 0xca, 0x81, 0x4b,
	0xb5, 0x8d, 0x87, 0xb0, 0x22, 0x89, 0xcd, 0xee, 0xba, 0xf0, 0x57, 0xfc, 0x09, 0x27, 0x3e, 0x08,
	0xed, 0x7a, 0x43, 0xd6, 0x49, 0xac, 0xb6, 0xe2, 0xe6, 0x99, 0x79, 0x6f, 0xfc, 0x66, 0xe6, 0x69,
	0x01, 0x0a, 0x89, 0x22, 0xce, 0x45, 0xa6, 0x32, 0xd2, 0x9a, 0x8a, 0x7c, 0x12, 0xeb, 0x44, 0x74,
	0x38, 0xcd, 0xb2, 0xe9, 0x0c, 0x8f, 0x4d, 0xe1, 0xb2, 0xf8, 0x72, 0xac, 0xf8, 0x1c, 0xa5, 0x62,
	0xf3, 0xbc, 0xc4, 0x46, 0x0f, 0xd7, 0x01, 0x3f, 0x04, 0xcb, 0x73, 0x14, 0xb2, 0xac, 0xd3, 0xdf,
	0x1e, 0xf8, 0xe7, 0x12, 0x05, 0xe9, 0x42, 0x83, 0xa7, 0xa1, 0x37, 0xf0, 0x8e, 0x5a, 0x49, 0x83,
	0xa7, 0x64, 0x1f, 0x9a, 0x38, 0x67, 0x7c, 0x16, 0x36, 0x4c, 0xaa, 0x0c, 0x08, 0x01, 0x7f, 0xc1,
	0xe6, 0x18, 0xee, 0x98, 0xa4, 0xf9, 0x26, 0x3d, 0x08, 0xd8, 0x15, 0x53, 0x4c, 0x84, 0xbe, 0xc9,
	0xda, 0x88, 0x9c, 0x00, 0x4c, 0x04, 0x32, 0x85, 0xe9, 0x05, 0x53, 0x61, 0x73, 0xe0, 0x1d, 0xb5,
	0x87, 0x51, 0x5c, 0xea, 0x89, 0x97, 0x7a, 0xe2, 0x8f, 0x4b, 0xc1, 0x49, 0xcb, 0xa2, 0x47, 0x4a,
	0x53, 0x8b, 0x3c, 0x5d, 0x52, 0x83, 0xeb, 0xa9, 0x16, 0x3d, 0x52, 0xf4, 0x97, 0x07, 0xa0, 0x07,
	0x3a, 0x37, 0x19, 0xf2, 0xc4, 0x0a, 0xf6, 0x4c, 0x8f, 0x07, 0x1b, 0x3d, 0xce, 0x94, 0xe0, 0x8b,
	0xe9, 0x27, 0x36, 0x2b, 0xd0, 0x8e, 0xf3, 0xac, 0x32, 0xce, 0x75, 0x1c, 0x67, 0x58, 0x47, 0xf1,
	0xee, 0x6d, 0x14, 0x9f, 0xc0, 0xde, 0xfb, 0x6c, 0xca, 0x17, 0x09, 0x7e, 0x2f, 0x50, 0x2a, 0x72,
	0x0f, 0x76, 0x58, 0xce, 0xed, 0x29, 0xf4, 0x27, 0xe9, 0xc3, 0x2e, 0x2b, 0xd4, 0xd7, 0x0b, 0x9e,
	0xda, 0x6b, 0x04, 0x3a, 0x7c, 0x97, 0xd2, 0x37, 0xd0, 0xb1, 0x54, 0x99, 0x67, 0x0b, 0x89, 0x5b,
	0xb8, 0x8f, 0xc0, 0xd7, 0x4e, 0x31, 0xc4, 0xf6, 0xf0, 0x6e, 0xfc, 0xcf, 0x3b, 0xb1, 0xde, 0x52,
	0x62, 0x8a, 0x74, 0x0c, 0xed, 0x04, 0x59, 0x5a, 0xaf, 0xa0, 0x74, 0x47, 0xc3, 0x75, 0x87, 0xca,
	0xbe, 0xe1, 0xc2, 0x1a, 0xa1, 0x0c, 0xe8, 0x18, 0xf6, 0xca, 0x36, 0xff, 0xa7, 0xe6, 0x05, 0xec,
	0x9f, 0x66, 0x28, 0x75, 0x66, 0xfc, 0x93, 0x4b, 0x75, 0x63, 0x59, 0x74, 0x04, 0x07, 0x6b, 0xcc,
	0x5a, 0x25, 0x3d, 0x08, 0x50, 0x43, 0xa4, 0xa1, 0xdf, 0x49, 0x6c, 0x44, 0x05, 0x74, 0x4a, 0xeb,
	0xdc, 0x7c, 0x19, 0x8f, 0xed, 0x50, 0x3b, 0x66, 0xa8, 0x83, 0xb5, 0xa1, 0x6c, 0x37, 0x03, 0x59,
	0xed, 0xcd, 0x77, 0xf7, 0x46, 0xa1, 0xbb, 0xfc, 0x67, 0x9d, 0x5e, 0xfa, 0x16, 0x3a, 0xa7, 0x38,
	0xc3, 0xdb, 0xe8, 0xda, 0x7e, 0x24, 0x0a, 0xdd, 0x65, 0xa3, 0xba, 0x9f, 0x0d, 0xff, 0x34, 0xa0,
	0xad, 0xb5, 0x9f, 0xa1, 0xb8, 0xe2, 0x13, 0x24, 0x2f, 0xa1, 0x69, 0x7c, 0x46, 0xfa, 0xce, 0x70,
	0xae, 0x69, 0xa3, 0x70, 0xb3, 0x60, 0xbb, 0x3f, 0x07, 0x5f, 0x9b, 0x82, 0xf4, 0x1c, 0x84, 0x63,
	0xb6, 0xa8, 0xbf, 0x91, 0xb7, 0xc4, 0x04, 0x3a, 0x95, 0x63, 0x92, 0x43, 0x07, 0xb9, 0xcd, 0x20,
	0xd1, 0xa0, 0x1e, 0x60, 0x7b, 0xbe, 0x82, 0xc0, 0x3e, 0x0c, 0xae, 0xe0, 0xca, 0xc1, 0xa3, 0xfb,
	0x5b, 0x2a, 0x2b, 0x7a, 0xb9, 0xbb, 0x0a, 0xbd, 0x72, 0x97, 0x0a, 0xbd, 0xba, 0xe8, 0xd7, 0x21,
	0xac, 0x9e, 0xee, 0x0f, 0xde, 0xe7, 0x55, 0x70, 0x19, 0x98, 0x27, 0xe2, 0xe9, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x71, 0x84, 0x5b, 0xe7, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	DoesUserExist(ctx context.Context, in *DoesUserExistRequest, opts ...grpc.CallOption) (*DoesUserExistResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/grpc.user.UserService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/grpc.user.UserService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DoesUserExist(ctx context.Context, in *DoesUserExistRequest, opts ...grpc.CallOption) (*DoesUserExistResponse, error) {
	out := new(DoesUserExistResponse)
	err := c.cc.Invoke(ctx, "/grpc.user.UserService/DoesUserExist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/grpc.user.UserService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/grpc.user.UserService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
type UserServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	DoesUserExist(context.Context, *DoesUserExistRequest) (*DoesUserExistResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.user.UserService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.user.UserService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DoesUserExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DoesUserExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DoesUserExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.user.UserService/DoesUserExist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DoesUserExist(ctx, req.(*DoesUserExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.user.UserService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.user.UserService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserService_Login_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _UserService_Read_Handler,
		},
		{
			MethodName: "DoesUserExist",
			Handler:    _UserService_DoesUserExist_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
