test:
	go test ./... -v -cover

build:
	docker build  --build-arg PORT=$(PORT) -t uber .
	docker run --env-file .env --network="host" -d --name uber uber

run:
	go run main.go