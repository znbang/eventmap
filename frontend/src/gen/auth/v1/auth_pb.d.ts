// @generated by protoc-gen-es v1.10.0 with parameter "target=js+dts"
// @generated from file auth/v1/auth.proto (package auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message auth.v1.Provider
 */
export declare class Provider extends Message<Provider> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string url = 2;
   */
  url: string;

  constructor(data?: PartialMessage<Provider>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "auth.v1.Provider";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Provider;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Provider;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Provider;

  static equals(a: Provider | PlainMessage<Provider> | undefined, b: Provider | PlainMessage<Provider> | undefined): boolean;
}

/**
 * @generated from message auth.v1.ListProviderRequest
 */
export declare class ListProviderRequest extends Message<ListProviderRequest> {
  constructor(data?: PartialMessage<ListProviderRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "auth.v1.ListProviderRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListProviderRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListProviderRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListProviderRequest;

  static equals(a: ListProviderRequest | PlainMessage<ListProviderRequest> | undefined, b: ListProviderRequest | PlainMessage<ListProviderRequest> | undefined): boolean;
}

/**
 * @generated from message auth.v1.ListProviderResponse
 */
export declare class ListProviderResponse extends Message<ListProviderResponse> {
  /**
   * @generated from field: repeated auth.v1.Provider providers = 1;
   */
  providers: Provider[];

  constructor(data?: PartialMessage<ListProviderResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "auth.v1.ListProviderResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListProviderResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListProviderResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListProviderResponse;

  static equals(a: ListProviderResponse | PlainMessage<ListProviderResponse> | undefined, b: ListProviderResponse | PlainMessage<ListProviderResponse> | undefined): boolean;
}

/**
 * @generated from message auth.v1.GetUserRequest
 */
export declare class GetUserRequest extends Message<GetUserRequest> {
  constructor(data?: PartialMessage<GetUserRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "auth.v1.GetUserRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserRequest;

  static equals(a: GetUserRequest | PlainMessage<GetUserRequest> | undefined, b: GetUserRequest | PlainMessage<GetUserRequest> | undefined): boolean;
}

/**
 * @generated from message auth.v1.GetUserResponse
 */
export declare class GetUserResponse extends Message<GetUserResponse> {
  /**
   * @generated from field: string email = 1;
   */
  email: string;

  /**
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * @generated from field: string token = 3;
   */
  token: string;

  constructor(data?: PartialMessage<GetUserResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "auth.v1.GetUserResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserResponse;

  static equals(a: GetUserResponse | PlainMessage<GetUserResponse> | undefined, b: GetUserResponse | PlainMessage<GetUserResponse> | undefined): boolean;
}

/**
 * @generated from message auth.v1.LogoutRequest
 */
export declare class LogoutRequest extends Message<LogoutRequest> {
  constructor(data?: PartialMessage<LogoutRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "auth.v1.LogoutRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutRequest;

  static equals(a: LogoutRequest | PlainMessage<LogoutRequest> | undefined, b: LogoutRequest | PlainMessage<LogoutRequest> | undefined): boolean;
}

/**
 * @generated from message auth.v1.LogoutResponse
 */
export declare class LogoutResponse extends Message<LogoutResponse> {
  constructor(data?: PartialMessage<LogoutResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "auth.v1.LogoutResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LogoutResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LogoutResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LogoutResponse;

  static equals(a: LogoutResponse | PlainMessage<LogoutResponse> | undefined, b: LogoutResponse | PlainMessage<LogoutResponse> | undefined): boolean;
}

