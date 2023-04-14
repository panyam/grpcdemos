// @generated by protoc-gen-es v1.2.0
// @generated from file musicservice/v1/labels.proto (package musicservice.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, FieldMask, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Album, Label } from "./models_pb.js";

/**
 *
 * Label creation request object
 *
 * @generated from message musicservice.v1.CreateLabelRequest
 */
export declare class CreateLabelRequest extends Message<CreateLabelRequest> {
  /**
   *
   * Label being updated
   *
   * @generated from field: musicservice.v1.Label label = 1;
   */
  label?: Label;

  constructor(data?: PartialMessage<CreateLabelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.CreateLabelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateLabelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateLabelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateLabelRequest;

  static equals(a: CreateLabelRequest | PlainMessage<CreateLabelRequest> | undefined, b: CreateLabelRequest | PlainMessage<CreateLabelRequest> | undefined): boolean;
}

/**
 *
 * Response of an label creation.
 *
 * @generated from message musicservice.v1.CreateLabelResponse
 */
export declare class CreateLabelResponse extends Message<CreateLabelResponse> {
  /**
   *
   * Label being created
   *
   * @generated from field: musicservice.v1.Label label = 1;
   */
  label?: Label;

  constructor(data?: PartialMessage<CreateLabelResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.CreateLabelResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateLabelResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateLabelResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateLabelResponse;

  static equals(a: CreateLabelResponse | PlainMessage<CreateLabelResponse> | undefined, b: CreateLabelResponse | PlainMessage<CreateLabelResponse> | undefined): boolean;
}

/**
 *
 * The request for (partially) updating an Label.
 *
 * @generated from message musicservice.v1.UpdateLabelRequest
 */
export declare class UpdateLabelRequest extends Message<UpdateLabelRequest> {
  /**
   *
   * Label being updated
   *
   * @generated from field: musicservice.v1.Label label = 1;
   */
  label?: Label;

  /**
   *
   * Mask of fields being updated in this Label to make partial changes.
   *
   * @generated from field: google.protobuf.FieldMask update_mask = 2;
   */
  updateMask?: FieldMask;

  constructor(data?: PartialMessage<UpdateLabelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.UpdateLabelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateLabelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateLabelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateLabelRequest;

  static equals(a: UpdateLabelRequest | PlainMessage<UpdateLabelRequest> | undefined, b: UpdateLabelRequest | PlainMessage<UpdateLabelRequest> | undefined): boolean;
}

/**
 *
 * The request for (partially) updating an Label.
 *
 * @generated from message musicservice.v1.UpdateLabelResponse
 */
export declare class UpdateLabelResponse extends Message<UpdateLabelResponse> {
  /**
   *
   * Label being updated
   *
   * @generated from field: musicservice.v1.Label label = 1;
   */
  label?: Label;

  constructor(data?: PartialMessage<UpdateLabelResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.UpdateLabelResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateLabelResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateLabelResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateLabelResponse;

  static equals(a: UpdateLabelResponse | PlainMessage<UpdateLabelResponse> | undefined, b: UpdateLabelResponse | PlainMessage<UpdateLabelResponse> | undefined): boolean;
}

/**
 *
 * Batch gets multiple labels.
 *
 * @generated from message musicservice.v1.GetLabelsRequest
 */
export declare class GetLabelsRequest extends Message<GetLabelsRequest> {
  /**
   * @generated from field: repeated string ids = 1;
   */
  ids: string[];

  constructor(data?: PartialMessage<GetLabelsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.GetLabelsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLabelsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLabelsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLabelsRequest;

  static equals(a: GetLabelsRequest | PlainMessage<GetLabelsRequest> | undefined, b: GetLabelsRequest | PlainMessage<GetLabelsRequest> | undefined): boolean;
}

/**
 *
 * Response for an label batch get.
 *
 * @generated from message musicservice.v1.GetLabelsResponse
 */
export declare class GetLabelsResponse extends Message<GetLabelsResponse> {
  /**
   *
   * Labels mapped by their ID.
   *
   * @generated from field: map<string, musicservice.v1.Label> labels = 1;
   */
  labels: { [key: string]: Label };

  constructor(data?: PartialMessage<GetLabelsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.GetLabelsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetLabelsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetLabelsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetLabelsResponse;

  static equals(a: GetLabelsResponse | PlainMessage<GetLabelsResponse> | undefined, b: GetLabelsResponse | PlainMessage<GetLabelsResponse> | undefined): boolean;
}

/**
 *
 * An label search request.  For now only paginations params are provided.
 *
 * @generated from message musicservice.v1.ListLabelsRequest
 */
export declare class ListLabelsRequest extends Message<ListLabelsRequest> {
  /**
   *
   * Instead of an offset an abstract  "page" key is provided that offers
   * an opaque "pointer" into some offset in a result set.
   *
   * @generated from field: string page_key = 1;
   */
  pageKey: string;

  /**
   *
   * Number of results to return.
   *
   * @generated from field: int32 count = 2;
   */
  count: number;

  constructor(data?: PartialMessage<ListLabelsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.ListLabelsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListLabelsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListLabelsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListLabelsRequest;

  static equals(a: ListLabelsRequest | PlainMessage<ListLabelsRequest> | undefined, b: ListLabelsRequest | PlainMessage<ListLabelsRequest> | undefined): boolean;
}

/**
 *
 * Response of a label search/listing.
 *
 * @generated from message musicservice.v1.ListLabelsResponse
 */
export declare class ListLabelsResponse extends Message<ListLabelsResponse> {
  /**
   *
   * The list of labels found as part of this response.
   *
   * @generated from field: repeated musicservice.v1.Label labels = 1;
   */
  labels: Label[];

  /**
   *
   * The key/pointer string that subsequent List requests should pass to
   * continue the pagination.
   *
   * @generated from field: string next_page_key = 2;
   */
  nextPageKey: string;

  constructor(data?: PartialMessage<ListLabelsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.ListLabelsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListLabelsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListLabelsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListLabelsResponse;

  static equals(a: ListLabelsResponse | PlainMessage<ListLabelsResponse> | undefined, b: ListLabelsResponse | PlainMessage<ListLabelsResponse> | undefined): boolean;
}

/**
 *
 * Request to delete an label.
 *
 * @generated from message musicservice.v1.DeleteLabelRequest
 */
export declare class DeleteLabelRequest extends Message<DeleteLabelRequest> {
  /**
   *
   * ID of the label to be deleted.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteLabelRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.DeleteLabelRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteLabelRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteLabelRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteLabelRequest;

  static equals(a: DeleteLabelRequest | PlainMessage<DeleteLabelRequest> | undefined, b: DeleteLabelRequest | PlainMessage<DeleteLabelRequest> | undefined): boolean;
}

/**
 *
 * Label deletion response
 *
 * @generated from message musicservice.v1.DeleteLabelResponse
 */
export declare class DeleteLabelResponse extends Message<DeleteLabelResponse> {
  constructor(data?: PartialMessage<DeleteLabelResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.DeleteLabelResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteLabelResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteLabelResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteLabelResponse;

  static equals(a: DeleteLabelResponse | PlainMessage<DeleteLabelResponse> | undefined, b: DeleteLabelResponse | PlainMessage<DeleteLabelResponse> | undefined): boolean;
}

/**
 *
 * Request to add an album to a label.
 *
 * @generated from message musicservice.v1.AddAlbumRequest
 */
export declare class AddAlbumRequest extends Message<AddAlbumRequest> {
  /**
   *
   * Label to add the album to.
   *
   * @generated from field: string label_id = 1;
   */
  labelId: string;

  /**
   *
   * The album to be added.
   *
   * @generated from field: string album_id = 2;
   */
  albumId: string;

  constructor(data?: PartialMessage<AddAlbumRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.AddAlbumRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddAlbumRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddAlbumRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddAlbumRequest;

  static equals(a: AddAlbumRequest | PlainMessage<AddAlbumRequest> | undefined, b: AddAlbumRequest | PlainMessage<AddAlbumRequest> | undefined): boolean;
}

/**
 *
 * Response for an AddAlbum request
 *
 * @generated from message musicservice.v1.AddAlbumResponse
 */
export declare class AddAlbumResponse extends Message<AddAlbumResponse> {
  constructor(data?: PartialMessage<AddAlbumResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.AddAlbumResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddAlbumResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddAlbumResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddAlbumResponse;

  static equals(a: AddAlbumResponse | PlainMessage<AddAlbumResponse> | undefined, b: AddAlbumResponse | PlainMessage<AddAlbumResponse> | undefined): boolean;
}

/**
 *
 * Request to remove an album to a label.
 *
 * @generated from message musicservice.v1.RemoveAlbumRequest
 */
export declare class RemoveAlbumRequest extends Message<RemoveAlbumRequest> {
  /**
   *
   * Label to remove the album to.
   *
   * @generated from field: string label_id = 1;
   */
  labelId: string;

  /**
   *
   * The album to be removeed.
   *
   * @generated from field: string album_id = 2;
   */
  albumId: string;

  constructor(data?: PartialMessage<RemoveAlbumRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.RemoveAlbumRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveAlbumRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveAlbumRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveAlbumRequest;

  static equals(a: RemoveAlbumRequest | PlainMessage<RemoveAlbumRequest> | undefined, b: RemoveAlbumRequest | PlainMessage<RemoveAlbumRequest> | undefined): boolean;
}

/**
 *
 * Response for an RemoveAlbum request
 *
 * @generated from message musicservice.v1.RemoveAlbumResponse
 */
export declare class RemoveAlbumResponse extends Message<RemoveAlbumResponse> {
  constructor(data?: PartialMessage<RemoveAlbumResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.RemoveAlbumResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RemoveAlbumResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RemoveAlbumResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RemoveAlbumResponse;

  static equals(a: RemoveAlbumResponse | PlainMessage<RemoveAlbumResponse> | undefined, b: RemoveAlbumResponse | PlainMessage<RemoveAlbumResponse> | undefined): boolean;
}

/**
 *
 * Request object for an album listing.
 *
 * @generated from message musicservice.v1.LabelServiceListAlbumsRequest
 */
export declare class LabelServiceListAlbumsRequest extends Message<LabelServiceListAlbumsRequest> {
  /**
   *
   * Instead of an offset an abstract  "page" key is provided that offers
   * an opaque "pointer" into some offset in a result set.
   *
   * @generated from field: string page_key = 1;
   */
  pageKey: string;

  /**
   *
   * Number of results to return.
   *
   * @generated from field: int32 count = 2;
   */
  count: number;

  constructor(data?: PartialMessage<LabelServiceListAlbumsRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.LabelServiceListAlbumsRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LabelServiceListAlbumsRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LabelServiceListAlbumsRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LabelServiceListAlbumsRequest;

  static equals(a: LabelServiceListAlbumsRequest | PlainMessage<LabelServiceListAlbumsRequest> | undefined, b: LabelServiceListAlbumsRequest | PlainMessage<LabelServiceListAlbumsRequest> | undefined): boolean;
}

/**
 *
 * Response for listing albums in a label.
 *
 * @generated from message musicservice.v1.LabelServiceListAlbumsResponse
 */
export declare class LabelServiceListAlbumsResponse extends Message<LabelServiceListAlbumsResponse> {
  /**
   *
   * List of albums.
   *
   * @generated from field: repeated musicservice.v1.Album albums = 1;
   */
  albums: Album[];

  /**
   *
   * The key/pointer string that subsequent List requests should pass to
   * continue the pagination.
   *
   * @generated from field: string next_page_key = 2;
   */
  nextPageKey: string;

  constructor(data?: PartialMessage<LabelServiceListAlbumsResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.v1.LabelServiceListAlbumsResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LabelServiceListAlbumsResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LabelServiceListAlbumsResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LabelServiceListAlbumsResponse;

  static equals(a: LabelServiceListAlbumsResponse | PlainMessage<LabelServiceListAlbumsResponse> | undefined, b: LabelServiceListAlbumsResponse | PlainMessage<LabelServiceListAlbumsResponse> | undefined): boolean;
}

