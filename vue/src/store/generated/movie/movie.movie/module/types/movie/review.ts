/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "movie.movie";

export interface Review {
  id: number;
  movieId: number;
  ratingUint: string;
  description: string;
  creator: string;
}

const baseReview: object = {
  id: 0,
  movieId: 0,
  ratingUint: "",
  description: "",
  creator: "",
};

export const Review = {
  encode(message: Review, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.movieId !== 0) {
      writer.uint32(16).uint64(message.movieId);
    }
    if (message.ratingUint !== "") {
      writer.uint32(26).string(message.ratingUint);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.creator !== "") {
      writer.uint32(42).string(message.creator);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Review {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseReview } as Review;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.movieId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.ratingUint = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Review {
    const message = { ...baseReview } as Review;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.movieId !== undefined && object.movieId !== null) {
      message.movieId = Number(object.movieId);
    } else {
      message.movieId = 0;
    }
    if (object.ratingUint !== undefined && object.ratingUint !== null) {
      message.ratingUint = String(object.ratingUint);
    } else {
      message.ratingUint = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: Review): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.movieId !== undefined && (obj.movieId = message.movieId);
    message.ratingUint !== undefined && (obj.ratingUint = message.ratingUint);
    message.description !== undefined &&
      (obj.description = message.description);
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(object: DeepPartial<Review>): Review {
    const message = { ...baseReview } as Review;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.movieId !== undefined && object.movieId !== null) {
      message.movieId = object.movieId;
    } else {
      message.movieId = 0;
    }
    if (object.ratingUint !== undefined && object.ratingUint !== null) {
      message.ratingUint = object.ratingUint;
    } else {
      message.ratingUint = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
