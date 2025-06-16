test:
	go test ./...

mocks:
	docker run -v "$$PWD":/src -w /src vektra/mockery:3
