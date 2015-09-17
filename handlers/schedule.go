package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

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

func PostProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	db := context.Get(r, "db").(*sqlx.DB)

	name := r.FormValue("Name")
	roastDate, err := time.Parse("2006-01-02 03:04", r.FormValue("RoastDate"))
	description := r.FormValue("Description")
	price, err := strconv.ParseFloat(r.FormValue("Price"), 64)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	_, err = dal.NewProduct(db).InsertNewProduct(nil, name, description, roastDate, price)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	http.Redirect(w, r, "/schedule", 302)
}
