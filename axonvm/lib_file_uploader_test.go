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

import (
	"testing"
)

func TestFileUploader(t *testing.T) {
	vm := NewVM(nil, nil, 0)
	vm.host = &MockHost{}

	lib := vm.newG3FileUploaderObject()
	if lib.Type != VTNativeObject {
		t.Fatalf("expected VTNativeObject, got %v", lib.Type)
	}

	obj := vm.fileUploaderItems[lib.Num]
	if obj == nil {
		t.Fatal("expected object in vm items")
	}

	obj.DispatchPropertySet("MaxFileSize", []Value{NewInteger(5000)})
	size := obj.DispatchPropertyGet("MaxFileSize")
	if size.Num != 5000 {
		t.Errorf("expected 5000, got %d", size.Num)
	}

	obj.DispatchMethod("BlockExtension", []Value{NewString(".exe")})
	blocked := obj.DispatchPropertyGet("BlockedExtensions")
	if blocked.Type != VTArray || blocked.Arr == nil {
		t.Fatal("expected array")
	}

	if len(blocked.Arr.Values) != 1 || blocked.Arr.Values[0].String() != ".exe" {
		t.Errorf("expected .exe in blocked list")
	}
}
