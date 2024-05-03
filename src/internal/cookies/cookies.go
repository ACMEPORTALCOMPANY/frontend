package cookies

import (
	"net/http"
	"time"
)

const TokenCookieName string = "acmeportalcookie"

func New(name, val string) *http.Cookie {
	return &http.Cookie{
		Name:    name,
		Value:   val,
		Expires: time.Now().Add(1 * time.Hour),
	}
}
