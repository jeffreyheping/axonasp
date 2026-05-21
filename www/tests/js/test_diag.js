// Diagnostic test to isolate the catch-scope variable corruption

console.log("=== A. Simple catch without inner function ===");
function testA(featureName) {
    try {
        throw new Error("test error");
    } catch (error) {
        var errMsg = error.message || error.toString();
        console.log(`[A] ${featureName} -> ${errMsg}`);
    }
}
testA("FeatureA");

console.log("=== B. Catch where inner function throws ===");
function testB(featureName, fn) {
    try {
        fn();
    } catch (error) {
        var errMsg = error.message || error.toString();
        console.log(`[B] ${featureName} -> ${errMsg}`);
    }
}
testB("FeatureB", function() { throw new Error("inner error"); });

console.log("=== C. Same as B but with multiple params ===");
function testC(name1, name2, fn) {
    try {
        fn();
    } catch (error) {
        var errMsg = error.message || error.toString();
        console.log(`[C] ${name1} / ${name2} -> ${errMsg}`);
    }
}
testC("First", "Second", function() { throw new Error("test C"); });

console.log("=== D. Check if local slot 0 is corrupted ===");
function testD(p1, p2, p3, fn) {
    try {
        fn();
    } catch (error) {
        var em = error.message || error.toString();
        console.log(`[D] p1=${p1}, p2=${p2}, p3=${p3} -> ${em}`);
    }
}
testD("A", "B", "C", function() { throw new Error("test D"); });

console.log("=== E. forEach inside function with catch ===");
function testE(featureName, fn) {
    try {
        fn();
    } catch (error) {
        var results = [{ name: featureName }];
        results.forEach(function(item) {
            console.log(`[E] ${item.name}`);
        });
    }
}
testE("FeatureE", function() { throw new Error("test E"); });

console.log("=== F. Arrow function inside catch ===");
function testF(featureName, fn) {
    try {
        fn();
    } catch (error) {
        var results = [{ name: featureName }];
        results.forEach(item => console.log(`[F] ${item.name}`));
    }
}
testF("FeatureF", function() { throw new Error("test F"); });

console.log("=== G. Test with error.message access separately ===");
function testG(featureName, fn) {
    try {
        fn();
    } catch (error) {
        var msg = error.message;
        console.log(`[G] featureName=${featureName}, msg=${msg}`);
    }
}
testG("FeatureG", function() { throw new Error("test G message"); });

console.log("=== H. Does the var declaration matter? ===");
function testH(featureName, fn) {
    try {
        fn();
    } catch (error) {
        var msg = "hardcoded";
        console.log(`[H] ${featureName} -> ${msg}`);
    }
}
testH("FeatureH", function() { throw new Error("test H"); });

console.log("=== I. Does the error.message expression matter? ===");
function testI(featureName, fn) {
    try {
        fn();
    } catch (error) {
        var x = 42;
        console.log(`[I] ${featureName} -> ${x}`);
    }
}
testI("FeatureI", function() { throw new Error("test I"); });

console.log("=== J. Minimal repro ===");
function testJ(featureName, fn) {
    try {
        fn();
    } catch (error) {
        console.log(`[J] ${featureName}`);
    }
}
testJ("FeatureJ", function() { throw new Error("test J"); });
