// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualmachine.proto

package virtualmachine

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type VMData struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Vmname               string   `protobuf:"bytes,2,opt,name=vmname,proto3" json:"vmname,omitempty"`
	Vcpu                 int64    `protobuf:"varint,3,opt,name=vcpu,proto3" json:"vcpu,omitempty"`
	Vmem                 int64    `protobuf:"varint,4,opt,name=vmem,proto3" json:"vmem,omitempty"`
	StoragePath          string   `protobuf:"bytes,5,opt,name=storage_path,json=storagePath,proto3" json:"storage_path,omitempty"`
	Storage              int64    `protobuf:"varint,6,opt,name=storage,proto3" json:"storage,omitempty"`
	CdromPath            string   `protobuf:"bytes,7,opt,name=cdrom_path,json=cdromPath,proto3" json:"cdrom_path,omitempty"`
	Vnet                 string   `protobuf:"bytes,8,opt,name=vnet,proto3" json:"vnet,omitempty"`
	Vnc                  int64    `protobuf:"varint,9,opt,name=vnc,proto3" json:"vnc,omitempty"`
	Change               bool     `protobuf:"varint,10,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMData) Reset()         { *m = VMData{} }
func (m *VMData) String() string { return proto.CompactTextString(m) }
func (*VMData) ProtoMessage()    {}
func (*VMData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f620cdfcc6a9426c, []int{0}
}

func (m *VMData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMData.Unmarshal(m, b)
}
func (m *VMData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMData.Marshal(b, m, deterministic)
}
func (m *VMData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMData.Merge(m, src)
}
func (m *VMData) XXX_Size() int {
	return xxx_messageInfo_VMData.Size(m)
}
func (m *VMData) XXX_DiscardUnknown() {
	xxx_messageInfo_VMData.DiscardUnknown(m)
}

var xxx_messageInfo_VMData proto.InternalMessageInfo

func (m *VMData) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VMData) GetVmname() string {
	if m != nil {
		return m.Vmname
	}
	return ""
}

func (m *VMData) GetVcpu() int64 {
	if m != nil {
		return m.Vcpu
	}
	return 0
}

func (m *VMData) GetVmem() int64 {
	if m != nil {
		return m.Vmem
	}
	return 0
}

func (m *VMData) GetStoragePath() string {
	if m != nil {
		return m.StoragePath
	}
	return ""
}

func (m *VMData) GetStorage() int64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *VMData) GetCdromPath() string {
	if m != nil {
		return m.CdromPath
	}
	return ""
}

func (m *VMData) GetVnet() string {
	if m != nil {
		return m.Vnet
	}
	return ""
}

func (m *VMData) GetVnc() int64 {
	if m != nil {
		return m.Vnc
	}
	return 0
}

func (m *VMData) GetChange() bool {
	if m != nil {
		return m.Change
	}
	return false
}

type VMName struct {
	Vmname               string   `protobuf:"bytes,1,opt,name=vmname,proto3" json:"vmname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMName) Reset()         { *m = VMName{} }
func (m *VMName) String() string { return proto.CompactTextString(m) }
func (*VMName) ProtoMessage()    {}
func (*VMName) Descriptor() ([]byte, []int) {
	return fileDescriptor_f620cdfcc6a9426c, []int{1}
}

func (m *VMName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMName.Unmarshal(m, b)
}
func (m *VMName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMName.Marshal(b, m, deterministic)
}
func (m *VMName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMName.Merge(m, src)
}
func (m *VMName) XXX_Size() int {
	return xxx_messageInfo_VMName.Size(m)
}
func (m *VMName) XXX_DiscardUnknown() {
	xxx_messageInfo_VMName.DiscardUnknown(m)
}

var xxx_messageInfo_VMName proto.InternalMessageInfo

func (m *VMName) GetVmname() string {
	if m != nil {
		return m.Vmname
	}
	return ""
}

type VMDataResponse struct {
	VMDatas              []*VMData `protobuf:"bytes,1,rep,name=VMDatas,proto3" json:"VMDatas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VMDataResponse) Reset()         { *m = VMDataResponse{} }
func (m *VMDataResponse) String() string { return proto.CompactTextString(m) }
func (*VMDataResponse) ProtoMessage()    {}
func (*VMDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f620cdfcc6a9426c, []int{2}
}

func (m *VMDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMDataResponse.Unmarshal(m, b)
}
func (m *VMDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMDataResponse.Marshal(b, m, deterministic)
}
func (m *VMDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMDataResponse.Merge(m, src)
}
func (m *VMDataResponse) XXX_Size() int {
	return xxx_messageInfo_VMDataResponse.Size(m)
}
func (m *VMDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VMDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VMDataResponse proto.InternalMessageInfo

func (m *VMDataResponse) GetVMDatas() []*VMData {
	if m != nil {
		return m.VMDatas
	}
	return nil
}

type VMID struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMID) Reset()         { *m = VMID{} }
func (m *VMID) String() string { return proto.CompactTextString(m) }
func (*VMID) ProtoMessage()    {}
func (*VMID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f620cdfcc6a9426c, []int{3}
}

func (m *VMID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMID.Unmarshal(m, b)
}
func (m *VMID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMID.Marshal(b, m, deterministic)
}
func (m *VMID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMID.Merge(m, src)
}
func (m *VMID) XXX_Size() int {
	return xxx_messageInfo_VMID.Size(m)
}
func (m *VMID) XXX_DiscardUnknown() {
	xxx_messageInfo_VMID.DiscardUnknown(m)
}

var xxx_messageInfo_VMID proto.InternalMessageInfo

func (m *VMID) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Result struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_f620cdfcc6a9426c, []int{4}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*VMData)(nil), "VMData")
	proto.RegisterType((*VMName)(nil), "VMName")
	proto.RegisterType((*VMDataResponse)(nil), "VMDataResponse")
	proto.RegisterType((*VMID)(nil), "VMID")
	proto.RegisterType((*Result)(nil), "Result")
}

func init() { proto.RegisterFile("virtualmachine.proto", fileDescriptor_f620cdfcc6a9426c) }

var fileDescriptor_f620cdfcc6a9426c = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x8b, 0xa3, 0x40,
	0x10, 0x4d, 0x6b, 0xe2, 0x47, 0x65, 0x09, 0x4b, 0xb3, 0x84, 0xde, 0xec, 0x97, 0xf1, 0xe4, 0xc9,
	0x43, 0xf2, 0x13, 0x36, 0xb0, 0xec, 0xc1, 0x61, 0x30, 0xe0, 0x75, 0xe8, 0x31, 0x45, 0x14, 0xd2,
	0x2a, 0xda, 0xfa, 0xd7, 0xe6, 0x7f, 0xcd, 0x2f, 0x18, 0xba, 0x34, 0x4c, 0x66, 0xc8, 0xad, 0xde,
	0x7b, 0xfd, 0xea, 0xf1, 0x8a, 0x86, 0x6f, 0x43, 0xd9, 0xea, 0x5e, 0x5e, 0x94, 0xcc, 0x8b, 0xb2,
	0xc2, 0xb8, 0x69, 0x6b, 0x5d, 0x87, 0xaf, 0x0c, 0x9c, 0x2c, 0x39, 0x48, 0x2d, 0xf9, 0x0a, 0xac,
	0xf2, 0x24, 0x58, 0xc0, 0x22, 0x3b, 0xb5, 0xca, 0x13, 0x5f, 0x83, 0x33, 0xa8, 0x4a, 0x2a, 0x14,
	0x56, 0xc0, 0x22, 0x3f, 0x9d, 0x10, 0xe7, 0x30, 0x1f, 0xf2, 0xa6, 0x17, 0x36, 0xbd, 0xa4, 0x99,
	0x38, 0x85, 0x4a, 0xcc, 0x27, 0x4e, 0xa1, 0xe2, 0x5b, 0xf8, 0xd2, 0xe9, 0xba, 0x95, 0x67, 0x7c,
	0x6a, 0xa4, 0x2e, 0xc4, 0x82, 0xb6, 0x2c, 0x27, 0xee, 0x51, 0xea, 0x82, 0x0b, 0x70, 0x27, 0x28,
	0x1c, 0x72, 0x5e, 0x21, 0xff, 0x05, 0x90, 0x9f, 0xda, 0x5a, 0x8d, 0x56, 0x97, 0xac, 0x3e, 0x31,
	0x64, 0x34, 0x79, 0x15, 0x6a, 0xe1, 0x91, 0x40, 0x33, 0xff, 0x0a, 0xf6, 0x50, 0xe5, 0xc2, 0xa7,
	0x45, 0x66, 0x34, 0x0d, 0xf2, 0x42, 0x56, 0x67, 0x14, 0x10, 0xb0, 0xc8, 0x4b, 0x27, 0x14, 0x06,
	0xa6, 0xf3, 0x83, 0xe9, 0xf2, 0xde, 0x91, 0xdd, 0x76, 0x0c, 0xf7, 0xb0, 0x1a, 0xaf, 0x92, 0x62,
	0xd7, 0xd4, 0x55, 0x87, 0x7c, 0x0b, 0xee, 0xc8, 0x74, 0x82, 0x05, 0x76, 0xb4, 0xdc, 0xb9, 0xf1,
	0xf4, 0xe2, 0xca, 0x87, 0x6b, 0x98, 0x67, 0xc9, 0xff, 0xc3, 0xe7, 0x43, 0x9a, 0xb8, 0x14, 0xbb,
	0xfe, 0xa2, 0x4d, 0x5c, 0xa7, 0xa5, 0xee, 0x3b, 0x52, 0xbd, 0x74, 0x42, 0xbb, 0x17, 0x06, 0x56,
	0x96, 0xf0, 0xdf, 0xe0, 0xfd, 0x6d, 0x51, 0x6a, 0xcc, 0x12, 0x7e, 0x5d, 0xbf, 0x71, 0xe3, 0xd1,
	0x1c, 0xce, 0xf8, 0x4f, 0xf0, 0x0e, 0x78, 0x41, 0xd2, 0x17, 0xb1, 0xc9, 0xba, 0x55, 0x7f, 0x80,
	0x7b, 0xd4, 0xb2, 0xd5, 0x77, 0xc5, 0x0d, 0x38, 0x47, 0x5d, 0x37, 0x77, 0xb5, 0xef, 0xb0, 0xf8,
	0x87, 0x1f, 0x6c, 0x63, 0x74, 0x38, 0xe3, 0x7f, 0xc0, 0x27, 0x89, 0x8e, 0x65, 0x78, 0x33, 0xdc,
	0x3c, 0x78, 0x76, 0xe8, 0x1b, 0xed, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x48, 0xda, 0xd1, 0x77,
	0x5e, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VMClient is the client API for VM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VMClient interface {
	CreateVM(ctx context.Context, in *VMData, opts ...grpc.CallOption) (*Result, error)
	DeleteVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*Result, error)
	StartVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*Result, error)
	StopVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*Result, error)
	GetVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*VMData, error)
	GetVMName(ctx context.Context, in *VMName, opts ...grpc.CallOption) (*VMData, error)
}

type vMClient struct {
	cc grpc.ClientConnInterface
}

func NewVMClient(cc grpc.ClientConnInterface) VMClient {
	return &vMClient{cc}
}

func (c *vMClient) CreateVM(ctx context.Context, in *VMData, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/VM/CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) DeleteVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/VM/DeleteVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) StartVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/VM/StartVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) StopVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/VM/StopVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) GetVM(ctx context.Context, in *VMID, opts ...grpc.CallOption) (*VMData, error) {
	out := new(VMData)
	err := c.cc.Invoke(ctx, "/VM/GetVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) GetVMName(ctx context.Context, in *VMName, opts ...grpc.CallOption) (*VMData, error) {
	out := new(VMData)
	err := c.cc.Invoke(ctx, "/VM/GetVMName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VMServer is the server API for VM service.
type VMServer interface {
	CreateVM(context.Context, *VMData) (*Result, error)
	DeleteVM(context.Context, *VMID) (*Result, error)
	StartVM(context.Context, *VMID) (*Result, error)
	StopVM(context.Context, *VMID) (*Result, error)
	GetVM(context.Context, *VMID) (*VMData, error)
	GetVMName(context.Context, *VMName) (*VMData, error)
}

// UnimplementedVMServer can be embedded to have forward compatible implementations.
type UnimplementedVMServer struct {
}

func (*UnimplementedVMServer) CreateVM(ctx context.Context, req *VMData) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVM not implemented")
}
func (*UnimplementedVMServer) DeleteVM(ctx context.Context, req *VMID) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVM not implemented")
}
func (*UnimplementedVMServer) StartVM(ctx context.Context, req *VMID) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartVM not implemented")
}
func (*UnimplementedVMServer) StopVM(ctx context.Context, req *VMID) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopVM not implemented")
}
func (*UnimplementedVMServer) GetVM(ctx context.Context, req *VMID) (*VMData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVM not implemented")
}
func (*UnimplementedVMServer) GetVMName(ctx context.Context, req *VMName) (*VMData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVMName not implemented")
}

func RegisterVMServer(s *grpc.Server, srv VMServer) {
	s.RegisterService(&_VM_serviceDesc, srv)
}

func _VM_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).CreateVM(ctx, req.(*VMData))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_DeleteVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).DeleteVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/DeleteVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).DeleteVM(ctx, req.(*VMID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_StartVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).StartVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/StartVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).StartVM(ctx, req.(*VMID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_StopVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).StopVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/StopVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).StopVM(ctx, req.(*VMID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_GetVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).GetVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/GetVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).GetVM(ctx, req.(*VMID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_GetVMName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).GetVMName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/GetVMName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).GetVMName(ctx, req.(*VMName))
	}
	return interceptor(ctx, in, info, handler)
}

var _VM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "VM",
	HandlerType: (*VMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVM",
			Handler:    _VM_CreateVM_Handler,
		},
		{
			MethodName: "DeleteVM",
			Handler:    _VM_DeleteVM_Handler,
		},
		{
			MethodName: "StartVM",
			Handler:    _VM_StartVM_Handler,
		},
		{
			MethodName: "StopVM",
			Handler:    _VM_StopVM_Handler,
		},
		{
			MethodName: "GetVM",
			Handler:    _VM_GetVM_Handler,
		},
		{
			MethodName: "GetVMName",
			Handler:    _VM_GetVMName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualmachine.proto",
}
