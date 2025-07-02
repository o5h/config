# config

`config` is a Go package for managing configuration files in YAML format.

## Features

- Load configuration from YAML files.
- Validate configuration structure.
- Easy-to-use API for accessing configuration values.

## Installation

```bash
# Install the package using go get

go get -u github.com/o5h/config
```

## Usage

Here is an example of how to use the `config` package:

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/o5h/config"
)

func main() {
	// Load configuration from a YAML file
	err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// Access configuration values
	fmt.Println("Server Port:", config.Get("server.port", ""))
	fmt.Println("Database Host:", config.Get("database.host", ""))
}
```

## Configuration File Format

The configuration file should be in YAML format. Below is an example:

```yaml
server:
  port: 8080

database:
  host: localhost
  port: 5432
  username: user
  password: pass
```

## Testing

Run the tests using the following command:

```bash
go test ./...
```

## License

This project is licensed under the MIT License. See the LICENSE file for details.
