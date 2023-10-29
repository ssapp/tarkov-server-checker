// Package location provides a function for retrieving the country name for a given IP address.
package ip_location

import (
	"encoding/json"
	"net"
	"net/http"
)

// Get retrieves the country name for a given IP address using an external service.
func Get(ip net.IP) (string, error) {
	// Construct the URL for fetching location information based on the IP.
	url := "https://api.iplocation.net/?ip=" + ip.String()

	// Fetch the location information.
	loc, err := fetchLocation(url)
	if err != nil {
		return "", err
	}

	// Return the country name from the location information.
	return loc.CountryName, nil
}

// loc represents the JSON structure for location information.
type loc struct {
	CountryName string `json:"country_name"`
}

// fetchLocation fetches location information from the specified URL.
func fetchLocation(url string) (*loc, error) {
	// Send an HTTP GET request to the URL.
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response body into the loc struct.
	var location *loc
	if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
		return nil, err
	}

	return location, nil
}
