
ANTLR = docker compose run --rm antlr4 antlr 

.PHONY: build
build:
	@echo "Building..."
	mkdir -p dist
	cd tool && go build -o ../dist/tdr

.PHONY: generate
generate:
	@echo "Generating parser..."
	$(ANTLR) tool/antlr/grammar/tdrl.g4