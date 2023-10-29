package systray

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/ssapp/tarkov-server-checker/internal/ip_location"
)

// updateTooltip sets the system tray tooltip with IP, country, and last updated information.
func (t *Systray) updateTooltip() {
	// Define the format for the tooltip.
	format := "Server: %s (%s)"

	// Fetch the IP information.
	ip, err := t.log.GetIP()
	if err != nil {
		// If an error occurred, set the tooltip to "Error".
		errText := fmt.Sprintf("Error reading IP: %v", err)
		systray.SetTooltip(errText)
		return
	}

	// Fetch the country name for the IP.
	country, err := ip_location.Get(ip)
	if err != nil {
		// If an error occurred, set the tooltip to "Error".
		errText := fmt.Sprintf("Error fetching country: %v", err)
		systray.SetTooltip(errText)
		return
	}

	// Create the tooltip string with the current values.
	tooltip := fmt.Sprintf(format, ip, country)

	// Set the tooltip in the system tray.
	systray.SetTooltip(tooltip)
}
