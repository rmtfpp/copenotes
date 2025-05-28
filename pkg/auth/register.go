package auth

import (
	"log"
	"net/http"

	"github.com/rmtfpp/copenotes/pkg/user"
	"github.com/rmtfpp/copenotes/pkg/utils/hash"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}

	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	if len(email) < 8 || len(password) < 8 { // will have to check if email exists
		err := http.StatusNotAcceptable
		http.Error(w, "Invalid email/password", err)
		return
	}

	if exists, _ := user.EmailExists(email); exists {
		err := http.StatusConflict
		http.Error(w, "email already registered", err)
		return
	}

	if exists, _ := user.UsernameExists(username); exists {
		err := http.StatusConflict
		http.Error(w, "email already registered", err)
		return
	}

	hashedPassword, _ := hash.HashPassword(password)

	user.CreateUser(firstname, lastname, username, email, hashedPassword)
	log.Printf("user %s registered succesfully", username)

}
