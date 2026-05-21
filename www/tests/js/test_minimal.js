// Minimal test to isolate template literal and forEach issues
console.log("=== Test 1: Simple template literal ===");
var name = "World";
console.log(`Hello ${name}`);

console.log("=== Test 2: Object literal with template ===");
var obj = { name: "TestObject", value: 42 };
console.log(`Object name is: ${obj.name}`);

console.log("=== Test 3: forEach with template ===");
var items = [{ name: "Item1" }, { name: "Item2" }, { name: "Item3" }];
items.forEach(function(item) {
    console.log(`- ${item.name}`);
});

console.log("=== Test 4: Arrow function with forEach ===");
items.forEach(item => console.log(`>> ${item.name}`));

console.log("=== Test 5: Try/catch with template ===");
function testFn(featureName, testFn) {
    try {
        testFn();
    } catch (error) {
        var errMsg = error.message || error.toString();
        console.log(`[ERROR] ${featureName} -> ${errMsg}`);
    }
}
testFn("Test Feature", function() { throw new Error("test error message"); });

console.log("=== Test 6: report.missing forEach ===");
var report = { missing: [] };
report.missing.push({ name: "Feature A", error: "Err A" });
report.missing.push({ name: "Feature B", error: "Err B" });
report.missing.push({ name: "Feature C", error: "Err C" });
report.missing.forEach(function(m) {
    console.log(`- ${m.name}`);
});

console.log("=== Test 7: forEach with var scope ===");
var report2 = { missing: [] };
var testFeature = function(featureName, testFn) {
    try {
        testFn();
    } catch (error) {
        report2.missing.push({ name: featureName, error: error.message || error.toString() });
    }
};
testFeature("Module: fs", function() { throw new Error("fs.readFileSync is not a function"); });
testFeature("Module: path", function() { throw new Error("path is not defined"); });
report2.missing.forEach(function(m) {
    console.log(`- ${m.name}`);
});

console.log("=== Test 8: const variable with try/catch ===");
const report3 = { missing: [] };
function testFeature2(featureName, testFn) {
    try {
        testFn();
    } catch (error) {
        report3.missing.push({ name: featureName, error: error.message || error.toString() });
    }
}
testFeature2("Module: fs", function() { throw new Error("is not a function"); });
testFeature2("Module: path", function() { throw new Error("is not defined"); });
testFeature2("Module: os", function() { throw new Error("is not a function"); });
report3.missing.forEach(function(m) {
    console.log(`- ${m.name}`);
});
