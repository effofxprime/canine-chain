// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/storage/payment_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StoragePaymentInfo struct {
	Start          time.Time                                `protobuf:"bytes,1,opt,name=start,proto3,stdtime" json:"start"`
	End            time.Time                                `protobuf:"bytes,2,opt,name=end,proto3,stdtime" json:"end"`
	SpaceAvailable int64                                    `protobuf:"varint,3,opt,name=spaceAvailable,proto3" json:"spaceAvailable,omitempty"`
	SpaceUsed      int64                                    `protobuf:"varint,4,opt,name=spaceUsed,proto3" json:"spaceUsed,omitempty"`
	Address        string                                   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Coins          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *StoragePaymentInfo) Reset()         { *m = StoragePaymentInfo{} }
func (m *StoragePaymentInfo) String() string { return proto.CompactTextString(m) }
func (*StoragePaymentInfo) ProtoMessage()    {}
func (*StoragePaymentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e2face38d88409, []int{0}
}
func (m *StoragePaymentInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoragePaymentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoragePaymentInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoragePaymentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoragePaymentInfo.Merge(m, src)
}
func (m *StoragePaymentInfo) XXX_Size() int {
	return m.Size()
}
func (m *StoragePaymentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_StoragePaymentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_StoragePaymentInfo proto.InternalMessageInfo

type PaymentGauge struct {
	Id    []byte                                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Start time.Time                                `protobuf:"bytes,2,opt,name=start,proto3,stdtime" json:"start"`
	End   time.Time                                `protobuf:"bytes,3,opt,name=end,proto3,stdtime" json:"end"`
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *PaymentGauge) Reset()         { *m = PaymentGauge{} }
func (m *PaymentGauge) String() string { return proto.CompactTextString(m) }
func (*PaymentGauge) ProtoMessage()    {}
func (*PaymentGauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1e2face38d88409, []int{1}
}
func (m *PaymentGauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PaymentGauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PaymentGauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PaymentGauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentGauge.Merge(m, src)
}
func (m *PaymentGauge) XXX_Size() int {
	return m.Size()
}
func (m *PaymentGauge) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentGauge.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentGauge proto.InternalMessageInfo

func (m *PaymentGauge) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PaymentGauge) GetStart() time.Time {
	if m != nil {
		return m.Start
	}
	return time.Time{}
}

func (m *PaymentGauge) GetEnd() time.Time {
	if m != nil {
		return m.End
	}
	return time.Time{}
}

func (m *PaymentGauge) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func init() {
	proto.RegisterType((*StoragePaymentInfo)(nil), "canine_chain.storage.StoragePaymentInfo")
	proto.RegisterType((*PaymentGauge)(nil), "canine_chain.storage.PaymentGauge")
}

func init() {
	proto.RegisterFile("canine_chain/storage/payment_info.proto", fileDescriptor_b1e2face38d88409)
}

var fileDescriptor_b1e2face38d88409 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xe3, 0xa4, 0x57, 0x5a, 0xb7, 0xea, 0x60, 0x75, 0x48, 0x4f, 0x28, 0x39, 0x75, 0x80,
	0x2c, 0xb5, 0xe9, 0x21, 0x31, 0x74, 0xe3, 0x18, 0x10, 0x52, 0x07, 0x14, 0x60, 0x61, 0x39, 0xbd,
	0x24, 0xbe, 0xd4, 0xf4, 0x62, 0x47, 0x67, 0x5f, 0x45, 0xbf, 0x01, 0x63, 0x3f, 0x42, 0x67, 0x3e,
	0x49, 0xc7, 0x2e, 0x48, 0x4c, 0x14, 0xdd, 0x2d, 0x7c, 0x01, 0x76, 0x14, 0xdb, 0xd7, 0x56, 0x6c,
	0x45, 0x62, 0x8a, 0xdf, 0xcb, 0xfb, 0xff, 0xff, 0xf6, 0x4f, 0x7a, 0xf8, 0x69, 0x09, 0x52, 0x48,
	0x3e, 0x2e, 0x4f, 0x40, 0x48, 0xa6, 0x8d, 0x9a, 0x41, 0xcd, 0x59, 0x0b, 0xe7, 0x0d, 0x97, 0x66,
	0x2c, 0xe4, 0x44, 0xd1, 0x76, 0xa6, 0x8c, 0x22, 0xbb, 0xf7, 0x07, 0xa9, 0x1f, 0xec, 0x27, 0xa5,
	0xd2, 0x8d, 0xd2, 0xac, 0x00, 0xcd, 0xd9, 0xd9, 0x61, 0xc1, 0x0d, 0x1c, 0xb2, 0x52, 0x09, 0xe9,
	0x54, 0xfd, 0xdd, 0x5a, 0xd5, 0xca, 0x1e, 0x59, 0x77, 0xf2, 0xdd, 0xb4, 0x56, 0xaa, 0x9e, 0x72,
	0x66, 0xab, 0x62, 0x3e, 0x61, 0x46, 0x34, 0x5c, 0x1b, 0x68, 0x5a, 0x3f, 0xb0, 0xe7, 0x6c, 0xc7,
	0x4e, 0xe9, 0x0a, 0xf7, 0x6b, 0xff, 0x5b, 0x88, 0xc9, 0x3b, 0x97, 0xfe, 0xd6, 0xdd, 0xf2, 0x8d,
	0x9c, 0x28, 0x72, 0x84, 0x7b, 0xda, 0xc0, 0xcc, 0xc4, 0x68, 0x80, 0xb2, 0xad, 0x61, 0x9f, 0xba,
	0x08, 0xba, 0x8a, 0xa0, 0xef, 0x57, 0x11, 0xa3, 0x8d, 0xab, 0x1f, 0x69, 0x70, 0x71, 0x93, 0xa2,
	0xdc, 0x49, 0xc8, 0x0b, 0x1c, 0x71, 0x59, 0xc5, 0xe1, 0x03, 0x94, 0x9d, 0x80, 0x3c, 0xc1, 0x3b,
	0xba, 0x85, 0x92, 0xbf, 0x3c, 0x03, 0x31, 0x85, 0x62, 0xca, 0xe3, 0x68, 0x80, 0xb2, 0x28, 0xff,
	0xab, 0x4b, 0x1e, 0xe3, 0x4d, 0xdb, 0xf9, 0xa0, 0x79, 0x15, 0xaf, 0xd9, 0x91, 0xbb, 0x06, 0x89,
	0xf1, 0x23, 0xa8, 0xaa, 0x19, 0xd7, 0x3a, 0xee, 0x0d, 0x50, 0xb6, 0x99, 0xaf, 0x4a, 0x02, 0xb8,
	0xd7, 0xa1, 0xd4, 0xf1, 0xfa, 0x20, 0xca, 0xb6, 0x86, 0x7b, 0xd4, 0x83, 0xe8, 0x60, 0x53, 0x0f,
	0x9b, 0xbe, 0x52, 0x42, 0x8e, 0x9e, 0x75, 0x17, 0xfb, 0x7a, 0x93, 0x66, 0xb5, 0x30, 0x27, 0xf3,
	0x82, 0x96, 0xaa, 0xf1, 0xd4, 0xfc, 0xe7, 0x40, 0x57, 0xa7, 0xcc, 0x9c, 0xb7, 0x5c, 0x5b, 0x81,
	0xce, 0x9d, 0xf3, 0xd1, 0xc6, 0x97, 0xcb, 0x34, 0xf8, 0x75, 0x99, 0x06, 0xfb, 0xbf, 0x11, 0xde,
	0xf6, 0x40, 0x5f, 0xc3, 0xbc, 0xe6, 0x64, 0x07, 0x87, 0xa2, 0xb2, 0x38, 0xb7, 0xf3, 0x50, 0x54,
	0x77, 0x84, 0xc3, 0x7f, 0x26, 0x1c, 0x3d, 0x94, 0xf0, 0x2d, 0x81, 0xb5, 0xff, 0x45, 0x60, 0x74,
	0x7c, 0xb5, 0x48, 0xd0, 0xf5, 0x22, 0x41, 0x3f, 0x17, 0x09, 0xba, 0x58, 0x26, 0xc1, 0xf5, 0x32,
	0x09, 0xbe, 0x2f, 0x93, 0xe0, 0xe3, 0xf0, 0x9e, 0xd5, 0x27, 0x28, 0x4f, 0x61, 0x7a, 0x0c, 0x85,
	0x66, 0x6e, 0x0f, 0x0e, 0xdc, 0xc2, 0x7c, 0xbe, 0x5d, 0x19, 0x6b, 0x5d, 0xac, 0xdb, 0x37, 0x3d,
	0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xe1, 0x88, 0x16, 0x57, 0x03, 0x00, 0x00,
}

func (m *StoragePaymentInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoragePaymentInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoragePaymentInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPaymentInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPaymentInfo(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SpaceUsed != 0 {
		i = encodeVarintPaymentInfo(dAtA, i, uint64(m.SpaceUsed))
		i--
		dAtA[i] = 0x20
	}
	if m.SpaceAvailable != 0 {
		i = encodeVarintPaymentInfo(dAtA, i, uint64(m.SpaceAvailable))
		i--
		dAtA[i] = 0x18
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPaymentInfo(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintPaymentInfo(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PaymentGauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PaymentGauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PaymentGauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPaymentInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.End, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.End):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintPaymentInfo(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	n4, err4 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Start, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Start):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintPaymentInfo(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintPaymentInfo(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPaymentInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovPaymentInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoragePaymentInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovPaymentInfo(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovPaymentInfo(uint64(l))
	if m.SpaceAvailable != 0 {
		n += 1 + sovPaymentInfo(uint64(m.SpaceAvailable))
	}
	if m.SpaceUsed != 0 {
		n += 1 + sovPaymentInfo(uint64(m.SpaceUsed))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPaymentInfo(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovPaymentInfo(uint64(l))
		}
	}
	return n
}

func (m *PaymentGauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovPaymentInfo(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Start)
	n += 1 + l + sovPaymentInfo(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.End)
	n += 1 + l + sovPaymentInfo(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovPaymentInfo(uint64(l))
		}
	}
	return n
}

func sovPaymentInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPaymentInfo(x uint64) (n int) {
	return sovPaymentInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoragePaymentInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPaymentInfo
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
			return fmt.Errorf("proto: StoragePaymentInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoragePaymentInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceAvailable", wireType)
			}
			m.SpaceAvailable = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceAvailable |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceUsed", wireType)
			}
			m.SpaceUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceUsed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
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
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types1.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPaymentInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PaymentGauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPaymentInfo
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
			return fmt.Errorf("proto: PaymentGauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PaymentGauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Start, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.End, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPaymentInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types1.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPaymentInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPaymentInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPaymentInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPaymentInfo
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
					return 0, ErrIntOverflowPaymentInfo
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
					return 0, ErrIntOverflowPaymentInfo
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
				return 0, ErrInvalidLengthPaymentInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPaymentInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPaymentInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPaymentInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPaymentInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPaymentInfo = fmt.Errorf("proto: unexpected end of group")
)
