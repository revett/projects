DEFAULT: run

build:
	hugo --buildDrafts --minify --verbose

js:
	MAPBOX_TOKEN=$(MAPBOX_TOKEN) ./node_modules/parcel-bundler/bin/cli.js watch -d static/js assets/js/*.js

run:
	hugo server --buildDrafts --verbose
