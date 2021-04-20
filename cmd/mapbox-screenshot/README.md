# Mapbox Screenshot

A CLI for taking screenshots of web elements using
[Selenium](https://github.com/SeleniumHQ/docker-selenium) (Docker) and the
[tebeka/selenium](https://github.com/tebeka/selenium) package.

## Setup

Start a Firefox Selenium container:

```bash
docker run -d -p 4444:4444 -p 7900:7900 selenium/standalone-firefox
```

> **[Full Documentation →](https://github.com/SeleniumHQ/docker-selenium)**

Build the binary:

```bash
make run cmd=mapbox-screenshot
```

> Using the **root** Makefile.

Change to that directory:

```bash
cd cmd/mapbox-screenshot/
```

## Usage

Add the following to your
[`mapboxgl`](https://docs.mapbox.com/mapbox-gl-js) implementation:

```js
map.on("idle", () => {
  localStorage.setItem("mapbox.StylesLoaded", true);
});
```

> Currently the most consistent way to know when the styles of a Mapbox map
> have loaded, see
> [mapbox/mapbox-gl-js#2268](https://github.com/mapbox/mapbox-gl-js/issues/2268#issuecomment-536782308)

## Example

```bash
mapbox-screenshot capture http://localhost:1313/ "#map"
```

## Debugging

View what is happening inside the Selenium container by visiting
[http://localhost:7900](http://localhost:7900), the default password is
`secret`.
