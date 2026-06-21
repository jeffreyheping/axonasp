<%@ Language=VBScript %>
<%
' Test: Date literal US parsing and locale-aware string formatting
' This test verifies:
' 1. Date literals (#...#) are always interpreted as US format
' 2. CStr(date) uses locale-aware formatting (not RFC3339)
' 3. Implicit string coercion (&) uses locale-aware formatting
' 4. Date-only and time-only values format correctly

Dim passedCount, failCount
passedCount = 0
failCount = 0

Sub AssertEqual(description, expected, actual)
    If expected = actual Then
        passedCount = passedCount + 1
    Else
        failCount = failCount + 1
        Response.Write "FAIL: " & description & " - expected: '" & expected & "', got: '" & actual & "'<br>"
    End If
End Sub

Sub AssertContains(description, haystack, needle)
    If InStr(1, haystack, needle, vbTextCompare) > 0 Then
        passedCount = passedCount + 1
    Else
        failCount = failCount + 1
        Response.Write "FAIL: " & description & " - expected '" & haystack & "' to contain '" & needle & "'<br>"
    End If
End Sub

Sub AssertNotContains(description, haystack, needle)
    If InStr(1, haystack, needle, vbTextCompare) = 0 Then
        passedCount = passedCount + 1
    Else
        failCount = failCount + 1
        Response.Write "FAIL: " & description & " - '" & haystack & "' should NOT contain '" & needle & "'<br>"
    End If
End Sub

' ============================================
' TEST 1: Date literals are always US format
' ============================================
Response.Write "<h2>Test 1: Date Literal US Parsing</h2>"

Dim d1, d2, d3

' #3/4/2026# must be March 4, 2026 (US), NOT April 3
d1 = #3/4/2026#
AssertEqual "Date literal #3/4/2026# Month is March", 3, Month(d1)
AssertEqual "Date literal #3/4/2026# Day is 4", 4, Day(d1)
AssertEqual "Date literal #3/4/2026# Year is 2026", 2026, Year(d1)

' #12/31/2025# must be December 31
d2 = #12/31/2025#
AssertEqual "Date literal #12/31/2025# Month is December", 12, Month(d2)
AssertEqual "Date literal #12/31/2025# Day is 31", 31, Day(d2)

' #1/2/2023# must be January 2
d3 = #1/2/2023#
AssertEqual "Date literal #1/2/2023# Month is January", 1, Month(d3)
AssertEqual "Date literal #1/2/2023# Day is 2", 2, Day(d3)

' ISO format #2026-03-04#
Dim d4
d4 = #2026-03-04#
AssertEqual "Date literal #2026-03-04# Month is March", 3, Month(d4)
AssertEqual "Date literal #2026-03-04# Day is 4", 4, Day(d4)

' Date with time
Dim d5
d5 = #3/4/2026 14:30:00#
AssertEqual "Date literal #3/4/2026 14:30:00# Hour is 14", 14, Hour(d5)
AssertEqual "Date literal #3/4/2026 14:30:00# Minute is 30", 30, Minute(d5)

' ============================================
' TEST 2: CStr(date) uses locale formatting
' ============================================
Response.Write "<h2>Test 2: CStr(date) Locale Formatting (US LCID 1033)</h2>"

Response.LCID = 1033

Dim testDate, strResult
testDate = #3/4/2026 14:30:00#
strResult = CStr(testDate)

' Must NOT be RFC3339 format (no T or Z)
AssertNotContains "CStr(date) does not use RFC3339 (T)", strResult, "T"
AssertNotContains "CStr(date) does not use RFC3339 (Z)", strResult, "Z"

' Must contain the year
AssertContains "CStr(date) contains year 2026", strResult, "2026"

' US locale should have M/D/YYYY
AssertContains "CStr(date) contains US date format with /", strResult, "/"

' ============================================
' TEST 3: Implicit string coercion via &
' ============================================
Response.Write "<h2>Test 3: Implicit Date String Coercion</h2>"

Response.LCID = 1033

Dim concatResult
concatResult = #3/4/2026# & ""

AssertNotContains "Date concatenation does not use RFC3339 (T)", concatResult, "T"
AssertContains "Date concatenation contains year", concatResult, "2026"

' Date-only (midnight) should NOT show time
Dim dateOnlyStr
dateOnlyStr = #3/4/2026# & ""
AssertNotContains "Date-only (midnight) should omit time separator", dateOnlyStr, ":"

' ============================================
' TEST 4: DateSerial (date-only) and TimeSerial (time-only)
' ============================================
Response.Write "<h2>Test 4: DateSerial / TimeSerial Formatting</h2>"

Response.LCID = 1033

Dim ds, dsStr, ts, tsStr
ds = DateSerial(2026, 3, 4)
dsStr = CStr(ds)
AssertContains "DateSerial string contains year", dsStr, "2026"
AssertNotContains "DateSerial string should omit time", dsStr, ":"

ts = TimeSerial(14, 30, 0)
tsStr = CStr(ts)
AssertNotContains "TimeSerial string should omit 1899", tsStr, "1899"
AssertContains "TimeSerial string contains time", tsStr, ":"

' ============================================
' TEST 5: Different LCID produces different formats
' ============================================
Response.Write "<h2>Test 5: Multi-Locale Date Formatting</h2>"

Dim usDateStr, brDateStr

Response.LCID = 1033
usDateStr = CStr(#3/4/2026#)
Response.Write "US (1033): " & usDateStr & "<br>"

Response.LCID = 1046
brDateStr = CStr(#3/4/2026#)
Response.Write "pt-BR (1046): " & brDateStr & "<br>"

' Brazilian format uses DD/MM/YYYY, so it should differ from US
If usDateStr <> brDateStr Then
    passedCount = passedCount + 1
    Response.Write "PASS: Different locales produce different date strings<br>"
Else
    failCount = failCount + 1
    Response.Write "FAIL: Expected different date strings for US and pt-BR locales<br>"
End If

' ============================================
' TEST 6: Response.Write uses locale formatting
' ============================================
Response.Write "<h2>Test 6: Response.Write with Date Value</h2>"

Response.LCID = 1033
Dim writeOutput
' We can't easily capture Response.Write output, but we can verify it doesn't error
On Error Resume Next
Response.Write #3/4/2026#
If Err.Number = 0 Then
    passedCount = passedCount + 1
    Response.Write " (no error - PASS)<br>"
Else
    failCount = failCount + 1
    Response.Write " (ERROR: " & Err.Description & " - FAIL)<br>"
End If
On Error GoTo 0

' ============================================
' SUMMARY
' ============================================
Response.Write "<h2>Results</h2>"
Response.Write "Passed: " & passedCount & "<br>"
Response.Write "Failed: " & failCount & "<br>"

If failCount = 0 Then
    Response.Write "<strong>ALL TESTS PASSED</strong>"
Else
    Response.Write "<strong>SOME TESTS FAILED</strong>"
End If
%>