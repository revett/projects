# ThruHike

## Usage

```bash
MAPBOX_TOKEN="..." make js
```

```bash
make run
```

## Adding A Route

> Follow the
> [Github Forking Guide](https://guides.github.com/activities/forking/) if you
> are contributing to the project.

Make sure that [Hugo](https://gohugo.io/getting-started/installing/) is
installed:

```bash
hugo version
```

Change to the `web/thruhike/` directory:

```bash
cd web/thruhike/
```

Add a new route content page (Markdown):

```bash
hugo new routes/glyndwrs-way.md
```

Add a new route data file (JSON):

```bash
hugo new data/glyndwrs-way.json
```

Add the GeoJSON blob to the new data file, making sure that it follows the
format of a `FeatureCollection` type:

```json
{
  "type": "FeatureCollection",
  "features": []
}
```

> Make use of the Mapbox
> [GPX/KML to GeoJSON](https://mapbox.github.io/togeojson/) tool for conversion.
