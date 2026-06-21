<%@ Language=VBScript %>
<%
' Test: OpLt (<) comparison operator semantics
' Regression test for the fix that replaced asFloat with compareValues
' Covers: Null propagation, string comparison, numeric comparison,
' mixed types, Empty coercion, and date comparison.

Dim testCount, passCount
testCount = 0
passCount = 0

Sub RunNullTest(name, actual)
	testCount = testCount + 1
	If IsNull(actual) Then
		passCount = passCount + 1
		Response.Write "<div class='pass'>PASS: " & name & "</div>"
	Else
		Response.Write "<div class='fail'>FAIL: " & name & " — expected Null, got '" & TypeName(actual) & "'</div>"
	End If
End Sub

Sub RunBoolTest(name, expected, actual)
	testCount = testCount + 1
	If CBool(actual) = CBool(expected) Then
		passCount = passCount + 1
		Response.Write "<div class='pass'>PASS: " & name & "</div>"
	Else
		Response.Write "<div class='fail'>FAIL: " & name & " — expected '" & expected & "', got '" & actual & "'</div>"
	End If
End Sub

Dim x, result

' --- Test 1: Null < number ---
x = Null
result = x < 5
RunNullTest "Null < 5 = Null", result

' --- Test 2: Number < Null ---
result = 5 < x
RunNullTest "5 < Null = Null", result

' --- Test 3: Null < Null ---
result = x < x
RunNullTest "Null < Null = Null", result

' --- Test 4: String comparison "b" < "a" ---
result = "b" < "a"
RunBoolTest """b"" < ""a"" = False", False, result

' --- Test 5: String comparison "a" < "b" ---
result = "a" < "b"
RunBoolTest """a"" < ""b"" = True", True, result

' --- Test 6: Numeric comparison 2 < 3 ---
result = 2 < 3
RunBoolTest "2 < 3 = True", True, result

' --- Test 7: Numeric comparison 5 < 3 ---
result = 5 < 3
RunBoolTest "5 < 3 = False", False, result

' --- Test 8: Numeric comparison 3 < 3 ---
result = 3 < 3
RunBoolTest "3 < 3 = False", False, result

' --- Test 9: Mixed numeric-string "3" < 5 ---
result = "3" < 5
RunBoolTest """3"" < 5 = True", True, result

' --- Test 10: Mixed numeric-string 5 < "3" ---
result = 5 < "3"
RunBoolTest "5 < ""3"" = False", False, result

' --- Test 11: Empty < 1 (Empty coerces to 0) ---
Dim e
e = Empty
result = e < 1
RunBoolTest "Empty < 1 = True", True, result

' --- Test 12: Empty < 0 ---
result = e < 0
RunBoolTest "Empty < 0 = False", False, result

' --- Test 13: Date comparison ---
result = #2025-01-01# < #2026-01-01#
RunBoolTest "#2025-01-01# < #2026-01-01# = True", True, result

' --- Test 14: Chained comparisons ---
result = 1 < 2 And 2 < 3
RunBoolTest "1 < 2 And 2 < 3 = True", True, result

' --- Test 15: Null in chained AND ---
result = (x < 5) And True
RunNullTest "(Null < 5) And True = Null", result

%>
<!DOCTYPE html>
<html>

    <head>
        <title>OpLt (&lt;) Comparison Regression Test</title>
        <style>
            body {
                font-family: monospace;
                margin: 20px;
                background: #1e1e1e;
                color: #d4d4d4;
            }

            .pass {
                color: #4ec9b0;
                padding: 2px 0;
            }

            .fail {
                color: #f44747;
                padding: 2px 0;
            }

            .summary {
                margin-top: 20px;
                font-weight: bold;
                font-size: 1.2em;
            }

            .summary.pass {
                color: #4ec9b0;
            }

            .summary.fail {
                color: #f44747;
            }
        </style>
    </head>

    <body>
        <h1>OpLt (&lt;) Comparison Regression Test</h1>
        <div class='summary <%= IIf(passCount = testCount, "pass", "fail") %>'>
            <%= passCount %> / <%= testCount %> tests passed
            <% If passCount <> testCount Then %>
            — <%= testCount - passCount %> FAILED
            <% End If %>
        </div>
    </body>

</html>