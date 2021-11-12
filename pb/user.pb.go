// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{0}
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

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type Users struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Users) Reset()         { *m = Users{} }
func (m *Users) String() string { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()    {}
func (*Users) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{1}
}

func (m *Users) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Users.Unmarshal(m, b)
}
func (m *Users) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Users.Marshal(b, m, deterministic)
}
func (m *Users) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Users.Merge(m, src)
}
func (m *Users) XXX_Size() int {
	return xxx_messageInfo_Users.Size(m)
}
func (m *Users) XXX_DiscardUnknown() {
	xxx_messageInfo_Users.DiscardUnknown(m)
}

var xxx_messageInfo_Users proto.InternalMessageInfo

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_d570e3e37e5899c5, []int{2}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*Users)(nil), "pb.Users")
	proto.RegisterType((*Empty)(nil), "pb.Empty")
}

func init() { proto.RegisterFile("proto/user.proto", fileDescriptor_d570e3e37e5899c5) }

var fileDescriptor_d570e3e37e5899c5 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x9a, 0xd8, 0x66, 0x0a, 0x22, 0x83, 0x87, 0x10, 0x50, 0x6b, 0x2e, 0xf6, 0x62,
	0x22, 0xf5, 0xe4, 0xb1, 0x82, 0xf4, 0x5e, 0xf1, 0x22, 0x5e, 0xf2, 0x31, 0xc4, 0x85, 0x5d, 0x77,
	0xd9, 0xdd, 0x08, 0xfe, 0x7b, 0xd9, 0x89, 0xe6, 0xd6, 0xdb, 0xf3, 0x3e, 0xf3, 0xee, 0x07, 0x03,
	0x17, 0xc6, 0x6a, 0xaf, 0xeb, 0xd1, 0x91, 0xad, 0x18, 0x31, 0x36, 0x6d, 0xf9, 0x01, 0xc9, 0x9b,
	0x23, 0x8b, 0xe7, 0x10, 0x8b, 0x3e, 0x8f, 0x36, 0xd1, 0x36, 0x3b, 0xc6, 0xa2, 0xc7, 0x02, 0x56,
	0xa1, 0xf9, 0xd5, 0x28, 0xca, 0x63, 0xb6, 0x73, 0x46, 0x84, 0x84, 0xfd, 0x82, 0x3d, 0x33, 0x5e,
	0x42, 0x4a, 0xaa, 0x11, 0x32, 0x4f, 0x58, 0x4e, 0xa1, 0xbc, 0x83, 0x34, 0xdc, 0xee, 0xf0, 0xfa,
	0x0f, 0xf2, 0x68, 0xb3, 0xd8, 0xae, 0x77, 0xab, 0xca, 0xb4, 0x55, 0x10, 0xc7, 0x49, 0x97, 0x4b,
	0x48, 0x5f, 0x94, 0xf1, 0x3f, 0x3b, 0x0d, 0xeb, 0x60, 0x5e, 0xc9, 0x7e, 0x8b, 0x8e, 0xf0, 0x0a,
	0x96, 0xfb, 0xbe, 0xe7, 0x1f, 0xce, 0x67, 0x8a, 0x99, 0xc2, 0xf8, 0x40, 0xfe, 0xe4, 0xb8, 0x04,
	0x38, 0x90, 0xdf, 0x4b, 0xc9, 0x29, 0x0b, 0x9e, 0x5f, 0x29, 0xb2, 0xff, 0x8a, 0x7b, 0xbe, 0x7d,
	0xbf, 0x19, 0x84, 0xff, 0x1c, 0xdb, 0xaa, 0xd3, 0xaa, 0x56, 0x8d, 0xa7, 0xd1, 0x49, 0xa1, 0x9a,
	0xa7, 0x87, 0x7a, 0xb0, 0xa6, 0xbb, 0x1f, 0x74, 0x7b, 0xc6, 0xeb, 0x7a, 0xfc, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x40, 0x17, 0x85, 0x0f, 0x42, 0x01, 0x00, 0x00,
}
