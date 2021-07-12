package result

type FileInfo struct {
	Path               string    `json:"path"`
	Licenses           []License `json:"licenses"`
	LicenseExpressions []string  `json:"license_expressions"`
	Copyrights         []CpInfo  `json:"copyrights"`
	Scan_Error         string    `json:"scan_error"`
}

type License struct {
	Key        string  `json:"key"`
	Score      float64 `json:"score"`
	StartLine  int     `json:"start_line"`
	EndLine    int     `json:"end_line"`
	StartIndex int     `json:"start_index"`
	EndIndex   int     `json:"end_index"`
}

type CpInfo struct {
	Notification string `json:"notification"`
	StartIndex   int    `json:"start_index"`
	EndIndex     int    `json:"end_index"`
	Holder       string `json:"holder"`
}
