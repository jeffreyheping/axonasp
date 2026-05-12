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
 */
package axonvm

import (
	"os"
	"path/filepath"
	"testing"
)

func TestJScriptModuleDefaultImportExport(t *testing.T) {
	dir := t.TempDir()
	depPath := filepath.Join(dir, "dep.js")
	entryPath := filepath.Join(dir, "entry.js")

	depSrc := `export default function(x) { return x * x; }`
	entrySrc := `import square from "./dep.js"; Response.Write(square(4));`

	if err := os.WriteFile(depPath, []byte(depSrc), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(entryPath, []byte(entrySrc), 0644); err != nil {
		t.Fatal(err)
	}

	out, err := runJScriptModuleEntry(t, entryPath)
	if err != nil {
		t.Fatal(err)
	}
	if out != "16" {
		t.Fatalf("expected 16, got %q", out)
	}
}

func TestJScriptModuleNamespaceImport(t *testing.T) {
	dir := t.TempDir()
	depPath := filepath.Join(dir, "dep.js")
	entryPath := filepath.Join(dir, "entry.js")

	depSrc := `export var a = 1; export var b = 2;`
	entrySrc := `import * as ns from "./dep.js"; Response.Write(ns.a + ns.b);`

	if err := os.WriteFile(depPath, []byte(depSrc), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(entryPath, []byte(entrySrc), 0644); err != nil {
		t.Fatal(err)
	}

	out, err := runJScriptModuleEntry(t, entryPath)
	if err != nil {
		t.Fatal(err)
	}
	if out != "3" {
		t.Fatalf("expected 3, got %q", out)
	}
}

func TestJScriptModuleExportAllFrom(t *testing.T) {
	dir := t.TempDir()
	aPath := filepath.Join(dir, "a.js")
	bPath := filepath.Join(dir, "b.js")
	entryPath := filepath.Join(dir, "entry.js")

	if err := os.WriteFile(aPath, []byte(`export var x = 10;`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(bPath, []byte(`export * from "./a.js";`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(entryPath, []byte(`import { x } from "./b.js"; Response.Write(x);`), 0644); err != nil {
		t.Fatal(err)
	}

	out, err := runJScriptModuleEntry(t, entryPath)
	if err != nil {
		t.Fatal(err)
	}
	if out != "10" {
		t.Fatalf("expected 10, got %q", out)
	}
}

func TestJScriptModuleCombinedImport(t *testing.T) {
	dir := t.TempDir()
	depPath := filepath.Join(dir, "dep.js")
	entryPath := filepath.Join(dir, "entry.js")

	depSrc := `export default 42; export var named = 100;`
	entrySrc := `import def, { named } from "./dep.js"; Response.Write(def + ":" + named);`

	if err := os.WriteFile(depPath, []byte(depSrc), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(entryPath, []byte(entrySrc), 0644); err != nil {
		t.Fatal(err)
	}

	out, err := runJScriptModuleEntry(t, entryPath)
	if err != nil {
		t.Fatal(err)
	}
	if out != "42:100" {
		t.Fatalf("expected 42:100, got %q", out)
	}
}

func TestJScriptModuleExportNamespaceFrom(t *testing.T) {
	dir := t.TempDir()
	aPath := filepath.Join(dir, "a.js")
	bPath := filepath.Join(dir, "b.js")
	entryPath := filepath.Join(dir, "entry.js")

	if err := os.WriteFile(aPath, []byte(`export var x = 1;`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(bPath, []byte(`export * as ns from "./a.js";`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(entryPath, []byte(`import { ns } from "./b.js"; Response.Write(ns.x);`), 0644); err != nil {
		t.Fatal(err)
	}

	out, err := runJScriptModuleEntry(t, entryPath)
	if err != nil {
		t.Fatal(err)
	}
	if out != "1" {
		t.Fatalf("expected 1, got %q", out)
	}
}

func TestJScriptModuleExportDefaultExpression(t *testing.T) {
	dir := t.TempDir()
	depPath := filepath.Join(dir, "dep.js")
	entryPath := filepath.Join(dir, "entry.js")

	depSrc := `export default (1 + 2);`
	entrySrc := `import val from "./dep.js"; Response.Write(val);`

	if err := os.WriteFile(depPath, []byte(depSrc), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(entryPath, []byte(entrySrc), 0644); err != nil {
		t.Fatal(err)
	}

	out, err := runJScriptModuleEntry(t, entryPath)
	if err != nil {
		t.Fatal(err)
	}
	if out != "3" {
		t.Fatalf("expected 3, got %q", out)
	}
}

func TestJScriptModuleMissingExportThrows(t *testing.T) {
	dir := t.TempDir()
	depPath := filepath.Join(dir, "dep.js")
	entryPath := filepath.Join(dir, "entry.js")

	if err := os.WriteFile(depPath, []byte(`export var x = 1;`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(entryPath, []byte(`import { missing } from "./dep.js";`), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := runJScriptModuleEntry(t, entryPath)
	if err == nil {
		t.Fatal("expected error for missing export, got nil")
	}
}
