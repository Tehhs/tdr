.PHONY: build
build:
	@echo "Building..."
	mkdir -p dist
	cd tool && go build -o ../dist/tdr