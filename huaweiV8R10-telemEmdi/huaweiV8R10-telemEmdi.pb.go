// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: huaweiV8R10-telemEmdi.proto

package huaweiV8R10_telemEmdi

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

type TelemEmdi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmdiTelemReps *TelemEmdi_EmdiTelemReps `protobuf:"bytes,1,opt,name=emdiTelemReps,proto3" json:"emdiTelemReps,omitempty"`
	EmdiTelemRtps *TelemEmdi_EmdiTelemRtps `protobuf:"bytes,2,opt,name=emdiTelemRtps,proto3" json:"emdiTelemRtps,omitempty"`
}

func (x *TelemEmdi) Reset() {
	*x = TelemEmdi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemEmdi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemEmdi) ProtoMessage() {}

func (x *TelemEmdi) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemEmdi.ProtoReflect.Descriptor instead.
func (*TelemEmdi) Descriptor() ([]byte, []int) {
	return file_huaweiV8R10_telemEmdi_proto_rawDescGZIP(), []int{0}
}

func (x *TelemEmdi) GetEmdiTelemReps() *TelemEmdi_EmdiTelemReps {
	if x != nil {
		return x.EmdiTelemReps
	}
	return nil
}

func (x *TelemEmdi) GetEmdiTelemRtps() *TelemEmdi_EmdiTelemRtps {
	if x != nil {
		return x.EmdiTelemRtps
	}
	return nil
}

type TelemEmdi_EmdiTelemReps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmdiTelemRep []*TelemEmdi_EmdiTelemReps_EmdiTelemRep `protobuf:"bytes,1,rep,name=emdiTelemRep,proto3" json:"emdiTelemRep,omitempty"`
}

func (x *TelemEmdi_EmdiTelemReps) Reset() {
	*x = TelemEmdi_EmdiTelemReps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemEmdi_EmdiTelemReps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemEmdi_EmdiTelemReps) ProtoMessage() {}

func (x *TelemEmdi_EmdiTelemReps) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemEmdi_EmdiTelemReps.ProtoReflect.Descriptor instead.
func (*TelemEmdi_EmdiTelemReps) Descriptor() ([]byte, []int) {
	return file_huaweiV8R10_telemEmdi_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TelemEmdi_EmdiTelemReps) GetEmdiTelemRep() []*TelemEmdi_EmdiTelemReps_EmdiTelemRep {
	if x != nil {
		return x.EmdiTelemRep
	}
	return nil
}

type TelemEmdi_EmdiTelemRtps struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmdiTelemRtp []*TelemEmdi_EmdiTelemRtps_EmdiTelemRtp `protobuf:"bytes,1,rep,name=emdiTelemRtp,proto3" json:"emdiTelemRtp,omitempty"`
}

func (x *TelemEmdi_EmdiTelemRtps) Reset() {
	*x = TelemEmdi_EmdiTelemRtps{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemEmdi_EmdiTelemRtps) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemEmdi_EmdiTelemRtps) ProtoMessage() {}

func (x *TelemEmdi_EmdiTelemRtps) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemEmdi_EmdiTelemRtps.ProtoReflect.Descriptor instead.
func (*TelemEmdi_EmdiTelemRtps) Descriptor() ([]byte, []int) {
	return file_huaweiV8R10_telemEmdi_proto_rawDescGZIP(), []int{0, 1}
}

func (x *TelemEmdi_EmdiTelemRtps) GetEmdiTelemRtp() []*TelemEmdi_EmdiTelemRtps_EmdiTelemRtp {
	if x != nil {
		return x.EmdiTelemRtp
	}
	return nil
}

type TelemEmdi_EmdiTelemReps_EmdiTelemRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BdId                     uint32 `protobuf:"varint,1,opt,name=bdId,proto3" json:"bdId,omitempty"`
	ChannelGroupAddress      uint32 `protobuf:"varint,2,opt,name=channelGroupAddress,proto3" json:"channelGroupAddress,omitempty"`
	ChannelName              string `protobuf:"bytes,3,opt,name=channelName,proto3" json:"channelName,omitempty"`
	ChannelSourceAddress     uint32 `protobuf:"varint,4,opt,name=channelSourceAddress,proto3" json:"channelSourceAddress,omitempty"`
	ChannelVpnType           uint32 `protobuf:"varint,5,opt,name=channelVpnType,proto3" json:"channelVpnType,omitempty"`
	MonitorPeriod            uint32 `protobuf:"varint,6,opt,name=monitorPeriod,proto3" json:"monitorPeriod,omitempty"`
	MonitorTime              uint64 `protobuf:"varint,7,opt,name=monitorTime,proto3" json:"monitorTime,omitempty"`
	ReportPeriod             uint32 `protobuf:"varint,8,opt,name=reportPeriod,proto3" json:"reportPeriod,omitempty"`
	ReportTime               uint64 `protobuf:"varint,9,opt,name=reportTime,proto3" json:"reportTime,omitempty"`
	TotalDisorderdPacketsNum uint64 `protobuf:"varint,10,opt,name=totalDisorderdPacketsNum,proto3" json:"totalDisorderdPacketsNum,omitempty"`
	TotalLossPacketsNum      uint64 `protobuf:"varint,11,opt,name=totalLossPacketsNum,proto3" json:"totalLossPacketsNum,omitempty"`
	TotalRecvPacketsNum      uint64 `protobuf:"varint,12,opt,name=totalRecvPacketsNum,proto3" json:"totalRecvPacketsNum,omitempty"`
	VlanId                   uint32 `protobuf:"varint,13,opt,name=vlanId,proto3" json:"vlanId,omitempty"`
	VpnName                  string `protobuf:"bytes,14,opt,name=vpnName,proto3" json:"vpnName,omitempty"`
	VrId                     uint32 `protobuf:"varint,15,opt,name=vrId,proto3" json:"vrId,omitempty"`
	VsiName                  string `protobuf:"bytes,16,opt,name=vsiName,proto3" json:"vsiName,omitempty"`
	WorstDisorderdPacketsNum uint64 `protobuf:"varint,17,opt,name=worstDisorderdPacketsNum,proto3" json:"worstDisorderdPacketsNum,omitempty"`
	WorstLossPacketsNum      uint64 `protobuf:"varint,18,opt,name=worstLossPacketsNum,proto3" json:"worstLossPacketsNum,omitempty"`
	WorstRecvPacketsNum      uint64 `protobuf:"varint,19,opt,name=worstRecvPacketsNum,proto3" json:"worstRecvPacketsNum,omitempty"`
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) Reset() {
	*x = TelemEmdi_EmdiTelemReps_EmdiTelemRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemEmdi_EmdiTelemReps_EmdiTelemRep) ProtoMessage() {}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemEmdi_EmdiTelemReps_EmdiTelemRep.ProtoReflect.Descriptor instead.
func (*TelemEmdi_EmdiTelemReps_EmdiTelemRep) Descriptor() ([]byte, []int) {
	return file_huaweiV8R10_telemEmdi_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetBdId() uint32 {
	if x != nil {
		return x.BdId
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetChannelGroupAddress() uint32 {
	if x != nil {
		return x.ChannelGroupAddress
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetChannelSourceAddress() uint32 {
	if x != nil {
		return x.ChannelSourceAddress
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetChannelVpnType() uint32 {
	if x != nil {
		return x.ChannelVpnType
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetMonitorPeriod() uint32 {
	if x != nil {
		return x.MonitorPeriod
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetMonitorTime() uint64 {
	if x != nil {
		return x.MonitorTime
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetReportPeriod() uint32 {
	if x != nil {
		return x.ReportPeriod
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetReportTime() uint64 {
	if x != nil {
		return x.ReportTime
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetTotalDisorderdPacketsNum() uint64 {
	if x != nil {
		return x.TotalDisorderdPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetTotalLossPacketsNum() uint64 {
	if x != nil {
		return x.TotalLossPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetTotalRecvPacketsNum() uint64 {
	if x != nil {
		return x.TotalRecvPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetVlanId() uint32 {
	if x != nil {
		return x.VlanId
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetVpnName() string {
	if x != nil {
		return x.VpnName
	}
	return ""
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetVrId() uint32 {
	if x != nil {
		return x.VrId
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetVsiName() string {
	if x != nil {
		return x.VsiName
	}
	return ""
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetWorstDisorderdPacketsNum() uint64 {
	if x != nil {
		return x.WorstDisorderdPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetWorstLossPacketsNum() uint64 {
	if x != nil {
		return x.WorstLossPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemReps_EmdiTelemRep) GetWorstRecvPacketsNum() uint64 {
	if x != nil {
		return x.WorstRecvPacketsNum
	}
	return 0
}

type TelemEmdi_EmdiTelemRtps_EmdiTelemRtp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BdId                 uint32 `protobuf:"varint,1,opt,name=bdId,proto3" json:"bdId,omitempty"`
	ChannelGroupAddress  uint32 `protobuf:"varint,2,opt,name=channelGroupAddress,proto3" json:"channelGroupAddress,omitempty"`
	ChannelName          string `protobuf:"bytes,3,opt,name=channelName,proto3" json:"channelName,omitempty"`
	ChannelSourceAddress uint32 `protobuf:"varint,4,opt,name=channelSourceAddress,proto3" json:"channelSourceAddress,omitempty"`
	ChannelVpnType       uint32 `protobuf:"varint,5,opt,name=channelVpnType,proto3" json:"channelVpnType,omitempty"`
	DisorderdPacketsNum  uint64 `protobuf:"varint,6,opt,name=disorderdPacketsNum,proto3" json:"disorderdPacketsNum,omitempty"`
	LossPacketsNum       uint64 `protobuf:"varint,7,opt,name=lossPacketsNum,proto3" json:"lossPacketsNum,omitempty"`
	MonitorPeriod        uint32 `protobuf:"varint,8,opt,name=monitorPeriod,proto3" json:"monitorPeriod,omitempty"`
	MonitorTime          uint64 `protobuf:"varint,9,opt,name=monitorTime,proto3" json:"monitorTime,omitempty"`
	RecvPacketsNum       uint64 `protobuf:"varint,10,opt,name=recvPacketsNum,proto3" json:"recvPacketsNum,omitempty"`
	ReportPeriod         uint32 `protobuf:"varint,11,opt,name=reportPeriod,proto3" json:"reportPeriod,omitempty"`
	ReportTime           uint64 `protobuf:"varint,12,opt,name=reportTime,proto3" json:"reportTime,omitempty"`
	VlanId               uint32 `protobuf:"varint,13,opt,name=vlanId,proto3" json:"vlanId,omitempty"`
	VpnName              string `protobuf:"bytes,14,opt,name=vpnName,proto3" json:"vpnName,omitempty"`
	VrId                 uint32 `protobuf:"varint,15,opt,name=vrId,proto3" json:"vrId,omitempty"`
	VsiName              string `protobuf:"bytes,16,opt,name=vsiName,proto3" json:"vsiName,omitempty"`
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) Reset() {
	*x = TelemEmdi_EmdiTelemRtps_EmdiTelemRtp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) ProtoMessage() {}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) ProtoReflect() protoreflect.Message {
	mi := &file_huaweiV8R10_telemEmdi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelemEmdi_EmdiTelemRtps_EmdiTelemRtp.ProtoReflect.Descriptor instead.
func (*TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) Descriptor() ([]byte, []int) {
	return file_huaweiV8R10_telemEmdi_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetBdId() uint32 {
	if x != nil {
		return x.BdId
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetChannelGroupAddress() uint32 {
	if x != nil {
		return x.ChannelGroupAddress
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetChannelSourceAddress() uint32 {
	if x != nil {
		return x.ChannelSourceAddress
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetChannelVpnType() uint32 {
	if x != nil {
		return x.ChannelVpnType
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetDisorderdPacketsNum() uint64 {
	if x != nil {
		return x.DisorderdPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetLossPacketsNum() uint64 {
	if x != nil {
		return x.LossPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetMonitorPeriod() uint32 {
	if x != nil {
		return x.MonitorPeriod
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetMonitorTime() uint64 {
	if x != nil {
		return x.MonitorTime
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetRecvPacketsNum() uint64 {
	if x != nil {
		return x.RecvPacketsNum
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetReportPeriod() uint32 {
	if x != nil {
		return x.ReportPeriod
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetReportTime() uint64 {
	if x != nil {
		return x.ReportTime
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetVlanId() uint32 {
	if x != nil {
		return x.VlanId
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetVpnName() string {
	if x != nil {
		return x.VpnName
	}
	return ""
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetVrId() uint32 {
	if x != nil {
		return x.VrId
	}
	return 0
}

func (x *TelemEmdi_EmdiTelemRtps_EmdiTelemRtp) GetVsiName() string {
	if x != nil {
		return x.VsiName
	}
	return ""
}

var File_huaweiV8R10_telemEmdi_proto protoreflect.FileDescriptor

var file_huaweiV8R10_telemEmdi_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x30, 0x2d, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68,
	0x75, 0x61, 0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x30, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d,
	0x45, 0x6d, 0x64, 0x69, 0x22, 0xe1, 0x0d, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d,
	0x64, 0x69, 0x12, 0x54, 0x0a, 0x0d, 0x65, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52,
	0x65, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x68, 0x75, 0x61, 0x77,
	0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x30, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64,
	0x69, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64, 0x69, 0x2e, 0x45, 0x6d, 0x64, 0x69,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x73, 0x52, 0x0d, 0x65, 0x6d, 0x64, 0x69, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x73, 0x12, 0x54, 0x0a, 0x0d, 0x65, 0x6d, 0x64, 0x69,
	0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x74, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x30, 0x5f, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64, 0x69, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64,
	0x69, 0x2e, 0x45, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x74, 0x70, 0x73, 0x52,
	0x0d, 0x65, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x74, 0x70, 0x73, 0x1a, 0xf1,
	0x06, 0x0a, 0x0d, 0x45, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x70, 0x73,
	0x12, 0x5f, 0x0a, 0x0c, 0x65, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x56,
	0x38, 0x52, 0x31, 0x30, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64, 0x69, 0x2e, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64, 0x69, 0x2e, 0x45, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x52, 0x65, 0x70, 0x73, 0x2e, 0x45, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x52, 0x65, 0x70, 0x52, 0x0c, 0x65, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x65,
	0x70, 0x1a, 0xfe, 0x05, 0x0a, 0x0c, 0x45, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52,
	0x65, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x62, 0x64, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x56, 0x70, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x56,
	0x70, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f,
	0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x18, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x69, 0x73, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x30,
	0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x6f, 0x73, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4c, 0x6f, 0x73, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d,
	0x12, 0x30, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x76, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x76, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e,
	0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x70,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x70, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x72, 0x49, 0x64, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x76, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x73, 0x69, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x73, 0x69, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x77, 0x6f, 0x72, 0x73, 0x74, 0x44, 0x69, 0x73, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x18, 0x77, 0x6f, 0x72, 0x73, 0x74, 0x44, 0x69, 0x73, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x30,
	0x0a, 0x13, 0x77, 0x6f, 0x72, 0x73, 0x74, 0x4c, 0x6f, 0x73, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x77, 0x6f, 0x72,
	0x73, 0x74, 0x4c, 0x6f, 0x73, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d,
	0x12, 0x30, 0x0a, 0x13, 0x77, 0x6f, 0x72, 0x73, 0x74, 0x52, 0x65, 0x63, 0x76, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x77,
	0x6f, 0x72, 0x73, 0x74, 0x52, 0x65, 0x63, 0x76, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e,
	0x75, 0x6d, 0x1a, 0xb3, 0x05, 0x0a, 0x0d, 0x45, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d,
	0x52, 0x74, 0x70, 0x73, 0x12, 0x5f, 0x0a, 0x0c, 0x65, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c, 0x65,
	0x6d, 0x52, 0x74, 0x70, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x68, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x56, 0x38, 0x52, 0x31, 0x30, 0x5f, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d,
	0x64, 0x69, 0x2e, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x45, 0x6d, 0x64, 0x69, 0x2e, 0x45, 0x6d, 0x64,
	0x69, 0x54, 0x65, 0x6c, 0x65, 0x6d, 0x52, 0x74, 0x70, 0x73, 0x2e, 0x45, 0x6d, 0x64, 0x69, 0x54,
	0x65, 0x6c, 0x65, 0x6d, 0x52, 0x74, 0x70, 0x52, 0x0c, 0x65, 0x6d, 0x64, 0x69, 0x54, 0x65, 0x6c,
	0x65, 0x6d, 0x52, 0x74, 0x70, 0x1a, 0xc0, 0x04, 0x0a, 0x0c, 0x45, 0x6d, 0x64, 0x69, 0x54, 0x65,
	0x6c, 0x65, 0x6d, 0x52, 0x74, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x64, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x62, 0x64, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x14, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x56, 0x70, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x56, 0x70, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x69,
	0x73, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x64, 0x69, 0x73, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x64, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x12, 0x26, 0x0a, 0x0e,
	0x6c, 0x6f, 0x73, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x6c, 0x6f, 0x73, 0x73, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x4e, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x72, 0x65, 0x63, 0x76, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x4e, 0x75, 0x6d, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x76, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x4e, 0x75, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6c, 0x61, 0x6e,
	0x49, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x76, 0x6c, 0x61, 0x6e, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x70, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x70, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x72,
	0x49, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x76, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x73, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x73, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_huaweiV8R10_telemEmdi_proto_rawDescOnce sync.Once
	file_huaweiV8R10_telemEmdi_proto_rawDescData = file_huaweiV8R10_telemEmdi_proto_rawDesc
)

func file_huaweiV8R10_telemEmdi_proto_rawDescGZIP() []byte {
	file_huaweiV8R10_telemEmdi_proto_rawDescOnce.Do(func() {
		file_huaweiV8R10_telemEmdi_proto_rawDescData = protoimpl.X.CompressGZIP(file_huaweiV8R10_telemEmdi_proto_rawDescData)
	})
	return file_huaweiV8R10_telemEmdi_proto_rawDescData
}

var file_huaweiV8R10_telemEmdi_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_huaweiV8R10_telemEmdi_proto_goTypes = []interface{}{
	(*TelemEmdi)(nil),                            // 0: huaweiV8R10_telemEmdi.TelemEmdi
	(*TelemEmdi_EmdiTelemReps)(nil),              // 1: huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemReps
	(*TelemEmdi_EmdiTelemRtps)(nil),              // 2: huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemRtps
	(*TelemEmdi_EmdiTelemReps_EmdiTelemRep)(nil), // 3: huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemReps.EmdiTelemRep
	(*TelemEmdi_EmdiTelemRtps_EmdiTelemRtp)(nil), // 4: huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemRtps.EmdiTelemRtp
}
var file_huaweiV8R10_telemEmdi_proto_depIdxs = []int32{
	1, // 0: huaweiV8R10_telemEmdi.TelemEmdi.emdiTelemReps:type_name -> huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemReps
	2, // 1: huaweiV8R10_telemEmdi.TelemEmdi.emdiTelemRtps:type_name -> huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemRtps
	3, // 2: huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemReps.emdiTelemRep:type_name -> huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemReps.EmdiTelemRep
	4, // 3: huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemRtps.emdiTelemRtp:type_name -> huaweiV8R10_telemEmdi.TelemEmdi.EmdiTelemRtps.EmdiTelemRtp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_huaweiV8R10_telemEmdi_proto_init() }
func file_huaweiV8R10_telemEmdi_proto_init() {
	if File_huaweiV8R10_telemEmdi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_huaweiV8R10_telemEmdi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemEmdi); i {
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
		file_huaweiV8R10_telemEmdi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemEmdi_EmdiTelemReps); i {
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
		file_huaweiV8R10_telemEmdi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemEmdi_EmdiTelemRtps); i {
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
		file_huaweiV8R10_telemEmdi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemEmdi_EmdiTelemReps_EmdiTelemRep); i {
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
		file_huaweiV8R10_telemEmdi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelemEmdi_EmdiTelemRtps_EmdiTelemRtp); i {
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
			RawDescriptor: file_huaweiV8R10_telemEmdi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_huaweiV8R10_telemEmdi_proto_goTypes,
		DependencyIndexes: file_huaweiV8R10_telemEmdi_proto_depIdxs,
		MessageInfos:      file_huaweiV8R10_telemEmdi_proto_msgTypes,
	}.Build()
	File_huaweiV8R10_telemEmdi_proto = out.File
	file_huaweiV8R10_telemEmdi_proto_rawDesc = nil
	file_huaweiV8R10_telemEmdi_proto_goTypes = nil
	file_huaweiV8R10_telemEmdi_proto_depIdxs = nil
}
