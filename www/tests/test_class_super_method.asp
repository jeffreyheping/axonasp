<script runat="server" language="JScript">
    // Test 1: Basic super method call
    class Base {
        greet() {
            return "Hello";
        }
    }

    class Derived extends Base {
        greet() {
            return super.greet() + " World";
        }
    }

    var d = new Derived();
    Response.Write("Test 1 - Basic super.method(): " + d.greet());
    if (d.greet() === "Hello World") {
        Response.Write(" - PASS");
    } else {
        Response.Write(" - FAIL");
    }
    Response.Write("<br>");

    // Test 2: Complex inheritance chain (A -> B -> C)
    class A {
        getValue() {
            return "A";
        }
    }

    class B extends A {
        getValue() {
            return super.getValue() + "->B";
        }
    }

    class C extends B {
        getValue() {
            return super.getValue() + "->C";
        }
    }

    var c = new C();
    Response.Write("Test 2 - Multi-level chain: " + c.getValue());
    if (c.getValue() === "A->B->C") {
        Response.Write(" - PASS");
    } else {
        Response.Write(" - FAIL");
    }
    Response.Write("<br>");

    // Test 3: Super method with arguments
    class BaseMath {
        add(a, b) {
            return a + b;
        }
    }

    class DerivedMath extends BaseMath {
        add(a, b) {
            var baseResult = super.add(a, b);
            return baseResult * 2;
        }
    }

    var m = new DerivedMath();
    Response.Write("Test 3 - Super with arguments: " + m.add(3, 4));
    if (m.add(3, 4) === 14) {  // (3 + 4) * 2 = 14
        Response.Write(" - PASS");
    } else {
        Response.Write(" - FAIL");
    }
    Response.Write("<br>");

    // Test 4: Super property access
    class BaseProp {
        constructor() {
            this.x = 10;
        }
        getX() {
            return this.x;
        }
    }

    class DerivedProp extends BaseProp {
        constructor() {
            super();
            this.x = 20;
        }
        getXWithSuper() {
            super.x = 15;
            return this.x;
        }
    }

    var p = new DerivedProp();
    var xVal = p.getXWithSuper();
    Response.Write("Test 4 - Super property access: " + xVal);
    if (p.x === 15) {
        Response.Write(" - PASS");
    } else {
        Response.Write(" - FAIL");
    }
    Response.Write("<br>");

    // Test 5: Multiple methods with super calls
    class Animal {
        speak() {
            return "sound";
        }
        move() {
            return "moving";
        }
    }

    class Dog extends Animal {
        speak() {
            return super.speak() + " - bark";
        }
        move() {
            return super.move() + " - running";
        }
    }

    var dog = new Dog();
    Response.Write("Test 5 - Multiple super methods: speak=" + dog.speak() + ", move=" + dog.move());
    if (dog.speak() === "sound - bark" && dog.move() === "moving - running") {
        Response.Write(" - PASS");
    } else {
        Response.Write(" - FAIL");
    }
    Response.Write("<br>");

    // Test 6: Super in nested method calls
    class BaseNested {
        getValue() {
            return 10;
        }
    }

    class DerivedNested extends BaseNested {
        getValue() {
            return super.getValue() + 5;
        }
        calculate() {
            return this.getValue() * 2;
        }
    }

    var dn = new DerivedNested();
    Response.Write("Test 6 - Nested method with super: " + dn.calculate());
    if (dn.calculate() === 30) {  // (10 + 5) * 2 = 30
        Response.Write(" - PASS");
    } else {
        Response.Write(" - FAIL");
    }
    Response.Write("<br>");

    // Test 7: Super method returns this
    class Builder {
        setValue(val) {
            this.value = val;
            return this;
        }
    }

    class FluentBuilder extends Builder {
        setValue(val) {
            super.setValue(val);
            this.modified = true;
            return this;
        }
    }

    var fb = new FluentBuilder();
    fb.setValue(42);
    Response.Write("Test 7 - Super returns this: value=" + fb.value + ", modified=" + fb.modified);
    if (fb.value === 42 && fb.modified === true) {
        Response.Write(" - PASS");
    } else {
        Response.Write(" - FAIL");
    }
    Response.Write("<br>");

    Response.Write("All super.method() tests completed successfully");
</script>