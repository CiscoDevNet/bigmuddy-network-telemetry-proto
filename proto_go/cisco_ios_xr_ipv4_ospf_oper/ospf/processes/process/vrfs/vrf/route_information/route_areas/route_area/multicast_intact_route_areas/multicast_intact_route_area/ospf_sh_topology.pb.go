// Code generated by protoc-gen-go.
// source: ospf_sh_topology.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_route_areas_route_area_multicast_intact_route_areas_multicast_intact_route_area is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_topology.proto

It has these top-level messages:
	OspfShTopology_KEYS
	OspfShTopology
	OspfShTime
	OspfShTopCommon
	OspfShTopPath
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_route_information_route_areas_route_area_multicast_intact_route_areas_multicast_intact_route_area

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

// OSPF Route Information
type OspfShTopology_KEYS struct {
	ProcessName  string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	VrfName      string `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName" json:"vrf_name,omitempty"`
	AreaId       uint32 `protobuf:"varint,3,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	Prefix       string `protobuf:"bytes,4,opt,name=prefix" json:"prefix,omitempty"`
	PrefixLength uint32 `protobuf:"varint,5,opt,name=prefix_length,json=prefixLength" json:"prefix_length,omitempty"`
}

func (m *OspfShTopology_KEYS) Reset()                    { *m = OspfShTopology_KEYS{} }
func (m *OspfShTopology_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology_KEYS) ProtoMessage()               {}
func (*OspfShTopology_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShTopology_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopology_KEYS) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *OspfShTopology_KEYS) GetPrefixLength() uint32 {
	if m != nil {
		return m.PrefixLength
	}
	return 0
}

type OspfShTopology struct {
	// Prefix
	RoutePrefix string `protobuf:"bytes,50,opt,name=route_prefix,json=routePrefix" json:"route_prefix,omitempty"`
	// Prefix length
	RoutePrefixLength uint32 `protobuf:"varint,51,opt,name=route_prefix_length,json=routePrefixLength" json:"route_prefix_length,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,52,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// Route type
	RouteType string `protobuf:"bytes,53,opt,name=route_type,json=routeType" json:"route_type,omitempty"`
	// If true, connected route
	RouteConnected bool `protobuf:"varint,54,opt,name=route_connected,json=routeConnected" json:"route_connected,omitempty"`
	// Route information
	RouteInfo *OspfShTopCommon `protobuf:"bytes,55,opt,name=route_info,json=routeInfo" json:"route_info,omitempty"`
	// List of paths to this route
	RoutePathList []*OspfShTopPath `protobuf:"bytes,56,rep,name=route_path_list,json=routePathList" json:"route_path_list,omitempty"`
}

func (m *OspfShTopology) Reset()                    { *m = OspfShTopology{} }
func (m *OspfShTopology) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopology) ProtoMessage()               {}
func (*OspfShTopology) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShTopology) GetRoutePrefix() string {
	if m != nil {
		return m.RoutePrefix
	}
	return ""
}

func (m *OspfShTopology) GetRoutePrefixLength() uint32 {
	if m != nil {
		return m.RoutePrefixLength
	}
	return 0
}

func (m *OspfShTopology) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopology) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *OspfShTopology) GetRouteConnected() bool {
	if m != nil {
		return m.RouteConnected
	}
	return false
}

func (m *OspfShTopology) GetRouteInfo() *OspfShTopCommon {
	if m != nil {
		return m.RouteInfo
	}
	return nil
}

func (m *OspfShTopology) GetRoutePathList() []*OspfShTopPath {
	if m != nil {
		return m.RoutePathList
	}
	return nil
}

type OspfShTime struct {
	Second     uint32 `protobuf:"varint,1,opt,name=second" json:"second,omitempty"`
	Nanosecond uint32 `protobuf:"varint,2,opt,name=nanosecond" json:"nanosecond,omitempty"`
}

func (m *OspfShTime) Reset()                    { *m = OspfShTime{} }
func (m *OspfShTime) String() string            { return proto.CompactTextString(m) }
func (*OspfShTime) ProtoMessage()               {}
func (*OspfShTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OspfShTime) GetSecond() uint32 {
	if m != nil {
		return m.Second
	}
	return 0
}

func (m *OspfShTime) GetNanosecond() uint32 {
	if m != nil {
		return m.Nanosecond
	}
	return 0
}

// OSPF Common Route Information
type OspfShTopCommon struct {
	// Area ID
	RouteAreaId uint32 `protobuf:"varint,1,opt,name=route_area_id,json=routeAreaId" json:"route_area_id,omitempty"`
	// TE metric
	RouteTeMetric uint32 `protobuf:"varint,2,opt,name=route_te_metric,json=routeTeMetric" json:"route_te_metric,omitempty"`
	// RIB version
	RouteRibVersion uint32 `protobuf:"varint,3,opt,name=route_rib_version,json=routeRibVersion" json:"route_rib_version,omitempty"`
	// SPF version
	RouteSpfVersion uint64 `protobuf:"varint,4,opt,name=route_spf_version,json=routeSpfVersion" json:"route_spf_version,omitempty"`
	// Forward distance
	RouteForwardDistance uint32 `protobuf:"varint,5,opt,name=route_forward_distance,json=routeForwardDistance" json:"route_forward_distance,omitempty"`
	// Protocol source
	RouteSource uint32 `protobuf:"varint,6,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// Last time updated
	RouteUpdateTime *OspfShTime `protobuf:"bytes,7,opt,name=route_update_time,json=routeUpdateTime" json:"route_update_time,omitempty"`
	// Last time update failed
	RouteFailTime *OspfShTime `protobuf:"bytes,8,opt,name=route_fail_time,json=routeFailTime" json:"route_fail_time,omitempty"`
	// SPF priority
	RouteSpfPriority uint32 `protobuf:"varint,9,opt,name=route_spf_priority,json=routeSpfPriority" json:"route_spf_priority,omitempty"`
	// If true, exclude from TE paths
	RouteAutoExcluded bool `protobuf:"varint,10,opt,name=route_auto_excluded,json=routeAutoExcluded" json:"route_auto_excluded,omitempty"`
	// If true, SRTE registered prefix route
	RouteSrtePrefixRegistered bool `protobuf:"varint,11,opt,name=route_srte_prefix_registered,json=routeSrtePrefixRegistered" json:"route_srte_prefix_registered,omitempty"`
	// SRTE registered neigbhor count on route
	RouteSrteNbrRegistered uint32 `protobuf:"varint,12,opt,name=route_srte_nbr_registered,json=routeSrteNbrRegistered" json:"route_srte_nbr_registered,omitempty"`
}

func (m *OspfShTopCommon) Reset()                    { *m = OspfShTopCommon{} }
func (m *OspfShTopCommon) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopCommon) ProtoMessage()               {}
func (*OspfShTopCommon) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OspfShTopCommon) GetRouteAreaId() uint32 {
	if m != nil {
		return m.RouteAreaId
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteTeMetric() uint32 {
	if m != nil {
		return m.RouteTeMetric
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteRibVersion() uint32 {
	if m != nil {
		return m.RouteRibVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSpfVersion() uint64 {
	if m != nil {
		return m.RouteSpfVersion
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteForwardDistance() uint32 {
	if m != nil {
		return m.RouteForwardDistance
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteSource() uint32 {
	if m != nil {
		return m.RouteSource
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteUpdateTime() *OspfShTime {
	if m != nil {
		return m.RouteUpdateTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteFailTime() *OspfShTime {
	if m != nil {
		return m.RouteFailTime
	}
	return nil
}

func (m *OspfShTopCommon) GetRouteSpfPriority() uint32 {
	if m != nil {
		return m.RouteSpfPriority
	}
	return 0
}

func (m *OspfShTopCommon) GetRouteAutoExcluded() bool {
	if m != nil {
		return m.RouteAutoExcluded
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrtePrefixRegistered() bool {
	if m != nil {
		return m.RouteSrtePrefixRegistered
	}
	return false
}

func (m *OspfShTopCommon) GetRouteSrteNbrRegistered() uint32 {
	if m != nil {
		return m.RouteSrteNbrRegistered
	}
	return 0
}

// OSPF Route Path Information
type OspfShTopPath struct {
	// Next hop Interface
	RouteInterfaceName string `protobuf:"bytes,1,opt,name=route_interface_name,json=routeInterfaceName" json:"route_interface_name,omitempty"`
	// Nexthop IP address
	RouteNextHopAddress string `protobuf:"bytes,2,opt,name=route_next_hop_address,json=routeNextHopAddress" json:"route_next_hop_address,omitempty"`
	// IP address of source of route
	RouteSource string `protobuf:"bytes,3,opt,name=route_source,json=routeSource" json:"route_source,omitempty"`
	// LSA ID, see RFC2328
	RouteLsaid string `protobuf:"bytes,4,opt,name=route_lsaid,json=routeLsaid" json:"route_lsaid,omitempty"`
	// Multicast-intact path
	RoutePathIsMcastIntact bool `protobuf:"varint,5,opt,name=route_path_is_mcast_intact,json=routePathIsMcastIntact" json:"route_path_is_mcast_intact,omitempty"`
	// UCMP path
	RoutePathIsUcmpPath bool `protobuf:"varint,6,opt,name=route_path_is_ucmp_path,json=routePathIsUcmpPath" json:"route_path_is_ucmp_path,omitempty"`
	// Metric
	RouteMetric uint32 `protobuf:"varint,7,opt,name=route_metric,json=routeMetric" json:"route_metric,omitempty"`
	// LSA type, see RFC2328 etc.
	LsaType uint32 `protobuf:"varint,8,opt,name=lsa_type,json=lsaType" json:"lsa_type,omitempty"`
	// Area ID
	AreaId uint32 `protobuf:"varint,9,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	// Area format IP or uint32
	AreaFormat bool `protobuf:"varint,10,opt,name=area_format,json=areaFormat" json:"area_format,omitempty"`
}

func (m *OspfShTopPath) Reset()                    { *m = OspfShTopPath{} }
func (m *OspfShTopPath) String() string            { return proto.CompactTextString(m) }
func (*OspfShTopPath) ProtoMessage()               {}
func (*OspfShTopPath) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OspfShTopPath) GetRouteInterfaceName() string {
	if m != nil {
		return m.RouteInterfaceName
	}
	return ""
}

func (m *OspfShTopPath) GetRouteNextHopAddress() string {
	if m != nil {
		return m.RouteNextHopAddress
	}
	return ""
}

func (m *OspfShTopPath) GetRouteSource() string {
	if m != nil {
		return m.RouteSource
	}
	return ""
}

func (m *OspfShTopPath) GetRouteLsaid() string {
	if m != nil {
		return m.RouteLsaid
	}
	return ""
}

func (m *OspfShTopPath) GetRoutePathIsMcastIntact() bool {
	if m != nil {
		return m.RoutePathIsMcastIntact
	}
	return false
}

func (m *OspfShTopPath) GetRoutePathIsUcmpPath() bool {
	if m != nil {
		return m.RoutePathIsUcmpPath
	}
	return false
}

func (m *OspfShTopPath) GetRouteMetric() uint32 {
	if m != nil {
		return m.RouteMetric
	}
	return 0
}

func (m *OspfShTopPath) GetLsaType() uint32 {
	if m != nil {
		return m.LsaType
	}
	return 0
}

func (m *OspfShTopPath) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShTopPath) GetAreaFormat() bool {
	if m != nil {
		return m.AreaFormat
	}
	return false
}

func init() {
	proto.RegisterType((*OspfShTopology_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_topology_KEYS")
	proto.RegisterType((*OspfShTopology)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_topology")
	proto.RegisterType((*OspfShTime)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_time")
	proto.RegisterType((*OspfShTopCommon)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_top_common")
	proto.RegisterType((*OspfShTopPath)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.route_information.route_areas.route_area.multicast_intact_route_areas.multicast_intact_route_area.ospf_sh_top_path")
}

func init() { proto.RegisterFile("ospf_sh_topology.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 865 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x4f, 0x6f, 0x23, 0x35,
	0x14, 0xd7, 0x6c, 0x4b, 0x93, 0x3a, 0x0d, 0x2c, 0x66, 0xc9, 0x4e, 0x10, 0xb0, 0x21, 0x48, 0x10,
	0x21, 0x34, 0x42, 0x6d, 0xf8, 0x7b, 0x41, 0x15, 0x6c, 0x44, 0x44, 0xb7, 0x5a, 0xa5, 0x5d, 0x24,
	0x4e, 0x96, 0xe3, 0xf1, 0x34, 0x96, 0x66, 0xec, 0x91, 0xed, 0x09, 0xc9, 0x87, 0xd9, 0x33, 0x87,
	0x3d, 0x20, 0x3e, 0x0a, 0x12, 0x07, 0xbe, 0x0d, 0xc8, 0xcf, 0x9e, 0xcc, 0xb4, 0x91, 0x38, 0x6f,
	0x2f, 0x23, 0xfb, 0xf7, 0x7b, 0xef, 0xf9, 0x3d, 0xfb, 0xfd, 0x19, 0x34, 0x50, 0xa6, 0xcc, 0x88,
	0x59, 0x11, 0xab, 0x4a, 0x95, 0xab, 0x9b, 0x6d, 0x52, 0x6a, 0x65, 0x15, 0x7e, 0x19, 0x31, 0x61,
	0x98, 0x22, 0x42, 0x19, 0xb2, 0xd1, 0x44, 0x94, 0xeb, 0x29, 0x01, 0x51, 0x55, 0x72, 0x9d, 0xb8,
	0x95, 0x13, 0x64, 0xdc, 0x18, 0x6e, 0xea, 0x55, 0xb2, 0xd6, 0x19, 0x7c, 0x12, 0xad, 0x2a, 0xcb,
	0x89, 0x90, 0x99, 0xd2, 0x05, 0xb5, 0x42, 0xc9, 0x80, 0x50, 0xcd, 0xa9, 0x69, 0xad, 0x93, 0xa2,
	0xca, 0xad, 0x60, 0xd4, 0x58, 0x22, 0xa4, 0xa5, 0xcc, 0x92, 0xb6, 0xdc, 0xff, 0x90, 0xe3, 0x57,
	0x11, 0x7a, 0xf7, 0xae, 0xeb, 0xe4, 0xe7, 0xa7, 0xbf, 0x5e, 0xe1, 0x8f, 0xd0, 0x49, 0xf0, 0x87,
	0x48, 0x5a, 0xf0, 0x38, 0x1a, 0x45, 0x93, 0xe3, 0x45, 0x2f, 0x60, 0x97, 0xb4, 0xe0, 0x78, 0x88,
	0xba, 0x6b, 0x9d, 0x79, 0xfa, 0x01, 0xd0, 0x9d, 0xb5, 0xce, 0x80, 0x7a, 0x8c, 0x3a, 0xce, 0x3e,
	0x11, 0x69, 0x7c, 0x30, 0x8a, 0x26, 0xfd, 0xc5, 0x91, 0xdb, 0xce, 0x53, 0x3c, 0x40, 0x47, 0xa5,
	0xe6, 0x99, 0xd8, 0xc4, 0x87, 0xa0, 0x11, 0x76, 0xf8, 0x63, 0xd4, 0xf7, 0x2b, 0x92, 0x73, 0x79,
	0x63, 0x57, 0xf1, 0x1b, 0xa0, 0x76, 0xe2, 0xc1, 0x0b, 0xc0, 0xc6, 0xff, 0x1e, 0xa2, 0x87, 0x77,
	0xbd, 0x75, 0x8e, 0xfa, 0x80, 0x82, 0xdd, 0x53, 0xef, 0x28, 0x60, 0xcf, 0xbd, 0xf1, 0x04, 0xbd,
	0xd3, 0x16, 0xa9, 0x8f, 0x38, 0x83, 0x23, 0xde, 0x6e, 0x49, 0xfa, 0x73, 0x1a, 0x93, 0x05, 0xb7,
	0x5a, 0xb0, 0x78, 0x0a, 0x82, 0xde, 0xe4, 0x33, 0x80, 0xf0, 0x07, 0x08, 0x79, 0x11, 0xbb, 0x2d,
	0x79, 0xfc, 0x25, 0x9c, 0x79, 0x0c, 0xc8, 0xf5, 0xb6, 0xe4, 0xf8, 0x53, 0xf4, 0x96, 0xa7, 0x99,
	0x92, 0x92, 0x33, 0xcb, 0xd3, 0xf8, 0xab, 0x51, 0x34, 0xe9, 0x2e, 0xde, 0x04, 0xf8, 0x87, 0x1a,
	0xc5, 0x7f, 0x45, 0xb5, 0x21, 0xf7, 0xcc, 0xf1, 0xd7, 0xa3, 0x68, 0xd2, 0x3b, 0xfd, 0x33, 0x4a,
	0x5e, 0xeb, 0xb4, 0x49, 0x5a, 0x8f, 0x40, 0x98, 0x2a, 0x0a, 0x25, 0x43, 0xf0, 0x73, 0x99, 0x29,
	0xfc, 0x4f, 0x54, 0x47, 0x5f, 0x52, 0xbb, 0x22, 0xb9, 0x30, 0x36, 0xfe, 0x66, 0x74, 0x30, 0xe9,
	0x9d, 0xfe, 0x71, 0x9f, 0x02, 0x73, 0xce, 0x2f, 0xfa, 0x3e, 0x3b, 0xa8, 0x5d, 0x5d, 0x08, 0x63,
	0xc7, 0x33, 0x74, 0xb2, 0x13, 0x11, 0x05, 0x77, 0xe9, 0x6c, 0x38, 0x53, 0x32, 0x85, 0xfa, 0xe8,
	0x2f, 0xc2, 0x0e, 0x7f, 0x88, 0x90, 0xa4, 0x52, 0x05, 0xee, 0x01, 0x70, 0x2d, 0x64, 0xfc, 0xb2,
	0x83, 0xf0, 0xfe, 0x25, 0xe2, 0x31, 0xea, 0x37, 0xce, 0xb8, 0xe2, 0x89, 0x5a, 0x99, 0x77, 0xee,
	0x2b, 0xe8, 0x93, 0xfa, 0x72, 0x9b, 0xfc, 0xf4, 0xf6, 0xbd, 0xea, 0x75, 0x9d, 0xa1, 0x9f, 0x21,
	0x9f, 0xd9, 0x44, 0x8b, 0x25, 0x59, 0x73, 0x6d, 0x84, 0x92, 0xa1, 0x18, 0xbd, 0x81, 0x85, 0x58,
	0xfe, 0xe2, 0xe1, 0x46, 0xd6, 0xb9, 0x54, 0xcb, 0xba, 0x02, 0x3d, 0x0c, 0xb2, 0x57, 0x65, 0x56,
	0xcb, 0x4e, 0xd1, 0xc0, 0xcb, 0x66, 0x4a, 0xff, 0x46, 0x75, 0x4a, 0x52, 0x61, 0x2c, 0x95, 0x8c,
	0x87, 0x92, 0x7d, 0x04, 0xec, 0xcc, 0x93, 0x3f, 0x06, 0xae, 0x29, 0x29, 0xa3, 0x2a, 0xcd, 0x78,
	0x7c, 0xd4, 0x0a, 0xec, 0x0a, 0x20, 0x97, 0x36, 0xc1, 0x8b, 0xaa, 0x4c, 0xa9, 0x0b, 0x50, 0x14,
	0x3c, 0xee, 0x40, 0x45, 0xbc, 0xba, 0x37, 0x89, 0x23, 0x0a, 0x1e, 0xee, 0xec, 0x05, 0x44, 0x71,
	0xed, 0xd2, 0xe4, 0xef, 0x5d, 0x45, 0x64, 0x54, 0xe4, 0x3e, 0xb0, 0xee, 0x3d, 0x0c, 0xcc, 0xa7,
	0xd8, 0x8c, 0x8a, 0x1c, 0xc2, 0xfa, 0x1c, 0xe1, 0x26, 0x6d, 0x4a, 0x2d, 0x94, 0x16, 0x76, 0x1b,
	0x1f, 0xc3, 0xd3, 0x3e, 0xac, 0xf3, 0xe6, 0x79, 0xc0, 0x9b, 0x2e, 0x4c, 0x2b, 0xab, 0x08, 0xdf,
	0xb0, 0xbc, 0x4a, 0x79, 0x1a, 0x23, 0xe8, 0x8b, 0xfe, 0xe5, 0xcf, 0x2b, 0xab, 0x9e, 0x06, 0x02,
	0x7f, 0x8f, 0xde, 0x0f, 0xd6, 0x75, 0xd3, 0xba, 0x35, 0xbf, 0x11, 0xc6, 0x72, 0xcd, 0xd3, 0xb8,
	0x07, 0x8a, 0x43, 0x7f, 0x8e, 0xae, 0x5b, 0xf8, 0x62, 0x27, 0x80, 0xbf, 0x45, 0xc3, 0x96, 0x01,
	0xb9, 0xd4, 0x6d, 0xed, 0x13, 0xf0, 0x72, 0xb0, 0xd3, 0xbe, 0x5c, 0xea, 0x46, 0x75, 0xfc, 0xfb,
	0xc1, 0xad, 0x49, 0x03, 0xbd, 0x00, 0x7f, 0x81, 0x1e, 0xd5, 0xf7, 0x6d, 0xb9, 0xce, 0x28, 0xe3,
	0xed, 0xd1, 0x88, 0x43, 0x03, 0x0c, 0x14, 0x8c, 0xc1, 0xb3, 0xba, 0x56, 0x24, 0xdf, 0x58, 0xb2,
	0x52, 0x25, 0xa1, 0x69, 0xaa, 0xb9, 0x31, 0x61, 0x5e, 0xfa, 0x0b, 0xb9, 0xe4, 0x1b, 0xfb, 0x93,
	0x2a, 0xcf, 0x3d, 0xb5, 0x57, 0x2a, 0x07, 0xad, 0x81, 0x16, 0x4a, 0xe5, 0x09, 0xf2, 0x5b, 0x92,
	0x1b, 0x2a, 0xd2, 0x30, 0x4a, 0xfd, 0x1c, 0xb9, 0x70, 0x08, 0xfe, 0x0e, 0xbd, 0xd7, 0xea, 0xc0,
	0xc2, 0x90, 0xa2, 0xf5, 0xbc, 0x50, 0xa8, 0xdd, 0x10, 0xbb, 0x6b, 0x6d, 0x73, 0xf3, 0xcc, 0xd1,
	0x73, 0x60, 0xf1, 0x14, 0x3d, 0xbe, 0xad, 0x5b, 0xb1, 0xc2, 0xdf, 0x00, 0x54, 0x6d, 0x37, 0x78,
	0xed, 0x15, 0x5f, 0xb0, 0xa2, 0x74, 0xab, 0xbd, 0x99, 0xd9, 0xd9, 0x9f, 0x99, 0x43, 0xd4, 0xcd,
	0x0d, 0xf5, 0x13, 0xb3, 0x0b, 0x74, 0x27, 0x37, 0x14, 0xe6, 0x65, 0xeb, 0x7f, 0xe1, 0xf8, 0xd6,
	0xff, 0xc2, 0x13, 0xd4, 0x03, 0xc2, 0x27, 0x78, 0x48, 0x16, 0xe4, 0xa0, 0x19, 0x20, 0xcb, 0x23,
	0xf8, 0xd1, 0x3a, 0xfb, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x42, 0x87, 0xe6, 0x82, 0x09, 0x00,
	0x00,
}
