// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: login.proto

package login

import (
	common "due-v2-example/shared/pb/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// 登录模式
type LoginMode int32

const (
	LoginMode_None     LoginMode = 0
	LoginMode_Guest    LoginMode = 1
	LoginMode_Mobile   LoginMode = 2
	LoginMode_Account  LoginMode = 3
	LoginMode_Wechat   LoginMode = 4
	LoginMode_Google   LoginMode = 5
	LoginMode_Facebook LoginMode = 6
	LoginMode_Token    LoginMode = 7
)

var LoginMode_name = map[int32]string{
	0: "None",
	1: "Guest",
	2: "Mobile",
	3: "Account",
	4: "Wechat",
	5: "Google",
	6: "Facebook",
	7: "Token",
}

var LoginMode_value = map[string]int32{
	"None":     0,
	"Guest":    1,
	"Mobile":   2,
	"Account":  3,
	"Wechat":   4,
	"Google":   5,
	"Facebook": 6,
	"Token":    7,
}

func (x LoginMode) String() string {
	return proto.EnumName(LoginMode_name, int32(x))
}

func (LoginMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

// 用户登录请求
type LoginReq struct {
	Mode                 LoginMode `protobuf:"varint,1,opt,name=Mode,proto3,enum=login.LoginMode" json:"Mode,omitempty"`
	DeviceID             string    `protobuf:"bytes,2,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
	Account              string    `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
	Password             string    `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	Openid               string    `protobuf:"bytes,5,opt,name=Openid,proto3" json:"Openid,omitempty"`
	Mobile               string    `protobuf:"bytes,6,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	Code                 string    `protobuf:"bytes,7,opt,name=Code,proto3" json:"Code,omitempty"`
	Token                string    `protobuf:"bytes,8,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}
func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return m.Size()
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetMode() LoginMode {
	if m != nil {
		return m.Mode
	}
	return LoginMode_None
}

func (m *LoginReq) GetDeviceID() string {
	if m != nil {
		return m.DeviceID
	}
	return ""
}

func (m *LoginReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginReq) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *LoginReq) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *LoginReq) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *LoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 用户登录响应
type LoginRes struct {
	Code                 common.Code `protobuf:"varint,1,opt,name=Code,proto3,enum=common.Code" json:"Code,omitempty"`
	Token                string      `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LoginRes) Reset()         { *m = LoginRes{} }
func (m *LoginRes) String() string { return proto.CompactTextString(m) }
func (*LoginRes) ProtoMessage()    {}
func (*LoginRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}
func (m *LoginRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LoginRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LoginRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LoginRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRes.Merge(m, src)
}
func (m *LoginRes) XXX_Size() int {
	return m.Size()
}
func (m *LoginRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRes proto.InternalMessageInfo

func (m *LoginRes) GetCode() common.Code {
	if m != nil {
		return m.Code
	}
	return common.Code_OK
}

func (m *LoginRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("login.LoginMode", LoginMode_name, LoginMode_value)
	proto.RegisterType((*LoginReq)(nil), "login.LoginReq")
	proto.RegisterType((*LoginRes)(nil), "login.LoginRes")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xdb, 0x6a, 0xea, 0x40,
	0x14, 0x86, 0x1d, 0x4d, 0x62, 0x1c, 0x65, 0x33, 0x2c, 0x36, 0x9b, 0xc1, 0x8b, 0x20, 0xb2, 0x2f,
	0x64, 0x83, 0x11, 0x76, 0xfb, 0x02, 0x55, 0xa9, 0x14, 0x6a, 0x5b, 0xa4, 0x50, 0xe8, 0x5d, 0x0e,
	0x0b, 0x8d, 0x87, 0x2c, 0x9b, 0x44, 0xfb, 0x2a, 0x7d, 0xa4, 0x5e, 0xf6, 0x01, 0x7a, 0x51, 0xec,
	0x8b, 0x94, 0x99, 0xd1, 0xd0, 0xde, 0xcd, 0xff, 0xfd, 0xfc, 0x6b, 0xd6, 0x81, 0x37, 0xd7, 0x34,
	0x4f, 0x52, 0x7f, 0x9b, 0x51, 0x41, 0x60, 0x6b, 0xd1, 0xe6, 0x11, 0xc5, 0x68, 0x50, 0xf7, 0x9d,
	0x71, 0xf7, 0x5a, 0xd1, 0x19, 0x3e, 0xc1, 0x5f, 0x6e, 0x4d, 0x29, 0x46, 0xc9, 0x3a, 0xac, 0xf7,
	0xeb, 0xbf, 0xf0, 0x4d, 0x56, 0xdb, 0x8a, 0xcf, 0xb4, 0x0b, 0x6d, 0xee, 0x8e, 0x71, 0x9f, 0x44,
	0x78, 0x35, 0x96, 0xd5, 0x0e, 0xeb, 0x35, 0x66, 0xa5, 0x06, 0xc9, 0xeb, 0x17, 0x51, 0x44, 0xbb,
	0xb4, 0x90, 0x35, 0x6d, 0x9d, 0xa4, 0x4a, 0xdd, 0x05, 0x79, 0xfe, 0x4c, 0x59, 0x2c, 0x2d, 0x93,
	0x3a, 0x69, 0xf8, 0xc3, 0x9d, 0xdb, 0x2d, 0xa6, 0x49, 0x2c, 0x6d, 0xed, 0x1c, 0x95, 0xe2, 0x53,
	0x0a, 0x93, 0x35, 0x4a, 0xc7, 0x70, 0xa3, 0x00, 0xb8, 0x35, 0x52, 0x7d, 0xd6, 0x35, 0xd5, 0x6f,
	0xf8, 0xcd, 0xed, 0x7b, 0x5a, 0x61, 0x2a, 0x5d, 0x0d, 0x8d, 0xe8, 0x0e, 0xcb, 0xe9, 0x72, 0xe8,
	0x1c, 0x53, 0x66, 0xba, 0x96, 0x1f, 0xd1, 0x66, 0x43, 0xa9, 0x3f, 0xd2, 0x93, 0xfd, 0xac, 0x51,
	0xfd, 0x56, 0xe3, 0xdf, 0x92, 0x37, 0xca, 0x15, 0x80, 0xcb, 0xad, 0x1b, 0x4a, 0x51, 0x54, 0xa0,
	0xc1, 0xed, 0xc9, 0x0e, 0xf3, 0x42, 0x30, 0xe0, 0xa7, 0x3e, 0x45, 0x15, 0x9a, 0xe5, 0x06, 0x44,
	0x4d, 0x19, 0x0f, 0x18, 0x2d, 0x82, 0x42, 0x58, 0xea, 0x3d, 0x21, 0x9a, 0xaf, 0x51, 0xd8, 0xd0,
	0xe2, 0xee, 0x65, 0x10, 0x61, 0x48, 0xb4, 0x12, 0x8e, 0xaa, 0xa4, 0x7f, 0x12, 0xf5, 0xe1, 0xf9,
	0xeb, 0xc1, 0x63, 0x6f, 0x07, 0x8f, 0x7d, 0x1c, 0x3c, 0xf6, 0xf2, 0xe9, 0x55, 0x1e, 0xbb, 0xf1,
	0x0e, 0xfb, 0x9b, 0x60, 0xb1, 0xa4, 0x74, 0xde, 0xcf, 0x31, 0xdb, 0x63, 0x36, 0xc8, 0x17, 0x41,
	0x86, 0xf1, 0x60, 0x1b, 0x0e, 0xf4, 0x85, 0x42, 0x47, 0xdf, 0xf2, 0xec, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0x0a, 0x99, 0x1c, 0x63, 0xed, 0x01, 0x00, 0x00,
}

func (m *LoginReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Mobile) > 0 {
		i -= len(m.Mobile)
		copy(dAtA[i:], m.Mobile)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Mobile)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Openid) > 0 {
		i -= len(m.Openid)
		copy(dAtA[i:], m.Openid)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Openid)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Password) > 0 {
		i -= len(m.Password)
		copy(dAtA[i:], m.Password)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Password)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DeviceID) > 0 {
		i -= len(m.DeviceID)
		copy(dAtA[i:], m.DeviceID)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.DeviceID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Mode != 0 {
		i = encodeVarintLogin(dAtA, i, uint64(m.Mode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *LoginRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LoginRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LoginRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintLogin(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintLogin(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLogin(dAtA []byte, offset int, v uint64) int {
	offset -= sovLogin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LoginReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Mode != 0 {
		n += 1 + sovLogin(uint64(m.Mode))
	}
	l = len(m.DeviceID)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.Openid)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.Mobile)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LoginRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovLogin(uint64(m.Code))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLogin(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovLogin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLogin(x uint64) (n int) {
	return sovLogin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LoginReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= LoginMode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeviceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeviceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Openid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Openid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mobile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mobile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LoginRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LoginRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LoginRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= common.Code(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLogin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLogin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLogin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLogin
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLogin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLogin
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLogin
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLogin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLogin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLogin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLogin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLogin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLogin = fmt.Errorf("proto: unexpected end of group")
)
