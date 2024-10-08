# Mole

## Project Overview

P2P File Sharing Service is a simple yet powerful peer-to-peer file sharing server written in Go, using only the standard library. The goal of the project is to create a distributed file sharing system where each node can act as both a client and a server, allowing users to securely and efficiently exchange files within a local network or over the Internet without relying on a central server.

## Project Idea

The main idea of the project is to create an easily scalable and secure P2P network that enables users to:
- Automatically discover other nodes in the local network using UDP broadcast.
- Directly exchange files over TCP connections, providing fast and easy access to data.
- Act as both a client (to request and send files) and a server (to receive files).

### Why Go?

Go was chosen as the development language due to its built-in support for networking operations, simplicity in handling concurrency (goroutines), and ease of building high-performance network applications. Using Go’s standard library ensures minimal dependencies and a straightforward installation process.

## Features

### Features of the near future

- **TCP server for receiving files**: Each node can listen for incoming connections and save transmitted files.
- **TCP client for sending files**: A node can connect to another node and send files directly.
- **Node discovery via UDP broadcast**: Nodes can automatically discover each other within the local network and exchange information.
- **Ease of use**: Minimal setup required, making it easy to deploy a node on a local network or over the Internet.

### Future Vision

The project has the potential to evolve into a more advanced and functional P2P service. Below are some directions we plan to take:

1. **Enhanced Node Discovery**:
   - Discover nodes not only within the local network but also over the Internet using DHT (Distributed Hash Table) and/or centralized trackers.
   - Support for IP and DNS-based connections so nodes can locate each other outside the local network.

2. **Improved Messaging Protocol**:
   - Introducing JSON or another format for exchanging information about requests, files, and connection statuses to ensure better compatibility and extendability of the protocol.
   - Support for transmitting file metadata (size, type, checksum) so nodes can verify data integrity.

3. **Encryption and Security**:
   - Implementing TLS for encrypting TCP connections to secure data during transmission.
   - Node authentication to prevent unauthorized access and network attacks.

4. **Data Recovery and Resiliency**:
   - Mechanisms for recovering file transfers in case of network errors or connection drops.
   - Tracking transfer progress to resume interrupted downloads from where they left off.

5. **Group File Transfers and Directory Support**:
   - Support for sending multiple files simultaneously and/or transferring entire directories.
   - The ability to create groups of nodes to share files with multiple peers at once.

6. **User Interface**:
   - Building a simple CLI interface to manage the node (request files, list available nodes and files).
   - Considering a web-based interface for node management and monitoring in the future.

## Installation and Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/Toolnado/mole.git

2. Build and run the server: 
    ``` bash
    make run

## Contributing

The project is open for contributions and suggestions. If you have ideas, bug reports, or enhancement proposals, feel free to create an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
