all: api cli

api:
	go build -o build/quiz-$@ cmd/api/main.go
.PHONY: api

cli:
	go build -o build/quiz-$@ cmd/cli/main.go
.PHONY: cli

clean:
	rm -rf build
.PHONY: clean

re: clean
	$(MAKE)
.PHONY: re
