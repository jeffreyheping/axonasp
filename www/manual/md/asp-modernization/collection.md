# Delegate For Each Iteration using Collection Object

## Overview
AxonASP introduces the built-in **Collection** object and **For Each** delegation support via COM-like interfaces. This feature enables developers to construct custom collection classes in VBScript and iterate them directly using the native `For Each...Next` loop syntax. This implementation can't be disabled during compilation, and although it is named in the source as a library object as it uses the Server.CreateObject call, it is not a traditional AxonASP library object.

Custom VBScript classes can delegate iteration to an internal `Collection` instance by marking a property or method with the compiler tag `[DispId(-4)]`.

## Syntax
To instantiate the built-in Collection object:
```vbscript
Set col = Server.CreateObject("Collection")
```

To invoke members on the Collection object:
```vbscript
col.Add item
col.Remove index
countVal = col.Count
itemVal = col.Item(index)
Set enumVal = col.[_NewEnum]
```

To delegate iteration in a custom VBScript class:
```vbscript
Class MyClass
    Private m_Items
    Private Sub Class_Initialize()
        Set m_Items = Server.CreateObject("Collection")
    End Sub
    
    [DispId(-4)]
    Public Property Get NewEnum()
        Set NewEnum = m_Items.[_NewEnum]
    End Property
End Class
```

## Parameters and Arguments
- **item**: Any valid VBScript type (scalars, records, arrays, or objects) to append to the collection. Required.
- **index**: An **Integer** representing a 1-based index pointing to the item's position in the collection. Must satisfy `1 <= index <= Count`. Required.

## Return Values
- **Server.CreateObject("Collection")**: Returns a **VTNativeObject** handle referencing the newly created collection instance.
- **Add**: Returns **Empty**.
- **Remove**: Returns **Empty**.
- **Count**: Returns an **Integer** representing the total number of items stored in the collection.
- **Item**: Returns the **Value** stored at the specified 1-based index.
- **[_NewEnum]** / **NewEnum**: Returns a **VTNativeObject** handle to an opaque enumerator snapshot of the collection.

## Remarks
- **1-Based Indexing**: The `Collection` object enforces 1-based indexing for `Item(index)` and `Remove(index)` to maintain parity with legacy Visual Basic 6 collection behavior.
- **Thread & GC Safety**: The `[_NewEnum]` property creates a stable snapshot of the collection items at the moment of request. Modifying the collection inside a `For Each` loop will not cause iterator invalidation or memory corruption.
- **DispId(-4) Metadata**: The compiler parses `[DispId(-4)]` attribute tags preceding property or method declarations in VBScript classes. This registers the tagged member under the special runtime identifier `__newenum__` allowing the VM to resolve and execute it when the class instance is iterated.
- **Compatibility**: If `For Each` is called on a custom class that lacks a `[DispId(-4)]` attribute or `NewEnum` property, the VM throws an **Invalid procedure call or argument** runtime exception.

## Code Example
```asp
<%
' Define a custom VBScript collection class
Class TaskList
    Private m_Items

    Private Sub Class_Initialize()
        Set m_Items = Server.CreateObject("Collection")
    End Sub

    Public Sub AddTask(ByVal taskName)
        m_Items.Add taskName
    End Sub

    [DispId(-4)]
    Public Property Get NewEnum()
        Set NewEnum = m_Items.[_NewEnum]
    End Property
End Class

Dim list, task
Set list = New TaskList

' Populate the collection
list.AddTask "Initialize Server"
list.AddTask "Load Configuration"
list.AddTask "Execute VM Bytecode"

' Iterate using the native For Each loop
For Each task In list
    Response.Write "Task: " & task & "<br>"
Next

Set list = Nothing
%>
```
