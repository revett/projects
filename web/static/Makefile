DEFAULT: run

build:
	hugo --buildDrafts --minify --verbose

compile-tailwind:
	npx tailwindcss-cli@latest build themes/kasparov/assets/css/tailwind.css -c tailwind.config.js -o themes/kasparov/assets/css/tailwind-build.css

run:
	hugo server --buildDrafts --verbose
