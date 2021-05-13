package models

type ResetDTO struct {
	Username string
	Hash     string
	Newhash  string
	N        int
	Salt     string
}

type RegisterDTO struct {
	Username string
	Hash     string
	N        int
	Salt     string
}

type LoginDTO struct {
	Username string
	Hash     string
}

type GetNDTO struct {
	Username string
}

type GetNResponse struct {
	N    int
	Salt string
}
