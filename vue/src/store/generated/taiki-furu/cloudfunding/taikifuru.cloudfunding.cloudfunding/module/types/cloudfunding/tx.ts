/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "taikifuru.cloudfunding.cloudfunding";

export interface MsgCreateProject {
  creator: string;
  target: string;
  period: string;
  description: string;
}

export interface MsgCreateProjectResponse {}

export interface MsgFund {
  creator: string;
  id: number;
  amt: string;
}

export interface MsgFundResponse {}

const baseMsgCreateProject: object = {
  creator: "",
  target: "",
  period: "",
  description: "",
};

export const MsgCreateProject = {
  encode(message: MsgCreateProject, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.target !== "") {
      writer.uint32(18).string(message.target);
    }
    if (message.period !== "") {
      writer.uint32(26).string(message.period);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateProject {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateProject } as MsgCreateProject;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.target = reader.string();
          break;
        case 3:
          message.period = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateProject {
    const message = { ...baseMsgCreateProject } as MsgCreateProject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.target !== undefined && object.target !== null) {
      message.target = String(object.target);
    } else {
      message.target = "";
    }
    if (object.period !== undefined && object.period !== null) {
      message.period = String(object.period);
    } else {
      message.period = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    return message;
  },

  toJSON(message: MsgCreateProject): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.target !== undefined && (obj.target = message.target);
    message.period !== undefined && (obj.period = message.period);
    message.description !== undefined &&
      (obj.description = message.description);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateProject>): MsgCreateProject {
    const message = { ...baseMsgCreateProject } as MsgCreateProject;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.target !== undefined && object.target !== null) {
      message.target = object.target;
    } else {
      message.target = "";
    }
    if (object.period !== undefined && object.period !== null) {
      message.period = object.period;
    } else {
      message.period = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    return message;
  },
};

const baseMsgCreateProjectResponse: object = {};

export const MsgCreateProjectResponse = {
  encode(
    _: MsgCreateProjectResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateProjectResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateProjectResponse,
    } as MsgCreateProjectResponse;
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

  fromJSON(_: any): MsgCreateProjectResponse {
    const message = {
      ...baseMsgCreateProjectResponse,
    } as MsgCreateProjectResponse;
    return message;
  },

  toJSON(_: MsgCreateProjectResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateProjectResponse>
  ): MsgCreateProjectResponse {
    const message = {
      ...baseMsgCreateProjectResponse,
    } as MsgCreateProjectResponse;
    return message;
  },
};

const baseMsgFund: object = { creator: "", id: 0, amt: "" };

export const MsgFund = {
  encode(message: MsgFund, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.amt !== "") {
      writer.uint32(26).string(message.amt);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFund {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFund } as MsgFund;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.amt = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgFund {
    const message = { ...baseMsgFund } as MsgFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.amt !== undefined && object.amt !== null) {
      message.amt = String(object.amt);
    } else {
      message.amt = "";
    }
    return message;
  },

  toJSON(message: MsgFund): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    message.amt !== undefined && (obj.amt = message.amt);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgFund>): MsgFund {
    const message = { ...baseMsgFund } as MsgFund;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.amt !== undefined && object.amt !== null) {
      message.amt = object.amt;
    } else {
      message.amt = "";
    }
    return message;
  },
};

const baseMsgFundResponse: object = {};

export const MsgFundResponse = {
  encode(_: MsgFundResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgFundResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgFundResponse } as MsgFundResponse;
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

  fromJSON(_: any): MsgFundResponse {
    const message = { ...baseMsgFundResponse } as MsgFundResponse;
    return message;
  },

  toJSON(_: MsgFundResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgFundResponse>): MsgFundResponse {
    const message = { ...baseMsgFundResponse } as MsgFundResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateProject(request: MsgCreateProject): Promise<MsgCreateProjectResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Fund(request: MsgFund): Promise<MsgFundResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateProject(request: MsgCreateProject): Promise<MsgCreateProjectResponse> {
    const data = MsgCreateProject.encode(request).finish();
    const promise = this.rpc.request(
      "taikifuru.cloudfunding.cloudfunding.Msg",
      "CreateProject",
      data
    );
    return promise.then((data) =>
      MsgCreateProjectResponse.decode(new Reader(data))
    );
  }

  Fund(request: MsgFund): Promise<MsgFundResponse> {
    const data = MsgFund.encode(request).finish();
    const promise = this.rpc.request(
      "taikifuru.cloudfunding.cloudfunding.Msg",
      "Fund",
      data
    );
    return promise.then((data) => MsgFundResponse.decode(new Reader(data)));
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
