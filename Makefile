.PHONY: run
run:
	bun run dev

.PHONY: build
build:
	bun run build

.PHONY: clear
clear:
	rm -rf dist

.PHONY: prod-run
prod-run: build
	bun run build && \
	python -m http.server --directory dist