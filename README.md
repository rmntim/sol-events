# Solana Events Monitor

A Go-based tool for monitoring and decoding events from Solana smart contracts. This tool continuously polls a specified Solana program and decodes any events emitted by it.

## Prerequisites

- Go 1.24 or later
- Access to a Solana RPC endpoint

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/sol-events.git
cd sol-events
```

2. Install dependencies:

```bash
go mod download
```

## Usage

Run the program with default settings:

```bash
task run
```

### Command Line Arguments

- `--program`: The Solana program ID to monitor (default: "6QSaTnP3biE27vP1gRFqmHr7hzpp3VRTdFnEonuMtgWw")
- `--interval`: Polling interval (default: 10s)
- `--endpoint`: Solana RPC endpoint (default: DevNet)

Example with custom parameters:

```bash
task build
sol-events --program YOUR_PROGRAM_ID --interval 5s --endpoint https://api.mainnet-beta.solana.com
```

## Dependencies

- [solana-go](https://github.com/gagliardetto/solana-go): Solana client library
- [solana-anchor-go](https://github.com/rmntim/solana-anchor-go): Anchor IDL support
- Other dependencies listed in `go.mod`

## License

This project is licensed under the MIT License - see the LICENSE file for details.
