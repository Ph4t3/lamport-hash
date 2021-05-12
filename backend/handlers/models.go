package handlers

type ResetDTO struct {
	Username string
	Hash     string
	Newhash  string
	N        int
}

type RegisterDTO struct {
	Username string
	Hash     string
	N        int
}

type LoginDTO struct {
	Username string
	Hash     string
}

type GetNDTO struct {
	Username string
}

type GetNResponse struct {
	N int
}
