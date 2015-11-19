// Code generated by protoc-gen-go.
// source: proto/v1/e2ekeys.proto
// DO NOT EDIT!

/*
Package google_security_e2ekeys_v1 is a generated protocol buffer package.

It is generated from these files:
	proto/v1/e2ekeys.proto

It has these top-level messages:
	HkpLookupRequest
	HttpResponse
*/
package google_security_e2ekeys_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_security_e2ekeys_v2 "github.com/google/e2e-key-server/proto/v2"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// HkpLookupRequest contains query parameters for retrieving PGP keys.
type HkpLookupRequest struct {
	// Op specifies the operation to be performed on the keyserver.
	// - "get" returns the pgp key specified in the search parameter.
	// - "index" returns 501 (not implemented).
	// - "vindex" returns 501 (not implemented).
	Op string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	// Search specifies the email address or key id being queried.
	Search string `protobuf:"bytes,2,opt,name=search" json:"search,omitempty"`
	// Options specifies what output format to use.
	// - "mr" machine readable will set the content type to "application/pgp-keys"
	// - other options will be ignored.
	Options string `protobuf:"bytes,3,opt,name=options" json:"options,omitempty"`
	// Exact specifies an exact match on search. Always on. If specified in the
	// URL, its value will be ignored.
	Exact string `protobuf:"bytes,4,opt,name=exact" json:"exact,omitempty"`
	// fingerprint is ignored.
	Fingerprint string `protobuf:"bytes,5,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *HkpLookupRequest) Reset()         { *m = HkpLookupRequest{} }
func (m *HkpLookupRequest) String() string { return proto.CompactTextString(m) }
func (*HkpLookupRequest) ProtoMessage()    {}

// HttpBody represents an http body.
type HttpResponse struct {
	// Header content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type" json:"content_type,omitempty"`
	// The http body itself.
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HttpResponse) Reset()         { *m = HttpResponse{} }
func (m *HttpResponse) String() string { return proto.CompactTextString(m) }
func (*HttpResponse) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for E2EKeyProxy service

type E2EKeyProxyClient interface {
	// GetEntry returns a user's current profile.
	GetEntry(ctx context.Context, in *google_security_e2ekeys_v2.GetEntryRequest, opts ...grpc.CallOption) (*google_security_e2ekeys_v2.Profile, error)
	HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error)
}

type e2EKeyProxyClient struct {
	cc *grpc.ClientConn
}

func NewE2EKeyProxyClient(cc *grpc.ClientConn) E2EKeyProxyClient {
	return &e2EKeyProxyClient{cc}
}

func (c *e2EKeyProxyClient) GetEntry(ctx context.Context, in *google_security_e2ekeys_v2.GetEntryRequest, opts ...grpc.CallOption) (*google_security_e2ekeys_v2.Profile, error) {
	out := new(google_security_e2ekeys_v2.Profile)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v1.E2EKeyProxy/GetEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *e2EKeyProxyClient) HkpLookup(ctx context.Context, in *HkpLookupRequest, opts ...grpc.CallOption) (*HttpResponse, error) {
	out := new(HttpResponse)
	err := grpc.Invoke(ctx, "/google.security.e2ekeys.v1.E2EKeyProxy/HkpLookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for E2EKeyProxy service

type E2EKeyProxyServer interface {
	// GetEntry returns a user's current profile.
	GetEntry(context.Context, *google_security_e2ekeys_v2.GetEntryRequest) (*google_security_e2ekeys_v2.Profile, error)
	HkpLookup(context.Context, *HkpLookupRequest) (*HttpResponse, error)
}

func RegisterE2EKeyProxyServer(s *grpc.Server, srv E2EKeyProxyServer) {
	s.RegisterService(&_E2EKeyProxy_serviceDesc, srv)
}

func _E2EKeyProxy_GetEntry_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(google_security_e2ekeys_v2.GetEntryRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyProxyServer).GetEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _E2EKeyProxy_HkpLookup_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(HkpLookupRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(E2EKeyProxyServer).HkpLookup(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _E2EKeyProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.security.e2ekeys.v1.E2EKeyProxy",
	HandlerType: (*E2EKeyProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEntry",
			Handler:    _E2EKeyProxy_GetEntry_Handler,
		},
		{
			MethodName: "HkpLookup",
			Handler:    _E2EKeyProxy_HkpLookup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}