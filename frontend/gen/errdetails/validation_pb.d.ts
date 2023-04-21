// @generated by protoc-gen-es v1.2.0 with parameter "target=js+dts"
// @generated from file errdetails/validation.proto (package errdetails.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message errdetails.v1.ValidationError
 */
export declare class ValidationError extends Message<ValidationError> {
  /**
   * @generated from field: map<string, string> errors = 1;
   */
  errors: { [key: string]: string };

  constructor(data?: PartialMessage<ValidationError>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "errdetails.v1.ValidationError";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ValidationError;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ValidationError;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ValidationError;

  static equals(a: ValidationError | PlainMessage<ValidationError> | undefined, b: ValidationError | PlainMessage<ValidationError> | undefined): boolean;
}

