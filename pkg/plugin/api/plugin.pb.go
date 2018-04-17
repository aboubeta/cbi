// Code generated by protoc-gen-gogo.
// source: plugin.proto
// DO NOT EDIT!

/*
Package cbi_plugin_v1 is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	InfoRequest
	InfoResponse
	SpecRequest
	SpecResponse
*/
package cbi_plugin_v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
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

type InfoRequest struct {
}

func (m *InfoRequest) Reset()                    { *m = InfoRequest{} }
func (m *InfoRequest) String() string            { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()               {}
func (*InfoRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{0} }

type InfoResponse struct {
	// Plugin labels.
	// This has nothing to do the labels of k8s objects of the pod.
	// e.g. pods, deployments, and services.
	//
	// However, it is still valid to implement a plugin that loads its labels from
	// k8s objects.
	//
	// See labels.go for the predefined labels.
	Labels map[string]string `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *InfoResponse) Reset()                    { *m = InfoResponse{} }
func (m *InfoResponse) String() string            { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()               {}
func (*InfoResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{1} }

func (m *InfoResponse) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SpecRequest struct {
	// JSON representation of CBI CRD BuildJob
	BuildJobJson []byte `protobuf:"bytes,1,opt,name=build_job_json,json=buildJobJson,proto3" json:"build_job_json,omitempty"`
}

func (m *SpecRequest) Reset()                    { *m = SpecRequest{} }
func (m *SpecRequest) String() string            { return proto.CompactTextString(m) }
func (*SpecRequest) ProtoMessage()               {}
func (*SpecRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{2} }

func (m *SpecRequest) GetBuildJobJson() []byte {
	if m != nil {
		return m.BuildJobJson
	}
	return nil
}

type SpecResponse struct {
	// JSON representation of Kubernetes PodTemplateSpec
	PodTemplateSpecJson []byte `protobuf:"bytes,1,opt,name=pod_template_spec_json,json=podTemplateSpecJson,proto3" json:"pod_template_spec_json,omitempty"`
}

func (m *SpecResponse) Reset()                    { *m = SpecResponse{} }
func (m *SpecResponse) String() string            { return proto.CompactTextString(m) }
func (*SpecResponse) ProtoMessage()               {}
func (*SpecResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{3} }

func (m *SpecResponse) GetPodTemplateSpecJson() []byte {
	if m != nil {
		return m.PodTemplateSpecJson
	}
	return nil
}

func init() {
	proto.RegisterType((*InfoRequest)(nil), "cbi.plugin.v1.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "cbi.plugin.v1.InfoResponse")
	proto.RegisterType((*SpecRequest)(nil), "cbi.plugin.v1.SpecRequest")
	proto.RegisterType((*SpecResponse)(nil), "cbi.plugin.v1.SpecResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Plugin service

type PluginClient interface {
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	Spec(ctx context.Context, in *SpecRequest, opts ...grpc.CallOption) (*SpecResponse, error)
}

type pluginClient struct {
	cc *grpc.ClientConn
}

func NewPluginClient(cc *grpc.ClientConn) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := grpc.Invoke(ctx, "/cbi.plugin.v1.Plugin/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) Spec(ctx context.Context, in *SpecRequest, opts ...grpc.CallOption) (*SpecResponse, error) {
	out := new(SpecResponse)
	err := grpc.Invoke(ctx, "/cbi.plugin.v1.Plugin/Spec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Plugin service

type PluginServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	Spec(context.Context, *SpecRequest) (*SpecResponse, error)
}

func RegisterPluginServer(s *grpc.Server, srv PluginServer) {
	s.RegisterService(&_Plugin_serviceDesc, srv)
}

func _Plugin_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbi.plugin.v1.Plugin/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_Spec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).Spec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cbi.plugin.v1.Plugin/Spec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).Spec(ctx, req.(*SpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Plugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cbi.plugin.v1.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _Plugin_Info_Handler,
		},
		{
			MethodName: "Spec",
			Handler:    _Plugin_Spec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "plugin.proto",
}

func (m *InfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *InfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InfoResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0xa
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovPlugin(uint64(len(k))) + 1 + len(v) + sovPlugin(uint64(len(v)))
			i = encodeVarintPlugin(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintPlugin(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintPlugin(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func (m *SpecRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.BuildJobJson) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.BuildJobJson)))
		i += copy(dAtA[i:], m.BuildJobJson)
	}
	return i, nil
}

func (m *SpecResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SpecResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.PodTemplateSpecJson) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintPlugin(dAtA, i, uint64(len(m.PodTemplateSpecJson)))
		i += copy(dAtA[i:], m.PodTemplateSpecJson)
	}
	return i, nil
}

func encodeFixed64Plugin(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Plugin(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintPlugin(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InfoRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *InfoResponse) Size() (n int) {
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovPlugin(uint64(len(k))) + 1 + len(v) + sovPlugin(uint64(len(v)))
			n += mapEntrySize + 1 + sovPlugin(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *SpecRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.BuildJobJson)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	return n
}

func (m *SpecResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.PodTemplateSpecJson)
	if l > 0 {
		n += 1 + l + sovPlugin(uint64(l))
	}
	return n
}

func sovPlugin(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPlugin(x uint64) (n int) {
	return sovPlugin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: InfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
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
func (m *InfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: InfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthPlugin
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPlugin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowPlugin
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthPlugin
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Labels[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Labels[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
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
func (m *SpecRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: SpecRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildJobJson", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuildJobJson = append(m.BuildJobJson[:0], dAtA[iNdEx:postIndex]...)
			if m.BuildJobJson == nil {
				m.BuildJobJson = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
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
func (m *SpecResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPlugin
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
			return fmt.Errorf("proto: SpecResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SpecResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PodTemplateSpecJson", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPlugin
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
				return ErrInvalidLengthPlugin
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PodTemplateSpecJson = append(m.PodTemplateSpecJson[:0], dAtA[iNdEx:postIndex]...)
			if m.PodTemplateSpecJson == nil {
				m.PodTemplateSpecJson = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPlugin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPlugin
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
func skipPlugin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
					return 0, ErrIntOverflowPlugin
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
				return 0, ErrInvalidLengthPlugin
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPlugin
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
				next, err := skipPlugin(dAtA[start:])
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
	ErrInvalidLengthPlugin = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPlugin   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("plugin.proto", fileDescriptorPlugin) }

var fileDescriptorPlugin = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x25, 0x9b, 0x0e, 0xfc, 0xda, 0x89, 0x44, 0x91, 0x51, 0x61, 0x8c, 0x22, 0xb8, 0x8b, 0x1d,
	0x6e, 0x17, 0xf5, 0x32, 0x50, 0x3c, 0x38, 0x3c, 0x48, 0xf5, 0x5e, 0x96, 0x2e, 0xab, 0x9d, 0x5d,
	0xbf, 0xb8, 0x24, 0x83, 0xfd, 0x85, 0xfd, 0x32, 0x8f, 0xfe, 0x04, 0xe9, 0x2f, 0x91, 0xa4, 0x3d,
	0x54, 0x26, 0xde, 0xbe, 0xf7, 0xf2, 0xde, 0xcb, 0xf7, 0x12, 0x70, 0x45, 0xa6, 0x93, 0x34, 0x0f,
	0xc4, 0x0a, 0x15, 0xd2, 0x76, 0xcc, 0xd2, 0xa0, 0x62, 0xd6, 0x57, 0xde, 0x65, 0x92, 0xaa, 0x37,
	0xcd, 0x82, 0x18, 0x97, 0x83, 0x04, 0x13, 0x1c, 0x58, 0x15, 0xd3, 0x73, 0x8b, 0x2c, 0xb0, 0x53,
	0xe9, 0xf6, 0xdb, 0xe0, 0x3c, 0xe6, 0x73, 0x0c, 0xf9, 0x87, 0xe6, 0x52, 0xf9, 0x5b, 0x02, 0x6e,
	0x89, 0xa5, 0xc0, 0x5c, 0x72, 0x3a, 0x86, 0x56, 0x36, 0x65, 0x3c, 0x93, 0x1d, 0xd2, 0x6b, 0xf6,
	0x9d, 0xe1, 0x45, 0xf0, 0xeb, 0xba, 0xa0, 0x2e, 0x0e, 0x9e, 0xac, 0xf2, 0x21, 0x57, 0xab, 0x4d,
	0x58, 0xd9, 0xbc, 0x1b, 0x70, 0x6a, 0x34, 0x3d, 0x82, 0xe6, 0x3b, 0xdf, 0x74, 0x48, 0x8f, 0xf4,
	0x0f, 0x42, 0x33, 0xd2, 0x13, 0xd8, 0x5f, 0x4f, 0x33, 0xcd, 0x3b, 0x0d, 0xcb, 0x95, 0xe0, 0xb6,
	0x71, 0x4d, 0xfc, 0x11, 0x38, 0x2f, 0x82, 0xc7, 0xd5, 0x6e, 0xf4, 0x1c, 0x0e, 0x99, 0x4e, 0xb3,
	0x59, 0xb4, 0x40, 0x16, 0x2d, 0x24, 0xe6, 0x36, 0xc5, 0x0d, 0x5d, 0xcb, 0x4e, 0x90, 0x4d, 0x24,
	0xe6, 0xfe, 0x3d, 0xb8, 0xa5, 0xa9, 0x2a, 0x30, 0x82, 0x53, 0x81, 0xb3, 0x48, 0xf1, 0xa5, 0xc8,
	0xa6, 0x8a, 0x47, 0x52, 0xf0, 0xb8, 0xee, 0x3e, 0x16, 0x38, 0x7b, 0xad, 0x0e, 0x8d, 0xd1, 0x84,
	0x0c, 0xb7, 0x04, 0x5a, 0xcf, 0xb6, 0x23, 0x1d, 0xc3, 0x9e, 0xe9, 0x48, 0xbd, 0x3f, 0x8b, 0xdb,
	0xcd, 0xbc, 0xb3, 0x7f, 0x1e, 0xc5, 0x04, 0x98, 0xdc, 0x9d, 0x80, 0x5a, 0xb5, 0x9d, 0x80, 0x7a,
	0x83, 0x3b, 0xf7, 0xb3, 0xe8, 0x92, 0xaf, 0xa2, 0x4b, 0xbe, 0x8b, 0x2e, 0x61, 0x2d, 0xfb, 0x6f,
	0xa3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x20, 0x30, 0x26, 0x05, 0x02, 0x00, 0x00,
}
