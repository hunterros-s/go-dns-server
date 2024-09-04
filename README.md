# Go DNS Server

Welcome to the Go DNS Server project! This repository houses a robust, performant DNS server written in Go. Designed to be lightweight and extensible (as well as for me to become more familiar with Go and other complex backend and networking concepts).

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Code Overview](#code-overview)

## Introduction

The Go DNS Server is a highly efficient and scalable DNS server implementation built to handle DNS requests and responses seamlessly. This project exemplifies advanced Go programming capabilities and demonstrates proficiency in networking, concurrency, and system design. It is meticulously structured and well-commented to serve as an educational tool and a starting point for further development.

## Features

- **High Performance**: Utilizes Go’s capabilities for handling concurrent network communications.
- **Modularity**: Well-organized codebase with clear separation of concerns across various packages.
- **Standard Compliance**: Adheres to DNS protocol standards ensuring compatibility with existing DNS infrastructure.
- **Record Types**: Supports various DNS record types such as A, NS, CNAME, MX, AAAA, and many more.
- **Extensible Architecture**: Easily extendable to support additional DNS features and record types.
- **Comprehensive Logging**: Detailed logging to aid in debugging and performance monitoring.
- **Error Handling**: Robust error handling mechanisms to ensure smooth operation.

## Installation

To get started with the Go DNS Server, ensure you have Go installed on your system. Follow these steps to install and run the server:

1. Clone the repository:
   ```sh
   git clone https://github.com/your-username/go-dns-server.git
   cd go-dns-server
   ```

2. Build the project:
   ```sh
   go build ./cmd/dns
   ```

3. Run the DNS server:
   ```sh
   ./dns
   ```

## Usage

After starting the server, it will listen on port 2053 for incoming DNS queries. To test the server, you can use a DNS client or tools like `dig`:

```sh
dig @127.0.0.1 -p 2053 example.com
```

The server will process the query and respond appropriately based on the implemented functionality.

## Code Overview

The project is divided into multiple packages, each handling specific functionalities:

- **`dns`**: Core DNS functionalities including interfaces and types.
- **`dns/packet`**: Packet structure and parsing logic for DNS messages.
- **`dns/udp`**: UDP socket implementation for handling DNS over UDP.
- **`dns/parser`**: Parses incoming DNS messages to extract useful information.
- **`dns/registry`**: Handling various DNS record types.
- **`dns/factory`**: Generators for DNS questions and records.
- **`dns/buffer`**: Buffer management for reading and writing DNS packets.
- **`cmd/dns`**: Contains the entry point of the application.

---