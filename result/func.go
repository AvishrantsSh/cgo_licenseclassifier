package result

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

func InitJSON(root string, size int) *JSON_struct {
	j := new(JSON_struct)
	j.header = Header{
		Tool_name:       "Golicense_classifier",
		Input:           root,
		Files_count:     size,
		Start_timestamp: time.Now().UTC(),
		Errors:          make([]string, 0),
	}
	return j
}

func (j *JSON_struct) WriteJSON(path string) error {
	j.header.End_timestamp = time.Now().UTC()
	j.header.Duration = float64(j.header.End_timestamp.Sub(j.header.Start_timestamp)) / float64(time.Second)
	file, err := json.MarshalIndent(&j, "", " ")

	if err != nil {
		return err
	}
	write_error := ioutil.WriteFile(path, file, 0644)
	if write_error != nil {
		return write_error
	}
	return nil
}

// Custom Marshalling for JSON_struct
func (j *JSON_struct) MarshalJSON() ([]byte, error) {
	info, err := json.Marshal(struct {
		Header []Header   `json:"headers"`
		Files  []FileInfo `json:"files"`
	}{
		Header: []Header{j.header},
		Files:  j.files,
	})

	if err != nil {
		return nil, err
	}
	return info, nil
}

func (j *JSON_struct) AddFile(file *FileInfo) {
	j.files = append(j.files, *file)
}

func InitFile(path string) *FileInfo {
	return &FileInfo{
		Path:        path,
		Licenses:    make([]License, 0),
		Copyrights:  make([]CpInfo, 0),
		Expression:  make([]string, 0),
		Scan_Errors: make([]string, 0),
	}
}

func (file *FileInfo) GetJSONString() (string, error) {
	jsonString, err := json.MarshalIndent(&file, "", " ")
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}
