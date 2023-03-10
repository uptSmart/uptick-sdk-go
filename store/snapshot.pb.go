// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/store/v1beta1/snapshot.proto

package store

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// SnapshotItem is an item contained in a rootmulti.Store snapshot.
type SnapshotItem struct {
	// item is the specific type of snapshot item.
	//
	// Types that are valid to be assigned to Item:
	//	*SnapshotItem_Store
	//	*SnapshotItem_IAVL
	Item isSnapshotItem_Item `protobuf_oneof:"item"`
}

func (m *SnapshotItem) Reset()         { *m = SnapshotItem{} }
func (m *SnapshotItem) String() string { return proto.CompactTextString(m) }
func (*SnapshotItem) ProtoMessage()    {}
func (*SnapshotItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a86cb47d6b1dddc7, []int{0}
}
func (m *SnapshotItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotItem.Merge(m, src)
}
func (m *SnapshotItem) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotItem.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotItem proto.InternalMessageInfo

type isSnapshotItem_Item interface {
	isSnapshotItem_Item()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SnapshotItem_Store struct {
	Store *SnapshotStoreItem `protobuf:"bytes,1,opt,name=store,proto3,oneof" json:"store,omitempty"`
}
type SnapshotItem_IAVL struct {
	IAVL *SnapshotIAVLItem `protobuf:"bytes,2,opt,name=iavl,proto3,oneof" json:"iavl,omitempty"`
}

func (*SnapshotItem_Store) isSnapshotItem_Item() {}
func (*SnapshotItem_IAVL) isSnapshotItem_Item()  {}

func (m *SnapshotItem) GetItem() isSnapshotItem_Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *SnapshotItem) GetStore() *SnapshotStoreItem {
	if x, ok := m.GetItem().(*SnapshotItem_Store); ok {
		return x.Store
	}
	return nil
}

func (m *SnapshotItem) GetIAVL() *SnapshotIAVLItem {
	if x, ok := m.GetItem().(*SnapshotItem_IAVL); ok {
		return x.IAVL
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SnapshotItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SnapshotItem_Store)(nil),
		(*SnapshotItem_IAVL)(nil),
	}
}

// SnapshotStoreItem contains metadata about a snapshotted store.
type SnapshotStoreItem struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *SnapshotStoreItem) Reset()         { *m = SnapshotStoreItem{} }
func (m *SnapshotStoreItem) String() string { return proto.CompactTextString(m) }
func (*SnapshotStoreItem) ProtoMessage()    {}
func (*SnapshotStoreItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a86cb47d6b1dddc7, []int{1}
}
func (m *SnapshotStoreItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotStoreItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotStoreItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotStoreItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotStoreItem.Merge(m, src)
}
func (m *SnapshotStoreItem) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotStoreItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotStoreItem.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotStoreItem proto.InternalMessageInfo

func (m *SnapshotStoreItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// SnapshotIAVLItem is an exported IAVL node.
type SnapshotIAVLItem struct {
	Key     []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value   []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Version int64  `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Height  int32  `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *SnapshotIAVLItem) Reset()         { *m = SnapshotIAVLItem{} }
func (m *SnapshotIAVLItem) String() string { return proto.CompactTextString(m) }
func (*SnapshotIAVLItem) ProtoMessage()    {}
func (*SnapshotIAVLItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a86cb47d6b1dddc7, []int{2}
}
func (m *SnapshotIAVLItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SnapshotIAVLItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SnapshotIAVLItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SnapshotIAVLItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotIAVLItem.Merge(m, src)
}
func (m *SnapshotIAVLItem) XXX_Size() int {
	return m.Size()
}
func (m *SnapshotIAVLItem) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotIAVLItem.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotIAVLItem proto.InternalMessageInfo

func (m *SnapshotIAVLItem) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SnapshotIAVLItem) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SnapshotIAVLItem) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SnapshotIAVLItem) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func init() {
	proto.RegisterType((*SnapshotItem)(nil), "cosmos.base.store.v1beta1.SnapshotItem")
	proto.RegisterType((*SnapshotStoreItem)(nil), "cosmos.base.store.v1beta1.SnapshotStoreItem")
	proto.RegisterType((*SnapshotIAVLItem)(nil), "cosmos.base.store.v1beta1.SnapshotIAVLItem")
}

func init() {
	proto.RegisterFile("cosmos/store/v1beta1/snapshot.proto", fileDescriptor_a86cb47d6b1dddc7)
}

var fileDescriptor_a86cb47d6b1dddc7 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4e, 0x32, 0x31,
	0x1c, 0xc5, 0xa7, 0x1f, 0x03, 0x9f, 0xfe, 0x65, 0x81, 0x0d, 0x31, 0xa3, 0x8b, 0x91, 0xe0, 0x02,
	0x12, 0xa5, 0x13, 0xf4, 0x04, 0x10, 0x17, 0x10, 0x5d, 0x95, 0xc4, 0x85, 0xbb, 0x0e, 0x36, 0x33,
	0x13, 0x28, 0x25, 0xd3, 0x32, 0x89, 0xb7, 0xf0, 0x1a, 0xde, 0xc4, 0x25, 0x4b, 0x57, 0xc6, 0x0c,
	0x17, 0x31, 0x6d, 0x87, 0x8d, 0xc6, 0xc4, 0xdd, 0x7b, 0xcd, 0x7b, 0xbf, 0xfe, 0x93, 0x07, 0x17,
	0x73, 0xa9, 0x84, 0x54, 0x91, 0xd2, 0x32, 0xe7, 0x51, 0x31, 0x8c, 0xb9, 0x66, 0xc3, 0x48, 0xad,
	0xd8, 0x5a, 0xa5, 0x52, 0x93, 0x75, 0x2e, 0xb5, 0xc4, 0xa7, 0x2e, 0x44, 0x62, 0xa6, 0x38, 0xb1,
	0x49, 0x52, 0x25, 0xcf, 0xda, 0x89, 0x4c, 0xa4, 0x4d, 0x45, 0x46, 0xb9, 0x42, 0xf7, 0x15, 0x41,
	0x73, 0x56, 0x31, 0xa6, 0x9a, 0x0b, 0x7c, 0x0b, 0x75, 0xdb, 0x0b, 0x50, 0x07, 0xf5, 0x8f, 0xae,
	0xaf, 0xc8, 0xaf, 0x44, 0xb2, 0xef, 0xcd, 0xcc, 0xab, 0x29, 0x4f, 0x3c, 0xea, 0xca, 0xf8, 0x0e,
	0xfc, 0x8c, 0x15, 0xcb, 0xe0, 0x9f, 0x85, 0x5c, 0xfe, 0x01, 0x32, 0x1d, 0x3d, 0xdc, 0x1b, 0xc6,
	0xf8, 0xa0, 0xfc, 0x38, 0xf7, 0x8d, 0x9b, 0x78, 0xd4, 0x42, 0xc6, 0x0d, 0xf0, 0x33, 0xcd, 0x45,
	0xb7, 0x07, 0xc7, 0x3f, 0xbe, 0xc4, 0x18, 0xfc, 0x15, 0x13, 0xee, 0xdc, 0x43, 0x6a, 0x75, 0x77,
	0x09, 0xad, 0xef, 0x58, 0xdc, 0x82, 0xda, 0x82, 0x3f, 0xdb, 0x58, 0x93, 0x1a, 0x89, 0xdb, 0x50,
	0x2f, 0xd8, 0x72, 0xc3, 0xed, 0x91, 0x4d, 0xea, 0x0c, 0x0e, 0xe0, 0x7f, 0xc1, 0x73, 0x95, 0xc9,
	0x55, 0x50, 0xeb, 0xa0, 0x7e, 0x8d, 0xee, 0x2d, 0x3e, 0x81, 0x46, 0xca, 0xb3, 0x24, 0xd5, 0x81,
	0xdf, 0x41, 0xfd, 0x3a, 0xad, 0xdc, 0x78, 0xf4, 0x56, 0x86, 0x68, 0x5b, 0x86, 0xe8, 0xb3, 0x0c,
	0xd1, 0xcb, 0x2e, 0xf4, 0xb6, 0xbb, 0xd0, 0x7b, 0xdf, 0x85, 0xde, 0x63, 0x2f, 0xc9, 0x74, 0xba,
	0x89, 0xc9, 0x5c, 0x8a, 0x68, 0xb3, 0xd6, 0x4a, 0xb0, 0x5c, 0x1b, 0x91, 0xcd, 0x17, 0x03, 0xf5,
	0xb4, 0x18, 0x24, 0xd2, 0xad, 0x19, 0x37, 0xec, 0x18, 0x37, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x1c, 0x20, 0x7b, 0xf5, 0xe4, 0x01, 0x00, 0x00,
}

func (m *SnapshotItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Item != nil {
		{
			size := m.Item.Size()
			i -= size
			if _, err := m.Item.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotItem_Store) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem_Store) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Store != nil {
		{
			size, err := m.Store.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSnapshot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *SnapshotItem_IAVL) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotItem_IAVL) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.IAVL != nil {
		{
			size, err := m.IAVL.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSnapshot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *SnapshotStoreItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotStoreItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotStoreItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SnapshotIAVLItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SnapshotIAVLItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SnapshotIAVLItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x20
	}
	if m.Version != 0 {
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintSnapshot(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	offset -= sovSnapshot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SnapshotItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Item != nil {
		n += m.Item.Size()
	}
	return n
}

func (m *SnapshotItem_Store) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Store != nil {
		l = m.Store.Size()
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}
func (m *SnapshotItem_IAVL) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.IAVL != nil {
		l = m.IAVL.Size()
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}
func (m *SnapshotStoreItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	return n
}

func (m *SnapshotIAVLItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovSnapshot(uint64(l))
	}
	if m.Version != 0 {
		n += 1 + sovSnapshot(uint64(m.Version))
	}
	if m.Height != 0 {
		n += 1 + sovSnapshot(uint64(m.Height))
	}
	return n
}

func sovSnapshot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SnapshotItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SnapshotStoreItem{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &SnapshotItem_Store{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IAVL", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SnapshotIAVLItem{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Item = &SnapshotItem_IAVL{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *SnapshotStoreItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotStoreItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotStoreItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *SnapshotIAVLItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: SnapshotIAVLItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SnapshotIAVLItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSnapshot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = append(m.Value[:0], dAtA[iNdEx:postIndex]...)
			if m.Value == nil {
				m.Value = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSnapshot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSnapshot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSnapshot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSnapshot = fmt.Errorf("proto: unexpected end of group")
)
