/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimarães - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 * Attribution Notice:
 * If this software is used in other projects, the name "AxonASP Server"
 * must be cited in the documentation or "About" section.
 *
 * Contribution Policy:
 * Modifications to the core source code of AxonASP Server must be
 * made available under this same license terms.
 */
package axonvm

import (
	"math"
	"strings"
	"time"
)

// jsIntlDateTimeFormatObject stores normalized Intl.DateTimeFormat state.
type jsIntlDateTimeFormatObject struct {
	localeTag string
	layout    string
}

// jsIntlNumberFormatObject stores normalized Intl.NumberFormat state.
type jsIntlNumberFormatObject struct {
	localeTag       string
	style           string
	digits          int
	useGrouping     bool
	currencySymbol  string
	currencySpacing string
}

// jsCreateIntlObject allocates the global Intl namespace and its constructor entries.
func (vm *VM) jsCreateIntlObject() Value {
	objID := vm.allocJSID()
	obj := make(map[string]Value, 4)
	obj["__js_type"] = NewString("Intl")
	obj["DateTimeFormat"] = vm.jsCreateIntlDateTimeFormatConstructor()
	obj["NumberFormat"] = vm.jsCreateIntlNumberFormatConstructor()
	vm.jsObjectItems[objID] = obj
	vm.jsPropertyItems[objID] = make(map[string]jsPropertyDescriptor, 4)
	return Value{Type: VTJSObject, Num: objID}
}

// jsCreateIntlDateTimeFormatConstructor allocates the Intl.DateTimeFormat constructor object.
func (vm *VM) jsCreateIntlDateTimeFormatConstructor() Value {
	objID := vm.allocJSID()
	obj := make(map[string]Value, 3)
	obj["__js_type"] = NewString("Intl.DateTimeFormat")
	obj["__js_ctor"] = NewString("IntlDateTimeFormat")
	obj["name"] = NewString("DateTimeFormat")
	vm.jsObjectItems[objID] = obj
	vm.jsPropertyItems[objID] = make(map[string]jsPropertyDescriptor, 4)
	return Value{Type: VTJSObject, Num: objID}
}

// jsCreateIntlNumberFormatConstructor allocates the Intl.NumberFormat constructor object.
func (vm *VM) jsCreateIntlNumberFormatConstructor() Value {
	objID := vm.allocJSID()
	obj := make(map[string]Value, 3)
	obj["__js_type"] = NewString("Intl.NumberFormat")
	obj["__js_ctor"] = NewString("IntlNumberFormat")
	obj["name"] = NewString("NumberFormat")
	vm.jsObjectItems[objID] = obj
	vm.jsPropertyItems[objID] = make(map[string]jsPropertyDescriptor, 4)
	return Value{Type: VTJSObject, Num: objID}
}

// jsIntlCreateFormatFunction allocates one bound formatter method for an Intl instance.
func (vm *VM) jsIntlCreateFormatFunction(ownerID int64, methodCtor string) Value {
	fn := vm.jsCreateIntrinsicFunction("format", methodCtor)
	if fn.Type == VTJSFunction {
		if obj, ok := vm.jsObjectItems[fn.Num]; ok {
			obj["__js_intl_owner"] = NewInteger(ownerID)
		}
	}
	return fn
}

// jsIntlCreateDateTimeFormat allocates one Intl.DateTimeFormat instance with normalized locale state.
func (vm *VM) jsIntlCreateDateTimeFormat(args []Value) Value {
	profile, localeTag := vm.jsIntlResolveLocaleProfile(jsArgOrUndefined(args, 0))
	options := Value{Type: VTJSUndefined}
	if len(args) > 1 {
		options = args[1]
	}
	layout := vm.jsIntlResolveDateTimeLayout(profile, options)
	objID := vm.allocJSID()
	obj := make(map[string]Value, 6)
	obj["__js_type"] = NewString("Intl.DateTimeFormat")
	obj["__js_ctor"] = NewString("IntlDateTimeFormat")
	obj["__js_intl_locale"] = NewString(localeTag)
	obj["__js_intl_layout"] = NewString(layout)
	obj["format"] = vm.jsIntlCreateFormatFunction(objID, "IntlDateTimeFormatFormat")
	vm.jsObjectItems[objID] = obj
	vm.jsIntlDateTimeFormatItems[objID] = &jsIntlDateTimeFormatObject{localeTag: localeTag, layout: layout}
	vm.jsPropertyItems[objID] = make(map[string]jsPropertyDescriptor, 6)
	vm.jsSetDescriptor(objID, "format", jsPropertyDescriptor{
		Value:        obj["format"],
		HasValue:     true,
		Enumerable:   false,
		Configurable: true,
		Writable:     true,
	})
	return Value{Type: VTJSObject, Num: objID}
}

// jsIntlCreateNumberFormat allocates one Intl.NumberFormat instance with normalized locale state.
func (vm *VM) jsIntlCreateNumberFormat(args []Value) Value {
	profile, localeTag := vm.jsIntlResolveLocaleProfile(jsArgOrUndefined(args, 0))
	options := Value{Type: VTJSUndefined}
	if len(args) > 1 {
		options = args[1]
	}
	style := strings.ToLower(strings.TrimSpace(vm.jsIntlOptionString(options, "style")))
	if style == "" {
		style = "decimal"
	}
	useGrouping := true
	if v, ok := vm.jsIntlOptionBool(options, "useGrouping"); ok {
		useGrouping = v
	}
	digits := vm.jsIntlResolveFractionDigits(style, options)
	currencyCode := strings.TrimSpace(vm.jsIntlOptionString(options, "currency"))
	currencySymbol, currencySpacing := jsIntlCurrencySymbol(currencyCode, profile)
	objID := vm.allocJSID()
	obj := make(map[string]Value, 8)
	obj["__js_type"] = NewString("Intl.NumberFormat")
	obj["__js_ctor"] = NewString("IntlNumberFormat")
	obj["__js_intl_locale"] = NewString(localeTag)
	obj["__js_intl_style"] = NewString(style)
	obj["__js_intl_use_grouping"] = NewBool(useGrouping)
	obj["__js_intl_digits"] = NewInteger(int64(digits))
	obj["__js_intl_currency_symbol"] = NewString(currencySymbol)
	obj["__js_intl_currency_spacing"] = NewString(currencySpacing)
	obj["format"] = vm.jsIntlCreateFormatFunction(objID, "IntlNumberFormatFormat")
	vm.jsObjectItems[objID] = obj
	vm.jsIntlNumberFormatItems[objID] = &jsIntlNumberFormatObject{
		localeTag:       localeTag,
		style:           style,
		digits:          digits,
		useGrouping:     useGrouping,
		currencySymbol:  currencySymbol,
		currencySpacing: currencySpacing,
	}
	vm.jsPropertyItems[objID] = make(map[string]jsPropertyDescriptor, 6)
	vm.jsSetDescriptor(objID, "format", jsPropertyDescriptor{
		Value:        obj["format"],
		HasValue:     true,
		Enumerable:   false,
		Configurable: true,
		Writable:     true,
	})
	return Value{Type: VTJSObject, Num: objID}
}

// jsIntlDateTimeFormatFormat renders one date value using the formatter state stored on the Intl instance.
func (vm *VM) jsIntlDateTimeFormatFormat(callee Value, thisVal Value, args []Value) Value {
	ownerID, ok := vm.jsIntlMethodOwnerID(callee)
	if !ok {
		ownerID = thisVal.Num
	}
	obj, ok := vm.jsObjectItems[ownerID]
	if !ok {
		return Value{Type: VTJSUndefined}
	}
	inst := vm.jsIntlDateTimeFormatItems[ownerID]
	localeTag := obj["__js_intl_locale"].String()
	layout := obj["__js_intl_layout"].String()
	if inst != nil {
		localeTag = inst.localeTag
		layout = inst.layout
	}
	profile := builtinLocaleProfileForTag(localeTag)
	value := time.Now().In(builtinCurrentLocation(vm))
	if len(args) > 0 {
		candidate := resolveCallable(vm, args[0])
		if candidate.Type == VTJSUndefined || candidate.Type == VTNull {
			value = time.Now().In(builtinCurrentLocation(vm))
		} else {
			parsed := valueToTimeInLocale(vm, candidate)
			if parsed.IsZero() {
				value = time.Now().In(builtinCurrentLocation(vm))
			} else {
				value = parsed.In(builtinCurrentLocation(vm))
			}
		}
	}
	if strings.TrimSpace(layout) == "" {
		layout = profile.shortDateLayout + " " + profile.longTimeLayout
	}
	return NewString(localizedFormat(value, layout, profile))
}

// jsIntlNumberFormatFormat renders one number using the formatter state stored on the Intl instance.
func (vm *VM) jsIntlNumberFormatFormat(callee Value, thisVal Value, args []Value) Value {
	ownerID, ok := vm.jsIntlMethodOwnerID(callee)
	if !ok {
		ownerID = thisVal.Num
	}
	obj, ok := vm.jsObjectItems[ownerID]
	if !ok {
		return Value{Type: VTJSUndefined}
	}
	inst := vm.jsIntlNumberFormatItems[ownerID]
	localeTag := obj["__js_intl_locale"].String()
	style := strings.ToLower(obj["__js_intl_style"].String())
	digits := int(obj["__js_intl_digits"].Num)
	useGrouping := obj["__js_intl_use_grouping"].Num != 0
	currencySymbol := obj["__js_intl_currency_symbol"].String()
	currencySpacing := obj["__js_intl_currency_spacing"].String()
	if inst != nil {
		localeTag = inst.localeTag
		style = inst.style
		digits = inst.digits
		useGrouping = inst.useGrouping
		currencySymbol = inst.currencySymbol
		currencySpacing = inst.currencySpacing
	}
	profile := builtinLocaleProfileForTag(localeTag)
	value := 0.0
	if len(args) > 0 {
		value = vm.jsToNumber(args[0]).Flt
	}
	switch style {
	case "currency":
		formatted := localizedNumberString(value, digits, profile, useGrouping)
		if strings.TrimSpace(currencySymbol) == "" {
			currencySymbol = profile.currencySymbol
			currencySpacing = profile.currencySpacing
		}
		result := currencySymbol + currencySpacing + formatted
		if value < 0 {
			return NewString("-" + result)
		}
		return NewString(result)
	case "percent":
		formatted := localizedNumberString(value*100, digits, profile, useGrouping)
		return NewString(formatted + "%")
	default:
		return NewString(localizedNumberString(value, digits, profile, useGrouping))
	}
}

// jsIntlResolveLocaleProfile resolves one locale profile from Intl constructor arguments.
func (vm *VM) jsIntlResolveLocaleProfile(locales Value) (builtinLocaleProfile, string) {
	defaultProfile := builtinLocaleProfileForVM(vm)
	tag := vm.jsIntlLocaleTagFromValue(locales)
	if strings.TrimSpace(tag) == "" {
		return defaultProfile, defaultProfile.tag
	}
	profile := builtinLocaleProfileForTag(tag)
	return profile, profile.tag
}

// jsIntlLocaleTagFromValue extracts the first usable locale tag from one Intl locales argument.
func (vm *VM) jsIntlLocaleTagFromValue(locales Value) string {
	switch locales.Type {
	case VTJSUndefined, VTNull, VTEmpty:
		return ""
	case VTString:
		return jsNormalizeLocaleTag(locales.Str)
	case VTArray:
		if locales.Arr == nil {
			return ""
		}
		for i := 0; i < len(locales.Arr.Values); i++ {
			tag := jsNormalizeLocaleTag(locales.Arr.Values[i].String())
			if tag != "" {
				return tag
			}
		}
		return ""
	case VTJSObject, VTJSFunction:
		if length, ok, deferred := vm.jsArrayLikeLength(locales); ok && !deferred {
			for i := 0; i < length; i++ {
				if v, exists := vm.jsArrayLikeGetIndex(locales, i); exists {
					tag := jsNormalizeLocaleTag(v.String())
					if tag != "" {
						return tag
					}
				}
			}
		}
		return jsNormalizeLocaleTag(locales.String())
	default:
		return jsNormalizeLocaleTag(locales.String())
	}
}

// jsIntlOptionString reads one string option from an Intl options object.
func (vm *VM) jsIntlOptionString(options Value, name string) string {
	if options.Type != VTJSObject && options.Type != VTJSFunction {
		return ""
	}
	value, deferred := vm.jsMemberGet(options, name)
	if deferred || value.Type == VTJSUndefined || value.Type == VTNull {
		return ""
	}
	return strings.TrimSpace(value.String())
}

// jsIntlOptionBool reads one boolean option from an Intl options object.
func (vm *VM) jsIntlOptionBool(options Value, name string) (bool, bool) {
	if options.Type != VTJSObject && options.Type != VTJSFunction {
		return false, false
	}
	value, deferred := vm.jsMemberGet(options, name)
	if deferred || value.Type == VTJSUndefined || value.Type == VTNull {
		return false, false
	}
	switch value.Type {
	case VTBool:
		return value.Num != 0, true
	case VTInteger, VTDouble:
		return vm.jsToNumber(value).Flt != 0, true
	default:
		return strings.TrimSpace(value.String()) != "", true
	}
}

// jsIntlResolveFractionDigits selects the formatter precision for Intl.NumberFormat.
func (vm *VM) jsIntlResolveFractionDigits(style string, options Value) int {
	defaultDigits := 2
	switch style {
	case "percent":
		defaultDigits = 0
	case "currency":
		defaultDigits = 2
	}
	if digits, ok := vm.jsIntlOptionInt(options, "maximumFractionDigits"); ok {
		return digits
	}
	if digits, ok := vm.jsIntlOptionInt(options, "minimumFractionDigits"); ok {
		return digits
	}
	if digits, ok := vm.jsIntlOptionInt(options, "fractionDigits"); ok {
		return digits
	}
	return defaultDigits
}

// jsIntlOptionInt reads one integer option from an Intl options object.
func (vm *VM) jsIntlOptionInt(options Value, name string) (int, bool) {
	if options.Type != VTJSObject && options.Type != VTJSFunction {
		return 0, false
	}
	value, deferred := vm.jsMemberGet(options, name)
	if deferred || value.Type == VTJSUndefined || value.Type == VTNull {
		return 0, false
	}
	return int(math.Round(vm.jsToNumber(value).Flt)), true
}

// jsIntlResolveDateTimeLayout builds one locale-aware layout string from Intl.DateTimeFormat options.
func (vm *VM) jsIntlResolveDateTimeLayout(profile builtinLocaleProfile, options Value) string {
	dateStyle := strings.ToLower(vm.jsIntlOptionString(options, "dateStyle"))
	timeStyle := strings.ToLower(vm.jsIntlOptionString(options, "timeStyle"))
	hasDateTokens := vm.jsIntlHasDateTimeToken(options, "year") || vm.jsIntlHasDateTimeToken(options, "month") || vm.jsIntlHasDateTimeToken(options, "day") || vm.jsIntlHasDateTimeToken(options, "weekday")
	hasTimeTokens := vm.jsIntlHasDateTimeToken(options, "hour") || vm.jsIntlHasDateTimeToken(options, "minute") || vm.jsIntlHasDateTimeToken(options, "second")
	hour12, _ := vm.jsIntlOptionBool(options, "hour12")

	dateLayout := ""
	switch dateStyle {
	case "full", "long":
		dateLayout = localizedLongDateLayout(profile)
	case "medium", "short":
		dateLayout = profile.shortDateLayout
	}
	if dateLayout == "" && hasDateTokens {
		dateLayout = profile.shortDateLayout
	}

	timeLayout := ""
	switch timeStyle {
	case "full", "long":
		timeLayout = profile.longTimeLayout
	case "medium", "short":
		timeLayout = profile.shortTimeLayout
	}
	if timeLayout == "" && hasTimeTokens {
		if hour12 && strings.Contains(profile.shortTimeLayout, "PM") {
			timeLayout = profile.shortTimeLayout
		} else if hour12 && profile.tag != "en-US" {
			timeLayout = profile.shortTimeLayout
		} else {
			timeLayout = profile.longTimeLayout
		}
	}

	if dateLayout != "" && timeLayout != "" {
		return dateLayout + " " + timeLayout
	}
	if dateLayout != "" {
		return dateLayout
	}
	if timeLayout != "" {
		return timeLayout
	}
	return profile.shortDateLayout + " " + profile.longTimeLayout
}

// jsIntlHasDateTimeToken reports whether one Intl.DateTimeFormat option is present.
func (vm *VM) jsIntlHasDateTimeToken(options Value, name string) bool {
	if options.Type != VTJSObject && options.Type != VTJSFunction {
		return false
	}
	value, deferred := vm.jsMemberGet(options, name)
	return !deferred && value.Type != VTJSUndefined && value.Type != VTNull && strings.TrimSpace(value.String()) != ""
}

// jsIntlMethodOwnerID resolves the hidden owner object ID for one Intl formatter method.
func (vm *VM) jsIntlMethodOwnerID(callee Value) (int64, bool) {
	if callee.Type != VTJSFunction {
		return 0, false
	}
	obj, ok := vm.jsObjectItems[callee.Num]
	if !ok {
		return 0, false
	}
	owner, ok := obj["__js_intl_owner"]
	if !ok || owner.Type != VTInteger {
		return 0, false
	}
	return owner.Num, true
}

// jsNormalizeLocaleTag normalizes one locale string for matcher lookup.
func jsNormalizeLocaleTag(tag string) string {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return ""
	}
	tag = strings.ReplaceAll(tag, "_", "-")
	return tag
}

// jsIntlCurrencySymbol resolves the currency symbol and spacing for one Intl.NumberFormat currency
// code. It delegates to builtinCurrencySymbolForCode so that builtinLocaleProfiles in
// locale_format.go is the single source of truth shared by both VBScript and JScript.
func jsIntlCurrencySymbol(code string, profile builtinLocaleProfile) (string, string) {
	if strings.TrimSpace(code) == "" {
		return profile.currencySymbol, profile.currencySpacing
	}
	if symbol, spacing, ok := builtinCurrencySymbolForCode(code); ok {
		return symbol, spacing
	}
	// Unknown currency code: use the code itself as the display symbol.
	return strings.ToUpper(strings.TrimSpace(code)), " "
}
