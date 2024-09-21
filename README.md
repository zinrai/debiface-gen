# debiface-gen

debiface-gen is a Debian Network Interface Configuration Generator. It provides a command-line interface and an HTTP API for generating network interface configurations compatible with Debian's interfaces(5) format.

## Features

- Generate configurations for:
  - Bonding interfaces
  - DSR (Direct Server Return) interfaces
  - Standard network interfaces
- Command-line interface
- HTTP API

## Installation

Build the project:

```bash
$ go build
```

## Usage

### Command-line Interface

debiface-gen provides three subcommands: `bonding`, `dsr`, and `standard`.

#### Bonding Configuration

```bash
debiface-gen bonding \
  -auto \
  -iface bond0 \
  -ip 192.168.2.10 \
  -netmask 255.255.255.0 \
  -gateway 192.168.2.254 \
  -bond-master eth0 \
  -bond-slaves "eth0 eth1" \
  -bond-miimon 100 \
  -bond-mode active-backup
```

#### DSR Configuration

```bash
debiface-gen dsr \
  -auto \
  -iface dsr0 \
  -ip 10.0.0.1
```

#### Standard Interface Configuration

```bash
debiface-gen standard \
  -auto \
  -iface eth0 \
  -ip 192.168.2.10 \
  -netmask 255.255.255.0 \
  -gateway 192.168.2.1
```

### HTTP API

To start the HTTP server:

```bash
debiface-gen -server
```

The server will start on port 8080 by default.

#### Bonding Configuration

```bash
curl -X POST http://localhost:8080/api/bonding \
  -H "Content-Type: application/json" \
  -d '{
    "AutoIfaceUp": true,
    "Iface": "bond0",
    "IP": "192.168.0.1",
    "Netmask": "255.255.255.0",
    "Gateway": "192.168.0.254",
    "BondMaster": "eth0",
    "BondSlaves": ["eth0", "eth1"],
    "BondMiimon": 100,
    "BondMode": "active-backup"
  }'
```

#### DSR Configuration

```bash
curl -X POST http://localhost:8080/api/dsr \
  -H "Content-Type: application/json" \
  -d '{
    "AutoIfaceUp": true,
    "Iface": "dsr0",
    "IP": "10.0.0.1"
  }'
```

#### Standard Interface Configuration

```bash
curl -X POST http://localhost:8080/api/standard \
  -H "Content-Type: application/json" \
  -d '{
    "AutoIfaceUp": true,
    "Iface": "eth0",
    "IP": "192.168.1.10",
    "Netmask": "255.255.255.0",
    "Gateway": "192.168.1.1"
  }'
```

## Project Structure and Clean Architecture

This project follows the principles of Clean Architecture to ensure separation of concerns and maintainability. Here's an overview of the directory structure and its alignment with Clean Architecture:

```
debiface-gen/
├── main.go           # Entry point of the application
├── config/           # Core business logic and entities
│   ├── types.go      # Defines the core data structures
│   └── generator.go  # Contains the core logic for generating configurations
├── cli/              # Command-line interface adapter
│   └── cli.go        # Handles CLI interactions
├── api/              # HTTP API adapter
│   └── handlers.go   # Handles HTTP requests and responses
└── go.mod            # Go module definition
```

### Clean Architecture Layers

1. **Entities (config/types.go)**:
   - Contains the core data structures (BondingConfig, DSRConfig, StandardConfig).
   - These are independent of any framework or external agency.

2. **Use Cases (config/generator.go)**:
   - Implements the core business logic for generating configurations.
   - Depends on the entities but not on any external frameworks.

3. **Interface Adapters (cli/ and api/)**:
   - CLI adapter: Handles command-line interactions.
   - API adapter: Manages HTTP requests and responses.
   - These adapt the core logic to external interfaces (CLI and HTTP).

4. **Frameworks and Drivers (main.go)**:
   - The entry point of the application.
   - Connects all the layers and starts the application.

This structure allows for:
- Independence of frameworks: The core logic doesn't depend on CLI or HTTP.
- Testability: Each layer can be tested independently.
- Independence of UI: The same core logic serves both CLI and HTTP interfaces.
- Independence of Database: In this case, no database is used, but if needed, it could be easily added without affecting the core logic.

## License

This project is licensed under the MIT License - see the [LICENSE](https://opensource.org/license/mit) for details.
