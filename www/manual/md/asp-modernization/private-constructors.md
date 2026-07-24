# Private Constructors and Encapsulation

AxonASP introduces strict class lifecycle encapsulation for VBScript classes, closing a common loophole in the Singleton pattern. By declaring `Class_Initialize` as `Private`, you can prevent external instantiation of a class using the `New` keyword, while still allowing the class itself to create instances internally (e.g., through a factory method).

## Standard VBScript Behavior vs AxonASP

In Classic ASP (IIS), declaring `Private Sub Class_Initialize` only affected the visibility of the constructor method itself. It did not prevent external code from successfully executing `Set obj = New Singleton` and obtaining a valid object reference. This broke the encapsulation needed for robust design patterns like Singleton or Factory.

AxonASP correctly enforces instantiation blocking when `Class_Initialize` is `Private`.

## Expected Behavior

When a class constructor (`Class_Initialize`) is marked as `Private`:

*   **Internal Instantiation:** The class itself can still create instances of itself. This is typically done inside a Public method of the same class (or in AxonASP, via internal host mechanisms or factory functions acting on behalf of the class).
*   **External Instantiation:** When external code attempts to use `New` on this class, the VM executes the `Class_Initialize` side effects (matching twinBASIC behavior), but the assignment to the external variable fails with a runtime error `91` ("Object variable or With block variable not set"). The object reference is never returned to the external caller.

## Example: Singleton Pattern Enforcement

```vbscript
Class DatabaseConnection
    ' The constructor is private, preventing 'New DatabaseConnection' from outside
    Private Sub Class_Initialize()
        ' Initialize connection parameters
    End Sub

    ' An internal public method can still instantiate the class
    Public Function CloneConnection()
        ' This is allowed because the caller scope is internal to DatabaseConnection
        Set CloneConnection = New DatabaseConnection
    End Function
End Class

' External instantiation attempt
On Error Resume Next
Dim conn
Set conn = New DatabaseConnection

' In AxonASP, Err.Number will be 91 (Object variable or With block variable not set)
If Err.Number <> 0 Then
    Response.Write "Cannot instantiate DatabaseConnection directly!"
End If
```

## Summary of Rules

1. Classes with `Public` (or implicit) constructors behave exactly as they did in IIS Classic ASP.
2. The scope/context check during instantiation is highly optimized within the AxonASP VM and does not cause heap allocations or degrade execution speed.
3. If instantiation fails due to a Private constructor, any partially constructed object or stack frame is properly cleaned up to prevent memory leaks.
