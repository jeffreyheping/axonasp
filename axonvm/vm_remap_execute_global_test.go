package axonvm

import (
	"encoding/binary"
	"testing"
)

// TestRemapExecuteGlobalBytecode_ObjectRestKeepsAlignment verifies remap scanning
// remains instruction-aligned when a variable-length OpJSObjectRest precedes a
// remapped constant-index opcode.
func TestRemapExecuteGlobalBytecode_ObjectRestKeepsAlignment(t *testing.T) {
	bytecode := []byte{
		byte(OpJSObjectRest),
		0x00, 0x01, // static key count = 1
		0xFE, 0x00, // static key const index (chosen to expose old misaligned size read)
		0x00, 0x00, // dynamic key count = 0
		byte(OpJSIncLocalInt),
		0x00, 0x02, // name const index to be remapped
		byte(OpHalt),
	}

	constBase := 5
	remapExecuteGlobalBytecode(bytecode, constBase, 0)

	got := int(binary.BigEndian.Uint16(bytecode[8:10]))
	want := 2 + constBase
	if got != want {
		t.Fatalf("expected OpJSIncLocalInt const index remapped to %d, got %d", want, got)
	}
}
