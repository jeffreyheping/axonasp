# The Client-Side Script

For the G3AxonLive ecosystem to function, the client-side JavaScript bridge (`g3axonlive.js`) must be included and initialized on every single page utilizing AxonLive features.

The script acts as the neural network of the frontend, binding to DOM elements, intercepting interactions, and communicating securely with the server-side procedural controller.

## Including the Script

You must include the script at the bottom of your HTML body and immediately call `G3AxonLive.init()`.

```html
<script src="/axonlive/g3axonlive.js"></script>
<script>
    // Initialize the engine
    G3AxonLive.init();
</script>
```

The optional `sessionId` argument is retained for backward compatibility, but it is no longer required.

## How It Works Under the Hood

The client bridge is written in vanilla JavaScript and performs several critical operations:

1. **Event Delegation:** It scans the DOM for elements marked with `data-g3al-*` attributes (e.g., `data-g3al-id`, `data-g3al-event`). It automatically attaches event listeners without requiring custom JavaScript for each button or input.
2. **Fetch API:** When an event is triggered, the script prevents the default browser action (like a full form submission) and instead utilizes the modern `fetch` API.
3. **Payload Generation:** It constructs a JSON payload containing the `componentId`, `eventName`, and contextual `eventArgs` (such as values from text inputs or custom data attributes).
4. **Targeted DOM Swapping:** The server responds with a JSON array of component patches. The script locates each component by its ID and precisely updates its `outerHTML`, preserving the rest of the application layout.
5. **Action Execution:** Beyond simple HTML swaps, the bridge processes server-instructed actions, such as `set_timer`, `redirect`, `add_class`, or `set_attribute`. This allows the server to manipulate the client dynamically.

## Session Binding and Security

The browser bridge does not control session authority.

* The server binds every async event to the authenticated `ASPSESSIONID` cookie.
* Async page routing and execution use the authenticated session registration.
* Any mismatched `sessionId` value in payload is rejected by the server.

## Handling Network Errors

The bridge includes built-in retry mechanisms and exponential backoff for transient network errors. If the backend becomes unreachable, it will gracefully log the failure to the console and attempt to recover without crashing the user interface.