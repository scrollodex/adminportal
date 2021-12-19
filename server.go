package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/scrollodex/adminportal/routes/admin"
	"github.com/scrollodex/adminportal/routes/callback"
	"github.com/scrollodex/adminportal/routes/edit"
	"github.com/scrollodex/adminportal/routes/editrow"
	"github.com/scrollodex/adminportal/routes/home"
	"github.com/scrollodex/adminportal/routes/login"
	"github.com/scrollodex/adminportal/routes/logout"
	"github.com/scrollodex/adminportal/routes/middlewares"
	"github.com/scrollodex/adminportal/routes/unauthorized"
	"github.com/scrollodex/adminportal/routes/zingdata"

	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.HandleFunc("/unauthorized", unauthorized.UnauthorizedHandler)
	//r.Handle("/user", negroni.New(
	//	negroni.HandlerFunc(middlewares.IsAuthenticated),
	//	negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	//))
	r.Handle("/admin", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.HandlerFunc(middlewares.IsRbacEditor),
		negroni.Wrap(http.HandlerFunc(admin.AdminHandler)),
	))
	r.Handle("/admin/editrow/{site:[a-z]+}/{table:[a-z]+}/{id:[0-9]+}", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.HandlerFunc(middlewares.IsRbacEditor),
		negroni.Wrap(http.HandlerFunc(editrow.EditrowHandler)),
	)).Methods("GET")
	r.Handle("/admin/edit/{site:[a-z]+}/{table:[a-z]+}", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.HandlerFunc(middlewares.IsRbacEditor),
		negroni.Wrap(http.HandlerFunc(edit.EditHandler)),
	)).Methods("GET")
	r.Handle("/admin/zingdata/{site:[a-z]+}/{table:[a-z]+}.json", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.HandlerFunc(middlewares.IsRbacEditor),
		negroni.Wrap(http.HandlerFunc(zingdata.ZingDataHandler)),
	)).Methods("GET")
	pubdir := filepath.Join(os.Getenv("ADMINPORTAL_TEMPLATE_BASEDIR"), "public") + "/"
	fmt.Printf("DEBUG: pubdir=%q\n", pubdir)
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir(pubdir)))) // static files
	http.Handle("/", r)
	tcpport := os.Getenv("ADMINPORTAL_TCPPORT")
	log.Printf("Server listening on %v", tcpport)
	log.Fatal(http.ListenAndServe(tcpport, nil))
}
