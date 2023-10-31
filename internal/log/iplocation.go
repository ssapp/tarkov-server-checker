package log

import (
    "encoding/json"
    "fmt"
    "net"
    "net/http"
    "strings"
    "sync"
)

const (
    // API endpoint for IP location information.
    apiEndpoint = "https://freeipapi.com/api/json"
)

// locationCache holds the cached location information.
type locationCache struct {
    IP        string `json:"ipAddress,omitempty"`
    Country   string `json:"countryName,omitempty"`
    Continent string `json:"continent,omitempty"`
    sync.RWMutex
}

// Location holds the location information.
var Cache = &locationCache{}

// GetLocation retrieves the country information based on the IP address.
func GetLocation(ip string) (string, error) {
    // Check if the IP address is cached.
    Cache.RLock()
    if Cache.IP == ip {
        defer Cache.RUnlock()
        return Cache.String(), nil
    }
    Cache.RUnlock()

    // Check if the IP address is empty or invalid.
    if ip == "" || net.ParseIP(ip) == nil {
        return "", ErrInvalidIP
    }

    // Retrieve the location information.
    loc, err := fetchLocation(ip)
    if err != nil {
        return "", err
    }

    // Cache the location information.
    loc.Store()

    return Cache.String(), nil
}

func (l *locationCache) Store() {
    l.Lock()
    defer l.Unlock()
    Cache.IP = l.IP
    Cache.Country = l.Country
    Cache.Continent = l.Continent
}

func (l *locationCache) String() string {
    if l.Continent == "" && l.Country == "" {
        return "-/-"
    }
    return fmt.Sprintf("%s/%s", l.Continent, l.Country)
}

func fetchLocation(ip string) (*locationCache, error) {
    url := fmt.Sprintf("%s/%s", apiEndpoint, strings.TrimSpace(ip))
    resp, err := http.Get(url)
    if err != nil {
        return nil, ErrUnableToGetLocation
    }
    defer resp.Body.Close()

    var loc locationCache
    if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
        return nil, ErrUnableToSerializeLocation
    }

    if loc.IP != ip {
        return nil, fmt.Errorf("IP address mismatch: %s != %s", loc.IP, ip)
    }

    return &loc, nil
}