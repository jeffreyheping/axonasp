# DocsPath Property

## Overview
Gets or sets a String value representing the root directory containing the documents that the **G3SEARCH** object will scan and index.

## Syntax
```asp
' Get the current value
path = search.DocsPath

' Set a new value
search.DocsPath = "../content/articles/"
```

## Return Values
Returns a **String** representing the path to the directory containing the source files for indexing.

## Remarks
- The indexing process initiated by **BuildIndex** is recursive; all subdirectories within the **DocsPath** will be scanned for files matching the **Extension** property.
- This property must be correctly set before calling the **BuildIndex** method.

## Code Example
```asp
<%
Dim search
Set search = Server.CreateObject("G3SEARCH")

' Define the source directory for documents
search.DocsPath = Server.MapPath("../content/articles/")

Response.Write "Indexing documents from: " & search.DocsPath

Set search = Nothing
%>
```