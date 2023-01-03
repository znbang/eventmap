// @generated by protoc-gen-connect-web v0.3.3 with parameter "target=js+dts"
// @generated from file event/v1/event.proto (package event.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateEventRequest, CreateEventResponse, DeleteEventRequest, DeleteEventResponse, GetEventRequest, GetEventResponse, ListActiveEventRequest, ListActiveEventResponse, ListEventRequest, ListEventResponse, ListUserEventRequest, ListUserEventResponse, UpdateEventRequest, UpdateEventResponse } from "./event_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service event.v1.EventService
 */
export const EventService = {
  typeName: "event.v1.EventService",
  methods: {
    /**
     * @generated from rpc event.v1.EventService.CreateEvent
     */
    createEvent: {
      name: "CreateEvent",
      I: CreateEventRequest,
      O: CreateEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc event.v1.EventService.UpdateEvent
     */
    updateEvent: {
      name: "UpdateEvent",
      I: UpdateEventRequest,
      O: UpdateEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc event.v1.EventService.DeleteEvent
     */
    deleteEvent: {
      name: "DeleteEvent",
      I: DeleteEventRequest,
      O: DeleteEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc event.v1.EventService.GetEvent
     */
    getEvent: {
      name: "GetEvent",
      I: GetEventRequest,
      O: GetEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc event.v1.EventService.ListActiveEvent
     */
    listActiveEvent: {
      name: "ListActiveEvent",
      I: ListActiveEventRequest,
      O: ListActiveEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc event.v1.EventService.ListEvent
     */
    listEvent: {
      name: "ListEvent",
      I: ListEventRequest,
      O: ListEventResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc event.v1.EventService.ListUserEvent
     */
    listUserEvent: {
      name: "ListUserEvent",
      I: ListUserEventRequest,
      O: ListUserEventResponse,
      kind: MethodKind.Unary,
    },
  }
};
