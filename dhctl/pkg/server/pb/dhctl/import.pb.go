// Copyright 2024 Flant JSC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: import.proto

package dhctl

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ImportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ImportRequest_Start
	//	*ImportRequest_Continue
	Message isImportRequest_Message `protobuf_oneof:"message"`
}

func (x *ImportRequest) Reset() {
	*x = ImportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportRequest) ProtoMessage() {}

func (x *ImportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportRequest.ProtoReflect.Descriptor instead.
func (*ImportRequest) Descriptor() ([]byte, []int) {
	return file_import_proto_rawDescGZIP(), []int{0}
}

func (m *ImportRequest) GetMessage() isImportRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ImportRequest) GetStart() *ImportStart {
	if x, ok := x.GetMessage().(*ImportRequest_Start); ok {
		return x.Start
	}
	return nil
}

func (x *ImportRequest) GetContinue() *ImportContinue {
	if x, ok := x.GetMessage().(*ImportRequest_Continue); ok {
		return x.Continue
	}
	return nil
}

type isImportRequest_Message interface {
	isImportRequest_Message()
}

type ImportRequest_Start struct {
	Start *ImportStart `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

type ImportRequest_Continue struct {
	Continue *ImportContinue `protobuf:"bytes,2,opt,name=continue,proto3,oneof"`
}

func (*ImportRequest_Start) isImportRequest_Message() {}

func (*ImportRequest_Continue) isImportRequest_Message() {}

type ImportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*ImportResponse_Result
	//	*ImportResponse_PhaseEnd
	//	*ImportResponse_Logs
	Message isImportResponse_Message `protobuf_oneof:"message"`
}

func (x *ImportResponse) Reset() {
	*x = ImportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResponse) ProtoMessage() {}

func (x *ImportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResponse.ProtoReflect.Descriptor instead.
func (*ImportResponse) Descriptor() ([]byte, []int) {
	return file_import_proto_rawDescGZIP(), []int{1}
}

func (m *ImportResponse) GetMessage() isImportResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *ImportResponse) GetResult() *ImportResult {
	if x, ok := x.GetMessage().(*ImportResponse_Result); ok {
		return x.Result
	}
	return nil
}

func (x *ImportResponse) GetPhaseEnd() *ImportPhaseEnd {
	if x, ok := x.GetMessage().(*ImportResponse_PhaseEnd); ok {
		return x.PhaseEnd
	}
	return nil
}

func (x *ImportResponse) GetLogs() *Logs {
	if x, ok := x.GetMessage().(*ImportResponse_Logs); ok {
		return x.Logs
	}
	return nil
}

type isImportResponse_Message interface {
	isImportResponse_Message()
}

type ImportResponse_Result struct {
	Result *ImportResult `protobuf:"bytes,1,opt,name=result,proto3,oneof"`
}

type ImportResponse_PhaseEnd struct {
	PhaseEnd *ImportPhaseEnd `protobuf:"bytes,2,opt,name=phase_end,json=phaseEnd,proto3,oneof"`
}

type ImportResponse_Logs struct {
	Logs *Logs `protobuf:"bytes,3,opt,name=logs,proto3,oneof"`
}

func (*ImportResponse_Result) isImportResponse_Message() {}

func (*ImportResponse_PhaseEnd) isImportResponse_Message() {}

func (*ImportResponse_Logs) isImportResponse_Message() {}

type ImportStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionConfig  string              `protobuf:"bytes,1,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	ScanOnly          *bool               `protobuf:"varint,2,opt,name=scan_only,json=scanOnly,proto3,oneof" json:"scan_only,omitempty"`
	ResourcesTemplate string              `protobuf:"bytes,3,opt,name=resources_template,json=resourcesTemplate,proto3" json:"resources_template,omitempty"`
	ResourcesValues   *structpb.Struct    `protobuf:"bytes,4,opt,name=resources_values,json=resourcesValues,proto3" json:"resources_values,omitempty"`
	Options           *ImportStartOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *ImportStart) Reset() {
	*x = ImportStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportStart) ProtoMessage() {}

func (x *ImportStart) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportStart.ProtoReflect.Descriptor instead.
func (*ImportStart) Descriptor() ([]byte, []int) {
	return file_import_proto_rawDescGZIP(), []int{2}
}

func (x *ImportStart) GetConnectionConfig() string {
	if x != nil {
		return x.ConnectionConfig
	}
	return ""
}

func (x *ImportStart) GetScanOnly() bool {
	if x != nil && x.ScanOnly != nil {
		return *x.ScanOnly
	}
	return false
}

func (x *ImportStart) GetResourcesTemplate() string {
	if x != nil {
		return x.ResourcesTemplate
	}
	return ""
}

func (x *ImportStart) GetResourcesValues() *structpb.Struct {
	if x != nil {
		return x.ResourcesValues
	}
	return nil
}

func (x *ImportStart) GetOptions() *ImportStartOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

type ImportPhaseEnd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompletedPhase      string            `protobuf:"bytes,1,opt,name=completed_phase,json=completedPhase,proto3" json:"completed_phase,omitempty"`
	CompletedPhaseState map[string][]byte `protobuf:"bytes,2,rep,name=completed_phase_state,json=completedPhaseState,proto3" json:"completed_phase_state,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CompletedPhaseData  string            `protobuf:"bytes,3,opt,name=completed_phase_data,json=completedPhaseData,proto3" json:"completed_phase_data,omitempty"`
	NextPhase           string            `protobuf:"bytes,4,opt,name=next_phase,json=nextPhase,proto3" json:"next_phase,omitempty"`
	NextPhaseCritical   bool              `protobuf:"varint,5,opt,name=next_phase_critical,json=nextPhaseCritical,proto3" json:"next_phase_critical,omitempty"`
}

func (x *ImportPhaseEnd) Reset() {
	*x = ImportPhaseEnd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportPhaseEnd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportPhaseEnd) ProtoMessage() {}

func (x *ImportPhaseEnd) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportPhaseEnd.ProtoReflect.Descriptor instead.
func (*ImportPhaseEnd) Descriptor() ([]byte, []int) {
	return file_import_proto_rawDescGZIP(), []int{3}
}

func (x *ImportPhaseEnd) GetCompletedPhase() string {
	if x != nil {
		return x.CompletedPhase
	}
	return ""
}

func (x *ImportPhaseEnd) GetCompletedPhaseState() map[string][]byte {
	if x != nil {
		return x.CompletedPhaseState
	}
	return nil
}

func (x *ImportPhaseEnd) GetCompletedPhaseData() string {
	if x != nil {
		return x.CompletedPhaseData
	}
	return ""
}

func (x *ImportPhaseEnd) GetNextPhase() string {
	if x != nil {
		return x.NextPhase
	}
	return ""
}

func (x *ImportPhaseEnd) GetNextPhaseCritical() bool {
	if x != nil {
		return x.NextPhaseCritical
	}
	return false
}

type ImportContinue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Continue Continue `protobuf:"varint,1,opt,name=continue,proto3,enum=dhctl.Continue" json:"continue,omitempty"`
	Err      string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ImportContinue) Reset() {
	*x = ImportContinue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportContinue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportContinue) ProtoMessage() {}

func (x *ImportContinue) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportContinue.ProtoReflect.Descriptor instead.
func (*ImportContinue) Descriptor() ([]byte, []int) {
	return file_import_proto_rawDescGZIP(), []int{4}
}

func (x *ImportContinue) GetContinue() Continue {
	if x != nil {
		return x.Continue
	}
	return Continue_CONTINUE_UNSPECIFIED
}

func (x *ImportContinue) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type ImportStartOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommanderMode    bool                 `protobuf:"varint,1,opt,name=commander_mode,json=commanderMode,proto3" json:"commander_mode,omitempty"`
	LogWidth         int32                `protobuf:"varint,2,opt,name=log_width,json=logWidth,proto3" json:"log_width,omitempty"`
	ResourcesTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=resources_timeout,json=resourcesTimeout,proto3" json:"resources_timeout,omitempty"`
	DeckhouseTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=deckhouse_timeout,json=deckhouseTimeout,proto3" json:"deckhouse_timeout,omitempty"`
}

func (x *ImportStartOptions) Reset() {
	*x = ImportStartOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportStartOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportStartOptions) ProtoMessage() {}

func (x *ImportStartOptions) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportStartOptions.ProtoReflect.Descriptor instead.
func (*ImportStartOptions) Descriptor() ([]byte, []int) {
	return file_import_proto_rawDescGZIP(), []int{5}
}

func (x *ImportStartOptions) GetCommanderMode() bool {
	if x != nil {
		return x.CommanderMode
	}
	return false
}

func (x *ImportStartOptions) GetLogWidth() int32 {
	if x != nil {
		return x.LogWidth
	}
	return 0
}

func (x *ImportStartOptions) GetResourcesTimeout() *durationpb.Duration {
	if x != nil {
		return x.ResourcesTimeout
	}
	return nil
}

func (x *ImportStartOptions) GetDeckhouseTimeout() *durationpb.Duration {
	if x != nil {
		return x.DeckhouseTimeout
	}
	return nil
}

type ImportResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State  string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Result string `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	Err    string `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ImportResult) Reset() {
	*x = ImportResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_import_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportResult) ProtoMessage() {}

func (x *ImportResult) ProtoReflect() protoreflect.Message {
	mi := &file_import_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportResult.ProtoReflect.Descriptor instead.
func (*ImportResult) Descriptor() ([]byte, []int) {
	return file_import_proto_rawDescGZIP(), []int{6}
}

func (x *ImportResult) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ImportResult) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *ImportResult) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_import_proto protoreflect.FileDescriptor

var file_import_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x64, 0x68, 0x63, 0x74, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x48, 0x00, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x33,
	0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x48, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69,
	0x6e, 0x75, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa3,
	0x01, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x34, 0x0a, 0x09, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x08, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x4c, 0x6f, 0x67,
	0x73, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x92, 0x02, 0x0a, 0x0b, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x73, 0x63, 0x61, 0x6e, 0x4f, 0x6e, 0x6c, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x73, 0x63, 0x61, 0x6e, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x22, 0xe6, 0x02, 0x0a, 0x0e, 0x49, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x27, 0x0a, 0x0f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x49, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50,
	0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x68, 0x61, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x68, 0x61,
	0x73, 0x65, 0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x1a, 0x46, 0x0a, 0x18, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x50, 0x68, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x4f, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x69, 0x6e, 0x75, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x22, 0xe8, 0x01, 0x0a, 0x12, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x46,
	0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x46, 0x0a, 0x11, 0x64, 0x65, 0x63, 0x6b, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x64, 0x65,
	0x63, 0x6b, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x4e,
	0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42, 0x0a,
	0x5a, 0x08, 0x70, 0x62, 0x2f, 0x64, 0x68, 0x63, 0x74, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_import_proto_rawDescOnce sync.Once
	file_import_proto_rawDescData = file_import_proto_rawDesc
)

func file_import_proto_rawDescGZIP() []byte {
	file_import_proto_rawDescOnce.Do(func() {
		file_import_proto_rawDescData = protoimpl.X.CompressGZIP(file_import_proto_rawDescData)
	})
	return file_import_proto_rawDescData
}

var file_import_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_import_proto_goTypes = []interface{}{
	(*ImportRequest)(nil),       // 0: dhctl.ImportRequest
	(*ImportResponse)(nil),      // 1: dhctl.ImportResponse
	(*ImportStart)(nil),         // 2: dhctl.ImportStart
	(*ImportPhaseEnd)(nil),      // 3: dhctl.ImportPhaseEnd
	(*ImportContinue)(nil),      // 4: dhctl.ImportContinue
	(*ImportStartOptions)(nil),  // 5: dhctl.ImportStartOptions
	(*ImportResult)(nil),        // 6: dhctl.ImportResult
	nil,                         // 7: dhctl.ImportPhaseEnd.CompletedPhaseStateEntry
	(*Logs)(nil),                // 8: dhctl.Logs
	(*structpb.Struct)(nil),     // 9: google.protobuf.Struct
	(Continue)(0),               // 10: dhctl.Continue
	(*durationpb.Duration)(nil), // 11: google.protobuf.Duration
}
var file_import_proto_depIdxs = []int32{
	2,  // 0: dhctl.ImportRequest.start:type_name -> dhctl.ImportStart
	4,  // 1: dhctl.ImportRequest.continue:type_name -> dhctl.ImportContinue
	6,  // 2: dhctl.ImportResponse.result:type_name -> dhctl.ImportResult
	3,  // 3: dhctl.ImportResponse.phase_end:type_name -> dhctl.ImportPhaseEnd
	8,  // 4: dhctl.ImportResponse.logs:type_name -> dhctl.Logs
	9,  // 5: dhctl.ImportStart.resources_values:type_name -> google.protobuf.Struct
	5,  // 6: dhctl.ImportStart.options:type_name -> dhctl.ImportStartOptions
	7,  // 7: dhctl.ImportPhaseEnd.completed_phase_state:type_name -> dhctl.ImportPhaseEnd.CompletedPhaseStateEntry
	10, // 8: dhctl.ImportContinue.continue:type_name -> dhctl.Continue
	11, // 9: dhctl.ImportStartOptions.resources_timeout:type_name -> google.protobuf.Duration
	11, // 10: dhctl.ImportStartOptions.deckhouse_timeout:type_name -> google.protobuf.Duration
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_import_proto_init() }
func file_import_proto_init() {
	if File_import_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_import_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportRequest); i {
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
		file_import_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResponse); i {
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
		file_import_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportStart); i {
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
		file_import_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportPhaseEnd); i {
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
		file_import_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportContinue); i {
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
		file_import_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportStartOptions); i {
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
		file_import_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportResult); i {
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
	file_import_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ImportRequest_Start)(nil),
		(*ImportRequest_Continue)(nil),
	}
	file_import_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ImportResponse_Result)(nil),
		(*ImportResponse_PhaseEnd)(nil),
		(*ImportResponse_Logs)(nil),
	}
	file_import_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_import_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_import_proto_goTypes,
		DependencyIndexes: file_import_proto_depIdxs,
		MessageInfos:      file_import_proto_msgTypes,
	}.Build()
	File_import_proto = out.File
	file_import_proto_rawDesc = nil
	file_import_proto_goTypes = nil
	file_import_proto_depIdxs = nil
}
