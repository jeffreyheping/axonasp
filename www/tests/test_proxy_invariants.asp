<%
' AxonASP Server
' Phase 5 - Proxy/Reflect Invariant Tests (ECMAScript 6 #10.5)
' Run with: axonasp-cli.exe -r www/tests/test_proxy_invariants.asp
%>
<script language="javascript" runat="server">
(function() {
    var pass = 0;
    var fail = 0;

    function assert(label, condition, details) {
        if (condition) {
            Response.Write("[PASS] " + label + "\n");
            pass++;
            return;
        }
        if (details && details.length > 0) {
            Response.Write("[FAIL] " + label + " (" + details + ")\n");
        } else {
            Response.Write("[FAIL] " + label + "\n");
        }
        fail++;
    }

    function hasFragment(message, fragment) {
        if (fragment === "") {
            return true;
        }
        return String(message).indexOf(fragment) !== -1;
    }

    // -----------------------------------------------------------------------
    // Subphase 5.1 - Trap Validation Engine
    // -----------------------------------------------------------------------

    // get: non-configurable non-writable must return exact value
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { value: 42, writable: false, configurable: false });
            var p = new Proxy(t, { get: function() { return 99; } });
            var v = p.x;
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("get: non-configurable non-writable wrong value", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // get: same value is OK
    (function() {
        var t = {};
        Object.defineProperty(t, "x", { value: 42, writable: false, configurable: false });
        var p = new Proxy(t, { get: function() { return 42; } });
        assert("get: non-configurable non-writable same value OK", p.x === 42, "expected 42");
    })();

    // get: accessor no getter must return undefined
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { set: function(v){}, configurable: false });
            var p = new Proxy(t, { get: function() { return 1; } });
            var v = p.x;
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("get: non-configurable no-getter accessor must be undefined", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // has: non-configurable cannot be absent
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { value: 1, configurable: false });
            var p = new Proxy(t, { has: function() { return false; } });
            var exists = ("x" in p);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("has: non-configurable property cannot be hidden", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // has: non-extensible target own property cannot be hidden
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = { x: 1 };
            Object.preventExtensions(t);
            var p = new Proxy(t, { has: function() { return false; } });
            var exists = ("x" in p);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("has: non-extensible target own key cannot be hidden", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // has: configurable property can be hidden
    (function() {
        var t = {};
        Object.defineProperty(t, "x", { value: 1, configurable: true });
        var p = new Proxy(t, { has: function() { return false; } });
        assert("has: configurable property can be hidden", !("x" in p), "expected hidden key");
    })();

    // set: non-configurable non-writable different value
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { value: 42, writable: false, configurable: false });
            var p = new Proxy(t, { set: function() { return true; } });
            p.x = 99;
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("set: non-configurable non-writable different value", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // set: same value OK
    (function() {
        var t = {};
        Object.defineProperty(t, "x", { value: 42, writable: false, configurable: false });
        var p = new Proxy(t, { set: function() { return true; } });
        p.x = 42;
        assert("set: non-configurable non-writable same value OK", true, "");
    })();

    // deleteProperty: non-configurable
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { value: 1, configurable: false });
            var p = new Proxy(t, { deleteProperty: function() { return true; } });
            delete p.x;
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("deleteProperty: cannot delete non-configurable property", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // ownKeys: must include non-configurable keys
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { value: 1, configurable: false });
            var p = new Proxy(t, { ownKeys: function() { return []; } });
            Object.getOwnPropertyNames(p);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("ownKeys: must include non-configurable keys", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // ownKeys: non-extensible target, extra key not allowed
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = { a: 1 };
            Object.preventExtensions(t);
            var p = new Proxy(t, { ownKeys: function() { return ["a", "b"]; } });
            Object.getOwnPropertyNames(p);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("ownKeys: non-extensible target cannot add extra keys", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // defineProperty: non-extensible target, new property
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.preventExtensions(t);
            var p = new Proxy(t, { defineProperty: function() { return true; } });
            Reflect.defineProperty(p, "x", { value: 1, configurable: true });
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("defineProperty: non-extensible target, new property", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // getOwnPropertyDescriptor: hide non-configurable
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { value: 1, configurable: false });
            var p = new Proxy(t, { getOwnPropertyDescriptor: function() { return undefined; } });
            Object.getOwnPropertyDescriptor(p, "x");
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("getOwnPropertyDescriptor: cannot hide non-configurable", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // getOwnPropertyDescriptor: report non-configurable as configurable
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.defineProperty(t, "x", { value: 1, configurable: false, writable: false });
            var p = new Proxy(t, {
                getOwnPropertyDescriptor: function() {
                    return { value: 1, configurable: true, writable: false, enumerable: false };
                }
            });
            Object.getOwnPropertyDescriptor(p, "x");
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("getOwnPropertyDescriptor: cannot report configurable for non-configurable", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // -----------------------------------------------------------------------
    // Subphase 5.2 - Prototype & Extensibility Safety
    // -----------------------------------------------------------------------

    // getPrototypeOf: non-extensible must return same prototype
    (function() {
        var threw = false;
        var msg = "";
        try {
            var proto = {};
            var t = Object.create(proto);
            Object.preventExtensions(t);
            var p = new Proxy(t, { getPrototypeOf: function() { return null; } });
            Object.getPrototypeOf(p);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("getPrototypeOf: non-extensible different prototype", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // getPrototypeOf: same prototype is OK
    (function() {
        var proto = {};
        var t = Object.create(proto);
        Object.preventExtensions(t);
        var p = new Proxy(t, { getPrototypeOf: function(target) { return Object.getPrototypeOf(target); } });
        assert("getPrototypeOf: non-extensible same prototype OK", Object.getPrototypeOf(p) === proto, "prototype mismatch");
    })();

    // setPrototypeOf: non-extensible different prototype
    (function() {
        var threw = false;
        var msg = "";
        try {
            var proto = {};
            var t = Object.create(proto);
            Object.preventExtensions(t);
            var newProto = {};
            var p = new Proxy(t, { setPrototypeOf: function() { return true; } });
            Reflect.setPrototypeOf(p, newProto);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("setPrototypeOf: non-extensible cannot change prototype", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // setPrototypeOf: non-extensible same prototype OK
    (function() {
        var proto = {};
        var t = Object.create(proto);
        Object.preventExtensions(t);
        var p = new Proxy(t, { setPrototypeOf: function() { return true; } });
        Object.setPrototypeOf(p, proto);
        assert("setPrototypeOf: non-extensible same prototype OK", true, "");
    })();

    // preventExtensions: trap returns true but target still extensible
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            var p = new Proxy(t, {
                preventExtensions: function(target) {
                    return true;
                }
            });
            Object.preventExtensions(p);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("preventExtensions: trap true but target still extensible", threw && hasFragment(msg, "invariant"), threw ? ("wrong error: " + msg) : "no error thrown");
    })();

    // preventExtensions: well-behaved trap
    (function() {
        var t = {};
        var p = new Proxy(t, {
            preventExtensions: function(target) {
                Object.preventExtensions(target);
                return true;
            }
        });
        Object.preventExtensions(p);
        assert("preventExtensions: well-behaved trap OK", !Object.isExtensible(t), "target still extensible");
    })();

    // isExtensible: regression - must match target
    (function() {
        var threw = false;
        var msg = "";
        try {
            var t = {};
            Object.preventExtensions(t);
            var p = new Proxy(t, { isExtensible: function() { return true; } });
            Object.isExtensible(p);
        } catch (e) {
            threw = true;
            msg = e && e.message ? e.message : String(e);
        }
        assert("isExtensible: must match target (regression)", threw, threw ? "" : "no error thrown");
    })();

    // -----------------------------------------------------------------------
    Response.Write("\n--- Phase 5 Results: " + pass + " passed, " + fail + " failed ---\n");
})();
</script>
