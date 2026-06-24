package axonvm

import (
	"bytes"
	"testing"
)

// TestCookieGitHubSubKeyEnum reproduces the exact ASP test scenario from
// www/tests/test_cookie_github.asp Test 3, where For Each is used on
// Request.Cookies(cn) when the cookie has sub-keys.
func TestCookieGitHubSubKeyEnum(t *testing.T) {
	source := `<%@LANGUAGE="VBSCRIPT" CODEPAGE="936"%>
<%Option Explicit%>
<%
' Create a main cookie named "user" containing multiple sub-keys
Response.Cookies("user")("firstname") = "John"
Response.Cookies("user")("lastname") = "Smith"
Response.Cookies("user")("country") = "Norway"
Response.Cookies("user")("age") = "25"

' Simulate the browser sending the cookie back on next request
' by directly populating Request.Cookies
Dim cn, kc, outStr
outStr = ""
For Each cn In Request.Cookies
    If Request.Cookies(cn).HasKeys Then
        For Each kc In Request.Cookies(cn)
            outStr = outStr & "[" & cn & "] " & kc & "=" & Request.Cookies(cn)(kc) & "|"
        Next
    End If
Next
Response.Write outStr
%>`

	compiler := NewASPCompiler(source)
	if err := compiler.Compile(); err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	vm := NewVM(compiler.Bytecode(), compiler.Constants(), compiler.GlobalsCount())
	host := NewMockHost()

	// Simulate browser sending the cookie with sub-keys
	host.Request().Cookies.AddCookie("user", "firstname=John&lastname=Smith&country=Norway&age=25")

	var output bytes.Buffer
	host.SetOutput(&output)
	vm.SetHost(host)

	if err := vm.Run(); err != nil {
		t.Fatalf("vm run failed: %v", err)
	}
	host.Response().Flush()

	actual := output.String()
	// Keys sorted: age, country, firstname, lastname
	const want = "[user] age=25|[user] country=Norway|[user] firstname=John|[user] lastname=Smith|"
	if actual != want {
		t.Fatalf("unexpected output:\n  got:  %q\n  want: %q", actual, want)
	}
	t.Logf("Output: %q", actual)
}
