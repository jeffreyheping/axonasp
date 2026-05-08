<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>AxonASP WASM Playground</title>
    <style>
        body { font-family: Arial, sans-serif; padding: 20px; background: #ece9d8; }
        .container { max-width: 800px; margin: 0 auto; background: white; padding: 20px; border-radius: 10px; box-shadow: 0 4px 14px rgba(0,0,0,0.20); }
        h1 { color: #003399; border-bottom: 3px solid #3366cc; padding-bottom: 10px; }
        textarea { width: 100%; height: 250px; font-family: monospace; padding: 10px; box-sizing: border-box; border: 1px solid #808080; border-radius: 6px; }
        button { background: linear-gradient(180deg, #3366cc 0%, #003399 100%); color: white; border: none; padding: 10px 20px; font-size: 16px; border-radius: 6px; cursor: pointer; margin-top: 10px; }
        button:hover { background: linear-gradient(180deg, #4076e0 0%, #0044cc 100%); }
        pre { background: #f4f4f4; padding: 15px; border-radius: 6px; border: 1px solid #ddd; min-height: 150px; overflow-x: auto; }
        #status { font-weight: bold; color: #c8a200; margin-bottom: 10px; }
    </style>
    <script src="wasm_exec.js"></script>
</head>
<body>
    <div class="container">
        <h1>AxonASP WASM Playground</h1>
        <div id="status">Loading WASM module...</div>
        
        <div>
            <textarea id="code"><%
Response.Write "<h1>Hello from AxonASP WASM!</h1>"
Response.Write "<p>The current time is: " & Now() & "</p>"

Dim i
For i = 1 to 5
    Response.Write "Line " & i & "<br>"
Next
%></textarea>
            <button onclick="runCode()" id="runBtn" disabled>Run ASP Code</button>
        </div>

        <h2>Output:</h2>
        <pre id="output"></pre>
    </div>

    <script>
        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("axonasp.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
            document.getElementById("status").innerText = "AxonASP module ready.";
            document.getElementById("status").style.color = "green";
            document.getElementById("runBtn").disabled = false;
        }).catch((err) => {
            document.getElementById("status").innerText = "Error loading AxonASP module: " + err;
            document.getElementById("status").style.color = "red";
        });

        function runCode() {
            const outEl = document.getElementById("output");
            if (typeof AxonASP !== "undefined") {
                const code = document.getElementById("code").value;
                try {
                    const result = AxonASP.execute(code);
                    outEl.innerHTML = result; // Allows rendering HTML from Response.Write
                } catch (e) {
                    outEl.innerText = "Exception: " + e;
                }
            } else {
                outEl.innerText = "AxonASP module not loaded yet.";
            }
        }
    </script>
</body>
</html>
