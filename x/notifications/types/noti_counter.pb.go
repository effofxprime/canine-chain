// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/notifications/noti_counter.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NotiCounter struct {
	Address        string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Counter        uint64 `protobuf:"varint,2,opt,name=counter,proto3" json:"counter,omitempty"`
	BlockedSenders string `protobuf:"bytes,3,opt,name=blockedSenders,proto3" json:"blockedSenders,omitempty"`
}

func (m *NotiCounter) Reset()         { *m = NotiCounter{} }
func (m *NotiCounter) String() string { return proto.CompactTextString(m) }
func (*NotiCounter) ProtoMessage()    {}
func (*NotiCounter) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfd479ec70d76f20, []int{0}
}

func (m *NotiCounter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *NotiCounter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NotiCounter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *NotiCounter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotiCounter.Merge(m, src)
}

func (m *NotiCounter) XXX_Size() int {
	return m.Size()
}

func (m *NotiCounter) XXX_DiscardUnknown() {
	xxx_messageInfo_NotiCounter.DiscardUnknown(m)
}

var xxx_messageInfo_NotiCounter proto.InternalMessageInfo

func (m *NotiCounter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NotiCounter) GetCounter() uint64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *NotiCounter) GetBlockedSenders() string {
	if m != nil {
		return m.BlockedSenders
	}
	return ""
}

func init() {
	proto.RegisterType((*NotiCounter)(nil), "canine_chain.notifications.NotiCounter")
}

func init() {
	proto.RegisterFile("canine_chain/notifications/noti_counter.proto", fileDescriptor_dfd479ec70d76f20)
}

var fileDescriptor_dfd479ec70d76f20 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x4e, 0xcc, 0xcb,
	0xcc, 0x4b, 0x8d, 0x4f, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c,
	0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x06, 0xf3, 0xe2, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x52, 0x8b,
	0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4, 0x90, 0x95, 0xeb, 0xa1, 0x28, 0x57, 0xca, 0xe4,
	0xe2, 0xf6, 0xcb, 0x2f, 0xc9, 0x74, 0x86, 0x68, 0x10, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0x49, 0x29,
	0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x41, 0x32, 0x50, 0x53,
	0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x60, 0x5c, 0x21, 0x35, 0x2e, 0xbe, 0xa4, 0x9c, 0xfc,
	0xe4, 0xec, 0xd4, 0x94, 0xe0, 0xd4, 0xbc, 0x94, 0xd4, 0xa2, 0x62, 0x09, 0x66, 0xb0, 0x56, 0x34,
	0x51, 0xa7, 0xa0, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x48, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4a, 0x4c, 0xce, 0x4e, 0xcc, 0xf1,
	0x49, 0x4c, 0x2a, 0xd6, 0x87, 0x38, 0x5b, 0x17, 0xe2, 0xcb, 0x0a, 0x34, 0x7f, 0x96, 0x54, 0x16,
	0xa4, 0x16, 0x27, 0xb1, 0x81, 0x7d, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xc7, 0x99,
	0xac, 0x12, 0x01, 0x00, 0x00,
}

func (m *NotiCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NotiCounter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NotiCounter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockedSenders) > 0 {
		i -= len(m.BlockedSenders)
		copy(dAtA[i:], m.BlockedSenders)
		i = encodeVarintNotiCounter(dAtA, i, uint64(len(m.BlockedSenders)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Counter != 0 {
		i = encodeVarintNotiCounter(dAtA, i, uint64(m.Counter))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNotiCounter(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNotiCounter(dAtA []byte, offset int, v uint64) int {
	offset -= sovNotiCounter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *NotiCounter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNotiCounter(uint64(l))
	}
	if m.Counter != 0 {
		n += 1 + sovNotiCounter(uint64(m.Counter))
	}
	l = len(m.BlockedSenders)
	if l > 0 {
		n += 1 + l + sovNotiCounter(uint64(l))
	}
	return n
}

func sovNotiCounter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozNotiCounter(x uint64) (n int) {
	return sovNotiCounter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *NotiCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNotiCounter
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
			return fmt.Errorf("proto: NotiCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NotiCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotiCounter
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
				return ErrInvalidLengthNotiCounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotiCounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counter", wireType)
			}
			m.Counter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotiCounter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Counter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockedSenders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNotiCounter
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
				return ErrInvalidLengthNotiCounter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNotiCounter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockedSenders = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNotiCounter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNotiCounter
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

func skipNotiCounter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNotiCounter
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
					return 0, ErrIntOverflowNotiCounter
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
					return 0, ErrIntOverflowNotiCounter
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
				return 0, ErrInvalidLengthNotiCounter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNotiCounter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNotiCounter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNotiCounter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNotiCounter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNotiCounter = fmt.Errorf("proto: unexpected end of group")
)