build:
	docker build -t apple-health-data-workflow .

run: build
	trap "make stop" EXIT; docker compose run app

stop:
	docker compose down

integration-tests: build
	trap "make stop" EXIT; make integration-tests-run

integration-tests-run:
	docker compose run \
		-v $(PWD)/test/integration:/go/src/app/test \
		app go test ./test/...

tests: integration-tests
