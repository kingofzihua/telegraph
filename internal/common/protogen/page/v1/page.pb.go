// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: page/v1/page.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Author      string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Thumbnail   string `protobuf:"bytes,4,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Content     string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreatePageRequest) Reset() {
	*x = CreatePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePageRequest) ProtoMessage() {}

func (x *CreatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePageRequest.ProtoReflect.Descriptor instead.
func (*CreatePageRequest) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePageRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePageRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreatePageRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreatePageRequest) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *CreatePageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type CreatePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *CreatePageResponse) Reset() {
	*x = CreatePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePageResponse) ProtoMessage() {}

func (x *CreatePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePageResponse.ProtoReflect.Descriptor instead.
func (*CreatePageResponse) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePageResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type UpdatePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author  string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdatePageRequest) Reset() {
	*x = UpdatePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageRequest) ProtoMessage() {}

func (x *UpdatePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageRequest.ProtoReflect.Descriptor instead.
func (*UpdatePageRequest) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePageRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UpdatePageRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePageRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *UpdatePageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdatePageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *UpdatePageResponse) Reset() {
	*x = UpdatePageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePageResponse) ProtoMessage() {}

func (x *UpdatePageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePageResponse.ProtoReflect.Descriptor instead.
func (*UpdatePageResponse) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePageResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type GetPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetPageRequest) Reset() {
	*x = GetPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageRequest) ProtoMessage() {}

func (x *GetPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageRequest.ProtoReflect.Descriptor instead.
func (*GetPageRequest) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{4}
}

func (x *GetPageRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetPageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetPageResponse) Reset() {
	*x = GetPageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageResponse) ProtoMessage() {}

func (x *GetPageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageResponse.ProtoReflect.Descriptor instead.
func (*GetPageResponse) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{5}
}

func (x *GetPageResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type GetPageListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetPageListRequest) Reset() {
	*x = GetPageListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageListRequest) ProtoMessage() {}

func (x *GetPageListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageListRequest.ProtoReflect.Descriptor instead.
func (*GetPageListRequest) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{6}
}

func (x *GetPageListRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *GetPageListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetPageListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetPageListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32       `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Pages []*PageList `protobuf:"bytes,2,rep,name=pages,proto3" json:"pages,omitempty"`
}

func (x *GetPageListResponse) Reset() {
	*x = GetPageListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPageListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPageListResponse) ProtoMessage() {}

func (x *GetPageListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPageListResponse.ProtoReflect.Descriptor instead.
func (*GetPageListResponse) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{7}
}

func (x *GetPageListResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetPageListResponse) GetPages() []*PageList {
	if x != nil {
		return x.Pages
	}
	return nil
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the page.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Title of the page.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Description of the page.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Name of the author, displayed below the title.
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	// Optional. Image URL of the page.
	Thumbnail string `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// Number of page views for the page.
	Views int32 `protobuf:"varint,6,opt,name=views,proto3" json:"views,omitempty"`
	// Optional. Content of the page.
	Content string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{8}
}

func (x *Page) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Page) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Page) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Page) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Page) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *Page) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *Page) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the page.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// Title of the page.
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	// Description of the page.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Optional. Name of the author, displayed below the title.
	Author string `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	// Optional. Image URL of the page.
	Thumbnail string `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	// Number of page views for the page.
	Views int32 `protobuf:"varint,6,opt,name=views,proto3" json:"views,omitempty"`
}

func (x *PageList) Reset() {
	*x = PageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_page_v1_page_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageList) ProtoMessage() {}

func (x *PageList) ProtoReflect() protoreflect.Message {
	mi := &file_page_v1_page_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageList.ProtoReflect.Descriptor instead.
func (*PageList) Descriptor() ([]byte, []int) {
	return file_page_v1_page_proto_rawDescGZIP(), []int{9}
}

func (x *PageList) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *PageList) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PageList) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PageList) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PageList) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

func (x *PageList) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

var File_page_v1_page_proto protoreflect.FileDescriptor

var file_page_v1_page_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x9b, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x6f, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x22, 0xb8, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x69, 0x65, 0x77, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xa2,
	0x01, 0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x32, 0x83, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x69, 0x6e, 0x67, 0x6f, 0x66, 0x7a, 0x69,
	0x68, 0x75, 0x61, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_page_v1_page_proto_rawDescOnce sync.Once
	file_page_v1_page_proto_rawDescData = file_page_v1_page_proto_rawDesc
)

func file_page_v1_page_proto_rawDescGZIP() []byte {
	file_page_v1_page_proto_rawDescOnce.Do(func() {
		file_page_v1_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_page_v1_page_proto_rawDescData)
	})
	return file_page_v1_page_proto_rawDescData
}

var file_page_v1_page_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_page_v1_page_proto_goTypes = []interface{}{
	(*CreatePageRequest)(nil),   // 0: v1.CreatePageRequest
	(*CreatePageResponse)(nil),  // 1: v1.CreatePageResponse
	(*UpdatePageRequest)(nil),   // 2: v1.UpdatePageRequest
	(*UpdatePageResponse)(nil),  // 3: v1.UpdatePageResponse
	(*GetPageRequest)(nil),      // 4: v1.GetPageRequest
	(*GetPageResponse)(nil),     // 5: v1.GetPageResponse
	(*GetPageListRequest)(nil),  // 6: v1.GetPageListRequest
	(*GetPageListResponse)(nil), // 7: v1.GetPageListResponse
	(*Page)(nil),                // 8: v1.Page
	(*PageList)(nil),            // 9: v1.PageList
}
var file_page_v1_page_proto_depIdxs = []int32{
	8, // 0: v1.CreatePageResponse.page:type_name -> v1.Page
	8, // 1: v1.UpdatePageResponse.page:type_name -> v1.Page
	8, // 2: v1.GetPageResponse.page:type_name -> v1.Page
	9, // 3: v1.GetPageListResponse.pages:type_name -> v1.PageList
	0, // 4: v1.PageService.CreatePage:input_type -> v1.CreatePageRequest
	2, // 5: v1.PageService.UpdatePage:input_type -> v1.UpdatePageRequest
	4, // 6: v1.PageService.GetPage:input_type -> v1.GetPageRequest
	6, // 7: v1.PageService.GetPageList:input_type -> v1.GetPageListRequest
	1, // 8: v1.PageService.CreatePage:output_type -> v1.CreatePageResponse
	3, // 9: v1.PageService.UpdatePage:output_type -> v1.UpdatePageResponse
	5, // 10: v1.PageService.GetPage:output_type -> v1.GetPageResponse
	7, // 11: v1.PageService.GetPageList:output_type -> v1.GetPageListResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_page_v1_page_proto_init() }
func file_page_v1_page_proto_init() {
	if File_page_v1_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_page_v1_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePageRequest); i {
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
		file_page_v1_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePageResponse); i {
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
		file_page_v1_page_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePageRequest); i {
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
		file_page_v1_page_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePageResponse); i {
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
		file_page_v1_page_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageRequest); i {
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
		file_page_v1_page_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageResponse); i {
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
		file_page_v1_page_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageListRequest); i {
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
		file_page_v1_page_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPageListResponse); i {
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
		file_page_v1_page_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_page_v1_page_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageList); i {
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
			RawDescriptor: file_page_v1_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_page_v1_page_proto_goTypes,
		DependencyIndexes: file_page_v1_page_proto_depIdxs,
		MessageInfos:      file_page_v1_page_proto_msgTypes,
	}.Build()
	File_page_v1_page_proto = out.File
	file_page_v1_page_proto_rawDesc = nil
	file_page_v1_page_proto_goTypes = nil
	file_page_v1_page_proto_depIdxs = nil
}
