package auth

import (
	"errors"
	"net/http"

	"github.com/rmtfpp/copenotes/pkg/user"
)

var ErrAuth1 = errors.New("Unauthorized1")
var ErrAuth2 = errors.New("Unauthorized2")
var ErrAuth3 = errors.New("Unauthorized3")

func Authorize(r *http.Request) (int, error) {

	username := r.FormValue("username")

	u, err := user.GetUserByUsername(username)
	if err != nil {
		return 1, ErrAuth1
	}

	st, err := r.Cookie("session_token")
	sessionTokenDb, _ := user.GetSessionToken(u.ID)
	if err != nil || st.Value == "" || st.Value != sessionTokenDb {
		return 2, ErrAuth2
	}

	csrf := r.Header.Get("X-CSRF-Token")
	csrfTokenDb, _ := user.GetCSRFToken(u.ID)
	if csrf != csrfTokenDb || csrf == "" {
		return 3, ErrAuth3
	}

	return 0, nil

}
