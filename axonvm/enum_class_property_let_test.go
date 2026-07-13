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
	"testing"
)

// TestEnumBareMemberAccess verifies that bare enum member names (e.g. "Red")
// resolve to their integer values at compile time.
func TestEnumBareMemberAccess(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Response.Write Red
Response.Write ","
Response.Write Green
Response.Write ","
Response.Write Blue
%>`
	out := runASPSourceForTest(t, source)
	if out != "0,1,2" {
		t.Fatalf("expected 0,1,2 got %q", out)
	}
}

// TestEnumPrefixedMemberAccess verifies that "EnumType.Member" syntax
// resolves to the correct integer value at compile time.
func TestEnumPrefixedMemberAccess(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Response.Write MyColor.Red
Response.Write ","
Response.Write MyColor.Green
Response.Write ","
Response.Write MyColor.Blue
%>`
	out := runASPSourceForTest(t, source)
	if out != "0,1,2" {
		t.Fatalf("expected 0,1,2 got %q", out)
	}
}

// TestEnumPrefixedMemberInExpression verifies that "EnumType.Member" works
// inside arithmetic expressions.
func TestEnumPrefixedMemberInExpression(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Dim x
x = MyColor.Green + MyColor.Blue
Response.Write x
%>`
	out := runASPSourceForTest(t, source)
	if out != "3" {
		t.Fatalf("expected 3 got %q", out)
	}
}

// TestEnumPropertyLetTyped verifies that an Enum value assigned to a
// Property Let whose parameter is typed "As <EnumType>" works correctly.
// This is the exact reproduction case from the bug report.
func TestEnumPropertyLetTyped(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Class Foo
    Private m_Color As MyColor

    Public Property Get Color As MyColor
        Color = m_Color
    End Property

    Public Property Let Color(v As MyColor)
        m_Color = v
    End Property

    Public Function Show
        Response.Write "Color = " & m_Color
    End Function
End Class

Dim f As Foo
Set f = New Foo
f.Color = MyColor.Green
f.Show
%>`
	out := runASPSourceForTest(t, source)
	if out != "Color = 1" {
		t.Fatalf("expected 'Color = 1' got %q", out)
	}
}

// TestEnumPropertyLetLiteralInt verifies that a literal integer assigned
// to a Property Let whose parameter is typed "As <EnumType>" works correctly
// (VBA parity: literal integers satisfy enum-typed parameters).
func TestEnumPropertyLetLiteralInt(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Class Foo
    Private m_Color As MyColor

    Public Property Get Color As MyColor
        Color = m_Color
    End Property

    Public Property Let Color(v As MyColor)
        m_Color = v
    End Property

    Public Function Show
        Response.Write "Color = " & m_Color
    End Function
End Class

Dim f As Foo
Set f = New Foo
f.Color = 2
f.Show
%>`
	out := runASPSourceForTest(t, source)
	if out != "Color = 2" {
		t.Fatalf("expected 'Color = 2' got %q", out)
	}
}

// TestEnumPropertyLetAllValues verifies that all enum values can be
// assigned and retrieved through Property Let/Get.
func TestEnumPropertyLetAllValues(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Class Foo
    Private m_Color As MyColor

    Public Property Get Color As MyColor
        Color = m_Color
    End Property

    Public Property Let Color(v As MyColor)
        m_Color = v
    End Property
End Class

Dim f As Foo
Set f = New Foo

f.Color = MyColor.Red
Response.Write f.Color & ","

f.Color = MyColor.Green
Response.Write f.Color & ","

f.Color = MyColor.Blue
Response.Write f.Color & ","
%>`
	out := runASPSourceForTest(t, source)
	if out != "0,1,2," {
		t.Fatalf("expected '0,1,2,' got %q", out)
	}
}

// TestEnumSubParameterTyped verifies that an Enum value passed to a Sub
// whose parameter is typed "As <EnumType>" works correctly.
func TestEnumSubParameterTyped(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Sub PrintColor(c As MyColor)
    Response.Write "C=" & c
End Sub

PrintColor MyColor.Green
%>`
	out := runASPSourceForTest(t, source)
	if out != "C=1" {
		t.Fatalf("expected 'C=1' got %q", out)
	}
}

// TestEnumFunctionParameterTyped verifies that an Enum value passed to a
// Function whose parameter is typed "As <EnumType>" works correctly.
func TestEnumFunctionParameterTyped(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Function GetColorName(c As MyColor)
    If c = MyColor.Red Then
        GetColorName = "Red"
    ElseIf c = MyColor.Green Then
        GetColorName = "Green"
    ElseIf c = MyColor.Blue Then
        GetColorName = "Blue"
    Else
        GetColorName = "Unknown"
    End If
End Function

Response.Write GetColorName(MyColor.Red) & ","
Response.Write GetColorName(MyColor.Green) & ","
Response.Write GetColorName(MyColor.Blue)
%>`
	out := runASPSourceForTest(t, source)
	if out != "Red,Green,Blue" {
		t.Fatalf("expected 'Red,Green,Blue' got %q", out)
	}
}

// TestEnumVariableAssignment verifies that Enum values can be assigned to
// variables typed "As <EnumType>".
func TestEnumVariableAssignment(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Dim x As MyColor
x = MyColor.Green
Response.Write x
%>`
	out := runASPSourceForTest(t, source)
	if out != "1" {
		t.Fatalf("expected '1' got %q", out)
	}
}

// TestEnumClassFieldAssignment verifies that Enum values can be assigned
// directly to class fields typed "As <EnumType>".
func TestEnumClassFieldAssignment(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
    Blue = 2
End Enum

Class Foo
    Public Color As MyColor
End Class

Dim f As Foo
Set f = New Foo
f.Color = MyColor.Blue
Response.Write f.Color
%>`
	out := runASPSourceForTest(t, source)
	if out != "2" {
		t.Fatalf("expected '2' got %q", out)
	}
}

// TestEnumCustomValues verifies that Enum members with explicit values work
// correctly via prefixed syntax.
func TestEnumCustomValues(t *testing.T) {
	source := `<%
Enum Colors
    Red = 10
    Green = 20
    Blue = 30
End Enum

Response.Write Colors.Red & ","
Response.Write Colors.Green & ","
Response.Write Colors.Blue
%>`
	out := runASPSourceForTest(t, source)
	if out != "10,20,30" {
		t.Fatalf("expected '10,20,30' got %q", out)
	}
}

// TestEnumMultipleEnums verifies that multiple Enum types with the same
// member names do not interfere with each other.
func TestEnumMultipleEnums(t *testing.T) {
	source := `<%
Enum ColorA
    Red = 1
    Green = 2
End Enum

Enum ColorB
    Red = 10
    Green = 20
End Enum

Response.Write ColorA.Red & "," & ColorA.Green & ","
Response.Write ColorB.Red & "," & ColorB.Green
%>`
	out := runASPSourceForTest(t, source)
	if out != "1,2,10,20" {
		t.Fatalf("expected '1,2,10,20' got %q", out)
	}
}

// TestEnumPropertyLetWithLegacyClass verifies that Property Let without
// typed parameters still works (regression test).
func TestEnumPropertyLetWithLegacyClass(t *testing.T) {
	source := `<%
Class Bar
    Private m_val

    Public Property Get Value
        Value = m_val
    End Property

    Public Property Let Value(v)
        m_val = v
    End Property
End Class

Dim b As Bar
Set b = New Bar
b.Value = 42
Response.Write b.Value
%>`
	out := runASPSourceForTest(t, source)
	if out != "42" {
		t.Fatalf("expected '42' got %q", out)
	}
}

// TestEnumAsIntComparison verifies that Enum values compare correctly
// with integer literals when the enum type is used as a constraint.
func TestEnumAsIntComparison(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
End Enum

Dim x As MyColor
x = MyColor.Green

If x = 1 Then
    Response.Write "OK"
Else
    Response.Write "FAIL"
End If
%>`
	out := runASPSourceForTest(t, source)
	if out != "OK" {
		t.Fatalf("expected 'OK' got %q", out)
	}
}

// TestEnumClassPropertyLetMixed verifies that a class with both typed
// and untyped properties works correctly (regression test).
func TestEnumClassPropertyLetMixed(t *testing.T) {
	source := `<%
Enum MyColor
    Red = 0
    Green = 1
End Enum

Class Mixed
    Private m_Color As MyColor
    Private m_Name

    Public Property Get Color As MyColor
        Color = m_Color
    End Property

    Public Property Let Color(v As MyColor)
        m_Color = v
    End Property

    Public Property Get Name
        Name = m_Name
    End Property

    Public Property Let Name(v)
        m_Name = v
    End Property
End Class

Dim m As Mixed
Set m = New Mixed
m.Color = MyColor.Green
m.Name = "test"
Response.Write m.Color & ":" & m.Name
%>`
	out := runASPSourceForTest(t, source)
	if out != "1:test" {
		t.Fatalf("expected '1:test' got %q", out)
	}
}
