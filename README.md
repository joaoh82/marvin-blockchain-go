# Marvin Blockchain - Go Implementation

[![Build Status](https://github.com/joaoh82/marvin-blockchain-go/workflows/Go/badge.svg)](https://github.com/joaoh82/marvin-blockchain-go/actions)
[![Coverage Status](https://coveralls.io/repos/github/joaoh82/marvin-blockchain-go/badge.svg?branch=main)](https://coveralls.io/github/joaoh82/marvin-blockchain-go?branch=main)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

Welcome to the Go implementation of the Marvin Blockchain. This project is part of a comparative study on building the same blockchain in both Go and Rust. Follow along as we explore the development process, performance, and features of a blockchain built in Go.

### Read the series of posts about it:
##### Crafting a Blockchain in Go and Rust: A Comparative Journey
[Crafting a Blockchain in Go and Rust: A Comparative Journey - Series](https://hashblog.thepolyglotprogrammer.com/series/crafting-a-blockchain-in-go-and-rust-a-comparative-journey)

## Project Overview

Marvin Blockchain is a distributed ledger and EVM Compatible inspired by Bitcoin and Ethereum, implemented in Go. This project aims to provide a robust and scalable blockchain solution while comparing the nuances of building the same system in Rust.

## Features (WIP)

- **Proof of Work (PoW) Consensus Mechanism** (Subject to Change)
- **Peer-to-Peer (P2P) Networking**
- **Storage and Persistence**
- **Transaction and Block Validation**
- **Smart Contract Support via EVM Compatibility**
- **JSON-RPC API**
- **Comprehensive Unit Tests and Benchmarks**

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git
- protobuf
```sh
# Install protoc (if not already installed)
brew install protobuf

# Install the Go plugin for protoc
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

### Installation

1. Clone the repository:
```sh
git clone https://github.com/joaoh82/marvin-blockchain-go.git
cd marvin-blockchain-go
```

2. Install dependencies:
```sh
go mod tidy
```

3. Build the project:
```sh
make build
```

4. Build the cli:
```sh
make build-cli
```

### Running the Blockchain
To start a node on the Marvin Blockchain:
```sh
./bin/marvin
```

### Running the CLI
To the CLI application to interact with the blockchain:
```sh
./bin/marvinctl --help
```

### Running Tests
To run the unit tests:
```sh
make test
```

### Project Progress (WIP)
- [x] Add CLI support for ease of interaction
- [x] Implemented key pair creation, sign and verify
- [x] Create key pair with mnemonic seed
- [x] Add address command to CLI
- [x] Basic transaction and block data structures and validation
- [x] Add better error handling and logging
- [x] Add protobuf enconding/decoding
- [x] Implement the basic blockchain data structure
- [ ] Peer-to-Peer (P2P) networking implementation (transport layer)

### Roadmap (Subject to Change)
- [ ] Proof of Work (PoW) consensus mechanism
- [ ] Storage and persistence for blockchain data
- [ ] EVM integration for smart contract support
- [ ] JSON-RPC API implementation
- [ ] Advanced transaction handling and validation
- [ ] Enhanced security measures and best practices
- [ ] Performance benchmarking and optimization
- [ ] Comprehensive documentation and examples
- [ ] Implement wallet functionalities
- [ ] Improve EVM compatibility and support
- [ ] Add more consensus mechanisms (e.g., PoS)
- [ ] Implement light client support
- [ ] Improve network protocol for better scalability
- [ ] Develop a robust test suite for security and performance
- [ ] Integration with Ethereum development tools
- [ ] Develop a block explorer
- [ ] Implement governance mechanisms
- [ ] Cross-chain interoperability solutions
- [ ] Improve documentation and developer guides

### Project Structure (WIP)
The project is structured as follows:
- `cmd/`: Contains the main entry point for the application and different binaries.
- `docs/`: Contains project documentation and guides.
- `core/`: Contains the core blockchain implementation and data structures.
- `network/`: Contains the networking and peer-to-peer communication logic.
- `crypto/`: Contains cryptographic utilities and security features.
- `wallet/`: Contains wallet and key management functionalities.
- `transactions/`: Contains transaction handling and validation logic.
- `internal/`: Contains internal packages and modules.
- `utils/`: Contains utility functions and helper methods.
- `tests/`: Contains unit and integration tests.

### Contributing
**Pull requests are warmly welcome!!!**

For major changes, please [open an issue](https://github.com/joaoh82/marvin-blockchain-go/issues/new) first and let's talk about it. We are all ears!

If you'd like to contribute, please fork the repository and make changes as you'd like and shoot a Pull Request our way!

**Please make sure to update tests as appropriate.**

If you feel like you need it go check the GitHub documentation on [creating a pull request](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

### Code of Conduct

Contribution to the project is organized under the terms of the
Contributor Covenant, the maintainer of Marvin Blockchain, [@joaoh82](https://github.com/joaoh82), promises to
intervene to uphold that code of conduct.

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

### Contact
For any inquiries or support, please open an issue on Github or contact me at Joao Henrique Machado Silva <joaoh82@gmail.com>.
