// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/claim/v1beta1/claim.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type Action int32

const (
	ActionAddLiquidity  Action = 0
	ActionSwap          Action = 1
	ActionVote          Action = 2
	ActionDelegateStake Action = 3
)

var Action_name = map[int32]string{
	0: "ActionAddLiquidity",
	1: "ActionSwap",
	2: "ActionVote",
	3: "ActionDelegateStake",
}

var Action_value = map[string]int32{
	"ActionAddLiquidity":  0,
	"ActionSwap":          1,
	"ActionVote":          2,
	"ActionDelegateStake": 3,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a48687913a68f9e3, []int{0}
}

// A Claim Records is the metadata of claim data per address
type ClaimRecord struct {
	// address of claim user
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial claimable amount for the user
	InitialClaimableAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_claimable_amount,json=initialClaimableAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_claimable_amount" yaml:"initial_claimable_amount"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,3,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *ClaimRecord) Reset()         { *m = ClaimRecord{} }
func (m *ClaimRecord) String() string { return proto.CompactTextString(m) }
func (*ClaimRecord) ProtoMessage()    {}
func (*ClaimRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_a48687913a68f9e3, []int{0}
}
func (m *ClaimRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRecord.Merge(m, src)
}
func (m *ClaimRecord) XXX_Size() int {
	return m.Size()
}
func (m *ClaimRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRecord.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRecord proto.InternalMessageInfo

func (m *ClaimRecord) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ClaimRecord) GetInitialClaimableAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialClaimableAmount
	}
	return nil
}

func (m *ClaimRecord) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterEnum("osmosis.claim.v1beta1.Action", Action_name, Action_value)
	proto.RegisterType((*ClaimRecord)(nil), "osmosis.claim.v1beta1.ClaimRecord")
}

func init() { proto.RegisterFile("osmosis/claim/v1beta1/claim.proto", fileDescriptor_a48687913a68f9e3) }

var fileDescriptor_a48687913a68f9e3 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x6e, 0xd4, 0x30,
	0x18, 0xc7, 0x93, 0x3b, 0x54, 0xc0, 0x95, 0x4a, 0x64, 0xa0, 0x0d, 0x37, 0x38, 0xd7, 0x4c, 0x11,
	0xa2, 0x31, 0x05, 0x26, 0xb6, 0x4b, 0x98, 0x2a, 0xa6, 0x9c, 0xc4, 0xc0, 0x12, 0x39, 0xb1, 0x15,
	0xac, 0x4b, 0xe2, 0x70, 0xf6, 0x15, 0xee, 0x0d, 0x18, 0x79, 0x07, 0x16, 0xc4, 0xc8, 0x53, 0x74,
	0xec, 0xc8, 0x74, 0xa0, 0xbb, 0x37, 0xe8, 0x13, 0xa0, 0xd8, 0x4e, 0x8b, 0x90, 0x3a, 0xc5, 0xdf,
	0xff, 0xfb, 0xe9, 0xef, 0xff, 0xe7, 0x7c, 0xe0, 0x58, 0xc8, 0x46, 0x48, 0x2e, 0x71, 0x59, 0x13,
	0xde, 0xe0, 0xf3, 0xd3, 0x82, 0x29, 0x72, 0x6a, 0xaa, 0xb8, 0x5b, 0x0a, 0x25, 0xe0, 0x63, 0x8b,
	0xc4, 0x46, 0xb4, 0xc8, 0xe4, 0x51, 0x25, 0x2a, 0xa1, 0x09, 0xdc, 0x9f, 0x0c, 0x3c, 0x41, 0xa5,
	0xa6, 0x71, 0x41, 0x24, 0xbb, 0x71, 0x13, 0xbc, 0xb5, 0xfd, 0xe3, 0xeb, 0x7e, 0xbb, 0xb8, 0xee,
	0x57, 0xac, 0x65, 0xfd, 0x05, 0x1a, 0x09, 0x7f, 0x8e, 0xc0, 0x7e, 0xda, 0x5f, 0x95, 0xb1, 0x52,
	0x2c, 0x29, 0x7c, 0x06, 0xee, 0x12, 0x4a, 0x97, 0x4c, 0x4a, 0xdf, 0x9d, 0xba, 0xd1, 0xfd, 0x04,
	0x5e, 0x6d, 0x82, 0x83, 0x35, 0x69, 0xea, 0xd7, 0xa1, 0x6d, 0x84, 0xd9, 0x80, 0xc0, 0xef, 0x2e,
	0xf0, 0x79, 0xcb, 0x15, 0x27, 0x75, 0xae, 0x03, 0x93, 0xa2, 0x66, 0x39, 0x69, 0xc4, 0xaa, 0x55,
	0xfe, 0x68, 0x3a, 0x8e, 0xf6, 0x5f, 0x3c, 0x89, 0x4d, 0x88, 0xb8, 0x0f, 0x39, 0xcc, 0x13, 0xa7,
	0x82, 0xb7, 0xc9, 0xfc, 0x62, 0x13, 0x38, 0x57, 0x9b, 0x20, 0x30, 0xf6, 0xb7, 0x19, 0x85, 0x3f,
	0x7e, 0x07, 0x51, 0xc5, 0xd5, 0x87, 0x55, 0x11, 0x97, 0xa2, 0xc1, 0x76, 0x28, 0xf3, 0x39, 0x91,
	0x74, 0x81, 0xd5, 0xba, 0x63, 0x52, 0x7b, 0xca, 0xec, 0xd0, 0xda, 0xa4, 0x83, 0xcb, 0x4c, 0x9b,
	0xc0, 0x33, 0xe0, 0x91, 0x52, 0x71, 0xd1, 0xe6, 0xa5, 0x68, 0xba, 0x9a, 0x29, 0x46, 0xfd, 0xf1,
	0x74, 0x1c, 0xdd, 0x4b, 0x02, 0x1b, 0xe3, 0xc8, 0x4e, 0xf9, 0x1f, 0x15, 0x66, 0x0f, 0x8c, 0x94,
	0x0e, 0xca, 0xd3, 0x1c, 0xec, 0xcd, 0xb4, 0x04, 0x0f, 0x01, 0x34, 0xa7, 0x19, 0xa5, 0x6f, 0xf9,
	0xc7, 0x15, 0xa7, 0x5c, 0xad, 0x3d, 0x07, 0x1e, 0x00, 0x60, 0xf4, 0xf9, 0x27, 0xd2, 0x79, 0xee,
	0x4d, 0xfd, 0x4e, 0x28, 0xe6, 0x8d, 0xe0, 0x11, 0x78, 0x68, 0xea, 0x37, 0xac, 0x66, 0x15, 0x51,
	0x6c, 0xae, 0xc8, 0x82, 0x79, 0xe3, 0xc9, 0x9d, 0x2f, 0xdf, 0x90, 0x93, 0x9c, 0x5d, 0x6c, 0x91,
	0x7b, 0xb9, 0x45, 0xee, 0x9f, 0x2d, 0x72, 0xbf, 0xee, 0x90, 0x73, 0xb9, 0x43, 0xce, 0xaf, 0x1d,
	0x72, 0xde, 0x3f, 0xff, 0xe7, 0x21, 0xec, 0xaa, 0x9c, 0xd4, 0xa4, 0x90, 0x43, 0x81, 0xcf, 0x5f,
	0xe1, 0xcf, 0x76, 0xbf, 0xf4, 0xb3, 0x14, 0x7b, 0xfa, 0x47, 0xbf, 0xfc, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x95, 0x8b, 0x4e, 0x6d, 0x7d, 0x02, 0x00, 0x00,
}

func (m *ClaimRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintClaim(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InitialClaimableAmount) > 0 {
		for iNdEx := len(m.InitialClaimableAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialClaimableAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaim(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClaimRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if len(m.InitialClaimableAmount) > 0 {
		for _, e := range m.InitialClaimableAmount {
			l = e.Size()
			n += 1 + l + sovClaim(uint64(l))
		}
	}
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovClaim(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClaimRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: ClaimRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialClaimableAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialClaimableAmount = append(m.InitialClaimableAmount, types.Coin{})
			if err := m.InitialClaimableAmount[len(m.InitialClaimableAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowClaim
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthClaim
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthClaim
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowClaim
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
