/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimarães - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 */
package axonvm

import (
	"strings"
	"testing"
)

func TestJScriptSymbolUnscopables(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
	}{
		{
			"Basic unscopables",
			`var obj = { a: 1, b: 2 };
			 obj[Symbol.unscopables] = { a: true };
			 var a = 10, b = 20;
			 with (obj) {
				 Response.Write(a + "," + b);
			 }`,
			"10,2",
		},
		{
			"Unscopables with inheritance",
			`var proto = { a: 1 };
			 proto[Symbol.unscopables] = { a: true };
			 var obj = Object.create(proto);
			 obj.a = 2;
			 var a = 10;
			 with (obj) {
				 Response.Write(a);
			 }`,
			"10", // obj.a is 2, but proto[Symbol.unscopables].a is true, so 'a' should resolve to 10
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := runJScript2(t, jscriptSrc(tt.code))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if out != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, out)
			}
		})
	}
}

func TestJScriptSymbolMatchAll(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
	}{
		{
			"String.prototype.matchAll basic",
			`var str = "test1test2";
			 var regex = /test(\d)/g;
			 var matches = str.matchAll(regex);
			 var res = "";
			 for (var m of matches) {
				 res += m[1] + ",";
			 }
			 Response.Write(res);`,
			"1,2,",
		},
		{
			"RegExp.prototype[Symbol.matchAll]",
			`var str = "abc";
			 var regex = /[a-z]/g;
			 var it = regex[Symbol.matchAll](str);
			 var res = "";
			 for (var m of it) {
				 res += m[0];
			 }
			 Response.Write(res);`,
			"abc",
		},
		{
			"String.prototype.matchAll error on non-global",
			`try { "abc".matchAll(/a/); } catch(e) { Response.Write(e.message); }`,
			"String.prototype.matchAll called with a non-global RegExp argument",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := runJScript2(t, jscriptSrc(tt.code))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !strings.Contains(out, tt.expected) {
				t.Errorf("expected %q to contain %q", out, tt.expected)
			}
		})
	}
}
