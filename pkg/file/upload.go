package file

import (
	"fmt"
	"net/http"

	"github.com/rmtfpp/copenotes/pkg/auth"
	"github.com/rmtfpp/copenotes/pkg/user"
)

func Upload(w http.ResponseWriter, r *http.Request) {
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

	filename := r.FormValue("filename")

	err := user.CreateFile(u.ID, filename)
	if err != nil {
		err := http.StatusInternalServerError
		http.Error(w, "failed to create file", err)
		return
	}

	fmt.Fprintf(w, "File created succesfully")
}
