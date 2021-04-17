# API Specification

Public REST APIs within this repo are documented using
[API Blueprint](https://github.com/apiaryio/api-blueprint).

## Mock API

Build the local [Drakov](https://github.com/Aconex/drakov) Docker image:

```bash
docker build -f build/package/drakov.dockerfile -t revett/projects/drakov .
```

Start a new container, watching for changes in the `.apib` file:

```bash
docker run -it --init --rm \
  -v $(pwd)/api/thruhike:/api \
  -p 4587:4587 \
  --name thruhike-content-mock-api \
  revett/projects/drakov -f /api/thruhike/content_api.apib --watch
```

Go to: `http://localhost:4587/...`

## Documentation

Build the local [Aglio](https://github.com/danielgtaylor/aglio) Docker image:

```bash
docker build -f build/package/aglio.dockerfile -t revett/projects/aglio .
```

Start a new container:

```bash
docker run -it --init --rm \
  -v $(pwd)/api/thruhike:/api \
  -p 5687:5687 \
  --name thruhike-content-api-docs \
  revett/projects/aglio -i /api/thruhike/content_api.apib
```

Go to: `http://localhost:5687`
