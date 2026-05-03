# IndexPath Property

## Overview
Gets or sets a String value specifying the file system path where the **G3SEARCH** engine stores its persistent index files.

## Syntax
```asp
' Get the current value
path = search.IndexPath

' Set a new value
search.IndexPath = "../temp/mysearch.index"
```

## Return Values
Returns a **String** representing the absolute or relative path to the directory that will contain the search index.

## Remarks
- The **IndexPath** must be set before calling either the **BuildIndex** or **Search** methods. 
- If the directory does not exist, the **G3SEARCH** engine will attempt to create it during the indexing process, provided the server process has sufficient permissions.
- It is recommended to use **Server.MapPath** to resolve relative paths to absolute system paths.

## Code Example
```asp
<%
Dim search
Set search = Server.CreateObject("G3SEARCH")

' Set the location for the index files
search.IndexPath = Server.MapPath("../temp/mysearch.index")

Response.Write "Index will be stored at: " & search.IndexPath

Set search = Nothing
%>
```