# G3AxonLive Ecosystem Overview

The G3AxonLive ecosystem is a high-performance Reactive Component Framework built directly into the AxonASP Virtual Machine. It empowers developers to create dynamic, stateful, and highly responsive web applications using Classic ASP (VBScript or Server-Side JavaScript) without requiring full page reloads.

By utilizing native Go mechanisms on the backend and modern `fetch` protocols on the frontend, AxonLive bridges the gap between legacy server-side scripting and modern single-page application (SPA) paradigms.

## Advantages of G3AxonLive

* **Zero Page Reloads:** All UI interactions (button clicks, form submissions, timers) are sent to the server asynchronously. The server responds with targeted JSON patches, swapping only the modified DOM elements.
* **Strict Backend Control:** All business logic, validation, and state mutation happen exclusively on the server. The client browser merely acts as a dumb terminal rendering the HTML patches, significantly reducing the attack surface.
* **Authenticated Session Binding:** The `/g3al` endpoint binds every async event to the authenticated `ASPSESSIONID` cookie. Client-provided session identifiers are not used as an authority for page routing.
* **Zero Additional Wrappers:** AxonLive is implemented directly inside the `axonvm` engine as a native procedural controller (`G3AXONLIVE`). This eliminates the need for bulky ASP class wrappers, providing bare-metal performance and zero garbage collection overhead.
* **Granular DOM Manipulation:** Instead of re-rendering entire components, developers can push targeted instructions to modify styles, attributes, classes, or trigger external redirects natively from ASP.

## State and Persistence

One of the most powerful features of G3AxonLive is its transparent **State and Persistence** model. 

When an AxonLive element or property is manipulated and saved in the session (e.g., using the `SetComponentProperty` method), **it persists across all pages for that user's session.** 

If you create an element state named `"main"` in the session, it will render identically and maintain its exact state across the user's entire session footprint, regardless of which physical `.asp` file they navigate to. 

This state is held in an optimized, thread-safe memory map within the Go server process, bypassing the overhead of standard ASP Session variables. The system includes a background garbage collector that automatically removes orphaned component states when a session times out.

## Security Model

G3AxonLive asynchronous events are resolved using the authenticated ASP session cookie (`ASPSESSIONID`).

* The server validates event ownership against the authenticated cookie session.
* A mismatched event `sessionId` in client payload is rejected.
* Page re-execution for async events always uses the authenticated session registration.

This model prevents cross-session routing when a client attempts to submit an event for a different user session.
