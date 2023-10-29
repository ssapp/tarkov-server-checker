package eftlog

import (
	"bufio"
	"errors"
	"net"
	"os"
	"regexp"
	"strings"
)

var (
	ErrorNoMatchingLine = errors.New("no matching line found in log file")
	ErrorNoLines        = errors.New("no lines to select from")
	ErrorNoIP           = errors.New("could not extract IP from line")
)

// ExtractIP extracts the IP from a log file.
func extractIP(filePath string) (net.IP, error) {
	// Open the log file.
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	// Close the file when done.
	defer file.Close()

	// Find lines containing "RaidMode: Online,"
	matchingLines, err := findMatchingLines(file, "RaidMode: Online,")
	if err != nil {
		return nil, err
	}

	// Find the latest matching line.
	latestLine, err := findLatestLine(matchingLines)
	if err != nil {
		return nil, err
	}

	// Extract the IP from the latest matching line.
	ip, err := extractIPFromLine(latestLine)
	if err != nil {
		return nil, err
	}

	return net.ParseIP(ip), nil
}

// findMatchingLines scans the file and returns lines containing the target string.
func findMatchingLines(file *os.File, target string) ([]string, error) {
	// Scan the file for lines containing the target string.
	var matchingLines []string

	// TODO: Use a buffered reader instead of a scanner.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, target) {
			matchingLines = append(matchingLines, line)
		}
	}

	// Check for errors.
	if len(matchingLines) == 0 {
		return nil, ErrorNoMatchingLine
	}

	return matchingLines, nil
}

// findLatestLine returns the last line from a slice of lines.
func findLatestLine(lines []string) (
	string, error) {
	if len(lines) == 0 {
		return "", ErrorNoLines
	}

	return lines[len(lines)-1], nil
}

// extractIPFromLine extracts the IP address from a line using regular expressions.
func extractIPFromLine(line string) (string, error) {
	re := regexp.MustCompile(`Ip: ([\d\.]+),`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 2 {
		return "", ErrorNoIP
	}

	return matches[1], nil
}
