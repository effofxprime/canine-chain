// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: canine_chain/amm/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the module.
type Params struct {
	MinInitPoolDeposit uint64 `protobuf:"varint,1,opt,name=MinInitPoolDeposit,proto3" json:"MinInitPoolDeposit,omitempty"`
	MaxPoolDenomCount  uint32 `protobuf:"varint,2,opt,name=MaxPoolDenomCount,proto3" json:"MaxPoolDenomCount,omitempty"`
	ProtocolFeeToAddr  string `protobuf:"bytes,3,opt,name=protocol_fee_to_addr,json=protocolFeeToAddr,proto3" json:"protocol_fee_to_addr,omitempty"`
	ProtocolFeeRate    string `protobuf:"bytes,4,opt,name=protocol_fee_rate,json=protocolFeeRate,proto3" json:"protocol_fee_rate,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_81e8eb1f608f1924, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMinInitPoolDeposit() uint64 {
	if m != nil {
		return m.MinInitPoolDeposit
	}
	return 0
}

func (m *Params) GetMaxPoolDenomCount() uint32 {
	if m != nil {
		return m.MaxPoolDenomCount
	}
	return 0
}

func (m *Params) GetProtocolFeeToAddr() string {
	if m != nil {
		return m.ProtocolFeeToAddr
	}
	return ""
}

func (m *Params) GetProtocolFeeRate() string {
	if m != nil {
		return m.ProtocolFeeRate
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "canine_chain.amm.Params")
}

func init() { proto.RegisterFile("canine_chain/amm/params.proto", fileDescriptor_81e8eb1f608f1924) }

var fileDescriptor_81e8eb1f608f1924 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4a, 0xc4, 0x40,
	0x14, 0x85, 0x33, 0xba, 0x2c, 0x38, 0x20, 0xba, 0xc3, 0x16, 0x41, 0x70, 0x0c, 0x56, 0x41, 0x34,
	0x53, 0xd8, 0xd9, 0xf9, 0x83, 0x28, 0xb8, 0xb0, 0x04, 0x2b, 0x9b, 0x70, 0x93, 0x8c, 0xd9, 0xd1,
	0xcc, 0xdc, 0x90, 0xcc, 0xc2, 0xfa, 0x16, 0x96, 0x96, 0x3e, 0x8e, 0x76, 0x5b, 0x5a, 0x4a, 0xf2,
	0x22, 0xe2, 0x04, 0x61, 0x45, 0xbb, 0xcb, 0xfd, 0xce, 0xe1, 0xc0, 0x47, 0x77, 0x33, 0x30, 0xca,
	0xc8, 0x24, 0x9b, 0x81, 0x32, 0x02, 0xb4, 0x16, 0x15, 0xd4, 0xa0, 0x9b, 0xa8, 0xaa, 0xd1, 0x22,
	0xdb, 0x5e, 0xc5, 0x11, 0x68, 0xbd, 0x33, 0x2e, 0xb0, 0x40, 0x07, 0xc5, 0xf7, 0xd5, 0xe7, 0xf6,
	0xdf, 0x09, 0x1d, 0x4e, 0x5d, 0x91, 0x45, 0x94, 0x4d, 0x94, 0xb9, 0x36, 0xca, 0x4e, 0x11, 0xcb,
	0x0b, 0x59, 0x61, 0xa3, 0xac, 0x4f, 0x02, 0x12, 0x0e, 0xe2, 0x7f, 0x08, 0x3b, 0xa4, 0xa3, 0x09,
	0x2c, 0xfa, 0x8f, 0x41, 0x7d, 0x8e, 0x73, 0x63, 0xfd, 0xb5, 0x80, 0x84, 0x9b, 0xf1, 0x5f, 0xc0,
	0x04, 0x1d, 0xbb, 0xc5, 0x0c, 0xcb, 0xe4, 0x5e, 0xca, 0xc4, 0x62, 0x02, 0x79, 0x5e, 0xfb, 0xeb,
	0x01, 0x09, 0x37, 0xe2, 0xd1, 0x0f, 0xbb, 0x94, 0xf2, 0x16, 0x4f, 0xf3, 0xbc, 0x66, 0x07, 0x74,
	0xf4, 0xab, 0x50, 0x83, 0x95, 0xfe, 0xc0, 0xa5, 0xb7, 0x56, 0xd2, 0x31, 0x58, 0x79, 0x32, 0x78,
	0x79, 0xdd, 0xf3, 0xce, 0xae, 0xde, 0x5a, 0x4e, 0x96, 0x2d, 0x27, 0x9f, 0x2d, 0x27, 0xcf, 0x1d,
	0xf7, 0x96, 0x1d, 0xf7, 0x3e, 0x3a, 0xee, 0xdd, 0x45, 0x85, 0xb2, 0xb3, 0x79, 0x1a, 0x65, 0xa8,
	0xc5, 0x03, 0x64, 0x8f, 0x50, 0xde, 0x40, 0xda, 0x88, 0xde, 0xd1, 0x51, 0xaf, 0x70, 0xe1, 0x24,
	0xda, 0xa7, 0x4a, 0x36, 0xe9, 0xd0, 0x0d, 0x1c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x85, 0x65,
	0xe5, 0x13, 0x65, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProtocolFeeRate) > 0 {
		i -= len(m.ProtocolFeeRate)
		copy(dAtA[i:], m.ProtocolFeeRate)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ProtocolFeeRate)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProtocolFeeToAddr) > 0 {
		i -= len(m.ProtocolFeeToAddr)
		copy(dAtA[i:], m.ProtocolFeeToAddr)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ProtocolFeeToAddr)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MaxPoolDenomCount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxPoolDenomCount))
		i--
		dAtA[i] = 0x10
	}
	if m.MinInitPoolDeposit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinInitPoolDeposit))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinInitPoolDeposit != 0 {
		n += 1 + sovParams(uint64(m.MinInitPoolDeposit))
	}
	if m.MaxPoolDenomCount != 0 {
		n += 1 + sovParams(uint64(m.MaxPoolDenomCount))
	}
	l = len(m.ProtocolFeeToAddr)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ProtocolFeeRate)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitPoolDeposit", wireType)
			}
			m.MinInitPoolDeposit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinInitPoolDeposit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPoolDenomCount", wireType)
			}
			m.MaxPoolDenomCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxPoolDenomCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolFeeToAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProtocolFeeToAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProtocolFeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProtocolFeeRate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)