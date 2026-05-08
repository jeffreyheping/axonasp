# Basic Page Creation & Routing

Building a page with G3AxonLive requires wiring a backend server script (VBScript or Server-Side JavaScript) with a frontend HTML layout that includes the `g3axonlive.js` client bridge.

AxonLive acts as a master controller. Every request starts by calling the `InitPage` method. If the request is a standard page load, the server renders the full HTML. If the request is an asynchronous event triggered by a component, the server processes the event logic, registers HTML patches, and cleanly halts execution via `EndAsyncResponse`.

## Creating Your First AxonLive Page

Below is a complete, step-by-step example demonstrating how to build a reactive counter. We provide examples in both Server-Side JavaScript (JScript) and VBScript.

### Server-Side JavaScript Example

```javascript
<%@ Language="Javascript" %>
<%
var AxonLive = Server.CreateObject("G3AXONLIVE");
AxonLive.InitPage();

var sessionID = Session.SessionID;

// Helper to safely read numbers from component state
function getCount(sid, cid, prop) {
    var val = AxonLive.GetComponentProperty(sid, cid, prop);
    return (val === undefined || val === null || val === "") ? 0 : parseInt(val, 10);
}

var currentCount = getCount(sessionID, "main", "clicks");

// Handle Asynchronous Requests
if (AxonLive.IsAsyncRequest) {
    var compID = AxonLive.EventComponentID;

    switch (compID) {
        case "btnIncrement":
            currentCount++;
            break;
        case "btnDecrement":
            currentCount--;
            break;
    }

    // Persist the state
    AxonLive.SetComponentProperty(sessionID, "main", "clicks", String(currentCount));

    // Register the component patch
    var newHtml = "<div id=\"counterDisplay\" class=\"pill pill-primary\">Current Count: " + currentCount + "</div>";
    AxonLive.RegisterComponent("counterDisplay", newHtml);

    // End the response (halts further execution)
    AxonLive.EndAsyncResponse();
}
%>
<!DOCTYPE html>
<html>
<head>
    <title>AxonLive JavaScript Counter</title>
    <link rel="stylesheet" href="/css/axonasp.css">
</head>
<body>
    <div id="content" style="padding: 20px;">
        <h2>AxonLive Counter</h2>
        
        <!-- The component that will be dynamically replaced -->
        <div id="counterDisplay" class="pill pill-primary">Current Count: <%= currentCount %></div>
        
        <div style="margin-top: 15px;">
            <button id="btnIncrement" class="btn btn-primary" data-g3al-id="btnIncrement" data-g3al-event="click" data-g3al-event-name="onclick">
                Increment
            </button>
            <button id="btnDecrement" class="btn btn-danger" data-g3al-id="btnDecrement" data-g3al-event="click" data-g3al-event-name="onclick">
                Decrement
            </button>
        </div>
    </div>

    <script src="/axonlive/g3axonlive.js"></script>
    <script>
        G3AxonLive.init();
    </script>
</body>
</html>
```

### VBScript Example

```vbscript
<%@ Language="VBScript" %>
<%
Dim AxonLive
Set AxonLive = Server.CreateObject("G3AXONLIVE")
AxonLive.InitPage()

Dim sessionID
sessionID = Session.SessionID

Dim currentCount
currentCount = AxonLive.GetComponentProperty(sessionID, "main", "clicks")
If IsEmpty(currentCount) Or currentCount = "" Then
    currentCount = 0
Else
    currentCount = CLng(currentCount)
End If

' Handle Asynchronous Requests
If AxonLive.IsAsyncRequest Then
    Dim compID
    compID = AxonLive.EventComponentID
    
    Select Case compID
        Case "btnIncrement"
            currentCount = currentCount + 1
        Case "btnDecrement"
            currentCount = currentCount - 1
    End Select
    
    ' Persist the state
    AxonLive.SetComponentProperty sessionID, "main", "clicks", CStr(currentCount)
    
    ' Register the component patch
    Dim newHtml
    newHtml = "<div id=""counterDisplay"" class=""pill pill-primary"">Current Count: " & currentCount & "</div>"
    AxonLive.RegisterComponent "counterDisplay", newHtml
    
    ' End the response (halts further execution)
    AxonLive.EndAsyncResponse()
End If
%>
<!DOCTYPE html>
<html>
<head>
    <title>AxonLive VBScript Counter</title>
    <link rel="stylesheet" href="/css/axonasp.css">
</head>
<body>
    <div id="content" style="padding: 20px;">
        <h2>AxonLive Counter</h2>
        
        <!-- The component that will be dynamically replaced -->
        <div id="counterDisplay" class="pill pill-primary">Current Count: <%= currentCount %></div>
        
        <div style="margin-top: 15px;">
            <button id="btnIncrement" class="btn btn-primary" data-g3al-id="btnIncrement" data-g3al-event="click" data-g3al-event-name="onclick">
                Increment
            </button>
            <button id="btnDecrement" class="btn btn-danger" data-g3al-id="btnDecrement" data-g3al-event="click" data-g3al-event-name="onclick">
                Decrement
            </button>
        </div>
    </div>

    <script src="/axonlive/g3axonlive.js"></script>
    <script>
        G3AxonLive.init();
    </script>
</body>
</html>
```

## Routing and Mechanics

1. **Initial Load:** When the user navigates to the page, `AxonLive.IsAsyncRequest` evaluates to `False`. The server renders the full HTML document and sends it to the browser.
2. **Client Interaction:** The user clicks a button. The client script (`g3axonlive.js`) intercepts the click, packages the event details (Component ID, Event Name, Arguments), and sends a JSON `POST` request back to the server.
3. **Async Execution:** The server processes the script from the top down. This time, `AxonLive.IsAsyncRequest` is `True`. The logic inside the `If/Then` block runs.
4. **Patch Response:** `AxonLive.EndAsyncResponse()` collects all registered HTML patches and actions, formats them as JSON, clears the output buffer, and halts the script. The client receives the JSON and dynamically updates the DOM.

## Session Security Behavior

AxonLive async routing is bound to the authenticated `ASPSESSIONID` cookie on the server.

* The client bridge can call `G3AxonLive.init()` without passing a session ID.
* If a payload includes a `sessionId`, it must match the authenticated cookie session.
* Async page lookup and execution always use the authenticated session registration.