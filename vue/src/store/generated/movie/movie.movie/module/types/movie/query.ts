/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../movie/params";
import { Movie } from "../movie/movie";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Review } from "../movie/review";

export const protobufPackage = "movie.movie";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetMovieRequest {
  id: number;
}

export interface QueryGetMovieResponse {
  Movie: Movie | undefined;
}

export interface QueryAllMovieRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMovieResponse {
  Movie: Movie[];
  pagination: PageResponse | undefined;
}

export interface QueryGetReviewRequest {
  id: number;
}

export interface QueryGetReviewResponse {
  Review: Review | undefined;
}

export interface QueryAllReviewRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllReviewResponse {
  Review: Review[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetMovieRequest: object = { id: 0 };

export const QueryGetMovieRequest = {
  encode(
    message: QueryGetMovieRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMovieRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMovieRequest } as QueryGetMovieRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMovieRequest {
    const message = { ...baseQueryGetMovieRequest } as QueryGetMovieRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetMovieRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetMovieRequest>): QueryGetMovieRequest {
    const message = { ...baseQueryGetMovieRequest } as QueryGetMovieRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetMovieResponse: object = {};

export const QueryGetMovieResponse = {
  encode(
    message: QueryGetMovieResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Movie !== undefined) {
      Movie.encode(message.Movie, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMovieResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMovieResponse } as QueryGetMovieResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Movie = Movie.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMovieResponse {
    const message = { ...baseQueryGetMovieResponse } as QueryGetMovieResponse;
    if (object.Movie !== undefined && object.Movie !== null) {
      message.Movie = Movie.fromJSON(object.Movie);
    } else {
      message.Movie = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMovieResponse): unknown {
    const obj: any = {};
    message.Movie !== undefined &&
      (obj.Movie = message.Movie ? Movie.toJSON(message.Movie) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMovieResponse>
  ): QueryGetMovieResponse {
    const message = { ...baseQueryGetMovieResponse } as QueryGetMovieResponse;
    if (object.Movie !== undefined && object.Movie !== null) {
      message.Movie = Movie.fromPartial(object.Movie);
    } else {
      message.Movie = undefined;
    }
    return message;
  },
};

const baseQueryAllMovieRequest: object = {};

export const QueryAllMovieRequest = {
  encode(
    message: QueryAllMovieRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMovieRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMovieRequest } as QueryAllMovieRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMovieRequest {
    const message = { ...baseQueryAllMovieRequest } as QueryAllMovieRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMovieRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllMovieRequest>): QueryAllMovieRequest {
    const message = { ...baseQueryAllMovieRequest } as QueryAllMovieRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMovieResponse: object = {};

export const QueryAllMovieResponse = {
  encode(
    message: QueryAllMovieResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Movie) {
      Movie.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMovieResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMovieResponse } as QueryAllMovieResponse;
    message.Movie = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Movie.push(Movie.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMovieResponse {
    const message = { ...baseQueryAllMovieResponse } as QueryAllMovieResponse;
    message.Movie = [];
    if (object.Movie !== undefined && object.Movie !== null) {
      for (const e of object.Movie) {
        message.Movie.push(Movie.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMovieResponse): unknown {
    const obj: any = {};
    if (message.Movie) {
      obj.Movie = message.Movie.map((e) => (e ? Movie.toJSON(e) : undefined));
    } else {
      obj.Movie = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMovieResponse>
  ): QueryAllMovieResponse {
    const message = { ...baseQueryAllMovieResponse } as QueryAllMovieResponse;
    message.Movie = [];
    if (object.Movie !== undefined && object.Movie !== null) {
      for (const e of object.Movie) {
        message.Movie.push(Movie.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetReviewRequest: object = { id: 0 };

export const QueryGetReviewRequest = {
  encode(
    message: QueryGetReviewRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetReviewRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetReviewRequest } as QueryGetReviewRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetReviewRequest {
    const message = { ...baseQueryGetReviewRequest } as QueryGetReviewRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetReviewRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetReviewRequest>
  ): QueryGetReviewRequest {
    const message = { ...baseQueryGetReviewRequest } as QueryGetReviewRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetReviewResponse: object = {};

export const QueryGetReviewResponse = {
  encode(
    message: QueryGetReviewResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Review !== undefined) {
      Review.encode(message.Review, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetReviewResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetReviewResponse } as QueryGetReviewResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Review = Review.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetReviewResponse {
    const message = { ...baseQueryGetReviewResponse } as QueryGetReviewResponse;
    if (object.Review !== undefined && object.Review !== null) {
      message.Review = Review.fromJSON(object.Review);
    } else {
      message.Review = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetReviewResponse): unknown {
    const obj: any = {};
    message.Review !== undefined &&
      (obj.Review = message.Review ? Review.toJSON(message.Review) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetReviewResponse>
  ): QueryGetReviewResponse {
    const message = { ...baseQueryGetReviewResponse } as QueryGetReviewResponse;
    if (object.Review !== undefined && object.Review !== null) {
      message.Review = Review.fromPartial(object.Review);
    } else {
      message.Review = undefined;
    }
    return message;
  },
};

const baseQueryAllReviewRequest: object = {};

export const QueryAllReviewRequest = {
  encode(
    message: QueryAllReviewRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllReviewRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllReviewRequest } as QueryAllReviewRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllReviewRequest {
    const message = { ...baseQueryAllReviewRequest } as QueryAllReviewRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllReviewRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllReviewRequest>
  ): QueryAllReviewRequest {
    const message = { ...baseQueryAllReviewRequest } as QueryAllReviewRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllReviewResponse: object = {};

export const QueryAllReviewResponse = {
  encode(
    message: QueryAllReviewResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Review) {
      Review.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllReviewResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllReviewResponse } as QueryAllReviewResponse;
    message.Review = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Review.push(Review.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllReviewResponse {
    const message = { ...baseQueryAllReviewResponse } as QueryAllReviewResponse;
    message.Review = [];
    if (object.Review !== undefined && object.Review !== null) {
      for (const e of object.Review) {
        message.Review.push(Review.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllReviewResponse): unknown {
    const obj: any = {};
    if (message.Review) {
      obj.Review = message.Review.map((e) =>
        e ? Review.toJSON(e) : undefined
      );
    } else {
      obj.Review = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllReviewResponse>
  ): QueryAllReviewResponse {
    const message = { ...baseQueryAllReviewResponse } as QueryAllReviewResponse;
    message.Review = [];
    if (object.Review !== undefined && object.Review !== null) {
      for (const e of object.Review) {
        message.Review.push(Review.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Movie by id. */
  Movie(request: QueryGetMovieRequest): Promise<QueryGetMovieResponse>;
  /** Queries a list of Movie items. */
  MovieAll(request: QueryAllMovieRequest): Promise<QueryAllMovieResponse>;
  /** Queries a Review by id. */
  Review(request: QueryGetReviewRequest): Promise<QueryGetReviewResponse>;
  /** Queries a list of Review items. */
  ReviewAll(request: QueryAllReviewRequest): Promise<QueryAllReviewResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("movie.movie.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Movie(request: QueryGetMovieRequest): Promise<QueryGetMovieResponse> {
    const data = QueryGetMovieRequest.encode(request).finish();
    const promise = this.rpc.request("movie.movie.Query", "Movie", data);
    return promise.then((data) =>
      QueryGetMovieResponse.decode(new Reader(data))
    );
  }

  MovieAll(request: QueryAllMovieRequest): Promise<QueryAllMovieResponse> {
    const data = QueryAllMovieRequest.encode(request).finish();
    const promise = this.rpc.request("movie.movie.Query", "MovieAll", data);
    return promise.then((data) =>
      QueryAllMovieResponse.decode(new Reader(data))
    );
  }

  Review(request: QueryGetReviewRequest): Promise<QueryGetReviewResponse> {
    const data = QueryGetReviewRequest.encode(request).finish();
    const promise = this.rpc.request("movie.movie.Query", "Review", data);
    return promise.then((data) =>
      QueryGetReviewResponse.decode(new Reader(data))
    );
  }

  ReviewAll(request: QueryAllReviewRequest): Promise<QueryAllReviewResponse> {
    const data = QueryAllReviewRequest.encode(request).finish();
    const promise = this.rpc.request("movie.movie.Query", "ReviewAll", data);
    return promise.then((data) =>
      QueryAllReviewResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
