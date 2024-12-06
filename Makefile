run:
	cd cmd/bin; go run main.go

test:
	go test -v -cover ./...

compose:
	sudo docker-compose up --build
