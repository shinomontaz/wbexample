package distancer

type osrmResponseStatus struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	DataVersion string `json:"data_version"`
}

type osrmTableResponse struct {
	osrmResponseStatus
	Durations [][]float32 `json:"durations"`
}
