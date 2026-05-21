// Minimal diagnostic test
var arr = [{name: "alpha", error: "err1"}, {name: "beta", error: "err2"}];

console.log("=== Test 1: Direct property access ===");
console.log(arr[0].name);
console.log(arr[0]);

console.log("=== Test 2: forEach iteration ===");
arr.forEach(function(m) {
    console.log("m:", m);
    console.log("m.name:", m.name);
});

console.log("=== Test 3: Arrow function ===");
// Note: arrow functions in ES6
