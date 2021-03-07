.PHONY: run watch build mocks rename test coover

run:
	go run main.go

test:
	go test -coverprofile=coverage.out ./...

docker:
	GOOS=linux GOARCH=386 go build cmd/upwork.go
	docker build . --tag upwork-scraping:local

docker_run:
	docker volume create docker_store_json
	docker run -v docker_store_json:/tmp/upwork-scrapping/store_json upwork-scraping:local