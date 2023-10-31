package main

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
