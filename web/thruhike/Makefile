DEFAULT: run

build:
	hugo --buildDrafts --minify --verbose

build-docker:
	hugo --buildDrafts --baseURL "http://nginx:80/" --minify --verbose

compose-down:
	docker-compose -f ../../deployments/thruhike/docker-compose.yml down

compose-up:
	docker-compose -f ../../deployments/thruhike/docker-compose.yml up -d

js-build:
	MAPBOX_TOKEN=$(MAPBOX_TOKEN) ./node_modules/parcel-bundler/bin/cli.js build -d static/js assets/js/*.js

js-watch:
	MAPBOX_TOKEN=$(MAPBOX_TOKEN) ./node_modules/parcel-bundler/bin/cli.js watch -d static/js assets/js/*.js

run:
	hugo server --buildDrafts --disableFastRender --verbose
