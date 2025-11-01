# KVD - Key Value Docker

```
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""\___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
         \    \        __/
          \____\______/

         _  ____   _____
        | |/ /\ \ / /   \
        | ' <  \ V /| |) |
        |_|\_\  \_/ |___/
```

> A dumb Redis-compatible key-value database using Docker containers as storage

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.23.2+-00ADD8?logo=go)](https://go.dev/)
[![Build Status](https://github.com/YungBricoCoop/kvd/workflows/build-and-release/badge.svg)](https://github.com/YungBricoCoop/kvd/actions)

## 🤔 What is this?

**KVD** is a inefficient but fun key-value database that uses Docker containers as its storage engine. It uses the RESP protocol, making it compatible with Redis clients and commands.

**Why "dumb"?** Because using Docker containers as a key-value store is one of the worst possible ways to implement a database. It's slow, resource-intensive, but fun.

### How it works

- **SET**: Creates a Docker container with the key as the container name and value stored as a label
- **GET**: Retrieves the value from the container label matching the key
- **DEL**: Removes the Docker container associated with the key

## 📋 Requirements

- **Docker** must be installed and running on your system

## 🚀 Installation

### macOS (Homebrew)

```bash
brew tap YungBricoCoop/tap
brew install kvd
```

### Linux (Debian/Ubuntu)

```bash
wget https://github.com/YungBricoCoop/kvd/releases/latest/download/kvd_Linux_x86_64.deb
sudo dpkg -i kvd_Linux_x86_64.deb
```

### Linux (RedHat/Fedora/CentOS)

```bash
wget https://github.com/YungBricoCoop/kvd/releases/latest/download/kvd_Linux_x86_64.rpm
sudo rpm -i kvd_Linux_x86_64.rpm
```

### Windows

```powershell
Invoke-WebRequest -Uri "https://github.com/YungBricoCoop/kvd/releases/latest/download/kvd_Windows_x86_64.zip" -OutFile "kvd.zip"
Expand-Archive -Path "kvd.zip" -DestinationPath "."
```

### Manual Binary Download

Download the latest binary for your platform from the [releases page](https://github.com/YungBricoCoop/kvd/releases).

Supported platforms:

- **Linux**: amd64, arm64 (`.deb`, `.rpm`, `.apk` packages available)
- **macOS**: amd64 (Intel), arm64 (Apple Silicon)
- **Windows**: amd64, arm64

### Build from Source

```bash
git clone https://github.com/YungBricoCoop/kvd.git
cd kvd
go build -o kvd
```

## 💻 Usage

### Start the Server

```bash
kvd serve --port 6379
```

Or use the default port (6379):

```bash
kvd serve
```

### Connect with Redis Client

Once the server is running, you can connect with any Redis client:

```bash
redis-cli -p 6379
```

### Supported Commands

```redis
127.0.0.1:6379> SET mykey "Hello, World!"
OK
127.0.0.1:6379> GET mykey
"Hello, World!"
127.0.0.1:6379> DEL mykey
(integer) 1
127.0.0.1:6379> GET mykey
(nil)
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
