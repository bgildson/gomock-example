test:
	@go test -v ./...

cover:
	@go test -coverprofile cover.out ./... && go tool cover -html=cover.out -o cover.html

clean:
	@rm cover.html cover.out
