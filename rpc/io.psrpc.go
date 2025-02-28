// Code generated by protoc-gen-psrpc v0.2.9, DO NOT EDIT.
// source: rpc/io.proto

package rpc

import context "context"
import psrpc "github.com/livekit/psrpc"
import version "github.com/livekit/psrpc/version"
import google_protobuf2 "google.golang.org/protobuf/types/known/emptypb"
import livekit "github.com/livekit/protocol/livekit"

var _ = version.PsrpcVersion_0_2_9

// =======================
// IOInfo Client Interface
// =======================

type IOInfoClient interface {
	UpdateEgressInfo(ctx context.Context, req *livekit.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf2.Empty, error)

	GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error)

	UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf2.Empty, error)
}

// ===========================
// IOInfo ServerImpl Interface
// ===========================

type IOInfoServerImpl interface {
	UpdateEgressInfo(context.Context, *livekit.EgressInfo) (*google_protobuf2.Empty, error)

	GetIngressInfo(context.Context, *GetIngressInfoRequest) (*GetIngressInfoResponse, error)

	UpdateIngressState(context.Context, *UpdateIngressStateRequest) (*google_protobuf2.Empty, error)
}

// =======================
// IOInfo Server Interface
// =======================

type IOInfoServer interface {

	// Close and wait for pending RPCs to complete
	Shutdown()

	// Close immediately, without waiting for pending RPCs
	Kill()
}

// =============
// IOInfo Client
// =============

type iOInfoClient struct {
	client *psrpc.RPCClient
}

// NewIOInfoClient creates a psrpc client that implements the IOInfoClient interface.
func NewIOInfoClient(clientID string, bus psrpc.MessageBus, opts ...psrpc.ClientOption) (IOInfoClient, error) {
	rpcClient, err := psrpc.NewRPCClient("IOInfo", clientID, bus, opts...)
	if err != nil {
		return nil, err
	}

	return &iOInfoClient{
		client: rpcClient,
	}, nil
}

func (c *iOInfoClient) UpdateEgressInfo(ctx context.Context, req *livekit.EgressInfo, opts ...psrpc.RequestOption) (*google_protobuf2.Empty, error) {
	return psrpc.RequestSingle[*google_protobuf2.Empty](ctx, c.client, "UpdateEgressInfo", nil, true, req, opts...)
}

func (c *iOInfoClient) GetIngressInfo(ctx context.Context, req *GetIngressInfoRequest, opts ...psrpc.RequestOption) (*GetIngressInfoResponse, error) {
	return psrpc.RequestSingle[*GetIngressInfoResponse](ctx, c.client, "GetIngressInfo", nil, true, req, opts...)
}

func (c *iOInfoClient) UpdateIngressState(ctx context.Context, req *UpdateIngressStateRequest, opts ...psrpc.RequestOption) (*google_protobuf2.Empty, error) {
	return psrpc.RequestSingle[*google_protobuf2.Empty](ctx, c.client, "UpdateIngressState", nil, true, req, opts...)
}

// =============
// IOInfo Server
// =============

type iOInfoServer struct {
	svc IOInfoServerImpl
	rpc *psrpc.RPCServer
}

// NewIOInfoServer builds a RPCServer that will route requests
// to the corresponding method in the provided svc implementation.
func NewIOInfoServer(serverID string, svc IOInfoServerImpl, bus psrpc.MessageBus, opts ...psrpc.ServerOption) (IOInfoServer, error) {
	s := psrpc.NewRPCServer("IOInfo", serverID, bus, opts...)

	var err error
	err = psrpc.RegisterHandler(s, "UpdateEgressInfo", nil, svc.UpdateEgressInfo, nil, true)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	err = psrpc.RegisterHandler(s, "GetIngressInfo", nil, svc.GetIngressInfo, nil, true)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	err = psrpc.RegisterHandler(s, "UpdateIngressState", nil, svc.UpdateIngressState, nil, true)
	if err != nil {
		s.Close(false)
		return nil, err
	}

	return &iOInfoServer{
		svc: svc,
		rpc: s,
	}, nil
}

func (s *iOInfoServer) Shutdown() {
	s.rpc.Close(false)
}

func (s *iOInfoServer) Kill() {
	s.rpc.Close(true)
}

var psrpcFileDescriptor2 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x61, 0x4b, 0x02, 0x41,
	0x10, 0xc5, 0x4c, 0xc1, 0x31, 0x22, 0x36, 0x4f, 0xec, 0xa4, 0x08, 0x3f, 0x49, 0xc5, 0x1e, 0xd8,
	0x0f, 0x08, 0x02, 0x89, 0xa3, 0x20, 0x30, 0xfc, 0xd2, 0x97, 0x43, 0xcf, 0xf1, 0x5a, 0xee, 0xbc,
	0xdd, 0x76, 0xf7, 0x14, 0xff, 0x6e, 0xbf, 0x24, 0xee, 0x76, 0xbd, 0xcc, 0x14, 0xfa, 0xb4, 0xec,
	0x7b, 0x33, 0x6f, 0xde, 0xcc, 0x83, 0x13, 0x29, 0x42, 0x8f, 0x71, 0x2a, 0x24, 0xd7, 0x9c, 0x54,
	0xa5, 0x08, 0xdd, 0x56, 0xc2, 0x96, 0x18, 0x33, 0x1d, 0x60, 0x24, 0x51, 0x29, 0x43, 0xb9, 0xce,
	0x06, 0x65, 0xe9, 0x36, 0xdc, 0x8d, 0x38, 0x8f, 0x12, 0xf4, 0x8a, 0xdf, 0x34, 0x9b, 0x7b, 0xb8,
	0x10, 0x7a, 0x6d, 0xc8, 0xde, 0x18, 0x9c, 0x27, 0xd4, 0xbe, 0x69, 0xf0, 0xd3, 0x39, 0x1f, 0xe1,
	0x67, 0x86, 0x4a, 0x93, 0x4b, 0x00, 0x2b, 0x13, 0xb0, 0x59, 0xa7, 0x72, 0x5d, 0xe9, 0x37, 0x46,
	0x0d, 0x8b, 0xf8, 0xb3, 0x9c, 0x56, 0x5a, 0xe2, 0x64, 0x11, 0xc4, 0xb8, 0xee, 0x1c, 0x19, 0xda,
	0x20, 0xcf, 0xb8, 0xee, 0x71, 0x68, 0xef, 0xca, 0x2a, 0xc1, 0x53, 0x85, 0xa4, 0x0f, 0xc7, 0x2c,
	0x9d, 0xf3, 0x42, 0xb1, 0x39, 0x68, 0x51, 0xeb, 0x99, 0x6e, 0xd7, 0x16, 0x15, 0xa4, 0x05, 0x35,
	0xcd, 0x63, 0x4c, 0xad, 0xba, 0xf9, 0x10, 0x07, 0xea, 0x2b, 0x15, 0x64, 0x32, 0xe9, 0x54, 0x0d,
	0xbc, 0x52, 0x63, 0x99, 0xf4, 0x22, 0xb8, 0x18, 0x8b, 0xd9, 0x44, 0xa3, 0xd5, 0x79, 0xd3, 0x13,
	0x8d, 0xff, 0xdc, 0xe5, 0x16, 0x6a, 0x2a, 0x2f, 0x2f, 0x06, 0x35, 0x07, 0xce, 0xae, 0x27, 0xa3,
	0x65, 0x6a, 0x06, 0x5f, 0x15, 0xa8, 0xfb, 0xaf, 0xb9, 0x4d, 0xf2, 0x00, 0x67, 0x66, 0xe6, 0xb0,
	0xb4, 0x4e, 0xce, 0xcb, 0xe6, 0x1f, 0xd0, 0x6d, 0x53, 0x13, 0x01, 0xdd, 0x44, 0x40, 0x87, 0x79,
	0x04, 0xc4, 0x87, 0xd3, 0xdf, 0x57, 0x22, 0x2e, 0x95, 0x22, 0xa4, 0x7b, 0x13, 0x71, 0xbb, 0x7b,
	0x39, 0x7b, 0xd6, 0x17, 0x20, 0x7f, 0xf7, 0x27, 0x57, 0x45, 0xcb, 0xc1, 0xc3, 0x1c, 0x32, 0xf6,
	0x78, 0xf7, 0x7e, 0x13, 0x31, 0xfd, 0x91, 0x4d, 0x69, 0xc8, 0x17, 0x9e, 0xdd, 0xa8, 0x7c, 0x45,
	0x1c, 0x79, 0x0a, 0xe5, 0x92, 0x85, 0xe8, 0x49, 0x11, 0x4e, 0xeb, 0x45, 0xf7, 0xfd, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xc0, 0x02, 0x76, 0x1c, 0xa9, 0x02, 0x00, 0x00,
}
