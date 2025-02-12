// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/kubeflow/tensorflow.proto

package plugins

import (
	fmt "fmt"
	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Proto for plugin that enables distributed training using https://github.com/kubeflow/tf-operator
type DistributedTensorflowTrainingTask struct {
	// Worker replicas spec
	WorkerReplicas *DistributedTensorflowTrainingReplicaSpec `protobuf:"bytes,1,opt,name=worker_replicas,json=workerReplicas,proto3" json:"worker_replicas,omitempty"`
	// Parameter server replicas spec
	PsReplicas *DistributedTensorflowTrainingReplicaSpec `protobuf:"bytes,2,opt,name=ps_replicas,json=psReplicas,proto3" json:"ps_replicas,omitempty"`
	// Chief replicas spec
	ChiefReplicas *DistributedTensorflowTrainingReplicaSpec `protobuf:"bytes,3,opt,name=chief_replicas,json=chiefReplicas,proto3" json:"chief_replicas,omitempty"`
	// RunPolicy encapsulates various runtime policies of the distributed training
	// job, for example how to clean up resources and how long the job can stay
	// active.
	RunPolicy *RunPolicy `protobuf:"bytes,4,opt,name=run_policy,json=runPolicy,proto3" json:"run_policy,omitempty"`
	// Evaluator replicas spec
	EvaluatorReplicas    *DistributedTensorflowTrainingReplicaSpec `protobuf:"bytes,5,opt,name=evaluator_replicas,json=evaluatorReplicas,proto3" json:"evaluator_replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *DistributedTensorflowTrainingTask) Reset()         { *m = DistributedTensorflowTrainingTask{} }
func (m *DistributedTensorflowTrainingTask) String() string { return proto.CompactTextString(m) }
func (*DistributedTensorflowTrainingTask) ProtoMessage()    {}
func (*DistributedTensorflowTrainingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_93de2bd764ddf01a, []int{0}
}

func (m *DistributedTensorflowTrainingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Unmarshal(m, b)
}
func (m *DistributedTensorflowTrainingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Marshal(b, m, deterministic)
}
func (m *DistributedTensorflowTrainingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedTensorflowTrainingTask.Merge(m, src)
}
func (m *DistributedTensorflowTrainingTask) XXX_Size() int {
	return xxx_messageInfo_DistributedTensorflowTrainingTask.Size(m)
}
func (m *DistributedTensorflowTrainingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedTensorflowTrainingTask.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedTensorflowTrainingTask proto.InternalMessageInfo

func (m *DistributedTensorflowTrainingTask) GetWorkerReplicas() *DistributedTensorflowTrainingReplicaSpec {
	if m != nil {
		return m.WorkerReplicas
	}
	return nil
}

func (m *DistributedTensorflowTrainingTask) GetPsReplicas() *DistributedTensorflowTrainingReplicaSpec {
	if m != nil {
		return m.PsReplicas
	}
	return nil
}

func (m *DistributedTensorflowTrainingTask) GetChiefReplicas() *DistributedTensorflowTrainingReplicaSpec {
	if m != nil {
		return m.ChiefReplicas
	}
	return nil
}

func (m *DistributedTensorflowTrainingTask) GetRunPolicy() *RunPolicy {
	if m != nil {
		return m.RunPolicy
	}
	return nil
}

func (m *DistributedTensorflowTrainingTask) GetEvaluatorReplicas() *DistributedTensorflowTrainingReplicaSpec {
	if m != nil {
		return m.EvaluatorReplicas
	}
	return nil
}

type DistributedTensorflowTrainingReplicaSpec struct {
	// Number of replicas
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// Image used for the replica group
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// Resources required for the replica group
	Resources *core.Resources `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
	// RestartPolicy Determines whether pods will be restarted when they exit
	RestartPolicy        RestartPolicy `protobuf:"varint,4,opt,name=restart_policy,json=restartPolicy,proto3,enum=flyteidl.plugins.kubeflow.RestartPolicy" json:"restart_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DistributedTensorflowTrainingReplicaSpec) Reset() {
	*m = DistributedTensorflowTrainingReplicaSpec{}
}
func (m *DistributedTensorflowTrainingReplicaSpec) String() string { return proto.CompactTextString(m) }
func (*DistributedTensorflowTrainingReplicaSpec) ProtoMessage()    {}
func (*DistributedTensorflowTrainingReplicaSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_93de2bd764ddf01a, []int{1}
}

func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Unmarshal(m, b)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Marshal(b, m, deterministic)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Merge(m, src)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_Size() int {
	return xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.Size(m)
}
func (m *DistributedTensorflowTrainingReplicaSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedTensorflowTrainingReplicaSpec proto.InternalMessageInfo

func (m *DistributedTensorflowTrainingReplicaSpec) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *DistributedTensorflowTrainingReplicaSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DistributedTensorflowTrainingReplicaSpec) GetResources() *core.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DistributedTensorflowTrainingReplicaSpec) GetRestartPolicy() RestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return RestartPolicy_RESTART_POLICY_NEVER
}

func init() {
	proto.RegisterType((*DistributedTensorflowTrainingTask)(nil), "flyteidl.plugins.kubeflow.DistributedTensorflowTrainingTask")
	proto.RegisterType((*DistributedTensorflowTrainingReplicaSpec)(nil), "flyteidl.plugins.kubeflow.DistributedTensorflowTrainingReplicaSpec")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/kubeflow/tensorflow.proto", fileDescriptor_93de2bd764ddf01a)
}

var fileDescriptor_93de2bd764ddf01a = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6a, 0xe3, 0x30,
	0x10, 0x86, 0xc9, 0x6e, 0xb2, 0x6c, 0x14, 0xe2, 0x65, 0xc5, 0x1e, 0xb2, 0x39, 0xed, 0x86, 0x65,
	0x09, 0x85, 0x5a, 0x90, 0x42, 0x6f, 0xa5, 0xd0, 0xf4, 0xde, 0xa2, 0xe6, 0xd4, 0x4b, 0x90, 0x15,
	0xc5, 0x51, 0x2d, 0x4b, 0x62, 0x24, 0x35, 0xe4, 0x5d, 0xfa, 0x7a, 0x7d, 0x8f, 0x12, 0x3b, 0xb6,
	0xd3, 0x42, 0x42, 0x0f, 0xb9, 0xcd, 0x58, 0xff, 0xfc, 0xdf, 0x78, 0xa4, 0x41, 0x67, 0x4b, 0xb5,
	0xf1, 0x42, 0x2e, 0x14, 0xb1, 0x2a, 0xa4, 0x52, 0x3b, 0x92, 0x85, 0x44, 0x2c, 0x95, 0x59, 0x13,
	0x2f, 0xb4, 0x33, 0xb0, 0x0d, 0x63, 0x0b, 0xc6, 0x1b, 0xfc, 0xbb, 0xd2, 0xc6, 0x3b, 0x6d, 0x5c,
	0x69, 0x87, 0xf5, 0x11, 0xe1, 0x06, 0x04, 0xf1, 0xcc, 0x65, 0xae, 0xac, 0x1a, 0xfe, 0x3f, 0x4c,
	0xe0, 0x26, 0xcf, 0x8d, 0x2e, 0x75, 0xa3, 0x97, 0x36, 0xfa, 0x7b, 0x2b, 0x9d, 0x07, 0x99, 0x04,
	0x2f, 0x16, 0xb3, 0x9a, 0x3e, 0x03, 0x26, 0xb5, 0xd4, 0xe9, 0x8c, 0xb9, 0x0c, 0x2b, 0xf4, 0x63,
	0x6d, 0x20, 0x13, 0x30, 0x07, 0x61, 0x95, 0xe4, 0xcc, 0x0d, 0x5a, 0x7f, 0x5a, 0xe3, 0xde, 0x64,
	0x1a, 0x1f, 0xec, 0x2e, 0x3e, 0x6a, 0x4b, 0x4b, 0x9f, 0x07, 0x2b, 0x38, 0x8d, 0x4a, 0xef, 0xdd,
	0x27, 0x87, 0x17, 0xa8, 0x67, 0x5d, 0x43, 0xfa, 0x72, 0x3a, 0x12, 0xb2, 0xae, 0xa6, 0x3c, 0xa1,
	0x88, 0xaf, 0xa4, 0x58, 0x36, 0xa0, 0xaf, 0xa7, 0x03, 0xf5, 0x0b, 0xeb, 0x9a, 0x35, 0x45, 0x08,
	0x82, 0x9e, 0x5b, 0xa3, 0x24, 0xdf, 0x0c, 0xda, 0x05, 0xe7, 0xdf, 0x11, 0x0e, 0x0d, 0xfa, 0xbe,
	0xd0, 0xd2, 0x2e, 0x54, 0x21, 0x06, 0x84, 0xc5, 0x33, 0x53, 0x81, 0x79, 0xb3, 0x77, 0x0f, 0x9d,
	0xd3, 0x35, 0xfd, 0xb3, 0xb6, 0xaf, 0x1a, 0x1f, 0xbd, 0xb6, 0xd0, 0xf8, 0xb3, 0xf5, 0x78, 0x88,
	0xbe, 0xbf, 0x7b, 0x1e, 0x1d, 0x5a, 0xe7, 0xf8, 0x17, 0xea, 0xc8, 0x9c, 0xa5, 0xa2, 0xb8, 0xcd,
	0x2e, 0x2d, 0x13, 0x7c, 0x89, 0xba, 0x20, 0x9c, 0x09, 0xc0, 0x45, 0x35, 0xfe, 0x41, 0xf3, 0x27,
	0xdb, 0x47, 0x1d, 0xd3, 0xea, 0x9c, 0x36, 0x52, 0x7c, 0x87, 0x22, 0x10, 0xce, 0x33, 0xf0, 0xfb,
	0x33, 0x8d, 0x26, 0xe3, 0x63, 0x33, 0x2d, 0x0b, 0x76, 0x73, 0xed, 0xc3, 0x7e, 0x7a, 0x73, 0xfd,
	0x78, 0x95, 0x4a, 0xbf, 0x0a, 0x49, 0xcc, 0x4d, 0x4e, 0x0a, 0x13, 0x03, 0x69, 0x19, 0x90, 0x7a,
	0x95, 0x52, 0xa1, 0x89, 0x4d, 0xce, 0x53, 0x43, 0x3e, 0x6e, 0x57, 0xf2, 0xad, 0x58, 0xa7, 0x8b,
	0xb7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0xcf, 0x60, 0xf4, 0xda, 0x03, 0x00, 0x00,
}
