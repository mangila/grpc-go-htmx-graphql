// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.12.4
// source: product.service.proto

package service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	model "shared/model"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAllProductsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*model.Product       `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllProductsResponse) Reset() {
	*x = GetAllProductsResponse{}
	mi := &file_product_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductsResponse) ProtoMessage() {}

func (x *GetAllProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductsResponse.ProtoReflect.Descriptor instead.
func (*GetAllProductsResponse) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllProductsResponse) GetProducts() []*model.Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type CommandProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price         float64                `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Stock         int32                  `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CommandProductRequest) Reset() {
	*x = CommandProductRequest{}
	mi := &file_product_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CommandProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandProductRequest) ProtoMessage() {}

func (x *CommandProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandProductRequest.ProtoReflect.Descriptor instead.
func (*CommandProductRequest) Descriptor() ([]byte, []int) {
	return file_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *CommandProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommandProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommandProductRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CommandProductRequest) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CommandProductRequest) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

var File_product_service_proto protoreflect.FileDescriptor

var file_product_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x44, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x32, 0x9a, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x67, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x47, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xe0, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x3d, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x49, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_service_proto_rawDescOnce sync.Once
	file_product_service_proto_rawDescData = file_product_service_proto_rawDesc
)

func file_product_service_proto_rawDescGZIP() []byte {
	file_product_service_proto_rawDescOnce.Do(func() {
		file_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_service_proto_rawDescData)
	})
	return file_product_service_proto_rawDescData
}

var file_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_service_proto_goTypes = []any{
	(*GetAllProductsResponse)(nil), // 0: model.GetAllProductsResponse
	(*CommandProductRequest)(nil),  // 1: model.CommandProductRequest
	(*model.Product)(nil),          // 2: model.Product
	(*wrappers.StringValue)(nil),   // 3: google.protobuf.StringValue
	(*empty.Empty)(nil),            // 4: google.protobuf.Empty
	(*wrappers.BoolValue)(nil),     // 5: google.protobuf.BoolValue
}
var file_product_service_proto_depIdxs = []int32{
	2, // 0: model.GetAllProductsResponse.products:type_name -> model.Product
	3, // 1: model.ProductQueryService.getProduct:input_type -> google.protobuf.StringValue
	4, // 2: model.ProductQueryService.getAllProducts:input_type -> google.protobuf.Empty
	1, // 3: model.ProductCommandService.createProduct:input_type -> model.CommandProductRequest
	1, // 4: model.ProductCommandService.updateProduct:input_type -> model.CommandProductRequest
	3, // 5: model.ProductCommandService.deleteProduct:input_type -> google.protobuf.StringValue
	2, // 6: model.ProductQueryService.getProduct:output_type -> model.Product
	0, // 7: model.ProductQueryService.getAllProducts:output_type -> model.GetAllProductsResponse
	2, // 8: model.ProductCommandService.createProduct:output_type -> model.Product
	2, // 9: model.ProductCommandService.updateProduct:output_type -> model.Product
	5, // 10: model.ProductCommandService.deleteProduct:output_type -> google.protobuf.BoolValue
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_service_proto_init() }
func file_product_service_proto_init() {
	if File_product_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_product_service_proto_goTypes,
		DependencyIndexes: file_product_service_proto_depIdxs,
		MessageInfos:      file_product_service_proto_msgTypes,
	}.Build()
	File_product_service_proto = out.File
	file_product_service_proto_rawDesc = nil
	file_product_service_proto_goTypes = nil
	file_product_service_proto_depIdxs = nil
}
