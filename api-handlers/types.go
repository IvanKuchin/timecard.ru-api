package apihandlers

type users struct {
	Login    string
	Password string
}

type login_response struct {
	Result      string
	Description string
	Sessid      string
}
