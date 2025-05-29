package auth

import (
	"log"
	"net/http"

	"github.com/rmtfpp/copenotes/pkg/user"
	"github.com/rmtfpp/copenotes/pkg/utils/hash"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "web/templates/register.html")
		return
	}

	if r.Method == http.MethodPost {
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if len(email) < 8 || len(password) < 8 {
			http.Error(w, "Invalid email/password", http.StatusNotAcceptable)
			return
		}

		if exists, _ := user.EmailExists(email); exists {
			http.Error(w, "Email already registered", http.StatusConflict)
			return
		}
		if exists, _ := user.UsernameExists(username); exists {
			http.Error(w, "Username already registered", http.StatusConflict)
			return
		}

		hashedPassword, _ := hash.HashPassword(password)
		user.CreateUser(firstname, lastname, username, email, hashedPassword)
		log.Printf("user %s registered successfully", username)

		w.Write([]byte("Registration successful!"))
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
