// @strict: true
// @declaration: true

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
