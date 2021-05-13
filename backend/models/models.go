package models

type RegisterDTO struct {
	Username string `validate:"required" json:"username"`
	Hash     string `validate:"required" json:"hash"`
	Salt     string `validate:"required" json:"salt"`
	N        int    `validate:"required,gt=2" json:"n"`
}

type LoginDTO struct {
	Username string `validate:"required" json:"username"`
	Hash     string `validate:"required" json:"hash"`
}

type GetNDTO struct {
	Username string `validate:"required" json:"username"`
}

type GetNResponse struct {
	N    int    `json:"n"`
	Salt string `json:"salt"`
}

type ResetDTO struct {
	Username string `validate:"required" json:"username"`
	Hash     string `validate:"required" json:"hash"`
	Newhash  string `validate:"required" json:"newhash"`
	Salt     string `validate:"required" json:"salt"`
	N        int    `validate:"required,gt=2" json:"n"`
}
