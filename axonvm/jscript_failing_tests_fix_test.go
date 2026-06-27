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

// TestJScriptPrimitiveWrappersAndFunctionCtor checks the soft equality comparisons
// of primitive wrapper objects, Object.prototype.toString.call(wrapper) behavior,
// and dynamic Function constructor compilation.
func TestJScriptPrimitiveWrappersAndFunctionCtor(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{`new String("x") == "x"`, "True"},
		{`new Number(1) == 1`, "True"},
		{`new Boolean(false) == false`, "True"},
		{`new String("") == false`, "True"},
		{`new Number(0) == false`, "True"},
		{`Object.prototype.toString.call(new String("x"))`, "[object String]"},
		{`Object.prototype.toString.call(new Number(1))`, "[object Number]"},
		{`Object.prototype.toString.call(new Boolean(false))`, "[object Boolean]"},
		{`Function("return 7")()`, "7"},
		{`typeof Function("return 1")`, "function"},
		{`Function("a, b", "return a + b")(2, 3)`, "5"},
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
