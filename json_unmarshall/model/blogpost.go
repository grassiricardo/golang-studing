package model

//BlogPost armazenada dados blogpost
type BlogPost struct {
	Usuario int    `json:"userid"`
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
}
