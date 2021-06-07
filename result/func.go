package result

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// Initialize JSON_struct
func (j *JSON_struct) Init(root string, size int) {
	j.header = Header{
		Tool_name:       "Golicense_classifier",
		Root:            root,
		Files_count:     size,
		Start_timestamp: time.Now().UTC(),
	}

	j.files = make([]FileInfo, size)
}

func (j *JSON_struct) Finish(path string) error {
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
		Header Header
		Files  []FileInfo
	}{
		Header: j.header,
		Files:  j.files,
	})

	if err != nil {
		return nil, err
	}
	return info, nil
}

func (j *JSON_struct) AddFile(index int, file *FileInfo) {
	j.files[index] = *file
}
