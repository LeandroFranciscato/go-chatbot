build:
	go build -o chatbot

run: build	
	./chatbot

mockery:
	go install github.com/vektra/mockery/v2@latest

mocks: mockery
	mockery --dir internal --with-expecter --keeptree --all