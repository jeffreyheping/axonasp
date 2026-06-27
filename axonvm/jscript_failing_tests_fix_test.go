package axonvm

import (
	"testing"
)

// TestJScriptInstanceOfChecks validates the correct prototype chain inheritance
// and JScript 'instanceof' behaviour for both functions, built-in constructor
// objects, wrapped primitives, and raw primitives.
func TestJScriptInstanceOfChecks(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{`((function () {}) instanceof Function)`, "True"},
		{`(Date instanceof Function)`, "True"},
		{`(Date instanceof Object)`, "True"},
		{`((new Number(1)) instanceof Number)`, "True"},
		{`((new Boolean(false)) instanceof Boolean)`, "True"},
		{`(1 instanceof Number)`, "False"},
		{`(true instanceof Boolean)`, "False"},
	}

	for _, tc := range tests {
		t.Run(tc.expr, func(t *testing.T) {
			aspSrc := jscriptSrc(`Response.Write(` + tc.expr + `);`)
			out, err := runJScript2(t, aspSrc)
			if err != nil {
				t.Fatalf("Expr %s failed to run: %v", tc.expr, err)
			}
			if out != tc.expected {
				t.Errorf("Expr %s: expected %q, got %q", tc.expr, tc.expected, out)
			}
		})
	}
}
