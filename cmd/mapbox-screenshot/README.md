# Screenshot

A CLI for taking screenshots of web elements using
[Selenium](https://github.com/SeleniumHQ/docker-selenium) (Docker) and the
[tebeka/selenium](https://github.com/tebeka/selenium) package.

## Setup

Start a Firefox Selenium container:

```bash
docker run -d -p 4444:4444 -p 7900:7900 selenium/standalone-firefox
```

> **[Full Documentation →](https://github.com/SeleniumHQ/docker-selenium)**

## Usage

Build the binary:

```bash
make run cmd=mapbox-screenshot
```

> Using the **root** Makefile.

Change to that director:

```bash
cd cmd/mapbox-screenshot/
```

View CLI usage:

```bash
mapbox-screenshot -h
```

### Example

```bash
mapbox-screenshot capture https://news.ycombinator.com/ "#hnmain"
```

### Debugging

View what is happening inside the Selenium container by visiting
[http://localhost:7900](http://localhost:7900), the default password is
`secret`.
