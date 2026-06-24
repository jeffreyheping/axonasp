<%@LANGUAGE="VBSCRIPT" CODEPAGE="936"%>
<%Option Explicit%>
<%
' Create a main cookie named "user" containing multiple sub-keys
Response.Cookies("user")("firstname") = "John"
Response.Cookies("user")("lastname") = "Smith"
Response.Cookies("user")("country") = "Norway"
Response.Cookies("user")("age") = "25"
%>
<html>

    <body>
        <h2>Test 1: For Each on request.Cookies (works)</h2>
        <table border="1">
            <%
dim cn
for each cn in request.Cookies
    response.Write("<tr><td>" & cn & "</td><td>" & request.Cookies(cn) & "</td></tr>")
next
%>
        </table>

        <h2>Test 2: HasKeys on request.Cookies(cn) (now works, thanks!)</h2>
        <table border="1">
            <%
for each cn in request.Cookies
    if request.Cookies(cn).HasKeys then
        response.Write("<tr><td>" & cn & "</td><td>Has sub-keys</td></tr>")
    else
        response.Write("<tr><td>" & cn & "</td><td>No sub-keys</td></tr>")
    end if
next
%>
        </table>

        <h2>Test 3: For Each on request.Cookies(cn) (now works, thanks!)</h2>
        <table border="1">
            <%
for each cn in request.Cookies
    if request.Cookies(cn).HasKeys then
        dim kc
        ' This line throws a runtime error:
        for each kc in request.Cookies(cn)
            response.Write("<tr><td>[" & cn & "] " & kc & "</td><td>" & request.Cookies(cn)(kc) & "</td></tr>")
        next
    end if
next
%>
        </table>
    </body>

</html>