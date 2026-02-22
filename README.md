# LogCheck Linter
A linter for checking code for compliance with user rules.  
Supports loggers:
* log/slog
* go.uber.org/zap
## Demonstration
**main.go**
```go 
package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	token := "wegegwe!@$#wgaw#13412"

	log.Error("Русский язык") // only english rule
	slog.Warn("token" + token) // no sensitive data rule
	slog.Info(",,,weg$%🚀") // no special symbols rule
	log.Error("First letter is upper") // lowercase rule
}
```

Run the checking
```bash
./custom-gcl run ./...
```

**logs**
```bash
ivan-sustatov@w11:/mnt/d/Documents/Learn Go/test_linter$ ./custom-gcl run ./...
main.go:13:12: log message should contain only english characters (logrules)
        log.Error("Русский язык") // only english rule
                  ^
main.go:14:12: log message may contain sensitive data (logrules)
        slog.Warn("token" + token) // no sensitive data rule
                  ^
main.go:15:12: log message contains forbidden characters or emoji (logrules)
        slog.Info(",,,weg$%🚀") // no special symbols rule
                  ^
main.go:16:12: log message should start with lowercase letter (logrules)
        log.Error("First letter is upper") // lowercase rule
                  ^
4 issues:
* logrules: 4
```

Run the fixes
```bash
./custom-gcl run --fix ./...
```

**fixed_main.go**
```go
package main

import (
	"log/slog"

	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	token := "wegegwe!@$#wgaw#13412"

	log.Error("russian language")      // only english rule
	slog.Warn("token")                 // no sensitive data rule
	slog.Info("weg")                   // no special symbols rule
	log.Error("first letter is upper") // lowercase rule
}
```


## Installation & Run
### Prerequisites
1. [Go (1.26.0+)](https://go.dev/doc/install)
2. [Bash/WSL](https://learn.microsoft.com/ru-ru/windows/wsl/install) (Для Windows)
3. [golangci-lint (v2.10.1+)](https://golangci-lint.run/docs/welcome/install/)

### Quick Start (Development)
1. Create a test project:
   ```bash
   go mod init test_linter
   ```
2. Include the linter as a dependency:
   ```bash
   go get github.com/sustatov027-max/logcheck_linter
   ```
3. Build the plugin (requires CGO):
   ```bash
   CGO_ENABLED=1 go build -buildmode=plugin -o logrules.so github.com/sustatov027-max/logcheck_linter
   ```
4. Build a custom golangci-lint binary ([set up the configs](#configuration-setup)):
    ```bash
    golangci-lint custom -v
    ````
5. Run code check:
    ```bash
    ./custom-gcl run main.go
    ```
#### Alternative method (without CGO/plugins)
If you want to use the linter without building the plugin:

1. Build the executable file:
    ```bash
    go build -o logcheck ./cmd/logcheck
    ```
2. Run the check via go vet:
    ```bash
    go vet -vettool=./logcheck ./...
    ```
### Configuration setup
Create .custom-gcl.yml in the project root:
```yaml
version: v2.10.1
plugins:
  - module: your_module_name_in_go.mod(example: test_linter)
    path: ./
    import: github.com/sustatov027-max/logcheck_linter/analyzer
```

Create .golangci.yml in the project root:
```yaml
version: "2"
linters:
  default: none
  enable:
    - logrules
  settings:
    custom:
      logrules:
        type: goplugin
        path: ./logrules.so
        description: "Checks log messages style"
  exclusions:
    generated: lax
    presets:
      - comments
      - common-false-positives
      - legacy
      - std-error-handling
    paths:
      - third_party$
      - builtin$
      - examples$
formatters:
  exclusions:
    generated: lax
    paths:
      - third_party$
      - builtin$
      - examples$
```

### Checking work
Checking a file
```bash
./custom-gcl run main.go
```
Checking the entire project
```bash
./custom-gcl run ./...
```

With autocorrection
```bash
./custom-gcl run --fix main.go
```