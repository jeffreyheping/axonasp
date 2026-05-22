//go:build wasm && !lib_g3axonlive_disabled

/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimaraes - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package axonvm

import (
	"encoding/json"
	"strings"
	"syscall/js"
)

// G3ALComponentProxy represents a specific reactive component in the DOM.
type G3ALComponentProxy struct {
	parent      *G3AXONLIVE
	componentID string
}

func (p *G3ALComponentProxy) getElement() js.Value {
	doc := js.Global().Get("document")
	if doc.IsUndefined() || doc.IsNull() {
		return js.Undefined()
	}
	el := doc.Call("getElementById", p.componentID)
	if el.IsUndefined() || el.IsNull() {
		println("✖ AxonLive WASM: Element not found:", p.componentID)
	}
	return el
}

func (p *G3ALComponentProxy) DispatchPropertyGet(propertyName string) Value {
	el := p.getElement()
	if el.IsUndefined() || el.IsNull() {
		return Value{Type: VTEmpty}
	}

	// 1. Try direct property access
	val := el.Get(propertyName)
	if val.IsUndefined() || val.IsNull() {
		val = el.Get(strings.ToLower(propertyName))
	}

	// 2. Try getAttribute (essential for data-* attributes used in counter)
	if val.IsUndefined() || val.IsNull() {
		attrVal := el.Call("getAttribute", propertyName)
		if !attrVal.IsNull() && !attrVal.IsUndefined() {
			val = attrVal
		}
	}

	if val.IsUndefined() || val.IsNull() {
		println("  → GetProperty:", propertyName, "= (Empty)")
		return Value{Type: VTEmpty}
	}

	println("  → GetProperty:", propertyName, "=", val.String())

	if val.Type() == js.TypeBoolean {
		if val.Bool() {
			return Value{Type: VTBool, Num: 1}
		}
		return Value{Type: VTBool, Num: 0}
	}

	return NewString(val.String())
}
func (p *G3ALComponentProxy) DispatchPropertySet(propertyName string, args []Value) {
	if len(args) < 1 {
		return
	}
	el := p.getElement()
	if el.IsUndefined() || el.IsNull() {
		return
	}

	val := args[0]
	// If it starts with data-, set attribute as well
	if strings.HasPrefix(strings.ToLower(propertyName), "data-") {
		el.Call("setAttribute", propertyName, val.String())
	}

	if val.Type == VTBool {
		el.Set(propertyName, val.Num != 0)
	} else {
		el.Set(propertyName, val.String())
	}
}

func (p *G3ALComponentProxy) DispatchMethod(methodName string, args []Value) Value {
	method := strings.ToLower(strings.TrimSpace(methodName))
	el := p.getElement()
	if el.IsUndefined() || el.IsNull() {
		return Value{Type: VTEmpty}
	}

	switch method {
	case "setstyle":
		if len(args) >= 2 {
			prop := args[0].String()
			val := args[1].String()
			println("  → SetStyle:", prop, "=", val)
			el.Get("style").Set(prop, val)
		}
	case "addclass":
		if len(args) >= 1 {
			el.Get("classList").Call("add", args[0].String())
		}
	case "removeclass":
		if len(args) >= 1 {
			el.Get("classList").Call("remove", args[0].String())
		}
	case "setattribute":
		if len(args) >= 2 {
			el.Call("setAttribute", args[0].String(), args[1].String())
		}
	case "removeattribute":
		if len(args) >= 1 {
			el.Call("removeAttribute", args[0].String())
		}
	case "setvalue":
		if len(args) >= 1 {
			val := args[0].String()
			println("  → SetValue:", val)
			tagName := el.Get("tagName").String()
			if tagName == "INPUT" || tagName == "TEXTAREA" || tagName == "SELECT" {
				el.Set("value", val)
			} else {
				el.Set("innerText", val)
			}
		}
	default:
		if len(args) == 0 {
			return p.DispatchPropertyGet(methodName)
		}
	}
	return Value{Type: VTEmpty}
}

type G3AXONLIVE struct {
	vm *VM

	initiated        bool
	isAsyncRequest   bool
	eventSessionID   string
	eventComponentID string
	eventName        string
	eventArgs        map[string]string
}

func (vm *VM) newG3AxonLiveObject() Value {
	obj := &G3AXONLIVE{vm: vm}
	id := vm.nextDynamicNativeID
	vm.nextDynamicNativeID++
	vm.g3axonliveItems[id] = obj
	return Value{Type: VTNativeObject, Num: id}
}

func (g *G3AXONLIVE) DispatchPropertyGet(propertyName string) Value {
	return g.DispatchMethod(propertyName, nil)
}

func (g *G3AXONLIVE) DispatchPropertySet(_ string, _ []Value) {}

func (g *G3AXONLIVE) DispatchMethod(methodName string, args []Value) Value {
	method := strings.ToLower(strings.TrimSpace(methodName))
	switch method {
	case "initpage":
		return g.initPage()
	case "isasyncrequest":
		if g.isAsyncRequest {
			return Value{Type: VTBool, Num: 1}
		}
		return Value{Type: VTBool, Num: 0}
	case "eventcomponentid":
		return NewString(g.eventComponentID)
	case "eventname":
		return NewString(g.eventName)
	case "eventargs":
		return g.getEventArgsJSON()
	case "geteventarg":
		return g.getEventArg(args)
	case "registercomponent":
		return g.registerComponent(args)
	case "getcomponent":
		return g.getComponent(args)
	case "endasyncresponse":
		return g.endAsyncResponse()
	case "settimer":
		return g.setTimer(args)
	case "redirect":
		return g.addRedirectAction(args)
	case "trigger":
		return g.addTriggerAction(args)
	case "addattribute":
		return g.addAttributeAction(args)
	case "setcomponentproperty":
		return g.setComponentProperty(args)
	case "getcomponentproperty":
		return g.getComponentProperty(args)
	case "removecomponentproperty":
		return g.removeComponentProperty(args)
	case "clearcomponentstate":
		return g.clearComponentState(args)
	case "getcomponentstate":
		return g.getComponentState(args)
	case "registerpage":
		return Value{Type: VTEmpty}
	case "removesession":
		return Value{Type: VTEmpty}
	case "version":
		return NewString("2.0.0-wasm")
	case "startcleanup":
		return Value{Type: VTEmpty}
	case "stopcleanup":
		return Value{Type: VTEmpty}
	default:
		return Value{Type: VTEmpty}
	}
}

func (g *G3AXONLIVE) initPage() Value {
	if g.initiated {
		println("  → InitPage: Already initiated, returning", g.isAsyncRequest)
		if g.isAsyncRequest {
			return Value{Type: VTBool, Num: 1}
		}
		return Value{Type: VTBool, Num: 0}
	}
	g.initiated = true

	req := g.vm.host.Request()
	// Read WASM-injected headers
	headerVal := req.ServerVars.Get("HTTP_X_G3AXONLIVE")
	println("  → InitPage: HTTP_X_G3AXONLIVE =", headerVal)

	if strings.EqualFold(strings.TrimSpace(headerVal), "true") {
		g.isAsyncRequest = true
		g.eventSessionID = strings.TrimSpace(req.ServerVars.Get("HTTP_X_G3AXONLIVE_SESSIONID"))
		g.eventComponentID = strings.TrimSpace(req.ServerVars.Get("HTTP_X_G3AXONLIVE_COMPONENTID"))
		g.eventName = strings.TrimSpace(req.ServerVars.Get("HTTP_X_G3AXONLIVE_EVENTNAME"))

		println("❖ AxonLive WASM: Async event detected - Component:", g.eventComponentID, "Event:", g.eventName)

		argsJSON := strings.TrimSpace(req.ServerVars.Get("HTTP_X_G3AXONLIVE_EVENTARGS"))
		if argsJSON != "" {
			var parsedArgs map[string]string
			if err := json.Unmarshal([]byte(argsJSON), &parsedArgs); err == nil && parsedArgs != nil {
				g.eventArgs = parsedArgs
			} else {
				g.eventArgs = map[string]string{}
			}
		} else {
			g.eventArgs = map[string]string{}
		}
		return Value{Type: VTBool, Num: 1}
	}

	g.isAsyncRequest = false
	return Value{Type: VTBool, Num: 0}
}

func (g *G3AXONLIVE) getEventArgsJSON() Value {
	if !g.isAsyncRequest || len(g.eventArgs) == 0 {
		return NewString("{}")
	}
	data, err := json.Marshal(g.eventArgs)
	if err != nil {
		return NewString("{}")
	}
	return NewString(string(data))
}

func (g *G3AXONLIVE) getEventArg(args []Value) Value {
	if len(args) < 1 || !g.isAsyncRequest {
		return Value{Type: VTEmpty}
	}
	name := strings.TrimSpace(args[0].String())
	if name == "" {
		return Value{Type: VTEmpty}
	}
	if val, ok := g.eventArgs[name]; ok {
		return NewString(val)
	}
	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) registerComponent(args []Value) Value {
	if len(args) < 2 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	html := args[1].String()
	if componentID == "" {
		return Value{Type: VTEmpty}
	}

	doc := js.Global().Get("document")
	if !doc.IsUndefined() && !doc.IsNull() {
		el := doc.Call("getElementById", componentID)
		if !el.IsUndefined() && !el.IsNull() {
			el.Set("outerHTML", html)
		}
	}

	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) getComponent(args []Value) Value {
	if len(args) < 1 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	if componentID == "" {
		return Value{Type: VTEmpty}
	}

	proxy := &G3ALComponentProxy{
		parent:      g,
		componentID: componentID,
	}
	id := g.vm.nextDynamicNativeID
	g.vm.nextDynamicNativeID++
	if g.vm.g3axonliveProxyItems == nil {
		g.vm.g3axonliveProxyItems = make(map[int64]*G3ALComponentProxy)
	}
	g.vm.g3axonliveProxyItems[id] = proxy
	return Value{Type: VTNativeObject, Num: id}
}

func (g *G3AXONLIVE) endAsyncResponse() Value {
	resp := g.vm.host.Response()
	resp.End()
	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) setTimer(args []Value) Value {
	if len(args) < 3 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	evtName := strings.TrimSpace(args[1].String())
	delay := g.vm.asInt(args[2])

	if componentID == "" || evtName == "" || delay <= 0 {
		return Value{Type: VTEmpty}
	}

	// In WASM, we can call setTimeout directly
	window := js.Global()
	if !window.IsUndefined() && !window.IsNull() {
		axonLive := window.Get("G3AxonLive")
		if !axonLive.IsUndefined() && !axonLive.IsNull() {
			cb := js.FuncOf(func(this js.Value, jsArgs []js.Value) interface{} {
				axonLive.Call("dispatch", componentID, evtName)
				return nil
			})
			window.Call("setTimeout", cb, delay)
		}
	}
	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) addRedirectAction(args []Value) Value {
	if len(args) < 1 {
		return Value{Type: VTEmpty}
	}
	url := strings.TrimSpace(args[0].String())
	if url == "" {
		return Value{Type: VTEmpty}
	}
	window := js.Global()
	if !window.IsUndefined() && !window.IsNull() {
		window.Get("location").Set("href", url)
	}
	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) addTriggerAction(args []Value) Value {
	if len(args) < 2 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	evtName := strings.TrimSpace(args[1].String())
	if componentID == "" || evtName == "" {
		return Value{Type: VTEmpty}
	}

	window := js.Global()
	if !window.IsUndefined() && !window.IsNull() {
		axonLive := window.Get("G3AxonLive")
		if !axonLive.IsUndefined() && !axonLive.IsNull() {
			axonLive.Call("dispatch", componentID, evtName)
		}
	}
	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) addAttributeAction(args []Value) Value {
	if len(args) < 3 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	attrName := strings.TrimSpace(args[1].String())
	attrValue := args[2].String()

	doc := js.Global().Get("document")
	if !doc.IsUndefined() && !doc.IsNull() {
		el := doc.Call("getElementById", componentID)
		if !el.IsUndefined() && !el.IsNull() {
			el.Call("setAttribute", attrName, attrValue)
		}
	}
	return Value{Type: VTEmpty}
}

var wasmState = make(map[string]string)

func (g *G3AXONLIVE) setComponentProperty(args []Value) Value {
	if len(args) < 3 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	propertyName := strings.ToLower(strings.TrimSpace(args[1].String()))
	value := args[2].String()

	sessionID := g.eventSessionID
	if sessionID == "" {
		sessionID = "wasm-global"
	}

	key := sessionID + "\x00" + componentID + "\x00" + propertyName
	wasmState[key] = value

	// In WASM, also persist data- attributes to the DOM for transparency and easy retrieval via proxy
	if strings.HasPrefix(propertyName, "data-") {
		doc := js.Global().Get("document")
		if !doc.IsUndefined() && !doc.IsNull() {
			el := doc.Call("getElementById", componentID)
			if !el.IsUndefined() && !el.IsNull() {
				el.Call("setAttribute", propertyName, value)
			}
		}
	}

	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) getComponentProperty(args []Value) Value {
	if len(args) < 2 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	propertyName := strings.ToLower(strings.TrimSpace(args[1].String()))

	sessionID := g.eventSessionID
	if sessionID == "" {
		sessionID = "wasm-global"
	}

	key := sessionID + "\x00" + componentID + "\x00" + propertyName
	if val, ok := wasmState[key]; ok {
		return NewString(val)
	}

	// Fallback: Check the DOM via proxy logic
	proxy := &G3ALComponentProxy{parent: g, componentID: componentID}
	return proxy.DispatchPropertyGet(propertyName)
}

func (g *G3AXONLIVE) removeComponentProperty(args []Value) Value {
	if len(args) < 2 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())
	propertyName := strings.ToLower(strings.TrimSpace(args[1].String()))

	sessionID := g.eventSessionID
	if sessionID == "" {
		sessionID = "wasm-global"
	}

	key := sessionID + "\x00" + componentID + "\x00" + propertyName
	delete(wasmState, key)
	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) clearComponentState(args []Value) Value {
	if len(args) < 1 {
		return Value{Type: VTEmpty}
	}
	componentID := strings.TrimSpace(args[0].String())

	sessionID := g.eventSessionID
	if sessionID == "" {
		sessionID = "wasm-global"
	}

	prefix := sessionID + "\x00" + componentID + "\x00"
	for k := range wasmState {
		if strings.HasPrefix(k, prefix) {
			delete(wasmState, k)
		}
	}
	return Value{Type: VTEmpty}
}

func (g *G3AXONLIVE) getComponentState(args []Value) Value {
	return NewString("State tracking in WASM")
}
