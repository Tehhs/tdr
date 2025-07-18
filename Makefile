
ANTLR = docker compose run --rm antlr4 antlr 

.PHONY: build
build:
	@echo "Building..."
	mkdir -p dist
	cd tool && go build -o ../dist/tdr

.PHONY: generate
generate:
	@echo "Generating parser..."
	$(ANTLR) -Dlanguage=Go -Xexact-output-dir \
		-o tool/pkg/tdrl \
		-package tdrl \
		tool/pkg/tdrl/grammar/tdrl.g4 