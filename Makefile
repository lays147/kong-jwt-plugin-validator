build-kong:
	docker build -t kong-local .

start-kong:
	$(MAKE) build-kong
	docker-compose up

