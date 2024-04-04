build:
	docker build -t apple-health-data-workflow .

run: build
	docker compose up frontend -d

stop:
	docker compose down

integration-tests: build
	trap "make stop" EXIT; make integration-tests-run

integration-tests-run:
	docker compose run \
		-v $(PWD)/test/integration:/go/src/test:ro \
		frontend go test ./test/...

tests: integration-tests
