# Initialize Go module
go mod init github.com/<username>/<repository-name>

# Download and organize dependencies
go mod tidy

# Build the application
go build

# Run the application
go run .

# Run tests
go test ./...

# Format Go code
go fmt ./...
