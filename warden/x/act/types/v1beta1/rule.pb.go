// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/act/v1beta1/rule.proto

package v1beta1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	proto "github.com/cosmos/gogoproto/proto"
	ast "github.com/warden-protocol/wardenprotocol/shield/ast"
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

type Rule struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The expression to be evaluated for this rule.
	Expression *ast.Expression `protobuf:"bytes,4,opt,name=expression,proto3" json:"expression,omitempty"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ebeb86c5a0f3ba, []int{0}
}
func (m *Rule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return m.Size()
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rule) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Rule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Rule) GetExpression() *ast.Expression {
	if m != nil {
		return m.Expression
	}
	return nil
}

func init() {
	proto.RegisterType((*Rule)(nil), "warden.act.v1beta1.Rule")
}

func init() { proto.RegisterFile("warden/act/v1beta1/rule.proto", fileDescriptor_53ebeb86c5a0f3ba) }

var fileDescriptor_53ebeb86c5a0f3ba = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x41, 0x4a, 0xc4, 0x30,
	0x18, 0x85, 0x9b, 0x5a, 0x14, 0x23, 0xb8, 0x08, 0x22, 0x75, 0xc0, 0x50, 0x5c, 0x75, 0x63, 0xc2,
	0x28, 0x78, 0x00, 0xd1, 0x0b, 0x74, 0xe9, 0x42, 0x48, 0xd3, 0xdf, 0x4e, 0x20, 0xd3, 0x94, 0x24,
	0xd5, 0x19, 0xf0, 0x10, 0x1e, 0xcb, 0xe5, 0x2c, 0x5d, 0x4a, 0x7b, 0x11, 0x99, 0x64, 0x46, 0xc4,
	0x45, 0x20, 0xef, 0xbd, 0x0f, 0x5e, 0xf2, 0xf0, 0xe5, 0x9b, 0xb0, 0x0d, 0x74, 0x5c, 0x48, 0xcf,
	0x5f, 0xe7, 0x35, 0x78, 0x31, 0xe7, 0x76, 0xd0, 0xc0, 0x7a, 0x6b, 0xbc, 0x21, 0x24, 0xc6, 0x4c,
	0x48, 0xcf, 0x76, 0xf1, 0xec, 0xa2, 0x35, 0xa6, 0xd5, 0xc0, 0x03, 0x51, 0x0f, 0x2f, 0x5c, 0x74,
	0xeb, 0x88, 0xcf, 0xce, 0xdc, 0x42, 0x81, 0x6e, 0xb8, 0x70, 0x7e, 0x7b, 0xa2, 0x7b, 0xf5, 0x8e,
	0xb3, 0x6a, 0xd0, 0x40, 0x4e, 0x71, 0xaa, 0x9a, 0x1c, 0x15, 0xa8, 0xcc, 0xaa, 0x54, 0x35, 0x24,
	0xc7, 0x47, 0xd2, 0x82, 0xf0, 0xc6, 0xe6, 0x69, 0x81, 0xca, 0xe3, 0x6a, 0x2f, 0x09, 0xc1, 0x59,
	0x27, 0x96, 0x90, 0x1f, 0x04, 0x3b, 0xdc, 0xc9, 0x1d, 0xc6, 0xb0, 0xea, 0x2d, 0x38, 0xa7, 0x4c,
	0x97, 0x67, 0x05, 0x2a, 0x4f, 0x6e, 0xce, 0x59, 0x2c, 0x64, 0xdb, 0xb2, 0xc7, 0xdf, 0xb4, 0xfa,
	0x43, 0xde, 0x3f, 0x7f, 0x8e, 0x14, 0x6d, 0x46, 0x8a, 0xbe, 0x47, 0x8a, 0x3e, 0x26, 0x9a, 0x6c,
	0x26, 0x9a, 0x7c, 0x4d, 0x34, 0x79, 0x7a, 0x68, 0x95, 0x5f, 0x0c, 0x35, 0x93, 0x66, 0xc9, 0xe3,
	0x3f, 0xaf, 0xc3, 0x83, 0xa5, 0xd1, 0x3b, 0xfd, 0x4f, 0xf2, 0x55, 0xd8, 0xc9, 0xaf, 0x7b, 0x70,
	0xfb, 0xb5, 0xea, 0xc3, 0x00, 0xdd, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x48, 0x45, 0x54, 0x9f,
	0x4a, 0x01, 0x00, 0x00,
}

func (m *Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expression != nil {
		{
			size, err := m.Expression.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRule(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRule(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRule(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintRule(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRule(dAtA []byte, offset int, v uint64) int {
	offset -= sovRule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Rule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRule(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRule(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRule(uint64(l))
	}
	if m.Expression != nil {
		l = m.Expression.Size()
		n += 1 + l + sovRule(uint64(l))
	}
	return n
}

func sovRule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRule(x uint64) (n int) {
	return sovRule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRule
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
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRule
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
				return ErrInvalidLengthRule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRule
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
				return ErrInvalidLengthRule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expression", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRule
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
				return ErrInvalidLengthRule
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expression == nil {
				m.Expression = &ast.Expression{}
			}
			if err := m.Expression.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRule
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
func skipRule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRule
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
					return 0, ErrIntOverflowRule
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
					return 0, ErrIntOverflowRule
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
				return 0, ErrInvalidLengthRule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRule = fmt.Errorf("proto: unexpected end of group")
)
