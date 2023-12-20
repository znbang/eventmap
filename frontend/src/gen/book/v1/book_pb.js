// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file book/v1/book.proto (package book.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum book.v1.BookJobStatus
 */
export const BookJobStatus = proto3.makeEnum(
  "book.v1.BookJobStatus",
  [
    {no: 0, name: "BOOK_JOB_STATUS_DONE_UNSPECIFIED", localName: "DONE_UNSPECIFIED"},
    {no: 1, name: "BOOK_JOB_STATUS_PENDING", localName: "PENDING"},
    {no: 2, name: "BOOK_JOB_STATUS_RUNNING", localName: "RUNNING"},
  ],
);

/**
 * @generated from message book.v1.Book
 */
export const Book = proto3.makeMessageType(
  "book.v1.Book",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "author", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "job", kind: "message", T: BookJob },
    { no: 7, name: "created_at", kind: "message", T: Timestamp },
    { no: 8, name: "updated_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message book.v1.Chapter
 */
export const Chapter = proto3.makeMessageType(
  "book.v1.Chapter",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "book_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "body", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "created_at", kind: "message", T: Timestamp },
    { no: 8, name: "updated_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message book.v1.BookJob
 */
export const BookJob = proto3.makeMessageType(
  "book.v1.BookJob",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "status", kind: "enum", T: proto3.getEnumType(BookJobStatus) },
    { no: 3, name: "book_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "created_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message book.v1.CreateBookRequest
 */
export const CreateBookRequest = proto3.makeMessageType(
  "book.v1.CreateBookRequest",
  () => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "author", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.CreateBookResponse
 */
export const CreateBookResponse = proto3.makeMessageType(
  "book.v1.CreateBookResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.UpdateBookRequest
 */
export const UpdateBookRequest = proto3.makeMessageType(
  "book.v1.UpdateBookRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "author", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.UpdateBookResponse
 */
export const UpdateBookResponse = proto3.makeMessageType(
  "book.v1.UpdateBookResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.DeleteBookRequest
 */
export const DeleteBookRequest = proto3.makeMessageType(
  "book.v1.DeleteBookRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.DeleteBookResponse
 */
export const DeleteBookResponse = proto3.makeMessageType(
  "book.v1.DeleteBookResponse",
  [],
);

/**
 * @generated from message book.v1.DeleteChapterRequest
 */
export const DeleteChapterRequest = proto3.makeMessageType(
  "book.v1.DeleteChapterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message book.v1.DeleteChapterResponse
 */
export const DeleteChapterResponse = proto3.makeMessageType(
  "book.v1.DeleteChapterResponse",
  [],
);

/**
 * @generated from message book.v1.GetBookRequest
 */
export const GetBookRequest = proto3.makeMessageType(
  "book.v1.GetBookRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.GetBookResponse
 */
export const GetBookResponse = proto3.makeMessageType(
  "book.v1.GetBookResponse",
  () => [
    { no: 1, name: "book", kind: "message", T: Book },
  ],
);

/**
 * @generated from message book.v1.GetTocRequest
 */
export const GetTocRequest = proto3.makeMessageType(
  "book.v1.GetTocRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message book.v1.GetTocResponse
 */
export const GetTocResponse = proto3.makeMessageType(
  "book.v1.GetTocResponse",
  () => [
    { no: 1, name: "book", kind: "message", T: Book },
    { no: 2, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "total", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "items", kind: "message", T: Chapter, repeated: true },
  ],
);

/**
 * @generated from message book.v1.GetChapterRequest
 */
export const GetChapterRequest = proto3.makeMessageType(
  "book.v1.GetChapterRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message book.v1.GetChapterResponse
 */
export const GetChapterResponse = proto3.makeMessageType(
  "book.v1.GetChapterResponse",
  () => [
    { no: 1, name: "book", kind: "message", T: Book },
    { no: 2, name: "chapter", kind: "message", T: Chapter },
    { no: 3, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "total", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ],
);

/**
 * @generated from message book.v1.ListBookRequest
 */
export const ListBookRequest = proto3.makeMessageType(
  "book.v1.ListBookRequest",
  () => [
    { no: 1, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "filter", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.ListBookResponse
 */
export const ListBookResponse = proto3.makeMessageType(
  "book.v1.ListBookResponse",
  () => [
    { no: 1, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "total", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "items", kind: "message", T: Book, repeated: true },
  ],
);

/**
 * @generated from message book.v1.SyncNewRequest
 */
export const SyncNewRequest = proto3.makeMessageType(
  "book.v1.SyncNewRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.SyncNewResponse
 */
export const SyncNewResponse = proto3.makeMessageType(
  "book.v1.SyncNewResponse",
  [],
);

/**
 * @generated from message book.v1.SyncAllRequest
 */
export const SyncAllRequest = proto3.makeMessageType(
  "book.v1.SyncAllRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.SyncAllResponse
 */
export const SyncAllResponse = proto3.makeMessageType(
  "book.v1.SyncAllResponse",
  [],
);

/**
 * @generated from message book.v1.StopSyncRequest
 */
export const StopSyncRequest = proto3.makeMessageType(
  "book.v1.StopSyncRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.StopSyncResponse
 */
export const StopSyncResponse = proto3.makeMessageType(
  "book.v1.StopSyncResponse",
  [],
);

/**
 * @generated from message book.v1.SyncStatusRequest
 */
export const SyncStatusRequest = proto3.makeMessageType(
  "book.v1.SyncStatusRequest",
  [],
);

/**
 * @generated from message book.v1.SyncStatusResponse
 */
export const SyncStatusResponse = proto3.makeMessageType(
  "book.v1.SyncStatusResponse",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "enum", T: proto3.getEnumType(BookJobStatus) },
  ],
);

/**
 * @generated from message book.v1.DownloadBookRequest
 */
export const DownloadBookRequest = proto3.makeMessageType(
  "book.v1.DownloadBookRequest",
  () => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message book.v1.DownloadBookResponse
 */
export const DownloadBookResponse = proto3.makeMessageType(
  "book.v1.DownloadBookResponse",
  () => [
    { no: 1, name: "title", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "content", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

