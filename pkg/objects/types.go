package objects

type AddBookRequest struct {
	Isbn    string `json:"isbn"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Country string `json:"country,omitempty"`
}

type BookIdentity struct {
	Id   string         `json:"id"`
	Data AddBookRequest `json:"data"`
}

type ModifyBookRequest struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Country string `json:"country,omitempty"`
}
