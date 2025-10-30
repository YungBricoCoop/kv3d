/*
Copyright © 2025 Elwan Mayencourt <mayencourt@elwan.ch>
Based on RESP protocol: https://redis.io/docs/latest/develop/reference/protocol-spec/
*/
package resp

import (
	"fmt"
	"kvd/internal/docker"
	"strings"
)

func ProcessCommand(command []string) string {
	if len(command) == 0 {
		return EncodeError("ERR empty command")
	}

	cmd := strings.ToUpper(command[0])

	switch cmd {
	case "PING":
		return handlePing(command)
	case "GET":
		return handleGet(command)
	case "SET":
		return handleSet(command)
	case "DEL":
		return handleDel(command)
	case "QUIT":
		return EncodeSimpleString("OK")
	default:
		return EncodeError(fmt.Sprintf("ERR unknown command '%s'", cmd))
	}
}

func handlePing(command []string) string {
	if len(command) == 1 {
		return EncodeSimpleString("PONG")
	}
	if len(command) == 2 {
		return EncodeBulkString(command[1])
	}
	return EncodeError("ERR wrong number of arguments for 'ping' command")
}

func handleGet(command []string) string {
	if len(command) != 2 {
		return EncodeError("ERR wrong number of arguments for 'get' command")
	}

	key := command[1]
	value, err := docker.GetContainerLabelValue(key)
	if err != nil {
		return EncodeNull()
	}

	return EncodeBulkString(value)
}

func handleSet(command []string) string {
	if len(command) < 3 {
		return EncodeError("ERR wrong number of arguments for 'set' command")
	}

	key := command[1]
	value := command[2]

	err := docker.RunContainer(key, value, 1)
	if err != nil {
		return EncodeError(fmt.Sprintf("ERR %v", err))
	}

	return EncodeSimpleString("OK")
}

func handleDel(command []string) string {
	if len(command) != 2 {
		return EncodeError("ERR wrong number of arguments for 'del' command")
	}

	key := command[1]
	err := docker.DeleteContainer(key, 1)
	if err != nil {
		return EncodeInteger(0)
	}

	return EncodeInteger(1)
}
