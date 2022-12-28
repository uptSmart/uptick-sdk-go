// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/crypto/multisig/keys.proto

package multisig

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/uptsmart/uptick-sdk-go/common/codec/types"
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

// LegacyAminoPubKey specifies a public key type
// which nests multiple public keys and a threshold,
// it uses legacy amino address rules.
type LegacyAminoPubKey struct {
	Threshold uint32       `protobuf:"varint,1,opt,name=threshold,proto3" json:"threshold,omitempty" yaml:"threshold"`
	PubKeys   []*types.Any `protobuf:"bytes,2,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty" yaml:"pubkeys"`
}

func (m *LegacyAminoPubKey) Reset()         { *m = LegacyAminoPubKey{} }
func (m *LegacyAminoPubKey) String() string { return proto.CompactTextString(m) }
func (*LegacyAminoPubKey) ProtoMessage()    {}
func (*LegacyAminoPubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_46b57537e097d47d, []int{0}
}
func (m *LegacyAminoPubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyAminoPubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LegacyAminoPubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LegacyAminoPubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyAminoPubKey.Merge(m, src)
}
func (m *LegacyAminoPubKey) XXX_Size() int {
	return m.Size()
}
func (m *LegacyAminoPubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyAminoPubKey.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyAminoPubKey proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LegacyAminoPubKey)(nil), "cosmos.crypto.multisig.LegacyAminoPubKey")
}

func init() { proto.RegisterFile("cosmos/crypto/multisig/keys.proto", fileDescriptor_46b57537e097d47d) }

var fileDescriptor_46b57537e097d47d = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x86, 0x67, 0xee, 0x15, 0xc5, 0x29, 0x8a, 0x96, 0x22, 0xb5, 0x60, 0xa6, 0xce, 0xaa, 0x9b,
	0x26, 0x50, 0x77, 0x05, 0x17, 0xed, 0x56, 0x17, 0x52, 0x04, 0xc1, 0x8d, 0x4c, 0xd2, 0x31, 0x0d,
	0x9d, 0xcc, 0x19, 0x26, 0xc9, 0x22, 0x6f, 0xe0, 0xd2, 0x47, 0x10, 0x7c, 0x19, 0x97, 0x5d, 0xba,
	0x2a, 0x32, 0x7d, 0x83, 0x3e, 0x81, 0xa4, 0x61, 0xea, 0xee, 0x84, 0xff, 0xcb, 0xf9, 0x0e, 0x7f,
	0x74, 0xcd, 0x40, 0x49, 0x50, 0x84, 0x55, 0xb6, 0xd4, 0x40, 0xa4, 0xc9, 0xb5, 0x50, 0x82, 0x93,
	0x65, 0x66, 0x15, 0x2e, 0x2b, 0xd0, 0xd0, 0xbe, 0xf0, 0x08, 0xf6, 0x08, 0x6e, 0x90, 0x5e, 0x87,
	0x03, 0x87, 0x1d, 0x42, 0xdc, 0xe4, 0xe9, 0xde, 0x25, 0x07, 0xe0, 0x79, 0x46, 0x76, 0x2f, 0x6a,
	0x5e, 0x49, 0x5a, 0x58, 0x1f, 0x25, 0x9f, 0x61, 0x74, 0x7e, 0x9f, 0xf1, 0x94, 0xd9, 0x89, 0x14,
	0x05, 0x3c, 0x18, 0x7a, 0x97, 0xd9, 0xf6, 0x28, 0x3a, 0xd6, 0x8b, 0x2a, 0x53, 0x0b, 0xc8, 0xe7,
	0xdd, 0xb0, 0x1f, 0x0e, 0x4e, 0xa6, 0x9d, 0xed, 0x3a, 0x3e, 0xb3, 0xa9, 0xcc, 0xc7, 0xc9, 0x3e,
	0x4a, 0x66, 0x7f, 0x58, 0xfb, 0x31, 0x6a, 0x95, 0x86, 0xe6, 0x82, 0xbd, 0xb8, 0x3b, 0xbb, 0xff,
	0xfa, 0xff, 0x07, 0xad, 0x51, 0x07, 0x7b, 0x35, 0x6e, 0xd4, 0x78, 0x52, 0xd8, 0xe9, 0x55, 0xbd,
	0x8e, 0x8f, 0xbc, 0x4a, 0x6d, 0xd7, 0xf1, 0xa9, 0x5f, 0x5b, 0x1a, 0xea, 0x7e, 0x26, 0xb3, 0xc8,
	0xef, 0x71, 0xe9, 0xf8, 0xe0, 0xed, 0x23, 0x0e, 0xa6, 0x4f, 0x5f, 0x35, 0x0a, 0x57, 0x35, 0x0a,
	0x7f, 0x6a, 0x14, 0xbe, 0x6f, 0x50, 0xb0, 0xda, 0xa0, 0xe0, 0x7b, 0x83, 0x82, 0xe7, 0x5b, 0x2e,
	0xf4, 0xc2, 0x50, 0xcc, 0x40, 0x12, 0x53, 0x6a, 0x25, 0xd3, 0x4a, 0xbb, 0x41, 0xb0, 0xe5, 0x50,
	0xcd, 0x97, 0x43, 0x0e, 0x84, 0x81, 0x94, 0x50, 0x34, 0x6d, 0x3a, 0xc5, 0xbe, 0x52, 0x7a, 0xb8,
	0xbb, 0xeb, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x6a, 0x3e, 0x42, 0x73, 0x01, 0x00, 0x00,
}

func (m *LegacyAminoPubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyAminoPubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LegacyAminoPubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKeys) > 0 {
		for iNdEx := len(m.PubKeys) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PubKeys[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintKeys(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Threshold != 0 {
		i = encodeVarintKeys(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeys(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeys(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LegacyAminoPubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Threshold != 0 {
		n += 1 + sovKeys(uint64(m.Threshold))
	}
	if len(m.PubKeys) > 0 {
		for _, e := range m.PubKeys {
			l = e.Size()
			n += 1 + l + sovKeys(uint64(l))
		}
	}
	return n
}

func sovKeys(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeys(x uint64) (n int) {
	return sovKeys(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LegacyAminoPubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeys
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
			return fmt.Errorf("proto: LegacyAminoPubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyAminoPubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKeys", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeys
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
				return ErrInvalidLengthKeys
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthKeys
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKeys = append(m.PubKeys, &types.Any{})
			if err := m.PubKeys[len(m.PubKeys)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKeys(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeys
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
func skipKeys(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
					return 0, ErrIntOverflowKeys
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
				return 0, ErrInvalidLengthKeys
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeys
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeys
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeys        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeys          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeys = fmt.Errorf("proto: unexpected end of group")
)