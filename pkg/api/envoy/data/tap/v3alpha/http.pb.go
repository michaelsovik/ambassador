// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/data/tap/v3alpha/http.proto

package envoy_data_tap_v3alpha

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/pkg/api/envoy/api/v3alpha/core"
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

// A fully buffered HTTP trace message.
type HttpBufferedTrace struct {
	// Request message.
	Request *HttpBufferedTrace_Message `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	// Response message.
	Response             *HttpBufferedTrace_Message `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpBufferedTrace) Reset()         { *m = HttpBufferedTrace{} }
func (m *HttpBufferedTrace) String() string { return proto.CompactTextString(m) }
func (*HttpBufferedTrace) ProtoMessage()    {}
func (*HttpBufferedTrace) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd2efb22af23c41, []int{0}
}
func (m *HttpBufferedTrace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpBufferedTrace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpBufferedTrace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpBufferedTrace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBufferedTrace.Merge(m, src)
}
func (m *HttpBufferedTrace) XXX_Size() int {
	return m.Size()
}
func (m *HttpBufferedTrace) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBufferedTrace.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBufferedTrace proto.InternalMessageInfo

func (m *HttpBufferedTrace) GetRequest() *HttpBufferedTrace_Message {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *HttpBufferedTrace) GetResponse() *HttpBufferedTrace_Message {
	if m != nil {
		return m.Response
	}
	return nil
}

// HTTP message wrapper.
type HttpBufferedTrace_Message struct {
	// Message headers.
	Headers []*core.HeaderValue `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	// Message body.
	Body *Body `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// Message trailers.
	Trailers             []*core.HeaderValue `protobuf:"bytes,3,rep,name=trailers,proto3" json:"trailers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HttpBufferedTrace_Message) Reset()         { *m = HttpBufferedTrace_Message{} }
func (m *HttpBufferedTrace_Message) String() string { return proto.CompactTextString(m) }
func (*HttpBufferedTrace_Message) ProtoMessage()    {}
func (*HttpBufferedTrace_Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd2efb22af23c41, []int{0, 0}
}
func (m *HttpBufferedTrace_Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpBufferedTrace_Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpBufferedTrace_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpBufferedTrace_Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpBufferedTrace_Message.Merge(m, src)
}
func (m *HttpBufferedTrace_Message) XXX_Size() int {
	return m.Size()
}
func (m *HttpBufferedTrace_Message) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpBufferedTrace_Message.DiscardUnknown(m)
}

var xxx_messageInfo_HttpBufferedTrace_Message proto.InternalMessageInfo

func (m *HttpBufferedTrace_Message) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpBufferedTrace_Message) GetBody() *Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HttpBufferedTrace_Message) GetTrailers() []*core.HeaderValue {
	if m != nil {
		return m.Trailers
	}
	return nil
}

// A streamed HTTP trace segment. Multiple segments make up a full trace.
type HttpStreamedTraceSegment struct {
	// Trace ID unique to the originating Envoy only. Trace IDs can repeat and should not be used
	// for long term stable uniqueness.
	TraceId uint64 `protobuf:"varint,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// Types that are valid to be assigned to MessagePiece:
	//	*HttpStreamedTraceSegment_RequestHeaders
	//	*HttpStreamedTraceSegment_RequestBodyChunk
	//	*HttpStreamedTraceSegment_RequestTrailers
	//	*HttpStreamedTraceSegment_ResponseHeaders
	//	*HttpStreamedTraceSegment_ResponseBodyChunk
	//	*HttpStreamedTraceSegment_ResponseTrailers
	MessagePiece         isHttpStreamedTraceSegment_MessagePiece `protobuf_oneof:"message_piece"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *HttpStreamedTraceSegment) Reset()         { *m = HttpStreamedTraceSegment{} }
func (m *HttpStreamedTraceSegment) String() string { return proto.CompactTextString(m) }
func (*HttpStreamedTraceSegment) ProtoMessage()    {}
func (*HttpStreamedTraceSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_1fd2efb22af23c41, []int{1}
}
func (m *HttpStreamedTraceSegment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpStreamedTraceSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpStreamedTraceSegment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HttpStreamedTraceSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStreamedTraceSegment.Merge(m, src)
}
func (m *HttpStreamedTraceSegment) XXX_Size() int {
	return m.Size()
}
func (m *HttpStreamedTraceSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStreamedTraceSegment.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStreamedTraceSegment proto.InternalMessageInfo

type isHttpStreamedTraceSegment_MessagePiece interface {
	isHttpStreamedTraceSegment_MessagePiece()
	MarshalTo([]byte) (int, error)
	Size() int
}

type HttpStreamedTraceSegment_RequestHeaders struct {
	RequestHeaders *core.HeaderMap `protobuf:"bytes,2,opt,name=request_headers,json=requestHeaders,proto3,oneof"`
}
type HttpStreamedTraceSegment_RequestBodyChunk struct {
	RequestBodyChunk *Body `protobuf:"bytes,3,opt,name=request_body_chunk,json=requestBodyChunk,proto3,oneof"`
}
type HttpStreamedTraceSegment_RequestTrailers struct {
	RequestTrailers *core.HeaderMap `protobuf:"bytes,4,opt,name=request_trailers,json=requestTrailers,proto3,oneof"`
}
type HttpStreamedTraceSegment_ResponseHeaders struct {
	ResponseHeaders *core.HeaderMap `protobuf:"bytes,5,opt,name=response_headers,json=responseHeaders,proto3,oneof"`
}
type HttpStreamedTraceSegment_ResponseBodyChunk struct {
	ResponseBodyChunk *Body `protobuf:"bytes,6,opt,name=response_body_chunk,json=responseBodyChunk,proto3,oneof"`
}
type HttpStreamedTraceSegment_ResponseTrailers struct {
	ResponseTrailers *core.HeaderMap `protobuf:"bytes,7,opt,name=response_trailers,json=responseTrailers,proto3,oneof"`
}

func (*HttpStreamedTraceSegment_RequestHeaders) isHttpStreamedTraceSegment_MessagePiece()    {}
func (*HttpStreamedTraceSegment_RequestBodyChunk) isHttpStreamedTraceSegment_MessagePiece()  {}
func (*HttpStreamedTraceSegment_RequestTrailers) isHttpStreamedTraceSegment_MessagePiece()   {}
func (*HttpStreamedTraceSegment_ResponseHeaders) isHttpStreamedTraceSegment_MessagePiece()   {}
func (*HttpStreamedTraceSegment_ResponseBodyChunk) isHttpStreamedTraceSegment_MessagePiece() {}
func (*HttpStreamedTraceSegment_ResponseTrailers) isHttpStreamedTraceSegment_MessagePiece()  {}

func (m *HttpStreamedTraceSegment) GetMessagePiece() isHttpStreamedTraceSegment_MessagePiece {
	if m != nil {
		return m.MessagePiece
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetTraceId() uint64 {
	if m != nil {
		return m.TraceId
	}
	return 0
}

func (m *HttpStreamedTraceSegment) GetRequestHeaders() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestHeaders); ok {
		return x.RequestHeaders
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestBodyChunk() *Body {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestBodyChunk); ok {
		return x.RequestBodyChunk
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetRequestTrailers() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_RequestTrailers); ok {
		return x.RequestTrailers
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseHeaders() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseHeaders); ok {
		return x.ResponseHeaders
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseBodyChunk() *Body {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseBodyChunk); ok {
		return x.ResponseBodyChunk
	}
	return nil
}

func (m *HttpStreamedTraceSegment) GetResponseTrailers() *core.HeaderMap {
	if x, ok := m.GetMessagePiece().(*HttpStreamedTraceSegment_ResponseTrailers); ok {
		return x.ResponseTrailers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpStreamedTraceSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpStreamedTraceSegment_RequestHeaders)(nil),
		(*HttpStreamedTraceSegment_RequestBodyChunk)(nil),
		(*HttpStreamedTraceSegment_RequestTrailers)(nil),
		(*HttpStreamedTraceSegment_ResponseHeaders)(nil),
		(*HttpStreamedTraceSegment_ResponseBodyChunk)(nil),
		(*HttpStreamedTraceSegment_ResponseTrailers)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpBufferedTrace)(nil), "envoy.data.tap.v3alpha.HttpBufferedTrace")
	proto.RegisterType((*HttpBufferedTrace_Message)(nil), "envoy.data.tap.v3alpha.HttpBufferedTrace.Message")
	proto.RegisterType((*HttpStreamedTraceSegment)(nil), "envoy.data.tap.v3alpha.HttpStreamedTraceSegment")
}

func init() { proto.RegisterFile("envoy/data/tap/v3alpha/http.proto", fileDescriptor_1fd2efb22af23c41) }

var fileDescriptor_1fd2efb22af23c41 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xeb, 0x26, 0xd4, 0xe5, 0x56, 0x50, 0x3a, 0x48, 0xc8, 0x44, 0x28, 0xea, 0x0f, 0x8b,
	0xae, 0x6c, 0x68, 0xd7, 0xa8, 0x92, 0xd9, 0x18, 0xd1, 0x54, 0x91, 0x5b, 0xb1, 0xb5, 0x6e, 0xec,
	0xdb, 0xc6, 0x22, 0xf6, 0x0c, 0xe3, 0x49, 0x45, 0xde, 0x87, 0xb7, 0xe0, 0x05, 0x58, 0xf2, 0x08,
	0x28, 0x7b, 0xde, 0x01, 0x8d, 0x3d, 0x33, 0x45, 0x82, 0x40, 0x60, 0xe9, 0xd1, 0x39, 0xe7, 0x9e,
	0xef, 0x7a, 0x06, 0x0e, 0xa8, 0xbe, 0xe5, 0x8b, 0xa8, 0x40, 0x85, 0x91, 0x42, 0x11, 0xdd, 0x9e,
	0xe2, 0x4c, 0x4c, 0x31, 0x9a, 0x2a, 0x25, 0x42, 0x21, 0xb9, 0xe2, 0xec, 0x49, 0x2b, 0x09, 0xb5,
	0x24, 0x54, 0x28, 0x42, 0x23, 0x19, 0x18, 0x2b, 0x8a, 0xd2, 0xb9, 0x72, 0x2e, 0x29, 0x9a, 0x60,
	0x43, 0x9d, 0x75, 0x70, 0xb4, 0x22, 0x3d, 0xe7, 0x55, 0xc5, 0xeb, 0x4e, 0x74, 0xf8, 0x7d, 0x13,
	0xf6, 0x12, 0xa5, 0x44, 0x3c, 0xbf, 0xbe, 0x26, 0x49, 0xc5, 0x95, 0xc4, 0x9c, 0xd8, 0x5b, 0xf0,
	0x25, 0x7d, 0x98, 0x53, 0xa3, 0x02, 0x6f, 0xdf, 0x3b, 0xde, 0x39, 0x79, 0x19, 0xfe, 0xbe, 0x47,
	0xf8, 0x8b, 0x37, 0x1c, 0x51, 0xd3, 0xe0, 0x0d, 0xa5, 0x36, 0x81, 0x8d, 0x60, 0x5b, 0x52, 0x23,
	0x78, 0xdd, 0x50, 0xb0, 0xf9, 0xbf, 0x69, 0x2e, 0x62, 0xf0, 0xd9, 0x03, 0xdf, 0x9c, 0xb2, 0x57,
	0xe0, 0x4f, 0x09, 0x0b, 0x92, 0x4d, 0xe0, 0xed, 0xf7, 0x8e, 0x77, 0x4e, 0x8e, 0x4c, 0x32, 0x8a,
	0xd2, 0x85, 0xea, 0xbd, 0x84, 0x49, 0x2b, 0x7b, 0x87, 0xb3, 0x39, 0xa5, 0xd6, 0xc3, 0x5e, 0x40,
	0x7f, 0xc2, 0x8b, 0x85, 0x69, 0xf5, 0x6c, 0x55, 0xab, 0x98, 0x17, 0x8b, 0xb4, 0x55, 0xb2, 0x33,
	0xd8, 0x56, 0x12, 0xcb, 0x99, 0x9e, 0xd8, 0x5b, 0x7f, 0xa2, 0x33, 0x1d, 0x7e, 0xea, 0x43, 0xa0,
	0x29, 0x2f, 0x95, 0x24, 0xac, 0x0c, 0xe5, 0x25, 0xdd, 0x54, 0x54, 0x2b, 0xf6, 0xb4, 0x4d, 0xcf,
	0x29, 0x2b, 0x8b, 0x76, 0xef, 0xfd, 0xd4, 0x6f, 0xbf, 0xdf, 0x14, 0xec, 0x1c, 0x76, 0xcd, 0x3e,
	0x33, 0x4b, 0xdc, 0xb5, 0x3e, 0xf8, 0xf3, 0xfc, 0x11, 0x8a, 0x64, 0x23, 0x7d, 0x68, 0xbc, 0x89,
	0x01, 0x3f, 0x07, 0x66, 0xd3, 0x34, 0x56, 0x96, 0x4f, 0xe7, 0xf5, 0xfb, 0xa0, 0xf7, 0xf7, 0x35,
	0x24, 0x1b, 0xe9, 0x23, 0xe3, 0xd4, 0x9f, 0xaf, 0xb5, 0x8f, 0x5d, 0x80, 0x3d, 0xcb, 0xdc, 0x72,
	0xfa, 0xeb, 0x97, 0xb3, 0x60, 0x57, 0xc6, 0xdb, 0xe5, 0x75, 0x7f, 0xdb, 0xc1, 0xde, 0xfb, 0xa7,
	0xbc, 0xce, 0x6c, 0x69, 0x2f, 0xe0, 0xb1, 0xcb, 0xfb, 0x09, 0x77, 0x6b, 0x2d, 0xdc, 0x3d, 0x6b,
	0xbd, 0xe3, 0x1d, 0x83, 0x3b, 0xbc, 0x03, 0xf6, 0xd7, 0x2f, 0xe8, 0xe8, 0x2c, 0x71, 0xbc, 0x0b,
	0x0f, 0xaa, 0xee, 0x4a, 0x67, 0xa2, 0xa4, 0x9c, 0xe2, 0xb3, 0x2f, 0xcb, 0xa1, 0xf7, 0x75, 0x39,
	0xf4, 0xbe, 0x2d, 0x87, 0x1e, 0x3c, 0x2f, 0x79, 0x97, 0x2b, 0x24, 0xff, 0xb8, 0x58, 0x51, 0x38,
	0xbe, 0xaf, 0xef, 0xd5, 0x58, 0xbf, 0xea, 0xb1, 0x37, 0xd9, 0x6a, 0x9f, 0xf7, 0xe9, 0x8f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0xe5, 0xab, 0x14, 0x63, 0x04, 0x00, 0x00,
}

func (m *HttpBufferedTrace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpBufferedTrace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpBufferedTrace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Response != nil {
		{
			size, err := m.Response.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HttpBufferedTrace_Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpBufferedTrace_Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpBufferedTrace_Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Trailers) > 0 {
		for iNdEx := len(m.Trailers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Trailers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHttp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Headers) > 0 {
		for iNdEx := len(m.Headers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Headers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintHttp(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *HttpStreamedTraceSegment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpStreamedTraceSegment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HttpStreamedTraceSegment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.MessagePiece != nil {
		{
			size := m.MessagePiece.Size()
			i -= size
			if _, err := m.MessagePiece.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.TraceId != 0 {
		i = encodeVarintHttp(dAtA, i, uint64(m.TraceId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HttpStreamedTraceSegment_RequestHeaders) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *HttpStreamedTraceSegment_RequestHeaders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestHeaders != nil {
		{
			size, err := m.RequestHeaders.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *HttpStreamedTraceSegment_RequestBodyChunk) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *HttpStreamedTraceSegment_RequestBodyChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestBodyChunk != nil {
		{
			size, err := m.RequestBodyChunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	return len(dAtA) - i, nil
}
func (m *HttpStreamedTraceSegment_RequestTrailers) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *HttpStreamedTraceSegment_RequestTrailers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.RequestTrailers != nil {
		{
			size, err := m.RequestTrailers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	return len(dAtA) - i, nil
}
func (m *HttpStreamedTraceSegment_ResponseHeaders) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *HttpStreamedTraceSegment_ResponseHeaders) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResponseHeaders != nil {
		{
			size, err := m.ResponseHeaders.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	return len(dAtA) - i, nil
}
func (m *HttpStreamedTraceSegment_ResponseBodyChunk) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *HttpStreamedTraceSegment_ResponseBodyChunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResponseBodyChunk != nil {
		{
			size, err := m.ResponseBodyChunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *HttpStreamedTraceSegment_ResponseTrailers) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *HttpStreamedTraceSegment_ResponseTrailers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.ResponseTrailers != nil {
		{
			size, err := m.ResponseTrailers.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHttp(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	return len(dAtA) - i, nil
}
func encodeVarintHttp(dAtA []byte, offset int, v uint64) int {
	offset -= sovHttp(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HttpBufferedTrace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.Response != nil {
		l = m.Response.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpBufferedTrace_Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovHttp(uint64(l))
		}
	}
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	if len(m.Trailers) > 0 {
		for _, e := range m.Trailers {
			l = e.Size()
			n += 1 + l + sovHttp(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpStreamedTraceSegment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TraceId != 0 {
		n += 1 + sovHttp(uint64(m.TraceId))
	}
	if m.MessagePiece != nil {
		n += m.MessagePiece.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *HttpStreamedTraceSegment_RequestHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestHeaders != nil {
		l = m.RequestHeaders.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	return n
}
func (m *HttpStreamedTraceSegment_RequestBodyChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestBodyChunk != nil {
		l = m.RequestBodyChunk.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	return n
}
func (m *HttpStreamedTraceSegment_RequestTrailers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestTrailers != nil {
		l = m.RequestTrailers.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	return n
}
func (m *HttpStreamedTraceSegment_ResponseHeaders) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseHeaders != nil {
		l = m.ResponseHeaders.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	return n
}
func (m *HttpStreamedTraceSegment_ResponseBodyChunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseBodyChunk != nil {
		l = m.ResponseBodyChunk.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	return n
}
func (m *HttpStreamedTraceSegment_ResponseTrailers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResponseTrailers != nil {
		l = m.ResponseTrailers.Size()
		n += 1 + l + sovHttp(uint64(l))
	}
	return n
}

func sovHttp(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHttp(x uint64) (n int) {
	return sovHttp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpBufferedTrace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttp
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
			return fmt.Errorf("proto: HttpBufferedTrace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpBufferedTrace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &HttpBufferedTrace_Message{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Response == nil {
				m.Response = &HttpBufferedTrace_Message{}
			}
			if err := m.Response.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttp
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
func (m *HttpBufferedTrace_Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttp
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
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &core.HeaderValue{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &Body{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trailers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trailers = append(m.Trailers, &core.HeaderValue{})
			if err := m.Trailers[len(m.Trailers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttp
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
func (m *HttpStreamedTraceSegment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttp
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
			return fmt.Errorf("proto: HttpStreamedTraceSegment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpStreamedTraceSegment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TraceId", wireType)
			}
			m.TraceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TraceId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestHeaders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &core.HeaderMap{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MessagePiece = &HttpStreamedTraceSegment_RequestHeaders{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestBodyChunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Body{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MessagePiece = &HttpStreamedTraceSegment_RequestBodyChunk{v}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTrailers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &core.HeaderMap{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MessagePiece = &HttpStreamedTraceSegment_RequestTrailers{v}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseHeaders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &core.HeaderMap{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MessagePiece = &HttpStreamedTraceSegment_ResponseHeaders{v}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseBodyChunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Body{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MessagePiece = &HttpStreamedTraceSegment_ResponseBodyChunk{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseTrailers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
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
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHttp
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &core.HeaderMap{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.MessagePiece = &HttpStreamedTraceSegment_ResponseTrailers{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttp
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHttp
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
func skipHttp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttp
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
					return 0, ErrIntOverflowHttp
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
					return 0, ErrIntOverflowHttp
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
				return 0, ErrInvalidLengthHttp
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHttp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttp
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
				next, err := skipHttp(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHttp
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
	ErrInvalidLengthHttp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttp   = fmt.Errorf("proto: integer overflow")
)
