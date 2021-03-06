// ---------------------------------------------------------------------------
// Protofilename : huawei-flow-recognition.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: huaweiV8R12-flow.proto

package huaweiV8R12_flow_recognition

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

type FlowRecognition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streaminfos *FlowRecognition_Streaminfos `protobuf:"bytes,1,opt,name=streaminfos,proto3" json:"streaminfos,omitempty"`
}

func (x *FlowRecognition) Reset() {
	*x = FlowRecognition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R12_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowRecognition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowRecognition) ProtoMessage() {}

func (x *FlowRecognition) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R12_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowRecognition.ProtoReflect.Descriptor instead.
func (*FlowRecognition) Descriptor() ([]byte, []int) {
	return file_huaweiV8R12_flow_proto_rawDescGZIP(), []int{0}
}

func (x *FlowRecognition) GetStreaminfos() *FlowRecognition_Streaminfos {
	if x != nil {
		return x.Streaminfos
	}
	return nil
}

type FlowRecognition_Streaminfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Streaminfo []*FlowRecognition_Streaminfos_Streaminfo `protobuf:"bytes,1,rep,name=streaminfo,proto3" json:"streaminfo,omitempty"`
}

func (x *FlowRecognition_Streaminfos) Reset() {
	*x = FlowRecognition_Streaminfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R12_flow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowRecognition_Streaminfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowRecognition_Streaminfos) ProtoMessage() {}

func (x *FlowRecognition_Streaminfos) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R12_flow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowRecognition_Streaminfos.ProtoReflect.Descriptor instead.
func (*FlowRecognition_Streaminfos) Descriptor() ([]byte, []int) {
	return file_huaweiV8R12_flow_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FlowRecognition_Streaminfos) GetStreaminfo() []*FlowRecognition_Streaminfos_Streaminfo {
	if x != nil {
		return x.Streaminfo
	}
	return nil
}

type FlowRecognition_Streaminfos_Streaminfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SrcMac        string `protobuf:"bytes,1,opt,name=src_mac,json=src-mac,proto3" json:"src_mac,omitempty"`
	DstMac        string `protobuf:"bytes,2,opt,name=dst_mac,json=dst-mac,proto3" json:"dst_mac,omitempty"`
	SrcIpAddr     string `protobuf:"bytes,3,opt,name=src_ip_addr,json=src-ip-addr,proto3" json:"src_ip_addr,omitempty"`
	DstIpAddr     string `protobuf:"bytes,4,opt,name=dst_ip_addr,json=dst-ip-addr,proto3" json:"dst_ip_addr,omitempty"`
	SrcPort       uint32 `protobuf:"varint,5,opt,name=src_port,json=src-port,proto3" json:"src_port,omitempty"`
	DstPort       uint32 `protobuf:"varint,6,opt,name=dst_port,json=dst-port,proto3" json:"dst_port,omitempty"`
	Protocol      uint32 `protobuf:"varint,7,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Direction     uint32 `protobuf:"varint,8,opt,name=direction,proto3" json:"direction,omitempty"`
	IfName        string `protobuf:"bytes,9,opt,name=if_name,json=if-name,proto3" json:"if_name,omitempty"`
	TimeStampSec  uint32 `protobuf:"varint,10,opt,name=time_stamp_sec,json=time-stamp-sec,proto3" json:"time_stamp_sec,omitempty"`
	TimeStampNsec uint32 `protobuf:"varint,11,opt,name=time_stamp_nsec,json=time-stamp-nsec,proto3" json:"time_stamp_nsec,omitempty"`
	PacketNum     uint64 `protobuf:"varint,12,opt,name=packet_num,json=packet-num,proto3" json:"packet_num,omitempty"`
	BytesNum      uint64 `protobuf:"varint,13,opt,name=bytes_num,json=bytes-num,proto3" json:"bytes_num,omitempty"`
}

func (x *FlowRecognition_Streaminfos_Streaminfo) Reset() {
	*x = FlowRecognition_Streaminfos_Streaminfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R12_flow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowRecognition_Streaminfos_Streaminfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowRecognition_Streaminfos_Streaminfo) ProtoMessage() {}

func (x *FlowRecognition_Streaminfos_Streaminfo) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R12_flow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowRecognition_Streaminfos_Streaminfo.ProtoReflect.Descriptor instead.
func (*FlowRecognition_Streaminfos_Streaminfo) Descriptor() ([]byte, []int) {
	return file_huaweiV8R12_flow_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetSrcMac() string {
	if x != nil {
		return x.SrcMac
	}
	return ""
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetDstMac() string {
	if x != nil {
		return x.DstMac
	}
	return ""
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetSrcIpAddr() string {
	if x != nil {
		return x.SrcIpAddr
	}
	return ""
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetDstIpAddr() string {
	if x != nil {
		return x.DstIpAddr
	}
	return ""
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetSrcPort() uint32 {
	if x != nil {
		return x.SrcPort
	}
	return 0
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetDstPort() uint32 {
	if x != nil {
		return x.DstPort
	}
	return 0
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetProtocol() uint32 {
	if x != nil {
		return x.Protocol
	}
	return 0
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetDirection() uint32 {
	if x != nil {
		return x.Direction
	}
	return 0
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetIfName() string {
	if x != nil {
		return x.IfName
	}
	return ""
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetTimeStampSec() uint32 {
	if x != nil {
		return x.TimeStampSec
	}
	return 0
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetTimeStampNsec() uint32 {
	if x != nil {
		return x.TimeStampNsec
	}
	return 0
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetPacketNum() uint64 {
	if x != nil {
		return x.PacketNum
	}
	return 0
}

func (x *FlowRecognition_Streaminfos_Streaminfo) GetBytesNum() uint64 {
	if x != nil {
		return x.BytesNum
	}
	return 0
}

var File_huaweiV8R12_flow_proto protoreflect.FileDescriptor

var file_huaweiV8R12_flow_proto_rawDesc = []byte{
	0x0a, 0x16, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x32, 0x2d, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69,
	0x56, 0x38, 0x52, 0x31, 0x32, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x87, 0x05, 0x0a, 0x0f, 0x46, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5b, 0x0a, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x32, 0x5f, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46,
	0x6c, 0x6f, 0x77, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x1a, 0x96, 0x04, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x64, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x68, 0x75,
	0x61, 0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x32, 0x5f, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x52,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x66,
	0x6f, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0xa0, 0x03,
	0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x72, 0x63, 0x5f, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x72, 0x63, 0x2d, 0x6d, 0x61, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x73, 0x74, 0x5f, 0x6d, 0x61,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x73, 0x74, 0x2d, 0x6d, 0x61, 0x63,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x72, 0x63, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x72, 0x63, 0x2d, 0x69, 0x70, 0x2d, 0x61, 0x64,
	0x64, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x73, 0x74, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x73, 0x74, 0x2d, 0x69, 0x70, 0x2d,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x72, 0x63, 0x5f, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x72, 0x63, 0x2d, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x64, 0x73, 0x74, 0x2d, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x66, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x66, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2d, 0x73, 0x65, 0x63, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x73, 0x65, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2d, 0x6e, 0x73,
	0x65, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2d, 0x6e,
	0x75, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x2d, 0x6e, 0x75, 0x6d,
	0x42, 0x1e, 0x5a, 0x1c, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x32, 0x5f,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_huaweiV8R12_flow_proto_rawDescOnce sync.Once
	file_huaweiV8R12_flow_proto_rawDescData = file_huaweiV8R12_flow_proto_rawDesc
)

func file_huaweiV8R12_flow_proto_rawDescGZIP() []byte {
	file_huaweiV8R12_flow_proto_rawDescOnce.Do(func() {
		file_huaweiV8R12_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_huaweiV8R12_flow_proto_rawDescData)
	})
	return file_huaweiV8R12_flow_proto_rawDescData
}

var file_huaweiV8R12_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_huaweiV8R12_flow_proto_goTypes = []interface{}{
	(*FlowRecognition)(nil),                        // 0: huaweiV8R12_flow_recognition.FlowRecognition
	(*FlowRecognition_Streaminfos)(nil),            // 1: huaweiV8R12_flow_recognition.FlowRecognition.Streaminfos
	(*FlowRecognition_Streaminfos_Streaminfo)(nil), // 2: huaweiV8R12_flow_recognition.FlowRecognition.Streaminfos.Streaminfo
}
var file_huaweiV8R12_flow_proto_depIdxs = []int32{
	1, // 0: huaweiV8R12_flow_recognition.FlowRecognition.streaminfos:type_name -> huaweiV8R12_flow_recognition.FlowRecognition.Streaminfos
	2, // 1: huaweiV8R12_flow_recognition.FlowRecognition.Streaminfos.streaminfo:type_name -> huaweiV8R12_flow_recognition.FlowRecognition.Streaminfos.Streaminfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_huaweiV8R12_flow_proto_init() }
func file_huaweiV8R12_flow_proto_init() {
	if File_huaweiV8R12_flow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_huaweiV8R12_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowRecognition); i {
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
		file_huaweiV8R12_flow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowRecognition_Streaminfos); i {
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
		file_huaweiV8R12_flow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowRecognition_Streaminfos_Streaminfo); i {
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
			RawDescriptor: file_huaweiV8R12_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_huaweiV8R12_flow_proto_goTypes,
		DependencyIndexes: file_huaweiV8R12_flow_proto_depIdxs,
		MessageInfos:      file_huaweiV8R12_flow_proto_msgTypes,
	}.Build()
	File_huaweiV8R12_flow_proto = out.File
	file_huaweiV8R12_flow_proto_rawDesc = nil
	file_huaweiV8R12_flow_proto_goTypes = nil
	file_huaweiV8R12_flow_proto_depIdxs = nil
}
