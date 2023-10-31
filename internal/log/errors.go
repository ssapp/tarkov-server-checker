package log

import "errors"

var (
	// ErrNoIPFound is returned when no IP address is found in the log file.
	ErrNoIPFound = errors.New("no IP address found in log file")
	// ErrNoLogFilesFound is returned when no log files are found in the log directory.
	ErrNoLogFilesFound = errors.New("no log files found")
	// ErrNoLogDirectoriesFound is returned when no log directories are found in the log path.
	ErrNoLogDirectoriesFound = errors.New("no log directories found")
	// ErrNoMatchingLinesFound is returned when no matching lines are found in the log file.
	ErrNoMatchingLinesFound = errors.New("no entries in log file")
	// ErrNoLinesToSelectFrom is returned when there are no lines to select from.
	ErrNoLinesToSelectFrom = errors.New("no lines to select from")
	// ErrNoMatchingLineFound is returned when no matching line is found in the log file.
	ErrNoMatchingLineFound = errors.New("no matching line found in log file")
	// ErrNoMatchingDirectoryFound is returned when no matching directory is found in the log path.
	ErrNoMatchingDirectoryFound = errors.New("no matching directory found in log path")
	// ErrNoMatchingFileFound is returned when no matching file is found in the log directory.
	ErrNoMatchingFileFound = errors.New("no matching file found in log directory")
	// ErrNoMatchingDirectoryOrFileFound is returned when no matching directory or file is found in the log path.
	ErrNoMatchingDirectoryOrFileFound = errors.New("no matching directory or file found in log path")
	// ErrNoMatchingIPFound is returned when no matching IP is found in the log file.
	ErrNoMatchingIPFound = errors.New("no matching IP found in log file")
	// ErrNoMatchingIPsFound is returned when no matching IPs are found in the log file
	ErrUnableToExtractIPFromLine = errors.New("could not extract IP from line")
	// ErrInvalidIP is returned when the IP address is invalid.
	ErrInvalidIP = errors.New("invalid IP address")
	// ErrIPEmpty is returned when the IP address is empty.
	ErrIPEmpty = errors.New("IP address is empty")
	// ErrUnableToGetLocation is returned when the location information cannot be retrieved.
	ErrUnableToGetLocation = errors.New("unable to get location information")
	// ErrUnableToSerializeLocation is returned when the location information cannot be serialized.
	ErrUnableToSerializeLocation = errors.New("unable to serialize location information")
)
