# Search Method

## Overview
Use **Search** to query the existing G3Pix AxonASP G3SEARCH index and retrieve matching documents with their relevance scores.

## Syntax
```asp
results = search.Search(term)
```

## Parameters and Arguments
- **term** (String, Required): A query term to match against indexed document content.

## Return Values
Returns a native VBScript **two-dimensional array** where each row contains exactly two values:
1. **Filename** as **String** (`results(i)(0)`)
2. **Score** as **Double** (`results(i)(1)`)

If no matches are found, the method returns an empty array.

## Remarks
- Set **IndexPath** to a valid index directory before calling this method.
- Set `g3search.g3search_enabled = true` in `config/axonasp.toml` before creating and using the object.
- The search runs against the `content` field and returns rows ordered by the internal relevance ranking from the search engine.

## Code Example
```asp
<%
Dim search, results, i
Dim filename, score
Set search = Server.CreateObject("G3SEARCH")

search.IndexPath = Server.MapPath("../temp/index")

' Execute search for the term "AxonASP"
results = search.Search("AxonASP")

If IsArray(results) Then
    If UBound(results) >= LBound(results) Then
        Response.Write "Found " & (UBound(results) - LBound(results) + 1) & " matches:<br>"
        For i = LBound(results) To UBound(results)
            If IsArray(results(i)) Then
                filename = CStr(results(i)(0))
                score = CDbl(results(i)(1))
                Response.Write " - " & filename & " | Score: " & CStr(score) & "<br>"
            End If
        Next
    Else
        Response.Write "No matches found."
    End If
End If

Set search = Nothing
%>
```