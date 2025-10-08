package testing_utils


type User struct {
	FirstName, LastName string
}

func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}