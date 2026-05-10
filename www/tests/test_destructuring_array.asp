<script runat="server" language="JScript">
/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimarães - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 */
    Response.Write("Testing Array Destructuring (Sub-Phase 5.3)\n");

    // 1. Basic array destructuring
    var [a, b, c] = [1, 2, 3];
    Response.Write("a: " + a + ", b: " + b + ", c: " + c + "\n");

    // 2. Nested array destructuring
    var [x, [y, z]] = ["outer", ["inner1", "inner2"]];
    Response.Write("x: " + x + ", y: " + y + ", z: " + z + "\n");

    // 3. String destructuring (using iteration protocol)
    var [h, e, l, l2, o] = "Hello";
    Response.Write("chars: " + h + e + l + l2 + o + "\n");

    // 4. Elision
    var [first, , last] = [10, 20, 30];
    Response.Write("first: " + first + ", last: " + last + "\n");

    // 5. Destructuring from Map (yields [key, value] pairs)
    var m = new Map();
    m.set("id", 123);
    var [[k, v]] = m;
    Response.Write("Map key: " + k + ", value: " + v + "\n");

    // 6. Non-iterable check
    try {
        var [fail] = true;
    } catch (err) {
        Response.Write("Caught expected error (non-iterable): " + (err.indexOf("not iterable") !== -1) + "\n");
    }
</script>
