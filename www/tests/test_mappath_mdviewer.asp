<%@ Language="JavaScript" %>
<%
var p = Server.MapPath("test.md");
var files = Server.CreateObject("G3FILES");
Response.Write("MapPath: " + p + "<br>");
Response.Write("Exists: " + files.Exists(p) + "<br>");
%>
