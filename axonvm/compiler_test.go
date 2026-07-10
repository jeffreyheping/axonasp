/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimarães - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 * Attribution Notice:
 * If this software is used in other projects, the name "AxonASP Server"
 * must be cited in the documentation or "About" section.
 *
 * Contribution Policy:
 * Modifications to the core source code of AxonASP Server must be
 * made available under this same license terms.
 */
package axonvm

import "testing"

// TestVBScriptConstantFolding validates integer and string folding in one
// compile pass and confirms OpNop replacement keeps opcode stream stable.
func TestVBScriptConstantFolding(t *testing.T) {
	source := `<% Dim x, y
	x = 5 + 10
	y = "a" & "b"
	Response.Write x & "|" & y
%>`
	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	has15 := false
	hasAB := false
	for _, c := range compiler.Constants() {
		if c.Type == VTInteger && c.Num == 15 {
			has15 = true
		}
		if c.Type == VTString && c.Str == "ab" {
			hasAB = true
		}
	}
	if !has15 || !hasAB {
		t.Fatalf("folded constants missing: has15=%v hasAB=%v", has15, hasAB)
	}
	if scanBytecodeForOp(compiler.Bytecode(), OpAdd) {
		t.Fatalf("unexpected OpAdd after folding")
	}
	if scanBytecodeForOp(compiler.Bytecode(), OpConcat) {
		t.Fatalf("unexpected OpConcat after folding")
	}
	if !scanBytecodeForOp(compiler.Bytecode(), OpNop) {
		t.Fatalf("expected OpNop placeholders after folding")
	}

	out := runVBSAndGetOutput(t, source)
	if out != "15|ab" {
		t.Fatalf("unexpected output: %q", out)
	}
}

// TestJScriptConstantFolding verifies AST-level JScript constant folding emits
// direct constants and preserves runtime output.
func TestJScriptConstantFolding(t *testing.T) {
	source := `<%@ Language="JScript" %>
<% var x = 5 + 10; Response.Write(x); %>`
	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	if scanBytecodeForOp(compiler.Bytecode(), OpJSAdd) {
		t.Fatalf("unexpected OpJSAdd for folded literal expression")
	}
	found := false
	for _, c := range compiler.Constants() {
		if (c.Type == VTInteger && c.Num == 15) || (c.Type == VTDouble && c.Flt == 15) {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("missing folded JScript constant 15")
	}

	out := runVBSAndGetOutput(t, source)
	if out != "15" {
		t.Fatalf("unexpected output: %q", out)
	}
}

// TestJumpIntegrity verifies VBScript absolute jump offsets remain valid after
// peephole folding replaces bytes with OpNop instead of shrinking bytecode.
func TestJumpIntegrity(t *testing.T) {
	source := `<% Dim x
If 10 > 5 Then
	x = 5 + 10
Else
	x = 1 + 1
End If
Response.Write x
%>`
	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	if !scanBytecodeForOp(compiler.Bytecode(), OpJumpIfFalse) && !scanBytecodeForOp(compiler.Bytecode(), OpJumpIfLte) {
		t.Fatalf("expected jump opcode (OpJumpIfFalse or OpJumpIfLte) in compiled If/Else bytecode")
	}
	if !scanBytecodeForOp(compiler.Bytecode(), OpNop) {
		t.Fatalf("expected OpNop placeholders in folded If/Else bytecode")
	}

	out := runVBSAndGetOutput(t, source)
	if out != "15" {
		t.Fatalf("jump integrity failed, got %q", out)
	}
}

// TestCompileZeroArgSubCall verifies that a zero-argument bare sub call (declared before) compiles and runs correctly.
func TestCompileZeroArgSubCall(t *testing.T) {
	source := `<%
	Sub Greet()
		Response.Write "Hello from Sub"
	End Sub
	Greet
	%>`
	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	if !scanBytecodeForOp(compiler.Bytecode(), OpCall) {
		t.Fatalf("expected OpCall in compiled zero-arg bare Sub bytecode")
	}

	out := runVBSAndGetOutput(t, source)
	if out != "Hello from Sub" {
		t.Fatalf("unexpected output: %q", out)
	}
}

// TestCompileZeroArgSubForwardRef verifies that a zero-argument bare sub call (declared after) compiles and runs correctly.
func TestCompileZeroArgSubForwardRef(t *testing.T) {
	source := `<%
	Greet
	Sub Greet()
		Response.Write "Hello from Sub Forward"
	End Sub
	%>`
	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	if !scanBytecodeForOp(compiler.Bytecode(), OpCall) {
		t.Fatalf("expected OpCall in compiled zero-arg bare Sub forward ref bytecode")
	}

	out := runVBSAndGetOutput(t, source)
	if out != "Hello from Sub Forward" {
		t.Fatalf("unexpected output: %q", out)
	}
}

// TestVBScriptStaticObjectReference verifies that local Static variables can hold
// object references, persist across executions, and evaluate correctly with "Is Nothing".
func TestVBScriptStaticObjectReference(t *testing.T) {
	source := `<%
	Class Counter
		Public Val
		Private Sub Class_Initialize
			Val = 0
		End Sub
	End Class

	Function GetCounter()
		Static c
		If c Is Nothing Then
			Set c = New Counter
		End If
		c.Val = c.Val + 1
		Set GetCounter = c
	End Function

	Dim c1, c2
	Set c1 = GetCounter()
	Set c2 = GetCounter()

	Response.Write "c1=" & c1.Val & "|c2=" & c2.Val & "|" & (c1 Is c2)
	%>`

	out := runVBSAndGetOutput(t, source)
	expected := "c1=2|c2=2|True"
	if out != expected {
		t.Fatalf("unexpected static object reference output: got %q, want %q", out, expected)
	}
}
