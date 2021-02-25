build: delete-dist
	node_modules/parcel-bundler/bin/cli.js build src/index.html

compile-markdown:
	node_modules/marked/bin/marked -i src/partials/header.md -o src/partials/header.partial.html --gfm

compile-tailwind:
	npx tailwindcss-cli@latest build src/css/styles.css -c tailwind.config.js -o src/css/tailwind.css

delete-dist:
	rm -rf dist/*
	mkdir -p dist/pages/
	cp -r src/img/ dist/pages/img/

run: delete-dist
	node_modules/parcel-bundler/bin/cli.js src/index.html

