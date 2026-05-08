# AxonLive Builder

The **AxonLive Builder** is an integrated visual IDE provided directly within the AxonASP ecosystem, designed to accelerate the development of reactive pages. 

Located within the `/www/axonlive/builder/` directory, this powerful Single Page Application (SPA) acts as a drag-and-drop interface where developers can visually construct their ASP pages with pre-built AxonLive components.

## Features

* **Visual Construction:** Drag and drop elements like buttons, counters, text inputs, grids, and timers onto a live canvas.
* **Server-Side JavaScript Generation:** The builder pivotally generates boilerplate **Server-Side JavaScript** (JScript) instead of VBScript. Server-side JS handles JSON payloads and mapping structures natively, making it a much cleaner output for event routing.
* **Automatic Boilerplate:** As you assemble your interface, the Builder simultaneously generates the required procedural routing logic, `InitPage` calls, and structural HTML mapping.
* **Instant Export:** You can instantly preview the code, copy it to the clipboard, or download it as a ready-to-run `.asp` file.

## Why Server-Side JavaScript?

While G3AxonLive fully supports VBScript, the Builder generates JScript because of its superior ability to handle `switch/case` structures and data dictionaries cleanly. The generated code strictly uses switch maps for event routing to guarantee that only developer-defined functions are triggered by client events, ensuring robust security against injection.

By utilizing the Builder, you can skip the repetitive tasks of defining `data-g3al-` attributes and writing standard switch statements, allowing you to focus entirely on writing the business logic inside the generated event handlers.

## Security Behavior

The Builder-generated pages use the same AxonLive security model as manual pages.

* Async event routing is bound to the authenticated `ASPSESSIONID` cookie.
* Builder-generated client events do not require a session ID payload field.
* If a `sessionId` is present in a custom payload, it must match the authenticated cookie session.