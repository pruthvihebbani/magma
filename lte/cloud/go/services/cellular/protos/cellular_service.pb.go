// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cellular_service.proto

package protos

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

// (0)Turning off NonEPS service, (1)Both CSFB and SMS, (2)only SMS
type GatewayNonEPSConfig_NonEPSServiceControl int32

const (
	GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_OFF      GatewayNonEPSConfig_NonEPSServiceControl = 0
	GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_CSFB_SMS GatewayNonEPSConfig_NonEPSServiceControl = 1
	GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_SMS      GatewayNonEPSConfig_NonEPSServiceControl = 2
)

var GatewayNonEPSConfig_NonEPSServiceControl_name = map[int32]string{
	0: "NON_EPS_SERVICE_CONTROL_OFF",
	1: "NON_EPS_SERVICE_CONTROL_CSFB_SMS",
	2: "NON_EPS_SERVICE_CONTROL_SMS",
}
var GatewayNonEPSConfig_NonEPSServiceControl_value = map[string]int32{
	"NON_EPS_SERVICE_CONTROL_OFF":      0,
	"NON_EPS_SERVICE_CONTROL_CSFB_SMS": 1,
	"NON_EPS_SERVICE_CONTROL_SMS":      2,
}

func (x GatewayNonEPSConfig_NonEPSServiceControl) String() string {
	return proto.EnumName(GatewayNonEPSConfig_NonEPSServiceControl_name, int32(x))
}
func (GatewayNonEPSConfig_NonEPSServiceControl) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{3, 0}
}

type GatewayNonEPSConfig_CSFBRat int32

const (
	GatewayNonEPSConfig_CSFBRAT_2G GatewayNonEPSConfig_CSFBRat = 0
	GatewayNonEPSConfig_CSFBRAT_3G GatewayNonEPSConfig_CSFBRat = 1
)

var GatewayNonEPSConfig_CSFBRat_name = map[int32]string{
	0: "CSFBRAT_2G",
	1: "CSFBRAT_3G",
}
var GatewayNonEPSConfig_CSFBRat_value = map[string]int32{
	"CSFBRAT_2G": 0,
	"CSFBRAT_3G": 1,
}

func (x GatewayNonEPSConfig_CSFBRat) String() string {
	return proto.EnumName(GatewayNonEPSConfig_CSFBRat_name, int32(x))
}
func (GatewayNonEPSConfig_CSFBRat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{3, 1}
}

type NetworkEPCConfig_NetworkServices int32

const (
	NetworkEPCConfig_METERING    NetworkEPCConfig_NetworkServices = 0
	NetworkEPCConfig_DPI         NetworkEPCConfig_NetworkServices = 1
	NetworkEPCConfig_ENFORCEMENT NetworkEPCConfig_NetworkServices = 2
)

var NetworkEPCConfig_NetworkServices_name = map[int32]string{
	0: "METERING",
	1: "DPI",
	2: "ENFORCEMENT",
}
var NetworkEPCConfig_NetworkServices_value = map[string]int32{
	"METERING":    0,
	"DPI":         1,
	"ENFORCEMENT": 2,
}

func (x NetworkEPCConfig_NetworkServices) String() string {
	return proto.EnumName(NetworkEPCConfig_NetworkServices_name, int32(x))
}
func (NetworkEPCConfig_NetworkServices) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{6, 0}
}

type CellularGatewayConfig struct {
	Ran                  *GatewayRANConfig    `protobuf:"bytes,1,opt,name=ran,proto3" json:"ran,omitempty"`
	Epc                  *GatewayEPCConfig    `protobuf:"bytes,2,opt,name=epc,proto3" json:"epc,omitempty"`
	NonEpsService        *GatewayNonEPSConfig `protobuf:"bytes,3,opt,name=non_eps_service,json=nonEpsService,proto3" json:"non_eps_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CellularGatewayConfig) Reset()         { *m = CellularGatewayConfig{} }
func (m *CellularGatewayConfig) String() string { return proto.CompactTextString(m) }
func (*CellularGatewayConfig) ProtoMessage()    {}
func (*CellularGatewayConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{0}
}
func (m *CellularGatewayConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellularGatewayConfig.Unmarshal(m, b)
}
func (m *CellularGatewayConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellularGatewayConfig.Marshal(b, m, deterministic)
}
func (dst *CellularGatewayConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellularGatewayConfig.Merge(dst, src)
}
func (m *CellularGatewayConfig) XXX_Size() int {
	return xxx_messageInfo_CellularGatewayConfig.Size(m)
}
func (m *CellularGatewayConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CellularGatewayConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CellularGatewayConfig proto.InternalMessageInfo

func (m *CellularGatewayConfig) GetRan() *GatewayRANConfig {
	if m != nil {
		return m.Ran
	}
	return nil
}

func (m *CellularGatewayConfig) GetEpc() *GatewayEPCConfig {
	if m != nil {
		return m.Epc
	}
	return nil
}

func (m *CellularGatewayConfig) GetNonEpsService() *GatewayNonEPSConfig {
	if m != nil {
		return m.NonEpsService
	}
	return nil
}

type GatewayRANConfig struct {
	// Physical cell ID (0-504)
	Pci int32 `protobuf:"varint,1,opt,name=pci,proto3" json:"pci,omitempty"`
	// Enable eNodeB
	TransmitEnabled      bool     `protobuf:"varint,2,opt,name=transmit_enabled,json=transmitEnabled,proto3" json:"transmit_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayRANConfig) Reset()         { *m = GatewayRANConfig{} }
func (m *GatewayRANConfig) String() string { return proto.CompactTextString(m) }
func (*GatewayRANConfig) ProtoMessage()    {}
func (*GatewayRANConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{1}
}
func (m *GatewayRANConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayRANConfig.Unmarshal(m, b)
}
func (m *GatewayRANConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayRANConfig.Marshal(b, m, deterministic)
}
func (dst *GatewayRANConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayRANConfig.Merge(dst, src)
}
func (m *GatewayRANConfig) XXX_Size() int {
	return xxx_messageInfo_GatewayRANConfig.Size(m)
}
func (m *GatewayRANConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayRANConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayRANConfig proto.InternalMessageInfo

func (m *GatewayRANConfig) GetPci() int32 {
	if m != nil {
		return m.Pci
	}
	return 0
}

func (m *GatewayRANConfig) GetTransmitEnabled() bool {
	if m != nil {
		return m.TransmitEnabled
	}
	return false
}

type GatewayEPCConfig struct {
	NatEnabled bool `protobuf:"varint,1,opt,name=nat_enabled,json=natEnabled,proto3" json:"nat_enabled,omitempty"`
	// IP block is only set if nat_enabled is false
	// An IP block is a range of IP addresses specified by a network address and
	// a prefix-length of the netmask. For example,
	//    IPv4 IP block:      "192.168.0.0/24"
	//    IPv6 IP block:      "2401:db00:1116:301b::/64"
	IpBlock              string   `protobuf:"bytes,2,opt,name=ip_block,json=ipBlock,proto3" json:"ip_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayEPCConfig) Reset()         { *m = GatewayEPCConfig{} }
func (m *GatewayEPCConfig) String() string { return proto.CompactTextString(m) }
func (*GatewayEPCConfig) ProtoMessage()    {}
func (*GatewayEPCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{2}
}
func (m *GatewayEPCConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayEPCConfig.Unmarshal(m, b)
}
func (m *GatewayEPCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayEPCConfig.Marshal(b, m, deterministic)
}
func (dst *GatewayEPCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayEPCConfig.Merge(dst, src)
}
func (m *GatewayEPCConfig) XXX_Size() int {
	return xxx_messageInfo_GatewayEPCConfig.Size(m)
}
func (m *GatewayEPCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayEPCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayEPCConfig proto.InternalMessageInfo

func (m *GatewayEPCConfig) GetNatEnabled() bool {
	if m != nil {
		return m.NatEnabled
	}
	return false
}

func (m *GatewayEPCConfig) GetIpBlock() string {
	if m != nil {
		return m.IpBlock
	}
	return ""
}

type GatewayNonEPSConfig struct {
	// Mobile country code for CSFB
	CsfbMcc string `protobuf:"bytes,1,opt,name=csfb_mcc,json=csfbMcc,proto3" json:"csfb_mcc,omitempty"`
	// Mobile network code for CSFB
	CsfbMnc string `protobuf:"bytes,2,opt,name=csfb_mnc,json=csfbMnc,proto3" json:"csfb_mnc,omitempty"`
	// Location area code. 16-bit
	Lac int32 `protobuf:"varint,3,opt,name=lac,proto3" json:"lac,omitempty"`
	// RAT type (2G/3G), used by eNB
	CsfbRat GatewayNonEPSConfig_CSFBRat `protobuf:"varint,4,opt,name=csfb_rat,json=csfbRat,proto3,enum=magma.cellular.GatewayNonEPSConfig_CSFBRat" json:"csfb_rat,omitempty"`
	// 2G RAT frequencies /ARFCNs for redirection, used by eNB
	Arfcn_2G []int32 `protobuf:"varint,5,rep,packed,name=arfcn_2g,json=arfcn2g,proto3" json:"arfcn_2g,omitempty"`
	// For indicating one of the three modes
	NonEpsServiceControl GatewayNonEPSConfig_NonEPSServiceControl `protobuf:"varint,6,opt,name=non_eps_service_control,json=nonEpsServiceControl,proto3,enum=magma.cellular.GatewayNonEPSConfig_NonEPSServiceControl" json:"non_eps_service_control,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *GatewayNonEPSConfig) Reset()         { *m = GatewayNonEPSConfig{} }
func (m *GatewayNonEPSConfig) String() string { return proto.CompactTextString(m) }
func (*GatewayNonEPSConfig) ProtoMessage()    {}
func (*GatewayNonEPSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{3}
}
func (m *GatewayNonEPSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayNonEPSConfig.Unmarshal(m, b)
}
func (m *GatewayNonEPSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayNonEPSConfig.Marshal(b, m, deterministic)
}
func (dst *GatewayNonEPSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayNonEPSConfig.Merge(dst, src)
}
func (m *GatewayNonEPSConfig) XXX_Size() int {
	return xxx_messageInfo_GatewayNonEPSConfig.Size(m)
}
func (m *GatewayNonEPSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayNonEPSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayNonEPSConfig proto.InternalMessageInfo

func (m *GatewayNonEPSConfig) GetCsfbMcc() string {
	if m != nil {
		return m.CsfbMcc
	}
	return ""
}

func (m *GatewayNonEPSConfig) GetCsfbMnc() string {
	if m != nil {
		return m.CsfbMnc
	}
	return ""
}

func (m *GatewayNonEPSConfig) GetLac() int32 {
	if m != nil {
		return m.Lac
	}
	return 0
}

func (m *GatewayNonEPSConfig) GetCsfbRat() GatewayNonEPSConfig_CSFBRat {
	if m != nil {
		return m.CsfbRat
	}
	return GatewayNonEPSConfig_CSFBRAT_2G
}

func (m *GatewayNonEPSConfig) GetArfcn_2G() []int32 {
	if m != nil {
		return m.Arfcn_2G
	}
	return nil
}

func (m *GatewayNonEPSConfig) GetNonEpsServiceControl() GatewayNonEPSConfig_NonEPSServiceControl {
	if m != nil {
		return m.NonEpsServiceControl
	}
	return GatewayNonEPSConfig_NON_EPS_SERVICE_CONTROL_OFF
}

type CellularNetworkConfig struct {
	Ran                  *NetworkRANConfig `protobuf:"bytes,1,opt,name=ran,proto3" json:"ran,omitempty"`
	Epc                  *NetworkEPCConfig `protobuf:"bytes,2,opt,name=epc,proto3" json:"epc,omitempty"`
	FegNetworkId         string            `protobuf:"bytes,3,opt,name=feg_network_id,json=fegNetworkId,proto3" json:"feg_network_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CellularNetworkConfig) Reset()         { *m = CellularNetworkConfig{} }
func (m *CellularNetworkConfig) String() string { return proto.CompactTextString(m) }
func (*CellularNetworkConfig) ProtoMessage()    {}
func (*CellularNetworkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{4}
}
func (m *CellularNetworkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellularNetworkConfig.Unmarshal(m, b)
}
func (m *CellularNetworkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellularNetworkConfig.Marshal(b, m, deterministic)
}
func (dst *CellularNetworkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellularNetworkConfig.Merge(dst, src)
}
func (m *CellularNetworkConfig) XXX_Size() int {
	return xxx_messageInfo_CellularNetworkConfig.Size(m)
}
func (m *CellularNetworkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CellularNetworkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CellularNetworkConfig proto.InternalMessageInfo

func (m *CellularNetworkConfig) GetRan() *NetworkRANConfig {
	if m != nil {
		return m.Ran
	}
	return nil
}

func (m *CellularNetworkConfig) GetEpc() *NetworkEPCConfig {
	if m != nil {
		return m.Epc
	}
	return nil
}

func (m *CellularNetworkConfig) GetFegNetworkId() string {
	if m != nil {
		return m.FegNetworkId
	}
	return ""
}

type NetworkRANConfig struct {
	Earfcndl int32 `protobuf:"varint,1,opt,name=earfcndl,proto3" json:"earfcndl,omitempty"`
	// Bandwidth in MHz, from set {1.4, 3, 5, 10, 15, 20}
	BandwidthMhz           int32                       `protobuf:"varint,2,opt,name=bandwidth_mhz,json=bandwidthMhz,proto3" json:"bandwidth_mhz,omitempty"`
	SubframeAssignment     int32                       `protobuf:"varint,3,opt,name=subframe_assignment,json=subframeAssignment,proto3" json:"subframe_assignment,omitempty"`
	SpecialSubframePattern int32                       `protobuf:"varint,5,opt,name=special_subframe_pattern,json=specialSubframePattern,proto3" json:"special_subframe_pattern,omitempty"`
	TddConfig              *NetworkRANConfig_TDDConfig `protobuf:"bytes,6,opt,name=tdd_config,json=tddConfig,proto3" json:"tdd_config,omitempty"`
	FddConfig              *NetworkRANConfig_FDDConfig `protobuf:"bytes,7,opt,name=fdd_config,json=fddConfig,proto3" json:"fdd_config,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                    `json:"-"`
	XXX_unrecognized       []byte                      `json:"-"`
	XXX_sizecache          int32                       `json:"-"`
}

func (m *NetworkRANConfig) Reset()         { *m = NetworkRANConfig{} }
func (m *NetworkRANConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkRANConfig) ProtoMessage()    {}
func (*NetworkRANConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{5}
}
func (m *NetworkRANConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRANConfig.Unmarshal(m, b)
}
func (m *NetworkRANConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRANConfig.Marshal(b, m, deterministic)
}
func (dst *NetworkRANConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRANConfig.Merge(dst, src)
}
func (m *NetworkRANConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkRANConfig.Size(m)
}
func (m *NetworkRANConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRANConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRANConfig proto.InternalMessageInfo

func (m *NetworkRANConfig) GetEarfcndl() int32 {
	if m != nil {
		return m.Earfcndl
	}
	return 0
}

func (m *NetworkRANConfig) GetBandwidthMhz() int32 {
	if m != nil {
		return m.BandwidthMhz
	}
	return 0
}

func (m *NetworkRANConfig) GetSubframeAssignment() int32 {
	if m != nil {
		return m.SubframeAssignment
	}
	return 0
}

func (m *NetworkRANConfig) GetSpecialSubframePattern() int32 {
	if m != nil {
		return m.SpecialSubframePattern
	}
	return 0
}

func (m *NetworkRANConfig) GetTddConfig() *NetworkRANConfig_TDDConfig {
	if m != nil {
		return m.TddConfig
	}
	return nil
}

func (m *NetworkRANConfig) GetFddConfig() *NetworkRANConfig_FDDConfig {
	if m != nil {
		return m.FddConfig
	}
	return nil
}

type NetworkRANConfig_FDDConfig struct {
	Earfcndl             int32    `protobuf:"varint,1,opt,name=earfcndl,proto3" json:"earfcndl,omitempty"`
	Earfcnul             int32    `protobuf:"varint,2,opt,name=earfcnul,proto3" json:"earfcnul,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkRANConfig_FDDConfig) Reset()         { *m = NetworkRANConfig_FDDConfig{} }
func (m *NetworkRANConfig_FDDConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkRANConfig_FDDConfig) ProtoMessage()    {}
func (*NetworkRANConfig_FDDConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{5, 0}
}
func (m *NetworkRANConfig_FDDConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRANConfig_FDDConfig.Unmarshal(m, b)
}
func (m *NetworkRANConfig_FDDConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRANConfig_FDDConfig.Marshal(b, m, deterministic)
}
func (dst *NetworkRANConfig_FDDConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRANConfig_FDDConfig.Merge(dst, src)
}
func (m *NetworkRANConfig_FDDConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkRANConfig_FDDConfig.Size(m)
}
func (m *NetworkRANConfig_FDDConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRANConfig_FDDConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRANConfig_FDDConfig proto.InternalMessageInfo

func (m *NetworkRANConfig_FDDConfig) GetEarfcndl() int32 {
	if m != nil {
		return m.Earfcndl
	}
	return 0
}

func (m *NetworkRANConfig_FDDConfig) GetEarfcnul() int32 {
	if m != nil {
		return m.Earfcnul
	}
	return 0
}

type NetworkRANConfig_TDDConfig struct {
	Earfcndl int32 `protobuf:"varint,1,opt,name=earfcndl,proto3" json:"earfcndl,omitempty"`
	// TDD subframe config parameter. See http://niviuk.free.fr/lte_tdd.php
	SubframeAssignment int32 `protobuf:"varint,2,opt,name=subframe_assignment,json=subframeAssignment,proto3" json:"subframe_assignment,omitempty"`
	// TDD subframe config parameter. See http://niviuk.free.fr/lte_tdd.php
	SpecialSubframePattern int32    `protobuf:"varint,3,opt,name=special_subframe_pattern,json=specialSubframePattern,proto3" json:"special_subframe_pattern,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *NetworkRANConfig_TDDConfig) Reset()         { *m = NetworkRANConfig_TDDConfig{} }
func (m *NetworkRANConfig_TDDConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkRANConfig_TDDConfig) ProtoMessage()    {}
func (*NetworkRANConfig_TDDConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{5, 1}
}
func (m *NetworkRANConfig_TDDConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkRANConfig_TDDConfig.Unmarshal(m, b)
}
func (m *NetworkRANConfig_TDDConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkRANConfig_TDDConfig.Marshal(b, m, deterministic)
}
func (dst *NetworkRANConfig_TDDConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkRANConfig_TDDConfig.Merge(dst, src)
}
func (m *NetworkRANConfig_TDDConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkRANConfig_TDDConfig.Size(m)
}
func (m *NetworkRANConfig_TDDConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkRANConfig_TDDConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkRANConfig_TDDConfig proto.InternalMessageInfo

func (m *NetworkRANConfig_TDDConfig) GetEarfcndl() int32 {
	if m != nil {
		return m.Earfcndl
	}
	return 0
}

func (m *NetworkRANConfig_TDDConfig) GetSubframeAssignment() int32 {
	if m != nil {
		return m.SubframeAssignment
	}
	return 0
}

func (m *NetworkRANConfig_TDDConfig) GetSpecialSubframePattern() int32 {
	if m != nil {
		return m.SpecialSubframePattern
	}
	return 0
}

type NetworkEPCConfig struct {
	// always 3 digits
	// Mobile country code
	Mcc string `protobuf:"bytes,1,opt,name=mcc,proto3" json:"mcc,omitempty"`
	// 2-3 digits
	// Mobile network code
	Mnc string `protobuf:"bytes,2,opt,name=mnc,proto3" json:"mnc,omitempty"`
	// Tracking area code. 16-bit
	Tac int32 `protobuf:"varint,3,opt,name=tac,proto3" json:"tac,omitempty"`
	// Operator configuration field for LTE
	LteAuthOp []byte `protobuf:"bytes,4,opt,name=lte_auth_op,json=lteAuthOp,proto3" json:"lte_auth_op,omitempty"`
	// Authentication management field for LTE
	LteAuthAmf    []byte                                           `protobuf:"bytes,5,opt,name=lte_auth_amf,json=lteAuthAmf,proto3" json:"lte_auth_amf,omitempty"`
	SubProfiles   map[string]*NetworkEPCConfig_SubscriptionProfile `protobuf:"bytes,6,rep,name=sub_profiles,json=subProfiles,proto3" json:"sub_profiles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	DefaultRuleId string                                           `protobuf:"bytes,7,opt,name=default_rule_id,json=defaultRuleId,proto3" json:"default_rule_id,omitempty"`
	// Enable relaying S6a messages via FeG RPC
	RelayEnabled bool `protobuf:"varint,8,opt,name=relay_enabled,json=relayEnabled,proto3" json:"relay_enabled,omitempty"`
	// Ordered list of network service that will be enabled
	NetworkServices []NetworkEPCConfig_NetworkServices `protobuf:"varint,9,rep,packed,name=network_services,json=networkServices,proto3,enum=magma.cellular.NetworkEPCConfig_NetworkServices" json:"network_services,omitempty"`
	// If relay_enabled is false, this determines whether cloud subscriberdb
	// or local subscriberdb is used for authentication requests.
	CloudSubscriberdbEnabled bool     `protobuf:"varint,10,opt,name=cloud_subscriberdb_enabled,json=cloudSubscriberdbEnabled,proto3" json:"cloud_subscriberdb_enabled,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *NetworkEPCConfig) Reset()         { *m = NetworkEPCConfig{} }
func (m *NetworkEPCConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkEPCConfig) ProtoMessage()    {}
func (*NetworkEPCConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{6}
}
func (m *NetworkEPCConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkEPCConfig.Unmarshal(m, b)
}
func (m *NetworkEPCConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkEPCConfig.Marshal(b, m, deterministic)
}
func (dst *NetworkEPCConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkEPCConfig.Merge(dst, src)
}
func (m *NetworkEPCConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkEPCConfig.Size(m)
}
func (m *NetworkEPCConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkEPCConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkEPCConfig proto.InternalMessageInfo

func (m *NetworkEPCConfig) GetMcc() string {
	if m != nil {
		return m.Mcc
	}
	return ""
}

func (m *NetworkEPCConfig) GetMnc() string {
	if m != nil {
		return m.Mnc
	}
	return ""
}

func (m *NetworkEPCConfig) GetTac() int32 {
	if m != nil {
		return m.Tac
	}
	return 0
}

func (m *NetworkEPCConfig) GetLteAuthOp() []byte {
	if m != nil {
		return m.LteAuthOp
	}
	return nil
}

func (m *NetworkEPCConfig) GetLteAuthAmf() []byte {
	if m != nil {
		return m.LteAuthAmf
	}
	return nil
}

func (m *NetworkEPCConfig) GetSubProfiles() map[string]*NetworkEPCConfig_SubscriptionProfile {
	if m != nil {
		return m.SubProfiles
	}
	return nil
}

func (m *NetworkEPCConfig) GetDefaultRuleId() string {
	if m != nil {
		return m.DefaultRuleId
	}
	return ""
}

func (m *NetworkEPCConfig) GetRelayEnabled() bool {
	if m != nil {
		return m.RelayEnabled
	}
	return false
}

func (m *NetworkEPCConfig) GetNetworkServices() []NetworkEPCConfig_NetworkServices {
	if m != nil {
		return m.NetworkServices
	}
	return nil
}

func (m *NetworkEPCConfig) GetCloudSubscriberdbEnabled() bool {
	if m != nil {
		return m.CloudSubscriberdbEnabled
	}
	return false
}

type NetworkEPCConfig_SubscriptionProfile struct {
	// Maximum uplink bit rate (AMBR-UL)
	MaxUlBitRate uint64 `protobuf:"varint,1,opt,name=max_ul_bit_rate,json=maxUlBitRate,proto3" json:"max_ul_bit_rate,omitempty"`
	// Maximum downlink bit rate (AMBR-DL)
	MaxDlBitRate         uint64   `protobuf:"varint,2,opt,name=max_dl_bit_rate,json=maxDlBitRate,proto3" json:"max_dl_bit_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkEPCConfig_SubscriptionProfile) Reset()         { *m = NetworkEPCConfig_SubscriptionProfile{} }
func (m *NetworkEPCConfig_SubscriptionProfile) String() string { return proto.CompactTextString(m) }
func (*NetworkEPCConfig_SubscriptionProfile) ProtoMessage()    {}
func (*NetworkEPCConfig_SubscriptionProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_cellular_service_3d8910d284a7d277, []int{6, 0}
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Unmarshal(m, b)
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Marshal(b, m, deterministic)
}
func (dst *NetworkEPCConfig_SubscriptionProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Merge(dst, src)
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_Size() int {
	return xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.Size(m)
}
func (m *NetworkEPCConfig_SubscriptionProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkEPCConfig_SubscriptionProfile proto.InternalMessageInfo

func (m *NetworkEPCConfig_SubscriptionProfile) GetMaxUlBitRate() uint64 {
	if m != nil {
		return m.MaxUlBitRate
	}
	return 0
}

func (m *NetworkEPCConfig_SubscriptionProfile) GetMaxDlBitRate() uint64 {
	if m != nil {
		return m.MaxDlBitRate
	}
	return 0
}

func init() {
	proto.RegisterType((*CellularGatewayConfig)(nil), "magma.cellular.CellularGatewayConfig")
	proto.RegisterType((*GatewayRANConfig)(nil), "magma.cellular.GatewayRANConfig")
	proto.RegisterType((*GatewayEPCConfig)(nil), "magma.cellular.GatewayEPCConfig")
	proto.RegisterType((*GatewayNonEPSConfig)(nil), "magma.cellular.GatewayNonEPSConfig")
	proto.RegisterType((*CellularNetworkConfig)(nil), "magma.cellular.CellularNetworkConfig")
	proto.RegisterType((*NetworkRANConfig)(nil), "magma.cellular.NetworkRANConfig")
	proto.RegisterType((*NetworkRANConfig_FDDConfig)(nil), "magma.cellular.NetworkRANConfig.FDDConfig")
	proto.RegisterType((*NetworkRANConfig_TDDConfig)(nil), "magma.cellular.NetworkRANConfig.TDDConfig")
	proto.RegisterType((*NetworkEPCConfig)(nil), "magma.cellular.NetworkEPCConfig")
	proto.RegisterMapType((map[string]*NetworkEPCConfig_SubscriptionProfile)(nil), "magma.cellular.NetworkEPCConfig.SubProfilesEntry")
	proto.RegisterType((*NetworkEPCConfig_SubscriptionProfile)(nil), "magma.cellular.NetworkEPCConfig.SubscriptionProfile")
	proto.RegisterEnum("magma.cellular.GatewayNonEPSConfig_NonEPSServiceControl", GatewayNonEPSConfig_NonEPSServiceControl_name, GatewayNonEPSConfig_NonEPSServiceControl_value)
	proto.RegisterEnum("magma.cellular.GatewayNonEPSConfig_CSFBRat", GatewayNonEPSConfig_CSFBRat_name, GatewayNonEPSConfig_CSFBRat_value)
	proto.RegisterEnum("magma.cellular.NetworkEPCConfig_NetworkServices", NetworkEPCConfig_NetworkServices_name, NetworkEPCConfig_NetworkServices_value)
}

func init() {
	proto.RegisterFile("cellular_service.proto", fileDescriptor_cellular_service_3d8910d284a7d277)
}

var fileDescriptor_cellular_service_3d8910d284a7d277 = []byte{
	// 1002 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdf, 0x6e, 0xe3, 0x44,
	0x17, 0xdf, 0x24, 0x4d, 0x93, 0x9c, 0xa4, 0x8d, 0x35, 0xdd, 0x6f, 0x3f, 0x13, 0x24, 0x36, 0xca,
	0x2e, 0xa8, 0x0b, 0x52, 0x80, 0x2c, 0x17, 0x05, 0x71, 0xd3, 0xa6, 0x4e, 0x15, 0xa0, 0x49, 0x34,
	0x09, 0x5c, 0xc0, 0xc5, 0x68, 0x6c, 0x8f, 0x53, 0xab, 0xf6, 0xd8, 0xb2, 0xc7, 0xdb, 0xcd, 0x5e,
	0xf0, 0x10, 0xbc, 0x03, 0xb7, 0x3c, 0x09, 0x8f, 0x80, 0x78, 0x16, 0x34, 0x93, 0xb1, 0xd3, 0x46,
	0xdb, 0x36, 0x70, 0x95, 0x39, 0xe7, 0xfc, 0x7e, 0xc7, 0x67, 0xce, 0xbf, 0x09, 0x3c, 0x73, 0x58,
	0x10, 0x64, 0x01, 0x4d, 0x48, 0xca, 0x92, 0x37, 0xbe, 0xc3, 0xfa, 0x71, 0x12, 0x89, 0x08, 0x1d,
	0x86, 0x74, 0x19, 0xd2, 0x7e, 0x6e, 0xed, 0xfd, 0x59, 0x82, 0xff, 0x0d, 0xb5, 0x70, 0x41, 0x05,
	0xbb, 0xa1, 0xab, 0x61, 0xc4, 0x3d, 0x7f, 0x89, 0x06, 0x50, 0x49, 0x28, 0x37, 0x4b, 0xdd, 0xd2,
	0x71, 0x73, 0xd0, 0xed, 0xdf, 0xe5, 0xf5, 0x35, 0x16, 0x9f, 0x4e, 0xd6, 0x70, 0x2c, 0xc1, 0x92,
	0xc3, 0x62, 0xc7, 0x2c, 0x3f, 0xc8, 0xb1, 0x66, 0xc3, 0x9c, 0xc3, 0x62, 0x07, 0x7d, 0x0f, 0x6d,
	0x1e, 0x71, 0xc2, 0xe2, 0x34, 0x0f, 0xd5, 0xac, 0x28, 0xfe, 0x8b, 0x7b, 0xf8, 0x93, 0x88, 0x5b,
	0xb3, 0xb9, 0x76, 0x71, 0xc0, 0x23, 0x6e, 0xc5, 0xe9, 0x7c, 0xcd, 0xec, 0x4d, 0xc1, 0xd8, 0x8e,
	0x0c, 0x19, 0x50, 0x89, 0x1d, 0x5f, 0x5d, 0xa4, 0x8a, 0xe5, 0x11, 0xbd, 0x02, 0x43, 0x24, 0x94,
	0xa7, 0xa1, 0x2f, 0x08, 0xe3, 0xd4, 0x0e, 0x98, 0xab, 0x62, 0xae, 0xe3, 0x76, 0xae, 0xb7, 0xd6,
	0xea, 0xde, 0xa4, 0x70, 0x58, 0x84, 0x8d, 0x9e, 0x43, 0x93, 0xd3, 0x0d, 0xb3, 0xa4, 0x98, 0xc0,
	0x69, 0x4e, 0x42, 0x1f, 0x40, 0xdd, 0x8f, 0x89, 0x1d, 0x44, 0xce, 0xb5, 0xf2, 0xdb, 0xc0, 0x35,
	0x3f, 0x3e, 0x93, 0x62, 0xef, 0xaf, 0x0a, 0x1c, 0xbd, 0xe7, 0x1e, 0x92, 0xe2, 0xa4, 0x9e, 0x4d,
	0x42, 0xc7, 0x51, 0x0e, 0x1b, 0xb8, 0x26, 0xe5, 0x4b, 0xc7, 0xd9, 0x98, 0xb8, 0x93, 0x7b, 0x53,
	0x26, 0xee, 0xc8, 0xab, 0x05, 0xd4, 0x51, 0xf9, 0xaa, 0x62, 0x79, 0x44, 0x23, 0x0d, 0x4e, 0xa8,
	0x30, 0xf7, 0xba, 0xa5, 0xe3, 0xc3, 0xc1, 0x67, 0x3b, 0xa4, 0xb1, 0x3f, 0x9c, 0x8f, 0xce, 0x30,
	0x15, 0x6b, 0xcf, 0x98, 0x0a, 0xf9, 0x51, 0x9a, 0x78, 0x0e, 0x27, 0x83, 0xa5, 0x59, 0xed, 0x56,
	0x8e, 0xab, 0xb8, 0xa6, 0xe4, 0xc1, 0x12, 0x45, 0xf0, 0xff, 0xad, 0x82, 0x11, 0x27, 0xe2, 0x22,
	0x89, 0x02, 0x73, 0x5f, 0x7d, 0xf1, 0x64, 0x97, 0x2f, 0xae, 0x05, 0x5d, 0xb7, 0xe1, 0x9a, 0x8f,
	0x9f, 0xde, 0xa9, 0xa6, 0xd6, 0xf6, 0x7e, 0x85, 0xa7, 0xef, 0x43, 0xa3, 0xe7, 0xf0, 0xe1, 0x64,
	0x3a, 0x21, 0xd6, 0x6c, 0x4e, 0xe6, 0x16, 0xfe, 0x69, 0x3c, 0xb4, 0xc8, 0x70, 0x3a, 0x59, 0xe0,
	0xe9, 0x0f, 0x64, 0x3a, 0x1a, 0x19, 0x4f, 0xd0, 0x4b, 0xe8, 0xde, 0x07, 0x90, 0x17, 0x26, 0xf3,
	0xcb, 0xb9, 0x51, 0x7a, 0xc8, 0x8d, 0x04, 0x94, 0x7b, 0xaf, 0xa0, 0xa6, 0xf3, 0x83, 0x0e, 0x01,
	0xd4, 0xf1, 0x74, 0x41, 0x06, 0x17, 0xc6, 0x93, 0xdb, 0xf2, 0xeb, 0x0b, 0xa3, 0xd4, 0xfb, 0xfd,
	0xd6, 0x38, 0x4d, 0x98, 0xb8, 0x89, 0x92, 0xeb, 0x9d, 0xc6, 0x49, 0x63, 0xff, 0xd5, 0x38, 0x69,
	0xce, 0xd6, 0x38, 0xbd, 0x84, 0x43, 0x8f, 0x2d, 0x09, 0x5f, 0x1b, 0x89, 0xef, 0xaa, 0xee, 0x68,
	0xe0, 0x96, 0xc7, 0x96, 0x9a, 0x31, 0x76, 0x7b, 0x7f, 0xec, 0x81, 0xb1, 0xfd, 0x4d, 0xd4, 0x81,
	0x3a, 0x53, 0x45, 0x76, 0x03, 0x3d, 0x2d, 0x85, 0x8c, 0x5e, 0xc0, 0x81, 0x4d, 0xb9, 0x7b, 0xe3,
	0xbb, 0xe2, 0x8a, 0x84, 0x57, 0xef, 0x54, 0x50, 0x55, 0xdc, 0x2a, 0x94, 0x97, 0x57, 0xef, 0xd0,
	0xe7, 0x70, 0x94, 0x66, 0xb6, 0x97, 0xd0, 0x90, 0x11, 0x9a, 0xa6, 0xfe, 0x92, 0x87, 0x8c, 0x0b,
	0xdd, 0x9e, 0x28, 0x37, 0x9d, 0x16, 0x16, 0x74, 0x02, 0x66, 0x1a, 0x33, 0xc7, 0xa7, 0x01, 0x29,
	0x88, 0x31, 0x15, 0x82, 0x25, 0xdc, 0xac, 0x2a, 0xd6, 0x33, 0x6d, 0x9f, 0x6b, 0xf3, 0x6c, 0x6d,
	0x45, 0x63, 0x00, 0xe1, 0xba, 0xb2, 0xf1, 0x3c, 0x7f, 0xa9, 0xfa, 0xae, 0x39, 0xf8, 0xf4, 0xb1,
	0xac, 0xf6, 0x17, 0xe7, 0xe7, 0x3a, 0x57, 0x0d, 0xe1, 0xba, 0xfa, 0xda, 0x63, 0x00, 0x6f, 0xe3,
	0xaa, 0xb6, 0xa3, 0xab, 0xd1, 0xc6, 0x95, 0x97, 0xbb, 0xea, 0x0c, 0xa1, 0x51, 0xe8, 0x1f, 0x4c,
	0x67, 0x61, 0xcb, 0x02, 0x9d, 0xc9, 0x42, 0xee, 0xfc, 0x56, 0x82, 0xc6, 0x62, 0x27, 0x2f, 0xf7,
	0xe4, 0xbb, 0xfc, 0x9f, 0xf2, 0x5d, 0x79, 0x28, 0xdf, 0xbd, 0xbf, 0xab, 0x45, 0xc3, 0x6c, 0x16,
	0xa1, 0x01, 0x95, 0xcd, 0xbe, 0x92, 0x47, 0xa5, 0x29, 0xd6, 0x94, 0x3c, 0x4a, 0x8d, 0xd8, 0xac,
	0x28, 0x41, 0x1d, 0xf4, 0x11, 0x34, 0x03, 0xc1, 0x08, 0xcd, 0xc4, 0x15, 0x89, 0x62, 0xb5, 0xa5,
	0x5a, 0xb8, 0x11, 0x08, 0x76, 0x9a, 0x89, 0xab, 0x69, 0x8c, 0xba, 0xd0, 0x2a, 0xec, 0x34, 0xf4,
	0x54, 0x23, 0xb4, 0x30, 0x68, 0xc0, 0x69, 0xe8, 0xa1, 0x05, 0xb4, 0xd2, 0xcc, 0x26, 0x71, 0x12,
	0x79, 0x7e, 0xc0, 0x52, 0x73, 0xbf, 0x5b, 0x39, 0x6e, 0x0e, 0xbe, 0x7c, 0x6c, 0x40, 0xfa, 0xf3,
	0xcc, 0x9e, 0x69, 0x8e, 0xc5, 0x45, 0xb2, 0xc2, 0xcd, 0x74, 0xa3, 0x41, 0x9f, 0x40, 0xdb, 0x65,
	0x1e, 0xcd, 0x02, 0x41, 0x92, 0x2c, 0x60, 0x72, 0x74, 0x6a, 0xea, 0x1e, 0x07, 0x5a, 0x8d, 0xb3,
	0x80, 0x8d, 0x5d, 0x39, 0x0a, 0x09, 0x0b, 0xe8, 0xaa, 0x78, 0x00, 0xea, 0xea, 0x01, 0x68, 0x29,
	0x65, 0xfe, 0x04, 0xfc, 0x02, 0x46, 0x3e, 0x82, 0x7a, 0x49, 0xa6, 0x66, 0xa3, 0x5b, 0x39, 0x3e,
	0x1c, 0x7c, 0xf1, 0x68, 0x98, 0x5a, 0xa1, 0xb7, 0x5d, 0x8a, 0xdb, 0xfc, 0xae, 0x02, 0x7d, 0x0b,
	0x1d, 0x27, 0x88, 0x32, 0x57, 0x16, 0x31, 0x75, 0x12, 0xdf, 0x66, 0x89, 0x6b, 0x17, 0xe1, 0x80,
	0x0a, 0xc7, 0x54, 0x88, 0xf9, 0x2d, 0x80, 0x0e, 0xad, 0xe3, 0xc0, 0x91, 0x56, 0xc7, 0xc2, 0x8f,
	0xb8, 0xbe, 0x3f, 0xfa, 0x18, 0xda, 0x21, 0x7d, 0x4b, 0xb2, 0x80, 0xd8, 0xbe, 0x90, 0xef, 0x07,
	0x53, 0x85, 0xdd, 0xc3, 0xad, 0x90, 0xbe, 0xfd, 0x31, 0x38, 0xf3, 0x05, 0xa6, 0xa2, 0x80, 0xb9,
	0xb7, 0x60, 0xe5, 0x02, 0x76, 0x9e, 0xc3, 0x3a, 0x02, 0x8c, 0xed, 0x6c, 0xcb, 0x56, 0xb8, 0x66,
	0xab, 0xbc, 0x5d, 0xae, 0xd9, 0x0a, 0x7d, 0x07, 0xd5, 0x37, 0x34, 0xc8, 0x98, 0x5e, 0x71, 0x5f,
	0xed, 0x52, 0xc1, 0xed, 0xc0, 0xf1, 0xda, 0xc5, 0x37, 0xe5, 0x93, 0x52, 0xef, 0x6b, 0x68, 0x6f,
	0x25, 0x0f, 0xb5, 0xa0, 0x7e, 0x69, 0x2d, 0x2c, 0x3c, 0x9e, 0xc8, 0x7d, 0x5d, 0x83, 0xca, 0xf9,
	0x6c, 0x6c, 0x94, 0x50, 0x1b, 0x9a, 0xd6, 0x64, 0x34, 0xc5, 0x43, 0xeb, 0xd2, 0x9a, 0x2c, 0x8c,
	0xf2, 0x59, 0xfd, 0xe7, 0x7d, 0xf5, 0x0f, 0x29, 0xb5, 0xd7, 0xbf, 0xaf, 0xff, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x9d, 0x1d, 0x4f, 0x79, 0x43, 0x09, 0x00, 0x00,
}
