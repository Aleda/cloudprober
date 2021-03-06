// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/probes/dns/proto/config.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	github.com/google/cloudprober/probes/dns/proto/config.proto

It has these top-level messages:
	ProbeConf
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type ProbeConf struct {
	// Domain to use when making DNS queries
	ResolvedDomain *string `protobuf:"bytes,1,opt,name=resolved_domain,json=resolvedDomain,def=www.google.com." json:"resolved_domain,omitempty"`
	// Export stats after these many milliseconds
	StatsExportIntervalMsec *int32 `protobuf:"varint,2,opt,name=stats_export_interval_msec,json=statsExportIntervalMsec,def=10000" json:"stats_export_interval_msec,omitempty"`
	XXX_unrecognized        []byte `json:"-"`
}

func (m *ProbeConf) Reset()                    { *m = ProbeConf{} }
func (m *ProbeConf) String() string            { return proto1.CompactTextString(m) }
func (*ProbeConf) ProtoMessage()               {}
func (*ProbeConf) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_ProbeConf_ResolvedDomain string = "www.google.com."
const Default_ProbeConf_StatsExportIntervalMsec int32 = 10000

func (m *ProbeConf) GetResolvedDomain() string {
	if m != nil && m.ResolvedDomain != nil {
		return *m.ResolvedDomain
	}
	return Default_ProbeConf_ResolvedDomain
}

func (m *ProbeConf) GetStatsExportIntervalMsec() int32 {
	if m != nil && m.StatsExportIntervalMsec != nil {
		return *m.StatsExportIntervalMsec
	}
	return Default_ProbeConf_StatsExportIntervalMsec
}

func init() {
	proto1.RegisterType((*ProbeConf)(nil), "cloudprober.probes.dns.ProbeConf")
}

func init() {
	proto1.RegisterFile("github.com/google/cloudprober/probes/dns/proto/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x07, 0x53, 0xc5, 0xfa, 0x29, 0x79,
	0xc5, 0x20, 0x66, 0x49, 0xbe, 0x7e, 0x72, 0x7e, 0x5e, 0x5a, 0x66, 0xba, 0x1e, 0x98, 0x23, 0x24,
	0x86, 0xa4, 0x54, 0x0f, 0xa2, 0x54, 0x2f, 0x25, 0xaf, 0x58, 0xa9, 0x93, 0x91, 0x8b, 0x33, 0x00,
	0xc4, 0x75, 0xce, 0xcf, 0x4b, 0x13, 0xb2, 0xe0, 0xe2, 0x2f, 0x4a, 0x2d, 0xce, 0xcf, 0x29, 0x4b,
	0x4d, 0x89, 0x4f, 0xc9, 0xcf, 0x4d, 0xcc, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0xb4, 0xe2,
	0x2f, 0x2f, 0x2f, 0xd7, 0x83, 0x58, 0x0a, 0xb2, 0x5f, 0x2f, 0x88, 0x0f, 0xa6, 0xce, 0x05, 0xac,
	0x4c, 0xc8, 0x89, 0x4b, 0xaa, 0xb8, 0x24, 0xb1, 0xa4, 0x38, 0x3e, 0xb5, 0xa2, 0x20, 0xbf, 0xa8,
	0x24, 0x3e, 0x33, 0xaf, 0x24, 0xb5, 0xa8, 0x2c, 0x31, 0x27, 0x3e, 0xb7, 0x38, 0x35, 0x59, 0x82,
	0x49, 0x81, 0x51, 0x83, 0xd5, 0x8a, 0xd5, 0xd0, 0xc0, 0xc0, 0xc0, 0x20, 0x48, 0x1c, 0xac, 0xd0,
	0x15, 0xac, 0xce, 0x13, 0xaa, 0xcc, 0xb7, 0x38, 0x35, 0x19, 0x10, 0x00, 0x00, 0xff, 0xff, 0x54,
	0x08, 0xfa, 0x19, 0xe1, 0x00, 0x00, 0x00,
}
