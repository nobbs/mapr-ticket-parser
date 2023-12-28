// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: security.proto

package proto

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

type CredentialsMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid                       *uint32       `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	Gids                      []uint32      `protobuf:"varint,2,rep,name=gids" json:"gids,omitempty"`
	IsRoot                    *bool         `protobuf:"varint,3,opt,name=isRoot" json:"isRoot,omitempty"`
	UserName                  *string       `protobuf:"bytes,4,opt,name=userName" json:"userName,omitempty"`
	ClientIp                  *uint32       `protobuf:"varint,5,opt,name=clientIp" json:"clientIp,omitempty"`
	IsPrivilegedProcess       *bool         `protobuf:"varint,6,opt,name=isPrivilegedProcess" json:"isPrivilegedProcess,omitempty"`
	LogInSeperateAuditPath    *bool         `protobuf:"varint,7,opt,name=logInSeperateAuditPath" json:"logInSeperateAuditPath,omitempty"`
	TenantUid                 *uint32       `protobuf:"varint,8,opt,name=tenantUid" json:"tenantUid,omitempty"`
	TenantGids                []uint32      `protobuf:"varint,9,rep,name=tenantGids" json:"tenantGids,omitempty"`
	EncryptedImpersonationMsg []byte        `protobuf:"bytes,10,opt,name=encryptedImpersonationMsg" json:"encryptedImpersonationMsg,omitempty"`
	FromMastGateway           *bool         `protobuf:"varint,11,opt,name=fromMastGateway" json:"fromMastGateway,omitempty"`
	Capabilities              *Capabilities `protobuf:"bytes,12,opt,name=capabilities" json:"capabilities,omitempty"`
}

func (x *CredentialsMsg) Reset() {
	*x = CredentialsMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialsMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialsMsg) ProtoMessage() {}

func (x *CredentialsMsg) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialsMsg.ProtoReflect.Descriptor instead.
func (*CredentialsMsg) Descriptor() ([]byte, []int) {
	return file_security_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialsMsg) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *CredentialsMsg) GetGids() []uint32 {
	if x != nil {
		return x.Gids
	}
	return nil
}

func (x *CredentialsMsg) GetIsRoot() bool {
	if x != nil && x.IsRoot != nil {
		return *x.IsRoot
	}
	return false
}

func (x *CredentialsMsg) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *CredentialsMsg) GetClientIp() uint32 {
	if x != nil && x.ClientIp != nil {
		return *x.ClientIp
	}
	return 0
}

func (x *CredentialsMsg) GetIsPrivilegedProcess() bool {
	if x != nil && x.IsPrivilegedProcess != nil {
		return *x.IsPrivilegedProcess
	}
	return false
}

func (x *CredentialsMsg) GetLogInSeperateAuditPath() bool {
	if x != nil && x.LogInSeperateAuditPath != nil {
		return *x.LogInSeperateAuditPath
	}
	return false
}

func (x *CredentialsMsg) GetTenantUid() uint32 {
	if x != nil && x.TenantUid != nil {
		return *x.TenantUid
	}
	return 0
}

func (x *CredentialsMsg) GetTenantGids() []uint32 {
	if x != nil {
		return x.TenantGids
	}
	return nil
}

func (x *CredentialsMsg) GetEncryptedImpersonationMsg() []byte {
	if x != nil {
		return x.EncryptedImpersonationMsg
	}
	return nil
}

func (x *CredentialsMsg) GetFromMastGateway() bool {
	if x != nil && x.FromMastGateway != nil {
		return *x.FromMastGateway
	}
	return false
}

func (x *CredentialsMsg) GetCapabilities() *Capabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type Capabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterOpsMask *uint64 `protobuf:"varint,1,opt,name=clusterOpsMask" json:"clusterOpsMask,omitempty"`
}

func (x *Capabilities) Reset() {
	*x = Capabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capabilities) ProtoMessage() {}

func (x *Capabilities) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capabilities.ProtoReflect.Descriptor instead.
func (*Capabilities) Descriptor() ([]byte, []int) {
	return file_security_proto_rawDescGZIP(), []int{1}
}

func (x *Capabilities) GetClusterOpsMask() uint64 {
	if x != nil && x.ClusterOpsMask != nil {
		return *x.ClusterOpsMask
	}
	return 0
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_security_proto_rawDescGZIP(), []int{2}
}

func (x *Key) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type TicketAndKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncryptedTicket       []byte          `protobuf:"bytes,1,opt,name=encryptedTicket" json:"encryptedTicket,omitempty"`
	UserKey               *Key            `protobuf:"bytes,2,opt,name=userKey" json:"userKey,omitempty"`
	UserCreds             *CredentialsMsg `protobuf:"bytes,3,opt,name=userCreds" json:"userCreds,omitempty"`
	ExpiryTime            *uint64         `protobuf:"varint,4,opt,name=expiryTime" json:"expiryTime,omitempty"`
	CreationTimeSec       *uint64         `protobuf:"varint,5,opt,name=creationTimeSec" json:"creationTimeSec,omitempty"`
	MaxRenewalDurationSec *uint64         `protobuf:"varint,6,opt,name=maxRenewalDurationSec" json:"maxRenewalDurationSec,omitempty"`
	CanUserImpersonate    *bool           `protobuf:"varint,7,opt,name=canUserImpersonate" json:"canUserImpersonate,omitempty"`
	Ips                   []int32         `protobuf:"varint,8,rep,name=ips" json:"ips,omitempty"`
	ImpersonatedUids      []int32         `protobuf:"varint,9,rep,name=impersonatedUids" json:"impersonatedUids,omitempty"`
	ImpersonatedGids      []int32         `protobuf:"varint,10,rep,name=impersonatedGids" json:"impersonatedGids,omitempty"`
	IsTenant              *bool           `protobuf:"varint,11,opt,name=isTenant" json:"isTenant,omitempty"`
	IsExternal            *bool           `protobuf:"varint,12,opt,name=isExternal" json:"isExternal,omitempty"`
	LastRenewalTime       *uint64         `protobuf:"varint,13,opt,name=lastRenewalTime" json:"lastRenewalTime,omitempty"`
	CanUserGenerateTicket *bool           `protobuf:"varint,14,opt,name=canUserGenerateTicket" json:"canUserGenerateTicket,omitempty"`
	IsRemoteTempTicket    *bool           `protobuf:"varint,15,opt,name=isRemoteTempTicket" json:"isRemoteTempTicket,omitempty"`
}

func (x *TicketAndKey) Reset() {
	*x = TicketAndKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TicketAndKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TicketAndKey) ProtoMessage() {}

func (x *TicketAndKey) ProtoReflect() protoreflect.Message {
	mi := &file_security_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TicketAndKey.ProtoReflect.Descriptor instead.
func (*TicketAndKey) Descriptor() ([]byte, []int) {
	return file_security_proto_rawDescGZIP(), []int{3}
}

func (x *TicketAndKey) GetEncryptedTicket() []byte {
	if x != nil {
		return x.EncryptedTicket
	}
	return nil
}

func (x *TicketAndKey) GetUserKey() *Key {
	if x != nil {
		return x.UserKey
	}
	return nil
}

func (x *TicketAndKey) GetUserCreds() *CredentialsMsg {
	if x != nil {
		return x.UserCreds
	}
	return nil
}

func (x *TicketAndKey) GetExpiryTime() uint64 {
	if x != nil && x.ExpiryTime != nil {
		return *x.ExpiryTime
	}
	return 0
}

func (x *TicketAndKey) GetCreationTimeSec() uint64 {
	if x != nil && x.CreationTimeSec != nil {
		return *x.CreationTimeSec
	}
	return 0
}

func (x *TicketAndKey) GetMaxRenewalDurationSec() uint64 {
	if x != nil && x.MaxRenewalDurationSec != nil {
		return *x.MaxRenewalDurationSec
	}
	return 0
}

func (x *TicketAndKey) GetCanUserImpersonate() bool {
	if x != nil && x.CanUserImpersonate != nil {
		return *x.CanUserImpersonate
	}
	return false
}

func (x *TicketAndKey) GetIps() []int32 {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *TicketAndKey) GetImpersonatedUids() []int32 {
	if x != nil {
		return x.ImpersonatedUids
	}
	return nil
}

func (x *TicketAndKey) GetImpersonatedGids() []int32 {
	if x != nil {
		return x.ImpersonatedGids
	}
	return nil
}

func (x *TicketAndKey) GetIsTenant() bool {
	if x != nil && x.IsTenant != nil {
		return *x.IsTenant
	}
	return false
}

func (x *TicketAndKey) GetIsExternal() bool {
	if x != nil && x.IsExternal != nil {
		return *x.IsExternal
	}
	return false
}

func (x *TicketAndKey) GetLastRenewalTime() uint64 {
	if x != nil && x.LastRenewalTime != nil {
		return *x.LastRenewalTime
	}
	return 0
}

func (x *TicketAndKey) GetCanUserGenerateTicket() bool {
	if x != nil && x.CanUserGenerateTicket != nil {
		return *x.CanUserGenerateTicket
	}
	return false
}

func (x *TicketAndKey) GetIsRemoteTempTicket() bool {
	if x != nil && x.IsRemoteTempTicket != nil {
		return *x.IsRemoteTempTicket
	}
	return false
}

var File_security_proto protoreflect.FileDescriptor

var file_security_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x6d, 0x61, 0x70, 0x72, 0x2e, 0x66, 0x73, 0x22, 0xd1, 0x03, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x67, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x04, 0x67, 0x69,
	0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x12, 0x30, 0x0a, 0x13, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67,
	0x65, 0x64, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x13, 0x69, 0x73, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x16, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x53, 0x65, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x6c, 0x6f, 0x67, 0x49, 0x6e, 0x53, 0x65, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x41, 0x75, 0x64, 0x69, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x55, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x47, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x47, 0x69, 0x64, 0x73, 0x12, 0x3c, 0x0a, 0x19, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x19, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d,
	0x4d, 0x61, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x4d, 0x61, 0x73, 0x74, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x39, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x70, 0x72, 0x2e,
	0x66, 0x73, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x36, 0x0a,
	0x0c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x73, 0x4d, 0x61, 0x73, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70,
	0x73, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x17, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0xfd,
	0x04, 0x0a, 0x0c, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x6e, 0x64, 0x4b, 0x65, 0x79, 0x12,
	0x28, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x61, 0x70,
	0x72, 0x2e, 0x66, 0x73, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x4b, 0x65,
	0x79, 0x12, 0x35, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x70, 0x72, 0x2e, 0x66, 0x73, 0x2e, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x4d, 0x73, 0x67, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x63, 0x12, 0x34, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x61, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x63, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x70, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x6d,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x64, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x10, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74,
	0x65, 0x64, 0x55, 0x69, 0x64, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x47, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x05,
	0x52, 0x10, 0x69, 0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x47, 0x69,
	0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x28,
	0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x6e,
	0x65, 0x77, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x63, 0x61, 0x6e, 0x55,
	0x73, 0x65, 0x72, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x63, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x72,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2e,
	0x0a, 0x12, 0x69, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x52, 0x65,
	0x6d, 0x6f, 0x74, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x42, 0x37,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x70, 0x72, 0x2e, 0x66, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x48, 0x03, 0x5a, 0x20, 0x65, 0x7a, 0x6d, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x68,
	0x70, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x66, 0x61, 0x62, 0x2f, 0x66,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_security_proto_rawDescOnce sync.Once
	file_security_proto_rawDescData = file_security_proto_rawDesc
)

func file_security_proto_rawDescGZIP() []byte {
	file_security_proto_rawDescOnce.Do(func() {
		file_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_proto_rawDescData)
	})
	return file_security_proto_rawDescData
}

var file_security_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_security_proto_goTypes = []interface{}{
	(*CredentialsMsg)(nil), // 0: mapr.fs.CredentialsMsg
	(*Capabilities)(nil),   // 1: mapr.fs.Capabilities
	(*Key)(nil),            // 2: mapr.fs.Key
	(*TicketAndKey)(nil),   // 3: mapr.fs.TicketAndKey
}
var file_security_proto_depIdxs = []int32{
	1, // 0: mapr.fs.CredentialsMsg.capabilities:type_name -> mapr.fs.Capabilities
	2, // 1: mapr.fs.TicketAndKey.userKey:type_name -> mapr.fs.Key
	0, // 2: mapr.fs.TicketAndKey.userCreds:type_name -> mapr.fs.CredentialsMsg
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_security_proto_init() }
func file_security_proto_init() {
	if File_security_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialsMsg); i {
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
		file_security_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capabilities); i {
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
		file_security_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_security_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TicketAndKey); i {
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
			RawDescriptor: file_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_proto_goTypes,
		DependencyIndexes: file_security_proto_depIdxs,
		MessageInfos:      file_security_proto_msgTypes,
	}.Build()
	File_security_proto = out.File
	file_security_proto_rawDesc = nil
	file_security_proto_goTypes = nil
	file_security_proto_depIdxs = nil
}