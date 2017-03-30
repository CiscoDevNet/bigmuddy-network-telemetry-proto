// Code generated by protoc-gen-go.
// source: ospfv3_edm_flood_list.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_flood_list_process_table_flood is a generated protocol buffer package.

It is generated from these files:
	ospfv3_edm_flood_list.proto

It has these top-level messages:
	Ospfv3EdmFloodList_KEYS
	Ospfv3EdmFloodList
	Ospfv3EdmLsaSum
*/
package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_flood_list_process_table_flood

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

// OSPFv3 flood list information
type Ospfv3EdmFloodList_KEYS struct {
	ProcessName   string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *Ospfv3EdmFloodList_KEYS) Reset()                    { *m = Ospfv3EdmFloodList_KEYS{} }
func (m *Ospfv3EdmFloodList_KEYS) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmFloodList_KEYS) ProtoMessage()               {}
func (*Ospfv3EdmFloodList_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ospfv3EdmFloodList_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmFloodList_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ospfv3EdmFloodList struct {
	// Time until next LS transmission (ms)
	LsTransmissionTimer uint32 `protobuf:"varint,50,opt,name=ls_transmission_timer,json=lsTransmissionTimer" json:"ls_transmission_timer,omitempty"`
	// Number of LSAs currently being flooded
	QueueLength uint32 `protobuf:"varint,51,opt,name=queue_length,json=queueLength" json:"queue_length,omitempty"`
	// Link floodlist
	LinkFloodList []*Ospfv3EdmLsaSum `protobuf:"bytes,52,rep,name=link_flood_list,json=linkFloodList" json:"link_flood_list,omitempty"`
	// Area scope floodlist
	AreaFloodList []*Ospfv3EdmLsaSum `protobuf:"bytes,53,rep,name=area_flood_list,json=areaFloodList" json:"area_flood_list,omitempty"`
	// AS scope floodlist
	AsFloodList []*Ospfv3EdmLsaSum `protobuf:"bytes,54,rep,name=as_flood_list,json=asFloodList" json:"as_flood_list,omitempty"`
}

func (m *Ospfv3EdmFloodList) Reset()                    { *m = Ospfv3EdmFloodList{} }
func (m *Ospfv3EdmFloodList) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmFloodList) ProtoMessage()               {}
func (*Ospfv3EdmFloodList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Ospfv3EdmFloodList) GetLsTransmissionTimer() uint32 {
	if m != nil {
		return m.LsTransmissionTimer
	}
	return 0
}

func (m *Ospfv3EdmFloodList) GetQueueLength() uint32 {
	if m != nil {
		return m.QueueLength
	}
	return 0
}

func (m *Ospfv3EdmFloodList) GetLinkFloodList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.LinkFloodList
	}
	return nil
}

func (m *Ospfv3EdmFloodList) GetAreaFloodList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.AreaFloodList
	}
	return nil
}

func (m *Ospfv3EdmFloodList) GetAsFloodList() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.AsFloodList
	}
	return nil
}

// LSA summary entry
type Ospfv3EdmLsaSum struct {
	// LSA type
	HeaderLsaType string `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType" json:"header_lsa_type,omitempty"`
	// Age of the LSA (seconds)
	HeaderLsaAge uint32 `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge" json:"header_lsa_age,omitempty"`
	// LSA ID
	HeaderLsaId string `protobuf:"bytes,3,opt,name=header_lsa_id,json=headerLsaId" json:"header_lsa_id,omitempty"`
	// Router ID of the advertising router
	HeaderAdvertisingRouter string `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter" json:"header_advertising_router,omitempty"`
	// Current LSA sequence number
	HeaderSequenceNumber int32 `protobuf:"zigzag32,5,opt,name=header_sequence_number,json=headerSequenceNumber" json:"header_sequence_number,omitempty"`
}

func (m *Ospfv3EdmLsaSum) Reset()                    { *m = Ospfv3EdmLsaSum{} }
func (m *Ospfv3EdmLsaSum) String() string            { return proto.CompactTextString(m) }
func (*Ospfv3EdmLsaSum) ProtoMessage()               {}
func (*Ospfv3EdmLsaSum) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Ospfv3EdmLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaId() string {
	if m != nil {
		return m.HeaderLsaId
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderSequenceNumber() int32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmFloodList_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.flood_list_process_table.flood.ospfv3_edm_flood_list_KEYS")
	proto.RegisterType((*Ospfv3EdmFloodList)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.flood_list_process_table.flood.ospfv3_edm_flood_list")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.flood_list_process_table.flood.ospfv3_edm_lsa_sum")
}

func init() { proto.RegisterFile("ospfv3_edm_flood_list.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x65, 0x5a, 0x90, 0xb8, 0xd4, 0xad, 0x38, 0x28, 0x18, 0x58, 0x42, 0x04, 0x28, 0x93,
	0x87, 0xa4, 0x74, 0x60, 0xeb, 0x00, 0x12, 0x22, 0xea, 0xe0, 0x66, 0x61, 0x7a, 0x7a, 0x89, 0x9f,
	0xd3, 0x13, 0x67, 0x9f, 0x7b, 0xef, 0x1c, 0xd1, 0x8f, 0xc2, 0xc0, 0x17, 0xe2, 0x13, 0x31, 0xa2,
	0x3b, 0x3b, 0x8e, 0x25, 0xba, 0xa2, 0x6e, 0xa7, 0xff, 0xff, 0x77, 0xbe, 0xdf, 0x59, 0xcf, 0x16,
	0xaf, 0x0d, 0xd7, 0xc5, 0x76, 0x0e, 0x94, 0x97, 0x50, 0x68, 0x63, 0x72, 0xd0, 0x8a, 0x5d, 0x5a,
	0x5b, 0xe3, 0x8c, 0xc4, 0xb5, 0xe2, 0xb5, 0x01, 0x65, 0x18, 0x7e, 0x58, 0x50, 0xf5, 0xf6, 0x1c,
	0x3a, 0xdc, 0xd4, 0x64, 0xd3, 0x76, 0xed, 0xd9, 0x35, 0x31, 0x13, 0xef, 0x56, 0x69, 0x4e, 0x05,
	0x36, 0xda, 0xc1, 0xd6, 0x16, 0xe9, 0xfe, 0xa1, 0xd0, 0xd5, 0xe0, 0x70, 0xa5, 0xa9, 0x2d, 0x26,
	0x85, 0x78, 0x75, 0xa7, 0x01, 0x7c, 0xfd, 0xf4, 0xed, 0x4a, 0xbe, 0x11, 0x47, 0xbb, 0x4d, 0x15,
	0x96, 0x94, 0x44, 0xe3, 0x68, 0xfa, 0x38, 0x1b, 0x75, 0xd9, 0x25, 0x96, 0x24, 0xdf, 0x89, 0x63,
	0x55, 0x39, 0xb2, 0x05, 0xae, 0xa9, 0x85, 0x1e, 0x04, 0x28, 0xee, 0x53, 0x8f, 0x4d, 0x7e, 0x1f,
	0x8a, 0xd3, 0x3b, 0x0f, 0x92, 0x33, 0x71, 0xaa, 0x19, 0x9c, 0xc5, 0x8a, 0x4b, 0xc5, 0xac, 0x4c,
	0x05, 0x4e, 0x95, 0x64, 0x93, 0xd9, 0x38, 0x9a, 0xc6, 0xd9, 0x53, 0xcd, 0xcb, 0x41, 0xb7, 0xf4,
	0x95, 0xf7, 0xba, 0x69, 0xa8, 0x21, 0xd0, 0x54, 0x6d, 0xdc, 0x75, 0x32, 0x0f, 0xe8, 0x28, 0x64,
	0x8b, 0x10, 0xc9, 0x5f, 0x91, 0x38, 0xd1, 0xaa, 0xfa, 0x3e, 0x38, 0x2a, 0x39, 0x1b, 0x1f, 0x4c,
	0x47, 0xb3, 0x26, 0xfd, 0xef, 0xaf, 0x35, 0x1d, 0x5c, 0x55, 0x33, 0x02, 0x37, 0x65, 0x16, 0x7b,
	0x9b, 0xcf, 0xbe, 0x5d, 0xf8, 0x6b, 0x7b, 0x3f, 0xb4, 0x84, 0x43, 0xbf, 0x0f, 0xf7, 0xea, 0xe7,
	0x6d, 0xf6, 0x7e, 0x3f, 0x23, 0x11, 0x23, 0x0f, 0xed, 0xce, 0xef, 0xd3, 0x6e, 0x84, 0xdc, 0xbb,
	0x4d, 0xfe, 0x44, 0x42, 0xfe, 0xcb, 0xc8, 0xf7, 0xe2, 0xe4, 0x9a, 0x30, 0x27, 0x1b, 0x12, 0x77,
	0x5b, 0xef, 0x06, 0x36, 0x6e, 0xe3, 0x05, 0xe3, 0xf2, 0xb6, 0x26, 0xf9, 0x56, 0x1c, 0x0f, 0x38,
	0xdc, 0xb4, 0x23, 0x1b, 0x67, 0x47, 0x3d, 0x76, 0xb1, 0x21, 0x39, 0x11, 0xf1, 0x80, 0x52, 0x79,
	0x72, 0xd0, 0x0e, 0x7f, 0x0f, 0x7d, 0xc9, 0xe5, 0x47, 0xf1, 0xb2, 0x63, 0x30, 0xdf, 0x92, 0x75,
	0x8a, 0x55, 0xb5, 0x01, 0x6b, 0x1a, 0x47, 0x36, 0x39, 0x0c, 0xfc, 0x8b, 0x16, 0xb8, 0xd8, 0xf7,
	0x59, 0xa8, 0xe5, 0x99, 0x78, 0xde, 0xed, 0x65, 0xba, 0x69, 0xa8, 0xf2, 0x9f, 0x4f, 0x53, 0xae,
	0xc8, 0x26, 0x0f, 0xc7, 0xd1, 0xf4, 0x49, 0xf6, 0xac, 0x6d, 0xaf, 0xba, 0xf2, 0x32, 0x74, 0xab,
	0x47, 0xe1, 0xcf, 0x30, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xee, 0x75, 0xff, 0x56, 0x38, 0x04,
	0x00, 0x00,
}
