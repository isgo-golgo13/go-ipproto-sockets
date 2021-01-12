 # Go parameters
    GOCMD=go
    GOBUILD=go build
    GOCLEAN=go clean
    GOTEST=go test
    CLIENT=client_svc
 
    all: build
    build: 
		$ env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(CLIENT) -v
    test: 
		$(GOTEST) -v ./...
    clean: 
		$(GOCLEAN)
			rm -f $(CLIENT)
    run:
		$ env GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(CLIENT) -v
		./$(CLIENT)