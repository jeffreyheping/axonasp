# Extension Property

## Overview
Gets or sets a String value that defines the file type filter used by the **G3SEARCH** object when scanning the document directory.

## Syntax
```asp
' Get the current value
ext = search.Extension

' Set a new value
search.Extension = ".txt"
```

## Return Values
Returns a **String** representing the file extension (e.g., ".md", ".txt", ".htm"). The default value is ".md".

## Remarks
- During the **BuildIndex** operation, only files with the specified extension will have their contents read and indexed. 
- If this property is set to an empty string, all files in the **DocsPath** will be indexed regardless of their extension.
- The extension can be provided with or without the leading period; the object will normalize it automatically.

## Code Example
```asp
<%
Dim search
Set search = Server.CreateObject("G3SEARCH")

' Index only text files
search.Extension = ".txt"

' Build the index (assuming IndexPath and DocsPath are set)
' search.BuildIndex()

Set search = Nothing
%>
```