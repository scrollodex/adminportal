module 01-Login/routes/editrow

go 1.12

require (
	app v0.0.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1 // indirect
	templates v0.0.0
)

replace admin => ../admin

replace app => ../../app

replace auth => ../../auth

replace callback => ../callback

replace edit => ../edit

replace home => ../home

replace login => ../login

replace logout => ../logout

replace middlewares => ../middlewares

replace rbac => ../../rbac

replace templates => ../templates

replace unauthorized => ../unauthorized

replace zingdata => ../zingdata
