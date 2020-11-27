export class Emitter {
  constructor(listeners = {}) {
    this.listeners = listeners;
  }
  /**
   * Register an event handler for the given type.
   *
   * @param  {String} type	Type of event to listen for, or `"*"` for all events
   * @param  {Function} handler Function to call in response to given event
   */
  on(type, handler) {
    this.listeners[type] = this.listeners[type] || [];
    this.listeners[type].push(handler);
  }

  /**
   * Remove an event handler for the given type.
   *
   * @param  {String} type	Type of event to unregister `handler` from, or `"*"`
   * @param  {Function} handler Handler function to remove
   */
  off(type, handler) {
    if (this.listeners[type]) {
      this.listeners[type].splice(
        this.listeners[type].indexOf(handler) >>> 0,
        1
      );
    }
  }

  /**
   * Invoke this.listeners handlers for the given type.
   * If present, `"*"` handlers are invoked after type-matched handlers.
   *
   * Note: Manually firing "*" handlers is not supported.
   *
   * @param {String} type  The event type to invoke
   * @param {Any} [evt]  Any value (object is recommended and powerful), passed to each handler
   */
  emit(type, evt) {
    (this.listeners[type] || []).slice().map((handler) => {
      handler(evt);
    });
    (this.listeners["*"] || []).slice().map((handler) => {
      handler(type, evt);
    });
  }
}
