module 01-Login

go 1.12

require (
	admin v0.0.0
	app v0.0.0
	callback v0.0.0
	edit v0.0.0
	github.com/codegangsta/negroni v1.0.0
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/gorilla/mux v1.8.0
	home v0.0.0
	login v0.0.0
	logout v0.0.0
	middlewares v0.0.0
	rbac v0.0.0 // indirect
	unauthorized v0.0.0
)

replace admin => ./routes/admin

replace app => ./app

replace auth => ./auth

replace callback => ./routes/callback

replace edit => ./routes/edit

replace home => ./routes/home

replace login => ./routes/login

replace logout => ./routes/logout

replace middlewares => ./routes/middlewares

replace rbac => ./rbac

replace templates => ./routes/templates

replace unauthorized => ./routes/unauthorized

replace user => ./routes/user
