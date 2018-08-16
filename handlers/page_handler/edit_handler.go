package page_handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/teimurjan/go-plain-server/dao"
)

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := dao.LoadPage(title)
	if err != nil {
		fmt.Fprint(w, "Not found.")
		return
	}
	t, _ := template.ParseFiles("templates/page/edit.html")
	t.Execute(w, p)
}
