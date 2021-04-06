# Screenshot

A CLI for taking a screenshot of a web element using
[Selenium](https://github.com/SeleniumHQ/docker-selenium) (Docker) and the
[tebeka/selenium](https://github.com/tebeka/selenium) package.

## Setup

Start a Firefox Selenium container:

```bash
docker run -d -p 4444:4444 -p 7900:7900 --shm-size 2g selenium/standalone-firefox:4.0.0-beta-3-prerelease-20210402
```

> **[Full Documentation →](https://github.com/SeleniumHQ/docker-selenium)**

## Usage

Build the binary:

```bash
make run cmd=screenshot
```

View CLI usage:

```bash
cmd/screenshot/screenshot
```

## Example

```bash
cmd/screenshot/screenshot -u https://news.ycombinator.com/ element -s="#hnmain"
```

## Debugging

View what is happening inside the Selenium container by visiting
[http://localhost:7900](http://localhost:7900), the default password is
`secret`.
