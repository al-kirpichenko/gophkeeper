package models

type SecretResponse struct {
	Title   string `json:"title"`
	Content []byte `json:"content"`
}

type CreateResponse struct {
	Message string `json:"message"`
}

type Secret struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Comment  string `json:"comment,omitempty"`
	Text     string `json:"text,omitempty"`
	Number   string `json:"number,omitempty"`
	Date     string `json:"date,omitempty"`
	Cvv      string `json:"cvv,omitempty"`
	Binary   []byte `json:"binary,omitempty"`
}

type Auth struct {
	Login    string
	Password string
	Comment  string
}
