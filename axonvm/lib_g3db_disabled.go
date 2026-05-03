//go:build lib_g3db_disabled

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

// G3DB is the disabled stub for the G3DB library.
type G3DB struct{}

type G3DBResultSet struct{}
type G3DBFields struct{}
type G3DBRow struct{}
type G3DBStatement struct{}
type G3DBTransaction struct{}
type G3DBResult struct{}

func (vm *VM) newG3DBObject() Value {
	panicLibraryDisabled("g3db", "G3DB library")
	return Value{Type: VTEmpty}
}

func (g *G3DB) DispatchPropertyGet(propertyName string) Value              { return Value{Type: VTEmpty} }
func (g *G3DB) DispatchPropertySet(propertyName string, args []Value) bool { return false }
func (g *G3DB) DispatchMethod(methodName string, args []Value) Value       { return Value{Type: VTEmpty} }

func (rs *G3DBResultSet) DispatchPropertyGet(propertyName string) Value              { return Value{Type: VTEmpty} }
func (rs *G3DBResultSet) DispatchPropertySet(propertyName string, args []Value) bool { return false }
func (rs *G3DBResultSet) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (f *G3DBFields) DispatchPropertyGet(propertyName string) Value              { return Value{Type: VTEmpty} }
func (f *G3DBFields) DispatchPropertySet(propertyName string, args []Value) bool { return false }
func (f *G3DBFields) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (r *G3DBRow) DispatchPropertyGet(propertyName string) Value              { return Value{Type: VTEmpty} }
func (r *G3DBRow) DispatchPropertySet(propertyName string, args []Value) bool { return false }
func (r *G3DBRow) DispatchMethod(methodName string, args []Value) Value       { return Value{Type: VTEmpty} }

func (s *G3DBStatement) DispatchPropertyGet(propertyName string) Value              { return Value{Type: VTEmpty} }
func (s *G3DBStatement) DispatchPropertySet(propertyName string, args []Value) bool { return false }
func (s *G3DBStatement) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (t *G3DBTransaction) DispatchPropertyGet(propertyName string) Value              { return Value{Type: VTEmpty} }
func (t *G3DBTransaction) DispatchPropertySet(propertyName string, args []Value) bool { return false }
func (t *G3DBTransaction) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (r *G3DBResult) DispatchPropertyGet(propertyName string) Value              { return Value{Type: VTEmpty} }
func (r *G3DBResult) DispatchPropertySet(propertyName string, args []Value) bool { return false }
func (r *G3DBResult) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}
