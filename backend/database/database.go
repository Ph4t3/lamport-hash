package database

// Data Structure to hold the details of a user
type UserCreds struct {
	Hash string
	Salt string
	N    int
}

// Map which maps the username to user detail
var Users = make(map[string]UserCreds)

// Set adds a user to the database
func Set(username string, creds UserCreds) {
	Users[username] = creds
}

// Get returns the details of a particular user
func Get(username string) UserCreds {
	return Users[username]
}

// Check checks if a user is present in the database
func Check(username string) bool {
	_, ok := Users[username]
	return ok
}
