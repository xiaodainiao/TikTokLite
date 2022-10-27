// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publish.proto

package message

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

type DouyinPublishActionRequest struct {
	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Data   []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DouyinPublishActionRequest) Reset()         { *m = DouyinPublishActionRequest{} }
func (m *DouyinPublishActionRequest) String() string { return proto.CompactTextString(m) }
func (*DouyinPublishActionRequest) ProtoMessage()    {}
func (*DouyinPublishActionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{0}
}
func (m *DouyinPublishActionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DouyinPublishActionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DouyinPublishActionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DouyinPublishActionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DouyinPublishActionRequest.Merge(m, src)
}
func (m *DouyinPublishActionRequest) XXX_Size() int {
	return m.Size()
}
func (m *DouyinPublishActionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DouyinPublishActionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DouyinPublishActionRequest proto.InternalMessageInfo

func (m *DouyinPublishActionRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DouyinPublishActionRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DouyinPublishActionRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type DouyinPublishActionResponse struct {
	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (m *DouyinPublishActionResponse) Reset()         { *m = DouyinPublishActionResponse{} }
func (m *DouyinPublishActionResponse) String() string { return proto.CompactTextString(m) }
func (*DouyinPublishActionResponse) ProtoMessage()    {}
func (*DouyinPublishActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{1}
}
func (m *DouyinPublishActionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DouyinPublishActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DouyinPublishActionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DouyinPublishActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DouyinPublishActionResponse.Merge(m, src)
}
func (m *DouyinPublishActionResponse) XXX_Size() int {
	return m.Size()
}
func (m *DouyinPublishActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DouyinPublishActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DouyinPublishActionResponse proto.InternalMessageInfo

func (m *DouyinPublishActionResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DouyinPublishActionResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

type DouyinPublishListRequest struct {
	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *DouyinPublishListRequest) Reset()         { *m = DouyinPublishListRequest{} }
func (m *DouyinPublishListRequest) String() string { return proto.CompactTextString(m) }
func (*DouyinPublishListRequest) ProtoMessage()    {}
func (*DouyinPublishListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{2}
}
func (m *DouyinPublishListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DouyinPublishListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DouyinPublishListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DouyinPublishListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DouyinPublishListRequest.Merge(m, src)
}
func (m *DouyinPublishListRequest) XXX_Size() int {
	return m.Size()
}
func (m *DouyinPublishListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DouyinPublishListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DouyinPublishListRequest proto.InternalMessageInfo

func (m *DouyinPublishListRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DouyinPublishListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DouyinPublishListResponse struct {
	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

func (m *DouyinPublishListResponse) Reset()         { *m = DouyinPublishListResponse{} }
func (m *DouyinPublishListResponse) String() string { return proto.CompactTextString(m) }
func (*DouyinPublishListResponse) ProtoMessage()    {}
func (*DouyinPublishListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_34180b7635741fb2, []int{3}
}
func (m *DouyinPublishListResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DouyinPublishListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DouyinPublishListResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DouyinPublishListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DouyinPublishListResponse.Merge(m, src)
}
func (m *DouyinPublishListResponse) XXX_Size() int {
	return m.Size()
}
func (m *DouyinPublishListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DouyinPublishListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DouyinPublishListResponse proto.InternalMessageInfo

func (m *DouyinPublishListResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *DouyinPublishListResponse) GetStatusMsg() string {
	if m != nil {
		return m.StatusMsg
	}
	return ""
}

func (m *DouyinPublishListResponse) GetVideoList() []*Video {
	if m != nil {
		return m.VideoList
	}
	return nil
}

func init() {
	proto.RegisterType((*DouyinPublishActionRequest)(nil), "douyin.core.douyin_publish_action_request")
	proto.RegisterType((*DouyinPublishActionResponse)(nil), "douyin.core.douyin_publish_action_response")
	proto.RegisterType((*DouyinPublishListRequest)(nil), "douyin.core.douyin_publish_list_request")
	proto.RegisterType((*DouyinPublishListResponse)(nil), "douyin.core.douyin_publish_list_response")
}

func init() { proto.RegisterFile("publish.proto", fileDescriptor_34180b7635741fb2) }

var fileDescriptor_34180b7635741fb2 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x6b, 0x42, 0x8b, 0x7a, 0x03, 0x0c, 0x16, 0x12, 0x11, 0x50, 0x13, 0x65, 0xca, 0x64,
	0x04, 0x8c, 0x6c, 0x30, 0x21, 0x95, 0x25, 0x03, 0x03, 0x8b, 0x49, 0x62, 0x2b, 0x58, 0x6d, 0xe3,
	0x90, 0xeb, 0x20, 0xf1, 0x16, 0xf0, 0x56, 0x8c, 0x1d, 0x19, 0x51, 0xf2, 0x22, 0x28, 0x3f, 0x03,
	0x42, 0xb0, 0x20, 0x36, 0x9f, 0x7b, 0xae, 0xcf, 0x67, 0xeb, 0xc0, 0x4e, 0x51, 0x25, 0x4b, 0x8d,
	0x0f, 0xbc, 0x28, 0x8d, 0x35, 0xd4, 0x95, 0xa6, 0x7a, 0xd6, 0x39, 0x4f, 0x4d, 0xa9, 0x0e, 0xa0,
	0x42, 0x55, 0xf6, 0x46, 0x90, 0xc0, 0xac, 0xb7, 0xc4, 0x70, 0x41, 0xc4, 0xa9, 0xd5, 0x26, 0x17,
	0xa5, 0x7a, 0xac, 0x14, 0x5a, 0xba, 0x0f, 0x5b, 0xed, 0xba, 0xd0, 0xd2, 0x23, 0x3e, 0x09, 0x9d,
	0x68, 0xd2, 0xca, 0x6b, 0x49, 0xf7, 0x60, 0x6c, 0xcd, 0x42, 0xe5, 0xde, 0x86, 0x4f, 0xc2, 0x69,
	0xd4, 0x0b, 0x4a, 0x61, 0x53, 0xc6, 0x36, 0xf6, 0x1c, 0x9f, 0x84, 0xdb, 0x51, 0x77, 0x0e, 0xee,
	0x81, 0xfd, 0xc6, 0xc0, 0xc2, 0xe4, 0xa8, 0xe8, 0x31, 0xb8, 0x68, 0x63, 0x5b, 0xa1, 0x48, 0x8d,
	0x54, 0x1d, 0x68, 0x1c, 0x41, 0x3f, 0xba, 0x32, 0x52, 0xd1, 0x19, 0x0c, 0x4a, 0xac, 0x30, 0x1b,
	0x88, 0xd3, 0x7e, 0x72, 0x83, 0x59, 0x30, 0x87, 0xc3, 0x6f, 0x84, 0xa5, 0x46, 0xfb, 0xc7, 0x3f,
	0x04, 0xaf, 0x04, 0x8e, 0x7e, 0x8e, 0xfb, 0x9f, 0xe7, 0xd2, 0x53, 0x80, 0x27, 0x2d, 0x95, 0xe9,
	0x62, 0x3d, 0xc7, 0x77, 0x42, 0xf7, 0x8c, 0xf2, 0x2f, 0x15, 0xf1, 0xdb, 0xd6, 0x8e, 0xa6, 0xdd,
	0xd6, 0x5c, 0xa3, 0xbd, 0x0c, 0xdf, 0x6a, 0x46, 0xd6, 0x35, 0x23, 0x1f, 0x35, 0x23, 0x2f, 0x0d,
	0x1b, 0xad, 0x1b, 0x36, 0x7a, 0x6f, 0xd8, 0xe8, 0x6e, 0x97, 0xf3, 0x93, 0x62, 0x91, 0x5d, 0xac,
	0x14, 0x62, 0x9c, 0xa9, 0x64, 0xd2, 0x15, 0x7b, 0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xd8,
	0xff, 0xa3, 0x02, 0x02, 0x00, 0x00,
}

func (m *DouyinPublishActionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DouyinPublishActionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DouyinPublishActionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintPublish(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintPublish(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserId != 0 {
		i = encodeVarintPublish(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DouyinPublishActionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DouyinPublishActionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DouyinPublishActionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StatusMsg) > 0 {
		i -= len(m.StatusMsg)
		copy(dAtA[i:], m.StatusMsg)
		i = encodeVarintPublish(dAtA, i, uint64(len(m.StatusMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.StatusCode != 0 {
		i = encodeVarintPublish(dAtA, i, uint64(m.StatusCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DouyinPublishListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DouyinPublishListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DouyinPublishListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintPublish(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if m.UserId != 0 {
		i = encodeVarintPublish(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DouyinPublishListResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DouyinPublishListResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DouyinPublishListResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VideoList) > 0 {
		for iNdEx := len(m.VideoList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VideoList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPublish(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StatusMsg) > 0 {
		i -= len(m.StatusMsg)
		copy(dAtA[i:], m.StatusMsg)
		i = encodeVarintPublish(dAtA, i, uint64(len(m.StatusMsg)))
		i--
		dAtA[i] = 0x12
	}
	if m.StatusCode != 0 {
		i = encodeVarintPublish(dAtA, i, uint64(m.StatusCode))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPublish(dAtA []byte, offset int, v uint64) int {
	offset -= sovPublish(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DouyinPublishActionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovPublish(uint64(m.UserId))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovPublish(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovPublish(uint64(l))
	}
	return n
}

func (m *DouyinPublishActionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StatusCode != 0 {
		n += 1 + sovPublish(uint64(m.StatusCode))
	}
	l = len(m.StatusMsg)
	if l > 0 {
		n += 1 + l + sovPublish(uint64(l))
	}
	return n
}

func (m *DouyinPublishListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserId != 0 {
		n += 1 + sovPublish(uint64(m.UserId))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovPublish(uint64(l))
	}
	return n
}

func (m *DouyinPublishListResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StatusCode != 0 {
		n += 1 + sovPublish(uint64(m.StatusCode))
	}
	l = len(m.StatusMsg)
	if l > 0 {
		n += 1 + l + sovPublish(uint64(l))
	}
	if len(m.VideoList) > 0 {
		for _, e := range m.VideoList {
			l = e.Size()
			n += 1 + l + sovPublish(uint64(l))
		}
	}
	return n
}

func sovPublish(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPublish(x uint64) (n int) {
	return sovPublish(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DouyinPublishActionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublish
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
			return fmt.Errorf("proto: douyin_publish_action_request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: douyin_publish_action_request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
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
					return ErrIntOverflowPublish
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
				return ErrInvalidLengthPublish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
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
				return ErrInvalidLengthPublish
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPublish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublish(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublish
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
func (m *DouyinPublishActionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublish
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
			return fmt.Errorf("proto: douyin_publish_action_response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: douyin_publish_action_response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusCode", wireType)
			}
			m.StatusCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
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
				return ErrInvalidLengthPublish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublish(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublish
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
func (m *DouyinPublishListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublish
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
			return fmt.Errorf("proto: douyin_publish_list_request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: douyin_publish_list_request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
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
					return ErrIntOverflowPublish
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
				return ErrInvalidLengthPublish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublish(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublish
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
func (m *DouyinPublishListResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPublish
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
			return fmt.Errorf("proto: douyin_publish_list_response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: douyin_publish_list_response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusCode", wireType)
			}
			m.StatusCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatusCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatusMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
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
				return ErrInvalidLengthPublish
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPublish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatusMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VideoList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPublish
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
				return ErrInvalidLengthPublish
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPublish
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VideoList = append(m.VideoList, &Video{})
			if err := m.VideoList[len(m.VideoList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPublish(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPublish
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
func skipPublish(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPublish
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
					return 0, ErrIntOverflowPublish
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
					return 0, ErrIntOverflowPublish
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
				return 0, ErrInvalidLengthPublish
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPublish
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPublish
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPublish        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPublish          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPublish = fmt.Errorf("proto: unexpected end of group")
)