import { client } from "./GameClient";
import { Map } from "./Map";
import "./leaflet";

const map = new Map("map");
map.planeMarker.rotate(30);

client.eventbus.on("data", (data) => {
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
