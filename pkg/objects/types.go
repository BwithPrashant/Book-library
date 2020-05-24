package objects

type AddBooksRequest struct {
	ISBN    string `json:"isbn"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Country string `json:"country,omitempty"`
}
