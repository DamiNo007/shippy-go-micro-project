// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: vessel.proto

package vessel

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vessel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity  int32  `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight int32  `protobuf:"varint,3,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Available bool   `protobuf:"varint,5,opt,name=available,proto3" json:"available,omitempty"`
	OwnerId   string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
}

func (x *Vessel) Reset() {
	*x = Vessel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vessel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vessel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vessel) ProtoMessage() {}

func (x *Vessel) ProtoReflect() protoreflect.Message {
	mi := &file_vessel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vessel.ProtoReflect.Descriptor instead.
func (*Vessel) Descriptor() ([]byte, []int) {
	return file_vessel_proto_rawDescGZIP(), []int{0}
}

func (x *Vessel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vessel) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Vessel) GetMaxWeight() int32 {
	if x != nil {
		return x.MaxWeight
	}
	return 0
}

func (x *Vessel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Vessel) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

func (x *Vessel) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

type Specification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capacity  int32 `protobuf:"varint,1,opt,name=capacity,proto3" json:"capacity,omitempty"`
	MaxWeight int32 `protobuf:"varint,2,opt,name=max_weight,json=maxWeight,proto3" json:"max_weight,omitempty"`
}

func (x *Specification) Reset() {
	*x = Specification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vessel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Specification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Specification) ProtoMessage() {}

func (x *Specification) ProtoReflect() protoreflect.Message {
	mi := &file_vessel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Specification.ProtoReflect.Descriptor instead.
func (*Specification) Descriptor() ([]byte, []int) {
	return file_vessel_proto_rawDescGZIP(), []int{1}
}

func (x *Specification) GetCapacity() int32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Specification) GetMaxWeight() int32 {
	if x != nil {
		return x.MaxWeight
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vessel  *Vessel   `protobuf:"bytes,1,opt,name=vessel,proto3" json:"vessel,omitempty"`
	Vessels []*Vessel `protobuf:"bytes,2,rep,name=vessels,proto3" json:"vessels,omitempty"`
	Created bool      `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vessel_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_vessel_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_vessel_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetVessel() *Vessel {
	if x != nil {
		return x.Vessel
	}
	return nil
}

func (x *Response) GetVessels() []*Vessel {
	if x != nil {
		return x.Vessels
	}
	return nil
}

func (x *Response) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

var File_vessel_proto protoreflect.FileDescriptor

var file_vessel_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x65, 0x73,
	0x73, 0x65, 0x6c, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x56, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x78, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x6d, 0x61, 0x78, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0d, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x57, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x22, 0x90, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x06, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x76,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x56, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x52, 0x06, 0x76, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x07, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x56, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x52, 0x07, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x32, 0xad, 0x01, 0x0a, 0x0d, 0x56, 0x65, 0x73, 0x73, 0x65, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x22, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1d, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x65, 0x73, 0x73,
	0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x73, 0x72, 0x76, 0x2e, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x56, 0x65,
	0x73, 0x73, 0x65, 0x6c, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x73, 0x72, 0x76, 0x2e, 0x76, 0x65, 0x73, 0x73, 0x65, 0x6c, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x65, 0x73, 0x73, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vessel_proto_rawDescOnce sync.Once
	file_vessel_proto_rawDescData = file_vessel_proto_rawDesc
)

func file_vessel_proto_rawDescGZIP() []byte {
	file_vessel_proto_rawDescOnce.Do(func() {
		file_vessel_proto_rawDescData = protoimpl.X.CompressGZIP(file_vessel_proto_rawDescData)
	})
	return file_vessel_proto_rawDescData
}

var file_vessel_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vessel_proto_goTypes = []interface{}{
	(*Vessel)(nil),        // 0: go.micro.srv.vessel.Vessel
	(*Specification)(nil), // 1: go.micro.srv.vessel.Specification
	(*Response)(nil),      // 2: go.micro.srv.vessel.Response
}
var file_vessel_proto_depIdxs = []int32{
	0, // 0: go.micro.srv.vessel.Response.vessel:type_name -> go.micro.srv.vessel.Vessel
	0, // 1: go.micro.srv.vessel.Response.vessels:type_name -> go.micro.srv.vessel.Vessel
	1, // 2: go.micro.srv.vessel.VesselService.FindAvailable:input_type -> go.micro.srv.vessel.Specification
	0, // 3: go.micro.srv.vessel.VesselService.Create:input_type -> go.micro.srv.vessel.Vessel
	2, // 4: go.micro.srv.vessel.VesselService.FindAvailable:output_type -> go.micro.srv.vessel.Response
	2, // 5: go.micro.srv.vessel.VesselService.Create:output_type -> go.micro.srv.vessel.Response
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vessel_proto_init() }
func file_vessel_proto_init() {
	if File_vessel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vessel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vessel); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vessel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Specification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vessel_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vessel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vessel_proto_goTypes,
		DependencyIndexes: file_vessel_proto_depIdxs,
		MessageInfos:      file_vessel_proto_msgTypes,
	}.Build()
	File_vessel_proto = out.File
	file_vessel_proto_rawDesc = nil
	file_vessel_proto_goTypes = nil
	file_vessel_proto_depIdxs = nil
}
