install go: sudo apt get golang-go
install templates tool : go get -t github.com/a-h/templ/cmd/templ@latest~
install templates : go get github.com/a-h/templ@latest
install hot reload: go install github.com/air-verse/air@latest

run: go run main.go

generate template: templ generate (or in watch mode: templ generate --watch)

air for hot reload