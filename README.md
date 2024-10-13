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
