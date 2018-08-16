package page_handler

import (
	"fmt"
	"net/http"

	"github.com/teimurjan/go-plain-server/dao"
	"github.com/teimurjan/go-plain-server/models"
)

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &models.Page{Title: title, Body: []byte(body)}
	err := dao.SavePage(p)
	if err != nil {
		fmt.Fprint(w, "Not found.")
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
