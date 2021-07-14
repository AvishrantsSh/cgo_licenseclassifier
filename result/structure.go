package result

type FileInfo struct {
	Path               string       `json:"path"`
	Licenses           []Licenses   `json:"licenses"`
	LicenseExpressions []string     `json:"license_expressions"`
	Copyrights         []Copyrights `json:"copyrights"`
	Holders            []Holder     `json:"holders"`
	Scan_Errors        []string     `json:"scan_errors"`
}

type Licenses struct {
	Key        string  `json:"key"`
	Score      float64 `json:"score"`
	StartLine  int     `json:"start_line"`
	EndLine    int     `json:"end_line"`
	StartIndex int     `json:"start_index"`
	EndIndex   int     `json:"end_index"`
}

type Copyrights struct {
	Notification string `json:"value"`
	// StartLine    int    `json:"start_line"`
	// EndLine      int    `json:"end_line"`
	StartIndex int `json:"start_index"`
	EndIndex   int `json:"end_index"`
}

type Holder struct {
	Holder string `json:"value"`
	// StartLine  int    `json:"start_line"`
	// EndLine    int    `json:"end_line"`
	StartIndex int `json:"start_index"`
	EndIndex   int `json:"end_index"`
}
