package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var (
	versionInfo          = &VersionInfo{}
	jsonFilePath         string
	buildNumber          int
	gitTag               string
	incrementBuildNumber bool
)

// rootCmd is the main Cobra command for the versioninfo_manager tool.
// It reads the versioninfo.json file if it exists and fills in the missing fields with default values if not specified by user input flags.
// It then updates the version information based on the provided git tag and build number flags.
// Finally, it writes the updated version information to the versioninfo.json file.
var rootCmd = &cobra.Command{
	Use:          "versioninfo_manager",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(jsonFilePath); err == nil {
			file, err := os.Open(jsonFilePath)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			dec := json.NewDecoder(file)
			dec.Decode(versionInfo)
		}

		if gitTag != "" {
			tag := strings.TrimPrefix(gitTag, "v")

			major, _ := strconv.Atoi(strings.Split(tag, ".")[0])
			minor, _ := strconv.Atoi(strings.Split(tag, ".")[1])
			patch, _ := strconv.Atoi(strings.Split(tag, ".")[2])

			var build int

			if buildNumber != 0 {
				build = buildNumber
			} else {
				build = versionInfo.FixedFileInfo.FileVersion.Build
			}

			if major < versionInfo.FixedFileInfo.FileVersion.Major ||
				minor < versionInfo.FixedFileInfo.FileVersion.Minor ||
				patch < versionInfo.FixedFileInfo.FileVersion.Patch ||
				build < versionInfo.FixedFileInfo.FileVersion.Build {
				return errors.New("version can't be lower than the current version")
			}

			if incrementBuildNumber && buildNumber == versionInfo.FixedFileInfo.FileVersion.Build {
				buildNumber++
			}

			versionInfo.FixedFileInfo.FileVersion.Major = major
			versionInfo.FixedFileInfo.FileVersion.Minor = minor
			versionInfo.FixedFileInfo.FileVersion.Patch = patch
			versionInfo.FixedFileInfo.FileVersion.Build = build
			versionInfo.FixedFileInfo.ProductVersion.Major = major
			versionInfo.FixedFileInfo.ProductVersion.Minor = minor
			versionInfo.FixedFileInfo.ProductVersion.Patch = patch
			versionInfo.FixedFileInfo.ProductVersion.Build = build
			versionInfo.StringFileInfo.ProductVersion = tag
			versionInfo.StringFileInfo.FileVersion = tag
			versionInfo.StringFileInfo.Comments = "Tarkov Server Checker " + tag
			versionInfo.StringFileInfo.CompanyName = "Tarkov Server Checker"
			versionInfo.StringFileInfo.OriginalFilename = "tarkov-server-checker.exe"
			versionInfo.StringFileInfo.FileDescription = "Tarkov Server Checker"
			versionInfo.StringFileInfo.InternalName = "tarkov-server-checker.exe"
			versionInfo.StringFileInfo.LegalTrademarks = "Tarkov Server Checker"
			versionInfo.StringFileInfo.ProductName = "Tarkov Server Checker"
			versionInfo.StringFileInfo.SpecialBuild = "Tarkov Server Checker"

			versionInfo.StringFileInfo.FileVersion = fmt.Sprintf(
				"%d.%d.%d.%d",
				versionInfo.FixedFileInfo.FileVersion.Major,
				versionInfo.FixedFileInfo.FileVersion.Minor,
				versionInfo.FixedFileInfo.FileVersion.Patch,
				versionInfo.FixedFileInfo.FileVersion.Build,
			)
		}

		file, err := os.OpenFile(
			jsonFilePath, os.O_CREATE|
				os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		enc := json.NewEncoder(file)
		enc.SetIndent("", "  ")
		enc.Encode(versionInfo)

		return nil
	},
}

var getVerCmd = &cobra.Command{
	Use:          "get",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(jsonFilePath); err == nil {
			file, err := os.Open(jsonFilePath)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			dec := json.NewDecoder(file)
			dec.Decode(versionInfo)
		}
		fmt.Println(versionInfo.StringFileInfo.FileVersion)
		return nil
	},
}

func init() {
	rootCmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileFlagsMask, "fileFlagsMask", "3f", "File flags mask")
	rootCmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileFlags, "fileFlags", "00", "File flags")
	rootCmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileOS, "fileOS", "040004b0", "File OS")
	rootCmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileType, "fileType", "01", "File type")
	rootCmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileSubType, "fileSubType", "00", "File sub type")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.FileVersion, "fileVersion", "", "File version")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.ProductVersion, "productVersion", "", "Product version")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.Comments, "comments", "Tarkov Server Checker", "Comments")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.CompanyName, "companyName", "", "Company name")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.OriginalFilename, "originalFilename", "tarkov-server-checker.exe", "Original filename")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.FileDescription, "fileDescription", "Tarkov Server Checker", "File description")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.InternalName, "internalName", "tarkov-server-checker.exe", "Internal name")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.LegalTrademarks, "legalTrademarks", "", "Legal trademarks")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.ProductName, "productName", "Tarkov Server Checker", "Product name")
	rootCmd.Flags().StringVar(&versionInfo.StringFileInfo.SpecialBuild, "specialBuild", "Tarkov Server Checker", "Special build")

	rootCmd.PersistentFlags().StringVarP(&gitTag, "gitTag", "t", "", "Git tag")
	rootCmd.Flags().IntVarP(&buildNumber, "buildNumber", "b", versionInfo.FixedFileInfo.FileVersion.Build, "Build number")
	rootCmd.PersistentFlags().StringVarP(&jsonFilePath, "jsonFilePath", "j", "versioninfo.json", "Json file path")
	rootCmd.Flags().BoolVarP(&incrementBuildNumber, "incrementBuildNumber", "i", false, "Increment build number")

	rootCmd.AddCommand(getVerCmd)
}

func main() {
	rootCmd.Execute()
}
