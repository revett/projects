import mapboxgl from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";

mapboxgl.accessToken = "...";

const coordinates = route.geometry.coordinates;
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
    data: route,
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

  const markerOpts = {
    color: "#197e33",
    scale: 0.5,
  };
  const markerLocs = [
    route.geometry.coordinates[0],
    route.geometry.coordinates[route.geometry.coordinates.length - 1],
  ];

  for (const loc of markerLocs) {
    new mapboxgl.Marker(markerOpts).setLngLat(loc).addTo(map);
  }
});
