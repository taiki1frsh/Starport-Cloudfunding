import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "taikifuru.cloudfunding.cloudfunding";
export interface MsgCreateProject {
    creator: string;
    target: string;
    deadline: string;
    description: string;
}
export interface MsgCreateProjectResponse {
}
export interface MsgFund {
    creator: string;
    id: number;
}
export interface MsgFundResponse {
}
export declare const MsgCreateProject: {
    encode(message: MsgCreateProject, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateProject;
    fromJSON(object: any): MsgCreateProject;
    toJSON(message: MsgCreateProject): unknown;
    fromPartial(object: DeepPartial<MsgCreateProject>): MsgCreateProject;
};
export declare const MsgCreateProjectResponse: {
    encode(_: MsgCreateProjectResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateProjectResponse;
    fromJSON(_: any): MsgCreateProjectResponse;
    toJSON(_: MsgCreateProjectResponse): unknown;
    fromPartial(_: DeepPartial<MsgCreateProjectResponse>): MsgCreateProjectResponse;
};
export declare const MsgFund: {
    encode(message: MsgFund, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFund;
    fromJSON(object: any): MsgFund;
    toJSON(message: MsgFund): unknown;
    fromPartial(object: DeepPartial<MsgFund>): MsgFund;
};
export declare const MsgFundResponse: {
    encode(_: MsgFundResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgFundResponse;
    fromJSON(_: any): MsgFundResponse;
    toJSON(_: MsgFundResponse): unknown;
    fromPartial(_: DeepPartial<MsgFundResponse>): MsgFundResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    CreateProject(request: MsgCreateProject): Promise<MsgCreateProjectResponse>;
    /** this line is used by starport scaffolding # proto/tx/rpc */
    Fund(request: MsgFund): Promise<MsgFundResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateProject(request: MsgCreateProject): Promise<MsgCreateProjectResponse>;
    Fund(request: MsgFund): Promise<MsgFundResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
