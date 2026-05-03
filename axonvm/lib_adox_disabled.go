//go:build lib_adodb_disabled

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

import "strings"

// ADOXCatalog is the tag-disabled fallback object used when ADODB is disabled.
type ADOXCatalog struct {
	vm               *VM
	activeConnection Value
	tables           Value
}

// newADOXCatalogObject instantiates the ADOX.Catalog fallback object.
// this must be used in place of the real ADOXCatalog when ADODB support is disabled, to ensure consistent error handling.
func (vm *VM) newADOXCatalogObject() Value {
	obj := &ADOXCatalog{
		vm:               vm,
		activeConnection: Value{Type: VTEmpty},
		tables:           Value{Type: VTEmpty},
	}
	id := vm.nextDynamicNativeID
	vm.nextDynamicNativeID++
	vm.adoxCatalogItems[id] = obj
	return Value{Type: VTNativeObject, Num: id}
}

// DispatchPropertyGet acts as a getter.
func (c *ADOXCatalog) DispatchPropertyGet(propertyName string) Value {
	switch strings.ToLower(strings.TrimSpace(propertyName)) {
	case "tables":
		return c.getTables()
	case "activeconnection":
		return c.activeConnection
	}
	return NewEmpty()
}

// DispatchPropertySet acts as a setter.
func (c *ADOXCatalog) DispatchPropertySet(propertyName string, args []Value) bool {
	if len(args) == 0 {
		return false
	}
	switch strings.ToLower(strings.TrimSpace(propertyName)) {
	case "activeconnection":
		c.activeConnection = args[0]
		c.tables = Value{Type: VTEmpty}
		return true
	}
	return false
}

// DispatchMethod handles methods for Catalog.
func (c *ADOXCatalog) DispatchMethod(methodName string, args []Value) Value {
	return NewEmpty()
}

func (c *ADOXCatalog) getTables() Value {
	if c.tables.Type != VTEmpty {
		return c.tables
	}

	c.tables = c.vm.newADOXTablesObject([]*ADOXTable{})
	return c.tables
}

// ADOXTables represents the catalog tables collection.
type ADOXTables struct {
	vm    *VM
	items []*ADOXTable
}

func (vm *VM) newADOXTablesObject(items []*ADOXTable) Value {
	obj := &ADOXTables{
		vm:    vm,
		items: items,
	}
	id := vm.nextDynamicNativeID
	vm.nextDynamicNativeID++
	vm.adoxTablesItems[id] = obj
	return Value{Type: VTNativeObject, Num: id}
}

func (t *ADOXTables) DispatchPropertyGet(name string) Value {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "count":
		return NewInteger(int64(len(t.items)))
	case "item":
		return NewEmpty()
	}
	return t.DispatchMethod(name, nil)
}

func (t *ADOXTables) DispatchMethod(name string, args []Value) Value {
	method := strings.ToLower(strings.TrimSpace(name))
	if method == "" || method == "item" {
		if len(args) < 1 {
			return NewEmpty()
		}

		val := args[0]
		switch val.Type {
		case VTInteger:
			idx := int(val.Num)
			if idx >= 0 && idx < len(t.items) {
				return t.vm.newADOXTableObject(t.items[idx])
			}
		case VTString:
			key := strings.ToLower(strings.TrimSpace(val.String()))
			for _, item := range t.items {
				if strings.ToLower(item.Name) == key {
					return t.vm.newADOXTableObject(item)
				}
			}
		}
		return NewEmpty()
	}

	return NewEmpty()
}

// ADOXTable represents a single ADOX table item.
type ADOXTable struct {
	Name string
	Type string
}

type ADOXTableWrapper struct {
	vm   *VM
	item *ADOXTable
}

func (vm *VM) newADOXTableObject(item *ADOXTable) Value {
	obj := &ADOXTableWrapper{
		vm:   vm,
		item: item,
	}
	id := vm.nextDynamicNativeID
	vm.nextDynamicNativeID++
	vm.adoxTableItems[id] = obj
	return Value{Type: VTNativeObject, Num: id}
}

func (t *ADOXTableWrapper) DispatchPropertyGet(name string) Value {
	switch strings.ToLower(strings.TrimSpace(name)) {
	case "name":
		return NewString(t.item.Name)
	case "type":
		return NewString(t.item.Type)
	}
	return NewEmpty()
}

func (t *ADOXTableWrapper) DispatchMethod(name string, args []Value) Value {
	return NewEmpty()
}
