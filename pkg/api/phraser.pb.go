// Code generated by protoc-gen-go. DO NOT EDIT.
// source: phraser.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Phrase struct {
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Collection           string               `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	Path                 string               `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Value                string               `protobuf:"bytes,5,opt,name=value,proto3" json:"value,omitempty"`
	Set                  bool                 `protobuf:"varint,6,opt,name=set,proto3" json:"set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Phrase) Reset()         { *m = Phrase{} }
func (m *Phrase) String() string { return proto.CompactTextString(m) }
func (*Phrase) ProtoMessage()    {}
func (*Phrase) Descriptor() ([]byte, []int) {
	return fileDescriptor_91a0bc29b5b936c9, []int{0}
}

func (m *Phrase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Phrase.Unmarshal(m, b)
}
func (m *Phrase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Phrase.Marshal(b, m, deterministic)
}
func (m *Phrase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Phrase.Merge(m, src)
}
func (m *Phrase) XXX_Size() int {
	return xxx_messageInfo_Phrase.Size(m)
}
func (m *Phrase) XXX_DiscardUnknown() {
	xxx_messageInfo_Phrase.DiscardUnknown(m)
}

var xxx_messageInfo_Phrase proto.InternalMessageInfo

func (m *Phrase) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Phrase) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Phrase) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *Phrase) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Phrase) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Phrase) GetSet() bool {
	if m != nil {
		return m.Set
	}
	return false
}

type SetPhraseRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPhraseRequest) Reset()         { *m = SetPhraseRequest{} }
func (m *SetPhraseRequest) String() string { return proto.CompactTextString(m) }
func (*SetPhraseRequest) ProtoMessage()    {}
func (*SetPhraseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91a0bc29b5b936c9, []int{1}
}

func (m *SetPhraseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPhraseRequest.Unmarshal(m, b)
}
func (m *SetPhraseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPhraseRequest.Marshal(b, m, deterministic)
}
func (m *SetPhraseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPhraseRequest.Merge(m, src)
}
func (m *SetPhraseRequest) XXX_Size() int {
	return xxx_messageInfo_SetPhraseRequest.Size(m)
}
func (m *SetPhraseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPhraseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetPhraseRequest proto.InternalMessageInfo

func (m *SetPhraseRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *SetPhraseRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *SetPhraseRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetPhraseRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPhraseRequest) Reset()         { *m = GetPhraseRequest{} }
func (m *GetPhraseRequest) String() string { return proto.CompactTextString(m) }
func (*GetPhraseRequest) ProtoMessage()    {}
func (*GetPhraseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91a0bc29b5b936c9, []int{2}
}

func (m *GetPhraseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPhraseRequest.Unmarshal(m, b)
}
func (m *GetPhraseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPhraseRequest.Marshal(b, m, deterministic)
}
func (m *GetPhraseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPhraseRequest.Merge(m, src)
}
func (m *GetPhraseRequest) XXX_Size() int {
	return xxx_messageInfo_GetPhraseRequest.Size(m)
}
func (m *GetPhraseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPhraseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPhraseRequest proto.InternalMessageInfo

func (m *GetPhraseRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *GetPhraseRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type ListPhrasesRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	PathPrefix           string   `protobuf:"bytes,2,opt,name=pathPrefix,proto3" json:"pathPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPhrasesRequest) Reset()         { *m = ListPhrasesRequest{} }
func (m *ListPhrasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPhrasesRequest) ProtoMessage()    {}
func (*ListPhrasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91a0bc29b5b936c9, []int{3}
}

func (m *ListPhrasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPhrasesRequest.Unmarshal(m, b)
}
func (m *ListPhrasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPhrasesRequest.Marshal(b, m, deterministic)
}
func (m *ListPhrasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPhrasesRequest.Merge(m, src)
}
func (m *ListPhrasesRequest) XXX_Size() int {
	return xxx_messageInfo_ListPhrasesRequest.Size(m)
}
func (m *ListPhrasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPhrasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPhrasesRequest proto.InternalMessageInfo

func (m *ListPhrasesRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *ListPhrasesRequest) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*Phrase)(nil), "api.Phrase")
	proto.RegisterType((*SetPhraseRequest)(nil), "api.SetPhraseRequest")
	proto.RegisterType((*GetPhraseRequest)(nil), "api.GetPhraseRequest")
	proto.RegisterType((*ListPhrasesRequest)(nil), "api.ListPhrasesRequest")
}

func init() { proto.RegisterFile("phraser.proto", fileDescriptor_91a0bc29b5b936c9) }

var fileDescriptor_91a0bc29b5b936c9 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0xaf, 0x9b, 0xb6, 0x97, 0x9e, 0x0a, 0xa9, 0xb2, 0x40, 0x44, 0x1d, 0x4a, 0x95, 0xa9,
	0x93, 0x0b, 0x65, 0x40, 0x8c, 0x5d, 0xc8, 0xc2, 0x50, 0x85, 0x8e, 0x48, 0xc8, 0x4d, 0x4f, 0x5b,
	0x4b, 0x29, 0x36, 0xb6, 0x83, 0x78, 0x24, 0x9e, 0x89, 0xa7, 0x41, 0xb1, 0x13, 0x14, 0x05, 0x50,
	0x19, 0xd8, 0x9c, 0xdf, 0xe7, 0xfb, 0x7e, 0x1d, 0x2b, 0x70, 0xac, 0x76, 0x9a, 0x1b, 0xd4, 0x4c,
	0x69, 0x69, 0x25, 0x0d, 0xb8, 0x12, 0xc3, 0xf3, 0xad, 0x94, 0xdb, 0x0c, 0xa7, 0x2e, 0x5a, 0xe5,
	0x9b, 0xa9, 0x15, 0x7b, 0x34, 0x96, 0xef, 0x95, 0x9f, 0x8a, 0xde, 0x09, 0x74, 0x17, 0x8e, 0xa3,
	0x37, 0x00, 0xa9, 0x46, 0x6e, 0x71, 0xfd, 0xc8, 0x6d, 0x48, 0xc6, 0x64, 0xd2, 0x9f, 0x0d, 0x99,
	0x17, 0xb0, 0x4a, 0xc0, 0x96, 0x95, 0x20, 0xe9, 0x95, 0xd3, 0x73, 0x5b, 0xa0, 0xb9, 0x5a, 0x57,
	0x68, 0xeb, 0x30, 0x5a, 0x4e, 0xcf, 0x2d, 0x1d, 0x01, 0xa4, 0x32, 0xcb, 0x30, 0xb5, 0x42, 0x3e,
	0x85, 0xc1, 0x98, 0x4c, 0x7a, 0x49, 0x2d, 0xa1, 0x14, 0xda, 0x8a, 0xdb, 0x5d, 0xd8, 0x76, 0x37,
	0xee, 0x4c, 0x4f, 0xa0, 0xf3, 0xc2, 0xb3, 0x1c, 0xc3, 0x8e, 0x0b, 0xfd, 0x07, 0x1d, 0x40, 0x60,
	0xd0, 0x86, 0xdd, 0x31, 0x99, 0x1c, 0x25, 0xc5, 0x31, 0x7a, 0x80, 0xc1, 0x3d, 0x5a, 0xbf, 0x5e,
	0x82, 0xcf, 0x39, 0x9a, 0x66, 0x1f, 0xf9, 0xb1, 0xaf, 0xf5, 0x5d, 0x5f, 0x50, 0xeb, 0x8b, 0x6e,
	0x61, 0x10, 0xff, 0x81, 0x3d, 0x5a, 0x02, 0xbd, 0x13, 0xa6, 0x14, 0x99, 0xdf, 0x9a, 0x46, 0x00,
	0x05, 0xbd, 0xd0, 0xb8, 0x11, 0xaf, 0xa5, 0xaf, 0x96, 0xcc, 0xde, 0x08, 0xfc, 0xf7, 0x4a, 0x4d,
	0x2f, 0xa1, 0xf7, 0xf9, 0x0e, 0xf4, 0x94, 0x71, 0x25, 0x58, 0xf3, 0x5d, 0x86, 0x7d, 0x17, 0xfb,
	0x2c, 0xfa, 0x57, 0x20, 0x71, 0x03, 0x89, 0x0f, 0x20, 0xd7, 0xd0, 0xaf, 0xed, 0x41, 0xcf, 0xdc,
	0xed, 0xd7, 0xcd, 0x1a, 0xd8, 0x05, 0x59, 0x75, 0xdd, 0x1f, 0x72, 0xf5, 0x11, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0xb9, 0x86, 0xde, 0xc1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PhraserClient is the client API for Phraser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhraserClient interface {
	SetPhrase(ctx context.Context, in *SetPhraseRequest, opts ...grpc.CallOption) (*Phrase, error)
	GetPhrase(ctx context.Context, in *GetPhraseRequest, opts ...grpc.CallOption) (*Phrase, error)
	ListPhrases(ctx context.Context, in *ListPhrasesRequest, opts ...grpc.CallOption) (Phraser_ListPhrasesClient, error)
}

type phraserClient struct {
	cc *grpc.ClientConn
}

func NewPhraserClient(cc *grpc.ClientConn) PhraserClient {
	return &phraserClient{cc}
}

func (c *phraserClient) SetPhrase(ctx context.Context, in *SetPhraseRequest, opts ...grpc.CallOption) (*Phrase, error) {
	out := new(Phrase)
	err := c.cc.Invoke(ctx, "/api.Phraser/SetPhrase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phraserClient) GetPhrase(ctx context.Context, in *GetPhraseRequest, opts ...grpc.CallOption) (*Phrase, error) {
	out := new(Phrase)
	err := c.cc.Invoke(ctx, "/api.Phraser/GetPhrase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phraserClient) ListPhrases(ctx context.Context, in *ListPhrasesRequest, opts ...grpc.CallOption) (Phraser_ListPhrasesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Phraser_serviceDesc.Streams[0], "/api.Phraser/ListPhrases", opts...)
	if err != nil {
		return nil, err
	}
	x := &phraserListPhrasesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Phraser_ListPhrasesClient interface {
	Recv() (*Phrase, error)
	grpc.ClientStream
}

type phraserListPhrasesClient struct {
	grpc.ClientStream
}

func (x *phraserListPhrasesClient) Recv() (*Phrase, error) {
	m := new(Phrase)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PhraserServer is the server API for Phraser service.
type PhraserServer interface {
	SetPhrase(context.Context, *SetPhraseRequest) (*Phrase, error)
	GetPhrase(context.Context, *GetPhraseRequest) (*Phrase, error)
	ListPhrases(*ListPhrasesRequest, Phraser_ListPhrasesServer) error
}

func RegisterPhraserServer(s *grpc.Server, srv PhraserServer) {
	s.RegisterService(&_Phraser_serviceDesc, srv)
}

func _Phraser_SetPhrase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPhraseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhraserServer).SetPhrase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Phraser/SetPhrase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhraserServer).SetPhrase(ctx, req.(*SetPhraseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Phraser_GetPhrase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPhraseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhraserServer).GetPhrase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Phraser/GetPhrase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhraserServer).GetPhrase(ctx, req.(*GetPhraseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Phraser_ListPhrases_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPhrasesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PhraserServer).ListPhrases(m, &phraserListPhrasesServer{stream})
}

type Phraser_ListPhrasesServer interface {
	Send(*Phrase) error
	grpc.ServerStream
}

type phraserListPhrasesServer struct {
	grpc.ServerStream
}

func (x *phraserListPhrasesServer) Send(m *Phrase) error {
	return x.ServerStream.SendMsg(m)
}

var _Phraser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Phraser",
	HandlerType: (*PhraserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetPhrase",
			Handler:    _Phraser_SetPhrase_Handler,
		},
		{
			MethodName: "GetPhrase",
			Handler:    _Phraser_GetPhrase_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPhrases",
			Handler:       _Phraser_ListPhrases_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "phraser.proto",
}