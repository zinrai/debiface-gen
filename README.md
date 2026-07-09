# debiface-gen

`debiface-gen` is a Debian Network Interface Configuration Generator. It provides a command-line interface and an HTTP API for generating network interface configurations compatible with Debian's interfaces(5) format.

## Features

- Generate configurations for:
  - Bonding interfaces
  - DSR (Direct Server Return) interfaces
  - Standard network interfaces
  - Bridge interfaces
- Command-line interface
- HTTP API

## Installation

Build the project:

```bash
$ go build
```

## Usage

### Command-line Interface

debiface-gen provides four subcommands: `bonding`, `dsr`, `standard`, and `bridge`.

#### Bonding Configuration

```bash
debiface-gen bonding \
  -auto \
  -iface bond0 \
  -ip 192.168.2.10 \
  -netmask 255.255.255.0 \
  -gateway 192.168.2.254 \
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

The DSR stanza only creates the dummy interface and assigns the /32 address. The `arp_ignore` / `arp_announce` sysctls that a working DSR setup needs are out of scope and must be managed separately.

#### Standard Interface Configuration

```bash
debiface-gen standard \
  -auto \
  -iface eth0 \
  -ip 192.168.2.10 \
  -netmask 255.255.255.0 \
  -gateway 192.168.2.1
```

#### Bridge Configuration

```bash
debiface-gen bridge \
  -auto \
  -iface br0 \
  -bridge-ports "eth0 eth1"
```

### HTTP API

To start the HTTP server:

```bash
debiface-gen -server
```

The server will start on port 8080 by default. Each endpoint validates the request and returns the generated stanza as JSON (`{"config": "..."}`) for a client to parse and apply. Invalid input returns 400 with the reason.

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

#### Bridge Configuration

```bash
curl -X POST http://localhost:8080/api/bridge \
  -H "Content-Type: application/json" \
  -d '{
    "AutoIfaceUp": true,
    "Iface": "br0",
    "BridgePorts": ["eth0", "eth1"]
  }'
```

## License

This project is licensed under the [MIT License](./LICENSE).
