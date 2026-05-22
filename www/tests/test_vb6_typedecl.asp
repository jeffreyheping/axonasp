<% @ Language = "VBScript" %>
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>VB6 As Type Declaration Test</title>
    <link rel="stylesheet" href="../css/axonasp.css">
</head>
<body>
<div id="content">
    <h1>VB6 As Type Declaration Test</h1>
    <p>Tests the VB6-style strong typing support (Dim x As Integer) in AxonASP.</p>

    <%
    Dim passCount
    Dim failCount
    passCount = 0
    failCount = 0

    Sub PrintResult(testName, condition)
        If condition Then
            Response.Write "<div class='alert alert-success'><span class='status-v'>PASS:</span> " & testName & "</div>"
            passCount = passCount + 1
        Else
            Response.Write "<div class='alert alert-error'><span class='status-x'>FAIL:</span> " & testName & "</div>"
            failCount = failCount + 1
        End If
    End Sub
    %>

    <div class="section">
        <h2>Integer Type</h2>
        <div class="results">
        <%
        On Error Resume Next
        Dim i As Integer
        i = 42
        PrintResult "Dim i As Integer; i = 42 -> " & i, (i = 42)

        Dim i2 As Integer
        PrintResult "Dim i2 As Integer (default zero): " & i2, (i2 = 0)

        ' Empty/Nothing assignment to Integer should produce 0
        Dim i3 As Integer
        i3 = Empty
        PrintResult "Dim i3 As Integer; i3 = Empty -> " & i3, (i3 = 0)
        On Error GoTo 0
        %>
        </div>
    </div>

    <div class="section">
        <h2>String Type</h2>
        <div class="results">
        <%
        On Error Resume Next
        Dim s As String
        s = "hello"
        PrintResult "Dim s As String; s = ""hello"" -> " & s, (s = "hello")

        Dim s2 As String
        PrintResult "Dim s2 As String (default empty): len=" & Len(s2), (Len(s2) = 0)

        ' Integer can coerce to String
        Dim s3 As String
        s3 = 123
        PrintResult "Dim s3 As String; s3 = 123 -> " & s3, (s3 = "123")
        On Error GoTo 0
        %>
        </div>
    </div>

    <div class="section">
        <h2>Boolean Type</h2>
        <div class="results">
        <%
        On Error Resume Next
        Dim b As Boolean
        b = True
        PrintResult "Dim b As Boolean; b = True -> " & b, (b = True)

        Dim b2 As Boolean
        PrintResult "Dim b2 As Boolean (default False): " & b2, (b2 = False)

        b2 = 1
        PrintResult "Dim b2 As Boolean; b2 = 1 -> " & b2, (b2 = True)

        b2 = 0
        PrintResult "Dim b2 As Boolean; b2 = 0 -> " & b2, (b2 = False)
        On Error GoTo 0
        %>
        </div>
    </div>

    <div class="section">
        <h2>Double Type</h2>
        <div class="results">
        <%
        On Error Resume Next
        Dim d As Double
        d = 3.14159
        PrintResult "Dim d As Double; d = 3.14159 -> " & d, (d > 3.14 And d < 3.15)

        Dim d2 As Double
        PrintResult "Dim d2 As Double (default 0): " & d2, (d2 = 0)

        ' Integer assignment to Double coerces
        d2 = 99
        PrintResult "Dim d2 As Double; d2 = 99 -> " & d2, (d2 = 99)
        On Error GoTo 0
        %>
        </div>
    </div>

    <div class="section">
        <h2>Type Mismatch Error Handling</h2>
        <div class="results">
        <%
        On Error Resume Next
        Dim xi As Integer
        xi = "not-a-number"
        Dim errNum1
        errNum1 = Err.Number
        Dim errDesc1
        errDesc1 = Err.Description
        Err.Clear

        PrintResult "Type mismatch error on invalid coercion to Integer", (errNum1 <> 0)
        If errNum1 <> 0 Then
            Response.Write "<div class='alert alert-info'>Error number: " & errNum1 & " - " & errDesc1 & "</div>"
        End If
        On Error GoTo 0
        %>
        </div>
    </div>

    <div class="section">
        <h2>Multiple Variables on One Dim</h2>
        <div class="results">
        <%
        On Error Resume Next
        Dim a As Integer, b As String
        a = 7
        b = "text"
        PrintResult "Dim a As Integer, b As String; a=7, b=""text"" -> " & a & "|" & b, (a = 7 And b = "text")
        On Error GoTo 0
        %>
        </div>
    </div>

    <div class="section">
        <h2>Variant (No As Clause) - Backward Compat</h2>
        <div class="results">
        <%
        On Error Resume Next
        Dim v
        v = "can hold any type"
        PrintResult "Dim v (Variant); v = ""string"" -> " & v, (v = "can hold any type")

        v = 999
        PrintResult "Dim v (Variant); v = 999 -> " & v, (v = 999)

        v = True
        PrintResult "Dim v (Variant); v = True -> " & v, (v = True)
        On Error GoTo 0
        %>
        </div>
    </div>

    <div class="section">
        <h2>Summary</h2>
        <div class="results">
        <%
        Response.Write "<p>Tests passed: <strong>" & passCount & "</strong></p>"
        Response.Write "<p>Tests failed: <strong>" & failCount & "</strong></p>"
        If failCount = 0 Then
            Response.Write "<div class='alert alert-success'>All tests passed.</div>"
        Else
            Response.Write "<div class='alert alert-error'>Some tests failed.</div>"
        End If
        %>
        </div>
    </div>
</div>
</body>
</html>
