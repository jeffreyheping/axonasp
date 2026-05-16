<%@ Language="JScript" %>
<%
Response.Write("<h2>Symbol.unscopables Test</h2>");
var obj = { a: 1, b: 2 };
obj[Symbol.unscopables] = { a: true };
var a = 10, b = 20;
with (obj) {
    Response.Write("a=" + a + " (expected 10)<br>");
    Response.Write("b=" + b + " (expected 2)<br>");
}

Response.Write("<h2>Symbol.matchAll Test</h2>");
var str = "test1test2";
var regex = /test(\d)/g;
var matches = str.matchAll(regex);
Response.Write("Type of matches: " + typeof matches + "<br>");
var res = "";
for (var m of matches) {
    Response.Write("Match: " + m[0] + " Group: " + m[1] + "<br>");
    res += m[1] + ",";
}
Response.Write("Result: " + res + " (expected 1,2,)<br>");

Response.Write("<h2>RegExp.prototype[Symbol.matchAll] Test</h2>");
var regex2 = /[a-z]/g;
Response.Write("RegExp exists: " + (typeof RegExp !== 'undefined') + "<br>");
Response.Write("RegExp.prototype[Symbol.matchAll]: " + typeof RegExp.prototype[Symbol.matchAll] + "<br>");
Response.Write("Symbol exists: " + (typeof Symbol !== 'undefined') + "<br>");
Response.Write("Symbol.matchAll: " + (Symbol.matchAll ? "EXISTS" : "MISSING") + "<br>");
Response.Write("Type of [Symbol.matchAll] on regex2: " + typeof regex2[Symbol.matchAll] + "<br>");
var it = regex2[Symbol.matchAll]("abc");
Response.Write("Type of it: " + typeof it + "<br>");
for (var x of it) {
    Response.Write("X: " + x[0] + "<br>");
}

Response.Write("<br><strong>Tests Completed.</strong>");
%>
