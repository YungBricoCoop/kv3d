/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>
*/
package cmd

import (
	"bufio"
	"fmt"
	"kvd/internal/docker"
	"kvd/internal/resp"
	"log"
	"net"
	"time"

	"github.com/spf13/cobra"
)

var port int
var pruneInterval int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a Redis-compatible server using Docker as the storage engine",
	Long: `Start a Redis-compatible server that accepts Redis protocol commands (GET, SET, DEL).
kvd uses Docker containers as the underlying key-value storage engine.

The server listens for Redis client connections and translates Redis commands into
Docker container operations:
  - SET: Creates a container with the key as name and value as label
  - GET: Retrieves value from the container label matching the key
  - DEL: Removes the container associated with the key

Examples:
  kvd serve
	kvd serve --port 6379
	kvd serve --prune-interval 600

Connect with any Redis client:
  redis-cli -p 6379
  > SET mykey myvalue
  > GET mykey
  > DEL mykey`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 6379, "Port to listen on")
	serveCmd.Flags().IntVarP(&pruneInterval, "prune-interval", "i", 30, "Interval in seconds between container pruning runs")
}

func startServer(port int) {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer func() {
		if err := listener.Close(); err != nil {
			log.Printf("Error closing listener: %v", err)
		}
	}()

	docker.StartPruner(time.Duration(pruneInterval) * time.Second)

	log.Printf("KVD server listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()
	log.Printf("New connection from %s", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		command, err := resp.ReadArray(reader)
		if err != nil {
			log.Printf("Connection closed: %v", err)
			return
		}

		if len(command) == 0 {
			continue
		}

		response := resp.ProcessCommand(command)
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Printf("Failed to write response: %v", err)
			return
		}
	}
}
