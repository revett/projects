build: delete-dist
	node_modules/parcel-bundler/bin/cli.js build index.html

compile-tailwind:
	npx tailwindcss-cli@latest build css/styles.css -c tailwind.config.js -o css/tailwind.css

delete-dist:
	rm dist/*

run: delete-dist
	node_modules/parcel-bundler/bin/cli.js index.html

