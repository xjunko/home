.PHONY: default
default:
	@echo "Available Targets:"
	@echo "  - setup      // Install dependencies"
	@echo "  - run        // Run the development server"
	@echo "  - build      // Build the project"
	@echo "  - clear      // Clear the build directory"
	@echo "  - prod-run   // Run the production server"

.PHONY: setup
setup:
	bun install

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