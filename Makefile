# Variables
MAIN_PACKAGE_PATH := .
BINARY_NAME := project-sprint-gogo-manager

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## build: build the application
.PHONY: build
build:
	go build -o=./tmp/bin/${BINARY_NAME} ${MAIN_PACKAGE_PATH}

## run: run the  application
.PHONY: run
run: build
	./tmp/bin/${BINARY_NAME}

## watch: run the application with reloading on file changes
.PHONY: watch
watch:
	go run github.com/air-verse/air@latest \
			--build.cmd "make build" --build.bin "./tmp/bin/${BINARY_NAME}" --build.full_bin "MODE=development ./tmp/bin/${BINARY_NAME}" \
			--build.delay "100" \
			--build.args_bin "--log-level debug" \
			--build.exclude_dir "" \
			--build.include_ext "go, mod, tpl, tmpl, html, env, toml" \
			--build.send_interrupt "true" \
			--build.kill_delay "5000000" \
			--misc.clean_on_exit "true"

## clean: remove the binary
.PHONY: clean
clean:
	rm -f ./tmp/bin/${binary_name}