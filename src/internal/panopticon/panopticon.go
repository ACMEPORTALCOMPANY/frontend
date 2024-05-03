package panopticon

import (
	"errors"
	"strings"
)

type Login struct {
	Token    string
	Handle   string
	Password string
}

type Registration struct {
	Handle          string
	DisplayName     string
	Password        string
	ConfirmPassword string
}

func LoginToken(tok string) (string, error) {
	l := Login{
		Token: tok,
	}

	return login(l)
}

func LoginCred(handle, password string) (string, error) {
	if strings.Trim(handle, " ") == "" {
		return "", errors.New("handle cannot be blank")
	}

	if strings.Trim(password, " ") == "" {
		return "", errors.New("password cannot be blank")
	}

	l := Login{
		Handle:   handle,
		Password: password,
	}

	return login(l)
}

func login(l Login) (string, error) {
	return "token", nil
}

func Register(r Registration) error {
	if strings.Trim(r.Handle, " ") == "" {
		return errors.New("handle cannot be blank")
	}

	if strings.Trim(r.Password, " ") == "" {
		return errors.New("password cannot be blank")
	}

	if r.Password != r.ConfirmPassword {
		return errors.New("passwords must match")
	}

	return nil
}
