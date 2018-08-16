package page_handler

import (
	"html/template"
	"net/http"

	"github.com/teimurjan/go-plain-server/dao"
)

func ViewAllHandler(w http.ResponseWriter, r *http.Request) {
	pages := dao.LoadAllPages()
	t, _ := template.ParseFiles("templates/page/view_all.html")

	t.Execute(w, pages)
}
