// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: cluster.proto

package cluster

import (
	actor "github.com/anthdm/hollywood/actor"
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

type CID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID    *actor.PID `protobuf:"bytes,1,opt,name=PID,proto3" json:"PID,omitempty"`
	ID     string     `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Kind   string     `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Region string     `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *CID) Reset() {
	*x = CID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CID) ProtoMessage() {}

func (x *CID) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CID.ProtoReflect.Descriptor instead.
func (*CID) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *CID) GetPID() *actor.PID {
	if x != nil {
		return x.PID
	}
	return nil
}

func (x *CID) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *CID) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *CID) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Host   string   `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Region string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Kinds  []string `protobuf:"bytes,4,rep,name=kinds,proto3" json:"kinds,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Member) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Member) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Member) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

type Members struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *Members) Reset() {
	*x = Members{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Members) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Members) ProtoMessage() {}

func (x *Members) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Members.ProtoReflect.Descriptor instead.
func (*Members) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *Members) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type MembersJoin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *MembersJoin) Reset() {
	*x = MembersJoin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembersJoin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembersJoin) ProtoMessage() {}

func (x *MembersJoin) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembersJoin.ProtoReflect.Descriptor instead.
func (*MembersJoin) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *MembersJoin) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type MembersLeave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *MembersLeave) Reset() {
	*x = MembersLeave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembersLeave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembersLeave) ProtoMessage() {}

func (x *MembersLeave) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembersLeave.ProtoReflect.Descriptor instead.
func (*MembersLeave) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *MembersLeave) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type Topology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash    uint64    `protobuf:"varint,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Members []*Member `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	Left    []*Member `protobuf:"bytes,3,rep,name=left,proto3" json:"left,omitempty"`
	Joined  []*Member `protobuf:"bytes,4,rep,name=joined,proto3" json:"joined,omitempty"`
	Blocked []*Member `protobuf:"bytes,5,rep,name=blocked,proto3" json:"blocked,omitempty"`
}

func (x *Topology) Reset() {
	*x = Topology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topology) ProtoMessage() {}

func (x *Topology) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topology.ProtoReflect.Descriptor instead.
func (*Topology) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{5}
}

func (x *Topology) GetHash() uint64 {
	if x != nil {
		return x.Hash
	}
	return 0
}

func (x *Topology) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *Topology) GetLeft() []*Member {
	if x != nil {
		return x.Left
	}
	return nil
}

func (x *Topology) GetJoined() []*Member {
	if x != nil {
		return x.Joined
	}
	return nil
}

func (x *Topology) GetBlocked() []*Member {
	if x != nil {
		return x.Blocked
	}
	return nil
}

type ActorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID *actor.PID `protobuf:"bytes,2,opt,name=PID,proto3" json:"PID,omitempty"`
}

func (x *ActorInfo) Reset() {
	*x = ActorInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorInfo) ProtoMessage() {}

func (x *ActorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorInfo.ProtoReflect.Descriptor instead.
func (*ActorInfo) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{6}
}

func (x *ActorInfo) GetPID() *actor.PID {
	if x != nil {
		return x.PID
	}
	return nil
}

type ActorTopology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Actors []*ActorInfo `protobuf:"bytes,1,rep,name=actors,proto3" json:"actors,omitempty"`
}

func (x *ActorTopology) Reset() {
	*x = ActorTopology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActorTopology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActorTopology) ProtoMessage() {}

func (x *ActorTopology) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActorTopology.ProtoReflect.Descriptor instead.
func (*ActorTopology) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{7}
}

func (x *ActorTopology) GetActors() []*ActorInfo {
	if x != nil {
		return x.Actors
	}
	return nil
}

type Activation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID *actor.PID `protobuf:"bytes,1,opt,name=PID,proto3" json:"PID,omitempty"`
}

func (x *Activation) Reset() {
	*x = Activation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activation) ProtoMessage() {}

func (x *Activation) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activation.ProtoReflect.Descriptor instead.
func (*Activation) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{8}
}

func (x *Activation) GetPID() *actor.PID {
	if x != nil {
		return x.PID
	}
	return nil
}

type Deactivation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID *actor.PID `protobuf:"bytes,1,opt,name=PID,proto3" json:"PID,omitempty"`
}

func (x *Deactivation) Reset() {
	*x = Deactivation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deactivation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deactivation) ProtoMessage() {}

func (x *Deactivation) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deactivation.ProtoReflect.Descriptor instead.
func (*Deactivation) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{9}
}

func (x *Deactivation) GetPID() *actor.PID {
	if x != nil {
		return x.PID
	}
	return nil
}

type ActivationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind         string `protobuf:"bytes,1,opt,name=Kind,proto3" json:"Kind,omitempty"`
	ID           string `protobuf:"bytes,2,opt,name=ID,proto3" json:"ID,omitempty"`
	Region       string `protobuf:"bytes,3,opt,name=Region,proto3" json:"Region,omitempty"`
	TopologyHash uint64 `protobuf:"varint,4,opt,name=topologyHash,proto3" json:"topologyHash,omitempty"`
}

func (x *ActivationRequest) Reset() {
	*x = ActivationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationRequest) ProtoMessage() {}

func (x *ActivationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationRequest.ProtoReflect.Descriptor instead.
func (*ActivationRequest) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{10}
}

func (x *ActivationRequest) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ActivationRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ActivationRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ActivationRequest) GetTopologyHash() uint64 {
	if x != nil {
		return x.TopologyHash
	}
	return 0
}

type ActivationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PID          *actor.PID `protobuf:"bytes,1,opt,name=PID,proto3" json:"PID,omitempty"`
	Success      bool       `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	TopologyHash uint64     `protobuf:"varint,3,opt,name=topologyHash,proto3" json:"topologyHash,omitempty"`
}

func (x *ActivationResponse) Reset() {
	*x = ActivationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cluster_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivationResponse) ProtoMessage() {}

func (x *ActivationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cluster_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivationResponse.ProtoReflect.Descriptor instead.
func (*ActivationResponse) Descriptor() ([]byte, []int) {
	return file_cluster_proto_rawDescGZIP(), []int{11}
}

func (x *ActivationResponse) GetPID() *actor.PID {
	if x != nil {
		return x.PID
	}
	return nil
}

func (x *ActivationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ActivationResponse) GetTopologyHash() uint64 {
	if x != nil {
		return x.TopologyHash
	}
	return 0
}

var File_cluster_proto protoreflect.FileDescriptor

var file_cluster_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x03, 0x43, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x03,
	0x50, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x03, 0x50, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x5a, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x6b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6b, 0x69, 0x6e,
	0x64, 0x73, 0x22, 0x34, 0x0a, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x29, 0x0a,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x38, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x22, 0x39, 0x0a, 0x0c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0xc2, 0x01,
	0x0a, 0x08, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x29,
	0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x65, 0x66,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x27,
	0x0a, 0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x06, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x22, 0x29, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1c, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x03, 0x50, 0x49, 0x44, 0x22, 0x3b, 0x0a,
	0x0d, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x2a,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x0a, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49,
	0x44, 0x52, 0x03, 0x50, 0x49, 0x44, 0x22, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52,
	0x03, 0x50, 0x49, 0x44, 0x22, 0x73, 0x0a, 0x11, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x48, 0x61, 0x73, 0x68, 0x22, 0x70, 0x0a, 0x12, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x03, 0x50, 0x49, 0x44, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x74,
	0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x48, 0x61, 0x73, 0x68, 0x42, 0x25, 0x5a, 0x23, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6e, 0x74, 0x68, 0x64, 0x6d,
	0x2f, 0x68, 0x6f, 0x6c, 0x6c, 0x79, 0x77, 0x6f, 0x6f, 0x64, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cluster_proto_rawDescOnce sync.Once
	file_cluster_proto_rawDescData = file_cluster_proto_rawDesc
)

func file_cluster_proto_rawDescGZIP() []byte {
	file_cluster_proto_rawDescOnce.Do(func() {
		file_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_cluster_proto_rawDescData)
	})
	return file_cluster_proto_rawDescData
}

var file_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_cluster_proto_goTypes = []interface{}{
	(*CID)(nil),                // 0: cluster.CID
	(*Member)(nil),             // 1: cluster.Member
	(*Members)(nil),            // 2: cluster.Members
	(*MembersJoin)(nil),        // 3: cluster.MembersJoin
	(*MembersLeave)(nil),       // 4: cluster.MembersLeave
	(*Topology)(nil),           // 5: cluster.Topology
	(*ActorInfo)(nil),          // 6: cluster.ActorInfo
	(*ActorTopology)(nil),      // 7: cluster.ActorTopology
	(*Activation)(nil),         // 8: cluster.Activation
	(*Deactivation)(nil),       // 9: cluster.Deactivation
	(*ActivationRequest)(nil),  // 10: cluster.ActivationRequest
	(*ActivationResponse)(nil), // 11: cluster.ActivationResponse
	(*actor.PID)(nil),          // 12: actor.PID
}
var file_cluster_proto_depIdxs = []int32{
	12, // 0: cluster.CID.PID:type_name -> actor.PID
	1,  // 1: cluster.Members.members:type_name -> cluster.Member
	1,  // 2: cluster.MembersJoin.members:type_name -> cluster.Member
	1,  // 3: cluster.MembersLeave.members:type_name -> cluster.Member
	1,  // 4: cluster.Topology.members:type_name -> cluster.Member
	1,  // 5: cluster.Topology.left:type_name -> cluster.Member
	1,  // 6: cluster.Topology.joined:type_name -> cluster.Member
	1,  // 7: cluster.Topology.blocked:type_name -> cluster.Member
	12, // 8: cluster.ActorInfo.PID:type_name -> actor.PID
	6,  // 9: cluster.ActorTopology.actors:type_name -> cluster.ActorInfo
	12, // 10: cluster.Activation.PID:type_name -> actor.PID
	12, // 11: cluster.Deactivation.PID:type_name -> actor.PID
	12, // 12: cluster.ActivationResponse.PID:type_name -> actor.PID
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cluster_proto_init() }
func file_cluster_proto_init() {
	if File_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CID); i {
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
		file_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Members); i {
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
		file_cluster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembersJoin); i {
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
		file_cluster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MembersLeave); i {
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
		file_cluster_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topology); i {
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
		file_cluster_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorInfo); i {
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
		file_cluster_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActorTopology); i {
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
		file_cluster_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activation); i {
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
		file_cluster_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deactivation); i {
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
		file_cluster_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationRequest); i {
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
		file_cluster_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivationResponse); i {
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
			RawDescriptor: file_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cluster_proto_goTypes,
		DependencyIndexes: file_cluster_proto_depIdxs,
		MessageInfos:      file_cluster_proto_msgTypes,
	}.Build()
	File_cluster_proto = out.File
	file_cluster_proto_rawDesc = nil
	file_cluster_proto_goTypes = nil
	file_cluster_proto_depIdxs = nil
}
