package priv

type user struct {
	name    string
	age     int
	animals []string
}

func NewUser() user {
	return user{
		name:    "admin",
		age:     50,
		animals: []string{"roger", "barry", "melissa"},
	}
}
