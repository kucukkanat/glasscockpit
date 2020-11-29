import { Emitter } from "./emitter";

/**
 * TeleportObject
 * @typedef {Object} TeleportObject
 * @property {string} ID - Airport ID
 * @property {string} Type - Airport Type
 * @property {string} Name - Airport Name
 * @property {number} Latitude - Airport Latitude
 * @property {number} Longitude - Airport Longitude as number
 * @property {number} Elevation - Airport Altitude as number
 * @property {number} Country - Airport Country Code
 * @property {string} Code - Airport GPS Code
 */

/**
 * Airport
 * @typedef {Object} Airport
 * @property {number} lat - Latitude as number
 * @property {number} lng - Longitude as number
 * @property {number} hdg - Heading as number
 * @property {number} alt - Altitude as number
 */

export class GameClient extends Emitter {
  constructor() {
    super();
    const {
      location: { host },
    } = window;
    this.socket = new WebSocket(`ws://${host}/ws`);
    this.apiURL = `http://${host}`;
    this.socket.onopen = (event) => {
      this.emit("socket:open", event);
    };
    this.socket.onclose = (event) => {
      this.emit("socket:close", event);
    };
    this.socket.onerror = (event) => {
      this.emit("socket:error", event);
    };
    this.socket.onmessage = (event) => {
      try {
        const data = JSON.parse(event.data);
        this.emit("data", data);
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
    return this.request("/event", event);
  }
  /**
   *
   * @param {TeleportObject} teleportObject
   * @returns {Promise<any>}
   */
  teleport(teleportObject) {
    return this.request("/teleport", teleportObject);
  }
  /**
   *
   * @param {number} southWest
   * @param {number} northEast
   * @returns {Promise<Airport[]>}
   */
  getNearbyAirports(southWest, northEast) {
    return this.request("/airports", {
      southWest,
      northEast,
    });
  }
  /**
   *
   * @param {string} path
   * @param {object} data
   */
  request(path, data) {
    return fetch(`${this.apiURL}${path}`, {
      method: "post",
      body: JSON.stringify(data),
    }).then((r) => r.json());
  }
}

export const client = new GameClient();
