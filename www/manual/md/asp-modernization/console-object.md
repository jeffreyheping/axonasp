# Write Diagnostic Output with the Global Console Object

## Overview
This page documents the global `console` object in G3Pix AxonASP. The object is available in both VBScript and JavaScript pages and provides runtime diagnostic methods.

## Syntax
```asp
<%
' VBScript
console.log "message"
console.info "message"
console.warn "message"
console.err "message"
console.error "message"
console.time "label"
console.timeEnd "label"
console.dir value
%>

<%@ Language=JScript %>
<%
// JScript
console.log("message");
console.info("message");
console.warn("message");
console.err("message");
console.error("message");
console.time("label");
console.timeEnd("label");
console.dir(value);
console.trace("message");
%>
```

## Parameters and Arguments
- **args** (required): one or more arguments for `console.log`, `console.info`, `console.warn`, `console.err`, and `console.error`.
- **Supported input forms:**
  - String: printed directly.
  - VBScript array: serialized to JSON and printed.
  - JScript array/object: serialized to JSON and printed.
- **Multi-argument behavior for log methods:**
  - `console.log`, `console.info`, `console.warn`, `console.err`, and `console.error` accept multiple arguments.
  - Each argument is serialized using AxonASP runtime conversion.
  - Serialized arguments are joined with one space in the output line.
- **Method names available in both VBScript and JavaScript:**
  - `console.info(args...)`
  - `console.log(args...)`
  - `console.warn(args...)`
  - `console.err(args...)`
  - `console.error(args...)`
  - `console.time(label)`
  - `console.timeEnd(label)`
  - `console.dir(value)`
- **JavaScript only:**
  - `console.trace([value])`

## Return Values
- `console.info` returns no value to the ASP page output.
- `console.log` returns no value to the ASP page output.
- `console.warn` returns no value to the ASP page output.
- `console.err` returns no value to the ASP page output.
- `console.error` returns no value to the ASP page output.
- `console.time` returns no value to the ASP page output.
- `console.timeEnd` returns no value to the ASP page output.
- `console.dir` returns no value to the ASP page output.
- `console.trace` returns no value to the ASP page output.

## Remarks
- Every console output line includes date and time.
- Stream routing:
  - `console.info` writes to standard output.
  - `console.log` writes to standard output.
  - `console.warn` writes to standard error.
  - `console.err` writes to standard error.
  - `console.error` writes to standard error.
  - `console.timeEnd` writes elapsed time output to standard output.
  - `console.dir` writes inspected values to standard output.
  - `console.trace` writes stack trace output to standard error.
- Console symbols in stream output:
  - `console.info`: `ℹ`
  - `console.log`: `⌨`
  - `console.warn`: `⚠`
  - `console.err`: `✖`
  - `console.error`: `✖`
  - `console.trace`: `↳`
- `console.err` is an alias for `console.error`.
- `console.time(label)` stores a high-precision timer for the label.
- `console.timeEnd(label)` prints elapsed milliseconds and removes the timer label.
- `console.dir(value)` prints a structured inspection view of the value.
- `console.trace` prints the JavaScript call stack with file, line, and column.
- `console.trace` is not available in VBScript runtime execution.
- File logging is controlled by `global.enable_log_files` in `config/axonasp.toml`.
- When enabled:
  - `console.log` and `console.info` are appended to `./temp/console.log`.
  - `console.warn`, `console.err`, and `console.error` are appended to `./temp/error.log`.
- File entries do not include decorative symbols. Files store timestamp, level, and message text.

## Code Example
```asp
<%
Dim items(2)
items(0) = "alpha"
items(1) = "beta"
items(2) = 3

console.info "Starting ASP page execution"
console.log "Items:", items
console.time "vb-load"
console.timeEnd "vb-load"
console.dir items
console.warn "Using fallback dataset", "code", 1001
console.err "Sample error line for diagnostics", "id", 42

Response.Write "Console sample completed."
%>
```

## JavaScript Stack Trace Example
```asp
<%@ Language=JScript %>
<%
function c() {
  console.trace("trace sample");
}

function b() {
  c();
}

function a() {
  b();
}

console.time("js-load");
a();
console.timeEnd("js-load");
%>
```
