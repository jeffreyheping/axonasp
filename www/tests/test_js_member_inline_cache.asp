<%@ Language="JScript" %>
<%
var o = { a: 1, b: 2 };
var sum = 0;
for (var i = 0; i < 8; i++) {
  sum += o.a;
  o.a = o.a + 1;
}
Response.Write("IC:" + o.a + "," + sum);
%>