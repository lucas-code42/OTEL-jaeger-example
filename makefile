run:
	docker compose up --build

test:
	curl localhost:8080/servicea/ping

stop:
	docker compose down