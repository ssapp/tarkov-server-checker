package log

import (
	"bufio"
	"net"
	"os"
	"regexp"
	"strings"
)

// ExtractIP extracts the IP address from the log file.
func extractIPFromLog(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	line, err := getLatestMatchingLine(f, "RaidMode: Online,")
	if err != nil {
		return "", err
	}

	return extractIPFromLine(line)
}

// getLatestMatchingLine scans the log file for the latest line containing the target string.
func getLatestMatchingLine(f *os.File, target string) (string, error) {
	scanner := bufio.NewScanner(f)
	var latestMatchingLine string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, target) {
			latestMatchingLine = line
		}
	}

	if latestMatchingLine == "" {
		return "", ErrNoMatchingLinesFound
	}

	return latestMatchingLine, nil
}

// extractIPFromLine extracts the IP address from a line using regular expressions.
func extractIPFromLine(line string) (string, error) {
	re := regexp.MustCompile(`Ip: ([\d\.]+),`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 2 {
		return "", ErrUnableToExtractIPFromLine
	}

	ip := matches[1]

	if !isValidIP(ip) {
		return "", ErrInvalidIP
	}

	return ip, nil
}

// isValidIP checks if a given string is a valid IP address.
func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
