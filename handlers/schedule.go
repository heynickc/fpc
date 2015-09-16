package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/heynickc/fpc/dal"
	"github.com/heynickc/fpc/libhttp"
	"github.com/jmoiron/sqlx"
)

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	db := context.Get(r, "db").(*sqlx.DB)
	p := dal.NewProduct(db)
	products, err := p.AllProducts(nil)

	data := struct {
		ProductRows []*dal.ProductRow
	}{
		products,
	}

	tmpl, err := template.ParseFiles("templates/schedule.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, data)
}
