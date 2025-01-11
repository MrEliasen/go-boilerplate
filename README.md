# GO Boilerplate

Nothing fancy, made it for myself for quick project hacking.

## Main features

- SQLite & LibSQL databases support
- Simple DB migration system
- [Echo](https://echo.labstack.com) web server with [Templ](https://templ.guide/) support
- Minimal [Goth Multi-Provider](https://github.com/markbates/goth) authentication implemented

## Setup/Dev

1. Install Go a-h/Templ: `go install github.com/a-h/templ/cmd/templ@latest`
2. Install project packages: `go mod tidy` 
3. Run Templ Watcher: `make dev`
4. Run Project: `go run ./cmd/boiler`
