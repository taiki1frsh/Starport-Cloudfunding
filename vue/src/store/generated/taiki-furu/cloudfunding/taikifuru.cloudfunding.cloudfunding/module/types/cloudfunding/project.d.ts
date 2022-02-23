import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "taikifuru.cloudfunding.cloudfunding";
export interface Project {
    id: number;
    target: string;
    collected: string;
    deadline: string;
    description: string;
    state: string;
    creator: string;
}
export declare const Project: {
    encode(message: Project, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Project;
    fromJSON(object: any): Project;
    toJSON(message: Project): unknown;
    fromPartial(object: DeepPartial<Project>): Project;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
