package reviews

type ReviewCreate struct {
	Details    string `json:"details"`
	AuthorName string `json:"authorName"`
}

type ReviewGet struct {
	ID int `json:"id"`
	ReviewCreate
}
