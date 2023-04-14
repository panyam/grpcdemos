// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: musicservice/v1/albums.proto

package protos

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Album creation request object
type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Album being updated
	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAlbumRequest) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

// Response of an album creation.
type CreateAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Album being created
	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *CreateAlbumResponse) Reset() {
	*x = CreateAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumResponse) ProtoMessage() {}

func (x *CreateAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumResponse.ProtoReflect.Descriptor instead.
func (*CreateAlbumResponse) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlbumResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

// The request for (partially) updating an Album.
type UpdateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Album being updated
	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
	// Mask of fields being updated in this Album to make partial changes.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateAlbumRequest) Reset() {
	*x = UpdateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlbumRequest) ProtoMessage() {}

func (x *UpdateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlbumRequest.ProtoReflect.Descriptor instead.
func (*UpdateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAlbumRequest) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

func (x *UpdateAlbumRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// The request for (partially) updating an Album.
type UpdateAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Album being updated
	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *UpdateAlbumResponse) Reset() {
	*x = UpdateAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlbumResponse) ProtoMessage() {}

func (x *UpdateAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlbumResponse.ProtoReflect.Descriptor instead.
func (*UpdateAlbumResponse) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAlbumResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

// Batch gets multiple albums.
type GetAlbumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetAlbumsRequest) Reset() {
	*x = GetAlbumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumsRequest) ProtoMessage() {}

func (x *GetAlbumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumsRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumsRequest) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{4}
}

func (x *GetAlbumsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// Response for an album batch get.
type GetAlbumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Albums mapped by their ID.
	Albums map[string]*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetAlbumsResponse) Reset() {
	*x = GetAlbumsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumsResponse) ProtoMessage() {}

func (x *GetAlbumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumsResponse.ProtoReflect.Descriptor instead.
func (*GetAlbumsResponse) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{5}
}

func (x *GetAlbumsResponse) GetAlbums() map[string]*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

// An album search request.  For now only paginations params are provided.
type ListAlbumsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Instead of an offset an abstract  "page" key is provided that offers
	// an opaque "pointer" into some offset in a result set.
	PageKey string `protobuf:"bytes,1,opt,name=page_key,json=pageKey,proto3" json:"page_key,omitempty"`
	// Number of results to return.
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListAlbumsRequest) Reset() {
	*x = ListAlbumsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumsRequest) ProtoMessage() {}

func (x *ListAlbumsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumsRequest.ProtoReflect.Descriptor instead.
func (*ListAlbumsRequest) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{6}
}

func (x *ListAlbumsRequest) GetPageKey() string {
	if x != nil {
		return x.PageKey
	}
	return ""
}

func (x *ListAlbumsRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// Response of a album search/listing.
type ListAlbumsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of albums found as part of this response.
	Albums []*Album `protobuf:"bytes,1,rep,name=albums,proto3" json:"albums,omitempty"`
	// The key/pointer string that subsequent List requests should pass to
	// continue the pagination.
	NextPageKey string `protobuf:"bytes,2,opt,name=next_page_key,json=nextPageKey,proto3" json:"next_page_key,omitempty"`
}

func (x *ListAlbumsResponse) Reset() {
	*x = ListAlbumsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlbumsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlbumsResponse) ProtoMessage() {}

func (x *ListAlbumsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlbumsResponse.ProtoReflect.Descriptor instead.
func (*ListAlbumsResponse) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{7}
}

func (x *ListAlbumsResponse) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

func (x *ListAlbumsResponse) GetNextPageKey() string {
	if x != nil {
		return x.NextPageKey
	}
	return ""
}

// Request to delete an album.
type DeleteAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the album to be deleted.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAlbumRequest) Reset() {
	*x = DeleteAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumRequest) ProtoMessage() {}

func (x *DeleteAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlbumRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlbumRequest) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteAlbumRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Album deletion response
type DeleteAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAlbumResponse) Reset() {
	*x = DeleteAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_musicservice_v1_albums_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlbumResponse) ProtoMessage() {}

func (x *DeleteAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_musicservice_v1_albums_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlbumResponse.ProtoReflect.Descriptor instead.
func (*DeleteAlbumResponse) Descriptor() ([]byte, []int) {
	return file_musicservice_v1_albums_proto_rawDescGZIP(), []int{9}
}

var File_musicservice_v1_albums_proto protoreflect.FileDescriptor

var file_musicservice_v1_albums_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x22, 0x43, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x7f, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x43, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c,
	0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x24, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69,
	0x64, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73,
	0x1a, 0x51, 0x0a, 0x0b, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x44, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12,
	0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x4b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xd2, 0x04, 0x0a, 0x0c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x12, 0x23, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x62, 0x75,
	0x6d, 0x73, 0x12, 0x6f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12,
	0x21, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x3a, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x12, 0x23, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x32, 0x17, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x2f, 0x7b, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x2e, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x72, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x23, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x69, 0x0a, 0x0a, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x22, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x42, 0x9e, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xa2, 0x02, 0x03, 0x4d, 0x58,
	0x58, 0xaa, 0x02, 0x0f, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_musicservice_v1_albums_proto_rawDescOnce sync.Once
	file_musicservice_v1_albums_proto_rawDescData = file_musicservice_v1_albums_proto_rawDesc
)

func file_musicservice_v1_albums_proto_rawDescGZIP() []byte {
	file_musicservice_v1_albums_proto_rawDescOnce.Do(func() {
		file_musicservice_v1_albums_proto_rawDescData = protoimpl.X.CompressGZIP(file_musicservice_v1_albums_proto_rawDescData)
	})
	return file_musicservice_v1_albums_proto_rawDescData
}

var file_musicservice_v1_albums_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_musicservice_v1_albums_proto_goTypes = []interface{}{
	(*CreateAlbumRequest)(nil),    // 0: musicservice.v1.CreateAlbumRequest
	(*CreateAlbumResponse)(nil),   // 1: musicservice.v1.CreateAlbumResponse
	(*UpdateAlbumRequest)(nil),    // 2: musicservice.v1.UpdateAlbumRequest
	(*UpdateAlbumResponse)(nil),   // 3: musicservice.v1.UpdateAlbumResponse
	(*GetAlbumsRequest)(nil),      // 4: musicservice.v1.GetAlbumsRequest
	(*GetAlbumsResponse)(nil),     // 5: musicservice.v1.GetAlbumsResponse
	(*ListAlbumsRequest)(nil),     // 6: musicservice.v1.ListAlbumsRequest
	(*ListAlbumsResponse)(nil),    // 7: musicservice.v1.ListAlbumsResponse
	(*DeleteAlbumRequest)(nil),    // 8: musicservice.v1.DeleteAlbumRequest
	(*DeleteAlbumResponse)(nil),   // 9: musicservice.v1.DeleteAlbumResponse
	nil,                           // 10: musicservice.v1.GetAlbumsResponse.AlbumsEntry
	(*Album)(nil),                 // 11: musicservice.v1.Album
	(*fieldmaskpb.FieldMask)(nil), // 12: google.protobuf.FieldMask
}
var file_musicservice_v1_albums_proto_depIdxs = []int32{
	11, // 0: musicservice.v1.CreateAlbumRequest.album:type_name -> musicservice.v1.Album
	11, // 1: musicservice.v1.CreateAlbumResponse.album:type_name -> musicservice.v1.Album
	11, // 2: musicservice.v1.UpdateAlbumRequest.album:type_name -> musicservice.v1.Album
	12, // 3: musicservice.v1.UpdateAlbumRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 4: musicservice.v1.UpdateAlbumResponse.album:type_name -> musicservice.v1.Album
	10, // 5: musicservice.v1.GetAlbumsResponse.albums:type_name -> musicservice.v1.GetAlbumsResponse.AlbumsEntry
	11, // 6: musicservice.v1.ListAlbumsResponse.albums:type_name -> musicservice.v1.Album
	11, // 7: musicservice.v1.GetAlbumsResponse.AlbumsEntry.value:type_name -> musicservice.v1.Album
	0,  // 8: musicservice.v1.AlbumService.CreateAlbum:input_type -> musicservice.v1.CreateAlbumRequest
	4,  // 9: musicservice.v1.AlbumService.GetAlbums:input_type -> musicservice.v1.GetAlbumsRequest
	2,  // 10: musicservice.v1.AlbumService.UpdateAlbum:input_type -> musicservice.v1.UpdateAlbumRequest
	8,  // 11: musicservice.v1.AlbumService.DeleteAlbum:input_type -> musicservice.v1.DeleteAlbumRequest
	6,  // 12: musicservice.v1.AlbumService.ListAlbums:input_type -> musicservice.v1.ListAlbumsRequest
	1,  // 13: musicservice.v1.AlbumService.CreateAlbum:output_type -> musicservice.v1.CreateAlbumResponse
	5,  // 14: musicservice.v1.AlbumService.GetAlbums:output_type -> musicservice.v1.GetAlbumsResponse
	3,  // 15: musicservice.v1.AlbumService.UpdateAlbum:output_type -> musicservice.v1.UpdateAlbumResponse
	9,  // 16: musicservice.v1.AlbumService.DeleteAlbum:output_type -> musicservice.v1.DeleteAlbumResponse
	7,  // 17: musicservice.v1.AlbumService.ListAlbums:output_type -> musicservice.v1.ListAlbumsResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_musicservice_v1_albums_proto_init() }
func file_musicservice_v1_albums_proto_init() {
	if File_musicservice_v1_albums_proto != nil {
		return
	}
	file_musicservice_v1_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_musicservice_v1_albums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumRequest); i {
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
		file_musicservice_v1_albums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumResponse); i {
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
		file_musicservice_v1_albums_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlbumRequest); i {
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
		file_musicservice_v1_albums_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlbumResponse); i {
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
		file_musicservice_v1_albums_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumsRequest); i {
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
		file_musicservice_v1_albums_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumsResponse); i {
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
		file_musicservice_v1_albums_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumsRequest); i {
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
		file_musicservice_v1_albums_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlbumsResponse); i {
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
		file_musicservice_v1_albums_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlbumRequest); i {
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
		file_musicservice_v1_albums_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlbumResponse); i {
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
			RawDescriptor: file_musicservice_v1_albums_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_musicservice_v1_albums_proto_goTypes,
		DependencyIndexes: file_musicservice_v1_albums_proto_depIdxs,
		MessageInfos:      file_musicservice_v1_albums_proto_msgTypes,
	}.Build()
	File_musicservice_v1_albums_proto = out.File
	file_musicservice_v1_albums_proto_rawDesc = nil
	file_musicservice_v1_albums_proto_goTypes = nil
	file_musicservice_v1_albums_proto_depIdxs = nil
}
