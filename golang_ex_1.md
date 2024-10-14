## Objective
Create a basic cross-platform command execution application that can be installed on Windows and macOS, capable of executing simple system commands and returning their results.

## Core Requirements (2 hours)

### 1. Installation Package (~30 minutes)
- Create a simple installer for either Windows (.exe) or macOS (.pkg) [choose one]
- Implement basic start-on-boot functionality

### 2. Command Execution (~45 minutes)
Implement support for two command types:
- Network ping
- Get system info (hostname, IP address)

```go
type Commander interface {
    Ping(host string) (PingResult, error)
    GetSystemInfo() (SystemInfo, error)
}

type PingResult struct {
    Successful bool
    Time       time.Duration
}

type SystemInfo struct {
    Hostname  string
    IPAddress string
}
```

### 3. Communication (~45 minutes)
- Implement a basic HTTP server to receive commands
- Return results as JSON

Example endpoint structure:
```go
// POST /execute
type CommandRequest struct {
    Type    string `json:"type"`    // "ping" or "sysinfo"
    Payload string `json:"payload"` // For ping, this is the host
}

type CommandResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data"`
    Error   string      `json:"error,omitempty"`
}
```

## Implementation Requirements

1. Main Application Structure
```go
func main() {
    commander := NewCommander()
    server := &http.Server{
        Addr:    ":8080",
        Handler: handleRequests(commander),
    }
    log.Fatal(server.ListenAndServe())
}

func handleRequests(cmdr Commander) http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/execute", handleCommand(cmdr))
    return mux
}

func handleCommand(cmdr Commander) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Parse request and execute command
    }
}
```

2. Platform-Specific Implementation
```go
type commander struct{}

func NewCommander() Commander {
    return &commander{}
}

func (c *commander) GetSystemInfo() (SystemInfo, error) {
    hostname, err := os.Hostname()
    if err != nil {
        return SystemInfo{}, err
    }
    
    // Get IP address (implement this)
    
    return SystemInfo{
        Hostname:  hostname,
        IPAddress: "implement me",
    }, nil
}
```

## Deliverables
1. Source code with:
   - Command execution implementation
   - HTTP server implementation
   - Basic installer script

2. README.md with:
   - Build instructions
   - API documentation
   - Installation guide
   - Testing
   - A short clip that shows the "app in action"

## Example Test Case
```go
func TestGetSystemInfo(t *testing.T) {
    cmdr := NewCommander()
    info, err := cmdr.GetSystemInfo()
    
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }
    
    if info.Hostname == "" {
        t.Error("Expected hostname to be non-empty")
    }
    
    if info.IPAddress == "" {
        t.Error("Expected IP address to be non-empty")
    }
}
```

## Time Breakdown
- Setup project and installer script: 30 minutes
- Implement command execution: 45 minutes
- Create HTTP server and endpoints: 45 minutes

## Evaluation Criteria
1. Code Quality
   - Clean, idiomatic Go code
   - Basic error handling
2. Functionality
   - Successfully executes required commands
   - Correctly returns results via HTTP
3. Completeness
   - All core requirements implemented
   - Basic documentation provided

## Notes for Candidates
- Focus on core functionality first
- Clean code is important, but perfect is the enemy of done
- Document any assumptions or limitations