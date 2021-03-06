// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeattestor.proto

package nodeattestor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
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

//* Represents a request to attest a node.
type AttestRequest struct {
	//* A type which contains attestation data for specific platform.
	AttestationData *common.AttestationData `protobuf:"bytes,1,opt,name=attestationData,proto3" json:"attestationData,omitempty"`
	//* Is true if the Base SPIFFE ID is present in the Attested Node table.
	AttestedBefore bool `protobuf:"varint,2,opt,name=attestedBefore,proto3" json:"attestedBefore,omitempty"`
	//* Challenge response
	Response             []byte   `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestRequest) Reset()         { *m = AttestRequest{} }
func (m *AttestRequest) String() string { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()    {}
func (*AttestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3b2582c38c076c, []int{0}
}

func (m *AttestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestRequest.Unmarshal(m, b)
}
func (m *AttestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestRequest.Marshal(b, m, deterministic)
}
func (m *AttestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestRequest.Merge(m, src)
}
func (m *AttestRequest) XXX_Size() int {
	return xxx_messageInfo_AttestRequest.Size(m)
}
func (m *AttestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttestRequest proto.InternalMessageInfo

func (m *AttestRequest) GetAttestationData() *common.AttestationData {
	if m != nil {
		return m.AttestationData
	}
	return nil
}

func (m *AttestRequest) GetAttestedBefore() bool {
	if m != nil {
		return m.AttestedBefore
	}
	return false
}

func (m *AttestRequest) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

//* Represents a response when attesting a node.
type AttestResponse struct {
	//* True/False
	Valid bool `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	//* Used by the Server to validate the SPIFFE Id in the Certificate signing request.
	BaseSPIFFEID string `protobuf:"bytes,2,opt,name=baseSPIFFEID,proto3" json:"baseSPIFFEID,omitempty"`
	//* Challenge required for attestation
	Challenge []byte `protobuf:"bytes,3,opt,name=challenge,proto3" json:"challenge,omitempty"`
	//* Optional list of selectors
	Selectors            []*common.Selector `protobuf:"bytes,4,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AttestResponse) Reset()         { *m = AttestResponse{} }
func (m *AttestResponse) String() string { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()    {}
func (*AttestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3b2582c38c076c, []int{1}
}

func (m *AttestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestResponse.Unmarshal(m, b)
}
func (m *AttestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestResponse.Marshal(b, m, deterministic)
}
func (m *AttestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestResponse.Merge(m, src)
}
func (m *AttestResponse) XXX_Size() int {
	return xxx_messageInfo_AttestResponse.Size(m)
}
func (m *AttestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttestResponse proto.InternalMessageInfo

func (m *AttestResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *AttestResponse) GetBaseSPIFFEID() string {
	if m != nil {
		return m.BaseSPIFFEID
	}
	return ""
}

func (m *AttestResponse) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *AttestResponse) GetSelectors() []*common.Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "spire.agent.nodeattestor.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "spire.agent.nodeattestor.AttestResponse")
}

func init() { proto.RegisterFile("nodeattestor.proto", fileDescriptor_9e3b2582c38c076c) }

var fileDescriptor_9e3b2582c38c076c = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdf, 0x6e, 0xda, 0x30,
	0x14, 0xc6, 0x65, 0xd8, 0x10, 0x31, 0x7f, 0x26, 0x59, 0xd3, 0x94, 0x45, 0x9b, 0x14, 0x21, 0x8d,
	0x65, 0xbb, 0x48, 0x2a, 0x5a, 0xa9, 0xea, 0x25, 0x94, 0x82, 0xb8, 0xa9, 0x50, 0xb8, 0xe3, 0xa6,
	0x32, 0xe4, 0x24, 0x44, 0x0a, 0x76, 0x6a, 0x3b, 0x3c, 0x4d, 0xd5, 0x97, 0xe8, 0x0b, 0x56, 0x8d,
	0x13, 0x20, 0xa8, 0x15, 0xbd, 0xb2, 0x7d, 0xbe, 0xdf, 0xf1, 0x77, 0x7c, 0x8e, 0x31, 0x61, 0x3c,
	0x00, 0xaa, 0x14, 0x48, 0xc5, 0x85, 0x9b, 0x0a, 0xae, 0x38, 0x31, 0x65, 0x1a, 0x0b, 0x70, 0x69,
	0x04, 0x4c, 0xb9, 0xc7, 0xba, 0x65, 0xe7, 0x8a, 0xb7, 0xe6, 0xdb, 0x2d, 0x67, 0x5e, 0x9a, 0x64,
	0x51, 0x5c, 0x2e, 0x3a, 0xd7, 0xfa, 0x59, 0x21, 0xf4, 0xa2, 0xa5, 0xde, 0x13, 0xc2, 0x9d, 0x61,
	0x7e, 0x93, 0x0f, 0x8f, 0x19, 0x48, 0x45, 0xa6, 0xf8, 0x9b, 0xbe, 0x9a, 0xaa, 0x98, 0xb3, 0x31,
	0x55, 0xd4, 0x44, 0x36, 0x72, 0x5a, 0x83, 0xdf, 0xae, 0x2e, 0xa1, 0xc8, 0x1f, 0x56, 0x21, 0xff,
	0x34, 0x8b, 0xf4, 0x71, 0x57, 0x87, 0x20, 0x18, 0x41, 0xc8, 0x05, 0x98, 0x35, 0x1b, 0x39, 0x4d,
	0xff, 0x24, 0x4a, 0x2c, 0xdc, 0x14, 0x20, 0x53, 0xce, 0x24, 0x98, 0x75, 0x1b, 0x39, 0x6d, 0x7f,
	0x7f, 0xee, 0x3d, 0x23, 0xdc, 0x2d, 0xcb, 0xd3, 0x21, 0xf2, 0x1d, 0x7f, 0xdd, 0xd1, 0x24, 0x0e,
	0xf2, 0xaa, 0x9a, 0xbe, 0x3e, 0x90, 0x1e, 0x6e, 0xaf, 0xa8, 0x84, 0xc5, 0x7c, 0x36, 0x99, 0xdc,
	0xcd, 0xc6, 0xb9, 0x95, 0xe1, 0x57, 0x62, 0xe4, 0x17, 0x36, 0xd6, 0x1b, 0x9a, 0x24, 0xc0, 0xa2,
	0xd2, 0xe9, 0x10, 0x20, 0x57, 0xd8, 0x90, 0x90, 0xc0, 0x5a, 0x71, 0x21, 0xcd, 0x2f, 0x76, 0xdd,
	0x69, 0x0d, 0x7e, 0x54, 0x5f, 0xbc, 0x28, 0x64, 0xff, 0x00, 0x0e, 0x5e, 0x6a, 0xb8, 0x7d, 0xcf,
	0x03, 0x18, 0x16, 0xd3, 0x20, 0x0f, 0xb8, 0xa1, 0xf7, 0xe4, 0xaf, 0xfb, 0xd1, 0xc8, 0xdc, 0x4a,
	0xc7, 0x2d, 0xe7, 0x3c, 0xa8, 0xdf, 0xee, 0xa0, 0x0b, 0x44, 0x96, 0xd8, 0xb8, 0xe5, 0x2c, 0x8c,
	0xa3, 0x4c, 0x00, 0xf9, 0x53, 0xad, 0xb0, 0x98, 0xfa, 0x5e, 0x2f, 0x1d, 0xfa, 0xe7, 0xb0, 0xa2,
	0xb7, 0x21, 0xee, 0x4c, 0x41, 0xcd, 0x73, 0x79, 0xc6, 0x42, 0x4e, 0xfe, 0xbd, 0x9b, 0x58, 0x61,
	0x4a, 0x8f, 0xff, 0x9f, 0x41, 0xb5, 0xcf, 0xe8, 0x66, 0x79, 0x1d, 0xc5, 0x6a, 0x93, 0xad, 0xde,
	0x68, 0x4f, 0xa6, 0x71, 0x18, 0x82, 0xa7, 0x3f, 0x69, 0xfe, 0x2d, 0x8b, 0xbd, 0x04, 0xb1, 0x03,
	0xe1, 0x1d, 0x77, 0x64, 0xd5, 0xc8, 0x81, 0xcb, 0xd7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x71,
	0xb6, 0x04, 0x24, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeAttestorClient is the client API for NodeAttestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeAttestorClient interface {
	//* Attesta a node.
	Attest(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_AttestClient, error)
	//* Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	//* Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) Attest(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_AttestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAttestor_serviceDesc.Streams[0], "/spire.agent.nodeattestor.NodeAttestor/Attest", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAttestorAttestClient{stream}
	return x, nil
}

type NodeAttestor_AttestClient interface {
	Send(*AttestRequest) error
	Recv() (*AttestResponse, error)
	grpc.ClientStream
}

type nodeAttestorAttestClient struct {
	grpc.ClientStream
}

func (x *nodeAttestorAttestClient) Send(m *AttestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeAttestorAttestClient) Recv() (*AttestResponse, error) {
	m := new(AttestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeAttestorServer is the server API for NodeAttestor service.
type NodeAttestorServer interface {
	//* Attesta a node.
	Attest(NodeAttestor_AttestServer) error
	//* Responsible for configuration of the plugin.
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	//* Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_Attest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeAttestorServer).Attest(&nodeAttestorAttestServer{stream})
}

type NodeAttestor_AttestServer interface {
	Send(*AttestResponse) error
	Recv() (*AttestRequest, error)
	grpc.ServerStream
}

type nodeAttestorAttestServer struct {
	grpc.ServerStream
}

func (x *nodeAttestorAttestServer) Send(m *AttestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeAttestorAttestServer) Recv() (*AttestRequest, error) {
	m := new(AttestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.nodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Attest",
			Handler:       _NodeAttestor_Attest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "nodeattestor.proto",
}
