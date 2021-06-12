package result

import (
	"time"
)

// JSON structure for storing scan results
type JSON_struct struct {
	header Header
	files  []FileInfo
}

// Header Info struct
type Header struct {
	Tool_name       string    `json:"tool_name"`
	Input           string    `json:"input"`
	Start_timestamp time.Time `json:"start_timestamp"`
	End_timestamp   time.Time `json:"end_timestamp"`
	Duration        float64   `json:"duration"`
	Files_count     int       `json:"files_count"`
}

type FileInfo struct {
	Path       string    `json:"path"`
	Licenses   []License `json:"licenses"`
	Copyrights []CpInfo  `json:"copyrights"`
	Errors     string    `json:"errors"`
}

type License struct {
	Expression string  `json:"expression"`
	Confidence float64 `json:"confidence"`
	StartLine  int     `json:"start_line"`
	EndLine    int     `json:"end_line"`
	StartIndex int     `json:"start_index"`
	EndIndex   int     `json:"end_index"`
}

type CpInfo struct {
	Expression string `json:"expression"`
	StartIndex int    `json:"start_index"`
	EndIndex   int    `json:"end_index"`
	Holder     string `json:"holder"`
}
