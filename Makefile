.PHONY: dev build
.DEFAULT_GOAL := dev

dev:
	@templ generate --watch

build:
	@templ generate

