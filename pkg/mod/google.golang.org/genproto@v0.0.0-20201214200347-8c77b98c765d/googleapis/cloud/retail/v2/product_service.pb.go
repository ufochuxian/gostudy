// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/retail/v2/product_service.proto

package retail

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Request message for [CreateProduct][] method.
type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent catalog resource name, such as
	// "projects/*/locations/global/catalogs/default_catalog/branches/default_branch".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The [Product][google.cloud.retail.v2.Product] to create.
	Product *Product `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
	// Required. The ID to use for the [Product][google.cloud.retail.v2.Product],
	// which will become the final component of the
	// [Product.name][google.cloud.retail.v2.Product.name].
	//
	// If the caller does not have permission to create the
	// [Product][google.cloud.retail.v2.Product], regardless of whether or not it
	// exists, a PERMISSION_DENIED error is returned.
	//
	// This field must be unique among all
	// [Product][google.cloud.retail.v2.Product]s with the same
	// [parent][google.cloud.retail.v2.CreateProductRequest.parent]. Otherwise, an
	// ALREADY_EXISTS error is returned.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	ProductId string `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_product_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateProductRequest) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *CreateProductRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

// Request message for [GetProduct][] method.
type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name of [Product][google.cloud.retail.v2.Product],
	// such as
	// "projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/some_product_id".
	//
	// If the caller does not have permission to access the
	// [Product][google.cloud.retail.v2.Product], regardless of whether or not it
	// exists, a PERMISSION_DENIED error is returned.
	//
	// If the requested [Product][google.cloud.retail.v2.Product] does not exist,
	// a NOT_FOUND error is returned.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_product_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for [UpdateProduct][] method.
type UpdateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The product to update/create.
	//
	// If the caller does not have permission to update the
	// [Product][google.cloud.retail.v2.Product], regardless of whether or not it
	// exists, a PERMISSION_DENIED error is returned.
	//
	// If the [Product][google.cloud.retail.v2.Product] to update does not exist,
	// a NOT_FOUND error is returned.
	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	// Indicates which fields in the provided
	// [Product][google.cloud.retail.v2.Product] to update. The immutable and
	// output only fields are NOT supported. If not set, all supported fields (the
	// fields that are neither immutable nor output only) are updated.
	//
	// If an unsupported or unknown field is provided, an INVALID_ARGUMENT error
	// is returned.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateProductRequest) Reset() {
	*x = UpdateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProductRequest) ProtoMessage() {}

func (x *UpdateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProductRequest.ProtoReflect.Descriptor instead.
func (*UpdateProductRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_product_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateProductRequest) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *UpdateProductRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// Request message for [DeleteProduct][] method.
type DeleteProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Full resource name of [Product][google.cloud.retail.v2.Product],
	// such as
	// "projects/*/locations/global/catalogs/default_catalog/branches/default_branch/products/some_product_id".
	//
	// If the caller does not have permission to delete the
	// [Product][google.cloud.retail.v2.Product], regardless of whether or not it
	// exists, a PERMISSION_DENIED error is returned.
	//
	// If the [Product][google.cloud.retail.v2.Product] to delete does not exist,
	// a NOT_FOUND error is returned.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteProductRequest) Reset() {
	*x = DeleteProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProductRequest) ProtoMessage() {}

func (x *DeleteProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_product_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProductRequest.ProtoReflect.Descriptor instead.
func (*DeleteProductRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_product_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_retail_v2_product_service_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2_product_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76,
	0x32, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x2f, 0x76, 0x32, 0x2f, 0x70, 0x75, 0x72, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x24, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1e, 0x0a, 0x1c, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x3e, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x4e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1f,
	0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x51, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x25, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1f, 0x0a, 0x1d, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xee,
	0x08, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0xcf, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x22, 0x6f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4d, 0x22, 0x42, 0x2f, 0x76, 0x32, 0x2f,
	0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x3a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0xda, 0x41, 0x19, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x2c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x12, 0xac, 0x01, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x52,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x12, 0x43, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0xd2, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x22, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x56, 0x32, 0x4b, 0x2f, 0x76, 0x32,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a,
	0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2a, 0x7d, 0x3a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0xda, 0x41, 0x13, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2c, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x12, 0xa9, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x45, 0x2a, 0x43, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x2f,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x8e, 0x02, 0x0a, 0x0e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6c,
	0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xad, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4e, 0x22, 0x49, 0x2f,
	0x76, 0x32, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x2a, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x2a, 0x2f, 0x62, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x73, 0x3a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x3a, 0x01, 0x2a, 0xca, 0x41, 0x56, 0x0a, 0x2d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x49, 0xca, 0x41, 0x15, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42,
	0xc8, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x42, 0x13,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c, 0xaa, 0x02, 0x16, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x32, 0xea, 0x02,
	0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2_product_service_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2_product_service_proto_rawDescData = file_google_cloud_retail_v2_product_service_proto_rawDesc
)

func file_google_cloud_retail_v2_product_service_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2_product_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2_product_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2_product_service_proto_rawDescData)
	})
	return file_google_cloud_retail_v2_product_service_proto_rawDescData
}

var file_google_cloud_retail_v2_product_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_retail_v2_product_service_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),  // 0: google.cloud.retail.v2.CreateProductRequest
	(*GetProductRequest)(nil),     // 1: google.cloud.retail.v2.GetProductRequest
	(*UpdateProductRequest)(nil),  // 2: google.cloud.retail.v2.UpdateProductRequest
	(*DeleteProductRequest)(nil),  // 3: google.cloud.retail.v2.DeleteProductRequest
	(*Product)(nil),               // 4: google.cloud.retail.v2.Product
	(*fieldmaskpb.FieldMask)(nil), // 5: google.protobuf.FieldMask
	(*ImportProductsRequest)(nil), // 6: google.cloud.retail.v2.ImportProductsRequest
	(*emptypb.Empty)(nil),         // 7: google.protobuf.Empty
	(*longrunning.Operation)(nil), // 8: google.longrunning.Operation
}
var file_google_cloud_retail_v2_product_service_proto_depIdxs = []int32{
	4, // 0: google.cloud.retail.v2.CreateProductRequest.product:type_name -> google.cloud.retail.v2.Product
	4, // 1: google.cloud.retail.v2.UpdateProductRequest.product:type_name -> google.cloud.retail.v2.Product
	5, // 2: google.cloud.retail.v2.UpdateProductRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 3: google.cloud.retail.v2.ProductService.CreateProduct:input_type -> google.cloud.retail.v2.CreateProductRequest
	1, // 4: google.cloud.retail.v2.ProductService.GetProduct:input_type -> google.cloud.retail.v2.GetProductRequest
	2, // 5: google.cloud.retail.v2.ProductService.UpdateProduct:input_type -> google.cloud.retail.v2.UpdateProductRequest
	3, // 6: google.cloud.retail.v2.ProductService.DeleteProduct:input_type -> google.cloud.retail.v2.DeleteProductRequest
	6, // 7: google.cloud.retail.v2.ProductService.ImportProducts:input_type -> google.cloud.retail.v2.ImportProductsRequest
	4, // 8: google.cloud.retail.v2.ProductService.CreateProduct:output_type -> google.cloud.retail.v2.Product
	4, // 9: google.cloud.retail.v2.ProductService.GetProduct:output_type -> google.cloud.retail.v2.Product
	4, // 10: google.cloud.retail.v2.ProductService.UpdateProduct:output_type -> google.cloud.retail.v2.Product
	7, // 11: google.cloud.retail.v2.ProductService.DeleteProduct:output_type -> google.protobuf.Empty
	8, // 12: google.cloud.retail.v2.ProductService.ImportProducts:output_type -> google.longrunning.Operation
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2_product_service_proto_init() }
func file_google_cloud_retail_v2_product_service_proto_init() {
	if File_google_cloud_retail_v2_product_service_proto != nil {
		return
	}
	file_google_cloud_retail_v2_import_config_proto_init()
	file_google_cloud_retail_v2_product_proto_init()
	file_google_cloud_retail_v2_purge_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2_product_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_google_cloud_retail_v2_product_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_google_cloud_retail_v2_product_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProductRequest); i {
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
		file_google_cloud_retail_v2_product_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProductRequest); i {
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
			RawDescriptor: file_google_cloud_retail_v2_product_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_retail_v2_product_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2_product_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_retail_v2_product_service_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2_product_service_proto = out.File
	file_google_cloud_retail_v2_product_service_proto_rawDesc = nil
	file_google_cloud_retail_v2_product_service_proto_goTypes = nil
	file_google_cloud_retail_v2_product_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductServiceClient interface {
	// Creates a [Product][google.cloud.retail.v2.Product].
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error)
	// Gets a [Product][google.cloud.retail.v2.Product].
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	// Updates a [Product][google.cloud.retail.v2.Product].
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*Product, error)
	// Deletes a [Product][google.cloud.retail.v2.Product].
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Bulk import of multiple [Product][google.cloud.retail.v2.Product]s.
	//
	// Request processing may be synchronous. No partial updating is supported.
	// Non-existing items are created.
	//
	// Note that it is possible for a subset of the
	// [Product][google.cloud.retail.v2.Product]s to be successfully updated.
	ImportProducts(ctx context.Context, in *ImportProductsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2.ProductService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2.ProductService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2.ProductService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2.ProductService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ImportProducts(ctx context.Context, in *ImportProductsRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.cloud.retail.v2.ProductService/ImportProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
type ProductServiceServer interface {
	// Creates a [Product][google.cloud.retail.v2.Product].
	CreateProduct(context.Context, *CreateProductRequest) (*Product, error)
	// Gets a [Product][google.cloud.retail.v2.Product].
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	// Updates a [Product][google.cloud.retail.v2.Product].
	UpdateProduct(context.Context, *UpdateProductRequest) (*Product, error)
	// Deletes a [Product][google.cloud.retail.v2.Product].
	DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error)
	// Bulk import of multiple [Product][google.cloud.retail.v2.Product]s.
	//
	// Request processing may be synchronous. No partial updating is supported.
	// Non-existing items are created.
	//
	// Note that it is possible for a subset of the
	// [Product][google.cloud.retail.v2.Product]s to be successfully updated.
	ImportProducts(context.Context, *ImportProductsRequest) (*longrunning.Operation, error)
}

// UnimplementedProductServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (*UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (*UnimplementedProductServiceServer) GetProduct(context.Context, *GetProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedProductServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (*UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (*UnimplementedProductServiceServer) ImportProducts(context.Context, *ImportProductsRequest) (*longrunning.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProducts not implemented")
}

func RegisterProductServiceServer(s *grpc.Server, srv ProductServiceServer) {
	s.RegisterService(&_ProductService_serviceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2.ProductService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2.ProductService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2.ProductService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2.ProductService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ImportProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ImportProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.retail.v2.ProductService/ImportProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ImportProducts(ctx, req.(*ImportProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.retail.v2.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductService_GetProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductService_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "ImportProducts",
			Handler:    _ProductService_ImportProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/retail/v2/product_service.proto",
}
