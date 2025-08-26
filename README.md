# Go Chat Application

A real-time chat application built with Go using the Chi router framework.

## Current Phase: Basic HTTP Server Setup

### What's Implemented

- **HTTP Server Foundation**: Basic HTTP server with Chi router and middleware
- **Health Check Endpoint**: `/v1/health` endpoint for monitoring server status
- **Environment Configuration**: Configurable server address via environment variables
- **Graceful Server Management**: Proper timeouts and server configuration

### Project Structure

```
go-chat/
├── cmd/api/           # Application entrypoint
│   ├── main.go       # Server startup and configuration
│   ├── api.go        # HTTP handlers and routing
│   └── health.go     # Health check endpoint
├── internal/env/     # Environment variable utilities
│   └── env.go        # Helper functions for env vars
├── bin/              # Compiled binaries (gitignored)
├── .env              # Environment variables
└── go.mod            # Go module dependencies
```

### Dependencies

- **Chi Router**: `github.com/go-chi/chi/v5` - HTTP router and middleware
- **GoDotEnv**: `github.com/joho/godotenv` - Environment file loading

### Configuration

The application uses environment variables for configuration:

- `ADDR`: Server address (default: `:8080`)

### Running the Application

1. **Install dependencies**:
   ```bash
   go mod tidy
   ```

2. **Set environment variables** (optional):
   ```bash
   # .env file is loaded automatically
   echo "ADDR=:3000" > .env
   ```

3. **Build and run**:
   ```bash
   go build -o bin/main ./cmd/api
   ./bin/main
   ```

4. **Test health endpoint**:
   ```bash
   curl http://localhost:8080/v1/health
   ```

### Next Steps

- [ ] Implement WebSocket connections for real-time messaging
- [ ] Add user authentication and session management
- [ ] Create chat room functionality
- [ ] Add message persistence (database integration)
- [ ] Implement message broadcasting
- [ ] Add frontend client interface

### Development Notes

- Server includes standard middleware: RequestID, RealIP, Logger, Recoverer, Timeout
- Follows Go project layout conventions
- Environment variables loaded with fallback defaults
- Proper error handling and logging implemented