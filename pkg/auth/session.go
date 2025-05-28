package auth

import (
	"errors"
	"net/http"

	"github.com/rmtfpp/copenotes/pkg/session"
	"github.com/rmtfpp/copenotes/pkg/user"
)

var ErrAuth = errors.New("Unauthorized")

func Authorize(r *http.Request) error {

	username := r.FormValue("username")

	u, err := user.GetUserByUsername(username)
	if err != nil {
		return ErrAuth
	}

	st, err := r.Cookie("session_token")
	sessionTokenDb, _ := session.GetSessionToken(u.ID)
	if err != nil || st.Value == "" || st.Value != sessionTokenDb {
		return ErrAuth
	}

	csrf := r.Header.Get("X-CSRF-Token")
	csrfTokenDb, _ := session.GetCSRFToken(u.ID)
	if csrf != csrfTokenDb || csrf == "" {
		return ErrAuth
	}

	return nil

}
