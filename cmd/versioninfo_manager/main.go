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

type VersionInfo struct {
	FixedFileInfo struct {
		FileVersion struct {
			Major int `json:"Major"`
			Minor int `json:"Minor"`
			Patch int `json:"Patch"`
			Build int `json:"Build"`
		} `json:"FileVersion"`
		ProductVersion struct {
			Major int `json:"Major"`
			Minor int `json:"Minor"`
			Patch int `json:"Patch"`
			Build int `json:"Build"`
		} `json:"ProductVersion"`
		FileFlagsMask string `json:"FileFlagsMask"`
		FileFlags     string `json:"FileFlags "`
		FileOS        string `json:"FileOS"`
		FileType      string `json:"FileType"`
		FileSubType   string `json:"FileSubType"`
	} `json:"FixedFileInfo"`
	StringFileInfo struct {
		Comments         string `json:"Comments"`
		CompanyName      string `json:"CompanyName"`
		OriginalFilename string `json:"OriginalFilename"`
		FileDescription  string `json:"FileDescription"`
		FileVersion      string `json:"FileVersion"`
		InternalName     string `json:"InternalName"`
		LegalCopyright   string `json:"LegalCopyright"`
		LegalTrademarks  string `json:"LegalTrademarks"`
		PrivateBuild     string `json:"PrivateBuild"`
		ProductName      string `json:"ProductName"`
		ProductVersion   string `json:"ProductVersion"`
		SpecialBuild     string `json:"SpecialBuild"`
	} `json:"StringFileInfo"`
	VarFileInfo struct {
		Translation struct {
			LangID    string `json:"LangID"`
			CharsetID string `json:"CharsetID"`
		} `json:"Translation"`
	} `json:"VarFileInfo"`
	IconPath     string `json:"IconPath"`
	ManifestPath string `json:"ManifestPath"`
}

var (
	versionInfo          = &VersionInfo{}
	jsonFilePath         string
	buildNumber          int
	gitTag               string
	gitCommit            string
	incrementBuildNumber bool
)

var cmd = &cobra.Command{
	Use:          "versioninfo_manager",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read versioninfo.json file if exists and fill in the missing fields with default values if not specified by user input flags
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

			if major < versionInfo.FixedFileInfo.FileVersion.Major {
				return errors.New("major version is lower than the current version")
			}

			if minor < versionInfo.FixedFileInfo.FileVersion.Minor {
				return errors.New("minor version is lower than the current version")
			}

			if patch < versionInfo.FixedFileInfo.FileVersion.Patch {
				return errors.New("patch version is lower than the current version")
			}

			if build < versionInfo.FixedFileInfo.FileVersion.Build {
				fmt.Println(build, versionInfo.FixedFileInfo.FileVersion.Build)
				return errors.New("build version is lower than the current version")
			}

			if incrementBuildNumber {
				if build == versionInfo.FixedFileInfo.FileVersion.Build {
					build++
				}
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

			versionInfo.StringFileInfo.FileVersion = fmt.Sprintf("%d.%d.%d.%d", versionInfo.FixedFileInfo.FileVersion.Major, versionInfo.FixedFileInfo.FileVersion.Minor, versionInfo.FixedFileInfo.FileVersion.Patch, versionInfo.FixedFileInfo.FileVersion.Build)
		}

		file, err := os.OpenFile(jsonFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
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
		// Read versioninfo.json file if exists and fill in the missing fields with default values if not specified by user input flags
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
	cmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileFlagsMask, "fileFlagsMask", "3f", "File flags mask")
	cmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileFlags, "fileFlags", "00", "File flags")
	cmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileOS, "fileOS", "040004b0", "File OS")
	cmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileType, "fileType", "01", "File type")
	cmd.Flags().StringVar(&versionInfo.FixedFileInfo.FileSubType, "fileSubType", "00", "File sub type")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.FileVersion, "fileVersion", "", "File version")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.ProductVersion, "productVersion", "", "Product version")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.PrivateBuild, "privateBuild", "", "Private build")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.SpecialBuild, "specialBuild", "", "Special build")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.LegalCopyright, "legalCopyright", "", "Legal copy right")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.Comments, "comments", "Tarkov Server Checker", "Comments")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.CompanyName, "companyName", "", "Company name")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.OriginalFilename, "originalFilename", "tarkov-server-checker.exe", "Original filename")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.FileDescription, "fileDescription", "Tarkov Server Checker", "File description")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.InternalName, "internalName", "tarkov-server-checker.exe", "Internal name")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.LegalTrademarks, "legalTrademarks", "", "Legal trademarks")
	cmd.Flags().StringVar(&versionInfo.StringFileInfo.ProductName, "productName", "Tarkov Server Checker", "Product name")
	cmd.Flags().StringVar(&versionInfo.VarFileInfo.Translation.LangID, "langID", "0409", "Lang ID")
	cmd.Flags().StringVar(&versionInfo.VarFileInfo.Translation.CharsetID, "charsetID", "04B0", "Charset ID")
	cmd.Flags().StringVar(&versionInfo.IconPath, "iconPath", "", "Icon path")
	cmd.Flags().StringVar(&versionInfo.ManifestPath, "manifestPath", "", "Manifest path")

	cmd.PersistentFlags().StringVarP(&gitTag, "gitTag", "t", "", "Git tag")
	cmd.Flags().IntVarP(&buildNumber, "buildNumber", "b", versionInfo.FixedFileInfo.FileVersion.Build, "Build number")
	cmd.Flags().StringVarP(&gitCommit, "gitCommit", "c", "", "Git commit")
	cmd.PersistentFlags().StringVarP(&jsonFilePath, "jsonFilePath", "j", "versioninfo.json", "Json file path")
	cmd.Flags().BoolVarP(&incrementBuildNumber, "incrementBuildNumber", "i", false, "Increment build number")

	cmd.AddCommand(getVerCmd)
}

func main() {
	cmd.Execute()
}
