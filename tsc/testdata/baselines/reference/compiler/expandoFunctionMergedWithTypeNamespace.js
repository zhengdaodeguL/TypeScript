//// [tests/cases/compiler/expandoFunctionMergedWithTypeNamespace.ts] ////

//// [expandoFunctionMergedWithTypeNamespace.ts]
export function f() {}
f.a = 0;
export namespace f {
    export type a = 0;
}
export const value: number = f.a;
export const type: f.a = 0;

export function Popover() { return null; }
Popover.Target = (x: number) => x;
Popover.displayName = "Popover";
export namespace Popover {
    export namespace Target {
        export type Props = number;
    }
}
export const result: number = Popover.Target(1);
export const name: string = Popover.displayName;
export const props: Popover.Target.Props = 1;


//// [expandoFunctionMergedWithTypeNamespace.js]
export function f() { }
f.a = 0;
export const value = f.a;
export const type = 0;
export function Popover() { return null; }
Popover.Target = (x) => x;
Popover.displayName = "Popover";
export const result = Popover.Target(1);
export const name = Popover.displayName;
export const props = 1;


//// [expandoFunctionMergedWithTypeNamespace.d.ts]
export declare function f(): void;
export declare namespace f {
    var a: number;
}
export declare namespace f {
    type a = 0;
}
export declare const value: number;
export declare const type: f.a;
export declare function Popover(): null;
export declare namespace Popover {
    var Target: (x: number) => number;
    var displayName: string;
}
export declare namespace Popover {
    namespace Target {
        type Props = number;
    }
}
export declare const result: number;
export declare const name: string;
export declare const props: Popover.Target.Props;
