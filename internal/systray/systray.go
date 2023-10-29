// Package tray provides a system tray application for Tarkov Server Checker.
package systray

import (
	"time"

	"github.com/getlantern/systray"
	"github.com/ssapp/tarkov-server-checker/icon"
	"github.com/ssapp/tarkov-server-checker/internal/eftlog"
)

const (
	// UpdateInterval is the interval at which the IP address is checked.
	updateInterval = 1 * time.Second
)

// Tray represents the system tray application for Tarkov Server Checker.
type Systray struct {
	log    *eftlog.Log
	ticker *time.Ticker
}

// New creates a new Tray instance with the specified log reader.
func New(log *eftlog.Log) *Systray {
	return &Systray{
		log:    log,
		ticker: time.NewTicker(updateInterval),
	}
}

// Run starts the system tray application.
func (t *Systray) Run() {
	systray.Run(t.onReady, nil)
}

// onReady is called when the system tray is ready to be used.
func (t *Systray) onReady() {
	// Set the title, tooltip, and icon.
	systray.SetTitle("Tarkov Server Checker")
	systray.SetIcon(icon.Data)
	systray.SetTooltip("Tarkov Server Checker is starting...")
	// Add a "Quit" menu item.
	mQuit := systray.AddMenuItem("Quit", "Quit Tarkov Server Checker")
	// Continuously update the IP and tooltip every minute.
	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			case <-t.ticker.C:
				t.updateTooltip()
			}
		}
	}()
}
