#Credit card statement analyser

## Purpose
Two reasons:
- To show how I organise and structure my code
- To figure out where my money is going

## General approach
I aim to make the majority of code unit testable so for most dependencies interfaces are passed
This project does not include a DB, but if it did then DB functions would include system tests
Each file (except main.go) has a unit test
Mockery is used to make mocking super easy


## Things not included here
Projects I've worked on using docker compose
This ensures my local development environment matches production
The environment includes a proxy (if developing a web or API app) and typically a DB
