//go:build lib_wscript_shell_disabled

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

// WScriptShell is the disabled placeholder for WScript.Shell.
type WScriptShell struct{}

// WScriptExecObject is the disabled placeholder for WScript Exec objects.
type WScriptExecObject struct{}

// ProcessTextStream is the disabled placeholder for WScript text streams.
type ProcessTextStream struct{}

// newWScriptShellObject fails because WScript.Shell is disabled at compile time.
func (vm *VM) newWScriptShellObject() Value {
	panicLibraryDisabled("wscript_shell", "WScript.Shell")
	return Value{Type: VTEmpty}
}

func (ws *WScriptShell) DispatchPropertyGet(propertyName string) Value { return Value{Type: VTEmpty} }
func (ws *WScriptShell) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (we *WScriptExecObject) DispatchPropertyGet(propertyName string) Value {
	return Value{Type: VTEmpty}
}
func (we *WScriptExecObject) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (ts *ProcessTextStream) DispatchPropertyGet(propertyName string) Value {
	return Value{Type: VTEmpty}
}
func (ts *ProcessTextStream) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

// expandEnvironmentStrings returns input unchanged when WScript.Shell is disabled.
func expandEnvironmentStrings(input string) string {
	return input
}
