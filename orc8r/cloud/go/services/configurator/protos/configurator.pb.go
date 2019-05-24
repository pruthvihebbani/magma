// Code generated by protoc-gen-go. DO NOT EDIT.
// source: configurator.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OWN grants READ and WRITE
type ACL_Permission int32

const (
	ACL_NO_PERM ACL_Permission = 0
	ACL_READ    ACL_Permission = 1
	ACL_WRITE   ACL_Permission = 2
	ACL_OWN     ACL_Permission = 3
)

var ACL_Permission_name = map[int32]string{
	0: "NO_PERM",
	1: "READ",
	2: "WRITE",
	3: "OWN",
}
var ACL_Permission_value = map[string]int32{
	"NO_PERM": 0,
	"READ":    1,
	"WRITE":   2,
	"OWN":     3,
}

func (x ACL_Permission) String() string {
	return proto.EnumName(ACL_Permission_name, int32(x))
}
func (ACL_Permission) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{4, 0}
}

type ACL_Wildcard int32

const (
	ACL_NO_WILDCARD  ACL_Wildcard = 0
	ACL_WILDCARD_ALL ACL_Wildcard = 1
)

var ACL_Wildcard_name = map[int32]string{
	0: "NO_WILDCARD",
	1: "WILDCARD_ALL",
}
var ACL_Wildcard_value = map[string]int32{
	"NO_WILDCARD":  0,
	"WILDCARD_ALL": 1,
}

func (x ACL_Wildcard) String() string {
	return proto.EnumName(ACL_Wildcard_name, int32(x))
}
func (ACL_Wildcard) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{4, 1}
}

// Network is the core tenancy concept in configurator. A network can have
// configurations which will hierarchically apply to entities in the network.
// Entities which don't belong to any specific tenant will be organized under
// the hood into an internal-only network.
type Network struct {
	// Network ID is unique across all tenants
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Description          string           `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	Configs              []*NetworkConfig `protobuf:"bytes,20,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Network) Reset()         { *m = Network{} }
func (m *Network) String() string { return proto.CompactTextString(m) }
func (*Network) ProtoMessage()    {}
func (*Network) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{0}
}
func (m *Network) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Network.Unmarshal(m, b)
}
func (m *Network) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Network.Marshal(b, m, deterministic)
}
func (dst *Network) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Network.Merge(dst, src)
}
func (m *Network) XXX_Size() int {
	return xxx_messageInfo_Network.Size(m)
}
func (m *Network) XXX_DiscardUnknown() {
	xxx_messageInfo_Network.DiscardUnknown(m)
}

var xxx_messageInfo_Network proto.InternalMessageInfo

func (m *Network) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Network) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Network) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Network) GetConfigs() []*NetworkConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

// The network entity is the core entity managed by configurator. A network
// entity can correspond to a physical asset like an access gateway or radio,
// in which case the physical_id field will be populated. A network entity can
// also represent an intangible part of a network like a configuration profile,
// a mesh, an API user, etc.
// Inside a network, network entities are organized into DAGs. Associations
// between entities are one-way and define a hierarchical relationship (e.g.,
// a mesh would be associated towards entities within the mesh).
type NetworkEntity struct {
	// (id, type) is a unique identifier for an entity within a network
	// type denotes the category of the network entity.
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name        string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,11,opt,name=description,proto3" json:"description,omitempty"`
	// physical_id ties the network entity back to a physical asset. This
	// field is optional
	PhysicalId string `protobuf:"bytes,20,opt,name=physical_id,json=physicalId,proto3" json:"physical_id,omitempty"`
	// Serialized representation of the network entity's configuration.
	// The value of the type field will point to the Serde implementation to
	// use to deserialize this field.
	Config  []byte `protobuf:"bytes,30,opt,name=config,proto3" json:"config,omitempty"`
	GraphID string `protobuf:"bytes,40,opt,name=graphID,proto3" json:"graphID,omitempty"`
	// assocs represents the related network entities as an adjacency list
	Assocs               []*EntityID `protobuf:"bytes,50,rep,name=assocs,proto3" json:"assocs,omitempty"`
	ParentAssocs         []*EntityID `protobuf:"bytes,60,rep,name=parent_assocs,json=parentAssocs,proto3" json:"parent_assocs,omitempty"`
	Permissions          []*ACL      `protobuf:"bytes,70,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NetworkEntity) Reset()         { *m = NetworkEntity{} }
func (m *NetworkEntity) String() string { return proto.CompactTextString(m) }
func (*NetworkEntity) ProtoMessage()    {}
func (*NetworkEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{1}
}
func (m *NetworkEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkEntity.Unmarshal(m, b)
}
func (m *NetworkEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkEntity.Marshal(b, m, deterministic)
}
func (dst *NetworkEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkEntity.Merge(dst, src)
}
func (m *NetworkEntity) XXX_Size() int {
	return xxx_messageInfo_NetworkEntity.Size(m)
}
func (m *NetworkEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkEntity.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkEntity proto.InternalMessageInfo

func (m *NetworkEntity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NetworkEntity) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NetworkEntity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *NetworkEntity) GetPhysicalId() string {
	if m != nil {
		return m.PhysicalId
	}
	return ""
}

func (m *NetworkEntity) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *NetworkEntity) GetGraphID() string {
	if m != nil {
		return m.GraphID
	}
	return ""
}

func (m *NetworkEntity) GetAssocs() []*EntityID {
	if m != nil {
		return m.Assocs
	}
	return nil
}

func (m *NetworkEntity) GetParentAssocs() []*EntityID {
	if m != nil {
		return m.ParentAssocs
	}
	return nil
}

func (m *NetworkEntity) GetPermissions() []*ACL {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type NetworkConfig struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkConfig) Reset()         { *m = NetworkConfig{} }
func (m *NetworkConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkConfig) ProtoMessage()    {}
func (*NetworkConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{2}
}
func (m *NetworkConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkConfig.Unmarshal(m, b)
}
func (m *NetworkConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkConfig.Marshal(b, m, deterministic)
}
func (dst *NetworkConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkConfig.Merge(dst, src)
}
func (m *NetworkConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkConfig.Size(m)
}
func (m *NetworkConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkConfig proto.InternalMessageInfo

func (m *NetworkConfig) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkConfig) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type EntityID struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntityID) Reset()         { *m = EntityID{} }
func (m *EntityID) String() string { return proto.CompactTextString(m) }
func (*EntityID) ProtoMessage()    {}
func (*EntityID) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{3}
}
func (m *EntityID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityID.Unmarshal(m, b)
}
func (m *EntityID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityID.Marshal(b, m, deterministic)
}
func (dst *EntityID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityID.Merge(dst, src)
}
func (m *EntityID) XXX_Size() int {
	return xxx_messageInfo_EntityID.Size(m)
}
func (m *EntityID) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityID.DiscardUnknown(m)
}

var xxx_messageInfo_EntityID proto.InternalMessageInfo

func (m *EntityID) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EntityID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// ACL specifies R/W/O permissions on a specific type of entity or a wildcard.
// Each ACL is bound to one or more networks, but may also be globally
// wildcarded (e.g. admin operators).
type ACL struct {
	// Unique identifier for this ACL
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// scope defines the networks to which this ACL applies, or a wildcard
	// specifying all networks
	//
	// Types that are valid to be assigned to Scope:
	//	*ACL_NetworkIds
	//	*ACL_ScopeWildcard
	Scope      isACL_Scope    `protobuf_oneof:"scope"`
	Permission ACL_Permission `protobuf:"varint,20,opt,name=permission,proto3,enum=magma.orc8r.configurator.ACL_Permission" json:"permission,omitempty"`
	// An ACL either applies to a specific entity type or all entities within
	// the scope
	//
	// Types that are valid to be assigned to Type:
	//	*ACL_EntityType
	//	*ACL_TypeWildcard
	Type isACL_Type `protobuf_oneof:"type"`
	// Optionally, a list of specific entity IDs on which this ACL applies.
	// If empty, the ACL will apply to all entities of the specified type or
	// wildcard within the scope.
	IdFilter             []string `protobuf:"bytes,40,rep,name=id_filter,json=idFilter,proto3" json:"id_filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACL) Reset()         { *m = ACL{} }
func (m *ACL) String() string { return proto.CompactTextString(m) }
func (*ACL) ProtoMessage()    {}
func (*ACL) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{4}
}
func (m *ACL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACL.Unmarshal(m, b)
}
func (m *ACL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACL.Marshal(b, m, deterministic)
}
func (dst *ACL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACL.Merge(dst, src)
}
func (m *ACL) XXX_Size() int {
	return xxx_messageInfo_ACL.Size(m)
}
func (m *ACL) XXX_DiscardUnknown() {
	xxx_messageInfo_ACL.DiscardUnknown(m)
}

var xxx_messageInfo_ACL proto.InternalMessageInfo

func (m *ACL) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isACL_Scope interface {
	isACL_Scope()
}

type ACL_NetworkIds struct {
	NetworkIds *ACL_NetworkIDs `protobuf:"bytes,10,opt,name=network_ids,json=networkIds,proto3,oneof"`
}

type ACL_ScopeWildcard struct {
	ScopeWildcard ACL_Wildcard `protobuf:"varint,11,opt,name=scope_wildcard,json=scopeWildcard,proto3,enum=magma.orc8r.configurator.ACL_Wildcard,oneof"`
}

func (*ACL_NetworkIds) isACL_Scope() {}

func (*ACL_ScopeWildcard) isACL_Scope() {}

func (m *ACL) GetScope() isACL_Scope {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (m *ACL) GetNetworkIds() *ACL_NetworkIDs {
	if x, ok := m.GetScope().(*ACL_NetworkIds); ok {
		return x.NetworkIds
	}
	return nil
}

func (m *ACL) GetScopeWildcard() ACL_Wildcard {
	if x, ok := m.GetScope().(*ACL_ScopeWildcard); ok {
		return x.ScopeWildcard
	}
	return ACL_NO_WILDCARD
}

func (m *ACL) GetPermission() ACL_Permission {
	if m != nil {
		return m.Permission
	}
	return ACL_NO_PERM
}

type isACL_Type interface {
	isACL_Type()
}

type ACL_EntityType struct {
	EntityType string `protobuf:"bytes,30,opt,name=entity_type,json=entityType,proto3,oneof"`
}

type ACL_TypeWildcard struct {
	TypeWildcard ACL_Wildcard `protobuf:"varint,31,opt,name=type_wildcard,json=typeWildcard,proto3,enum=magma.orc8r.configurator.ACL_Wildcard,oneof"`
}

func (*ACL_EntityType) isACL_Type() {}

func (*ACL_TypeWildcard) isACL_Type() {}

func (m *ACL) GetType() isACL_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *ACL) GetEntityType() string {
	if x, ok := m.GetType().(*ACL_EntityType); ok {
		return x.EntityType
	}
	return ""
}

func (m *ACL) GetTypeWildcard() ACL_Wildcard {
	if x, ok := m.GetType().(*ACL_TypeWildcard); ok {
		return x.TypeWildcard
	}
	return ACL_NO_WILDCARD
}

func (m *ACL) GetIdFilter() []string {
	if m != nil {
		return m.IdFilter
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ACL) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ACL_OneofMarshaler, _ACL_OneofUnmarshaler, _ACL_OneofSizer, []interface{}{
		(*ACL_NetworkIds)(nil),
		(*ACL_ScopeWildcard)(nil),
		(*ACL_EntityType)(nil),
		(*ACL_TypeWildcard)(nil),
	}
}

func _ACL_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ACL)
	// scope
	switch x := m.Scope.(type) {
	case *ACL_NetworkIds:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NetworkIds); err != nil {
			return err
		}
	case *ACL_ScopeWildcard:
		b.EncodeVarint(11<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ScopeWildcard))
	case nil:
	default:
		return fmt.Errorf("ACL.Scope has unexpected type %T", x)
	}
	// type
	switch x := m.Type.(type) {
	case *ACL_EntityType:
		b.EncodeVarint(30<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.EntityType)
	case *ACL_TypeWildcard:
		b.EncodeVarint(31<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.TypeWildcard))
	case nil:
	default:
		return fmt.Errorf("ACL.Type has unexpected type %T", x)
	}
	return nil
}

func _ACL_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ACL)
	switch tag {
	case 10: // scope.network_ids
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ACL_NetworkIDs)
		err := b.DecodeMessage(msg)
		m.Scope = &ACL_NetworkIds{msg}
		return true, err
	case 11: // scope.scope_wildcard
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Scope = &ACL_ScopeWildcard{ACL_Wildcard(x)}
		return true, err
	case 30: // type.entity_type
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Type = &ACL_EntityType{x}
		return true, err
	case 31: // type.type_wildcard
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &ACL_TypeWildcard{ACL_Wildcard(x)}
		return true, err
	default:
		return false, nil
	}
}

func _ACL_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ACL)
	// scope
	switch x := m.Scope.(type) {
	case *ACL_NetworkIds:
		s := proto.Size(x.NetworkIds)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ACL_ScopeWildcard:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ScopeWildcard))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	// type
	switch x := m.Type.(type) {
	case *ACL_EntityType:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(len(x.EntityType)))
		n += len(x.EntityType)
	case *ACL_TypeWildcard:
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(x.TypeWildcard))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ACL_NetworkIDs struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACL_NetworkIDs) Reset()         { *m = ACL_NetworkIDs{} }
func (m *ACL_NetworkIDs) String() string { return proto.CompactTextString(m) }
func (*ACL_NetworkIDs) ProtoMessage()    {}
func (*ACL_NetworkIDs) Descriptor() ([]byte, []int) {
	return fileDescriptor_configurator_320aa35db679fe65, []int{4, 0}
}
func (m *ACL_NetworkIDs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACL_NetworkIDs.Unmarshal(m, b)
}
func (m *ACL_NetworkIDs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACL_NetworkIDs.Marshal(b, m, deterministic)
}
func (dst *ACL_NetworkIDs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACL_NetworkIDs.Merge(dst, src)
}
func (m *ACL_NetworkIDs) XXX_Size() int {
	return xxx_messageInfo_ACL_NetworkIDs.Size(m)
}
func (m *ACL_NetworkIDs) XXX_DiscardUnknown() {
	xxx_messageInfo_ACL_NetworkIDs.DiscardUnknown(m)
}

var xxx_messageInfo_ACL_NetworkIDs proto.InternalMessageInfo

func (m *ACL_NetworkIDs) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func init() {
	proto.RegisterType((*Network)(nil), "magma.orc8r.configurator.Network")
	proto.RegisterType((*NetworkEntity)(nil), "magma.orc8r.configurator.NetworkEntity")
	proto.RegisterType((*NetworkConfig)(nil), "magma.orc8r.configurator.NetworkConfig")
	proto.RegisterType((*EntityID)(nil), "magma.orc8r.configurator.EntityID")
	proto.RegisterType((*ACL)(nil), "magma.orc8r.configurator.ACL")
	proto.RegisterType((*ACL_NetworkIDs)(nil), "magma.orc8r.configurator.ACL.NetworkIDs")
	proto.RegisterEnum("magma.orc8r.configurator.ACL_Permission", ACL_Permission_name, ACL_Permission_value)
	proto.RegisterEnum("magma.orc8r.configurator.ACL_Wildcard", ACL_Wildcard_name, ACL_Wildcard_value)
}

func init() { proto.RegisterFile("configurator.proto", fileDescriptor_configurator_320aa35db679fe65) }

var fileDescriptor_configurator_320aa35db679fe65 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcb, 0x6a, 0xdb, 0x40,
	0x14, 0xb5, 0xac, 0xc4, 0x8f, 0x2b, 0xdb, 0x15, 0x43, 0x28, 0xd3, 0x94, 0x26, 0xae, 0x16, 0xad,
	0x36, 0x75, 0xc0, 0x5d, 0xf4, 0x41, 0xa1, 0x28, 0x96, 0x53, 0x8b, 0x3a, 0x76, 0x18, 0x02, 0x86,
	0x6e, 0x84, 0x22, 0x29, 0xce, 0x50, 0x5b, 0x23, 0x66, 0x94, 0x06, 0x6f, 0xfb, 0x09, 0xfd, 0x87,
	0xfe, 0x67, 0xd1, 0x8c, 0x64, 0x1b, 0x1a, 0xa7, 0x8f, 0x95, 0xe6, 0x3e, 0xce, 0x9d, 0x33, 0xe7,
	0x5c, 0x04, 0x28, 0x64, 0xc9, 0x35, 0x9d, 0xdf, 0xf2, 0x20, 0x63, 0xbc, 0x97, 0x72, 0x96, 0x31,
	0x84, 0x97, 0xc1, 0x7c, 0x19, 0xf4, 0x18, 0x0f, 0xdf, 0xf2, 0xde, 0x76, 0xfd, 0xf0, 0xc9, 0x9c,
	0xb1, 0xf9, 0x22, 0x3e, 0x91, 0x7d, 0x57, 0xb7, 0xd7, 0x27, 0x41, 0xb2, 0x52, 0x20, 0xeb, 0x87,
	0x06, 0xf5, 0x49, 0x9c, 0xdd, 0x31, 0xfe, 0x15, 0x75, 0xa0, 0x4a, 0x23, 0xac, 0x75, 0x35, 0xbb,
	0x49, 0xaa, 0x34, 0x42, 0x08, 0xf6, 0x92, 0x60, 0x19, 0x63, 0x90, 0x19, 0x79, 0x46, 0x5d, 0x30,
	0xa2, 0x58, 0x84, 0x9c, 0xa6, 0x19, 0x65, 0x09, 0x36, 0x64, 0x69, 0x3b, 0x85, 0x1c, 0xa8, 0xab,
	0xcb, 0x05, 0x3e, 0xe8, 0xea, 0xb6, 0xd1, 0x7f, 0xd9, 0xdb, 0x45, 0xac, 0x57, 0xdc, 0x3c, 0x90,
	0x39, 0x52, 0xe2, 0xac, 0xef, 0x3a, 0xb4, 0x8b, 0xd2, 0x30, 0xc9, 0x68, 0xb6, 0xba, 0x8f, 0x5a,
	0xb6, 0x4a, 0x63, 0x5c, 0x55, 0xd4, 0xf2, 0xf3, 0x7f, 0xd2, 0x3d, 0x06, 0x23, 0xbd, 0x59, 0x09,
	0x1a, 0x06, 0x0b, 0x9f, 0x46, 0xf8, 0x40, 0x76, 0x40, 0x99, 0xf2, 0x22, 0xf4, 0x18, 0x6a, 0x8a,
	0x17, 0x3e, 0xea, 0x6a, 0x76, 0x8b, 0x14, 0x11, 0xc2, 0x50, 0x9f, 0xf3, 0x20, 0xbd, 0xf1, 0x5c,
	0x6c, 0x4b, 0x50, 0x19, 0xa2, 0xf7, 0x50, 0x0b, 0x84, 0x60, 0xa1, 0xc0, 0x7d, 0x29, 0x80, 0xb5,
	0x5b, 0x00, 0xf5, 0x3c, 0xcf, 0x25, 0x05, 0x02, 0x7d, 0x82, 0x76, 0x1a, 0xf0, 0x38, 0xc9, 0xfc,
	0x62, 0xc4, 0x87, 0xbf, 0x1e, 0xd1, 0x52, 0x40, 0x47, 0x0d, 0xfa, 0x08, 0x46, 0x1a, 0xf3, 0x25,
	0x15, 0x82, 0xb2, 0x44, 0xe0, 0x33, 0x39, 0xe6, 0xd9, 0xee, 0x31, 0xce, 0x60, 0x4c, 0xb6, 0x11,
	0xd6, 0xbb, 0xb5, 0x07, 0xca, 0x9e, 0xb5, 0xe6, 0xda, 0x96, 0xe6, 0x07, 0xb0, 0xff, 0x2d, 0x58,
	0xdc, 0x2a, 0x23, 0x5a, 0x44, 0x05, 0x56, 0x0f, 0x1a, 0x25, 0xab, 0x7b, 0x51, 0xca, 0xcd, 0x6a,
	0xe9, 0xa6, 0xf5, 0x73, 0x0f, 0x74, 0x67, 0x30, 0xfe, 0xcd, 0xe5, 0xcf, 0x60, 0x24, 0x8a, 0x82,
	0x4f, 0x23, 0x21, 0x8d, 0x35, 0xfa, 0xf6, 0x83, 0x6f, 0x28, 0x57, 0xca, 0x73, 0xc5, 0xa8, 0x42,
	0xa0, 0x80, 0x7b, 0x91, 0x40, 0x53, 0xe8, 0x88, 0x90, 0xa5, 0xb1, 0x7f, 0x47, 0x17, 0x51, 0x18,
	0xf0, 0x48, 0x6e, 0x43, 0xa7, 0xff, 0xe2, 0xe1, 0x79, 0xb3, 0xa2, 0x7b, 0x54, 0x21, 0x6d, 0x89,
	0x2f, 0x13, 0x68, 0x04, 0xb0, 0xd1, 0x4b, 0x2e, 0x4e, 0xe7, 0x4f, 0xe4, 0x2e, 0xd6, 0xfd, 0x64,
	0x0b, 0x8b, 0x9e, 0x83, 0x11, 0x4b, 0xbd, 0x7c, 0x29, 0x55, 0xbe, 0x67, 0xcd, 0x91, 0x46, 0x40,
	0x25, 0x2f, 0x73, 0xc9, 0xce, 0xa1, 0x9d, 0xd7, 0x36, 0xe4, 0x8f, 0xff, 0x89, 0xbc, 0x46, 0x5a,
	0x39, 0x7c, 0xcd, 0xfd, 0x29, 0x34, 0x69, 0xe4, 0x5f, 0xd3, 0x45, 0x16, 0x73, 0x6c, 0x77, 0x75,
	0xbb, 0x49, 0x1a, 0x34, 0x3a, 0x93, 0xf1, 0xe1, 0x11, 0xc0, 0x46, 0x45, 0x64, 0x82, 0x9e, 0x8b,
	0xaf, 0xc9, 0xa6, 0xfc, 0x68, 0xbd, 0x01, 0xd8, 0x3c, 0x04, 0x19, 0x50, 0x9f, 0x4c, 0xfd, 0x8b,
	0x21, 0x39, 0x37, 0x2b, 0xa8, 0x01, 0x7b, 0x64, 0xe8, 0xb8, 0xa6, 0x86, 0x9a, 0xb0, 0x3f, 0x23,
	0xde, 0xe5, 0xd0, 0xac, 0xa2, 0x3a, 0xe8, 0xd3, 0xd9, 0xc4, 0xd4, 0xad, 0x57, 0xd0, 0x58, 0x33,
	0x78, 0x04, 0xc6, 0x64, 0xea, 0xcf, 0xbc, 0xb1, 0x3b, 0x70, 0x88, 0x6b, 0x56, 0x90, 0x09, 0xad,
	0x32, 0xf2, 0x9d, 0xf1, 0xd8, 0xd4, 0x4e, 0xeb, 0xb0, 0x2f, 0x15, 0x3f, 0xad, 0xa9, 0x1d, 0x3a,
	0x6d, 0x7c, 0xa9, 0xc9, 0xbf, 0x96, 0xb8, 0x52, 0xdf, 0xd7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0x4f, 0xf3, 0xd9, 0x08, 0x05, 0x00, 0x00,
}
