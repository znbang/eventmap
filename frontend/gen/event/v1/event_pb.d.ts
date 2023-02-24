// @generated by protoc-gen-es v1.0.0 with parameter "target=js+dts"
// @generated from file event/v1/event.proto (package event.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message event.v1.Event
 */
export declare class Event extends Message<Event> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string user_id = 2;
   */
  userId: string;

  /**
   * @generated from field: google.protobuf.Timestamp start_date = 3;
   */
  startDate?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end_date = 4;
   */
  endDate?: Timestamp;

  /**
   * @generated from field: string name = 5;
   */
  name: string;

  /**
   * @generated from field: string location = 6;
   */
  location: string;

  /**
   * @generated from field: string url = 7;
   */
  url: string;

  /**
   * @generated from field: string detail = 8;
   */
  detail: string;

  /**
   * @generated from field: double lat = 9;
   */
  lat: number;

  /**
   * @generated from field: double lng = 10;
   */
  lng: number;

  /**
   * @generated from field: int32 zoom = 11;
   */
  zoom: number;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 12;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 13;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Event>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.Event";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Event;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Event;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Event;

  static equals(a: Event | PlainMessage<Event> | undefined, b: Event | PlainMessage<Event> | undefined): boolean;
}

/**
 * @generated from message event.v1.CreateEventRequest
 */
export declare class CreateEventRequest extends Message<CreateEventRequest> {
  /**
   * @generated from field: google.protobuf.Timestamp start_date = 1;
   */
  startDate?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end_date = 2;
   */
  endDate?: Timestamp;

  /**
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * @generated from field: string location = 4;
   */
  location: string;

  /**
   * @generated from field: string url = 5;
   */
  url: string;

  /**
   * @generated from field: string detail = 6;
   */
  detail: string;

  /**
   * @generated from field: double lat = 7;
   */
  lat: number;

  /**
   * @generated from field: double lng = 8;
   */
  lng: number;

  /**
   * @generated from field: int32 zoom = 9;
   */
  zoom: number;

  constructor(data?: PartialMessage<CreateEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.CreateEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateEventRequest;

  static equals(a: CreateEventRequest | PlainMessage<CreateEventRequest> | undefined, b: CreateEventRequest | PlainMessage<CreateEventRequest> | undefined): boolean;
}

/**
 * @generated from message event.v1.CreateEventResponse
 */
export declare class CreateEventResponse extends Message<CreateEventResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<CreateEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.CreateEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateEventResponse;

  static equals(a: CreateEventResponse | PlainMessage<CreateEventResponse> | undefined, b: CreateEventResponse | PlainMessage<CreateEventResponse> | undefined): boolean;
}

/**
 * @generated from message event.v1.UpdateEventRequest
 */
export declare class UpdateEventRequest extends Message<UpdateEventRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: google.protobuf.Timestamp start_date = 2;
   */
  startDate?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp end_date = 3;
   */
  endDate?: Timestamp;

  /**
   * @generated from field: string name = 4;
   */
  name: string;

  /**
   * @generated from field: string location = 5;
   */
  location: string;

  /**
   * @generated from field: string url = 6;
   */
  url: string;

  /**
   * @generated from field: string detail = 7;
   */
  detail: string;

  /**
   * @generated from field: double lat = 8;
   */
  lat: number;

  /**
   * @generated from field: double lng = 9;
   */
  lng: number;

  /**
   * @generated from field: int32 zoom = 10;
   */
  zoom: number;

  constructor(data?: PartialMessage<UpdateEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.UpdateEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateEventRequest;

  static equals(a: UpdateEventRequest | PlainMessage<UpdateEventRequest> | undefined, b: UpdateEventRequest | PlainMessage<UpdateEventRequest> | undefined): boolean;
}

/**
 * @generated from message event.v1.UpdateEventResponse
 */
export declare class UpdateEventResponse extends Message<UpdateEventResponse> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<UpdateEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.UpdateEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateEventResponse;

  static equals(a: UpdateEventResponse | PlainMessage<UpdateEventResponse> | undefined, b: UpdateEventResponse | PlainMessage<UpdateEventResponse> | undefined): boolean;
}

/**
 * @generated from message event.v1.DeleteEventRequest
 */
export declare class DeleteEventRequest extends Message<DeleteEventRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<DeleteEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.DeleteEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteEventRequest;

  static equals(a: DeleteEventRequest | PlainMessage<DeleteEventRequest> | undefined, b: DeleteEventRequest | PlainMessage<DeleteEventRequest> | undefined): boolean;
}

/**
 * @generated from message event.v1.DeleteEventResponse
 */
export declare class DeleteEventResponse extends Message<DeleteEventResponse> {
  constructor(data?: PartialMessage<DeleteEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.DeleteEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteEventResponse;

  static equals(a: DeleteEventResponse | PlainMessage<DeleteEventResponse> | undefined, b: DeleteEventResponse | PlainMessage<DeleteEventResponse> | undefined): boolean;
}

/**
 * @generated from message event.v1.GetEventRequest
 */
export declare class GetEventRequest extends Message<GetEventRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  constructor(data?: PartialMessage<GetEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.GetEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventRequest;

  static equals(a: GetEventRequest | PlainMessage<GetEventRequest> | undefined, b: GetEventRequest | PlainMessage<GetEventRequest> | undefined): boolean;
}

/**
 * @generated from message event.v1.GetEventResponse
 */
export declare class GetEventResponse extends Message<GetEventResponse> {
  /**
   * @generated from field: event.v1.Event event = 1;
   */
  event?: Event;

  constructor(data?: PartialMessage<GetEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.GetEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetEventResponse;

  static equals(a: GetEventResponse | PlainMessage<GetEventResponse> | undefined, b: GetEventResponse | PlainMessage<GetEventResponse> | undefined): boolean;
}

/**
 * @generated from message event.v1.ListEventRequest
 */
export declare class ListEventRequest extends Message<ListEventRequest> {
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

  constructor(data?: PartialMessage<ListEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.ListEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventRequest;

  static equals(a: ListEventRequest | PlainMessage<ListEventRequest> | undefined, b: ListEventRequest | PlainMessage<ListEventRequest> | undefined): boolean;
}

/**
 * @generated from message event.v1.ListEventResponse
 */
export declare class ListEventResponse extends Message<ListEventResponse> {
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
   * @generated from field: repeated event.v1.Event items = 4;
   */
  items: Event[];

  constructor(data?: PartialMessage<ListEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.ListEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListEventResponse;

  static equals(a: ListEventResponse | PlainMessage<ListEventResponse> | undefined, b: ListEventResponse | PlainMessage<ListEventResponse> | undefined): boolean;
}

/**
 * @generated from message event.v1.ListActiveEventRequest
 */
export declare class ListActiveEventRequest extends Message<ListActiveEventRequest> {
  constructor(data?: PartialMessage<ListActiveEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.ListActiveEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListActiveEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListActiveEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListActiveEventRequest;

  static equals(a: ListActiveEventRequest | PlainMessage<ListActiveEventRequest> | undefined, b: ListActiveEventRequest | PlainMessage<ListActiveEventRequest> | undefined): boolean;
}

/**
 * @generated from message event.v1.ListActiveEventResponse
 */
export declare class ListActiveEventResponse extends Message<ListActiveEventResponse> {
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
   * @generated from field: repeated event.v1.Event items = 4;
   */
  items: Event[];

  constructor(data?: PartialMessage<ListActiveEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.ListActiveEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListActiveEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListActiveEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListActiveEventResponse;

  static equals(a: ListActiveEventResponse | PlainMessage<ListActiveEventResponse> | undefined, b: ListActiveEventResponse | PlainMessage<ListActiveEventResponse> | undefined): boolean;
}

/**
 * @generated from message event.v1.ListUserEventRequest
 */
export declare class ListUserEventRequest extends Message<ListUserEventRequest> {
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

  constructor(data?: PartialMessage<ListUserEventRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.ListUserEventRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUserEventRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUserEventRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUserEventRequest;

  static equals(a: ListUserEventRequest | PlainMessage<ListUserEventRequest> | undefined, b: ListUserEventRequest | PlainMessage<ListUserEventRequest> | undefined): boolean;
}

/**
 * @generated from message event.v1.ListUserEventResponse
 */
export declare class ListUserEventResponse extends Message<ListUserEventResponse> {
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
   * @generated from field: repeated event.v1.Event items = 4;
   */
  items: Event[];

  constructor(data?: PartialMessage<ListUserEventResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "event.v1.ListUserEventResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUserEventResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUserEventResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUserEventResponse;

  static equals(a: ListUserEventResponse | PlainMessage<ListUserEventResponse> | undefined, b: ListUserEventResponse | PlainMessage<ListUserEventResponse> | undefined): boolean;
}

