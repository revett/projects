build: delete-dist
	node_modules/parcel-bundler/bin/cli.js build src/index.html

compile-markdown:
	node_modules/marked/bin/marked -i partials/test.md -o partials/test.partial.html --gfm

compile-tailwind:
	npx tailwindcss-cli@latest build src/css/styles.css -c tailwind.config.js -o src/css/tailwind.css

delete-dist:
	rm -f dist/*

run: delete-dist
	node_modules/parcel-bundler/bin/cli.js src/index.html

