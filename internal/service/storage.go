package service

type Word struct {
	Title       string `json:"title"`
	Translation string `json:"translation"`
	Id          int    `json:"id"`
}

type Report struct {
	WordId      int    `json:"word_id"`
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
