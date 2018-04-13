build:
	docker build -t npx/mods-api .
run:
	docker run --rm -it -p 8000:8000 npx/mods-api