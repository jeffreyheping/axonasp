# Native File I/O (VB6 Compatibility)

AxonASP now supports native VB6-style file I/O operations in the CLI environment. These operations provide faster and more direct file access compared to the standard `FileSystemObject` (FSO), which remains available for all environments including the HTTP server and FastCGI.

**NOTE:** Native File I/O is currently restricted to the **CLI** environment only. Attempting to use these statements in the HTTP server or FastCGI modules will result in a `Permission Denied` error.

## Availability

| Feature | CLI | HTTP Server | FastCGI | WASM |
| :--- | :---: | :---: | :---: | :---: |
| Native File I/O | YES | NO | NO | NO |

## Statements and Functions

### Open Statement
Enables I/O to a file.

**Syntax:**
`Open pathname For mode [Access access] [lock] As [#]filenumber`

- `pathname`: String expression specifying the file name. Can include directory and drive.
- `mode`: Keyword specifying the file mode: `Input`, `Output`, `Append`, `Binary`, or `Random`.
- `access`: (Optional) Keyword specifying the operations permitted on the open file: `Read`, `Write`, or `Read Write`.
- `lock`: (Optional) Keyword specifying the operations permitted on the open file by other processes: `Shared`, `Lock Read`, `Lock Write`, and `Lock Read Write`.
- `filenumber`: A valid file number in the range 1 to 511, inclusive. Use the `FreeFile` function to obtain the next available file number.

### Print # Statement
Writes display-formatted data to a sequential file.

**Syntax:**
`Print #filenumber, [outputlist]`

- `filenumber`: Any valid file number.
- `outputlist`: (Optional) Expression or list of expressions to write to the file.

### Write # Statement
Writes raw data to a sequential file. Similar to `Print #`, but inserts commas between items and puts quotes around strings.

**Syntax:**
`Write #filenumber, [outputlist]`

### Close Statement
Ends I/O to a file. Always close your files to prevent data loss and free system resources.

**Syntax:**
`Close [#filenumber1, #filenumber2, ...]`

If no file numbers are provided, all open files are closed.

### Line Input # Statement
Reads a single line from an open sequential file and assigns it to a String variable.

**Syntax:**
`Line Input #filenumber, varname`

### FreeFile Function
Returns an Integer representing the next file number available for use by the `Open` statement.

**Syntax:**
`variable = FreeFile()`

## Examples

### Writing and Reading a Text File

```vbscript
Dim fNum, path, lineData
fNum = FreeFile
path = "C:\temp\data.txt"

' Open for Output (creates or overwrites)
Open path For Output As #fNum
Print #fNum, "Hello AxonASP"
Print #1, "Another Line"
Close #fNum

' Open for Input
Open path For Input As #fNum
Line Input #fNum, lineData
Response.Write "First line: " & lineData
Close #fNum
```

### Appending to a File

```vbscript
Open "log.txt" For Append As #1
Print #1, Now & " - Application Started"
Close #1
```

## Error Codes

The following error codes may be raised during File I/O operations:

- `52`: Bad file name or number.
- `53`: File not found.
- `70`: Permission denied (raised when used outside CLI).
- `76`: Path not found.
