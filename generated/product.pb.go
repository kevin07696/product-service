// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: protos/product.proto

package generated

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductCategory int32

const (
	ProductCategory_CATEGORY_UNSPECIFIED ProductCategory = 0
	ProductCategory_CATEGORY_CLOTHING    ProductCategory = 1
	ProductCategory_CATEGORY_BAGS        ProductCategory = 2
)

// Enum value maps for ProductCategory.
var (
	ProductCategory_name = map[int32]string{
		0: "CATEGORY_UNSPECIFIED",
		1: "CATEGORY_CLOTHING",
		2: "CATEGORY_BAGS",
	}
	ProductCategory_value = map[string]int32{
		"CATEGORY_UNSPECIFIED": 0,
		"CATEGORY_CLOTHING":    1,
		"CATEGORY_BAGS":        2,
	}
)

func (x ProductCategory) Enum() *ProductCategory {
	p := new(ProductCategory)
	*p = x
	return p
}

func (x ProductCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_product_proto_enumTypes[0].Descriptor()
}

func (ProductCategory) Type() protoreflect.EnumType {
	return &file_protos_product_proto_enumTypes[0]
}

func (x ProductCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductCategory.Descriptor instead.
func (ProductCategory) EnumDescriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{0}
}

type ProductSort int32

const (
	ProductSort_SORT_UNSPECIFIED ProductSort = 0
	ProductSort_SORT_TIMESTAMP   ProductSort = 1
	ProductSort_SORT_RATING      ProductSort = 2
	ProductSort_SORT_PRICE       ProductSort = 3
)

// Enum value maps for ProductSort.
var (
	ProductSort_name = map[int32]string{
		0: "SORT_UNSPECIFIED",
		1: "SORT_TIMESTAMP",
		2: "SORT_RATING",
		3: "SORT_PRICE",
	}
	ProductSort_value = map[string]int32{
		"SORT_UNSPECIFIED": 0,
		"SORT_TIMESTAMP":   1,
		"SORT_RATING":      2,
		"SORT_PRICE":       3,
	}
)

func (x ProductSort) Enum() *ProductSort {
	p := new(ProductSort)
	*p = x
	return p
}

func (x ProductSort) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductSort) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_product_proto_enumTypes[1].Descriptor()
}

func (ProductSort) Type() protoreflect.EnumType {
	return &file_protos_product_proto_enumTypes[1]
}

func (x ProductSort) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductSort.Descriptor instead.
func (ProductSort) EnumDescriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{1}
}

type ProductStock int32

const (
	ProductStock_STOCK_UNSPECIFIED ProductStock = 0
	ProductStock_STOCK_IN          ProductStock = 1
	ProductStock_STOCK_OUT         ProductStock = 2
)

// Enum value maps for ProductStock.
var (
	ProductStock_name = map[int32]string{
		0: "STOCK_UNSPECIFIED",
		1: "STOCK_IN",
		2: "STOCK_OUT",
	}
	ProductStock_value = map[string]int32{
		"STOCK_UNSPECIFIED": 0,
		"STOCK_IN":          1,
		"STOCK_OUT":         2,
	}
)

func (x ProductStock) Enum() *ProductStock {
	p := new(ProductStock)
	*p = x
	return p
}

func (x ProductStock) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductStock) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_product_proto_enumTypes[2].Descriptor()
}

func (ProductStock) Type() protoreflect.EnumType {
	return &file_protos_product_proto_enumTypes[2]
}

func (x ProductStock) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductStock.Descriptor instead.
func (ProductStock) EnumDescriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{2}
}

type ProductSummaryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sort          ProductSort            `protobuf:"varint,1,opt,name=Sort,proto3,enum=ProductSort" json:"Sort,omitempty"`
	Category      ProductCategory        `protobuf:"varint,2,opt,name=Category,proto3,enum=ProductCategory" json:"Category,omitempty"`
	Stock         ProductStock           `protobuf:"varint,3,opt,name=Stock,proto3,enum=ProductStock" json:"Stock,omitempty"`
	MinPrice      uint32                 `protobuf:"varint,4,opt,name=MinPrice,proto3" json:"MinPrice,omitempty"`
	MaxPrice      uint32                 `protobuf:"varint,5,opt,name=MaxPrice,proto3" json:"MaxPrice,omitempty"`
	Desc          bool                   `protobuf:"varint,6,opt,name=Desc,proto3" json:"Desc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductSummaryRequest) Reset() {
	*x = ProductSummaryRequest{}
	mi := &file_protos_product_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductSummaryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSummaryRequest) ProtoMessage() {}

func (x *ProductSummaryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSummaryRequest.ProtoReflect.Descriptor instead.
func (*ProductSummaryRequest) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductSummaryRequest) GetSort() ProductSort {
	if x != nil {
		return x.Sort
	}
	return ProductSort_SORT_UNSPECIFIED
}

func (x *ProductSummaryRequest) GetCategory() ProductCategory {
	if x != nil {
		return x.Category
	}
	return ProductCategory_CATEGORY_UNSPECIFIED
}

func (x *ProductSummaryRequest) GetStock() ProductStock {
	if x != nil {
		return x.Stock
	}
	return ProductStock_STOCK_UNSPECIFIED
}

func (x *ProductSummaryRequest) GetMinPrice() uint32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *ProductSummaryRequest) GetMaxPrice() uint32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

func (x *ProductSummaryRequest) GetDesc() bool {
	if x != nil {
		return x.Desc
	}
	return false
}

type ProductSummary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ThumbnailURL  string                 `protobuf:"bytes,3,opt,name=ThumbnailURL,proto3" json:"ThumbnailURL,omitempty"`
	Price         string                 `protobuf:"bytes,4,opt,name=Price,proto3" json:"Price,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductSummary) Reset() {
	*x = ProductSummary{}
	mi := &file_protos_product_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSummary) ProtoMessage() {}

func (x *ProductSummary) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSummary.ProtoReflect.Descriptor instead.
func (*ProductSummary) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductSummary) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductSummary) GetThumbnailURL() string {
	if x != nil {
		return x.ThumbnailURL
	}
	return ""
}

func (x *ProductSummary) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

type ProductSummaryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Products      []*ProductSummary      `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductSummaryResponse) Reset() {
	*x = ProductSummaryResponse{}
	mi := &file_protos_product_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductSummaryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductSummaryResponse) ProtoMessage() {}

func (x *ProductSummaryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductSummaryResponse.ProtoReflect.Descriptor instead.
func (*ProductSummaryResponse) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductSummaryResponse) GetProducts() []*ProductSummary {
	if x != nil {
		return x.Products
	}
	return nil
}

type ProductDetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductDetailRequest) Reset() {
	*x = ProductDetailRequest{}
	mi := &file_protos_product_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailRequest) ProtoMessage() {}

func (x *ProductDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetailRequest.ProtoReflect.Descriptor instead.
func (*ProductDetailRequest) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{3}
}

func (x *ProductDetailRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ProductDetailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ID            uint32                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Quantity      uint32                 `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Category      ProductCategory        `protobuf:"varint,3,opt,name=Category,proto3,enum=ProductCategory" json:"Category,omitempty"`
	Price         string                 `protobuf:"bytes,4,opt,name=Price,proto3" json:"Price,omitempty"`
	Name          string                 `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	ImagesURLs    string                 `protobuf:"bytes,6,opt,name=ImagesURLs,proto3" json:"ImagesURLs,omitempty"`
	Description   string                 `protobuf:"bytes,7,opt,name=Description,proto3" json:"Description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductDetailResponse) Reset() {
	*x = ProductDetailResponse{}
	mi := &file_protos_product_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetailResponse) ProtoMessage() {}

func (x *ProductDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_product_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetailResponse.ProtoReflect.Descriptor instead.
func (*ProductDetailResponse) Descriptor() ([]byte, []int) {
	return file_protos_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductDetailResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProductDetailResponse) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *ProductDetailResponse) GetCategory() ProductCategory {
	if x != nil {
		return x.Category
	}
	return ProductCategory_CATEGORY_UNSPECIFIED
}

func (x *ProductDetailResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *ProductDetailResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductDetailResponse) GetImagesURLs() string {
	if x != nil {
		return x.ImagesURLs
	}
	return ""
}

func (x *ProductDetailResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_protos_product_proto protoreflect.FileDescriptor

var file_protos_product_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x04, 0x53, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x52, 0x04, 0x53, 0x6f,
	0x72, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x23, 0x0a, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x05,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x4d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x4d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x44, 0x65, 0x73,
	0x63, 0x22, 0x6e, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x45, 0x0a, 0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x26, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44,
	0x22, 0xdd, 0x01, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x51, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x55, 0x52, 0x4c, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x55, 0x52, 0x4c, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2a, 0x55, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x54, 0x48, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59,
	0x5f, 0x42, 0x41, 0x47, 0x53, 0x10, 0x02, 0x2a, 0x58, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x53, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x4f, 0x52, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x43, 0x45, 0x10,
	0x03, 0x2a, 0x42, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x4f, 0x43, 0x4b, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x43,
	0x4b, 0x5f, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54, 0x4f, 0x43, 0x4b, 0x5f,
	0x4f, 0x55, 0x54, 0x10, 0x02, 0x32, 0x89, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x41, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x75, 0x6d,
	0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x15, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_protos_product_proto_rawDescOnce sync.Once
	file_protos_product_proto_rawDescData []byte
)

func file_protos_product_proto_rawDescGZIP() []byte {
	file_protos_product_proto_rawDescOnce.Do(func() {
		file_protos_product_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protos_product_proto_rawDesc), len(file_protos_product_proto_rawDesc)))
	})
	return file_protos_product_proto_rawDescData
}

var file_protos_product_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protos_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_product_proto_goTypes = []any{
	(ProductCategory)(0),           // 0: ProductCategory
	(ProductSort)(0),               // 1: ProductSort
	(ProductStock)(0),              // 2: ProductStock
	(*ProductSummaryRequest)(nil),  // 3: ProductSummaryRequest
	(*ProductSummary)(nil),         // 4: ProductSummary
	(*ProductSummaryResponse)(nil), // 5: ProductSummaryResponse
	(*ProductDetailRequest)(nil),   // 6: ProductDetailRequest
	(*ProductDetailResponse)(nil),  // 7: ProductDetailResponse
}
var file_protos_product_proto_depIdxs = []int32{
	1, // 0: ProductSummaryRequest.Sort:type_name -> ProductSort
	0, // 1: ProductSummaryRequest.Category:type_name -> ProductCategory
	2, // 2: ProductSummaryRequest.Stock:type_name -> ProductStock
	4, // 3: ProductSummaryResponse.products:type_name -> ProductSummary
	0, // 4: ProductDetailResponse.Category:type_name -> ProductCategory
	3, // 5: Product.FilterProducts:input_type -> ProductSummaryRequest
	6, // 6: Product.GetDetails:input_type -> ProductDetailRequest
	5, // 7: Product.FilterProducts:output_type -> ProductSummaryResponse
	7, // 8: Product.GetDetails:output_type -> ProductDetailResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_product_proto_init() }
func file_protos_product_proto_init() {
	if File_protos_product_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protos_product_proto_rawDesc), len(file_protos_product_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_product_proto_goTypes,
		DependencyIndexes: file_protos_product_proto_depIdxs,
		EnumInfos:         file_protos_product_proto_enumTypes,
		MessageInfos:      file_protos_product_proto_msgTypes,
	}.Build()
	File_protos_product_proto = out.File
	file_protos_product_proto_goTypes = nil
	file_protos_product_proto_depIdxs = nil
}
