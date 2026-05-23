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
