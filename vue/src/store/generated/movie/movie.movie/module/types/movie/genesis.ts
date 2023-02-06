/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../movie/params";
import { Movie } from "../movie/movie";
import { Review } from "../movie/review";

export const protobufPackage = "movie.movie";

/** GenesisState defines the movie module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  movieList: Movie[];
  movieCount: number;
  reviewList: Review[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  reviewCount: number;
}

const baseGenesisState: object = { movieCount: 0, reviewCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.movieList) {
      Movie.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.movieCount !== 0) {
      writer.uint32(24).uint64(message.movieCount);
    }
    for (const v of message.reviewList) {
      Review.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.reviewCount !== 0) {
      writer.uint32(40).uint64(message.reviewCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.movieList = [];
    message.reviewList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.movieList.push(Movie.decode(reader, reader.uint32()));
          break;
        case 3:
          message.movieCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.reviewList.push(Review.decode(reader, reader.uint32()));
          break;
        case 5:
          message.reviewCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.movieList = [];
    message.reviewList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.movieList !== undefined && object.movieList !== null) {
      for (const e of object.movieList) {
        message.movieList.push(Movie.fromJSON(e));
      }
    }
    if (object.movieCount !== undefined && object.movieCount !== null) {
      message.movieCount = Number(object.movieCount);
    } else {
      message.movieCount = 0;
    }
    if (object.reviewList !== undefined && object.reviewList !== null) {
      for (const e of object.reviewList) {
        message.reviewList.push(Review.fromJSON(e));
      }
    }
    if (object.reviewCount !== undefined && object.reviewCount !== null) {
      message.reviewCount = Number(object.reviewCount);
    } else {
      message.reviewCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.movieList) {
      obj.movieList = message.movieList.map((e) =>
        e ? Movie.toJSON(e) : undefined
      );
    } else {
      obj.movieList = [];
    }
    message.movieCount !== undefined && (obj.movieCount = message.movieCount);
    if (message.reviewList) {
      obj.reviewList = message.reviewList.map((e) =>
        e ? Review.toJSON(e) : undefined
      );
    } else {
      obj.reviewList = [];
    }
    message.reviewCount !== undefined &&
      (obj.reviewCount = message.reviewCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.movieList = [];
    message.reviewList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.movieList !== undefined && object.movieList !== null) {
      for (const e of object.movieList) {
        message.movieList.push(Movie.fromPartial(e));
      }
    }
    if (object.movieCount !== undefined && object.movieCount !== null) {
      message.movieCount = object.movieCount;
    } else {
      message.movieCount = 0;
    }
    if (object.reviewList !== undefined && object.reviewList !== null) {
      for (const e of object.reviewList) {
        message.reviewList.push(Review.fromPartial(e));
      }
    }
    if (object.reviewCount !== undefined && object.reviewCount !== null) {
      message.reviewCount = object.reviewCount;
    } else {
      message.reviewCount = 0;
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
