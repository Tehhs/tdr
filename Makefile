
ANTLR = docker compose run --rm antlr4 antlr 

.PHONY: build
build:
	@echo "Building..."
	mkdir -p dist
	go build -o ./dist/tdr

.PHONY: generate
generate:
	@echo "Generating parser..."
	$(ANTLR) -Dlanguage=Go -Xexact-output-dir \
		-o pkg/tdrl \
		-package tdrl \
		pkg/tdrl/grammar/tdrl.g4 


.PHONY: test
test:
	go test ./...