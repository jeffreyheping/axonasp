<%
@ Language = "JScript"
%>
<script runat="server" language="JScript">
    var s = "A😀B";
    Response.Write("cp0=" + s.codePointAt(0) + "\n");
    Response.Write("cp1=" + s.codePointAt(1) + "\n");
    Response.Write("cp2=" + s.codePointAt(2) + "\n");
    Response.Write("cp99=" + (s.codePointAt(99) === undefined ? "undef" : "bad") + "\n");

    var decomposed = "e\u0301";
    Response.Write("nfc=" + (decomposed.normalize("NFC") === "é" ? "yes" : "no") + "\n");
    Response.Write("nfd=" + ("é".normalize("NFD") === decomposed ? "yes" : "no") + "\n");

    try {
        "x".normalize("BAD");
        Response.Write("invalid=fail");
    } catch (e) {
        Response.Write("invalid=" + (("" + e).indexOf("RangeError") >= 0 ? "range" : "other"));
    }
</script>