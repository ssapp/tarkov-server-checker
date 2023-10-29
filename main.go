//go:build windows

// Package main is the entry point for the Tarkov Server Checker application.
package main

import "github.com/ssapp/tarkov-server-checker/cmd"

func main() {
	cmd.Execute()
}
