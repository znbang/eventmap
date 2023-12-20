// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file book/v1/book.proto (package book.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum book.v1.BookJobStatus
 */
export declare enum BookJobStatus {
  /**
   * @generated from enum value: BOOK_JOB_STATUS_DONE_UNSPECIFIED = 0;
   */
  DONE_UNSPECIFIED = 0,

  /**
   * @generated from enum value: BOOK_JOB_STATUS_PENDING = 1;
   */
  PENDING = 1,

  /**
   * @generated from enum value: BOOK_JOB_STATUS_RUNNING = 2;
   */
  RUNNING = 2,
}

/**
 * @generated from message book.v1.Book
 */
export declare class Book extends Message<Book> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string user_id = 2;
   */
  userId: string;

  /**
   * @generated from field: string title = 3;
   */
  title: string;

  /**
   * @generated from field: string author = 4;
   */
  author: string;

  /**
   * @generated from field: string url = 5;
   */
  url: string;

  /**
   * @generated from field: book.v1.BookJob job = 6;
   */
  job?: BookJob;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 7;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 8;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Book>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.Book";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Book;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Book;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Book;

  static equals(a: Book | PlainMessage<Book> | undefined, b: Book | PlainMessage<Book> | undefined): boolean;
}

/**
 * @generated from message book.v1.Chapter
 */
export declare class Chapter extends Message<Chapter> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string book_id = 2;
   */
  bookId: string;

  /**
   * @generated from field: int32 page = 3;
   */
  page: number;

  /**
   * @generated from field: string url = 4;
   */
  url: string;

  /**
   * @generated from field: string title = 5;
   */
  title: string;

  /**
   * @generated from field: string body = 6;
   */
  body: string;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 7;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 8;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Chapter>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.Chapter";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Chapter;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Chapter;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Chapter;

  static equals(a: Chapter | PlainMessage<Chapter> | undefined, b: Chapter | PlainMessage<Chapter> | undefined): boolean;
}

/**
 * @generated from message book.v1.BookJob
 */
export declare class BookJob extends Message<BookJob> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: book.v1.BookJobStatus status = 2;
   */
  status: BookJobStatus;

  /**
   * @generated from field: string book_id = 3;
   */
  bookId: string;

  /**
   * @generated from field: string message = 4;
   */
  message: string;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 5;
   */
  createdAt?: Timestamp;

  constructor(data?: PartialMessage<BookJob>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.BookJob";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): BookJob;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): BookJob;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): BookJob;

  static equals(a: BookJob | PlainMessage<BookJob> | undefined, b: BookJob | PlainMessage<BookJob> | undefined): boolean;
}

/**
 * @generated from message book.v1.CreateBookRequest
 */
export declare class CreateBookRequest extends Message<CreateBookRequest> {
  /**
   * @generated from field: string title = 1;
   */
  title: string;

  /**
   * @generated from field: string author = 2;
   */
  author: string;

  /**
   * @generated from field: string url = 3;
   */
  url: string;

  constructor(data?: PartialMessage<CreateBookRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.CreateBookRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateBookRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateBookRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateBookRequest;

  static equals(a: CreateBookRequest | PlainMessage<CreateBookRequest> | undefined, b: CreateBookRequest | PlainMessage<CreateBookRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.CreateBookResponse
 */
export declare class CreateBookResponse extends Message<CreateBookResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<CreateBookResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.CreateBookResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateBookResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateBookResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateBookResponse;

  static equals(a: CreateBookResponse | PlainMessage<CreateBookResponse> | undefined, b: CreateBookResponse | PlainMessage<CreateBookResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.UpdateBookRequest
 */
export declare class UpdateBookRequest extends Message<UpdateBookRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string title = 2;
   */
  title: string;

  /**
   * @generated from field: string author = 3;
   */
  author: string;

  /**
   * @generated from field: string url = 4;
   */
  url: string;

  constructor(data?: PartialMessage<UpdateBookRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.UpdateBookRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateBookRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateBookRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateBookRequest;

  static equals(a: UpdateBookRequest | PlainMessage<UpdateBookRequest> | undefined, b: UpdateBookRequest | PlainMessage<UpdateBookRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.UpdateBookResponse
 */
export declare class UpdateBookResponse extends Message<UpdateBookResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<UpdateBookResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.UpdateBookResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateBookResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateBookResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateBookResponse;

  static equals(a: UpdateBookResponse | PlainMessage<UpdateBookResponse> | undefined, b: UpdateBookResponse | PlainMessage<UpdateBookResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.DeleteBookRequest
 */
export declare class DeleteBookRequest extends Message<DeleteBookRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteBookRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.DeleteBookRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteBookRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteBookRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteBookRequest;

  static equals(a: DeleteBookRequest | PlainMessage<DeleteBookRequest> | undefined, b: DeleteBookRequest | PlainMessage<DeleteBookRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.DeleteBookResponse
 */
export declare class DeleteBookResponse extends Message<DeleteBookResponse> {
  constructor(data?: PartialMessage<DeleteBookResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.DeleteBookResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteBookResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteBookResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteBookResponse;

  static equals(a: DeleteBookResponse | PlainMessage<DeleteBookResponse> | undefined, b: DeleteBookResponse | PlainMessage<DeleteBookResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.DeleteChapterRequest
 */
export declare class DeleteChapterRequest extends Message<DeleteChapterRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: int32 page = 2;
   */
  page: number;

  constructor(data?: PartialMessage<DeleteChapterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.DeleteChapterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteChapterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteChapterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteChapterRequest;

  static equals(a: DeleteChapterRequest | PlainMessage<DeleteChapterRequest> | undefined, b: DeleteChapterRequest | PlainMessage<DeleteChapterRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.DeleteChapterResponse
 */
export declare class DeleteChapterResponse extends Message<DeleteChapterResponse> {
  constructor(data?: PartialMessage<DeleteChapterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.DeleteChapterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteChapterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteChapterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteChapterResponse;

  static equals(a: DeleteChapterResponse | PlainMessage<DeleteChapterResponse> | undefined, b: DeleteChapterResponse | PlainMessage<DeleteChapterResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.GetBookRequest
 */
export declare class GetBookRequest extends Message<GetBookRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<GetBookRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.GetBookRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetBookRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetBookRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetBookRequest;

  static equals(a: GetBookRequest | PlainMessage<GetBookRequest> | undefined, b: GetBookRequest | PlainMessage<GetBookRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.GetBookResponse
 */
export declare class GetBookResponse extends Message<GetBookResponse> {
  /**
   * @generated from field: book.v1.Book book = 1;
   */
  book?: Book;

  constructor(data?: PartialMessage<GetBookResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.GetBookResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetBookResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetBookResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetBookResponse;

  static equals(a: GetBookResponse | PlainMessage<GetBookResponse> | undefined, b: GetBookResponse | PlainMessage<GetBookResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.GetTocRequest
 */
export declare class GetTocRequest extends Message<GetTocRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: int32 page = 2;
   */
  page: number;

  /**
   * @generated from field: int32 size = 3;
   */
  size: number;

  constructor(data?: PartialMessage<GetTocRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.GetTocRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTocRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTocRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTocRequest;

  static equals(a: GetTocRequest | PlainMessage<GetTocRequest> | undefined, b: GetTocRequest | PlainMessage<GetTocRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.GetTocResponse
 */
export declare class GetTocResponse extends Message<GetTocResponse> {
  /**
   * @generated from field: book.v1.Book book = 1;
   */
  book?: Book;

  /**
   * @generated from field: int32 page = 2;
   */
  page: number;

  /**
   * @generated from field: int32 size = 3;
   */
  size: number;

  /**
   * @generated from field: int32 total = 4;
   */
  total: number;

  /**
   * @generated from field: repeated book.v1.Chapter items = 5;
   */
  items: Chapter[];

  constructor(data?: PartialMessage<GetTocResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.GetTocResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetTocResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetTocResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetTocResponse;

  static equals(a: GetTocResponse | PlainMessage<GetTocResponse> | undefined, b: GetTocResponse | PlainMessage<GetTocResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.GetChapterRequest
 */
export declare class GetChapterRequest extends Message<GetChapterRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: int32 page = 2;
   */
  page: number;

  constructor(data?: PartialMessage<GetChapterRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.GetChapterRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetChapterRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetChapterRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetChapterRequest;

  static equals(a: GetChapterRequest | PlainMessage<GetChapterRequest> | undefined, b: GetChapterRequest | PlainMessage<GetChapterRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.GetChapterResponse
 */
export declare class GetChapterResponse extends Message<GetChapterResponse> {
  /**
   * @generated from field: book.v1.Book book = 1;
   */
  book?: Book;

  /**
   * @generated from field: book.v1.Chapter chapter = 2;
   */
  chapter?: Chapter;

  /**
   * @generated from field: int32 page = 3;
   */
  page: number;

  /**
   * @generated from field: int32 total = 4;
   */
  total: number;

  constructor(data?: PartialMessage<GetChapterResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.GetChapterResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetChapterResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetChapterResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetChapterResponse;

  static equals(a: GetChapterResponse | PlainMessage<GetChapterResponse> | undefined, b: GetChapterResponse | PlainMessage<GetChapterResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.ListBookRequest
 */
export declare class ListBookRequest extends Message<ListBookRequest> {
  /**
   * @generated from field: int32 page = 1;
   */
  page: number;

  /**
   * @generated from field: int32 size = 2;
   */
  size: number;

  /**
   * @generated from field: string filter = 3;
   */
  filter: string;

  constructor(data?: PartialMessage<ListBookRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.ListBookRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListBookRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListBookRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListBookRequest;

  static equals(a: ListBookRequest | PlainMessage<ListBookRequest> | undefined, b: ListBookRequest | PlainMessage<ListBookRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.ListBookResponse
 */
export declare class ListBookResponse extends Message<ListBookResponse> {
  /**
   * @generated from field: int32 page = 1;
   */
  page: number;

  /**
   * @generated from field: int32 size = 2;
   */
  size: number;

  /**
   * @generated from field: int32 total = 3;
   */
  total: number;

  /**
   * @generated from field: repeated book.v1.Book items = 4;
   */
  items: Book[];

  constructor(data?: PartialMessage<ListBookResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.ListBookResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListBookResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListBookResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListBookResponse;

  static equals(a: ListBookResponse | PlainMessage<ListBookResponse> | undefined, b: ListBookResponse | PlainMessage<ListBookResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.SyncNewRequest
 */
export declare class SyncNewRequest extends Message<SyncNewRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<SyncNewRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.SyncNewRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SyncNewRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SyncNewRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SyncNewRequest;

  static equals(a: SyncNewRequest | PlainMessage<SyncNewRequest> | undefined, b: SyncNewRequest | PlainMessage<SyncNewRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.SyncNewResponse
 */
export declare class SyncNewResponse extends Message<SyncNewResponse> {
  constructor(data?: PartialMessage<SyncNewResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.SyncNewResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SyncNewResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SyncNewResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SyncNewResponse;

  static equals(a: SyncNewResponse | PlainMessage<SyncNewResponse> | undefined, b: SyncNewResponse | PlainMessage<SyncNewResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.SyncAllRequest
 */
export declare class SyncAllRequest extends Message<SyncAllRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<SyncAllRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.SyncAllRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SyncAllRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SyncAllRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SyncAllRequest;

  static equals(a: SyncAllRequest | PlainMessage<SyncAllRequest> | undefined, b: SyncAllRequest | PlainMessage<SyncAllRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.SyncAllResponse
 */
export declare class SyncAllResponse extends Message<SyncAllResponse> {
  constructor(data?: PartialMessage<SyncAllResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.SyncAllResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SyncAllResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SyncAllResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SyncAllResponse;

  static equals(a: SyncAllResponse | PlainMessage<SyncAllResponse> | undefined, b: SyncAllResponse | PlainMessage<SyncAllResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.StopSyncRequest
 */
export declare class StopSyncRequest extends Message<StopSyncRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<StopSyncRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.StopSyncRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StopSyncRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StopSyncRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StopSyncRequest;

  static equals(a: StopSyncRequest | PlainMessage<StopSyncRequest> | undefined, b: StopSyncRequest | PlainMessage<StopSyncRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.StopSyncResponse
 */
export declare class StopSyncResponse extends Message<StopSyncResponse> {
  constructor(data?: PartialMessage<StopSyncResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.StopSyncResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): StopSyncResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): StopSyncResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): StopSyncResponse;

  static equals(a: StopSyncResponse | PlainMessage<StopSyncResponse> | undefined, b: StopSyncResponse | PlainMessage<StopSyncResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.SyncStatusRequest
 */
export declare class SyncStatusRequest extends Message<SyncStatusRequest> {
  constructor(data?: PartialMessage<SyncStatusRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.SyncStatusRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SyncStatusRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SyncStatusRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SyncStatusRequest;

  static equals(a: SyncStatusRequest | PlainMessage<SyncStatusRequest> | undefined, b: SyncStatusRequest | PlainMessage<SyncStatusRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.SyncStatusResponse
 */
export declare class SyncStatusResponse extends Message<SyncStatusResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string message = 2;
   */
  message: string;

  /**
   * @generated from field: book.v1.BookJobStatus status = 3;
   */
  status: BookJobStatus;

  constructor(data?: PartialMessage<SyncStatusResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.SyncStatusResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SyncStatusResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SyncStatusResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SyncStatusResponse;

  static equals(a: SyncStatusResponse | PlainMessage<SyncStatusResponse> | undefined, b: SyncStatusResponse | PlainMessage<SyncStatusResponse> | undefined): boolean;
}

/**
 * @generated from message book.v1.DownloadBookRequest
 */
export declare class DownloadBookRequest extends Message<DownloadBookRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DownloadBookRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.DownloadBookRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DownloadBookRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DownloadBookRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DownloadBookRequest;

  static equals(a: DownloadBookRequest | PlainMessage<DownloadBookRequest> | undefined, b: DownloadBookRequest | PlainMessage<DownloadBookRequest> | undefined): boolean;
}

/**
 * @generated from message book.v1.DownloadBookResponse
 */
export declare class DownloadBookResponse extends Message<DownloadBookResponse> {
  /**
   * @generated from field: string title = 1;
   */
  title: string;

  /**
   * @generated from field: string content = 2;
   */
  content: string;

  constructor(data?: PartialMessage<DownloadBookResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "book.v1.DownloadBookResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DownloadBookResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DownloadBookResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DownloadBookResponse;

  static equals(a: DownloadBookResponse | PlainMessage<DownloadBookResponse> | undefined, b: DownloadBookResponse | PlainMessage<DownloadBookResponse> | undefined): boolean;
}

