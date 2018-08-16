package page_handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/teimurjan/go-plain-server/dao"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := dao.LoadPage(title)
	if err != nil {
		fmt.Fprint(w, "Not found.")
		return
	}
	t, _ := template.ParseFiles("templates/page/view.html")
	t.Execute(w, p)
}
