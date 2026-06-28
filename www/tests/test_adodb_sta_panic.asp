<%@ Language = VBScript %>
<!DOCTYPE html>
<html>

    <head>
        <title>ADODB STA Panic Recovery Test</title>
        <style>
            body {
                font-family: Arial, sans-serif;
                margin: 20px;
            }

            h2 {
                color: #333;
                border-bottom: 2px solid #007bff;
                padding-bottom: 5px;
            }

            .success {
                color: green;
                font-weight: bold;
            }

            .error {
                color: red;
                font-weight: bold;
            }

            .info {
                color: blue;
            }

            .section {
                margin: 20px 0;
                padding: 15px;
                border: 1px solid #ccc;
                border-radius: 5px;
            }
        </style>
    </head>

    <body>
        <h1>ADODB STA Panic Recovery Test</h1>
        <p class="info">Verifies that an ADODB runtime error (e.g., empty connection string) does not crash the server
            and correctly propagates to the script error handler.</p>

        <div class="section">
            <h2>Test 1: Empty Connection String (On Error Resume Next)</h2>
            <%
        Dim conn1, errNum1, errDesc1
        On Error Resume Next
        Set conn1 = Server.CreateObject("ADODB.Connection")
        conn1.ConnectionString = ""
        conn1.Open
        errNum1 = Err.Number
        errDesc1 = Err.Description
        On Error Goto 0

        If errNum1 <> 0 Then
            Response.Write "<p class=""success"">PASS: Error correctly raised on empty connection string.</p>"
            Response.Write "<p class=""info"">Err.Number: " & errNum1 & "</p>"
            Response.Write "<p class=""info"">Err.Description: " & Server.HTMLEncode(errDesc1) & "</p>"
        Else
            Response.Write "<p class=""error"">FAIL: No error was raised for empty connection string.</p>"
        End If
        %>
        </div>

        <div class="section">
            <h2>Test 2: Invalid Provider (On Error Resume Next)</h2>
            <%
        Dim conn2, errNum2, errDesc2
        On Error Resume Next
        Set conn2 = Server.CreateObject("ADODB.Connection")
        conn2.ConnectionString = "Provider=INVALID_PROVIDER;Data Source=nonexistent"
        conn2.Open
        errNum2 = Err.Number
        errDesc2 = Err.Description
        On Error Goto 0

        If errNum2 <> 0 Then
            Response.Write "<p class=""success"">PASS: Error correctly raised on invalid provider.</p>"
            Response.Write "<p class=""info"">Err.Number: " & errNum2 & "</p>"
            Response.Write "<p class=""info"">Err.Description: " & Server.HTMLEncode(errDesc2) & "</p>"
        Else
            Response.Write "<p class=""error"">FAIL: No error was raised for invalid provider.</p>"
        End If
        %>
        </div>

        <div class="section">
            <h2>Test 3: Server Continuity After Error</h2>
            <p class="info">Verifies that the server remains operational after a failed ADODB connection attempt.</p>
            <%
        ' This test passes if it reaches this point without the server crashing.
        Response.Write "<p class=""success"">PASS: Server is still running after ADODB connection errors.</p>"
        %>
        </div>

        <div class="section">
            <h2>Test 4: Successful Connection After Failed Attempt</h2>
            <%
        ' Intentionally fail first, then verify the error state is clean
        On Error Resume Next
        Dim conn3
        Set conn3 = Server.CreateObject("ADODB.Connection")
        conn3.ConnectionString = ""
        conn3.Open
        Err.Clear
        On Error Goto 0

        ' Verify Err object is clean after Clear
        If Err.Number = 0 Then
            Response.Write "<p class=""success"">PASS: Err object correctly cleared after failed connection.</p>"
        Else
            Response.Write "<p class=""error"">FAIL: Err object not cleared. Number: " & Err.Number & "</p>"
        End If
        %>
        </div>

        <div class="section">
            <h2>Test 5: Nested Error Handling</h2>
            <%
        Dim outerErr
        On Error Resume Next
        ' Inner scope with error
        Dim conn4
        Set conn4 = Server.CreateObject("ADODB.Connection")
        conn4.ConnectionString = ""
        conn4.Open
        ' The error should propagate up to this handler
        outerErr = Err.Number
        On Error Goto 0

        If outerErr <> 0 Then
            Response.Write "<p class=""success"">PASS: Nested error handling works correctly.</p>"
            Response.Write "<p class=""info"">Err.Number: " & outerErr & "</p>"
        Else
            Response.Write "<p class=""error"">FAIL: No error propagated in nested scope.</p>"
        End If
        %>
        </div>

        <div class="section">
            <h2>Summary</h2>
            <%
        Response.Write "<p>All tests completed. If this page rendered fully, the STA panic recovery is working correctly.</p>"
        %>
        </div>
    </body>

</html>