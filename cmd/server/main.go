package main

import (
	"fmt"
	"net/http"

	"github.com/rmtfpp/copenotes/pkg/auth"
	"github.com/rmtfpp/copenotes/pkg/database"
	"github.com/rmtfpp/copenotes/pkg/file"
	"github.com/rmtfpp/copenotes/pkg/user"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

func main() {
	database.InitializeDatabase()
	user.MigrateUsers()
	user.MigrateSessions()
	user.MigrateFiles()

	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.HandleFunc("/upload", file.Upload)
	http.ListenAndServe(":8080", nil)
	//user.DeleteUser("a1329985-516f-4396-abe7-3f801a320fb6")
	//user.CreateUser("Luca", "Martinetti", "lucamarti@gmail.com", "passapa")
}

func protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid method", err)
		return
	}

	if i, err := auth.Authorize(r); err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "unauthorized", err)
		fmt.Fprintf(w, "%d", i)
		return
	}

	username := r.FormValue("username")
	u, _ := user.GetUserByUsername(username)
	fmt.Fprintf(w, "CSFR validation succesful! Welcome %s", u.FirstName)
}
func logout(w http.ResponseWriter, r *http.Request) {}
