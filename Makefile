build:
	go build -o chatbot

run: build	
	./chatbot

mockery:
	go install github.com/vektra/mockery/v2@latest

mocks: mockery
	mockery --dir internal --with-expecter --keeptree --all

.PHONY: test
test:
	go test -v -cover  ./internal/delivery/... ./internal/repo/... ./internal/usecase/... > test_report.txt
	go test -coverprofile=coverage.out  ./internal/delivery/... ./internal/repo/... ./internal/usecase/...
	go tool cover -html=coverage.out -o coverage.html		
	explorer.exe coverage.html