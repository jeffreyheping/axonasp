<%
@ Language = "JScript"
%>
<script runat="server" language="JScript">
    var full = "https://example.com/a path/?q=hello world&x=1+2#frag";
    var encoded = encodeURI(full);
    var decoded = decodeURI(encoded);

    var component = "q=hello world&x=1+2";
    var encodedComponent = encodeURIComponent(component);
    var decodedComponent = decodeURIComponent(encodedComponent);

    Response.Write("ENC=" + encoded + "\n");
    Response.Write("DEC_OK=" + (decoded === full ? "yes" : "no") + "\n");
    Response.Write("COMP_ENC=" + encodedComponent + "\n");
    Response.Write("COMP_DEC_OK=" + (decodedComponent === component ? "yes" : "no") + "\n");

    try {
        decodeURIComponent("%");
        Response.Write("MALFORMED=fail");
    } catch (e) {
        Response.Write("MALFORMED=" + (e.name || "error"));
    }
</script>