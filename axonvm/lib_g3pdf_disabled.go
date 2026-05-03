//go:build lib_g3pdf_disabled

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

// G3PDF is the disabled stub for the G3PDF library.
type G3PDF struct{}

func NewG3PDF(ctx *VM) *G3PDF {
	panicLibraryDisabled("g3pdf", "G3PDF library")
	return nil
}

func (p *G3PDF) DispatchPropertyGet(name string) Value {
	return Value{Type: VTEmpty}
}

func (p *G3PDF) DispatchPropertySet(name string, args []Value) bool {
	return false
}

func (p *G3PDF) DispatchMethod(name string, args []Value) Value {
	return Value{Type: VTEmpty}
}
