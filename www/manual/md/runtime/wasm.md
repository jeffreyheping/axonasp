# WebAssembly (WASM) Support

## Overview
AxonASP provides experimental support for compilation to WebAssembly (WASM). This allows the entire ASP Virtual Machine and compiler to run directly inside a modern web browser, enabling serverless execution of Classic ASP scripts on the client side.

The WASM module encapsulates the single-pass compiler and the stack-based VM, exposing a simple JavaScript API (`AxonASP.execute`) to compile and run ASP code and capture its output.

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
- `G3HTTP`
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

*Note: Core language features, basic intrinsics (Response, Request, Server, Application, Session), and safe data structure objects (e.g., Scripting.Dictionary) remain fully functional.*

## Using the WASM Playground
The AxonASP project includes a built-in playground to test ASP code natively in your browser. 

Once compiled, navigate to the `wasm/` directory and host it using a local web server (e.g., using `python -m http.server 8080`). Open the `default.asp` file (which acts as a standard HTML page for the playground) in your browser.

The playground demonstrates how to load the module using `wasm_exec.js` and interact with the engine using JavaScript:

```javascript
const go = new Go();
WebAssembly.instantiateStreaming(fetch("axonasp.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
    console.log("AxonASP module ready.");
});

function runCode() {
    const code = "<% Response.Write \"Hello from WASM!\" %>";
    const result = AxonASP.execute(code);
    console.log(result);
}
```