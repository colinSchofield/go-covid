package history

type TableDetails struct {
	Flag     string   `json:"flag"`
	Labels   []string `json:"labels"`
	NewCases []int    `json:"newCases"`
	Deaths   []int    `json:"deaths"`
}
