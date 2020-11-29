import * as L from "leaflet";
import "leaflet-rotatedmarker";
import "leaflet.markercluster";
import { Emitter } from "./emitter";

const defaultPosition = [51.505, -0.09];
export class Map extends Emitter {
  /**
   * @param {string} divID
   */
  constructor(divID) {
    super();
    this.followPlane = true;
    this.teleportDestination = defaultPosition;

    // Create the map
    this.map = L.map(divID).setView(defaultPosition, 13);
    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
      maxZoom: 19,
      attribution:
        '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors',
    }).addTo(this.map);

    const OpenTopoMap = L.tileLayer(
      "https://{s}.tile.opentopomap.org/{z}/{x}/{y}.png",
      {
        maxZoom: 17,
        attribution:
          'Map data: &copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors, <a href="http://viewfinderpanoramas.org">SRTM</a> | Map style: &copy; <a href="https://opentopomap.org">OpenTopoMap</a> (<a href="https://creativecommons.org/licenses/by-sa/3.0/">CC-BY-SA</a>)',
      }
    );

    // Create plane marker
    this.planeMarker = new PlaneMarker(this.map);
    this.planeMarker.marker.on("click", () => {
      this.followPlane = true;
    });
    this.map.on("drag", (event) => {
      this.followPlane = false;
      this.emit("drag", event);
    });

    // Teleport marker functionality
    this.teleportmarker = L.marker([50.5, 30.5]).addTo(this.map);

    this.map.on("click", (event) => {
      this.emit("click", event);
      const { lat, lng } = event.latlng;
      this.teleportmarker.setLatLng([lat, lng]);
      this.teleportDestination = [lat, lng];
    });
    this.map.on("move", (e) => this.emit("move", e));
  }
  /**
   *
   * @param {[number,number]} center
   * @param {number} zoom
   */
  setView(center) {
    this.map.setView(center);
  }
}

class PlaneMarker {
  constructor(map) {
    const iconSize = 80;
    const planeIcon = L.icon({
      iconUrl: "/assets/images/plane.svg",
      iconSize: [iconSize, iconSize],
      iconAnchor: [iconSize / 2, iconSize / 2],
      popupAnchor: [0, iconSize / 2],
    });
    this.marker = L.marker(defaultPosition, {
      icon: planeIcon,
      rotationAngle: 0,
      rotationOrigin: "center",
    });
    this.marker.addTo(map);
  }
  /**
   *
   * @param {[number, number]} latLngTuple
   * @param {number} rotation
   */
  update(latLngTuple, rotation = 0) {
    this.marker.setLatLng(latLngTuple);
    this.rotate(rotation);
  }
  /**
   * @param {number} angle
   */
  rotate(angle) {
    this.marker.setRotationAngle(angle);
  }
}
