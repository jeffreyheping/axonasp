<%@ language="JScript" %>
<%
    function report(name, actual, expected) {
        Response.Write("Testing " + name + ": ");
        if (String(actual) === String(expected)) {
            Response.Write("PASS\n");
        } else {
            Response.Write("FAIL (Expected " + expected + ", got " + actual + ")\n");
        }
    }

    function sum(n, acc) {
        if (n === 0) {
            return acc;
        }
        return sum(n - 1, acc + 1);
    }

    function depth(n) {
        if (n === 0) {
            return 0;
        }
        return 1 + depth(n - 1);
    }

    report("Tail recursion depth 100000", sum(100000, 0), 100000);

    try {
        depth(10101);
        Response.Write("Testing depth guard at 10101: FAIL (Expected Out of stack space)\n");
    } catch (e) {
        var msg = "";
        if (e && e.message) {
            msg = e.message;
        }
        if (msg.indexOf("Out of stack space") !== -1) {
            Response.Write("Testing depth guard at 10101: PASS (TCO Error Prevention system action)\n");
        } else {
            Response.Write("Testing depth guard at 10101: FAIL (Unexpected error: " + msg + ")\n");
        }
    }

    Response.Write("\nTCO LIMIT TESTS COMPLETED\n");
%>
