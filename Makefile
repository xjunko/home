BINARY_FILE=eva.bin

all: $(BINARY_FILE)

$(BINARY_FILE):
	@if [ ! -f $(BINARY_FILE) ]; then \
		echo "Building $(BINARY_FILE)!"; \
		go build -o $(BINARY_FILE) ./internal/cmd/eva/eva.go; \
	fi

clean:
	rm -f $(BINARY_FILE)

run: $(BINARY_FILE)
	./$(BINARY_FILE)