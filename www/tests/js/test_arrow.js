// Test arrow function forEach
var arr = [{name: "alpha", error: "err1"}, {name: "beta", error: "err2"}];

console.log("=== Arrow function forEach ===");
arr.forEach(m => console.log("- " + m.name));

console.log("=== Arrow function with template ===");
arr.forEach(m => console.log(`- ${m.name}`));

console.log("=== Regular function forEach ===");
arr.forEach(function(m) { console.log("- " + m.name); });
