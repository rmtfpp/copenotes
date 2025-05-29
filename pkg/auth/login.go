package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/rmtfpp/copenotes/pkg/user"
	"github.com/rmtfpp/copenotes/pkg/utils/hash"
	"github.com/rmtfpp/copenotes/pkg/utils/tokens"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/templates/login.html")
		return
	}

	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		u, err := user.GetUserByEmail(email)
		if err != nil {
			err := http.StatusBadRequest
			http.Error(w, "email not registered", err)
			return
		}

		if !hash.CheckPasswordHash(password, u.Password) {
			err := http.StatusUnauthorized
			http.Error(w, "incorrect username or password", err)
			return
		}

		if exists, _ := user.HasSession(u.ID); exists {
			log.Printf("user %s already logged in", u.UserName)
			return
		}

		sessionToken := tokens.GenerateToken(32)
		csrfToken := tokens.GenerateToken(32)

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    csrfToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: false,
		})

		user.CreateSession(u.ID, sessionToken, csrfToken)

		log.Printf("login succesful for user %s", u.UserName)
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

}
