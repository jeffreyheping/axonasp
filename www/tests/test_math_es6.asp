<%
@ Language = "JScript"
%>
<script runat="server" language="JScript">
    Response.Write("acosh=" + Math.acosh(1) + "\n");
    Response.Write("asinh=" + Math.asinh(0) + "\n");
    Response.Write("atanh=" + Math.atanh(0.5).toFixed(3) + "\n");

    Response.Write("expm1=" + Math.expm1(1).toFixed(6) + "\n");
    Response.Write("log1p=" + Math.log1p(1).toFixed(6) + "\n");
    Response.Write("log10=" + Math.log10(1000) + "\n");
    Response.Write("log2=" + Math.log2(8) + "\n");

    Response.Write("hypot=" + Math.hypot(3, 4) + "\n");
    Response.Write("fround_diff=" + (Math.fround(1.337) !== 1.337 ? "yes" : "no") + "\n");
    Response.Write("imul=" + Math.imul(0xffffffff, 5) + "\n");
    Response.Write("clz32_1=" + Math.clz32(1) + "\n");
    Response.Write("clz32_0=" + Math.clz32(0) + "\n");

    Response.Write("nan_case=" + (isNaN(Math.log1p(-2)) ? "yes" : "no") + "\n");
    Response.Write("inf_case=" + (!isFinite(Math.hypot(Infinity, 3)) ? "yes" : "no"));
</script>