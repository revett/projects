import data from "./geojson.json";

import mapboxgl from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";

mapboxgl.accessToken =
  "pk.eyJ1IjoicmV2Y2QiLCJhIjoiY2ttcnQxYmNyMGI1cjJxcGJ1dHlhdXF6diJ9.Bp0j1asGrBZ9DjZ2LYqplQ";

const p = window.location.pathname.split("/");
const filename = p.pop() || p.pop();
const geojson = data[filename];

const coordinates = geojson.geometry.coordinates;
const bounds = coordinates.reduce(function (bounds, coord) {
  return bounds.extend(coord);
}, new mapboxgl.LngLatBounds(coordinates[0], coordinates[0]));

const map = new mapboxgl.Map({
  bounds: bounds,
  container: "map",
  fitBoundsOptions: {
    padding: 80,
  },
  style: "mapbox://styles/mapbox/outdoors-v11",
});

map.on("load", function () {
  map.addSource("route", {
    type: "geojson",
    data: geojson,
  });

  map.addLayer({
    id: "route",
    type: "line",
    source: "route",
    layout: {
      "line-join": "round",
      "line-cap": "round",
    },
    paint: {
      "line-color": "#1a602f",
      "line-width": 3,
    },
  });
});
