PROJECT_NAME = apple-health-data-workflow

build:
	docker build -t $(PROJECT_NAME) .

run: build
	docker run $(PROJECT_NAME)
