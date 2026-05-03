# Use JavaScript (JScript) in AxonASP Pages

## Overview
AxonASP provides a high-performance JavaScript (JScript) execution engine that allows you to write server-side logic using ECMAScript 5 (ES5) standards. This page covers how to enable JavaScript (JScript), use ASP intrinsic objects, and leverage modern JavaScript features within your ASP applications.

## Syntax
To set JavaScript (JScript) as the default language for an entire page, use the language directive at the very first line of your file:

```asp
<%@ Language="javascript" %>
```

or

```asp
<%@ Language="jscript" %>
```

Alternatively, you can use JavaScript (JScript) within specific script blocks:

```html
<script runat="server" language="javascript">
    // JavaScript (JScript) code here
</script>
```

## Parameters and Arguments
- **Language Directive** (Required for page-level): The value must be `"JavaScript (JScript)"` or `"Javascript"`.
- **runat="server"** (Required for script tags): Ensures the code executes on the server rather than the client browser.
- **ASP Intrinsic Objects**: Native access to **Request**, **Response**, **Server**, **Session**, **Application**, and **Err**. Note that in JavaScript (JScript), these object names and their members are **case-sensitive**.

## Return Values
The JavaScript (JScript) engine returns standard JavaScript values (String, Number, Boolean, Object, Array, null, undefined). When communicating with the AxonASP VM or VBScript components:
- JavaScript objects are automatically converted to their closest AxonASP **Value** equivalent.
- **undefined** and **null** map to **Empty** in the VM context.

## Remarks
- **ECMAScript 5 Support**: AxonASP's JavaScript (JScript) engine supports ES5 features, including JSON support (`JSON.parse`, `JSON.stringify`), and standard Array methods (`map`, `filter`, `reduce`).
- **Case Sensitivity**: Unlike VBScript, JavaScript (JScript) is strictly case-sensitive. You must use `Response.Write`, not `response.write`.
- **Engine Architecture**: JavaScript (JScript) execution in AxonASP utilizes a sophisticated Abstract Syntax Tree (AST) parser and interpreter, providing optimized performance for complex logic.
- **Global Console**: The engine includes a built-in **console** object (`console.log`, `console.warn`, `console.error`) for server-side debugging and diagnostics. Output is directed to the system console or log files depending on your `axonasp.toml` configuration.
- **Interoperability**: You can mix VBScript and JavaScript (JScript) in the same application by using separate `<script runat="server">` blocks, though global variable sharing follows standard ASP scoping rules.

## Code Example
The following example demonstrates using ES5 features and ASP objects within a JavaScript (JScript) page:

```asp
<%@ Language="javascript" %>
<%

// Using ES5 Array methods
var data = [1, 2, 3, 4, 5];
var doubled = data.map(function(n) {
    return n * 2;
});

// Using the JSON object
var responseData = {
    status: "success",
    processed: doubled,
    timestamp: new Date().toISOString()
};

Response.ContentType = "application/json";
Response.Write(JSON.stringify(responseData));

// Server-side logging
console.log("JSON response sent for timestamp: " + responseData.timestamp);
%>
```
