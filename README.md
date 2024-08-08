# Marvin Blockchain - Go Implementation

[![Build Status](https://github.com/joaoh82/marvin-blockchain-go/workflows/Go/badge.svg)](https://github.com/joaoh82/marvin-blockchain-go/actions)
[![Coverage Status](https://coveralls.io/repos/github/joaoh82/marvin-blockchain-go/badge.svg?branch=main)](https://coveralls.io/github/joaoh82/marvin-blockchain-go?branch=main)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

Welcome to the Go implementation of the Marvin Blockchain. This project is part of a comparative study on building the same blockchain in both Go and Rust. Follow along as we explore the development process, performance, and features of a blockchain built in Go.

### Read the series of posts about it:
##### Crafting a Blockchain in Go and Rust: A Comparative Journey
* [Part 0 - Introduction and Overview](https://blog.thepolyglotprogrammer.com/)

## Project Overview

Marvin Blockchain is a distributed ledger inspired by Bitcoin, implemented in Go. This project aims to provide a robust and scalable blockchain solution while comparing the nuances of building the same system in Rust.

## Features (WIP)

- **Proof of Work (PoW) Consensus Mechanism**
- **Peer-to-Peer (P2P) Networking**
- **Transaction and Block Validation**
- **Smart Contract Support via EVM Compatibility**
- **Comprehensive Unit Tests and Benchmarks**

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

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
go build -o marvin-blockchain-go
```

### Running the Blockchain
To start a node on the Marvin Blockchain:
```sh
./marvin-blockchain-go
```

### Running Tests
To run the unit tests:
```sh
go test ./...
```

### Project Progress (WIP)
- [ ] Add CLI support for ease of interaction
- [ ] Implement the basic blockchain data structure
- [ ] Proof of Work (PoW) consensus mechanism
- [ ] Peer-to-Peer (P2P) networking setup
- [ ] Basic transaction and block validation
- [ ] EVM integration for smart contract support
- [ ] JSON-RPC API implementation
- [ ] Advanced transaction handling and validation
- [ ] Enhanced security measures and best practices
- [ ] Performance benchmarking and optimization
- [ ] Comprehensive documentation and examples

### Roadmap (Subject to Change)
- [ ] Features and improvements planned for the project:
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
`cmd/`: Contains the main entry point for the application and different binaries.
`docs/`: Contains project documentation and guides.
`core/`: Contains the core blockchain implementation and data structures.
`network/`: Contains the networking and peer-to-peer communication logic.
`crypto/`: Contains cryptographic utilities and security features.
`wallet/`: Contains wallet and key management functionalities.
`transactions/`: Contains transaction handling and validation logic.
`internal/`: Contains internal packages and modules.
`utils/`: Contains utility functions and helper methods.
`tests/`: Contains unit and integration tests.

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
