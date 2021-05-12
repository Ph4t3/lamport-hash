package database

type UserCreds struct {
	Hash string
	N    int
}

var Users = make(map[string]UserCreds)

func Set(username string, creds UserCreds) {
	Users[username] = creds
}

func Get(username string) UserCreds {
	return Users[username]
}

func Check(username string) bool {
	_, ok := Users[username]
	return ok
}
