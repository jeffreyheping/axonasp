//go:build lib_msxml_disabled

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

// MsXML2ServerXMLHTTP is the disabled placeholder for MSXML HTTP object.
type MsXML2ServerXMLHTTP struct{ ctx *VM }

// MsXML2DOMDocument is the disabled placeholder for MSXML DOM document.
type MsXML2DOMDocument struct{ ctx *VM }

// XMLNodeList is the disabled placeholder for MSXML node-list objects.
type XMLNodeList struct{ ctx *VM }

// ParseError is the disabled placeholder for MSXML parse-error objects.
type ParseError struct {
	ctx         *VM
	ErrorCode   int
	ErrorReason string
	FilePos     int
	Line        int
	LinePos     int
	SrcText     string
	URL         string
}

// XMLElement is the disabled placeholder for MSXML element objects.
type XMLElement struct{ ctx *VM }

func NewMsXML2ServerXMLHTTP(ctx *VM) *MsXML2ServerXMLHTTP {
	panicLibraryDisabled("msxml", "MSXML2.ServerXMLHTTP")
	return nil
}

func NewMsXML2DOMDocument(ctx *VM) *MsXML2DOMDocument {
	panicLibraryDisabled("msxml", "MSXML2.DOMDocument")
	return nil
}

func (s *MsXML2ServerXMLHTTP) legacyGetProperty(name string) interface{}              { return nil }
func (s *MsXML2ServerXMLHTTP) legacySetProperty(name string, value interface{}) error { return nil }
func (s *MsXML2ServerXMLHTTP) legacyCallMethod(name string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

func (d *MsXML2DOMDocument) legacyGetProperty(name string) interface{}              { return nil }
func (d *MsXML2DOMDocument) legacySetProperty(name string, value interface{}) error { return nil }
func (d *MsXML2DOMDocument) legacyCallMethod(name string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

func (l *XMLNodeList) legacyGetProperty(name string) interface{}              { return nil }
func (l *XMLNodeList) legacySetProperty(name string, value interface{}) error { return nil }
func (l *XMLNodeList) legacyCallMethod(name string, args ...interface{}) (interface{}, error) {
	return nil, nil
}
func (l *XMLNodeList) Enumeration() []interface{} { return nil }

func (p *ParseError) legacyGetProperty(name string) interface{}              { return nil }
func (p *ParseError) legacySetProperty(name string, value interface{}) error { return nil }
func (p *ParseError) legacyCallMethod(name string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

func (e *XMLElement) legacyGetProperty(name string) interface{}              { return nil }
func (e *XMLElement) legacySetProperty(name string, value interface{}) error { return nil }
func (e *XMLElement) legacyCallMethod(name string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

func (x *MsXML2ServerXMLHTTP) DispatchPropertyGet(name string) Value              { return Value{Type: VTEmpty} }
func (x *MsXML2ServerXMLHTTP) DispatchPropertySet(name string, args []Value) bool { return false }
func (x *MsXML2ServerXMLHTTP) DispatchMethod(name string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (x *MsXML2DOMDocument) DispatchPropertyGet(name string) Value              { return Value{Type: VTEmpty} }
func (x *MsXML2DOMDocument) DispatchPropertySet(name string, args []Value) bool { return false }
func (x *MsXML2DOMDocument) DispatchMethod(name string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (x *XMLNodeList) DispatchPropertyGet(name string) Value              { return Value{Type: VTEmpty} }
func (x *XMLNodeList) DispatchPropertySet(name string, args []Value) bool { return false }
func (x *XMLNodeList) DispatchMethod(name string, args []Value) Value     { return Value{Type: VTEmpty} }

func (x *ParseError) DispatchPropertyGet(name string) Value              { return Value{Type: VTEmpty} }
func (x *ParseError) DispatchPropertySet(name string, args []Value) bool { return false }
func (x *ParseError) DispatchMethod(name string, args []Value) Value     { return Value{Type: VTEmpty} }

func (x *XMLElement) DispatchPropertyGet(name string) Value              { return Value{Type: VTEmpty} }
func (x *XMLElement) DispatchPropertySet(name string, args []Value) bool { return false }
func (x *XMLElement) DispatchMethod(name string, args []Value) Value     { return Value{Type: VTEmpty} }
