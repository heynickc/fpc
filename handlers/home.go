package handlers

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/heynickc/fpc/dal"
	"github.com/heynickc/fpc/libhttp"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	cookieStore := context.Get(r, "cookieStore").(*sessions.CookieStore)

	session, _ := cookieStore.Get(r, "fpc-session")
	currentUser, ok := session.Values["user"].(*dal.UserRow)
	if !ok {
		http.Redirect(w, r, "/logout", 302)
		return
	}

	data := struct {
		CurrentUser *dal.UserRow
	}{
		currentUser,
	}

	tmpl, err := template.ParseFiles("templates/dashboard.html.tmpl", "templates/home.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, data)
}

func GetLanding(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	tmpl, err := template.ParseFiles("templates/landing.html.tmpl")
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	tmpl.Execute(w, nil)
}
