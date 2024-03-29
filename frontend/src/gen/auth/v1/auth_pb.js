// @generated by protoc-gen-es v1.8.0 with parameter "target=js+dts"
// @generated from file auth/v1/auth.proto (package auth.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message auth.v1.Provider
 */
export const Provider = /*@__PURE__*/ proto3.makeMessageType(
  "auth.v1.Provider",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message auth.v1.ListProviderRequest
 */
export const ListProviderRequest = /*@__PURE__*/ proto3.makeMessageType(
  "auth.v1.ListProviderRequest",
  [],
);

/**
 * @generated from message auth.v1.ListProviderResponse
 */
export const ListProviderResponse = /*@__PURE__*/ proto3.makeMessageType(
  "auth.v1.ListProviderResponse",
  () => [
    { no: 1, name: "providers", kind: "message", T: Provider, repeated: true },
  ],
);

/**
 * @generated from message auth.v1.GetUserRequest
 */
export const GetUserRequest = /*@__PURE__*/ proto3.makeMessageType(
  "auth.v1.GetUserRequest",
  [],
);

/**
 * @generated from message auth.v1.GetUserResponse
 */
export const GetUserResponse = /*@__PURE__*/ proto3.makeMessageType(
  "auth.v1.GetUserResponse",
  () => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message auth.v1.LogoutRequest
 */
export const LogoutRequest = /*@__PURE__*/ proto3.makeMessageType(
  "auth.v1.LogoutRequest",
  [],
);

/**
 * @generated from message auth.v1.LogoutResponse
 */
export const LogoutResponse = /*@__PURE__*/ proto3.makeMessageType(
  "auth.v1.LogoutResponse",
  [],
);

