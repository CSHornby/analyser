default:
	docker compose up

test:
	docker compose exec analyser go test ./...

mocks:
	docker run -v "$$PWD":/src -w /src vektra/mockery:3
