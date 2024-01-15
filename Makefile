run: 
	@templ generate
	@go run main.go

test: 
	@templ generate
	@go test /test