# url-shortener written in Go

## Development

- Install go from https://go.dev/dl/.
- Download the standalone version of tailwind from https://github.com/tailwindlabs/tailwindcss/releases/latest, copy the executable to `internal/tailwind/` and rename it to `tailwind.run`.
- Run `go generate` to generate the css file.
- Run `go run server/main.go` to start the server.
