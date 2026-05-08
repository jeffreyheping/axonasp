//go:build wasm

package main

import (
	"bytes"
	"syscall/js"

	"g3pix.com.br/axonasp/axonvm"
	"g3pix.com.br/axonasp/axonvm/asp"
)

var (
	sharedApplication = asp.NewApplication()
	sharedSession     = asp.NewSession()
)

// executeASP handles the invocation from JS: AxonASP.execute(code)
func executeASP(this js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return "Error: No code provided"
	}
	code := args[0].String()

	compiler := axonvm.NewASPCompiler(code)
	compiler.SetSourceName("/wasm.asp")

	if err := compiler.Compile(); err != nil {
		return "Compile Error: " + err.Error()
	}

	vm := axonvm.AcquireVMFromCompiler(compiler)
	defer vm.Release()

	var outBuf bytes.Buffer
	host := axonvm.NewMockHost()
	host.SetOutput(&outBuf)

	host.SetApplication(sharedApplication)
	host.SetSession(sharedSession)
	vm.SetHost(host)

	runErr := vm.Run()
	host.Response().Flush()
	host.Response().ReleaseBuffer()

	output := outBuf.String()
	if runErr != nil {
		output += "\nRuntime Error: " + runErr.Error()
	}

	return output
}

func main() {
	js.Global().Set("AxonASP", js.ValueOf(map[string]interface{}{
		"execute": js.FuncOf(executeASP),
	}))

	c := make(chan struct{}, 0)
	println("G3pix ❖ AxonASP initialized")
	<-c
}
