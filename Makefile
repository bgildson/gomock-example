test:
	@go test -v ./...

cover:
	@go test -coverprofile cover.out ./... && go tool cover -html=cover.out -o cover.html

clean:
	@rm cover.html cover.out

mockgen:
	@which mockgen > /dev/null || go get github.com/golang/mock/gomock && go install github.com/golang/mock/mockgen
	mockgen -source ./client/finalspace3/client.go -destination ./client/finalspace3/mock.go -package finalspace3
	mockgen -source ./pkg/net/http/client.go -destination ./pkg/net/http/mock.go -package http
