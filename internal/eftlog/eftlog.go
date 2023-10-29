// Package tarkovlog provides a function for retrieving the IP address from the Tarkov log file.
package eftlog

import (
	"net"
	"path/filepath"
)

// LogPathDefault is the default log file path.
const LogPathDefault = "C:\\Battlestate Games\\EFT\\Logs"

// Config holds configuration options for log extraction.
type Log struct {
	Path string
}

type LogOption func(*Log)

func WithPath(path string) LogOption {
	return func(l *Log) {
		l.Path = filepath.FromSlash(path)
	}
}

// New creates a new configuration for log reading.
func New(opts ...LogOption) *Log {
	l := new(Log)

	for _, opt := range opts {
		opt(l)
	}

	if l.Path == "" {
		l.Path = LogPathDefault
	}

	return l
}

// ReadIP reads the IP address from the log file specified in the configuration.
func (c *Log) GetIP() (net.IP, error) {
	// Find the log file path.
	filePath, err := c.FindLog()
	if err != nil {
		return nil, err
	}

	// Extract the IP from the log file.
	return extractIP(filePath)
}
