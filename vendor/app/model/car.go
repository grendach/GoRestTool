package model

type Car struct {
	Producent string `json:"producent,omitempty"`
	Model     string `json:"model,omitempty"`
	Colour    string `json:"colour,omitempty"`
	Year      int    `json:"year,omitempty"`
}
