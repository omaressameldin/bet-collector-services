// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bet.proto

package grpc_bet

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

type Bet struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,json=authId,proto3" json:"id,omitempty"`
	Description          string               `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Payment              string               `protobuf:"bytes,3,opt,name=payment,proto3" json:"payment,omitempty"`
	BetterId             string               `protobuf:"bytes,4,opt,name=better_id,json=betterId,proto3" json:"better_id,omitempty"`
	AccepterId           string               `protobuf:"bytes,5,opt,name=accepter_id,proto3" json:"accepter_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Bet) Reset()         { *m = Bet{} }
func (m *Bet) String() string { return proto.CompactTextString(m) }
func (*Bet) ProtoMessage()    {}
func (*Bet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{0}
}

func (m *Bet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bet.Unmarshal(m, b)
}
func (m *Bet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bet.Marshal(b, m, deterministic)
}
func (m *Bet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bet.Merge(m, src)
}
func (m *Bet) XXX_Size() int {
	return xxx_messageInfo_Bet.Size(m)
}
func (m *Bet) XXX_DiscardUnknown() {
	xxx_messageInfo_Bet.DiscardUnknown(m)
}

var xxx_messageInfo_Bet proto.InternalMessageInfo

func (m *Bet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Bet) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Bet) GetPayment() string {
	if m != nil {
		return m.Payment
	}
	return ""
}

func (m *Bet) GetBetterId() string {
	if m != nil {
		return m.BetterId
	}
	return ""
}

func (m *Bet) GetAccepterId() string {
	if m != nil {
		return m.AccepterId
	}
	return ""
}

func (m *Bet) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Bet) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type BetUpdate struct {
	Description          *wrappers.StringValue `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Payment              *wrappers.StringValue `protobuf:"bytes,2,opt,name=payment,proto3" json:"payment,omitempty"`
	AccepterId           *wrappers.StringValue `protobuf:"bytes,5,opt,name=accepter_id,proto3" json:"accepter_id,omitempty"`
	UpdatedAt            *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BetUpdate) Reset()         { *m = BetUpdate{} }
func (m *BetUpdate) String() string { return proto.CompactTextString(m) }
func (*BetUpdate) ProtoMessage()    {}
func (*BetUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{1}
}

func (m *BetUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BetUpdate.Unmarshal(m, b)
}
func (m *BetUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BetUpdate.Marshal(b, m, deterministic)
}
func (m *BetUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BetUpdate.Merge(m, src)
}
func (m *BetUpdate) XXX_Size() int {
	return xxx_messageInfo_BetUpdate.Size(m)
}
func (m *BetUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_BetUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_BetUpdate proto.InternalMessageInfo

func (m *BetUpdate) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *BetUpdate) GetPayment() *wrappers.StringValue {
	if m != nil {
		return m.Payment
	}
	return nil
}

func (m *BetUpdate) GetAccepterId() *wrappers.StringValue {
	if m != nil {
		return m.AccepterId
	}
	return nil
}

func (m *BetUpdate) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreateRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Bet                  *Bet     `protobuf:"bytes,2,opt,name=bet,proto3" json:"bet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{2}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CreateRequest) GetBet() *Bet {
	if m != nil {
		return m.Bet
	}
	return nil
}

type CreateResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{3}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

type ReadRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{4}
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

type ReadResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Bet                  *Bet     `protobuf:"bytes,2,opt,name=bet,proto3" json:"bet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{5}
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

func (m *ReadResponse) GetBet() *Bet {
	if m != nil {
		return m.Bet
	}
	return nil
}

type ReadAllRequest struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 int32    `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllRequest) Reset()         { *m = ReadAllRequest{} }
func (m *ReadAllRequest) String() string { return proto.CompactTextString(m) }
func (*ReadAllRequest) ProtoMessage()    {}
func (*ReadAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{6}
}

func (m *ReadAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllRequest.Unmarshal(m, b)
}
func (m *ReadAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllRequest.Marshal(b, m, deterministic)
}
func (m *ReadAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllRequest.Merge(m, src)
}
func (m *ReadAllRequest) XXX_Size() int {
	return xxx_messageInfo_ReadAllRequest.Size(m)
}
func (m *ReadAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllRequest proto.InternalMessageInfo

func (m *ReadAllRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ReadAllRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReadAllRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ReadAllResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Bet                  []*Bet   `protobuf:"bytes,2,rep,name=bet,proto3" json:"bet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadAllResponse) Reset()         { *m = ReadAllResponse{} }
func (m *ReadAllResponse) String() string { return proto.CompactTextString(m) }
func (*ReadAllResponse) ProtoMessage()    {}
func (*ReadAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{7}
}

func (m *ReadAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadAllResponse.Unmarshal(m, b)
}
func (m *ReadAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadAllResponse.Marshal(b, m, deterministic)
}
func (m *ReadAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadAllResponse.Merge(m, src)
}
func (m *ReadAllResponse) XXX_Size() int {
	return xxx_messageInfo_ReadAllResponse.Size(m)
}
func (m *ReadAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadAllResponse proto.InternalMessageInfo

func (m *ReadAllResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadAllResponse) GetBet() []*Bet {
	if m != nil {
		return m.Bet
	}
	return nil
}

type UpdateRequest struct {
	Api                  string     `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id                   string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	BetUpdate            *BetUpdate `protobuf:"bytes,3,opt,name=BetUpdate,proto3" json:"BetUpdate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{8}
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

func (m *UpdateRequest) GetBetUpdate() *BetUpdate {
	if m != nil {
		return m.BetUpdate
	}
	return nil
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
	return fileDescriptor_8bffc34e5f5c53ed, []int{9}
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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bffc34e5f5c53ed, []int{10}
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
	return fileDescriptor_8bffc34e5f5c53ed, []int{11}
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
	proto.RegisterType((*Bet)(nil), "grpc.bet.Bet")
	proto.RegisterType((*BetUpdate)(nil), "grpc.bet.BetUpdate")
	proto.RegisterType((*CreateRequest)(nil), "grpc.bet.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "grpc.bet.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "grpc.bet.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "grpc.bet.ReadResponse")
	proto.RegisterType((*ReadAllRequest)(nil), "grpc.bet.ReadAllRequest")
	proto.RegisterType((*ReadAllResponse)(nil), "grpc.bet.ReadAllResponse")
	proto.RegisterType((*UpdateRequest)(nil), "grpc.bet.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "grpc.bet.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "grpc.bet.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "grpc.bet.DeleteResponse")
}

func init() { proto.RegisterFile("bet.proto", fileDescriptor_8bffc34e5f5c53ed) }

var fileDescriptor_8bffc34e5f5c53ed = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0xc4, 0x69, 0x26, 0x24, 0xa0, 0x05, 0x5a, 0x63, 0x10, 0x8d, 0x7c, 0xea, 0xc9,
	0x51, 0x83, 0x40, 0x42, 0x91, 0x22, 0xc5, 0xf4, 0xd2, 0x1b, 0x72, 0x81, 0x03, 0x97, 0x6a, 0x63,
	0x0f, 0x61, 0x25, 0x27, 0x5e, 0xec, 0x35, 0x88, 0x5f, 0xe0, 0x93, 0xf8, 0x38, 0x84, 0xbc, 0xbb,
	0x8e, 0xed, 0xa4, 0x56, 0xda, 0xdb, 0xee, 0xcc, 0xbc, 0x37, 0x33, 0x6f, 0x5f, 0x62, 0x18, 0xac,
	0x50, 0x78, 0x3c, 0x4d, 0x44, 0x42, 0x4e, 0xd6, 0x29, 0x0f, 0xbd, 0x15, 0x0a, 0xe7, 0x7c, 0x9d,
	0x24, 0xeb, 0x18, 0xa7, 0x32, 0xbe, 0xca, 0xbf, 0x4d, 0x05, 0xdb, 0x60, 0x26, 0xe8, 0x86, 0xab,
	0x52, 0xe7, 0xf5, 0x7e, 0xc1, 0xaf, 0x94, 0x72, 0x8e, 0x69, 0xa6, 0xf2, 0xee, 0x1f, 0x13, 0x3a,
	0x3e, 0x0a, 0x42, 0xc0, 0x64, 0x91, 0x6d, 0x4c, 0x8c, 0x8b, 0x41, 0x60, 0xd1, 0x5c, 0x7c, 0xbf,
	0x8e, 0xc8, 0x04, 0x86, 0x11, 0x66, 0x61, 0xca, 0xb8, 0x60, 0xc9, 0xd6, 0x36, 0x65, 0xb2, 0x1e,
	0x22, 0x36, 0xf4, 0x39, 0xfd, 0xbd, 0xc1, 0xad, 0xb0, 0x3b, 0x32, 0x5b, 0x5e, 0xc9, 0x4b, 0x39,
	0xaf, 0xc0, 0xf4, 0x96, 0x45, 0x76, 0x57, 0xe6, 0x4e, 0x54, 0x40, 0x11, 0xd3, 0x30, 0x44, 0xae,
	0xd3, 0x3d, 0x45, 0x5c, 0x0b, 0x91, 0xf7, 0x00, 0x61, 0x8a, 0x54, 0x60, 0x74, 0x4b, 0x85, 0x6d,
	0x4d, 0x8c, 0x8b, 0xe1, 0xcc, 0xf1, 0xd4, 0x2e, 0x5e, 0xb9, 0x8b, 0xf7, 0xa9, 0x5c, 0x36, 0x18,
	0xe8, 0xea, 0xa5, 0x28, 0xa0, 0x39, 0x8f, 0x4a, 0x68, 0xff, 0x38, 0x54, 0x57, 0x2f, 0x85, 0xfb,
	0xcf, 0x80, 0x81, 0x8f, 0xe2, 0xb3, 0x0c, 0x90, 0x45, 0x73, 0x7d, 0x43, 0x32, 0xbd, 0x3a, 0x60,
	0xba, 0x11, 0x29, 0xdb, 0xae, 0xbf, 0xd0, 0x38, 0xc7, 0xa6, 0x38, 0xef, 0x2a, 0x71, 0xcc, 0x7b,
	0x60, 0x77, 0xd2, 0x2d, 0x0e, 0xd5, 0x39, 0xda, 0x77, 0x4f, 0xbb, 0x9a, 0x00, 0xdd, 0x87, 0x08,
	0xe0, 0xc3, 0xe8, 0x83, 0x14, 0x32, 0xc0, 0x1f, 0x39, 0x66, 0x82, 0x3c, 0x81, 0x0e, 0xe5, 0x4c,
	0xfb, 0xa2, 0x38, 0x92, 0x73, 0xe8, 0xac, 0xb0, 0xdc, 0x68, 0xe4, 0x95, 0x4e, 0xf4, 0x7c, 0x14,
	0x41, 0x91, 0x71, 0x5d, 0x18, 0x97, 0x1c, 0x19, 0x4f, 0xb6, 0x19, 0x1e, 0x92, 0xb8, 0x53, 0x18,
	0x06, 0x48, 0xa3, 0xf6, 0x2e, 0x63, 0x69, 0x47, 0xe5, 0x38, 0x93, 0x45, 0xee, 0x12, 0x1e, 0x29,
	0x40, 0x1b, 0xe5, 0xf1, 0xb9, 0x10, 0xc6, 0x05, 0xc5, 0x32, 0x8e, 0xdb, 0xdb, 0x9e, 0x41, 0x3f,
	0xcf, 0x94, 0xec, 0xaa, 0xb7, 0x55, 0x5c, 0xaf, 0x23, 0xf2, 0x0c, 0x7a, 0x31, 0xdb, 0x30, 0x65,
	0xf3, 0x5e, 0xa0, 0x2e, 0x84, 0x40, 0x97, 0xd3, 0x35, 0x4a, 0x8d, 0x7b, 0x81, 0x3c, 0xbb, 0x57,
	0xf0, 0x78, 0xd7, 0xe6, 0xf8, 0xb0, 0x9d, 0x96, 0x61, 0x23, 0x18, 0x29, 0x17, 0xde, 0x5b, 0x22,
	0x72, 0x59, 0xf3, 0xae, 0x1c, 0x73, 0x38, 0x7b, 0xda, 0x60, 0xd6, 0x84, 0x55, 0x55, 0xf1, 0x54,
	0x65, 0x97, 0xd6, 0xa7, 0xba, 0x84, 0xd1, 0x15, 0xc6, 0xf8, 0x80, 0x49, 0x0a, 0xda, 0x12, 0xd2,
	0x46, 0x3b, 0xfb, 0x6b, 0x02, 0xf8, 0x28, 0x6e, 0x30, 0xfd, 0xc9, 0x42, 0x24, 0x73, 0xb0, 0x94,
	0x69, 0xc8, 0x59, 0x35, 0x73, 0xc3, 0x8a, 0x8e, 0x7d, 0x98, 0xd0, 0xec, 0x6f, 0xa1, 0x5b, 0x48,
	0x4e, 0x9e, 0x57, 0x15, 0x35, 0x77, 0x39, 0xa7, 0xfb, 0x61, 0x0d, 0x5b, 0x40, 0x5f, 0xbf, 0x14,
	0xb1, 0x9b, 0x25, 0x95, 0x47, 0x9c, 0x17, 0x77, 0x64, 0x34, 0x7e, 0x0e, 0x96, 0xfe, 0xa7, 0xa8,
	0xcd, 0xdc, 0x78, 0xb5, 0xfa, 0xcc, 0x7b, 0x42, 0xcf, 0xc1, 0x52, 0x1a, 0xd5, 0xc1, 0x0d, 0xa1,
	0xeb, 0xe0, 0xa6, 0x9c, 0xfe, 0x29, 0xec, 0xbe, 0x00, 0x1f, 0x8d, 0xaf, 0xbb, 0xf3, 0xca, 0x92,
	0xbf, 0xee, 0x37, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xde, 0x2d, 0x49, 0x37, 0x2b, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BetServiceClient is the client API for BetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BetServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type betServiceClient struct {
	cc *grpc.ClientConn
}

func NewBetServiceClient(cc *grpc.ClientConn) BetServiceClient {
	return &betServiceClient{cc}
}

func (c *betServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/grpc.bet.BetService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/grpc.bet.BetService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) ReadAll(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/grpc.bet.BetService/ReadAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/grpc.bet.BetService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *betServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/grpc.bet.BetService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BetServiceServer is the server API for BetService service.
type BetServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

func RegisterBetServiceServer(s *grpc.Server, srv BetServiceServer) {
	s.RegisterService(&_BetService_serviceDesc, srv)
}

func _BetService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.bet.BetService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.bet.BetService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_ReadAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).ReadAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.bet.BetService/ReadAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).ReadAll(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.bet.BetService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BetService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BetServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.bet.BetService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BetServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.bet.BetService",
	HandlerType: (*BetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BetService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _BetService_Read_Handler,
		},
		{
			MethodName: "ReadAll",
			Handler:    _BetService_ReadAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BetService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BetService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bet.proto",
}
