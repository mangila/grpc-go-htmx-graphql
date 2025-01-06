// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: common/enum_model.proto

package model

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

type OrderStatus int32

const (
	OrderStatus_PENDING   OrderStatus = 0
	OrderStatus_APPROVED  OrderStatus = 1
	OrderStatus_SHIPPED   OrderStatus = 2
	OrderStatus_DELIVERED OrderStatus = 3
	OrderStatus_CANCELLED OrderStatus = 4
)

// Enum value maps for OrderStatus.
var (
	OrderStatus_name = map[int32]string{
		0: "PENDING",
		1: "APPROVED",
		2: "SHIPPED",
		3: "DELIVERED",
		4: "CANCELLED",
	}
	OrderStatus_value = map[string]int32{
		"PENDING":   0,
		"APPROVED":  1,
		"SHIPPED":   2,
		"DELIVERED": 3,
		"CANCELLED": 4,
	}
)

func (x OrderStatus) Enum() *OrderStatus {
	p := new(OrderStatus)
	*p = x
	return p
}

func (x OrderStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_model_proto_enumTypes[0].Descriptor()
}

func (OrderStatus) Type() protoreflect.EnumType {
	return &file_common_enum_model_proto_enumTypes[0]
}

func (x OrderStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus.Descriptor instead.
func (OrderStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_model_proto_rawDescGZIP(), []int{0}
}

type PaymentMethod int32

const (
	PaymentMethod_CREDIT_CARD   PaymentMethod = 0
	PaymentMethod_DEBIT_CARD    PaymentMethod = 1
	PaymentMethod_PAYPAL        PaymentMethod = 2
	PaymentMethod_BANK_TRANSFER PaymentMethod = 3
)

// Enum value maps for PaymentMethod.
var (
	PaymentMethod_name = map[int32]string{
		0: "CREDIT_CARD",
		1: "DEBIT_CARD",
		2: "PAYPAL",
		3: "BANK_TRANSFER",
	}
	PaymentMethod_value = map[string]int32{
		"CREDIT_CARD":   0,
		"DEBIT_CARD":    1,
		"PAYPAL":        2,
		"BANK_TRANSFER": 3,
	}
)

func (x PaymentMethod) Enum() *PaymentMethod {
	p := new(PaymentMethod)
	*p = x
	return p
}

func (x PaymentMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_model_proto_enumTypes[1].Descriptor()
}

func (PaymentMethod) Type() protoreflect.EnumType {
	return &file_common_enum_model_proto_enumTypes[1]
}

func (x PaymentMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethod.Descriptor instead.
func (PaymentMethod) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_model_proto_rawDescGZIP(), []int{1}
}

var File_common_enum_model_proto protoreflect.FileDescriptor

var file_common_enum_model_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2a, 0x53, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x48, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x49,
	0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45,
	0x4c, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x2a, 0x4f, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45, 0x44, 0x49,
	0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x45, 0x42, 0x49,
	0x54, 0x5f, 0x43, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x59, 0x50,
	0x41, 0x4c, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x54, 0x52, 0x41,
	0x4e, 0x53, 0x46, 0x45, 0x52, 0x10, 0x03, 0x42, 0x0e, 0x5a, 0x0c, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enum_model_proto_rawDescOnce sync.Once
	file_common_enum_model_proto_rawDescData = file_common_enum_model_proto_rawDesc
)

func file_common_enum_model_proto_rawDescGZIP() []byte {
	file_common_enum_model_proto_rawDescOnce.Do(func() {
		file_common_enum_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enum_model_proto_rawDescData)
	})
	return file_common_enum_model_proto_rawDescData
}

var file_common_enum_model_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_common_enum_model_proto_goTypes = []any{
	(OrderStatus)(0),   // 0: common.OrderStatus
	(PaymentMethod)(0), // 1: common.PaymentMethod
}
var file_common_enum_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enum_model_proto_init() }
func file_common_enum_model_proto_init() {
	if File_common_enum_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enum_model_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enum_model_proto_goTypes,
		DependencyIndexes: file_common_enum_model_proto_depIdxs,
		EnumInfos:         file_common_enum_model_proto_enumTypes,
	}.Build()
	File_common_enum_model_proto = out.File
	file_common_enum_model_proto_rawDesc = nil
	file_common_enum_model_proto_goTypes = nil
	file_common_enum_model_proto_depIdxs = nil
}