/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../cloudfunding/params";
import { Project } from "../cloudfunding/project";
export const protobufPackage = "taikifuru.cloudfunding.cloudfunding";
const baseGenesisState = { projectCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        if (message.params !== undefined) {
            Params.encode(message.params, writer.uint32(10).fork()).ldelim();
        }
        for (const v of message.projectList) {
            Project.encode(v, writer.uint32(18).fork()).ldelim();
        }
        if (message.projectCount !== 0) {
            writer.uint32(24).uint64(message.projectCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.projectList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.params = Params.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.projectList.push(Project.decode(reader, reader.uint32()));
                    break;
                case 3:
                    message.projectCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.projectList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromJSON(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.projectList !== undefined && object.projectList !== null) {
            for (const e of object.projectList) {
                message.projectList.push(Project.fromJSON(e));
            }
        }
        if (object.projectCount !== undefined && object.projectCount !== null) {
            message.projectCount = Number(object.projectCount);
        }
        else {
            message.projectCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.params !== undefined &&
            (obj.params = message.params ? Params.toJSON(message.params) : undefined);
        if (message.projectList) {
            obj.projectList = message.projectList.map((e) => e ? Project.toJSON(e) : undefined);
        }
        else {
            obj.projectList = [];
        }
        message.projectCount !== undefined &&
            (obj.projectCount = message.projectCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.projectList = [];
        if (object.params !== undefined && object.params !== null) {
            message.params = Params.fromPartial(object.params);
        }
        else {
            message.params = undefined;
        }
        if (object.projectList !== undefined && object.projectList !== null) {
            for (const e of object.projectList) {
                message.projectList.push(Project.fromPartial(e));
            }
        }
        if (object.projectCount !== undefined && object.projectCount !== null) {
            message.projectCount = object.projectCount;
        }
        else {
            message.projectCount = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
