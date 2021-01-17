import { client } from "./GameClient";
import { Map } from "./Map";
import { debounce } from "./debounce";

const map = new Map("map");
map.planeMarker.rotate(30);
// Request nearby airports and cluster
const markerClusterGroup = L.markerClusterGroup();
const addedAirports = [];
const getNearbyAirports = debounce(function () {
  const { _southWest, _northEast } = map.map.getBounds();
  client.getNearbyAirports(_southWest, _northEast).then((nearbyAirports) => {
    nearbyAirports.forEach((nearbyAirport) => {
      if (!addedAirports.includes(nearbyAirport.ID) && nearbyAirport.Code) {
        const title = nearbyAirport.Code;
        const marker = L.marker(
          new L.LatLng(nearbyAirport.Latitude, nearbyAirport.Longitude),
          { title }
        );
        marker.bindPopup(title);
        markerClusterGroup.addLayer(marker);
        addedAirports.push(nearbyAirport.ID);
      }
    });
    // remove and redraw
    map.map.addLayer(markerClusterGroup);
  });
}, 500);

map.on("move", async () => {
  if (map.map.getZoom() >= 7) {
    // Fetch airports only if zoomed enough
    getNearbyAirports();
  }
});

client.on("data", (data) => {
  const lat = data["PLANE LATITUDE"];
  const lng = data["PLANE LONGITUDE"];
  const heading = data["PLANE HEADING DEGREES TRUE"];
  map.planeMarker.update([lat, lng], heading);
  if (map.followPlane) {
    map.setView([lat, lng]);
  }
});

document.querySelector(".teleport_button").addEventListener("click", () => {
  const [lat, lng] = map.teleportDestination;
  client.teleport({
    lat: lat,
    lng: lng,
    hdg: parseInt(document.querySelector("[name=heading]").value),
    alt: parseInt(document.querySelector("[name=altitude]").value),
  });
});

document.getElementById("sim_rate_plus").addEventListener("click", async () => {
  await client.sendEvent({ eventname: "KEY_SIM_RATE" });
  await client.sendEvent({ eventname: "PLUS" });
});

document
  .getElementById("sim_rate_minus")
  .addEventListener("click", async () => {
    await client.sendEvent({ eventname: "KEY_SIM_RATE" });
    await client.sendEvent({ eventname: "MINUS" });
  });
