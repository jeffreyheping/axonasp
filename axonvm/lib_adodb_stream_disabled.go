//go:build lib_adodb_stream_disabled

/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 */
package axonvm

const (
	adodbTypeBinary = 1
	adodbTypeText   = 2

	adodbStateClosed = 0
	adodbStateOpen   = 1

	adodbLineLF   = 10
	adodbLineCR   = 13
	adodbLineCRLF = -1
)

// adodbStreamNativeObject is the disabled ADODB.Stream runtime placeholder.
type adodbStreamNativeObject struct {
	typ           int
	state         int
	position      int64
	size          int64
	buffer        []byte
	charset       string
	lineSeparator int
}

// newADODBStreamObject fails because ADODB.Stream is disabled at compile time.
func (vm *VM) newADODBStreamObject() Value {
	panicLibraryDisabled("adodb_stream", "ADODB.Stream")
	return Value{Type: VTEmpty}
}

// dispatchADODBStreamMethod is unhandled when ADODB.Stream is disabled.
func (vm *VM) dispatchADODBStreamMethod(objID int64, member string, args []Value) (Value, bool) {
	return Value{Type: VTEmpty}, false
}

// dispatchADODBStreamPropertyGet is unhandled when ADODB.Stream is disabled.
func (vm *VM) dispatchADODBStreamPropertyGet(objID int64, member string) (Value, bool) {
	return Value{Type: VTEmpty}, false
}

// dispatchADODBStreamPropertySet is unhandled when ADODB.Stream is disabled.
func (vm *VM) dispatchADODBStreamPropertySet(objID int64, member string, val Value) bool {
	return false
}
