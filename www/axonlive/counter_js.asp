<%@ Language="JScript" %>
<%
/*
 * G3AxonLive Counter Example (JScript - Procedural)
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimaraes - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 * Classic ASP Server-Side JScript example demonstrating the G3AxonLive
 * reactive component framework. Implements the same counter as counter.asp
 * using the procedural Go-controller pattern with JScript syntax.
 */

function renderCounterLabel(currentCount) {
    var className = 'counter-value';
    var styleAttr = '';

    if (currentCount < 0) {
        className += ' status-p';
        styleAttr = ' style="color:red;"';
    } else if (currentCount > 0) {
        styleAttr = ' style="color:green;"';
    }

    return '<span id="lblCounter" class="' + className + '"' + styleAttr + '>' + Server.HTMLEncode(String(currentCount)) + '</span>';
}

function renderResetButton(isDisabled) {
    var extraAttrs = '';
    if (isDisabled) {
        extraAttrs = ' disabled="disabled" title="Counter is already zero"';
    }

    return '<button id="btnReset" class="btn btn-danger" data-g3al-id="btnReset" data-g3al-event="click" data-g3al-event-name="onclick"' + extraAttrs + '>Reset</button>';
}

var AxonLive = Server.CreateObject("G3AXONLIVE");
AxonLive.InitPage();

var sessionID = Session.SessionID;
var rawCount = AxonLive.GetComponentProperty(sessionID, "counter", "count");
var count = 0;
if (rawCount !== "" && rawCount !== null && !isNaN(parseInt(rawCount, 10))) {
    count = parseInt(rawCount, 10);
}

if (AxonLive.IsAsyncRequest) {
    var compID  = AxonLive.EventComponentID;
    var evtName = AxonLive.EventName;

    if (compID === "btnIncrement" && evtName === "onclick") {
        count = count + 1;
    } else if (compID === "btnDecrement" && evtName === "onclick") {
        count = count - 1;
    } else if (compID === "btnReset" && evtName === "onclick") {
        count = 0;
    }

    AxonLive.SetComponentProperty(sessionID, "counter", "count", String(count));
    AxonLive.RegisterComponent("lblCounter", renderCounterLabel(count));
    AxonLive.RegisterComponent("btnReset", renderResetButton(count === 0));
    AxonLive.EndAsyncResponse();
}
%>
<!DOCTYPE html>
<html lang="en">

    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>G3AxonLive Counter (JScript) - AxonASP</title>
        <link rel="stylesheet" href="/css/axonasp.css">
        <style>
            .counter-panel {
                text-align: center;
                padding: 32px 24px;
                max-width: 420px;
                margin: 32px auto;
            }

            .counter-value {
                display: block;
                font-size: 64px;
                font-weight: bold;
                font-family: Tahoma, Verdana, Arial, sans-serif;
                color: var(--win-blue-dark);
                margin: 16px 0 24px;
                line-height: 1;
            }

            .counter-actions {
                display: flex;
                gap: 10px;
                justify-content: center;
                flex-wrap: wrap;
            }
        </style>
    </head>

    <body>

        <div id="header">
            <span
                style="color:#fff; font-family:Tahoma,Verdana,Arial,sans-serif; font-size:18px; font-weight:bold; line-height:60px; padding-left:18px;">
                G3AxonLive &mdash; Reactive Counter (JScript)
            </span>
        </div>

        <div id="main-container">
            <div id="content">

                <div class="card counter-panel">
                    <h2>Live Counter</h2>
                    <p>Click the buttons to update the counter without a full page reload.<br>
                        All logic runs server-side &mdash; only the changed HTML is returned.</p>

                    <%
    Response.Write(renderCounterLabel(count));
    %>

                    <div class="counter-actions">
                        <button id="btnDecrement" class="btn btn-secondary" data-g3al-id="btnDecrement"
                            data-g3al-event="click" data-g3al-event-name="onclick">
                            &minus; Decrement
                        </button>
                        <%
        Response.Write(renderResetButton(count === 0));
        %>
                        <button id="btnIncrement" class="btn btn-primary" data-g3al-id="btnIncrement"
                            data-g3al-event="click" data-g3al-event-name="onclick">
                            + Increment
                        </button>
                    </div>
                </div>

                <div class="card" style="max-width:420px; margin:0 auto 24px;">
                    <h3>How it works (JScript)</h3>
                    <ul>
                        <li>On page load, <code>AxonLive.InitPage()</code> registers the session.</li>
                        <li>When a button is clicked, the JS engine POSTs to <code>/g3al/</code>.</li>
                        <li>The server re-runs this page, detects <code>AxonLive.IsAsyncRequest</code>,
                            applies the counter mutation, and calls <code>AxonLive.EndAsyncResponse()</code>.</li>
                        <li>The browser replaces the <code>lblCounter</code> and <code>btnReset</code> markup via
                            <code>outerHTML</code>.
                        </li>
                        <li>The server-side logic is identical to the VBScript version &mdash; same Go API,
                            different scripting language.</li>
                    </ul>
                </div>

            </div>
        </div>

        <div id="status-bar">AxonASP &mdash; G3AxonLive Counter Example (JScript)</div>

        <script src="/axonlive/g3axonlive.js"></script>
        <script>
            G3AxonLive.init();
        </script>
    </body>

</html>