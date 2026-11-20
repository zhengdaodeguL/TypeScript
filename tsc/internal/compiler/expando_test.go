package compiler_test

import (
	"testing"

	"github.com/microsoft/TypeScript/tsc/internal/compiler"
	"github.com/microsoft/TypeScript/tsc/internal/core"
	"github.com/microsoft/TypeScript/tsc/internal/tsoptions"
	"github.com/microsoft/TypeScript/tsc/internal/tspath"
	"github.com/microsoft/TypeScript/tsc/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestExpandoPropertyWithMergedNamespace(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		name   string
		source string
	}{
		{
			name: "type alias",
			source: `
function f() {}
f.a = 0;
namespace f { export type a = 0; }
f.a = 1;
const n: number = f.a;
const a: f.a = 0;
// @ts-expect-error
const s: string = f.a;
// @ts-expect-error
const b: f.a = 1;
`,
		},
		{
			name: "interface",
			source: `
function f() {}
f.a = 0;
namespace f { export interface a { value: string; } }
const n: number = f.a;
const a: f.a = { value: "a" };
// @ts-expect-error
const b: f.a = 0;
`,
		},
		{
			name: "empty namespace",
			source: `
function f() {}
f.a = 0;
namespace f { export namespace a {} }
const n: number = f.a;
`,
		},
		{
			name: "namespace before function",
			source: `
namespace f { export type a = 0; }
function f() {}
f.a = 0;
const n: number = f.a;
const a: f.a = 0;
`,
		},
		{
			name: "nested type-only namespace",
			source: `
export function Popover() { return null; }
Popover.Target = (x: number) => x;
Popover.displayName = "Popover";
export namespace Popover {
    export namespace Target {
        export type Props = number;
    }
}
const n: number = Popover.Target(1);
const name: string = Popover.displayName;
const props: Popover.Target.Props = 1;
// @ts-expect-error
Popover.Target("a");
// @ts-expect-error
const wrongProps: Popover.Target.Props = "a";
`,
		},
		{
			name: "value declaration retains priority",
			source: `
function f() {}
// @ts-expect-error
f.a = 0;
namespace f { export let a: string; }
const s: string = f.a;
`,
		},
		{
			name: "value alias retains priority",
			source: `
namespace other { export let a: string; }
function f() {}
f.a = "a";
namespace f { export import a = other.a; }
const s: string = f.a;
// @ts-expect-error
f.a = 0;
`,
		},
		{
			name: "class static member retains priority",
			source: `
declare class Thing { static stuff(): number; }
function Thing() {}
Thing.stuff = () => 7;
const n: number = Thing.stuff();
// @ts-expect-error
Thing.stuff = () => "a";
`,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			fs := vfstest.FromMap(map[string]string{
				"/src/index.ts": test.source,
			}, tspath.CaseSensitive)
			program, err := compiler.NewProgram(t.Context(), compiler.ProgramOptions{
				Config: tsoptions.NewParsedCommandLine(
					&core.CompilerOptions{NoLib: core.TSTrue, Strict: core.TSTrue},
					testFileNames("/src/index.ts"), nil, "/src", tspath.CaseSensitive,
				),
				Host: compiler.NewCompilerHost(fs, "/", nil, nil, nil),
			})
			assert.NilError(t, err)
			assert.Equal(t, len(program.GetSyntacticDiagnostics(t.Context(), nil)), 0)
			for _, diagnostic := range program.GetSemanticDiagnostics(t.Context(), nil) {
				t.Errorf("unexpected TS%d: %s", diagnostic.Code(), diagnostic.String())
			}
		})
	}
}
