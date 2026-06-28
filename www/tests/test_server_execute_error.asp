<%@ Language="VBScript" %>
<%
Response.Write "<h3>Server.Execute Error Integration Test</h3>"

' Enable error handling
On Error Resume Next

' Try to execute a non-existent file
Server.Execute "nonexistent_file_for_test.asp"

' Assert that an error was captured and has a non-zero number
If Err.Number <> 0 Then
    Response.Write "[SUCCESS] Non-existent file error captured: " & Err.Number & " - " & Err.Description & "<br>"
Else
    Response.Write "[FAIL] Error not captured, Err.Number is 0!<br>"
End If

' Reset error object
Err.Clear

' Try to execute a file with compilation/runtime error
' First, write a temporary file with a division by zero error
Dim fso, filePath, outStream
Set fso = Server.CreateObject("Scripting.FileSystemObject")
filePath = Server.MapPath("temp_child_error.asp")

' Write division by zero code
Set outStream = fso.CreateTextFile(filePath, True)
outStream.WriteLine("<" & "%")
outStream.WriteLine("dim y")
outStream.WriteLine("y = 1 / 0")
outStream.WriteLine("%" & ">")
outStream.Close()

' Execute the faulty file
Server.Execute "temp_child_error.asp"

' Assert that the division by zero error was propagated
If Err.Number = 11 Or Err.Number = -2146828277 Then
    Response.Write "[SUCCESS] Division by zero error propagated: " & Err.Number & " - " & Err.Description & "<br>"
Else
    Response.Write "[FAIL] Expected error 11 or HRESULT (division by zero), but got: " & Err.Number & " - " & Err.Description & "<br>"
End If

' Clean up temp file
If fso.FileExists(filePath) Then
    fso.DeleteFile(filePath)
End If
%>
