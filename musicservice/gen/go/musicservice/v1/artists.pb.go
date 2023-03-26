// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: artists.proto

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

// *
// Artist creation request object
type CreateArtistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Artist being updated
	Artist *Artist `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
}

func (x *CreateArtistRequest) Reset() {
	*x = CreateArtistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArtistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArtistRequest) ProtoMessage() {}

func (x *CreateArtistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArtistRequest.ProtoReflect.Descriptor instead.
func (*CreateArtistRequest) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{0}
}

func (x *CreateArtistRequest) GetArtist() *Artist {
	if x != nil {
		return x.Artist
	}
	return nil
}

// *
// Response of an artist creation.
type CreateArtistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Artist being created
	Artist *Artist `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
}

func (x *CreateArtistResponse) Reset() {
	*x = CreateArtistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateArtistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArtistResponse) ProtoMessage() {}

func (x *CreateArtistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateArtistResponse.ProtoReflect.Descriptor instead.
func (*CreateArtistResponse) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArtistResponse) GetArtist() *Artist {
	if x != nil {
		return x.Artist
	}
	return nil
}

// *
// The request for (partially) updating an Artist.
type UpdateArtistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Artist being updated
	Artist *Artist `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
	// *
	// Mask of fields being updated in this Artist to make partial changes.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateArtistRequest) Reset() {
	*x = UpdateArtistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArtistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArtistRequest) ProtoMessage() {}

func (x *UpdateArtistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArtistRequest.ProtoReflect.Descriptor instead.
func (*UpdateArtistRequest) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateArtistRequest) GetArtist() *Artist {
	if x != nil {
		return x.Artist
	}
	return nil
}

func (x *UpdateArtistRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

// *
// The request for (partially) updating an Artist.
type UpdateArtistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Artist being updated
	Artist *Artist `protobuf:"bytes,1,opt,name=artist,proto3" json:"artist,omitempty"`
}

func (x *UpdateArtistResponse) Reset() {
	*x = UpdateArtistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateArtistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateArtistResponse) ProtoMessage() {}

func (x *UpdateArtistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateArtistResponse.ProtoReflect.Descriptor instead.
func (*UpdateArtistResponse) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateArtistResponse) GetArtist() *Artist {
	if x != nil {
		return x.Artist
	}
	return nil
}

// *
// Batch gets multiple artists.
type GetArtistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *GetArtistsRequest) Reset() {
	*x = GetArtistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArtistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArtistsRequest) ProtoMessage() {}

func (x *GetArtistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArtistsRequest.ProtoReflect.Descriptor instead.
func (*GetArtistsRequest) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{4}
}

func (x *GetArtistsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

// *
// Response for an artist batch get.
type GetArtistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Artists mapped by their ID.
	Artists map[string]*Artist `protobuf:"bytes,1,rep,name=artists,proto3" json:"artists,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetArtistsResponse) Reset() {
	*x = GetArtistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetArtistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArtistsResponse) ProtoMessage() {}

func (x *GetArtistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArtistsResponse.ProtoReflect.Descriptor instead.
func (*GetArtistsResponse) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{5}
}

func (x *GetArtistsResponse) GetArtists() map[string]*Artist {
	if x != nil {
		return x.Artists
	}
	return nil
}

// *
// An artist search request.  For now only paginations params are provided.
type ListArtistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// Instead of an offset an abstract  "page" key is provided that offers
	// an opaque "pointer" into some offset in a result set.
	PageKey string `protobuf:"bytes,1,opt,name=page_key,json=pageKey,proto3" json:"page_key,omitempty"`
	// *
	// Number of results to return.
	Count int32 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *ListArtistsRequest) Reset() {
	*x = ListArtistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArtistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtistsRequest) ProtoMessage() {}

func (x *ListArtistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtistsRequest.ProtoReflect.Descriptor instead.
func (*ListArtistsRequest) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{6}
}

func (x *ListArtistsRequest) GetPageKey() string {
	if x != nil {
		return x.PageKey
	}
	return ""
}

func (x *ListArtistsRequest) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

// *
// Response of a artist search/listing.
type ListArtistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// The list of artists found as part of this response.
	Artists []*Artist `protobuf:"bytes,1,rep,name=artists,proto3" json:"artists,omitempty"`
	// *
	// The key/pointer string that subsequent List requests should pass to
	// continue the pagination.
	NextPageKey string `protobuf:"bytes,2,opt,name=next_page_key,json=nextPageKey,proto3" json:"next_page_key,omitempty"`
}

func (x *ListArtistsResponse) Reset() {
	*x = ListArtistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListArtistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListArtistsResponse) ProtoMessage() {}

func (x *ListArtistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListArtistsResponse.ProtoReflect.Descriptor instead.
func (*ListArtistsResponse) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{7}
}

func (x *ListArtistsResponse) GetArtists() []*Artist {
	if x != nil {
		return x.Artists
	}
	return nil
}

func (x *ListArtistsResponse) GetNextPageKey() string {
	if x != nil {
		return x.NextPageKey
	}
	return ""
}

// *
// Request to delete an artist.
type DeleteArtistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// *
	// ID of the artist to be deleted.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteArtistRequest) Reset() {
	*x = DeleteArtistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArtistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArtistRequest) ProtoMessage() {}

func (x *DeleteArtistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArtistRequest.ProtoReflect.Descriptor instead.
func (*DeleteArtistRequest) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteArtistRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// *
// Artist deletion response
type DeleteArtistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteArtistResponse) Reset() {
	*x = DeleteArtistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artists_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteArtistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteArtistResponse) ProtoMessage() {}

func (x *DeleteArtistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_artists_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteArtistResponse.ProtoReflect.Descriptor instead.
func (*DeleteArtistResponse) Descriptor() ([]byte, []int) {
	return file_artists_proto_rawDescGZIP(), []int{9}
}

var File_artists_proto protoreflect.FileDescriptor

var file_artists_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x22, 0x44, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x06,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x22, 0x80, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c,
	0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x44, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x22,
	0x25, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xaf, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a,
	0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x1a, 0x50, 0x0a, 0x0c, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x45, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x69, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x25, 0x0a, 0x13, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xca, 0x04, 0x0a, 0x0d, 0x41, 0x72,
	0x74, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x6d, 0x75,
	0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x6d, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x75, 0x73, 0x69,
	0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73,
	0x3a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x12, 0x80, 0x01, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x6d, 0x75, 0x73,
	0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x32, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x2e, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x6f, 0x0a, 0x0c,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x67, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x6d,
	0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artists_proto_rawDescOnce sync.Once
	file_artists_proto_rawDescData = file_artists_proto_rawDesc
)

func file_artists_proto_rawDescGZIP() []byte {
	file_artists_proto_rawDescOnce.Do(func() {
		file_artists_proto_rawDescData = protoimpl.X.CompressGZIP(file_artists_proto_rawDescData)
	})
	return file_artists_proto_rawDescData
}

var file_artists_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_artists_proto_goTypes = []interface{}{
	(*CreateArtistRequest)(nil),   // 0: musicservice.CreateArtistRequest
	(*CreateArtistResponse)(nil),  // 1: musicservice.CreateArtistResponse
	(*UpdateArtistRequest)(nil),   // 2: musicservice.UpdateArtistRequest
	(*UpdateArtistResponse)(nil),  // 3: musicservice.UpdateArtistResponse
	(*GetArtistsRequest)(nil),     // 4: musicservice.GetArtistsRequest
	(*GetArtistsResponse)(nil),    // 5: musicservice.GetArtistsResponse
	(*ListArtistsRequest)(nil),    // 6: musicservice.ListArtistsRequest
	(*ListArtistsResponse)(nil),   // 7: musicservice.ListArtistsResponse
	(*DeleteArtistRequest)(nil),   // 8: musicservice.DeleteArtistRequest
	(*DeleteArtistResponse)(nil),  // 9: musicservice.DeleteArtistResponse
	nil,                           // 10: musicservice.GetArtistsResponse.ArtistsEntry
	(*Artist)(nil),                // 11: musicservice.Artist
	(*fieldmaskpb.FieldMask)(nil), // 12: google.protobuf.FieldMask
}
var file_artists_proto_depIdxs = []int32{
	11, // 0: musicservice.CreateArtistRequest.artist:type_name -> musicservice.Artist
	11, // 1: musicservice.CreateArtistResponse.artist:type_name -> musicservice.Artist
	11, // 2: musicservice.UpdateArtistRequest.artist:type_name -> musicservice.Artist
	12, // 3: musicservice.UpdateArtistRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 4: musicservice.UpdateArtistResponse.artist:type_name -> musicservice.Artist
	10, // 5: musicservice.GetArtistsResponse.artists:type_name -> musicservice.GetArtistsResponse.ArtistsEntry
	11, // 6: musicservice.ListArtistsResponse.artists:type_name -> musicservice.Artist
	11, // 7: musicservice.GetArtistsResponse.ArtistsEntry.value:type_name -> musicservice.Artist
	0,  // 8: musicservice.ArtistService.CreateArtist:input_type -> musicservice.CreateArtistRequest
	4,  // 9: musicservice.ArtistService.GetArtists:input_type -> musicservice.GetArtistsRequest
	2,  // 10: musicservice.ArtistService.UpdateArtist:input_type -> musicservice.UpdateArtistRequest
	8,  // 11: musicservice.ArtistService.DeleteArtist:input_type -> musicservice.DeleteArtistRequest
	6,  // 12: musicservice.ArtistService.ListArtists:input_type -> musicservice.ListArtistsRequest
	1,  // 13: musicservice.ArtistService.CreateArtist:output_type -> musicservice.CreateArtistResponse
	5,  // 14: musicservice.ArtistService.GetArtists:output_type -> musicservice.GetArtistsResponse
	3,  // 15: musicservice.ArtistService.UpdateArtist:output_type -> musicservice.UpdateArtistResponse
	9,  // 16: musicservice.ArtistService.DeleteArtist:output_type -> musicservice.DeleteArtistResponse
	7,  // 17: musicservice.ArtistService.ListArtists:output_type -> musicservice.ListArtistsResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_artists_proto_init() }
func file_artists_proto_init() {
	if File_artists_proto != nil {
		return
	}
	file_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_artists_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArtistRequest); i {
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
		file_artists_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateArtistResponse); i {
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
		file_artists_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArtistRequest); i {
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
		file_artists_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateArtistResponse); i {
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
		file_artists_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArtistsRequest); i {
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
		file_artists_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetArtistsResponse); i {
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
		file_artists_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArtistsRequest); i {
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
		file_artists_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListArtistsResponse); i {
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
		file_artists_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArtistRequest); i {
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
		file_artists_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteArtistResponse); i {
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
			RawDescriptor: file_artists_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_artists_proto_goTypes,
		DependencyIndexes: file_artists_proto_depIdxs,
		MessageInfos:      file_artists_proto_msgTypes,
	}.Build()
	File_artists_proto = out.File
	file_artists_proto_rawDesc = nil
	file_artists_proto_goTypes = nil
	file_artists_proto_depIdxs = nil
}
