import * as L from "leaflet";
import "leaflet-rotatedmarker";
const defaultPosition = [51.505, -0.09];
export class Map {
  /**
   * @param {string} divID
   */
  constructor(divID) {
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
    this.planeMarker = new PlaneMarker();
    this.planeMarker.addToMap(this.map);
  }
  /**
   *
   * @param {[number,number]} center
   * @param {number} zoom
   */
  setView(center, zoom = 13) {
    this.map.setView(center, zoom);
  }
}

class PlaneMarker {
  constructor() {
    const iconSize = 40;
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
  }
  addToMap(map) {
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
