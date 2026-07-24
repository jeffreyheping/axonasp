<%
Option Explicit
Response.Buffer = True

Class Singleton
    Private Sub Class_Initialize()
        Response.Write "InitSingleton;"
    End Sub

    Public Function CreateInternal()
        Set CreateInternal = New Singleton
    End Function
End Class

Class PublicClass
    Public Sub Class_Initialize()
        Response.Write "InitPublic;"
    End Sub

    Public Function TryCreateSingleton()
        On Error Resume Next
        Dim s
        Set s = New Singleton
        If Err.Number <> 0 Then
            Response.Write "FactoryErr:" & Err.Number & ";"
        End If
        Set TryCreateSingleton = s
    End Function
End Class

On Error Resume Next
Dim obj, pubObj, sObj

Response.Write "1. External instantiation: "
Set obj = New Singleton
If Err.Number <> 0 Then
    Response.Write "Err:" & Err.Number & ";"
End If
If IsEmpty(obj) Then
    Response.Write "IsEmpty;"
End If
Err.Clear

Response.Write "<br>2. Public instantiation: "
Set pubObj = New PublicClass
If Not pubObj Is Nothing Then
    Response.Write "SuccessPublic;"
End If

Response.Write "<br>3. Factory external instantiation: "
Set sObj = pubObj.TryCreateSingleton()
If IsEmpty(sObj) Then
    Response.Write "FactoryIsEmpty;"
End If

Response.Write "<br>Done"
%>
