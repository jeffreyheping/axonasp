//go:build lib_g3axonfunctions_disabled

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

// AxonGlobalFunctionNames is empty when G3Axon.Functions is disabled.
var AxonGlobalFunctionNames = []string{}

// AxonGlobalFunctionPointers is empty when G3Axon.Functions is disabled.
var AxonGlobalFunctionPointers = []BuiltinFunc{}

// AxonLibrary is the disabled stub for the G3AXON library.
type AxonLibrary struct{}

func (vm *VM) newAxonLibrary() Value {
	panicLibraryDisabled("g3axonfunctions", "G3Axon.Functions library")
	return Value{Type: VTEmpty}
}

func (al *AxonLibrary) DispatchPropertyGet(propertyName string) Value {
	return Value{Type: VTEmpty}
}

func (al *AxonLibrary) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

// loadAxConfigValue is disabled with the G3Axon.Functions library.
func loadAxConfigValue(configKey string) (interface{}, bool) {
	return nil, false
}
