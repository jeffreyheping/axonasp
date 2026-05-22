# WebAssembly (WASM) Support

## Overview
AxonASP provides experimental support for compilation to WebAssembly (WASM). This allows the entire ASP Virtual Machine and compiler to run directly inside a modern web browser, enabling serverless execution of Classic ASP scripts on the client side, using VBScript and JavaScript for logic.

You can use this to create client-side applications using VBScript/ASP as the scripting language. The WASM playground is available in the `wasm/` directory of the project, demonstrating how to load and run ASP code natively in the browser.

The WASM module encapsulates the single-pass compiler and the stack-based VM, exposing a simple JavaScript API to compile and run ASP code asynchronously.

## AxonLive in WASM (Blazor-like Architecture)
AxonASP WASM supports **AxonLive**, the native reactive component framework. This allows you to build Single-Page Applications (SPAs) where Classic ASP logic directly manipulates the browser's Document Object Model (DOM) without network latency.

When `G3AXONLIVE` is instantiated in the WASM environment, methods like `SetStyle`, `AddClass`, and `SetAttribute` bypass JSON payloads and use `syscall/js` to mutate HTML elements instantaneously.

## Building for WASM
To compile AxonASP for WebAssembly, use the provided build scripts with the `wasm` platform target:

Windows PowerShell:
```powershell
./build.ps1 -Platform wasm
```

Linux and macOS Bash:
```bash
./build.sh --platform wasm
```

This process generates two files in the `wasm/` directory:
- `axonasp.wasm`: The compiled WebAssembly binary.
- `wasm_exec.js`: The required Go WebAssembly runtime environment script.

## Disabled Libraries
Due to the constraints of the browser's sandbox environment (such as lack of direct file system access, network socket restrictions, and unsupported CGO bindings), several native libraries are strictly disabled in the WASM build.

The following objects cannot be instantiated (`Server.CreateObject`) in the WASM runtime:
- `ADODB.Connection`
- `ADODB.Recordset`
- `ADODB.Command`
- `ADODB.Stream`
- `ADOX.Catalog`
- `G3DB`
- `G3FC`
- `G3FILES`
- `G3FILEUPLOADER`
- `G3IMAGE`
- `G3MAIL`
- `G3PDF`
- `G3SEARCH`
- `G3TAR`
- `G3ZIP`
- `G3ZLIB`
- `G3ZSTD`
- `Scripting.FileSystemObject`
- `WScript.Shell`

*Note: Most of the core language features, most of basic intrinsics (Application, Session), and safe data structure objects (e.g., Scripting.Dictionary) remain fully functional. Response, Request, Server will be limited in functionality. The `G3AXONLIVE` object is fully supported and optimized for the DOM.*

## WASM-Specific Built-ins
AxonASP WASM re-introduces standard VBScript interactive UI functions that are typically disabled in server-side ASP:

- **MsgBox(prompt, [buttons])**: 
    - If `vbOKCancel` (1) or `vbYesNo` (4) flags are set in the `buttons` argument, it triggers a browser `confirm()` dialog and returns the corresponding VBScript constant (`vbOK`, `vbCancel`, `vbYes`, `vbNo`).
    - Otherwise, it triggers a standard browser `alert()`.
- **InputBox(prompt, [title], [default])**: 
    - Triggers a browser `prompt()` dialog. Returns the user's input string, or an empty string if cancelled.

These functions allow your Classic ASP logic to interact directly with the user while running natively in the browser.

Once compiled, navigate to the `wasm/` directory and host it using a local web server (e.g., using `python -m http.server 8080`). Open `index.htm` in your browser.

The playground demonstrates how to load the module using `wasm_exec.js` and interact with the engine using JavaScript Promises:

```javascript
const go = new Go();
WebAssembly.instantiateStreaming(fetch("axonasp.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    console.log("AxonASP module ready.");
});

async function runCode() {
    const code = "<% Response.Write \"Hello from WASM!\" %>";
    try {
        const result = await AxonASP.execute(code);
        console.log(result);
    } catch (err) {
        console.error("Execution failed:", err);
    }
}
```