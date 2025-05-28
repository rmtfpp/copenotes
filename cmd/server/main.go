package main

import (
	"net/http"

	"github.com/rmtfpp/copenotes/pkg/auth"
	"github.com/rmtfpp/copenotes/pkg/database"
	"github.com/rmtfpp/copenotes/pkg/session"
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
	session.MigrateSessions()

	http.HandleFunc("/register", auth.Register)
	http.HandleFunc("/login", auth.Login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":8080", nil)
	//user.DeleteUser("a1329985-516f-4396-abe7-3f801a320fb6")
	//user.CreateUser("Luca", "Martinetti", "lucamarti@gmail.com", "passapa")
}

func logout(w http.ResponseWriter, r *http.Request)    {}
func protected(w http.ResponseWriter, r *http.Request) {}
