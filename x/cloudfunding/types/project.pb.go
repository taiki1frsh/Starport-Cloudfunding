// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cloudfunding/project.proto

package types

import (
	fmt "fmt"
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

type Project struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Target      string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Collected   string `protobuf:"bytes,3,opt,name=collected,proto3" json:"collected,omitempty"`
	Deadline    string `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	State       string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	Creator     string `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_a705c5ea7860e2be, []int{0}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Project.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return m.Size()
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Project) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Project) GetCollected() string {
	if m != nil {
		return m.Collected
	}
	return ""
}

func (m *Project) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Project) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*Project)(nil), "taikifuru.cloudfunding.cloudfunding.Project")
}

func init() { proto.RegisterFile("cloudfunding/project.proto", fileDescriptor_a705c5ea7860e2be) }

var fileDescriptor_a705c5ea7860e2be = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe3, 0xd0, 0x26, 0xd4, 0x48, 0x0c, 0x16, 0x42, 0x56, 0x85, 0xac, 0x08, 0x96, 0x2e,
	0x24, 0x03, 0x03, 0x3b, 0x4f, 0x00, 0x1d, 0xd9, 0x52, 0xfb, 0x1a, 0x0e, 0x42, 0x1c, 0x39, 0x67,
	0x09, 0xde, 0x82, 0x17, 0x62, 0x67, 0xec, 0xc8, 0x88, 0x92, 0x17, 0x41, 0xb8, 0x40, 0x9b, 0xf1,
	0xfb, 0xbf, 0xbb, 0xe5, 0xe3, 0x73, 0x5d, 0x5b, 0x6f, 0xd6, 0xbe, 0x31, 0xd8, 0x54, 0x45, 0xeb,
	0xec, 0x23, 0x68, 0xca, 0x5b, 0x67, 0xc9, 0x8a, 0x0b, 0x2a, 0xf1, 0x09, 0xd7, 0xde, 0xf9, 0x7c,
	0xff, 0x6a, 0x04, 0xe7, 0xef, 0x8c, 0xa7, 0xb7, 0xdb, 0x37, 0x71, 0xcc, 0x63, 0x34, 0x92, 0x65,
	0x6c, 0x31, 0x59, 0xc6, 0x68, 0xc4, 0x29, 0x4f, 0xa8, 0x74, 0x15, 0x90, 0x8c, 0x33, 0xb6, 0x98,
	0x2d, 0x7f, 0x49, 0x9c, 0xf1, 0x99, 0xb6, 0x75, 0x0d, 0x9a, 0xc0, 0xc8, 0x83, 0xa0, 0x76, 0x83,
	0x98, 0xf3, 0x43, 0x03, 0xa5, 0xa9, 0xb1, 0x01, 0x39, 0x09, 0xf2, 0x9f, 0x45, 0xc6, 0x8f, 0x0c,
	0x74, 0xda, 0x61, 0x4b, 0x68, 0x1b, 0x39, 0x0d, 0x7a, 0x7f, 0x12, 0x27, 0x7c, 0xda, 0x51, 0x49,
	0x20, 0x93, 0xe0, 0xb6, 0x20, 0x24, 0x4f, 0xb5, 0x83, 0x92, 0xac, 0x93, 0x69, 0xd8, 0xff, 0xf0,
	0xe6, 0xee, 0xa3, 0x57, 0x6c, 0xd3, 0x2b, 0xf6, 0xd5, 0x2b, 0xf6, 0x36, 0xa8, 0x68, 0x33, 0xa8,
	0xe8, 0x73, 0x50, 0xd1, 0xfd, 0x75, 0x85, 0xf4, 0xe0, 0x57, 0xb9, 0xb6, 0xcf, 0x45, 0x28, 0x71,
	0xf9, 0x93, 0xa2, 0x18, 0x05, 0x7b, 0x19, 0x23, 0xbd, 0xb6, 0xd0, 0xad, 0x92, 0x90, 0xef, 0xea,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x00, 0xdc, 0x08, 0xbf, 0x5c, 0x01, 0x00, 0x00,
}

func (m *Project) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Project) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Project) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintProject(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Collected) > 0 {
		i -= len(m.Collected)
		copy(dAtA[i:], m.Collected)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Collected)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintProject(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintProject(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProject(dAtA []byte, offset int, v uint64) int {
	offset -= sovProject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Project) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovProject(uint64(m.Id))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Collected)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProject(uint64(l))
	}
	return n
}

func sovProject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProject(x uint64) (n int) {
	return sovProject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Project) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProject
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
			return fmt.Errorf("proto: Project: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Project: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collected", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collected = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProject
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
				return ErrInvalidLengthProject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProject
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
func skipProject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProject
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
					return 0, ErrIntOverflowProject
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
					return 0, ErrIntOverflowProject
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
				return 0, ErrInvalidLengthProject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProject = fmt.Errorf("proto: unexpected end of group")
)