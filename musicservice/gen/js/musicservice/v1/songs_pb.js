// @generated by protoc-gen-es v1.2.0
// @generated from file musicservice/v1/songs.proto (package musicservice.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3 } from "@bufbuild/protobuf";
import { Song } from "./models_pb.js";

/**
 *
 * Song creation request object
 *
 * @generated from message musicservice.v1.CreateSongRequest
 */
export const CreateSongRequest = proto3.makeMessageType(
  "musicservice.v1.CreateSongRequest",
  () => [
    { no: 1, name: "song", kind: "message", T: Song },
  ],
);

/**
 *
 * Response of an song creation.
 *
 * @generated from message musicservice.v1.CreateSongResponse
 */
export const CreateSongResponse = proto3.makeMessageType(
  "musicservice.v1.CreateSongResponse",
  () => [
    { no: 1, name: "song", kind: "message", T: Song },
  ],
);

/**
 *
 * The request for (partially) updating an Song.
 *
 * @generated from message musicservice.v1.UpdateSongRequest
 */
export const UpdateSongRequest = proto3.makeMessageType(
  "musicservice.v1.UpdateSongRequest",
  () => [
    { no: 1, name: "song", kind: "message", T: Song },
    { no: 2, name: "update_mask", kind: "message", T: FieldMask },
  ],
);

/**
 *
 * The request for (partially) updating an Song.
 *
 * @generated from message musicservice.v1.UpdateSongResponse
 */
export const UpdateSongResponse = proto3.makeMessageType(
  "musicservice.v1.UpdateSongResponse",
  () => [
    { no: 1, name: "song", kind: "message", T: Song },
  ],
);

/**
 *
 * Batch gets multiple songs.
 *
 * @generated from message musicservice.v1.GetSongsRequest
 */
export const GetSongsRequest = proto3.makeMessageType(
  "musicservice.v1.GetSongsRequest",
  () => [
    { no: 1, name: "ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 *
 * Response for an song batch get.
 *
 * @generated from message musicservice.v1.GetSongsResponse
 */
export const GetSongsResponse = proto3.makeMessageType(
  "musicservice.v1.GetSongsResponse",
  () => [
    { no: 1, name: "songs", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "message", T: Song} },
  ],
);

/**
 *
 * An song search request.  For now only paginations params are provided.
 *
 * @generated from message musicservice.v1.ListSongsRequest
 */
export const ListSongsRequest = proto3.makeMessageType(
  "musicservice.v1.ListSongsRequest",
  () => [
    { no: 1, name: "page_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 *
 * Response of a song search/listing.
 *
 * @generated from message musicservice.v1.ListSongsResponse
 */
export const ListSongsResponse = proto3.makeMessageType(
  "musicservice.v1.ListSongsResponse",
  () => [
    { no: 1, name: "songs", kind: "message", T: Song, repeated: true },
    { no: 2, name: "next_page_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 *
 * Request to delete an song.
 *
 * @generated from message musicservice.v1.DeleteSongRequest
 */
export const DeleteSongRequest = proto3.makeMessageType(
  "musicservice.v1.DeleteSongRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 *
 * Song deletion response
 *
 * @generated from message musicservice.v1.DeleteSongResponse
 */
export const DeleteSongResponse = proto3.makeMessageType(
  "musicservice.v1.DeleteSongResponse",
  [],
);

