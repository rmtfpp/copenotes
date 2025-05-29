package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/rmtfpp/copenotes/pkg/user"
)

var ErrAuth1 = errors.New("user doesnt exist")
var ErrAuth2 = errors.New("unauthorized")
var ErrAuth3 = errors.New("unauthorized")
var ErrSessionExpired = errors.New("session expired")

func Authorize(r *http.Request) (int, error) {

	username := r.FormValue("username")

	u, err := user.GetUserByUsername(username)
	if err != nil {
		return 1, ErrAuth1
	}

	session, _ := user.GetSession(u.ID)
	if session.ExpiresAt.Before(time.Now()) {
		return 4, ErrSessionExpired
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != session.SessionToken {
		return 2, ErrAuth2
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != session.CSRFToken || csrf == "" {
		return 3, ErrAuth3
	}

	return 0, nil

}
