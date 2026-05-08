<%@ Language="VBScript" %>
<%
'
' G3AxonLive Full Capability Showcase
' Copyright (C) 2026 G3pix Ltda. All rights reserved.
'
' This example demonstrates all major features of the G3AxonLive framework:
' 1. Full HTML Replacement (RegisterComponent)
' 2. Granular property and attribute updates on real DOM properties
' 3. Style and CSS class management through re-rendered components
' 4. Server-triggered client actions (Timer, Redirect, Trigger)
' 5. Event Arguments
' 6. Persistent cross-request state
'

Function ReadLongOrDefault(rawValue, defaultValue)
    If IsEmpty(rawValue) Or CStr(rawValue) = "" Then
        ReadLongOrDefault = defaultValue
    Else
        ReadLongOrDefault = CLng(rawValue)
    End If
End Function

Function ReadBoolOrDefault(rawValue, defaultValue)
    If IsEmpty(rawValue) Or CStr(rawValue) = "" Then
        ReadBoolOrDefault = defaultValue
    Else
        ReadBoolOrDefault = (LCase(CStr(rawValue)) = "true" Or CStr(rawValue) = "1")
    End If
End Function

Function AppendLog(logText, message)
    Dim lineText

    lineText = "[" & Now & "] " & message
    If IsEmpty(logText) Or CStr(logText) = "" Then
        AppendLog = lineText
    Else
        AppendLog = CStr(logText) & vbCrLf & lineText
    End If
End Function

Function RenderStatusPill(clickCount, lastAction)
    Dim cssClass, styleAttr, labelText

    cssClass = "status-pill-big"
    styleAttr = " style=""background:#808080;color:#fff;"""
    labelText = "Ready to start..."

    If clickCount > 0 Then
        labelText = lastAction & " (Total: " & clickCount & ")"
        If clickCount Mod 2 = 0 Then
            cssClass = cssClass & " status-v"
            styleAttr = " style=""background:#3366cc;color:#fff;"""
        Else
            cssClass = cssClass & " status-x"
            styleAttr = " style=""background:#cc3300;color:#fff;"""
        End If
    ElseIf lastAction <> "None" Then
        labelText = lastAction
    End If

    RenderStatusPill = "<div id=""statusPill"" class=""" & cssClass & """" & styleAttr & ">" & Server.HTMLEncode(labelText) & "</div>"
End Function

Function RenderGhostButton(ghostPending)
    Dim extraAttrs, labelText

    extraAttrs = ""
    labelText = "Ghost (Delayed Trigger)"

    If ghostPending Then
        extraAttrs = " disabled=""disabled"" title=""Waiting for delayed trigger"""
        labelText = "Ghost (Waiting...)"
    End If

    RenderGhostButton = "<button id=""btnGhost"" class=""btn btn-info"" data-g3al-id=""btnGhost"" data-g3al-event=""click"" data-g3al-event-name=""onclick""" & extraAttrs & ">" & labelText & "</button>"
End Function

Function RenderServerLog(logText)
    RenderServerLog = "<textarea id=""txtLog"" readonly>" & Server.HTMLEncode(CStr(logText)) & "</textarea>"
End Function

Function RenderActionCard(lastAction)
    RenderActionCard = _
        "<div id=""cardAction"" class=""card"">" & _
        "<h3>Last Server Action</h3>" & _
        "<p class=""pill pill-primary"">" & Server.HTMLEncode(CStr(lastAction)) & "</p>" & _
        "<p>Timestamp: " & Now & "</p>" & _
        "</div>"
End Function

Function RenderStatusSummary(currentSessionID, clickCount)
    RenderStatusSummary = "<span id=""statusSummary"">Session: " & Server.HTMLEncode(currentSessionID) & " | Total Clicks: " & clickCount & "</span>"
End Function

Dim AxonLive
Set AxonLive = Server.CreateObject("G3AXONLIVE")
AxonLive.InitPage()

Dim sessionID : sessionID = Session.SessionID

Dim clickCount
clickCount = ReadLongOrDefault(AxonLive.GetComponentProperty(sessionID, "main", "clicks"), 0)

Dim lastAction
lastAction = AxonLive.GetComponentProperty(sessionID, "main", "lastAction")
If IsEmpty(lastAction) Then lastAction = "None"

Dim logText
logText = AxonLive.GetComponentProperty(sessionID, "txtLog", "val")
If IsEmpty(logText) Or CStr(logText) = "" Then logText = "Page Loaded: " & Now

Dim ghostPending
ghostPending = ReadBoolOrDefault(AxonLive.GetComponentProperty(sessionID, "main", "ghostPending"), False)

If AxonLive.IsAsyncRequest Then
    Dim compID : compID = AxonLive.EventComponentID
    Dim btnAction
    Set btnAction = AxonLive.GetComponent("btnAction")

    Select Case compID
        Case "btnAction"
            clickCount = clickCount + 1

            If ghostPending Then
                lastAction = "Ghost Trigger Completed"
                ghostPending = False
            Else
                lastAction = "Button Clicked"
            End If

            btnAction.SetAttribute "data-count", CStr(clickCount)
            btnAction.AddTitle "You have clicked " & clickCount & " times"
            logText = AppendLog(logText, "Action performed. Total: " & clickCount)

        Case "btnGhost"
            lastAction = "Ghost Trigger Scheduled"
            ghostPending = True
            logText = AppendLog(logText, "Ghost trigger scheduled.")
            AxonLive.SetTimer "btnAction", "onclick", 1000

        Case "btnRedirect"
            lastAction = "Redirecting"
            logText = AppendLog(logText, "Redirect requested.")
            AxonLive.Redirect "https://g3pix.com.br/axonasp"

        Case "btnTrigger"
            AxonLive.Trigger "btnAction", "onclick"
            lastAction = "Remote Trigger Fired"
            logText = AppendLog(logText, "Remote trigger fired.")

        Case "btnArgs"
            Dim stepVal
            stepVal = AxonLive.GetEventArg("step")
            If stepVal = "" Then stepVal = 1 Else stepVal = CLng(stepVal)

            clickCount = clickCount + stepVal
            lastAction = "Custom Step Addition: " & stepVal
            logText = AppendLog(logText, "Arguments received. Step=" & stepVal & ", total=" & clickCount)

        Case "timer1"
            clickCount = clickCount + 1
            lastAction = "Auto Timer Tick"
            logText = AppendLog(logText, "Timer tick. Total: " & clickCount)
            AxonLive.SetTimer "timer1", "ontimer", 5000
    End Select

    Call AxonLive.SetComponentProperty(sessionID, "main", "clicks", CStr(clickCount))
    Call AxonLive.SetComponentProperty(sessionID, "main", "lastAction", lastAction)
    Call AxonLive.SetComponentProperty(sessionID, "main", "ghostPending", CStr(ghostPending))
    Call AxonLive.SetComponentProperty(sessionID, "txtLog", "val", logText)

    AxonLive.RegisterComponent "statusPill", RenderStatusPill(clickCount, lastAction)
    AxonLive.RegisterComponent "txtLog", RenderServerLog(logText)
    AxonLive.RegisterComponent "cardAction", RenderActionCard(lastAction)
    AxonLive.RegisterComponent "btnGhost", RenderGhostButton(ghostPending)
    AxonLive.RegisterComponent "statusSummary", RenderStatusSummary(sessionID, clickCount)

    AxonLive.EndAsyncResponse()
End If
%>
<!DOCTYPE html>
<html>

    <head>
        <title>AxonLive Full Capability Demo</title>
        <link rel="stylesheet" href="/css/axonasp.css">
        <style>
            .demo-grid {
                display: grid;
                grid-template-columns: 1fr 1fr;
                gap: 20px;
            }

            #txtLog {
                width: 100%;
                height: 150px;
                font-family: monospace;
                font-size: 11px;
                margin-top: 10px;
            }

            .status-pill-big {
                font-size: 18px;
                padding: 15px;
                text-align: center;
                font-weight: bold;
                border-radius: 8px;
                color: #fff;
                background: #808080;
                transition: all 0.3s;
            }
        </style>
    </head>

    <body>

        <div id="header">
            <span
                style="color:#fff; font-family:Tahoma,Verdana,Arial,sans-serif; font-size:18px; font-weight:bold; line-height:60px; padding-left:18px;">
                G3AxonLive &mdash; Full Capability Demonstration
            </span>
        </div>

        <div id="main-container">
            <div id="content">

                <div class="info-banner">
                    This page demonstrates the supported AxonLive patterns for markup patches, client actions, event
                    arguments, and persisted state.
                </div>

                <div class="demo-grid">
                    <div class="col">
                        <div class="card">
                            <h3>Interactive Controls</h3>
                            <p>These buttons show the current procedural controller API.</p>

                            <div class="actions-row">
                                <button id="btnAction" class="btn btn-primary" data-g3al-id="btnAction"
                                    data-g3al-event="click" data-g3al-event-name="onclick">
                                    Primary Action
                                </button>

                                <button id="btnArgs" class="btn btn-secondary" data-g3al-id="btnArgs"
                                    data-g3al-event="click" data-g3al-event-name="onclick" data-g3al-arg-step="5">
                                    Add +5 (Args)
                                </button>
                            </div>

                            <div class="actions-row" style="margin-top:10px;">
                                <%=RenderGhostButton(ghostPending)%>

                                <button id="btnTrigger" class="btn btn-info" data-g3al-id="btnTrigger"
                                    data-g3al-event="click" data-g3al-event-name="onclick">
                                    Remote Trigger
                                </button>
                            </div>

                            <div class="actions-row" style="margin-top:10px;">
                                <button id="btnRedirect" class="btn btn-download" data-g3al-id="btnRedirect"
                                    data-g3al-event="click" data-g3al-event-name="onclick">
                                    External Redirect
                                </button>
                            </div>
                        </div>

                        <div class="card">
                            <h3>Server Log</h3>
                            <%=RenderServerLog(logText)%>
                        </div>
                    </div>

                    <div class="col">
                        <%=RenderStatusPill(clickCount, lastAction)%>
                        <%=RenderActionCard(lastAction)%>

                        <div class="card">
                            <h3>Framework Capabilities</h3>
                            <ul class="treeview">
                                <li class="folder">Markup Patching
                                    <ul class="submenu">
                                        <li class="file">RegisterComponent for text and markup surfaces</li>
                                        <li class="file">Persistent state restored on every async run</li>
                                    </ul>
                                </li>
                                <li class="folder">Client Actions
                                    <ul class="submenu">
                                        <li class="file">SetTimer for delayed events</li>
                                        <li class="file">Trigger for immediate follow-up requests</li>
                                        <li class="file">Redirect for navigation</li>
                                    </ul>
                                </li>
                                <li class="folder">Arguments and State
                                    <ul class="submenu">
                                        <li class="file">GetEventArg for button parameters</li>
                                        <li class="file">GetComponentProperty and SetComponentProperty</li>
                                    </ul>
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>

                <div id="timer1" data-g3al-id="timer1" data-g3al-event="timer" data-g3al-event-name="ontimer"
                    style="display:none"></div>

            </div>
        </div>

        <div id="status-bar">
            <%=RenderStatusSummary(Session.SessionID, clickCount)%>
        </div>

        <script src="/axonlive/g3axonlive.js"></script>
        <script>
            G3AxonLive.init();

            setTimeout(function () {
                G3AxonLive.trigger('timer1', 'ontimer');
            }, 5000);
        </script>

    </body>

</html>