// Package main is the entry point for the Tarkov Server Checker application.
package main

import (
	"github.com/ssapp/tarkov-server-checker/internal/log"
	"github.com/ssapp/tarkov-server-checker/internal/systray"
)

//go:generate goversioninfo -icon=..\..\resources\icon\icon.ico
func main() {
	logConfig := log.Config{}
	log := log.NewLog(logConfig)
	systray := systray.NewSystray(log)
	systray.Run()
}
