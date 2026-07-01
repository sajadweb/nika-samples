# Nika
## Project Overview
Nika is a small web service scaffold written in Go, intended as a lightweight starter for building APIs and microservices.

## Framework
Nika uses the Gin web framework for routing and middleware and relies on Go modules for dependency management.

## Requirements
- Go 1.20 or newer
- Git

## Local Installation & Run
1. Clone the repository:
    `git clone <repository-url>`
2. Enter the project directory:
    `cd nika-app`
3. Download dependencies:
    `go mod download`
4. Run for production:  
```bash
go install github.com/sajadweb/nika-cli@latest
nika swagger init
nika start .
```
5. Run for development: 
```bash
go install github.com/sajadweb/nika-cli@latest
nika swagger init
nika start . --watch
```
 


## Environment Configuration
Configuration is read from environment variables. Common examples:
- `PORT` — service port (default: 3007)


## Tests
Run tests with:
`go test ./...`

## Quick Usage
Start example:
```bash
http://localhost:3007/swagger/index.html
```



## License
See the LICENSE file for licensing details.