package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ssapp/tarkov-server-checker/internal/eftlog"
	"github.com/ssapp/tarkov-server-checker/internal/ip_location"
	"github.com/ssapp/tarkov-server-checker/internal/systray"
)

var (
	// Commit is the commit hash of the application.
	Commit = "unknown"
	// Version is the version of the application.
	Version = "0.0.0-dev"
	// FullVersion is the full version string of the application.
	FullVersion = fmt.Sprintf("%s (%s)", Version, Commit)
	// logPath is the path to the Tarkov log files directory.
	logPath string
)

var rootCmd = &cobra.Command{
	Use: "tarkov-server-checker",
	Long: `Tarkov Server Checker is a system tray application that checks
the IP address of the Tarkov server you are connected to and displays it
in the system tray.`,
	Version:      FullVersion,
	SilenceUsage: true,
	Run: func(cmd *cobra.Command, args []string) {
		systemtray := systray.New(eftlog.New())
		systemtray.Run()
	},
}

var runOnceCmd = &cobra.Command{
	Use:   "run-once",
	Short: "Check the server IP address once, and print the results.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var log *eftlog.Log

		log = eftlog.New()

		if logPath != "" {
			log = eftlog.New(eftlog.WithPath(logPath))
		}

		ip, err := log.GetIP()
		if err != nil {
			return err
		}

		country, err := ip_location.Get(ip)
		if err != nil {
			return err
		}

		fmt.Printf("Server: %s (%s)\n", ip, country)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(runOnceCmd)
	rootCmd.AddGroup(
		&cobra.Group{
			ID:    "app",
			Title: "Application commands",
		},
	)

	rootCmd.
		CompletionOptions.
		DisableDefaultCmd = true

	runOnceCmd.GroupID = "app"

	// Add the log path flag. This is used to specify the path to the Tarkov log files directory.
	logPathDesc := "`path` to the Tarkov log file (default: %s)"
	logPathFormatted := fmt.Sprintf(logPathDesc, eftlog.LogPathDefault)
	rootCmd.PersistentFlags().StringVarP(&logPath, "log-path", "l", "", logPathFormatted)
}

func Execute() error {
	return rootCmd.Execute()
}
