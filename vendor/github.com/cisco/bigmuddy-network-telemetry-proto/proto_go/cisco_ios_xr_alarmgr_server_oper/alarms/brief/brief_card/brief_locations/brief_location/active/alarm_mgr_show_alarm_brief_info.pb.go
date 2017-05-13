// Code generated by protoc-gen-go.
// source: alarm_mgr_show_alarm_brief_info.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_alarmgr_server_oper_alarms_brief_brief_card_brief_locations_brief_location_active is a generated protocol buffer package.

It is generated from these files:
	alarm_mgr_show_alarm_brief_info.proto

It has these top-level messages:
	AlarmMgrShowAlarmBriefInfo_KEYS
	AlarmMgrShowAlarmBriefInfo
	AlarmMgrShowAlarmBriefData
*/
package cisco_ios_xr_alarmgr_server_oper_alarms_brief_brief_card_brief_locations_brief_location_active

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Alarm mgr show alarm brief info
type AlarmMgrShowAlarmBriefInfo_KEYS struct {
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
}

func (m *AlarmMgrShowAlarmBriefInfo_KEYS) Reset()                    { *m = AlarmMgrShowAlarmBriefInfo_KEYS{} }
func (m *AlarmMgrShowAlarmBriefInfo_KEYS) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmBriefInfo_KEYS) ProtoMessage()               {}
func (*AlarmMgrShowAlarmBriefInfo_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AlarmMgrShowAlarmBriefInfo_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type AlarmMgrShowAlarmBriefInfo struct {
	// Alarm List
	AlarmInfo []*AlarmMgrShowAlarmBriefData `protobuf:"bytes,50,rep,name=alarm_info,json=alarmInfo" json:"alarm_info,omitempty"`
}

func (m *AlarmMgrShowAlarmBriefInfo) Reset()                    { *m = AlarmMgrShowAlarmBriefInfo{} }
func (m *AlarmMgrShowAlarmBriefInfo) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmBriefInfo) ProtoMessage()               {}
func (*AlarmMgrShowAlarmBriefInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AlarmMgrShowAlarmBriefInfo) GetAlarmInfo() []*AlarmMgrShowAlarmBriefData {
	if m != nil {
		return m.AlarmInfo
	}
	return nil
}

// Alarm mgr show alarm brief data
type AlarmMgrShowAlarmBriefData struct {
	// Alarm location
	Location string `protobuf:"bytes,1,opt,name=location" json:"location,omitempty"`
	// Alarm severity
	Severity string `protobuf:"bytes,2,opt,name=severity" json:"severity,omitempty"`
	// Alarm group
	Group string `protobuf:"bytes,3,opt,name=group" json:"group,omitempty"`
	// Alarm set time
	SetTime string `protobuf:"bytes,4,opt,name=set_time,json=setTime" json:"set_time,omitempty"`
	// Alarm set time(timestamp format)
	SetTimestamp uint64 `protobuf:"varint,5,opt,name=set_timestamp,json=setTimestamp" json:"set_timestamp,omitempty"`
	// Alarm clear time
	ClearTime string `protobuf:"bytes,6,opt,name=clear_time,json=clearTime" json:"clear_time,omitempty"`
	// Alarm clear time(timestamp format)
	ClearTimestamp uint64 `protobuf:"varint,7,opt,name=clear_timestamp,json=clearTimestamp" json:"clear_timestamp,omitempty"`
	// Alarm description
	Description string `protobuf:"bytes,8,opt,name=description" json:"description,omitempty"`
}

func (m *AlarmMgrShowAlarmBriefData) Reset()                    { *m = AlarmMgrShowAlarmBriefData{} }
func (m *AlarmMgrShowAlarmBriefData) String() string            { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmBriefData) ProtoMessage()               {}
func (*AlarmMgrShowAlarmBriefData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AlarmMgrShowAlarmBriefData) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetSetTime() string {
	if m != nil {
		return m.SetTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetSetTimestamp() uint64 {
	if m != nil {
		return m.SetTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmBriefData) GetClearTime() string {
	if m != nil {
		return m.ClearTime
	}
	return ""
}

func (m *AlarmMgrShowAlarmBriefData) GetClearTimestamp() uint64 {
	if m != nil {
		return m.ClearTimestamp
	}
	return 0
}

func (m *AlarmMgrShowAlarmBriefData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*AlarmMgrShowAlarmBriefInfo_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_card.brief_locations.brief_location.active.alarm_mgr_show_alarm_brief_info_KEYS")
	proto.RegisterType((*AlarmMgrShowAlarmBriefInfo)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_card.brief_locations.brief_location.active.alarm_mgr_show_alarm_brief_info")
	proto.RegisterType((*AlarmMgrShowAlarmBriefData)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.brief.brief_card.brief_locations.brief_location.active.alarm_mgr_show_alarm_brief_data")
}

func init() { proto.RegisterFile("alarm_mgr_show_alarm_brief_info.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0xfe, 0xef, 0x95, 0x3f, 0x92, 0x85, 0x44, 0x40, 0x42, 0x44, 0x05, 0x44, 0xa7,
	0x0c, 0xe5, 0x01, 0x98, 0x18, 0x2a, 0xb6, 0xc2, 0xc2, 0x82, 0xe5, 0x3a, 0xd7, 0x62, 0xa9, 0x89,
	0xa3, 0xb3, 0x29, 0x30, 0xf1, 0x1c, 0x2c, 0x3c, 0x0c, 0x4f, 0x86, 0x72, 0x6e, 0x0b, 0x62, 0x68,
	0x27, 0x96, 0x28, 0xdf, 0xf7, 0xdd, 0xfd, 0x7c, 0x67, 0x19, 0x2e, 0xd4, 0x5c, 0x51, 0x2e, 0xf3,
	0x19, 0x49, 0xf7, 0x64, 0x5f, 0x64, 0x90, 0x13, 0x32, 0x38, 0x95, 0xa6, 0x98, 0xda, 0xb4, 0x24,
	0xeb, 0xad, 0x78, 0xd4, 0xc6, 0x69, 0x2b, 0x8d, 0x75, 0xf2, 0x95, 0x42, 0x51, 0xd5, 0x81, 0xb4,
	0x40, 0x92, 0xb6, 0x44, 0x4a, 0xd9, 0x73, 0x29, 0x77, 0x86, 0xaf, 0xd4, 0x8a, 0xb2, 0xe5, 0xef,
	0xdc, 0x6a, 0xe5, 0x8d, 0x2d, 0xdc, 0x1f, 0x9d, 0x2a, 0xed, 0xcd, 0x02, 0xfb, 0xd7, 0x70, 0xbe,
	0x65, 0x10, 0x79, 0x7b, 0xf3, 0x70, 0x27, 0x0e, 0xa1, 0x5d, 0xd8, 0x0c, 0xa5, 0xc9, 0xe2, 0x28,
	0x89, 0x06, 0xdd, 0x71, 0xab, 0x92, 0xa3, 0xac, 0xff, 0x15, 0xc1, 0xe9, 0x16, 0x82, 0xf8, 0x8c,
	0x00, 0x82, 0x59, 0xc9, 0x78, 0x98, 0xd4, 0x07, 0xbd, 0xe1, 0x7b, 0xfa, 0xbf, 0xab, 0xa5, 0x1b,
	0xa6, 0xca, 0x94, 0x57, 0xe3, 0x2e, 0x3b, 0xa3, 0x62, 0x6a, 0xfb, 0x1f, 0xb5, 0x8d, 0x4b, 0x54,
	0xe5, 0xe2, 0x18, 0x3a, 0xab, 0x13, 0x96, 0x57, 0xb0, 0xd6, 0x55, 0xe6, 0x70, 0x81, 0x64, 0xfc,
	0x5b, 0x5c, 0x0b, 0xd9, 0x4a, 0x8b, 0x03, 0x68, 0xce, 0xc8, 0x3e, 0x97, 0x71, 0x9d, 0x83, 0x20,
	0xc4, 0x51, 0xd5, 0xe1, 0xa5, 0x37, 0x39, 0xc6, 0x0d, 0x0e, 0xda, 0x0e, 0xfd, 0xbd, 0xc9, 0x51,
	0x9c, 0xc1, 0xee, 0x2a, 0x72, 0x5e, 0xe5, 0x65, 0xdc, 0x4c, 0xa2, 0x41, 0x63, 0xbc, 0xb3, 0xcc,
	0xd9, 0x13, 0x27, 0x00, 0x7a, 0x8e, 0x8a, 0x02, 0xa1, 0xc5, 0x84, 0x2e, 0x3b, 0xcc, 0xb8, 0x84,
	0xfd, 0x9f, 0x38, 0x50, 0xda, 0x4c, 0xd9, 0x5b, 0xd7, 0x04, 0x4e, 0x02, 0xbd, 0x0c, 0x9d, 0x26,
	0x53, 0xf2, 0x62, 0x1d, 0x06, 0xfd, 0xb6, 0x26, 0x2d, 0x7e, 0x88, 0x57, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xeb, 0x00, 0x79, 0xf0, 0xb1, 0x02, 0x00, 0x00,
}
