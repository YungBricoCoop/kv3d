/*
copyright © 2025 elwan mayencourt <mayencourt@elwan.ch>
based on resp protocol: https://redis.io/docs/latest/develop/reference/protocol-spec/
*/
package resp

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const crlf = "\r\n"

func ReadArray(reader *bufio.Reader) ([]string, error) {
	// read first line: *<count>\r\n
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// remove trailing \r\n
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")

	// arrays must start with *
	if !strings.HasPrefix(line, "*") {
		return nil, fmt.Errorf("invalid resp array: expected '*', got '%s'", line)
	}

	// parse number of elements
	count, err := strconv.Atoi(line[1:])
	if err != nil {
		return nil, fmt.Errorf("invalid array count: %w", err)
	}
	if count < 0 {
		return nil, fmt.Errorf("invalid array count: %d", count)
	}

	result := make([]string, count)

	for i := 0; i < count; i++ {
		element, err := readBulkString(reader)
		if err != nil {
			return nil, fmt.Errorf("failed to read element %d: %w", i, err)
		}
		result[i] = element
	}

	return result, nil
}

func readBulkString(reader *bufio.Reader) (string, error) {
	// read bulk string header: $<len>\r\n
	lengthLine, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	lengthLine = strings.TrimSuffix(lengthLine, "\n")
	lengthLine = strings.TrimSuffix(lengthLine, "\r")

	if !strings.HasPrefix(lengthLine, "$") {
		return "", fmt.Errorf("invalid bulk string: expected '$', got '%s'", lengthLine)
	}

	length, err := strconv.Atoi(lengthLine[1:])
	if err != nil {
		return "", fmt.Errorf("invalid bulk string length: %w", err)
	}

	// handle null string
	if length == -1 {
		return "", nil
	}
	if length < -1 {
		return "", fmt.Errorf("invalid bulk string length: %d", length)
	}

	// read string content (exactly 'length' bytes)
	data := make([]byte, length)
	if _, err := io.ReadFull(reader, data); err != nil {
		return "", err
	}

	// read and validate trailing \r\n
	trailer := make([]byte, 2)
	if _, err := io.ReadFull(reader, trailer); err != nil {
		return "", err
	}
	if string(trailer) != crlf {
		return "", fmt.Errorf("invalid bulk string terminator: expected %q", crlf)
	}

	return string(data), nil
}

func EncodeSimpleString(s string) string {
	return fmt.Sprintf("+%s%s", s, crlf)
}

func EncodeBulkString(s string) string {
	return fmt.Sprintf("$%d%s%s%s", len(s), crlf, s, crlf)
}

func EncodeInteger(i int) string {
	return fmt.Sprintf(":%d%s", i, crlf)
}

func EncodeNull() string {
	return fmt.Sprintf("$-1%s", crlf)
}

func EncodeError(msg string) string {
	return fmt.Sprintf("-%s%s", msg, crlf)
}
