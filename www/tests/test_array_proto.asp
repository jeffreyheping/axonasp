<%
@ Language = "JScript"
%>
<script runat="server" language="JScript">
    var a = [1, 2, 3, 4, 5];
    a.copyWithin(0, 3);
    Response.Write("copyMain=" + a.join(",") + "\n");

    var b = [1, 2, 3, 4, 5];
    b.copyWithin(-2, 0, 2);
    Response.Write("copyNeg=" + b.join(",") + "\n");

    var base = [10, 20, 30];
    var keys = [];
    for (var k of base.keys()) {
        keys.push(k);
    }

    var entries = [];
    for (var e of base.entries()) {
        entries.push(e[0] + ":" + e[1]);
    }

    Response.Write("keys=" + keys.join(",") + "\n");
    Response.Write("entries=" + entries.join(","));
</script>