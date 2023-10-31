// Package systray provides a system tray application for Tarkov Server Checker.
package systray

import (
	"fmt"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/ssapp/tarkov-server-checker/internal/log"
	"github.com/ssapp/tarkov-server-checker/resources/icon"
)

// Systray represents the system tray application for Tarkov Server Checker.
type Systray struct {
	logReader logReader
	ticker    *time.Ticker
	mInfo     *systray.MenuItem
	mQuit     *systray.MenuItem
}

type logReader interface {
	GetIP() (string, error)
}

// NewSystray creates a new Systray instance with the specified log reader.
func NewSystray(logReader logReader) *Systray {
	return &Systray{
		logReader: logReader,
		ticker:    time.NewTicker(time.Second),
	}
}

// Run starts the system tray application.
func (s *Systray) Run() {
	systray.Run(s.onReady, s.onExit)
}

// onReady is called when the system tray is ready to be used.
func (s *Systray) onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTooltip("Tarkov Server Checker is starting...")

	s.mInfo = systray.AddMenuItem("Loading...", "Loading...")
	s.mQuit = systray.AddMenuItem("Quit", "Quit Tarkov Server Checker")

	s.mInfo.Disable()
	s.mQuit.Enable()

	go s.handleQuit()
	go s.startTickerLoop()
}

// onExit is called when the system tray application is exiting.
func (s *Systray) onExit() {
	s.ticker.Stop()
}

// handleQuit handles the quit menu item click.
func (s *Systray) handleQuit() {
	<-s.mQuit.ClickedCh
	systray.Quit()
}

// handleError handles the error and updates the state.
func (s *Systray) handleError(err error, msg string) {
	if err != nil {
		s.updateState(formatTooltip(msg, err.Error(), ""))
	}
}

// startTickerLoop starts the ticker loop.
func (s *Systray) startTickerLoop() {
	for range s.ticker.C {
		ip, err := s.logReader.GetIP()
		s.handleError(err, "Unable to get Server's IP address")
		if err != nil {
			continue
		}

		loc, err := log.GetLocation(ip)
		s.handleError(err, "Unable to get Server's Location")
		if err != nil {
			continue
		}

		s.updateState(formatTooltip(ip, loc))
	}
}

// updateState updates the menu item and tooltip with the specified info.
func (s *Systray) updateState(info string) {
	s.mInfo.SetTitle(info)
	systray.SetTooltip(info)
}

// formatTooltip formats the tooltip with the specified key, value, and info.
func formatTooltip(key, val string, info ...string) string {
	if len(info) != 0 {
		return fmt.Sprintf("%s - %s | %s", key, val, strings.Join(info, " "))
	}
	return fmt.Sprintf("%s - %s", key, val)
}
