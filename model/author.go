package model

type Author struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func NewAuthor(name string) *Author {
	return &Author{
		Name: name,
	}
}
