test:
<<<<<<< Updated upstream
	go test ./...
=======
	docker compose exec analyser go test ./...
>>>>>>> Stashed changes

mocks:
	docker run -v "$$PWD":/src -w /src vektra/mockery:3
