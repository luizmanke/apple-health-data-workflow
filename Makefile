build:
	docker build -t apple-health-data-workflow .

start: build
	docker compose up frontend -d

logs:
	docker compose logs --follow frontend backend

stop:
	docker compose down

integration-tests: build
	trap "make stop" EXIT; make integration-tests-run

integration-tests-run:
	docker compose run \
		-v $(PWD)/test/integration:/go/src/test:ro \
		backend go test ./test/...

tests: integration-tests
