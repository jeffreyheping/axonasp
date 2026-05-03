//go:build lib_scripting_filesystemobject_disabled

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

const (
	fsoKindRoot int = iota + 1
	fsoKindFile
	fsoKindFolder
	fsoKindTextStream
	fsoKindFilesCollection
	fsoKindSubFoldersCollection
	fsoKindDrive
	fsoKindDrivesCollection
)

// fsoNativeObject is the disabled placeholder for FSO runtime objects.
type fsoNativeObject struct {
	kind int
	path string
}

// fsoTextStreamState is the disabled placeholder for text stream state.
type fsoTextStreamState struct{}

// newFSORootObject fails because Scripting.FileSystemObject is disabled.
func (vm *VM) newFSORootObject() Value {
	panicLibraryDisabled("scripting_filesystemobject", "Scripting.FileSystemObject")
	return Value{Type: VTEmpty}
}

func (vm *VM) newFSONativeObject(kind int, path string, stream *fsoTextStreamState) Value {
	return Value{Type: VTEmpty}
}

func (vm *VM) newFSODriveObject(drive string) Value { return Value{Type: VTEmpty} }
func (vm *VM) newFSODrivesCollectionObject() Value  { return Value{Type: VTEmpty} }

func (vm *VM) dispatchFSOMethod(objID int64, member string, args []Value) (Value, bool) {
	return Value{Type: VTEmpty}, false
}

func (vm *VM) dispatchFSOPropertySet(objID int64, member string, val Value) bool { return false }

func (vm *VM) dispatchFSOPropertyGet(objID int64, member string) (Value, bool) {
	return Value{Type: VTEmpty}, false
}

func (vm *VM) fsoEnumerateDriveNames() []string {
	return nil
}

func (vm *VM) fsoResolvePath(path string) (string, bool) {
	return "", false
}
