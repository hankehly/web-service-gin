build:
	docker build -t hankehly/web-service-gin -f Containerfile .

run:
	docker run --rm -p 8080:8080 hankehly/web-service-gin
