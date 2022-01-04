BASEDIR := $(CURDIR)
APP_TARGETS := Printer

ifeq ($(OS), Windows_NT)
	BUILD_OS := windows
	OUTPUT_EXT = .exe
	RM_WRAPPER = @cmd.exe /c rmdir /s /q $(subst /,\,$(1))
else
	BUILD_OS := unix
	OUTPUT_EXT :=
	RM_WRAPPER = @rm -rf $(1)
endif

APPS_DIR := $(BASEDIR)/cmd
OUTPUT_DIR := $(BASEDIR)/out

CLEAN_TARGETS := $(addprefix $(OUTPUT_DIR)/,$(APP_TARGETS))

.DEFAULT_GOAL := build
.PHONY: build clean lint modules run vet

lint:
	go fmt ./...

vet: lint
	go vet ./...

modules:
	go mod tidy

$(APP_TARGETS): vet modules

build: $(APP_TARGETS)
	go build -v -o $(OUTPUT_DIR)/$^ $(APPS_DIR)/main.go

run: vet
	go run $(APPS_DIR)/main.go

clean: $(OUTPUT_DIR)
	$(call RM_WRAPPER,$^)
