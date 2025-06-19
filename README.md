#Credit card statement analyser

## Purpose
Two reasons:
- To show how I organise and structure my code
- To figure out where my money is going

## Prerequisites and running
Make sure docker is running
To run app:
```bash
make
```

To run tests:
```bash
make test
```

To regenerate mocks
```bash
make mocks
```

air is used to enable hot reloads during development

## Usage
Visit http://localhost:8000
Upload a CSV with three columns: Date, entry description, amount
Click "Upload and Analyse"
Refine analysis by adding to services/analyse::getCategoryLookup

## General approach
I aim to make the majority of code unit testable so most dependencies are passed as interfaces alowing dependencies to be mocked in tests.
This project does not include a DB, but if it did then DB functions would include system tests
Each file (except main.go) has a unit test
https://vektra.github.io/mockery is used to make mocking super easy
