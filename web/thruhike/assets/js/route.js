import axios from "axios";
import mapbox from "mapbox-gl";
import "mapbox-gl/dist/mapbox-gl.css";

mapbox.accessToken = process.env.MAPBOX_TOKEN;

const addMarkers = (m, d) => {
  const coordinates = mergeCoordinatesFromFeatures(d);
  const markerLocs = [coordinates[0], coordinates[coordinates.length - 1]];

  const markerOpts = {
    color: "#197e33",
    scale: 0.5,
  };

  for (const loc of markerLocs) {
    new mapbox.Marker(markerOpts).setLngLat(loc).addTo(m);
  }
};

const addRouteLayer = (m, d) => {
  m.on("load", function () {
    m.addSource("route", {
      type: "geojson",
      data: d,
    });

    m.addLayer({
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

    addMarkers(m, d);
  });
};

const calculateBounds = (d) => {
  const coordinates = mergeCoordinatesFromFeatures(d);

  return coordinates.reduce(function (bounds, coord) {
    return bounds.extend(coord);
  }, new mapbox.LngLatBounds(coordinates[0], coordinates[0]));
};

const initialiseMap = (d) => {
  return new mapbox.Map({
    bounds: calculateBounds(d),
    container: "map",
    fitBoundsOptions: {
      padding: 80,
    },
    maxBounds: [-11.97, 49.2, 3.31, 59.76],
    style: "mapbox://styles/mapbox/outdoors-v11",
  });
};

const mergeCoordinatesFromFeatures = (d) => {
  let c = [];

  for (const f of d.features) {
    c.push(...f.geometry.coordinates);
  }

  return c;
};

const parseLastPathPart = () => {
  const p = window.location.pathname.split("/");
  return p.pop() || p.pop();
};

const renderRoute = () => {
  const routeName = parseLastPathPart();

  axios
    .get(`/data/${routeName}.json`)
    .then((r) => {
      const map = initialiseMap(r.data);
      addRouteLayer(map, r.data);
    })
    .catch((e) => {
      console.log(e);
    });
};

renderRoute();
