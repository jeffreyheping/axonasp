<%
    ' Phase 3: Native File I/O Test
    
    Dim fNum, path, content
    fNum = FreeFile
    path = Server.MapPath("test_file_io.txt")
    
    Response.Write "Opening file for output...<br>"
    Open path For Output As #fNum
    Print #fNum, "AxonASP VB6 File I/O"
    Print #fNum, "Line 2"
    Close #fNum
    Response.Write "File written and closed.<br>"
    
    Response.Write "Opening file for input...<br>"
    Dim s1, s2
    Open path For Input As #fNum
    Line Input #fNum, s1
    Line Input #fNum, s2
    Close #fNum
    Response.Write "File read and closed.<br>"
    
    Response.Write "Content read:<br>"
    Response.Write "S1: " & s1 & "<br>"
    Response.Write "S2: " & s2 & "<br>"
    
    If s1 = "AxonASP VB6 File I/O" And s2 = "Line 2" Then
        Response.Write "<b>TEST PASSED</b>"
    Else
        Response.Write "<b>TEST FAILED</b>"
    End If
%>