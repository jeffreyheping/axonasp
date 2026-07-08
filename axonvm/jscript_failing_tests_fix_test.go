package axonvm

import (
	"strings"
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

func TestJScriptStringCastingTTRegressions(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{`new Array(5).toString()`, `,,,,`},
		{`new String(true).toString()`, `true`},
		{`new String({}).toString()`, `[object Object]`},
		{`new String([]).toString()`, ``},
		{`new Object(5).toString()`, `5`},
		{`new Object("test").toString()`, `test`},
		{`new Object([]).toString()`, ``},
		{`new Object(true).toString()`, `true`},
		{`new Function().toString()`, "function anonymous() {\n\n}"},
		{`new Function("return 5").toString()`, "function anonymous() {\nreturn 5\n}"},
		{`new Function("intParam1", "intParam2", "return intParam1 + intParam2").toString()`, "function anonymous(intParam1, intParam2) {\nreturn intParam1 + intParam2\n}"},
		{`new RegExp().toString()`, `//`},
		{`new Enumerator().toString()`, `[object Object]`},
		{`(function hi(){return "hi";}).toString()`, `(function hi(){return "hi";})`},
		{`(function(strName){return "hi " + strName;}).toString()`, `(function(strName){return "hi " + strName;})`},
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

func TestJScriptResponseWriteImplicitCoercion(t *testing.T) {
	aspSrc := jscriptSrc(`
var arrTest = new Array(5);
Response.Write(arrTest);

var intTest = new Number(5);
Response.Write(intTest);

var strTest = new String("Hello, world");
Response.Write(strTest);

var objTest = new Object();
Response.Write(objTest);

var fnTest = new Function();
Response.Write(fnTest);

var dteTest = new Date("Tue Jul 7 10:35:27 UTC+0100 2026");
Response.Write(dteTest);

var blnTest = new Boolean(false);
Response.Write(blnTest);
`)
	out, err := runJScript2(t, aspSrc)
	if err != nil {
		t.Fatalf("script failed: %v", err)
	}
	if strings.Contains(out, "NaN") {
		t.Fatalf("unexpected output: Date coercion still produced NaN: %q", out)
	}
	if !strings.Contains(out, ",,,,5Hello, world[object Object]function anonymous() {\n\n}Tue Jul") {
		t.Fatalf("unexpected output: missing implicit string coercion prefix: %q", out)
	}
	if !strings.HasSuffix(out, "False") {
		t.Fatalf("unexpected output: missing Boolean object coercion suffix: %q", out)
	}
}
