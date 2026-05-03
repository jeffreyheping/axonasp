//go:build lib_g3json_disabled

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

// G3JSON is the disabled stub for the G3JSON library.
type G3JSON struct{}

func (vm *VM) newG3JSONObject() Value {
	panicLibraryDisabled("g3json", "G3JSON library")
	return Value{Type: VTEmpty}
}

func (j *G3JSON) DispatchPropertyGet(propertyName string) Value {
	return Value{Type: VTEmpty}
}

func (j *G3JSON) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}
