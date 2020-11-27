import { GameClient } from "./GameClient";
import { Map } from "./Map";
import "./leaflet";

const client = new GameClient();
const map = new Map("map");
map.planeMarker.rotate(30);

client.eventbus.on("data", (data) => {
  const lat = data["PLANE LATITUDE"];
  const lng = data["PLANE LONGITUDE"];
  map.planeMarker.update([lat, lng]);
  console.log({ data });
});
