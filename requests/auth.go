package requests

type Register struct {
	Email           string
	Password        string
	ConfirmPassword string
	FirstName       string
	LastName        string
	City            string
	Country         string
	CompanyName     string
}

type Login struct {
	Email    string
	Password string
}
