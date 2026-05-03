//go:build lib_mswc_disabled

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

// MSWC library object placeholders.
type G3AdRotator struct{}
type G3BrowserType struct{}
type G3NextLink struct{}
type G3ContentRotator struct{}
type G3Counters struct{}
type G3PageCounter struct{}
type G3Tools struct{}
type G3MyInfo struct{}
type G3PermissionChecker struct{}

func (vm *VM) newG3AdRotatorObject() Value {
	panicLibraryDisabled("mswc", "MSWC.AdRotator")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3BrowserTypeObject() Value {
	panicLibraryDisabled("mswc", "MSWC.BrowserType")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3NextLinkObject() Value {
	panicLibraryDisabled("mswc", "MSWC.NextLink")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3ContentRotatorObject() Value {
	panicLibraryDisabled("mswc", "MSWC.ContentRotator")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3CountersObject() Value {
	panicLibraryDisabled("mswc", "MSWC.Counters")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3PageCounterObject() Value {
	panicLibraryDisabled("mswc", "MSWC.PageCounter")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3ToolsObject() Value {
	panicLibraryDisabled("mswc", "MSWC.Tools")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3MyInfoObject() Value {
	panicLibraryDisabled("mswc", "MSWC.MyInfo")
	return Value{Type: VTEmpty}
}
func (vm *VM) newG3PermissionCheckerObject() Value {
	panicLibraryDisabled("mswc", "MSWC.PermissionChecker")
	return Value{Type: VTEmpty}
}

func (lib *G3AdRotator) DispatchPropertyGet(name string) Value              { return Value{Type: VTEmpty} }
func (lib *G3AdRotator) DispatchPropertySet(name string, args []Value) bool { return false }
func (lib *G3AdRotator) DispatchMethod(name string, args []Value) Value     { return Value{Type: VTEmpty} }

func (lib *G3BrowserType) DispatchPropertyGet(name string) Value { return Value{Type: VTEmpty} }
func (lib *G3BrowserType) DispatchMethod(name string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (lib *G3NextLink) DispatchPropertyGet(name string) Value          { return Value{Type: VTEmpty} }
func (lib *G3NextLink) DispatchMethod(name string, args []Value) Value { return Value{Type: VTEmpty} }

func (lib *G3ContentRotator) DispatchPropertyGet(name string) Value { return Value{Type: VTEmpty} }
func (lib *G3ContentRotator) DispatchMethod(name string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (lib *G3Counters) DispatchPropertyGet(name string) Value          { return Value{Type: VTEmpty} }
func (lib *G3Counters) DispatchMethod(name string, args []Value) Value { return Value{Type: VTEmpty} }

func (lib *G3PageCounter) DispatchPropertyGet(name string) Value { return Value{Type: VTEmpty} }
func (lib *G3PageCounter) DispatchMethod(name string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (lib *G3Tools) DispatchPropertyGet(name string) Value          { return Value{Type: VTEmpty} }
func (lib *G3Tools) DispatchMethod(name string, args []Value) Value { return Value{Type: VTEmpty} }

func (lib *G3MyInfo) DispatchPropertyGet(name string) Value              { return Value{Type: VTEmpty} }
func (lib *G3MyInfo) DispatchPropertySet(name string, args []Value) bool { return false }
func (lib *G3MyInfo) DispatchMethod(name string, args []Value) Value     { return Value{Type: VTEmpty} }

func (lib *G3PermissionChecker) DispatchPropertyGet(name string) Value { return Value{Type: VTEmpty} }
func (lib *G3PermissionChecker) DispatchMethod(name string, args []Value) Value {
	return Value{Type: VTEmpty}
}
