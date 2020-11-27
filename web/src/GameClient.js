import { Emitter } from "./emitter";

/**
 * TeleportObject
 * @typedef {Object} TeleportObject
 * @property {number} lat - Latitude as number
 * @property {number} lng - Longitude as number
 * @property {number} hdg - Heading as number
 * @property {number} alt - Altitude as number
 */

export class GameClient {
  constructor() {
    const {
      location: { host },
    } = window;
    this.socket = new WebSocket(`ws://${host}/ws`);
    this.apiURL = `http://${host}`;
    this.eventbus = new Emitter();
    this.socket.onopen = (event) => {
      this.eventbus.emit("socket:open", event);
    };
    this.socket.onclose = (event) => {
      this.eventbus.emit("socket:close", event);
    };
    this.socket.onerror = (event) => {
      this.eventbus.emit("socket:error", event);
    };
    this.socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        this.eventbus.emit("data", data);
      } catch (err) {
        console.error({ err });
      }
    };
  }
  /**
   *
   * @param {string} event  - Event name to send
   */
  sendEvent(event) {
    return fetch(this.apiURL, {
      method: "post",
      body: JSON.stringify({ event }),
    });
  }
  /**
   *
   * @param {TeleportObject} teleportObject
   */
  teleport(teleportObject) {
    return fetch(`${this.apiURL}/teleport`, {
      method: "post",
      body: JSON.stringify(teleportObject),
    });
  }
}
