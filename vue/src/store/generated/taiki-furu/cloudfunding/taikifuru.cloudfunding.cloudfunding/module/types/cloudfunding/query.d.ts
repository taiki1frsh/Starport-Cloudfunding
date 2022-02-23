import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cloudfunding/params";
import { Project } from "../cloudfunding/project";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "taikifuru.cloudfunding.cloudfunding";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryGetProjectRequest {
    id: number;
}
export interface QueryGetProjectResponse {
    Project: Project | undefined;
}
export interface QueryAllProjectRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllProjectResponse {
    Project: Project[];
    pagination: PageResponse | undefined;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryGetProjectRequest: {
    encode(message: QueryGetProjectRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProjectRequest;
    fromJSON(object: any): QueryGetProjectRequest;
    toJSON(message: QueryGetProjectRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetProjectRequest>): QueryGetProjectRequest;
};
export declare const QueryGetProjectResponse: {
    encode(message: QueryGetProjectResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetProjectResponse;
    fromJSON(object: any): QueryGetProjectResponse;
    toJSON(message: QueryGetProjectResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetProjectResponse>): QueryGetProjectResponse;
};
export declare const QueryAllProjectRequest: {
    encode(message: QueryAllProjectRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProjectRequest;
    fromJSON(object: any): QueryAllProjectRequest;
    toJSON(message: QueryAllProjectRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllProjectRequest>): QueryAllProjectRequest;
};
export declare const QueryAllProjectResponse: {
    encode(message: QueryAllProjectResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllProjectResponse;
    fromJSON(object: any): QueryAllProjectResponse;
    toJSON(message: QueryAllProjectResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllProjectResponse>): QueryAllProjectResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a Project by id. */
    Project(request: QueryGetProjectRequest): Promise<QueryGetProjectResponse>;
    /** Queries a list of Project items. */
    ProjectAll(request: QueryAllProjectRequest): Promise<QueryAllProjectResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Project(request: QueryGetProjectRequest): Promise<QueryGetProjectResponse>;
    ProjectAll(request: QueryAllProjectRequest): Promise<QueryAllProjectResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
