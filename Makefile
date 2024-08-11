# Makefile

# Definir el directorio de salida para el binario
BINARY_DIR=./tmp
BINARY_NAME=main

all: generate build

generate:
	@echo "Generando código..."
	templ generate

build:
	@echo "Construyendo el proyecto..."
	go build -o $(BINARY_DIR)/$(BINARY_NAME) cmd/main.go

clean:
	@echo "Limpiando..."
	rm -rf $(BINARY_DIR)/$(BINARY_NAME)

.PHONY: generate build clean
