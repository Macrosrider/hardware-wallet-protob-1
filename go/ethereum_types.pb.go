// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethereum_types.proto

package messages

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// *
// Structure representing Ethereum transaction input
type EthereumTransactionInput struct {
	AddressN             *uint32  `protobuf:"varint,1,req,name=address_n,json=addressN" json:"address_n,omitempty"`
	Index                *uint32  `protobuf:"varint,2,req,name=index" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumTransactionInput) Reset()         { *m = EthereumTransactionInput{} }
func (m *EthereumTransactionInput) String() string { return proto.CompactTextString(m) }
func (*EthereumTransactionInput) ProtoMessage()    {}
func (*EthereumTransactionInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethereum_types_3608c9324be5e1d3, []int{0}
}
func (m *EthereumTransactionInput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthereumTransactionInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthereumTransactionInput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EthereumTransactionInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumTransactionInput.Merge(dst, src)
}
func (m *EthereumTransactionInput) XXX_Size() int {
	return m.Size()
}
func (m *EthereumTransactionInput) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumTransactionInput.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumTransactionInput proto.InternalMessageInfo

func (m *EthereumTransactionInput) GetAddressN() uint32 {
	if m != nil && m.AddressN != nil {
		return *m.AddressN
	}
	return 0
}

func (m *EthereumTransactionInput) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

// *
// Structure representing transaction output
type EthereumTransactionOutput struct {
	Address              *string  `protobuf:"bytes,1,req,name=address" json:"address,omitempty"`
	Coin                 *uint64  `protobuf:"varint,2,req,name=coin" json:"coin,omitempty"`
	AddressIndex         *uint32  `protobuf:"varint,3,opt,name=address_index,json=addressIndex" json:"address_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumTransactionOutput) Reset()         { *m = EthereumTransactionOutput{} }
func (m *EthereumTransactionOutput) String() string { return proto.CompactTextString(m) }
func (*EthereumTransactionOutput) ProtoMessage()    {}
func (*EthereumTransactionOutput) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethereum_types_3608c9324be5e1d3, []int{1}
}
func (m *EthereumTransactionOutput) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthereumTransactionOutput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthereumTransactionOutput.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EthereumTransactionOutput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumTransactionOutput.Merge(dst, src)
}
func (m *EthereumTransactionOutput) XXX_Size() int {
	return m.Size()
}
func (m *EthereumTransactionOutput) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumTransactionOutput.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumTransactionOutput proto.InternalMessageInfo

func (m *EthereumTransactionOutput) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *EthereumTransactionOutput) GetCoin() uint64 {
	if m != nil && m.Coin != nil {
		return *m.Coin
	}
	return 0
}

func (m *EthereumTransactionOutput) GetAddressIndex() uint32 {
	if m != nil && m.AddressIndex != nil {
		return *m.AddressIndex
	}
	return 0
}

// *
// Structure representing transaction
type EthereumTransactionType struct {
	InputsCnt            *uint32                      `protobuf:"varint,2,opt,name=inputs_cnt,json=inputsCnt" json:"inputs_cnt,omitempty"`
	Inputs               []*EthereumTransactionInput  `protobuf:"bytes,3,rep,name=inputs" json:"inputs,omitempty"`
	OutputsCnt           *uint32                      `protobuf:"varint,4,opt,name=outputs_cnt,json=outputsCnt" json:"outputs_cnt,omitempty"`
	Outputs              []*EthereumTransactionOutput `protobuf:"bytes,5,rep,name=outputs" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *EthereumTransactionType) Reset()         { *m = EthereumTransactionType{} }
func (m *EthereumTransactionType) String() string { return proto.CompactTextString(m) }
func (*EthereumTransactionType) ProtoMessage()    {}
func (*EthereumTransactionType) Descriptor() ([]byte, []int) {
	return fileDescriptor_ethereum_types_3608c9324be5e1d3, []int{2}
}
func (m *EthereumTransactionType) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthereumTransactionType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthereumTransactionType.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *EthereumTransactionType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumTransactionType.Merge(dst, src)
}
func (m *EthereumTransactionType) XXX_Size() int {
	return m.Size()
}
func (m *EthereumTransactionType) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumTransactionType.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumTransactionType proto.InternalMessageInfo

func (m *EthereumTransactionType) GetInputsCnt() uint32 {
	if m != nil && m.InputsCnt != nil {
		return *m.InputsCnt
	}
	return 0
}

func (m *EthereumTransactionType) GetInputs() []*EthereumTransactionInput {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *EthereumTransactionType) GetOutputsCnt() uint32 {
	if m != nil && m.OutputsCnt != nil {
		return *m.OutputsCnt
	}
	return 0
}

func (m *EthereumTransactionType) GetOutputs() []*EthereumTransactionOutput {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*EthereumTransactionInput)(nil), "EthereumTransactionInput")
	proto.RegisterType((*EthereumTransactionOutput)(nil), "EthereumTransactionOutput")
	proto.RegisterType((*EthereumTransactionType)(nil), "EthereumTransactionType")
}
func (m *EthereumTransactionInput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthereumTransactionInput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.AddressN == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("address_n")
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEthereumTypes(dAtA, i, uint64(*m.AddressN))
	}
	if m.Index == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("index")
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEthereumTypes(dAtA, i, uint64(*m.Index))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *EthereumTransactionOutput) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthereumTransactionOutput) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Address == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("address")
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEthereumTypes(dAtA, i, uint64(len(*m.Address)))
		i += copy(dAtA[i:], *m.Address)
	}
	if m.Coin == nil {
		return 0, github_com_gogo_protobuf_proto.NewRequiredNotSetError("coin")
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEthereumTypes(dAtA, i, uint64(*m.Coin))
	}
	if m.AddressIndex != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintEthereumTypes(dAtA, i, uint64(*m.AddressIndex))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *EthereumTransactionType) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthereumTransactionType) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.InputsCnt != nil {
		dAtA[i] = 0x10
		i++
		i = encodeVarintEthereumTypes(dAtA, i, uint64(*m.InputsCnt))
	}
	if len(m.Inputs) > 0 {
		for _, msg := range m.Inputs {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintEthereumTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.OutputsCnt != nil {
		dAtA[i] = 0x20
		i++
		i = encodeVarintEthereumTypes(dAtA, i, uint64(*m.OutputsCnt))
	}
	if len(m.Outputs) > 0 {
		for _, msg := range m.Outputs {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintEthereumTypes(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintEthereumTypes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EthereumTransactionInput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AddressN != nil {
		n += 1 + sovEthereumTypes(uint64(*m.AddressN))
	}
	if m.Index != nil {
		n += 1 + sovEthereumTypes(uint64(*m.Index))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EthereumTransactionOutput) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		l = len(*m.Address)
		n += 1 + l + sovEthereumTypes(uint64(l))
	}
	if m.Coin != nil {
		n += 1 + sovEthereumTypes(uint64(*m.Coin))
	}
	if m.AddressIndex != nil {
		n += 1 + sovEthereumTypes(uint64(*m.AddressIndex))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *EthereumTransactionType) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InputsCnt != nil {
		n += 1 + sovEthereumTypes(uint64(*m.InputsCnt))
	}
	if len(m.Inputs) > 0 {
		for _, e := range m.Inputs {
			l = e.Size()
			n += 1 + l + sovEthereumTypes(uint64(l))
		}
	}
	if m.OutputsCnt != nil {
		n += 1 + sovEthereumTypes(uint64(*m.OutputsCnt))
	}
	if len(m.Outputs) > 0 {
		for _, e := range m.Outputs {
			l = e.Size()
			n += 1 + l + sovEthereumTypes(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovEthereumTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEthereumTypes(x uint64) (n int) {
	return sovEthereumTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthereumTransactionInput) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthereumTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EthereumTransactionInput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthereumTransactionInput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressN", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddressN = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Index = &v
			hasFields[0] |= uint64(0x00000002)
		default:
			iNdEx = preIndex
			skippy, err := skipEthereumTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEthereumTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("address_n")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("index")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EthereumTransactionOutput) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthereumTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EthereumTransactionOutput: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthereumTransactionOutput: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEthereumTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Address = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Coin = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressIndex", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AddressIndex = &v
		default:
			iNdEx = preIndex
			skippy, err := skipEthereumTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEthereumTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("address")
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return github_com_gogo_protobuf_proto.NewRequiredNotSetError("coin")
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EthereumTransactionType) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthereumTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EthereumTransactionType: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthereumTransactionType: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputsCnt", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InputsCnt = &v
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEthereumTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inputs = append(m.Inputs, &EthereumTransactionInput{})
			if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputsCnt", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OutputsCnt = &v
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEthereumTypes
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Outputs = append(m.Outputs, &EthereumTransactionOutput{})
			if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEthereumTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEthereumTypes
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
func skipEthereumTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthereumTypes
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
					return 0, ErrIntOverflowEthereumTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEthereumTypes
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEthereumTypes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEthereumTypes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEthereumTypes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEthereumTypes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthereumTypes   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("ethereum_types.proto", fileDescriptor_ethereum_types_3608c9324be5e1d3)
}

var fileDescriptor_ethereum_types_3608c9324be5e1d3 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0x83, 0x40,
	0x10, 0x87, 0xb3, 0x85, 0xda, 0x76, 0x2a, 0xc6, 0x6c, 0x9a, 0xb8, 0xd5, 0x88, 0x04, 0x2f, 0x9c,
	0x48, 0x34, 0x3e, 0x41, 0x4d, 0x0f, 0xf5, 0xa0, 0xc9, 0xb6, 0x27, 0x2f, 0x04, 0x61, 0x55, 0x62,
	0xbb, 0x10, 0x76, 0x89, 0xf2, 0x7a, 0x9e, 0x3c, 0xfa, 0x08, 0x86, 0x27, 0x31, 0xfb, 0x87, 0x5b,
	0x7b, 0x82, 0xdf, 0xc7, 0xf0, 0xcd, 0xce, 0x2c, 0xcc, 0x98, 0x7c, 0x67, 0x35, 0x6b, 0x76, 0x89,
	0x6c, 0x2b, 0x26, 0xe2, 0xaa, 0x2e, 0x65, 0x19, 0xae, 0x81, 0x2c, 0x2d, 0xdf, 0xd4, 0x29, 0x17,
	0x69, 0x26, 0x8b, 0x92, 0xaf, 0x78, 0xd5, 0x48, 0x7c, 0x01, 0x93, 0x34, 0xcf, 0x6b, 0x26, 0x44,
	0xc2, 0x09, 0x0a, 0x06, 0x91, 0x47, 0xc7, 0x16, 0x3c, 0xe2, 0x19, 0x0c, 0x0b, 0x9e, 0xb3, 0x2f,
	0x32, 0xd0, 0x1f, 0x4c, 0x78, 0x70, 0xc7, 0xce, 0xe9, 0x49, 0xc8, 0x61, 0xbe, 0x47, 0xfa, 0xd4,
	0x48, 0x65, 0x25, 0x30, 0xb2, 0x12, 0xed, 0x9c, 0xd0, 0x3e, 0x62, 0x0c, 0x6e, 0x56, 0x16, 0x5c,
	0x1b, 0x5d, 0xaa, 0xdf, 0xf1, 0x35, 0x78, 0xfd, 0x19, 0x4c, 0x3b, 0x27, 0x40, 0x91, 0x47, 0x8f,
	0x2d, 0x5c, 0x29, 0x16, 0x7e, 0x23, 0x38, 0xdb, 0xd3, 0x70, 0xd3, 0x56, 0x0c, 0x5f, 0x02, 0x14,
	0x6a, 0x1a, 0x91, 0x64, 0x5c, 0x92, 0x81, 0xfe, 0x7b, 0x62, 0xc8, 0x3d, 0x97, 0xf8, 0x06, 0x8e,
	0x4c, 0x20, 0x4e, 0xe0, 0x44, 0xd3, 0xdb, 0x79, 0x7c, 0x68, 0x1d, 0xd4, 0x16, 0xe2, 0x2b, 0x98,
	0x96, 0x7a, 0x14, 0xa3, 0x74, 0xb5, 0x12, 0x2c, 0x52, 0xce, 0x3b, 0x18, 0xd9, 0x44, 0x86, 0x5a,
	0x7a, 0x1e, 0x1f, 0x5c, 0x07, 0xed, 0x4b, 0x17, 0xcb, 0x9f, 0xce, 0x47, 0xbf, 0x9d, 0x8f, 0xfe,
	0x3a, 0x1f, 0x81, 0xcf, 0x99, 0x8c, 0xc5, 0x47, 0xab, 0x96, 0xa0, 0x9e, 0x9f, 0xe9, 0x76, 0xcb,
	0xa4, 0xb9, 0xb4, 0x97, 0xe6, 0x75, 0xe1, 0xad, 0x7b, 0xa6, 0xa6, 0x7c, 0x1e, 0xef, 0x98, 0x10,
	0xe9, 0x1b, 0x13, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x3c, 0x96, 0xe2, 0xe7, 0x01, 0x00,
	0x00,
}
