// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: mars/mars.proto

package mars

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ElementType int32

const (
	ElementType_ElementTypeUnknown     ElementType = 0
	ElementType_ElementTypeInput       ElementType = 1
	ElementType_ElementTypeInputNumber ElementType = 2
	ElementType_ElementTypeSelect      ElementType = 3
	ElementType_ElementTypeRadio       ElementType = 4
	ElementType_ElementTypeSwitch      ElementType = 5
	ElementType_ElementTypeTextArea    ElementType = 6
)

// Enum value maps for ElementType.
var (
	ElementType_name = map[int32]string{
		0: "ElementTypeUnknown",
		1: "ElementTypeInput",
		2: "ElementTypeInputNumber",
		3: "ElementTypeSelect",
		4: "ElementTypeRadio",
		5: "ElementTypeSwitch",
		6: "ElementTypeTextArea",
	}
	ElementType_value = map[string]int32{
		"ElementTypeUnknown":     0,
		"ElementTypeInput":       1,
		"ElementTypeInputNumber": 2,
		"ElementTypeSelect":      3,
		"ElementTypeRadio":       4,
		"ElementTypeSwitch":      5,
		"ElementTypeTextArea":    6,
	}
)

func (x ElementType) Enum() *ElementType {
	p := new(ElementType)
	*p = x
	return p
}

func (x ElementType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ElementType) Descriptor() protoreflect.EnumDescriptor {
	return file_mars_mars_proto_enumTypes[0].Descriptor()
}

func (ElementType) Type() protoreflect.EnumType {
	return &file_mars_mars_proto_enumTypes[0]
}

func (x ElementType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ElementType.Descriptor instead.
func (ElementType) EnumDescriptor() ([]byte, []int) {
	return file_mars_mars_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// config_file 指定项目下的默认配置文件, 也可以是别的项目的文件，格式为 "pid|branch|filename"
	ConfigFile string `protobuf:"bytes,1,opt,name=config_file,json=configFile,proto3" json:"config_file,omitempty"`
	// config_file_values 全局配置文件，如果没有 ConfigFile 则使用这个
	ConfigFileValues string `protobuf:"bytes,2,opt,name=config_file_values,json=configFileValues,proto3" json:"config_file_values,omitempty"`
	ConfigField      string `protobuf:"bytes,3,opt,name=config_field,json=configField,proto3" json:"config_field,omitempty"`
	IsSimpleEnv      bool   `protobuf:"varint,4,opt,name=is_simple_env,json=isSimpleEnv,proto3" json:"is_simple_env,omitempty"`
	// config_file_type 配置文件类型，php/env/yaml...
	ConfigFileType string `protobuf:"bytes,5,opt,name=config_file_type,json=configFileType,proto3" json:"config_file_type,omitempty"`
	// local_chart_path helm charts 目录, charts 文件在项目中存放的目录(必填), 也可以是别的项目的文件，格式为 "pid|branch|path"
	LocalChartPath string `protobuf:"bytes,6,opt,name=local_chart_path,json=localChartPath,proto3" json:"local_chart_path,omitempty"`
	// branches 启用的分支
	Branches []string `protobuf:"bytes,7,rep,name=branches,proto3" json:"branches,omitempty"`
	// values_yaml 和 values.yaml 一样
	ValuesYaml string `protobuf:"bytes,8,opt,name=values_yaml,json=valuesYaml,proto3" json:"values_yaml,omitempty"`
	// elements 自定义字段
	Elements []*Element `protobuf:"bytes,9,rep,name=elements,proto3" json:"elements,omitempty"`
	// 显示的名称 (helm app name), 不填就使用 git server project name
	// 以字母开头结尾，中间可以有 '_' '-'
	DisplayName string `protobuf:"bytes,10,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_mars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_mars_mars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_mars_mars_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetConfigFile() string {
	if x != nil {
		return x.ConfigFile
	}
	return ""
}

func (x *Config) GetConfigFileValues() string {
	if x != nil {
		return x.ConfigFileValues
	}
	return ""
}

func (x *Config) GetConfigField() string {
	if x != nil {
		return x.ConfigField
	}
	return ""
}

func (x *Config) GetIsSimpleEnv() bool {
	if x != nil {
		return x.IsSimpleEnv
	}
	return false
}

func (x *Config) GetConfigFileType() string {
	if x != nil {
		return x.ConfigFileType
	}
	return ""
}

func (x *Config) GetLocalChartPath() string {
	if x != nil {
		return x.LocalChartPath
	}
	return ""
}

func (x *Config) GetBranches() []string {
	if x != nil {
		return x.Branches
	}
	return nil
}

func (x *Config) GetValuesYaml() string {
	if x != nil {
		return x.ValuesYaml
	}
	return ""
}

func (x *Config) GetElements() []*Element {
	if x != nil {
		return x.Elements
	}
	return nil
}

func (x *Config) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type Element struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path         string      `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Type         ElementType `protobuf:"varint,2,opt,name=type,proto3,enum=mars.ElementType" json:"type,omitempty"`
	Default      string      `protobuf:"bytes,3,opt,name=default,proto3" json:"default,omitempty"`
	Description  string      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	SelectValues []string    `protobuf:"bytes,6,rep,name=select_values,json=selectValues,proto3" json:"select_values,omitempty"`
	Order        uint32      `protobuf:"varint,7,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *Element) Reset() {
	*x = Element{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mars_mars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Element) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Element) ProtoMessage() {}

func (x *Element) ProtoReflect() protoreflect.Message {
	mi := &file_mars_mars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Element.ProtoReflect.Descriptor instead.
func (*Element) Descriptor() ([]byte, []int) {
	return file_mars_mars_proto_rawDescGZIP(), []int{1}
}

func (x *Element) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Element) GetType() ElementType {
	if x != nil {
		return x.Type
	}
	return ElementType_ElementTypeUnknown
}

func (x *Element) GetDefault() string {
	if x != nil {
		return x.Default
	}
	return ""
}

func (x *Element) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Element) GetSelectValues() []string {
	if x != nil {
		return x.SelectValues
	}
	return nil
}

func (x *Element) GetOrder() uint32 {
	if x != nil {
		return x.Order
	}
	return 0
}

var File_mars_mars_proto protoreflect.FileDescriptor

var file_mars_mars_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6d, 0x61, 0x72, 0x73, 0x2f, 0x6d, 0x61, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6d, 0x61, 0x72, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa8, 0x03, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x22, 0x0a,
	0x0d, 0x69, 0x73, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x6e,
	0x76, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x5f, 0x79, 0x61, 0x6d, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x59, 0x61,
	0x6d, 0x6c, 0x12, 0x29, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x61, 0x72, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4c, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xfa, 0x42, 0x26, 0x72, 0x24, 0x28, 0x40, 0x32, 0x1d, 0x5e, 0x5b,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x5b, 0x41, 0x2d, 0x5a, 0x2d, 0x5f, 0x61, 0x2d,
	0x7a, 0x5d, 0x2a, 0x5b, 0x5e, 0x5f, 0x2d, 0x5d, 0x29, 0x2a, 0x24, 0xd0, 0x01, 0x01, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x07,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x6d, 0x61, 0x72, 0x73, 0x2e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52,
	0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2a, 0xb4, 0x01, 0x0a, 0x0b, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x12,
	0x15, 0x0a, 0x11, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x61, 0x64, 0x69, 0x6f, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x53, 0x77, 0x69, 0x74, 0x63,
	0x68, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x54, 0x65, 0x78, 0x74, 0x41, 0x72, 0x65, 0x61, 0x10, 0x06, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x2d, 0x63,
	0x6e, 0x7a, 0x6a, 0x2f, 0x6d, 0x61, 0x72, 0x73, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x34, 0x2f, 0x6d, 0x61, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mars_mars_proto_rawDescOnce sync.Once
	file_mars_mars_proto_rawDescData = file_mars_mars_proto_rawDesc
)

func file_mars_mars_proto_rawDescGZIP() []byte {
	file_mars_mars_proto_rawDescOnce.Do(func() {
		file_mars_mars_proto_rawDescData = protoimpl.X.CompressGZIP(file_mars_mars_proto_rawDescData)
	})
	return file_mars_mars_proto_rawDescData
}

var file_mars_mars_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mars_mars_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mars_mars_proto_goTypes = []interface{}{
	(ElementType)(0), // 0: mars.ElementType
	(*Config)(nil),   // 1: mars.Config
	(*Element)(nil),  // 2: mars.Element
}
var file_mars_mars_proto_depIdxs = []int32{
	2, // 0: mars.Config.elements:type_name -> mars.Element
	0, // 1: mars.Element.type:type_name -> mars.ElementType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_mars_mars_proto_init() }
func file_mars_mars_proto_init() {
	if File_mars_mars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mars_mars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_mars_mars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Element); i {
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
			RawDescriptor: file_mars_mars_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mars_mars_proto_goTypes,
		DependencyIndexes: file_mars_mars_proto_depIdxs,
		EnumInfos:         file_mars_mars_proto_enumTypes,
		MessageInfos:      file_mars_mars_proto_msgTypes,
	}.Build()
	File_mars_mars_proto = out.File
	file_mars_mars_proto_rawDesc = nil
	file_mars_mars_proto_goTypes = nil
	file_mars_mars_proto_depIdxs = nil
}
