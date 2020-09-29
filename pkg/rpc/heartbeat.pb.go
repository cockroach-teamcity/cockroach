// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpc/heartbeat.proto

package rpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import github_com_cockroachdb_cockroach_pkg_util_uuid "github.com/cockroachdb/cockroach/pkg/util/uuid"
import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

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

// RemoteOffset keeps track of this client's estimate of its offset from a
// remote server. Uncertainty is the maximum error in the reading of this
// offset, so that the real offset should be in the interval
// [Offset - Uncertainty, Offset + Uncertainty]. If the last heartbeat timed
// out, Offset = 0.
//
// Offset and Uncertainty are measured using the remote clock reading technique
// described in http://se.inf.tu-dresden.de/pubs/papers/SRDS1994.pdf, page 6.
type RemoteOffset struct {
	// The estimated offset from the remote server, in nanoseconds.
	Offset int64 `protobuf:"varint,1,opt,name=offset" json:"offset"`
	// The maximum error of the measured offset, in nanoseconds.
	Uncertainty int64 `protobuf:"varint,2,opt,name=uncertainty" json:"uncertainty"`
	// Measurement time, in nanoseconds from unix epoch.
	MeasuredAt int64 `protobuf:"varint,3,opt,name=measured_at,json=measuredAt" json:"measured_at"`
}

func (m *RemoteOffset) Reset()      { *m = RemoteOffset{} }
func (*RemoteOffset) ProtoMessage() {}
func (*RemoteOffset) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_c48a0290adb87462, []int{0}
}
func (m *RemoteOffset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoteOffset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *RemoteOffset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteOffset.Merge(dst, src)
}
func (m *RemoteOffset) XXX_Size() int {
	return m.Size()
}
func (m *RemoteOffset) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteOffset.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteOffset proto.InternalMessageInfo

// A PingRequest specifies the string to echo in response.
// Fields are exported so that they will be serialized in the rpc call.
type PingRequest struct {
	// Echo this string with PingResponse.
	Ping string `protobuf:"bytes,1,opt,name=ping" json:"ping"`
	// The last offset the client measured with the server.
	Offset RemoteOffset `protobuf:"bytes,2,opt,name=offset" json:"offset"`
	// The address of the client.
	OriginAddr string `protobuf:"bytes,3,opt,name=origin_addr,json=originAddr" json:"origin_addr"`
	// The configured maximum clock offset (in nanoseconds) on the server.
	OriginMaxOffsetNanos int64 `protobuf:"varint,4,opt,name=origin_max_offset_nanos,json=originMaxOffsetNanos" json:"origin_max_offset_nanos"`
	// Cluster ID to prevent connections between nodes in different clusters.
	ClusterID     *github_com_cockroachdb_cockroach_pkg_util_uuid.UUID `protobuf:"bytes,5,opt,name=origin_cluster_id,json=originClusterId,customtype=github.com/cockroachdb/cockroach/pkg/util/uuid.UUID" json:"origin_cluster_id,omitempty"`
	ServerVersion roachpb.Version                                      `protobuf:"bytes,6,opt,name=server_version,json=serverVersion" json:"server_version"`
	// Node ID to prevent connections from being misrouted to an invalid node inside the cluster.
	TargetNodeID github_com_cockroachdb_cockroach_pkg_roachpb.NodeID `protobuf:"varint,7,opt,name=target_node_id,json=targetNodeId,customtype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"target_node_id"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_c48a0290adb87462, []int{1}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
}
func (m *PingRequest) XXX_Size() int {
	return m.Size()
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

// A PingResponse contains the echoed ping request string.
type PingResponse struct {
	// An echo of value sent with PingRequest.
	Pong          string          `protobuf:"bytes,1,opt,name=pong" json:"pong"`
	ServerTime    int64           `protobuf:"varint,2,opt,name=server_time,json=serverTime" json:"server_time"`
	ServerVersion roachpb.Version `protobuf:"bytes,3,opt,name=server_version,json=serverVersion" json:"server_version"`
	// Cluster name to prevent joining a new node to the wrong cluster.
	ClusterName string `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName" json:"cluster_name"`
	// Skip cluster name check if either side's name is empty / not configured.
	DisableClusterNameVerification bool `protobuf:"varint,5,opt,name=disable_cluster_name_verification,json=disableClusterNameVerification" json:"disable_cluster_name_verification"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_heartbeat_c48a0290adb87462, []int{2}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
}
func (m *PingResponse) XXX_Size() int {
	return m.Size()
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RemoteOffset)(nil), "cockroach.rpc.RemoteOffset")
	proto.RegisterType((*PingRequest)(nil), "cockroach.rpc.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "cockroach.rpc.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HeartbeatClient is the client API for Heartbeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type heartbeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatClient(cc *grpc.ClientConn) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/cockroach.rpc.Heartbeat/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServer is the server API for Heartbeat service.
type HeartbeatServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.rpc.Heartbeat/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.rpc.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Heartbeat_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/heartbeat.proto",
}

// TestingHeartbeatStreamClient is the client API for TestingHeartbeatStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestingHeartbeatStreamClient interface {
	PingStream(ctx context.Context, opts ...grpc.CallOption) (TestingHeartbeatStream_PingStreamClient, error)
}

type testingHeartbeatStreamClient struct {
	cc *grpc.ClientConn
}

func NewTestingHeartbeatStreamClient(cc *grpc.ClientConn) TestingHeartbeatStreamClient {
	return &testingHeartbeatStreamClient{cc}
}

func (c *testingHeartbeatStreamClient) PingStream(ctx context.Context, opts ...grpc.CallOption) (TestingHeartbeatStream_PingStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TestingHeartbeatStream_serviceDesc.Streams[0], "/cockroach.rpc.TestingHeartbeatStream/PingStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testingHeartbeatStreamPingStreamClient{stream}
	return x, nil
}

type TestingHeartbeatStream_PingStreamClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type testingHeartbeatStreamPingStreamClient struct {
	grpc.ClientStream
}

func (x *testingHeartbeatStreamPingStreamClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testingHeartbeatStreamPingStreamClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestingHeartbeatStreamServer is the server API for TestingHeartbeatStream service.
type TestingHeartbeatStreamServer interface {
	PingStream(TestingHeartbeatStream_PingStreamServer) error
}

func RegisterTestingHeartbeatStreamServer(s *grpc.Server, srv TestingHeartbeatStreamServer) {
	s.RegisterService(&_TestingHeartbeatStream_serviceDesc, srv)
}

func _TestingHeartbeatStream_PingStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestingHeartbeatStreamServer).PingStream(&testingHeartbeatStreamPingStreamServer{stream})
}

type TestingHeartbeatStream_PingStreamServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type testingHeartbeatStreamPingStreamServer struct {
	grpc.ServerStream
}

func (x *testingHeartbeatStreamPingStreamServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testingHeartbeatStreamPingStreamServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TestingHeartbeatStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.rpc.TestingHeartbeatStream",
	HandlerType: (*TestingHeartbeatStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PingStream",
			Handler:       _TestingHeartbeatStream_PingStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rpc/heartbeat.proto",
}

func (m *RemoteOffset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoteOffset) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.Offset))
	dAtA[i] = 0x10
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.Uncertainty))
	dAtA[i] = 0x18
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.MeasuredAt))
	return i, nil
}

func (m *PingRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(len(m.Ping)))
	i += copy(dAtA[i:], m.Ping)
	dAtA[i] = 0x12
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.Offset.Size()))
	n1, err := m.Offset.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(len(m.OriginAddr)))
	i += copy(dAtA[i:], m.OriginAddr)
	dAtA[i] = 0x20
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.OriginMaxOffsetNanos))
	if m.ClusterID != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHeartbeat(dAtA, i, uint64(m.ClusterID.Size()))
		n2, err := m.ClusterID.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	dAtA[i] = 0x32
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.ServerVersion.Size()))
	n3, err := m.ServerVersion.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x38
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.TargetNodeID))
	return i, nil
}

func (m *PingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(len(m.Pong)))
	i += copy(dAtA[i:], m.Pong)
	dAtA[i] = 0x10
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.ServerTime))
	dAtA[i] = 0x1a
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(m.ServerVersion.Size()))
	n4, err := m.ServerVersion.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x22
	i++
	i = encodeVarintHeartbeat(dAtA, i, uint64(len(m.ClusterName)))
	i += copy(dAtA[i:], m.ClusterName)
	dAtA[i] = 0x28
	i++
	if m.DisableClusterNameVerification {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i++
	return i, nil
}

func encodeVarintHeartbeat(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RemoteOffset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovHeartbeat(uint64(m.Offset))
	n += 1 + sovHeartbeat(uint64(m.Uncertainty))
	n += 1 + sovHeartbeat(uint64(m.MeasuredAt))
	return n
}

func (m *PingRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ping)
	n += 1 + l + sovHeartbeat(uint64(l))
	l = m.Offset.Size()
	n += 1 + l + sovHeartbeat(uint64(l))
	l = len(m.OriginAddr)
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 1 + sovHeartbeat(uint64(m.OriginMaxOffsetNanos))
	if m.ClusterID != nil {
		l = m.ClusterID.Size()
		n += 1 + l + sovHeartbeat(uint64(l))
	}
	l = m.ServerVersion.Size()
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 1 + sovHeartbeat(uint64(m.TargetNodeID))
	return n
}

func (m *PingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pong)
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 1 + sovHeartbeat(uint64(m.ServerTime))
	l = m.ServerVersion.Size()
	n += 1 + l + sovHeartbeat(uint64(l))
	l = len(m.ClusterName)
	n += 1 + l + sovHeartbeat(uint64(l))
	n += 2
	return n
}

func sovHeartbeat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeartbeat(x uint64) (n int) {
	return sovHeartbeat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RemoteOffset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
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
			return fmt.Errorf("proto: RemoteOffset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoteOffset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			m.Offset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Offset |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uncertainty", wireType)
			}
			m.Uncertainty = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uncertainty |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MeasuredAt", wireType)
			}
			m.MeasuredAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MeasuredAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func (m *PingRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
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
			return fmt.Errorf("proto: PingRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ping", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ping = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Offset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginMaxOffsetNanos", wireType)
			}
			m.OriginMaxOffsetNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OriginMaxOffsetNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cockroachdb_cockroach_pkg_util_uuid.UUID
			m.ClusterID = &v
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServerVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetNodeID", wireType)
			}
			m.TargetNodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetNodeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func (m *PingResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeartbeat
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
			return fmt.Errorf("proto: PingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pong", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pong = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerTime", wireType)
			}
			m.ServerTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ServerTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServerVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServerVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
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
				return ErrInvalidLengthHeartbeat
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableClusterNameVerification", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeartbeat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableClusterNameVerification = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHeartbeat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeartbeat
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
func skipHeartbeat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeartbeat
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
					return 0, ErrIntOverflowHeartbeat
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
					return 0, ErrIntOverflowHeartbeat
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
				return 0, ErrInvalidLengthHeartbeat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeartbeat
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
				next, err := skipHeartbeat(dAtA[start:])
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
	ErrInvalidLengthHeartbeat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeartbeat   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("rpc/heartbeat.proto", fileDescriptor_heartbeat_c48a0290adb87462) }

var fileDescriptor_heartbeat_c48a0290adb87462 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x41, 0x4f, 0xd4, 0x40,
	0x14, 0x6e, 0xd9, 0x05, 0x65, 0xb6, 0x60, 0x1c, 0x09, 0x36, 0x8b, 0xe9, 0xe2, 0x26, 0xe8, 0x9e,
	0x5a, 0x83, 0x27, 0xf5, 0xc4, 0x42, 0xa2, 0x84, 0xb8, 0x98, 0x15, 0x38, 0x78, 0x69, 0x66, 0x3b,
	0x8f, 0x32, 0x81, 0xce, 0x94, 0xe9, 0x94, 0xe0, 0xd1, 0x7f, 0x60, 0x3c, 0x79, 0xf4, 0xe7, 0x70,
	0xe4, 0x48, 0x3c, 0x10, 0x5d, 0xfe, 0x82, 0x3f, 0xc0, 0x4c, 0xa7, 0x85, 0x82, 0x1c, 0x0c, 0xb7,
	0x37, 0xef, 0x7d, 0xef, 0xbd, 0xef, 0xbd, 0xef, 0x0d, 0x7a, 0x24, 0xd3, 0x28, 0xd8, 0x03, 0x22,
	0xd5, 0x08, 0x88, 0xf2, 0x53, 0x29, 0x94, 0xc0, 0x33, 0x91, 0x88, 0xf6, 0xa5, 0x20, 0xd1, 0x9e,
	0x2f, 0xd3, 0xa8, 0x3d, 0x5f, 0x98, 0xe9, 0x28, 0x48, 0x40, 0x11, 0x4a, 0x14, 0x31, 0xb0, 0xf6,
	0x5c, 0x2c, 0x62, 0x51, 0x98, 0x81, 0xb6, 0x8c, 0xb7, 0xfb, 0xc5, 0x46, 0xce, 0x10, 0x12, 0xa1,
	0x60, 0x73, 0x77, 0x37, 0x03, 0x85, 0x9f, 0xa0, 0x29, 0x51, 0x58, 0xae, 0xbd, 0x68, 0xf7, 0x1a,
	0xfd, 0xe6, 0xc9, 0x79, 0xc7, 0x1a, 0x96, 0x3e, 0xfc, 0x0c, 0xb5, 0x72, 0x1e, 0x81, 0x54, 0x84,
	0x71, 0xf5, 0xd9, 0x9d, 0xa8, 0x41, 0xea, 0x01, 0xbc, 0x84, 0x5a, 0x09, 0x90, 0x2c, 0x97, 0x40,
	0x43, 0xa2, 0xdc, 0x46, 0x0d, 0x87, 0xaa, 0xc0, 0x8a, 0x7a, 0xdd, 0xfc, 0xfe, 0xa3, 0x63, 0x75,
	0xff, 0x34, 0x50, 0xeb, 0x03, 0xe3, 0xf1, 0x10, 0x0e, 0x73, 0xc8, 0x14, 0x76, 0x51, 0x33, 0x65,
	0x3c, 0x2e, 0x08, 0x4c, 0x97, 0x59, 0x85, 0x07, 0xbf, 0xba, 0x24, 0xa7, 0x3b, 0xb7, 0x96, 0x17,
	0xfc, 0x6b, 0xb3, 0xfb, 0xf5, 0x49, 0x6e, 0x30, 0x5f, 0x42, 0x2d, 0x21, 0x59, 0xcc, 0x78, 0x48,
	0x28, 0x95, 0x05, 0xa3, 0xaa, 0x36, 0x32, 0x81, 0x15, 0x4a, 0x25, 0x7e, 0x83, 0x1e, 0x97, 0xb0,
	0x84, 0x1c, 0x87, 0x26, 0x37, 0xe4, 0x84, 0x8b, 0xcc, 0x6d, 0xd6, 0x86, 0x98, 0x33, 0xa0, 0xf7,
	0xe4, 0xd8, 0x34, 0x1b, 0x68, 0x04, 0x4e, 0xd1, 0xc3, 0x32, 0x39, 0x3a, 0xc8, 0x33, 0x05, 0x32,
	0x64, 0xd4, 0x9d, 0x5c, 0xb4, 0x7b, 0x4e, 0x7f, 0xed, 0xe7, 0x79, 0xe7, 0x65, 0xcc, 0xd4, 0x5e,
	0x3e, 0xf2, 0x23, 0x91, 0x04, 0x97, 0xbc, 0xe9, 0xe8, 0xca, 0x0e, 0xd2, 0xfd, 0x38, 0xc8, 0x15,
	0x3b, 0x08, 0xf2, 0x9c, 0x51, 0x7f, 0x7b, 0x7b, 0x7d, 0x6d, 0x7c, 0xde, 0x99, 0x5e, 0x35, 0xc5,
	0xd6, 0xd7, 0x86, 0x0f, 0x4c, 0xf9, 0xca, 0x41, 0xf1, 0x5b, 0x34, 0x9b, 0x81, 0x3c, 0x02, 0x19,
	0x1e, 0x81, 0xcc, 0x98, 0xe0, 0xee, 0x54, 0xb1, 0x98, 0x76, 0x7d, 0x31, 0xe6, 0x1e, 0xfc, 0x1d,
	0x83, 0x28, 0x27, 0x98, 0x31, 0x79, 0xa5, 0x13, 0x1f, 0xa2, 0x59, 0x45, 0x64, 0xac, 0x87, 0x15,
	0x14, 0x34, 0xef, 0x7b, 0x8b, 0x76, 0x6f, 0xb2, 0xbf, 0xa1, 0xc1, 0xff, 0xcd, 0xbd, 0x6a, 0x35,
	0x10, 0x14, 0x0a, 0xee, 0xce, 0x56, 0x51, 0xd4, 0xbc, 0x87, 0x8e, 0xba, 0x7a, 0xd1, 0xee, 0xb7,
	0x09, 0xe4, 0x18, 0xd9, 0xb3, 0x54, 0xf0, 0x0c, 0x0a, 0xdd, 0xc5, 0x3f, 0xba, 0x0b, 0x1e, 0x6b,
	0xf1, 0xca, 0x31, 0x15, 0x4b, 0xe0, 0xda, 0xd9, 0x21, 0x13, 0xd8, 0x62, 0x09, 0xdc, 0xb2, 0x8d,
	0xc6, 0xdd, 0xb6, 0xf1, 0x1c, 0x39, 0x95, 0x82, 0x9c, 0x24, 0x50, 0x48, 0x5f, 0x31, 0x6a, 0x95,
	0x91, 0x01, 0x49, 0x00, 0x6f, 0xa2, 0xa7, 0x94, 0x65, 0x64, 0x74, 0x00, 0x61, 0x3d, 0x41, 0xf7,
	0x67, 0xbb, 0x2c, 0x22, 0x4a, 0x93, 0xd0, 0x17, 0x70, 0xbf, 0xcc, 0xf6, 0x4a, 0xf8, 0xea, 0x55,
	0x91, 0x9d, 0x1a, 0x76, 0x79, 0x80, 0xa6, 0xdf, 0x55, 0xff, 0x1b, 0xaf, 0xa0, 0xa6, 0x5e, 0x10,
	0x6e, 0xdf, 0x38, 0xf3, 0xda, 0x67, 0x69, 0x2f, 0xdc, 0x1a, 0x33, 0x1b, 0xed, 0x5a, 0xcb, 0x80,
	0xe6, 0xb7, 0x20, 0x53, 0x8c, 0xc7, 0x97, 0x65, 0x3f, 0x2a, 0x09, 0x24, 0xc1, 0x1b, 0x08, 0x69,
	0x6c, 0xf9, 0xba, 0x7b, 0x8b, 0x9e, 0xfd, 0xc2, 0xee, 0x2f, 0x9d, 0xfc, 0xf6, 0xac, 0x93, 0xb1,
	0x67, 0x9f, 0x8e, 0x3d, 0xfb, 0x6c, 0xec, 0xd9, 0xbf, 0xc6, 0x9e, 0xfd, 0xf5, 0xc2, 0xb3, 0x4e,
	0x2f, 0x3c, 0xeb, 0xec, 0xc2, 0xb3, 0x3e, 0x35, 0x64, 0x1a, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xfb, 0xa1, 0xf8, 0xf1, 0xc0, 0x04, 0x00, 0x00,
}
