//go:build lib_g3crypto_disabled

/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 */
package axonvm

// G3Crypto is the disabled stub for the G3Crypto library.
type G3Crypto struct{}

func NewG3Crypto() *G3Crypto {
	panicLibraryDisabled("g3crypto", "G3Crypto library")
	return nil
}

func NewG3CryptoWithAlgorithm(algorithm string) *G3Crypto {
	panicLibraryDisabled("g3crypto", "G3Crypto library")
	return nil
}

func (c *G3Crypto) DispatchMethod(methodName string, args []Value) Value {
	return Value{Type: VTEmpty}
}

func (c *G3Crypto) DispatchPropertyGet(propertyName string) Value {
	return Value{Type: VTEmpty}
}

func (c *G3Crypto) DispatchPropertySet(propertyName string, val Value) {}

func g3cryptoResolveProgID(progID string) (string, bool) {
	return "", false
}
