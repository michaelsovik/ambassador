// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v3alpha/core/http_uri.proto

package envoy_api_v3alpha_core

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// Envoy external URI descriptor
type HttpUri struct {
	// The HTTP server URI. It should be a full FQDN with protocol, host and path.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//    uri: https://www.googleapis.com/oauth2/v1/certs
	//
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Specify how `uri` is to be fetched. Today, this requires an explicit
	// cluster, but in the future we may support dynamic cluster creation or
	// inline DNS resolution. See `issue
	// <https://github.com/envoyproxy/envoy/issues/1606>`_.
	//
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	// Sets the maximum duration in milliseconds that a response can take to arrive upon request.
	Timeout              *types.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a09f7c0d394634e, []int{0}
}
func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return m.Size()
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *types.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.api.v3alpha.core.HttpUri")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/core/http_uri.proto", fileDescriptor_6a09f7c0d394634e)
}

var fileDescriptor_6a09f7c0d394634e = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0x3b, 0x5d, 0x6b, 0x6d, 0xf4, 0x14, 0x44, 0x6b, 0x0b, 0xcb, 0x22, 0x16, 0x7a, 0x4a,
	0xa0, 0x7d, 0x83, 0xe0, 0xa1, 0xde, 0xca, 0x82, 0xe7, 0x92, 0xb6, 0xb1, 0x0d, 0x6c, 0x9b, 0x90,
	0x4e, 0x16, 0xf7, 0x95, 0xbc, 0x0b, 0xe2, 0xa9, 0x47, 0x8f, 0x3e, 0x82, 0xec, 0xad, 0x6f, 0x21,
	0xfb, 0xef, 0x22, 0xde, 0x86, 0xcc, 0x6f, 0xc2, 0xef, 0xfb, 0xc8, 0x48, 0xed, 0x53, 0x93, 0x71,
	0x69, 0x35, 0x4f, 0xa7, 0x32, 0xb1, 0x5b, 0xc9, 0x57, 0xc6, 0x29, 0xbe, 0x45, 0xb4, 0x0b, 0xef,
	0x34, 0xb3, 0xce, 0xa0, 0xa1, 0x37, 0x25, 0xc6, 0xa4, 0xd5, 0xac, 0xc6, 0x58, 0x81, 0x0d, 0xc2,
	0x8d, 0x31, 0x9b, 0x44, 0xf1, 0x92, 0x5a, 0xfa, 0x17, 0xbe, 0xf6, 0x4e, 0xa2, 0x36, 0xfb, 0xea,
	0x6e, 0x70, 0x9b, 0xca, 0x44, 0xaf, 0x25, 0x2a, 0xde, 0x0c, 0xd5, 0xe2, 0xfe, 0x1d, 0x48, 0x77,
	0x86, 0x68, 0x9f, 0x9d, 0xa6, 0x43, 0x12, 0x78, 0xa7, 0xfb, 0x10, 0xc1, 0xb8, 0x27, 0x7a, 0x9f,
	0xa7, 0x63, 0x70, 0xe6, 0xda, 0x11, 0xc4, 0xc5, 0x2b, 0x1d, 0x91, 0xee, 0x2a, 0xf1, 0x07, 0x54,
	0xae, 0xdf, 0xfe, 0x03, 0xcc, 0x5a, 0x71, 0xb3, 0xa3, 0x4f, 0xa4, 0x8b, 0x7a, 0xa7, 0x8c, 0xc7,
	0x7e, 0x10, 0xc1, 0xf8, 0x72, 0x72, 0xc7, 0x2a, 0x35, 0xd6, 0xa8, 0xb1, 0xc7, 0x5a, 0x4d, 0x5c,
	0x17, 0x3f, 0x74, 0xde, 0xa0, 0x3d, 0x69, 0x35, 0xd3, 0x05, 0xc4, 0xcd, 0xbd, 0x18, 0x12, 0x5a,
	0xa5, 0xb7, 0x07, 0x74, 0x4a, 0xee, 0x16, 0x98, 0x59, 0x45, 0x3b, 0x1f, 0xa7, 0x63, 0x00, 0x42,
	0x7c, 0xe5, 0x21, 0x7c, 0xe7, 0x21, 0xfc, 0xe4, 0x21, 0x90, 0x07, 0x6d, 0x58, 0xd9, 0x8c, 0x75,
	0xe6, 0x35, 0x63, 0xff, 0x97, 0x24, 0xae, 0xea, 0xa0, 0xf3, 0xc2, 0x64, 0x0e, 0xcb, 0xf3, 0x52,
	0x69, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xb1, 0xfb, 0x0d, 0x7c, 0x01, 0x00, 0x00,
}

func (m *HttpUri) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpUri) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpUri) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Timeout != nil {
		{
			size, err := m.Timeout.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttpUri(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.HttpUpstreamType != nil {
		{
			size := m.HttpUpstreamType.Size()
			i -= size
			if _, err := m.HttpUpstreamType.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintHttpUri(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HttpUri_Cluster) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *HttpUri_Cluster) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.Cluster)
	copy(dAtA[i:], m.Cluster)
	i = encodeVarintHttpUri(dAtA, i, uint64(len(m.Cluster)))
	i--
	dAtA[i] = 0x12
	return len(dAtA) - i, nil
}
func encodeVarintHttpUri(dAtA []byte, offset int, v uint64) int {
	offset -= sovHttpUri(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HttpUri) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovHttpUri(uint64(l))
	}
	if m.HttpUpstreamType != nil {
		n += m.HttpUpstreamType.Size()
	}
	if m.Timeout != nil {
		l = m.Timeout.Size()
		n += 1 + l + sovHttpUri(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpUri_Cluster) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Cluster)
	n += 1 + l + sovHttpUri(uint64(l))
	return n
}

func sovHttpUri(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHttpUri(x uint64) (n int) {
	return sovHttpUri(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpUri) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpUri
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
			return fmt.Errorf("proto: HttpUri: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpUri: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpUri
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
				return ErrInvalidLengthHttpUri
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHttpUri
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpUri
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
				return ErrInvalidLengthHttpUri
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHttpUri
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HttpUpstreamType = &HttpUri_Cluster{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timeout", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpUri
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
				return ErrInvalidLengthHttpUri
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttpUri
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timeout == nil {
				m.Timeout = &types.Duration{}
			}
			if err := m.Timeout.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpUri(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpUri
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttpUri
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
func skipHttpUri(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpUri
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
					return 0, ErrIntOverflowHttpUri
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
					return 0, ErrIntOverflowHttpUri
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
				return 0, ErrInvalidLengthHttpUri
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHttpUri
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpUri
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
				next, err := skipHttpUri(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHttpUri
				}
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
	ErrInvalidLengthHttpUri = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpUri   = fmt.Errorf("proto: integer overflow")
)
