# BuildIndex Method

## Overview
The **BuildIndex** method triggers a recursive scan of the specified document directory and populates the search index with the content and filenames of the matching documents.

## Syntax
```asp
search.BuildIndex()
```

## Parameters and Arguments
This method does not take any arguments.

## Return Values
Returns **Empty** upon successful completion.

## Remarks
- Before calling **BuildIndex**, the **IndexPath** and **DocsPath** properties must be properly configured. 
- The method performs the following actions:
  1. Recursively traverses the **DocsPath**.
  2. Identifies files matching the **Extension** filter.
  3. Reads the full text content of each file.
  4. Updates or creates entries in the search index at **IndexPath**.
- Note that large document sets may take several seconds to index. It is recommended to run this operation during maintenance windows or via an administrative interface.

## Code Example
```asp
<%
Dim search
Set search = Server.CreateObject("G3SEARCH")

search.IndexPath = Server.MapPath("../temp/index")
search.DocsPath = Server.MapPath("../docs")
search.Extension = ".md"

' Build the index
search.BuildIndex()

Response.Write "Indexing complete."

Set search = Nothing
%>
```