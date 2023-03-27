// @generated by protoc-gen-es v1.1.1
// @generated from file musicservice/v1/models.proto (package musicservice, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * A song - eg Jingle Bells
 *
 * @generated from message musicservice.Song
 */
export declare class Song extends Message<Song> {
  /**
   * @generated from field: google.protobuf.Timestamp created_at = 1;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 2;
   */
  updatedAt?: Timestamp;

  /**
   * ID of this song
   *
   * @generated from field: string id = 3;
   */
  id: string;

  /**
   * Name for this song
   *
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * ID of the composer/artist for this Song
   *
   * @generated from field: string composer_id = 5;
   */
  composerId: string;

  /**
   * A homepage containing all other info about the Song
   *
   * @generated from field: string homepage = 6;
   */
  homepage: string;

  /**
   * All performing artists for this song
   *
   * @generated from field: repeated string performer_ids = 7;
   */
  performerIds: string[];

  constructor(data?: PartialMessage<Song>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.Song";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Song;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Song;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Song;

  static equals(a: Song | PlainMessage<Song> | undefined, b: Song | PlainMessage<Song> | undefined): boolean;
}

/**
 * Artists perform/play/sing songs
 *
 * @generated from message musicservice.Artist
 */
export declare class Artist extends Message<Artist> {
  /**
   * @generated from field: google.protobuf.Timestamp created_at = 1;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 2;
   */
  updatedAt?: Timestamp;

  /**
   * ID of the artist
   *
   * @generated from field: string id = 3;
   */
  id: string;

  /**
   * Full name of the artist
   *
   * @generated from field: string name = 5;
   */
  name: string;

  /**
   * Artist's date of birth
   *
   * @generated from field: google.protobuf.Timestamp date_of_birth = 6;
   */
  dateOfBirth?: Timestamp;

  /**
   * Country of residence
   *
   * @generated from field: string country = 7;
   */
  country: string;

  constructor(data?: PartialMessage<Artist>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.Artist";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Artist;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Artist;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Artist;

  static equals(a: Artist | PlainMessage<Artist> | undefined, b: Artist | PlainMessage<Artist> | undefined): boolean;
}

/**
 * Album showcasing a bunch of artists performing a bunch of songs.
 *
 * @generated from message musicservice.Album
 */
export declare class Album extends Message<Album> {
  /**
   * @generated from field: google.protobuf.Timestamp created_at = 1;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 2;
   */
  updatedAt?: Timestamp;

  /**
   * ID of the Album
   *
   * @generated from field: string id = 3;
   */
  id: string;

  /**
   * Album name
   *
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * Songs performed in this album
   *
   * @generated from field: repeated string song_ids = 5;
   */
  songIds: string[];

  /**
   * Where this album was performed
   *
   * @generated from field: string venue = 6;
   */
  venue: string;

  /**
   * Date released on
   *
   * @generated from field: google.protobuf.Timestamp release_date = 7;
   */
  releaseDate?: Timestamp;

  constructor(data?: PartialMessage<Album>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.Album";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Album;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Album;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Album;

  static equals(a: Album | PlainMessage<Album> | undefined, b: Album | PlainMessage<Album> | undefined): boolean;
}

/**
 * A record label company.   This entity is the owner of albums as they produce and release them.
 * Note unlike the Album entity the Label does not contain a (repeated) field of all albums it
 * owns.   This is because the way the albums are added to or removed from a Label's collection
 * is via the API.
 *
 * @generated from message musicservice.Label
 */
export declare class Label extends Message<Label> {
  /**
   * @generated from field: google.protobuf.Timestamp created_at = 1;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 2;
   */
  updatedAt?: Timestamp;

  /**
   * ID of this label
   *
   * @generated from field: string id = 3;
   */
  id: string;

  /**
   * Name of this record label company (eg, "Virgin Records")
   *
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * When this company was established
   *
   * @generated from field: google.protobuf.Timestamp established_date = 5;
   */
  establishedDate?: Timestamp;

  constructor(data?: PartialMessage<Label>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "musicservice.Label";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Label;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Label;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Label;

  static equals(a: Label | PlainMessage<Label> | undefined, b: Label | PlainMessage<Label> | undefined): boolean;
}

