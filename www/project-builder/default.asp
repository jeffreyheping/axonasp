<%
Option Explicit
Response.ContentType = "text/html; charset=utf-8"
%>
<!DOCTYPE html>
<html lang="en" class="project-builder-root">

    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>ASP Project Builder - AxonASP Code Generation Assistant</title>
        <link rel="stylesheet" href="../css/axonasp.css" />
    </head>

    <body class="project-builder-page">
        <div id="header">
            <div class="logo">
                <img src="data:image/svg+xml; charset=utf-8;;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiIHN0YW5kYWxvbmU9Im5vIj8+CjwhLS0gQ3JlYXRlZCB3aXRoIElua3NjYXBlIChodHRwOi8vd3d3Lmlua3NjYXBlLm9yZy8pIC0tPgoKPHN2ZwogICB2ZXJzaW9uPSIxLjEiCiAgIGlkPSJzdmcxIgogICB3aWR0aD0iNjc3LjM4Mzg1MDEiCiAgIGhlaWdodD0iNjc3LjI3NzE2MDYiCiAgIHZpZXdCb3g9IjAgMCA2NzcuMzgzODUwMSA2NzcuMjc3MTYwNiIKICAgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIgogICB4bWxuczpzdmc9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KICA8ZGVmcwogICAgIGlkPSJkZWZzMSIgLz4KICA8ZwogICAgIGlkPSJnMSIKICAgICB0cmFuc2Zvcm09InRyYW5zbGF0ZSgtMi42NDE0MjQ2NTYsLTIuNjk0Nzg2MDcyKSI+CiAgICA8ZwogICAgICAgaWQ9ImcxMyIKICAgICAgIHRyYW5zZm9ybT0ibWF0cml4KDIuMjc2NzU2MDQ0LDAsMCwyLjI3Njc1NjA0NCwtNjM0LjQ5NDUyNDQsLTcxOC40MjMyMjI4KSI+CiAgICAgIDxwYXRoCiAgICAgICAgIGQ9Ik0gNTEyLjgwMDc4MTMsNDAwLjk5MjE4NzUgNDQ4LjM3NSw0NjUuNDE3OTY4NyBoIDY0LjQ4ODI4MTMgNjQuMzYzMjgxMiB6IG0gLTY0LjQzMTY0MDcsNjQuNDMxNjQwNiAtMC4wMDM5MSwwLjAwMzkxIDY0LjQzNTU0NjksNjIuNzgzMjAzMSAwLjAzMzIwMywtMC4wMzEyNSAtMzIuMjIwNzAzMiwtMzEuMzY1MjM0NCB6IgogICAgICAgICBzdHlsZT0iZm9udC13ZWlnaHQ6Ym9sZDtmb250LXNpemU6NjE5LjI4NHB4O2ZvbnQtZmFtaWx5OlRhaG9tYTstaW5rc2NhcGUtZm9udC1zcGVjaWZpY2F0aW9uOidUYWhvbWEgQm9sZCc7bGV0dGVyLXNwYWNpbmc6MHB4O2Jhc2VsaW5lLXNoaWZ0OmJhc2VsaW5lO2Rpc3BsYXk6aW5saW5lO292ZXJmbG93OnZpc2libGU7dmVjdG9yLWVmZmVjdDpub25lO2ZpbGw6IzA0OWM0ZDtlbmFibGUtYmFja2dyb3VuZDphY2N1bXVsYXRlO3N0b3AtY29sb3I6IzAwMDAwMDtzdG9wLW9wYWNpdHk6MSIKICAgICAgICAgaWQ9InBhdGg5IiAvPgogICAgICA8cGF0aAogICAgICAgICBkPSJtIDM0NC4yNzkyOTY5LDQwMC45OTIxODc1IC02NC4zODQ3NjU3LDY0LjM4NDc2NTYgaCA2NC40NTUwNzgyIDY0LjMxMjUgeiBtIC02NC40MDYyNSw2NC40MDYyNSAtMC4wMjkyOTcsMC4wMjkyOTcgNjAuNDg0Mzc1LDU4LjkzMzU5MzcgLTI4LjIyODUxNTYsLTI3LjUzMTI1IHoiCiAgICAgICAgIHN0eWxlPSJmb250LXdlaWdodDpib2xkO2ZvbnQtc2l6ZTo2MTkuMjg0cHg7Zm9udC1mYW1pbHk6VGFob21hOy1pbmtzY2FwZS1mb250LXNwZWNpZmljYXRpb246J1RhaG9tYSBCb2xkJztsZXR0ZXItc3BhY2luZzowcHg7YmFzZWxpbmUtc2hpZnQ6YmFzZWxpbmU7ZGlzcGxheTppbmxpbmU7b3ZlcmZsb3c6dmlzaWJsZTt2ZWN0b3ItZWZmZWN0Om5vbmU7ZmlsbDojNjA5NDFhO2VuYWJsZS1iYWNrZ3JvdW5kOmFjY3VtdWxhdGU7c3RvcC1jb2xvcjojMDAwMDAwO3N0b3Atb3BhY2l0eToxIgogICAgICAgICBpZD0icGF0aDExIiAvPgogICAgICA8cGF0aAogICAgICAgICBkPSJtIDQyOC41MzkwNjI1LDQ4NS4yNTM5MDYyIC02Mi42NTYyNSw2NC4zMDI3MzQ0IGggNjIuNjQ4NDM3NSA2Mi42NjQwNjI1IHogbSA2Mi43NjU2MjUsNjQuNDE2MDE1NyAtMTkuNTc0MjE4OCwyMC4xMjUgMTkuNTkxNzk2OSwtMjAuMTA3NDIxOSB6IG0gLTEyNS41MzkwNjI1LDAuMDA3ODEgLTAuMDA5NzcsMC4wMDk3NyAxMC44NjEzMjgxLDExLjE0NjQ4NDQgeiIKICAgICAgICAgc3R5bGU9ImZvbnQtd2VpZ2h0OmJvbGQ7Zm9udC1zaXplOjYxOS4yODRweDtmb250LWZhbWlseTpUYWhvbWE7LWlua3NjYXBlLWZvbnQtc3BlY2lmaWNhdGlvbjonVGFob21hIEJvbGQnO2xldHRlci1zcGFjaW5nOjBweDtiYXNlbGluZS1zaGlmdDpiYXNlbGluZTtkaXNwbGF5OmlubGluZTtvdmVyZmxvdzp2aXNpYmxlO3ZlY3Rvci1lZmZlY3Q6bm9uZTtmaWxsOiNlZGMwMGY7ZW5hYmxlLWJhY2tncm91bmQ6YWNjdW11bGF0ZTtzdG9wLWNvbG9yOiMwMDAwMDA7c3RvcC1vcGFjaXR5OjEiCiAgICAgICAgIGlkPSJwYXRoMTMiIC8+CiAgICAgIDxwYXRoCiAgICAgICAgIGQ9Im0gNDI4LjUzOTA2MjUsMzE2LjczMDQ2ODcgLTYyLjc4MzIwMzEsNjQuNDM1NTQ2OSAwLjA1MjczNCwwLjA1MjczNCBoIDYyLjc1OTc2NTcgNjIuNzAxMTcxOCBsIDAuMDUyNzM0LC0wLjA1MjczNCB6IG0gLTI2Ljc5Njg3NSwxMDAuNDIxODc1IDI2Ljc5Njg3NSwyNi43OTY4NzUgMy44NjEzMjgxLC0zLjg2MTMyODEgLTMuODMyMDMxMiwzLjgyMjI2NTYgeiIKICAgICAgICAgc3R5bGU9ImZvbnQtd2VpZ2h0OmJvbGQ7Zm9udC1zaXplOjYxOS4yODRweDtmb250LWZhbWlseTpUYWhvbWE7LWlua3NjYXBlLWZvbnQtc3BlY2lmaWNhdGlvbjonVGFob21hIEJvbGQnO2xldHRlci1zcGFjaW5nOjBweDtiYXNlbGluZS1zaGlmdDpiYXNlbGluZTtkaXNwbGF5OmlubGluZTtvdmVyZmxvdzp2aXNpYmxlO3ZlY3Rvci1lZmZlY3Q6bm9uZTtmaWxsOiMwMDRhYWQ7ZW5hYmxlLWJhY2tncm91bmQ6YWNjdW11bGF0ZTtzdG9wLWNvbG9yOiMwMDAwMDA7c3RvcC1vcGFjaXR5OjEiCiAgICAgICAgIGlkPSJwYXRoNyIgLz4KICAgICAgPHBhdGgKICAgICAgICAgZD0ibSA0NDguMzYxMzI4MSw0NjUuNDE3OTY4NyAzMi4yNTE5NTMxLDMxLjM5NjQ4NDQgMzIuMjUwMDAwMSwzMS4zOTQ1MzEzIDMyLjI1MTk1MzEsLTMxLjM5NDUzMTMgMzIuMjUsLTMxLjM5NjQ4NDQgaCAtNjQuNTAxOTUzMSB6IgogICAgICAgICBzdHlsZT0iYmFzZWxpbmUtc2hpZnQ6YmFzZWxpbmU7ZGlzcGxheTppbmxpbmU7b3ZlcmZsb3c6dmlzaWJsZTt2ZWN0b3ItZWZmZWN0Om5vbmU7ZmlsbDojMDA0YWFkO2VuYWJsZS1iYWNrZ3JvdW5kOmFjY3VtdWxhdGU7c3RvcC1jb2xvcjojMDAwMDAwO3N0b3Atb3BhY2l0eToxIgogICAgICAgICBpZD0icGF0aDgiIC8+CiAgICAgIDxwYXRoCiAgICAgICAgIGQ9Im0gMjc5Ljg0OTYwOTQsNDY1LjM3Njk1MzEgMzIuMjUsMzEuNDUzMTI1IDMyLjI1LDMxLjQ1MzEyNSAzMi4yNDgwNDY4LC0zMS40NTMxMjUgMzIuMjUsLTMxLjQ1MzEyNSBoIC02NC40OTgwNDY4IHoiCiAgICAgICAgIHN0eWxlPSJiYXNlbGluZS1zaGlmdDpiYXNlbGluZTtkaXNwbGF5OmlubGluZTtvdmVyZmxvdzp2aXNpYmxlO3ZlY3Rvci1lZmZlY3Q6bm9uZTtmaWxsOiNlZGMwMGY7ZW5hYmxlLWJhY2tncm91bmQ6YWNjdW11bGF0ZTtzdG9wLWNvbG9yOiMwMDAwMDA7c3RvcC1vcGFjaXR5OjEiCiAgICAgICAgIGlkPSJwYXRoMTAiIC8+CiAgICAgIDxwYXRoCiAgICAgICAgIGQ9Im0gMzY1LjcxNjc5NjksMzgxLjIxODc1IDMxLjQyNTc4MTIsMzEuMzQ1NzAzMSAzMS40MjU3ODEzLDMxLjM0NTcwMzEgMzEuNDI1NzgxMiwtMzEuMzQ1NzAzMSAzMS40MjM4MjgxLC0zMS4zNDU3MDMxIGggLTYyLjg0OTYwOTMgeiIKICAgICAgICAgc3R5bGU9ImJhc2VsaW5lLXNoaWZ0OmJhc2VsaW5lO2Rpc3BsYXk6aW5saW5lO292ZXJmbG93OnZpc2libGU7dmVjdG9yLWVmZmVjdDpub25lO2ZpbGw6IzYwOTQxYTtlbmFibGUtYmFja2dyb3VuZDphY2N1bXVsYXRlO3N0b3AtY29sb3I6IzAwMDAwMDtzdG9wLW9wYWNpdHk6MSIKICAgICAgICAgaWQ9InBhdGg1IiAvPgogICAgICA8cGF0aAogICAgICAgICBkPSJtIDM2NS42NDY0ODQ0LDU0OS41NTY2NDA2IDMxLjQ0MzM1OTMsMzIuMzI0MjE4OCAzMS40NDE0MDYzLDMyLjMyNDIxODcgMzEuNDQzMzU5NCwtMzIuMzI0MjE4NyAzMS40NDE0MDYyLC0zMi4zMjQyMTg4IEggNDI4LjUzMTI1IFoiCiAgICAgICAgIHN0eWxlPSJiYXNlbGluZS1zaGlmdDpiYXNlbGluZTtkaXNwbGF5OmlubGluZTtvdmVyZmxvdzp2aXNpYmxlO3ZlY3Rvci1lZmZlY3Q6bm9uZTtmaWxsOiMwNDljNGQ7ZW5hYmxlLWJhY2tncm91bmQ6YWNjdW11bGF0ZTtzdG9wLWNvbG9yOiMwMDAwMDA7c3RvcC1vcGFjaXR5OjEiCiAgICAgICAgIGlkPSJwYXRoMTIiIC8+CiAgICA8L2c+CiAgPC9nPgo8L3N2Zz4K"
                    alt="AxonASP" width="43">
            </div>
            <h1>ASP Code Generator</h1>
        </div>

        <div id="main-container">
            <div class="sidebar" id="sidebar">
                <div class="section-title">User Guide</div>
                <ul>
                    <li><a href="#section-app">Application Details</a></li>
                    <li><a href="#section-arch">Architecture</a></li>
                    <li><a href="#section-features">Features</a></li>
                    <li><a href="#section-lang">Localization</a></li>
                    <li><a href="#section-extra">Advanced Options</a></li>
                    <li><a href="#section-output">Generated Output</a></li>
                </ul>

                <div class="section-title">About ASP Builder</div>
                <p class="pb-sidebar-text">
                    This tool generates structured prompts for AI coding agents
                    building applications with AxonASP and ASP Classic. Provide
                    your requirements and receive a comprehensive markdown
                    document ready for your chosen agent.
                </p>

                <div class="section-title">AxonASP Resources</div>
                <ul>
                    <li><a href="/">Home</a></li>
                    <li><a href="/manual/">Documentation</a></li>
                </ul>
            </div>

            <div id="content">
                <h1>Code Generation Assistant</h1>
                <p class="intro-text">
                    Build structured AI prompts for developing web applications
                    using AxonASP. Describe your vision, choose your
                    preferences, and generate a comprehensive guideline
                    document. Perfect for collaboration with development agents
                    and teams.
                </p>

                <!-- Application Details Section -->
                <div id="section-app">
                    <h2>Core Application Definition</h2>
                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Project Name
                        </h3>
                        <p>Assign a name for your application.</p>
                        <div class="form-input-area">
                            <input type="text" id="appname" placeholder="Example: My Blog" />
                        </div>
                    </div>
                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Requirements Statement
                        </h3>
                        <p>
                            Describe the application concept, target users, and
                            primary functionality.
                        </p>
                        <div class="form-input-area">
                            <textarea id="description"
                                placeholder="Example: A time-tracking platform for freelancers to log billable hours, organize projects by client, generate invoices, track expenses per project, and create monthly billing reports."></textarea>
                        </div>
                        <p class="note">
                            Be specific about core features, workflows, and
                            business logic.
                        </p>
                    </div>
                </div>

                <!-- Architecture Section -->
                <div id="section-arch">
                    <h2>Architecture & Technology Stack</h2>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Server-Side Language
                        </h3>
                        <p>
                            Select the primary Classic ASP language for generated server-side code.
                        </p>
                        <div class="radio-group">
                            <div class="radio-item">
                                <input type="radio" name="serverlang" value="VBScript" id="lang-vbscript" checked />
                                <label for="lang-vbscript">VBScript</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="serverlang" value="JavaScript (JScript ASP)"
                                    id="lang-jscript" />
                                <label for="lang-jscript">JavaScript (JScript ASP)</label>
                            </div>
                        </div>
                    </div>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Design Pattern
                        </h3>
                        <p>
                            Selection determines code organization and
                            development workflow.
                        </p>
                        <div class="radio-group">
                            <div class="radio-item">
                                <input type="radio" name="style" value="mvc" id="style-mvc" checked />
                                <label for="style-mvc">Model-View-Controller</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="style" value="mvvm" id="style-mvvm" />
                                <label for="style-mvvm">Model-View-ViewModel</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="style" value="mixed" id="style-mixed" />
                                <label for="style-mixed">Inline Code</label>
                            </div>
                        </div>
                        <p class="note">
                            MVC recommended for larger projects. Inline suitable
                            for simple pages.
                        </p>
                    </div>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Data Persistence
                        </h3>
                        <p>
                            Backend database technology for application data
                            storage.
                        </p>
                        <div class="radio-group">
                            <div class="radio-item">
                                <input type="radio" name="database" value="sqlite" id="db-sqlite" checked />
                                <label for="db-sqlite">SQLite</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="database" value="mysql" id="db-mysql" />
                                <label for="db-mysql">MySQL</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="database" value="postgresql" id="db-psql" />
                                <label for="db-psql">PostgreSQL</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="database" value="mssql" id="db-mssql" />
                                <label for="db-mssql">MS SQL Server</label>
                            </div>
                        </div>
                    </div>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            User Interface Framework
                        </h3>
                        <p>
                            CSS framework for responsive design and component
                            styling.
                        </p>
                        <div class="radio-group">
                            <div class="radio-item">
                                <input type="radio" name="ui" value="axonasp" id="ui-axonasp" />
                                <label for="ui-axonasp">AxonASP Native</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="ui" value="bootstrap" id="ui-bootstrap" />
                                <label for="ui-bootstrap">Bootstrap 5</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="ui" value="tailwind" id="ui-tailwind" checked />
                                <label for="ui-tailwind">Tailwind CSS</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="ui" value="None" id="ui-none" />
                                <label for="ui-none">None</label>
                            </div>
                        </div>
                        <div id="axonasp-native-hint" class="pb-native-hint">
                            <strong>AxonASP Native Style Directives:</strong><br />
                            Uses the application local
                            stylesheet <code>axonasp.css</code> when available,
                            or download and use
                            <code>https://g3pix.com.br/axonasp/css/axonasp.css</code>.
                        </div>
                    </div>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            JavaScript Framework
                        </h3>
                        <p>
                            Client-side JavaScript framework for interactivity
                            and dynamic behavior.
                        </p>
                        <div class="radio-group">
                            <div class="radio-item">
                                <input type="radio" name="jsframework" value="Vanilla JavaScript" id="js-vanilla"
                                    checked />
                                <label for="js-vanilla">Vanilla JavaScript</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="jsframework" value="jQuery" id="js-jquery" />
                                <label for="js-jquery">jQuery</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="jsframework" value="React" id="js-react" />
                                <label for="js-react">React</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="jsframework" value="Vue" id="js-vue" />
                                <label for="js-vue">Vue</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="jsframework" value="Next.js" id="js-nextjs" />
                                <label for="js-nextjs">Next.js</label>
                            </div>
                            <div class="radio-item">
                                <input type="radio" name="jsframework" value="Angular" id="js-angular" />
                                <label for="js-angular">Angular</label>
                            </div>
                        </div>
                    </div>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Visual Appearance
                        </h3>
                        <div class="two-col">
                            <div>
                                <p class="pb-option-title">
                                    Color Palette
                                </p>
                                <div class="radio-group pb-options-stack">
                                    <div class="radio-item pb-radio-compact">
                                        <input type="radio" name="color" value="light" id="color-light" checked />
                                        <label for="color-light">Light Theme</label>
                                    </div>
                                    <div class="radio-item pb-radio-compact">
                                        <input type="radio" name="color" value="dark" id="color-dark" />
                                        <label for="color-dark">Dark Theme</label>
                                    </div>
                                    <div class="radio-item pb-radio-compact">
                                        <input type="radio" name="color" value="auto" id="color-auto" />
                                        <label for="color-auto">Auto-Detect</label>
                                    </div>
                                </div>
                            </div>
                            <div>
                                <p class="pb-option-title">
                                    Interactive Elements
                                </p>
                                <div class="radio-group pb-options-stack">
                                    <div class="radio-item pb-radio-compact">
                                        <input type="radio" name="emoji" value="enabled" id="emoji-enabled" checked />
                                        <label for="emoji-enabled">Include Icons</label>
                                    </div>
                                    <div class="radio-item pb-radio-compact">
                                        <input type="radio" name="emoji" value="minimal" id="emoji-minimal" />
                                        <label for="emoji-minimal">Minimal Icons</label>
                                    </div>
                                    <div class="radio-item pb-radio-compact">
                                        <input type="radio" name="emoji" value="none" id="emoji-none" />
                                        <label for="emoji-none">Text Only</label>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Features Section -->
                <div id="section-features">
                    <h2>Functional Components</h2>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Select Enabled Capabilities
                        </h3>
                        <p>
                            Choose which features to include in your generated
                            specification.
                        </p>
                        <div class="checkbox-group">
                            <div class="check-item">
                                <input type="checkbox" id="feat-auth" checked />
                                <label for="feat-auth">User Authentication</label>
                            </div>
                            <div class="check-item">
                                <input type="checkbox" id="feat-sample" />
                                <label for="feat-sample">Sample Data Initialization</label>
                            </div>
                            <div class="check-item">
                                <input type="checkbox" id="feat-crud" checked />
                                <label for="feat-crud">Data Management (CRUD)</label>
                            </div>
                            <br />
                            <div class="check-item">
                                <input type="checkbox" id="feat-search" />
                                <label for="feat-search">Search Capability</label>
                            </div>
                            <div class="check-item">
                                <input type="checkbox" id="feat-upload" />
                                <label for="feat-upload">File Handling</label>
                            </div>
                            <div class="check-item">
                                <input type="checkbox" id="feat-dash" />
                                <label for="feat-dash">Dashboard Page</label>
                            </div>
                            <br />
                            <div class="check-item">
                                <input type="checkbox" id="feat-pdf" />
                                <label for="feat-pdf">PDF Export</label>
                            </div>
                            <div class="check-item">
                                <input type="checkbox" id="feat-email" />
                                <label for="feat-email">Email Integration</label>
                            </div>
                            <div class="check-item">
                                <input type="checkbox" id="feat-api" />
                                <label for="feat-api">REST API Endpoints</label>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Localization Section -->
                <div id="section-lang">
                    <h2>Localization & Content Language</h2>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Application Language
                        </h3>
                        <p>
                            Primary language for all user-facing content and
                            interface text.
                        </p>
                        <div class="form-input-area">
                            <select id="lang-select">
                                <option value="English">English</option>
                                <option value="Brazilian Portuguese">
                                    Brazilian Portuguese
                                </option>
                                <option value="Portuguese">Portuguese</option>
                                <option value="Spanish">Spanish</option>
                                <option value="French">French</option>
                                <option value="German">German</option>
                                <option value="Italian">Italian</option>
                                <option value="Russian">Russian</option>
                                <option value="Japanese">Japanese</option>
                                <option value="Chinese">
                                    Chinese (Simplified)
                                </option>
                                <option value="Chinese (Traditional)">
                                    Chinese (Traditional)
                                </option>
                                <option value="Korean">Korean</option>
                                <option value="Arabic">Arabic</option>
                                <option value="Hindi">Hindi</option>
                                <option value="Vietnamese">Vietnamese</option>
                                <option value="Thai">Thai</option>
                                <option value="Indonesian">Indonesian</option>
                            </select>
                        </div>
                    </div>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Multi-Language Support
                        </h3>
                        <p>
                            Enable translation framework for content in
                            additional languages.
                        </p>
                        <div class="check-item">
                            <input type="checkbox" id="multilingual" />
                            <label for="multilingual">Enable Multi-Language Interface</label>
                        </div>
                        <div id="multi-langs-list" class="pb-multilang-list">
                            <p class="pb-multilang-text">
                                Additional language support (in addition to
                                selected primary):
                            </p>
                            <div class="checkbox-group" id="multi-lang-checks"></div>
                        </div>
                        <p class="note">
                            Framework will include language switching and
                            translation management.
                        </p>
                    </div>
                </div>

                <!-- Advanced Options Section -->
                <div id="section-extra">
                    <h2>Advanced Configuration</h2>

                    <div class="form-section">
                        <h3 class="pb-subtitle">
                            Custom Requirements
                            <span class="tag">(optional)</span>
                        </h3>
                        <p>
                            Specify additional constraints, architectural
                            decisions, or special implementation details.
                        </p>
                        <div class="form-input-area">
                            <textarea id="extra"
                                placeholder="Example: Implement role-based access control with Admin/User roles. Use stored procedures for complex queries. Include activity logging for all data modifications."></textarea>
                        </div>
                    </div>
                </div>

                <!-- Output Section -->
                <div id="section-output">
                    <h2>Generated Documentation</h2>
                    <div class="button-group">
                        <button class="btn btn-primary" onclick="generatePrompt()">
                            Prepare Markdown Document
                        </button>
                        <button class="btn" onclick="copyToClipboard()">
                            Copy Instructions
                        </button>
                    </div>

                    <div id="output-area">
                        <h3>Your Generated Agent Prompt</h3>
                        <p class="pb-output-note">
                            Copy this markdown and paste it into your AI coding
                            agent.
                        </p>
                        <textarea id="output-textarea" readonly></textarea>
                    </div>
                </div>
            </div>
        </div>

        <div class="copy-toast" id="copy-toast">
            Copied to clipboard successfully!
        </div>
        <script>
            // Helper functions
            function val(id) {
                return (document.getElementById(id) || {}).value || "";
            }
            function chk(id) {
                return (document.getElementById(id) || {}).checked || false;
            }
            function getRadio(name) {
                var el = document.querySelector(
                    'input[name="' + name + '"]:checked'
                );
                return el ? el.value : "";
            }

            // Initialize multi-language support
            function initMultiLang() {
                var langs = [
                    "Brazilian Portuguese",
                    "Portuguese",
                    "Spanish",
                    "French",
                    "German",
                    "Italian",
                    "Russian",
                    "Japanese",
                    "Chinese (Traditional)",
                    "Korean",
                    "Arabic",
                    "Hindi",
                    "Vietnamese",
                    "Thai",
                    "Indonesian",
                ];

                var container = document.getElementById("multi-lang-checks");
                if (container) {
                    container.innerHTML = "";
                    langs.forEach(function (lang) {
                        var id =
                            "ml-" +
                            lang.replace(/[^a-zA-Z]/g, "").toLowerCase();
                        var div = document.createElement("div");
                        div.className = "check-item";
                        div.innerHTML =
                            '<input type="checkbox" id="' +
                            id +
                            '" data-lang="' +
                            lang +
                            '">' +
                            '<label for="' +
                            id +
                            '">' +
                            lang +
                            "</label>";
                        container.appendChild(div);
                    });
                }
            }

            // Multi-language toggle
            if (document.getElementById("multilingual")) {
                document
                    .getElementById("multilingual")
                    .addEventListener("change", function () {
                        var list = document.getElementById("multi-langs-list");
                        if (list) {
                            list.style.display = this.checked
                                ? "block"
                                : "none";
                        }
                    });
            }

            // Show/hide AxonASP Native hint
            document
                .querySelectorAll('input[name="ui"]')
                .forEach(function (radio) {
                    radio.addEventListener("change", function () {
                        var hint = document.getElementById(
                            "axonasp-native-hint"
                        );
                        if (hint) {
                            hint.style.display =
                                this.value === "axonasp" ? "block" : "none";
                        }
                    });
                });

            // Radio styling
            document
                .querySelectorAll(".radio-item input, .check-item input")
                .forEach(function (input) {
                    input.addEventListener("change", function () {
                        var parent =
                            this.closest(".radio-item") ||
                            this.closest(".check-item");
                        if (parent) {
                            if (this.type === "radio") {
                                var group = this.getAttribute("name");
                                document
                                    .querySelectorAll(
                                        'input[name="' + group + '"]'
                                    )
                                    .forEach(function (el) {
                                        var p = el.closest(".radio-item");
                                        if (p) p.classList.remove("selected");
                                    });
                            }
                            parent.classList.toggle("selected", this.checked);
                        }
                    });

                    // Initial styling
                    var parent =
                        input.closest(".radio-item") ||
                        input.closest(".check-item");
                    if (parent && input.checked) {
                        parent.classList.add("selected");
                    }
                });

            function generatePrompt() {
                var desc = val("description").trim();
                if (!desc) {
                    alert("Please provide an application description.");
                    document.getElementById("description").focus();
                    return;
                }

                var features = [];
                if (chk("feat-auth"))
                    features.push(
                        "Secure user authentication with session management"
                    );
                if (chk("feat-sample"))
                    features.push("Pre-populated sample data for testing");
                if (chk("feat-crud"))
                    features.push(
                        "Complete data management (create, read, update, delete operations)-CRUD"
                    );
                if (chk("feat-search"))
                    features.push(
                        "Data search and advanced filtering capabilities"
                    );
                if (chk("feat-upload"))
                    features.push("File upload and attachment handling");
                if (chk("feat-dash"))
                    features.push(
                        "Administrative dashboard with key statistics"
                    );
                if (chk("feat-pdf"))
                    features.push("PDF document generation and export");
                if (chk("feat-email"))
                    features.push("Email integration and notification system");
                if (chk("feat-api"))
                    features.push(
                        "RESTful API endpoints for external integration"
                    );

                var primaryLang = val("lang-select");
                var isMulti = chk("multilingual");
                var supportedLangs = [primaryLang];

                if (isMulti) {
                    document
                        .querySelectorAll("#multi-lang-checks input:checked")
                        .forEach(function (el) {
                            var lang = el.getAttribute("data-lang");
                            if (lang && supportedLangs.indexOf(lang) === -1) {
                                supportedLangs.push(lang);
                            }
                        });
                }

                var style = getRadio("style");
                var db = getRadio("database");
                var ui = getRadio("ui");
                var serverLang = getRadio("serverlang") || "VBScript";
                var jsframework = getRadio("jsframework");
                var color = getRadio("color");
                var emoji = getRadio("emoji");

                var md =
                    "# " +
                    (val("appname") || "Web Application") +
                    " - AxonASP Development Specification\n\n";

                md += "## Project Overview\n\n";
                md += desc + "\n\n";

                md += "## Technical Requirements\n\n";

                md += "### Platform & Language\n";
                md += "- **Language:** Classic ASP (" + serverLang + ")\n";
                md += "- **Runtime:** AxonASP Virtual Machine\n\n";

                md += "### Architecture\n";
                md += "- **Pattern:** " + style.toUpperCase() + "\n";
                md +=
                    "- **Database:** " +
                    db.charAt(0).toUpperCase() +
                    db.slice(1) +
                    "\n";
                md +=
                    "- **UI Framework:** " +
                    ui.charAt(0).toUpperCase() +
                    ui.slice(1) +
                    "\n";
                if (ui === "axonasp") {
                    md +=
                        "  - _Stylesheet Source:_ Use `axonasp.css` from the application directory first. If unavailable, download and use https://g3pix.com.br/axonasp/css/axonasp.css\n";
                }
                md += "- **JavaScript Framework:** " + jsframework + "\n";
                md +=
                    "- **Theme:** " +
                    color.charAt(0).toUpperCase() +
                    color.slice(1) +
                    "\n";
                md +=
                    "- **Icons:** " +
                    emoji.charAt(0).toUpperCase() +
                    emoji.slice(1) +
                    "\n\n";

                md += "### Supported Languages\n";
                md += "- Primary: " + primaryLang + "\n";
                if (isMulti && supportedLangs.length > 1) {
                    md +=
                        "- Additional: " +
                        supportedLangs.slice(1).join(", ") +
                        "\n";
                }
                md +=
                    "- All UI text must be translatable and centralized. English is always supported.\n\n";

                if (features.length > 0) {
                    md += "### Required Features\n\n";
                    features.forEach(function (feat) {
                        md += "- " + feat + "\n";
                    });
                    md += "\n";
                }

                if (val("extra").trim()) {
                    md += "### Custom Requirements\n\n";
                    md += val("extra").trim() + "\n\n";
                }

                md += "## AxonASP Implementation Rules (Compact)\n\n";
                md += "### Source of Truth & Research Order\n";
                md += "1. Follow AxonASP patterns and compatibility rules first.\n";
                md += "2. Instruct the coding agent to try MCP documentation first.\n";
                md += "3. If MCP is unavailable, use AxonASP manual: https://g3pix.com.br/axonasp/manual/\n";
                md += "4. If needed, use GitHub manual index and follow links: https://github.com/guimaraeslucas/axonasp/blob/main/www/manual/menu.md\n\n";

                md += "### Language-Specific Server Rules\n";
                if (serverLang === "VBScript") {
                    md += "- Use Classic ASP VBScript page directive on top: <" + "%@ Language=\"VBSCRIPT\" %" + ">.\n";
                    md += "- Use Option Explicit and declare variables before use.\n";
                    md += "- Never write inline single-line If; always use block If...Then...End If.\n";
                    md += "- Close loops correctly (Next, Loop, Wend).\n";
                    md += "- Use Set for objects and release with Set obj = Nothing.\n";
                    md += "- Use & for string concatenation (never +).\n";
                    md += "- No short-circuit behavior; protect risky expressions with nested If blocks.\n";
                    md += "- Use On Error Resume Next only with immediate Err.Number checks and On Error GoTo 0 reset.\n";
                    md += "- Read the manual of AxonASP to avoid reinventing the wheel. Prefer to use AxonASP native implementations.\n\n";
                } else {
                    md += "- YOU MUST Use Classic ASP JavaScript (JScript ASP) directive on top of every page: <" + "%@ Language=\"JavaScript\" %" + ">.\n";
                    md += "- Keep compatibility with classic JScript behavior used by ASP pages.\n";
                    md += "- Avoid modern browser-only APIs in server-side code unless explicitly supported.\n";
                    md += "- Keep code deterministic, defensive, and aligned with AxonASP server runtime constraints.\n";
                    md += "- Release native objects and close resources explicitly when applicable.\n";
                    md += "- Read the manual of AxonASP to avoid reinventing the wheel. Prefer to use AxonASP native implementations.\n\n";
                }

                md += "### AxonASP Native Libraries (Prefer Over Reinventing)\n";
                md += "- Prefer built-in libraries when they already solve the task: G3JSON, G3DB, G3Files, G3HTTP, G3MAIL, G3Image, G3Template, G3PDF, G3ZIP, G3Crypto.\n";
                md += "- Use exact ProgID/function names as documented.\n";
                md += "- Handle errors explicitly and close/cleanup resources.\n\n";

                md += "### Development Workflow\n";
                md += "1. Map requirements to modules/pages/routes.\n";
                md += "2. Design schema and data-access strategy.\n";
                md += "3. Implement server logic first, then UI pages.\n";
                md += "4. Validate error handling, edge cases, and resource cleanup.\n";
                md += "5. Review compatibility against AxonASP documentation references above.\n\n";

                md += "---\n## Final Verification Checklist\n\n";
                md += "- [ ] All variables declared with Option Explicit\n";
                md +=
                    "- [ ] Objects assigned with Set and released with Nothing\n";
                md += "- [ ] String concatenation uses & operator\n";
                md += "- [ ] Array access protected with bounds checking\n";
                md += "- [ ] Error handling implemented throughout\n";
                md +=
                    "- [ ] No single-line/Inline If statements in code - this is critical, if fail the code will break \n";
                md += "- [ ] All loops properly closed (Next, Loop, Wend)\n";
                md += "- [ ] Method calls use correct parenthesis rules\n";
                md += "- [ ] Database connections properly closed\n";
                md += "- [ ] No hardcoded SQL values (use parameters)\n";
                md += "- [ ] File handles properly closed\n";
                md += "- [ ] Code tested and following Classic ASP pattern\n";
                md += "- [ ] Coding agent attempted MCP docs first\n";
                md += "- [ ] If MCP unavailable, manual fallback links were used\n\n";

                document.getElementById("output-textarea").value = md;
                document.getElementById("output-area").classList.add("show");
                document
                    .getElementById("output-area")
                    .scrollIntoView({ behavior: "smooth", block: "nearest" });
            }

            function copyToClipboard() {
                var textarea = document.getElementById("output-textarea");
                if (!textarea.value) {
                    alert("Please generate document first.");
                    return;
                }

                textarea.select();
                try {
                    document.execCommand("copy");
                    var toast = document.getElementById("copy-toast");
                    toast.classList.add("show");
                    setTimeout(function () {
                        toast.classList.remove("show");
                    }, 2500);
                } catch (err) {
                    alert("Failed to copy to clipboard.");
                }
            }

            function clearForm() {
                if (confirm("Reset all fields to defaults?")) {
                    document.getElementById("description").value = "";
                    document.getElementById("appname").value = "";
                    document.getElementById("extra").value = "";
                    document.getElementById("style-mvc").checked = true;
                    document.getElementById("db-sqlite").checked = true;
                    document.getElementById("lang-vbscript").checked = true;
                    document.getElementById("ui-axonasp").checked = true;
                    document.getElementById("js-vanilla").checked = true;
                    document.getElementById("color-light").checked = true;
                    document.getElementById("emoji-enabled").checked = true;
                    document.getElementById("feat-auth").checked = true;
                    document.getElementById("feat-sample").checked = true;
                    document.getElementById("feat-crud").checked = true;
                    document.getElementById("multilingual").checked = false;
                    document
                        .getElementById("output-area")
                        .classList.remove("show");
                    document
                        .querySelectorAll(
                            ".radio-item.selected, .check-item.selected"
                        )
                        .forEach(function (el) {
                            el.classList.remove("selected");
                        });
                    initMultiLang();
                }
            }

            // Initialize on load
            document.addEventListener("DOMContentLoaded", function () {
                initMultiLang();
            });
        </script>
    </body>

</html>