package result

import (
	"encoding/json"
)

func InitFile(path string) *FileInfo {
	return &FileInfo{
		Path:               path,
		Licenses:           make([]Licenses, 0),
		Copyrights:         make([]Copyrights, 0),
		Holders:            make([]Holder, 0),
		LicenseExpressions: make([]string, 0),
		Scan_Errors:        make([]string, 0),
	}
}

func (file *FileInfo) GetJSONString() (string, error) {
	jsonString, err := json.MarshalIndent(&file, "", " ")
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}
