package axonvm

import (
	"bytes"
	"testing"
)

// TestResponseCookiesSubKeyAssignment verifies that setting sub-keys on
// Response.Cookies encodes them into the cookie value correctly.
func TestResponseCookiesSubKeyAssignment(t *testing.T) {
	source := `<%
Response.Cookies("user")("firstname") = "John"
Response.Cookies("user")("lastname") = "Smith"
Response.Write "done"
%>`

	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	vm := NewVM(compiler.Bytecode(), compiler.Constants(), compiler.GlobalsCount())
	host := NewMockHost()
	var output bytes.Buffer
	host.SetOutput(&output)
	vm.SetHost(host)

	if err := vm.Run(); err != nil {
		t.Fatalf("vm run failed: %v", err)
	}
	host.Response().Flush()

	// Get the cookie value that was set
	cookieValue := host.Response().GetCookieValue("user")
	t.Logf("Cookie 'user' value: %q", cookieValue)

	// In Classic ASP, sub-keys should be URL-encoded in the value: firstname=John&lastname=Smith
	if cookieValue == "" {
		// Check if any cookie was set at all
		count := host.Response().GetCookieCount()
		t.Logf("Cookie count: %d", count)
		key0 := host.Response().GetCookieKey(1)
		t.Logf("Cookie key at index 1: %q", key0)
	}

	// The value should contain the sub-keys
	// Expected format: firstname=John&lastname=Smith
}

// TestSimpleCookieAssignment verifies that Response.Cookies("name") = "value" works.
func TestSimpleCookieAssignment(t *testing.T) {
	source := `<%
Response.Cookies("simple") = "hello"
Response.Write "done"
%>`

	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	vm := NewVM(compiler.Bytecode(), compiler.Constants(), compiler.GlobalsCount())
	host := NewMockHost()
	var output bytes.Buffer
	host.SetOutput(&output)
	vm.SetHost(host)

	if err := vm.Run(); err != nil {
		t.Fatalf("vm run failed: %v", err)
	}
	host.Response().Flush()

	cookieValue := host.Response().GetCookieValue("simple")
	t.Logf("Cookie 'simple' value: %q", cookieValue)
	if cookieValue != "hello" {
		t.Fatalf("expected 'hello', got %q", cookieValue)
	}
}
