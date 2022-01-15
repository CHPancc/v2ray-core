// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: proxy/shadowsocks/config.proto

package shadowsocks

import (
	net "github.com/v2fly/v2ray-core/v5/common/net"
	packetaddr "github.com/v2fly/v2ray-core/v5/common/net/packetaddr"
	protocol "github.com/v2fly/v2ray-core/v5/common/protocol"
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

type CipherType int32

const (
	CipherType_UNKNOWN                 CipherType = 0
	CipherType_AES_128_GCM             CipherType = 1
	CipherType_AES_192_GCM             CipherType = 2
	CipherType_AES_256_GCM             CipherType = 3
	CipherType_CHACHA20_IETF_POLY1305  CipherType = 4
	CipherType_XCHACHA20_IETF_POLY1305 CipherType = 5
	CipherType_NONE                    CipherType = 6
	CipherType_AES_128_CTR             CipherType = 7
	CipherType_AES_192_CTR             CipherType = 8
	CipherType_AES_256_CTR             CipherType = 9
	CipherType_AES_128_CFB             CipherType = 10
	CipherType_AES_192_CFB             CipherType = 11
	CipherType_AES_256_CFB             CipherType = 12
	CipherType_AES_128_CFB8            CipherType = 13
	CipherType_AES_192_CFB8            CipherType = 14
	CipherType_AES_256_CFB8            CipherType = 15
	CipherType_AES_128_OFB             CipherType = 16
	CipherType_AES_192_OFB             CipherType = 17
	CipherType_AES_256_OFB             CipherType = 18
	CipherType_RC4                     CipherType = 19
	CipherType_RC4_MD5                 CipherType = 20
	CipherType_BF_CFB                  CipherType = 21
	CipherType_CAST5_CFB               CipherType = 22
	CipherType_DES_CFB                 CipherType = 23
	CipherType_IDEA_CFB                CipherType = 24
	CipherType_RC2_CFB                 CipherType = 25
	CipherType_SEED_CFB                CipherType = 26
	CipherType_CAMELLIA_128_CFB        CipherType = 27
	CipherType_CAMELLIA_192_CFB        CipherType = 28
	CipherType_CAMELLIA_256_CFB        CipherType = 29
	CipherType_CAMELLIA_128_CFB8       CipherType = 30
	CipherType_CAMELLIA_192_CFB8       CipherType = 31
	CipherType_CAMELLIA_256_CFB8       CipherType = 32
	CipherType_SALSA20                 CipherType = 33
	CipherType_CHACHA20                CipherType = 34
	CipherType_CHACHA20_IETF           CipherType = 35
	CipherType_XCHACHA20               CipherType = 36
)

// Enum value maps for CipherType.
var (
	CipherType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "AES_128_GCM",
		2:  "AES_192_GCM",
		3:  "AES_256_GCM",
		4:  "CHACHA20_IETF_POLY1305",
		5:  "XCHACHA20_IETF_POLY1305",
		6:  "NONE",
		7:  "AES_128_CTR",
		8:  "AES_192_CTR",
		9:  "AES_256_CTR",
		10: "AES_128_CFB",
		11: "AES_192_CFB",
		12: "AES_256_CFB",
		13: "AES_128_CFB8",
		14: "AES_192_CFB8",
		15: "AES_256_CFB8",
		16: "AES_128_OFB",
		17: "AES_192_OFB",
		18: "AES_256_OFB",
		19: "RC4",
		20: "RC4_MD5",
		21: "BF_CFB",
		22: "CAST5_CFB",
		23: "DES_CFB",
		24: "IDEA_CFB",
		25: "RC2_CFB",
		26: "SEED_CFB",
		27: "CAMELLIA_128_CFB",
		28: "CAMELLIA_192_CFB",
		29: "CAMELLIA_256_CFB",
		30: "CAMELLIA_128_CFB8",
		31: "CAMELLIA_192_CFB8",
		32: "CAMELLIA_256_CFB8",
		33: "SALSA20",
		34: "CHACHA20",
		35: "CHACHA20_IETF",
		36: "XCHACHA20",
	}
	CipherType_value = map[string]int32{
		"UNKNOWN":                 0,
		"AES_128_GCM":             1,
		"AES_192_GCM":             2,
		"AES_256_GCM":             3,
		"CHACHA20_IETF_POLY1305":  4,
		"XCHACHA20_IETF_POLY1305": 5,
		"NONE":                    6,
		"AES_128_CTR":             7,
		"AES_192_CTR":             8,
		"AES_256_CTR":             9,
		"AES_128_CFB":             10,
		"AES_192_CFB":             11,
		"AES_256_CFB":             12,
		"AES_128_CFB8":            13,
		"AES_192_CFB8":            14,
		"AES_256_CFB8":            15,
		"AES_128_OFB":             16,
		"AES_192_OFB":             17,
		"AES_256_OFB":             18,
		"RC4":                     19,
		"RC4_MD5":                 20,
		"BF_CFB":                  21,
		"CAST5_CFB":               22,
		"DES_CFB":                 23,
		"IDEA_CFB":                24,
		"RC2_CFB":                 25,
		"SEED_CFB":                26,
		"CAMELLIA_128_CFB":        27,
		"CAMELLIA_192_CFB":        28,
		"CAMELLIA_256_CFB":        29,
		"CAMELLIA_128_CFB8":       30,
		"CAMELLIA_192_CFB8":       31,
		"CAMELLIA_256_CFB8":       32,
		"SALSA20":                 33,
		"CHACHA20":                34,
		"CHACHA20_IETF":           35,
		"XCHACHA20":               36,
	}
)

func (x CipherType) Enum() *CipherType {
	p := new(CipherType)
	*p = x
	return p
}

func (x CipherType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CipherType) Descriptor() protoreflect.EnumDescriptor {
	return file_proxy_shadowsocks_config_proto_enumTypes[0].Descriptor()
}

func (CipherType) Type() protoreflect.EnumType {
	return &file_proxy_shadowsocks_config_proto_enumTypes[0]
}

func (x CipherType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CipherType.Descriptor instead.
func (CipherType) EnumDescriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password                       string     `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	CipherType                     CipherType `protobuf:"varint,2,opt,name=cipher_type,json=cipherType,proto3,enum=v2ray.core.proxy.shadowsocks.CipherType" json:"cipher_type,omitempty"`
	IvCheck                        bool       `protobuf:"varint,3,opt,name=iv_check,json=ivCheck,proto3" json:"iv_check,omitempty"`
	ExperimentReducedIvHeadEntropy bool       `protobuf:"varint,90001,opt,name=experiment_reduced_iv_head_entropy,json=experimentReducedIvHeadEntropy,proto3" json:"experiment_reduced_iv_head_entropy,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Account) GetCipherType() CipherType {
	if x != nil {
		return x.CipherType
	}
	return CipherType_UNKNOWN
}

func (x *Account) GetIvCheck() bool {
	if x != nil {
		return x.IvCheck
	}
	return false
}

func (x *Account) GetExperimentReducedIvHeadEntropy() bool {
	if x != nil {
		return x.ExperimentReducedIvHeadEntropy
	}
	return false
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UdpEnabled specified whether or not to enable UDP for Shadowsocks.
	// Deprecated. Use 'network' field.
	//
	// Deprecated: Do not use.
	UdpEnabled     bool                      `protobuf:"varint,1,opt,name=udp_enabled,json=udpEnabled,proto3" json:"udp_enabled,omitempty"`
	User           *protocol.User            `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Network        []net.Network             `protobuf:"varint,3,rep,packed,name=network,proto3,enum=v2ray.core.common.net.Network" json:"network,omitempty"`
	PacketEncoding packetaddr.PacketAddrType `protobuf:"varint,4,opt,name=packet_encoding,json=packetEncoding,proto3,enum=v2ray.core.net.packetaddr.PacketAddrType" json:"packet_encoding,omitempty"`
	Plugin         string                    `protobuf:"bytes,5,opt,name=plugin,proto3" json:"plugin,omitempty"`
	PluginOpts     string                    `protobuf:"bytes,6,opt,name=plugin_opts,json=pluginOpts,proto3" json:"plugin_opts,omitempty"`
	PluginArgs     []string                  `protobuf:"bytes,7,rep,name=plugin_args,json=pluginArgs,proto3" json:"plugin_args,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{1}
}

// Deprecated: Do not use.
func (x *ServerConfig) GetUdpEnabled() bool {
	if x != nil {
		return x.UdpEnabled
	}
	return false
}

func (x *ServerConfig) GetUser() *protocol.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *ServerConfig) GetNetwork() []net.Network {
	if x != nil {
		return x.Network
	}
	return nil
}

func (x *ServerConfig) GetPacketEncoding() packetaddr.PacketAddrType {
	if x != nil {
		return x.PacketEncoding
	}
	return packetaddr.PacketAddrType(0)
}

func (x *ServerConfig) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *ServerConfig) GetPluginOpts() string {
	if x != nil {
		return x.PluginOpts
	}
	return ""
}

func (x *ServerConfig) GetPluginArgs() []string {
	if x != nil {
		return x.PluginArgs
	}
	return nil
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server                         []*protocol.ServerEndpoint `protobuf:"bytes,1,rep,name=server,proto3" json:"server,omitempty"`
	Plugin                         string                     `protobuf:"bytes,2,opt,name=plugin,proto3" json:"plugin,omitempty"`
	PluginOpts                     string                     `protobuf:"bytes,3,opt,name=plugin_opts,json=pluginOpts,proto3" json:"plugin_opts,omitempty"`
	PluginArgs                     []string                   `protobuf:"bytes,4,rep,name=plugin_args,json=pluginArgs,proto3" json:"plugin_args,omitempty"`
	ExperimentReducedIvHeadEntropy bool                       `protobuf:"varint,90001,opt,name=experiment_reduced_iv_head_entropy,json=experimentReducedIvHeadEntropy,proto3" json:"experiment_reduced_iv_head_entropy,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_config_proto_rawDescGZIP(), []int{2}
}

func (x *ClientConfig) GetServer() []*protocol.ServerEndpoint {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *ClientConfig) GetPlugin() string {
	if x != nil {
		return x.Plugin
	}
	return ""
}

func (x *ClientConfig) GetPluginOpts() string {
	if x != nil {
		return x.PluginOpts
	}
	return ""
}

func (x *ClientConfig) GetPluginArgs() []string {
	if x != nil {
		return x.PluginArgs
	}
	return nil
}

func (x *ClientConfig) GetExperimentReducedIvHeadEntropy() bool {
	if x != nil {
		return x.ExperimentReducedIvHeadEntropy
	}
	return false
}

var File_proxy_shadowsocks_config_proto protoreflect.FileDescriptor

var file_proxy_shadowsocks_config_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f,
	0x63, 0x6b, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x1a, 0x18,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6e, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x61, 0x64, 0x64, 0x72, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x07,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x49, 0x0a, 0x0b, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64,
	0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x76, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x76, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x4c, 0x0a, 0x22, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x5f,
	0x69, 0x76, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x18,
	0x91, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x49, 0x76, 0x48, 0x65, 0x61, 0x64,
	0x45, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x22, 0xd1, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x23, 0x0a, 0x0b, 0x75, 0x64, 0x70, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18,
	0x01, 0x52, 0x0a, 0x75, 0x64, 0x70, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x34, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x52, 0x0a,
	0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x61, 0x64,
	0x64, 0x72, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x22, 0xfa, 0x01, 0x0a, 0x0c,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x42, 0x0a, 0x06,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4f, 0x70, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x4c, 0x0a, 0x22, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64,
	0x5f, 0x69, 0x76, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79,
	0x18, 0x91, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x49, 0x76, 0x48, 0x65, 0x61,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x2a, 0x89, 0x05, 0x0a, 0x0a, 0x43, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x32, 0x38, 0x5f,
	0x47, 0x43, 0x4d, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x39, 0x32,
	0x5f, 0x47, 0x43, 0x4d, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35,
	0x36, 0x5f, 0x47, 0x43, 0x4d, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x48, 0x41, 0x43, 0x48,
	0x41, 0x32, 0x30, 0x5f, 0x49, 0x45, 0x54, 0x46, 0x5f, 0x50, 0x4f, 0x4c, 0x59, 0x31, 0x33, 0x30,
	0x35, 0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x58, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30,
	0x5f, 0x49, 0x45, 0x54, 0x46, 0x5f, 0x50, 0x4f, 0x4c, 0x59, 0x31, 0x33, 0x30, 0x35, 0x10, 0x05,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45,
	0x53, 0x5f, 0x31, 0x32, 0x38, 0x5f, 0x43, 0x54, 0x52, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x41,
	0x45, 0x53, 0x5f, 0x31, 0x39, 0x32, 0x5f, 0x43, 0x54, 0x52, 0x10, 0x08, 0x12, 0x0f, 0x0a, 0x0b,
	0x41, 0x45, 0x53, 0x5f, 0x32, 0x35, 0x36, 0x5f, 0x43, 0x54, 0x52, 0x10, 0x09, 0x12, 0x0f, 0x0a,
	0x0b, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x32, 0x38, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x0a, 0x12, 0x0f,
	0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x39, 0x32, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x0b, 0x12,
	0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35, 0x36, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x0c,
	0x12, 0x10, 0x0a, 0x0c, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x32, 0x38, 0x5f, 0x43, 0x46, 0x42, 0x38,
	0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x39, 0x32, 0x5f, 0x43, 0x46,
	0x42, 0x38, 0x10, 0x0e, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35, 0x36, 0x5f,
	0x43, 0x46, 0x42, 0x38, 0x10, 0x0f, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x31, 0x32,
	0x38, 0x5f, 0x4f, 0x46, 0x42, 0x10, 0x10, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x31,
	0x39, 0x32, 0x5f, 0x4f, 0x46, 0x42, 0x10, 0x11, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f,
	0x32, 0x35, 0x36, 0x5f, 0x4f, 0x46, 0x42, 0x10, 0x12, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x43, 0x34,
	0x10, 0x13, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x43, 0x34, 0x5f, 0x4d, 0x44, 0x35, 0x10, 0x14, 0x12,
	0x0a, 0x0a, 0x06, 0x42, 0x46, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x15, 0x12, 0x0d, 0x0a, 0x09, 0x43,
	0x41, 0x53, 0x54, 0x35, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x16, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45,
	0x53, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x17, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x44, 0x45, 0x41, 0x5f,
	0x43, 0x46, 0x42, 0x10, 0x18, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x43, 0x32, 0x5f, 0x43, 0x46, 0x42,
	0x10, 0x19, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x45, 0x44, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x1a,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x45, 0x4c, 0x4c, 0x49, 0x41, 0x5f, 0x31, 0x32, 0x38,
	0x5f, 0x43, 0x46, 0x42, 0x10, 0x1b, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4d, 0x45, 0x4c, 0x4c,
	0x49, 0x41, 0x5f, 0x31, 0x39, 0x32, 0x5f, 0x43, 0x46, 0x42, 0x10, 0x1c, 0x12, 0x14, 0x0a, 0x10,
	0x43, 0x41, 0x4d, 0x45, 0x4c, 0x4c, 0x49, 0x41, 0x5f, 0x32, 0x35, 0x36, 0x5f, 0x43, 0x46, 0x42,
	0x10, 0x1d, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x4d, 0x45, 0x4c, 0x4c, 0x49, 0x41, 0x5f, 0x31,
	0x32, 0x38, 0x5f, 0x43, 0x46, 0x42, 0x38, 0x10, 0x1e, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x4d,
	0x45, 0x4c, 0x4c, 0x49, 0x41, 0x5f, 0x31, 0x39, 0x32, 0x5f, 0x43, 0x46, 0x42, 0x38, 0x10, 0x1f,
	0x12, 0x15, 0x0a, 0x11, 0x43, 0x41, 0x4d, 0x45, 0x4c, 0x4c, 0x49, 0x41, 0x5f, 0x32, 0x35, 0x36,
	0x5f, 0x43, 0x46, 0x42, 0x38, 0x10, 0x20, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x41, 0x4c, 0x53, 0x41,
	0x32, 0x30, 0x10, 0x21, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30,
	0x10, 0x22, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41, 0x32, 0x30, 0x5f, 0x49,
	0x45, 0x54, 0x46, 0x10, 0x23, 0x12, 0x0d, 0x0a, 0x09, 0x58, 0x43, 0x48, 0x41, 0x43, 0x48, 0x41,
	0x32, 0x30, 0x10, 0x24, 0x42, 0x75, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61,
	0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0xaa, 0x02, 0x1c, 0x56,
	0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proxy_shadowsocks_config_proto_rawDescOnce sync.Once
	file_proxy_shadowsocks_config_proto_rawDescData = file_proxy_shadowsocks_config_proto_rawDesc
)

func file_proxy_shadowsocks_config_proto_rawDescGZIP() []byte {
	file_proxy_shadowsocks_config_proto_rawDescOnce.Do(func() {
		file_proxy_shadowsocks_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_shadowsocks_config_proto_rawDescData)
	})
	return file_proxy_shadowsocks_config_proto_rawDescData
}

var file_proxy_shadowsocks_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proxy_shadowsocks_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxy_shadowsocks_config_proto_goTypes = []interface{}{
	(CipherType)(0),                 // 0: v2ray.core.proxy.shadowsocks.CipherType
	(*Account)(nil),                 // 1: v2ray.core.proxy.shadowsocks.Account
	(*ServerConfig)(nil),            // 2: v2ray.core.proxy.shadowsocks.ServerConfig
	(*ClientConfig)(nil),            // 3: v2ray.core.proxy.shadowsocks.ClientConfig
	(*protocol.User)(nil),           // 4: v2ray.core.common.protocol.User
	(net.Network)(0),                // 5: v2ray.core.common.net.Network
	(packetaddr.PacketAddrType)(0),  // 6: v2ray.core.net.packetaddr.PacketAddrType
	(*protocol.ServerEndpoint)(nil), // 7: v2ray.core.common.protocol.ServerEndpoint
}
var file_proxy_shadowsocks_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.proxy.shadowsocks.Account.cipher_type:type_name -> v2ray.core.proxy.shadowsocks.CipherType
	4, // 1: v2ray.core.proxy.shadowsocks.ServerConfig.user:type_name -> v2ray.core.common.protocol.User
	5, // 2: v2ray.core.proxy.shadowsocks.ServerConfig.network:type_name -> v2ray.core.common.net.Network
	6, // 3: v2ray.core.proxy.shadowsocks.ServerConfig.packet_encoding:type_name -> v2ray.core.net.packetaddr.PacketAddrType
	7, // 4: v2ray.core.proxy.shadowsocks.ClientConfig.server:type_name -> v2ray.core.common.protocol.ServerEndpoint
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proxy_shadowsocks_config_proto_init() }
func file_proxy_shadowsocks_config_proto_init() {
	if File_proxy_shadowsocks_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_shadowsocks_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_proxy_shadowsocks_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
		file_proxy_shadowsocks_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
			RawDescriptor: file_proxy_shadowsocks_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_shadowsocks_config_proto_goTypes,
		DependencyIndexes: file_proxy_shadowsocks_config_proto_depIdxs,
		EnumInfos:         file_proxy_shadowsocks_config_proto_enumTypes,
		MessageInfos:      file_proxy_shadowsocks_config_proto_msgTypes,
	}.Build()
	File_proxy_shadowsocks_config_proto = out.File
	file_proxy_shadowsocks_config_proto_rawDesc = nil
	file_proxy_shadowsocks_config_proto_goTypes = nil
	file_proxy_shadowsocks_config_proto_depIdxs = nil
}
