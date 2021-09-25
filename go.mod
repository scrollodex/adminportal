module github.com/scrollodex/scrollodex

go 1.17

require (
	admin v0.0.0-00010101000000-000000000000
	app v0.0.0
	callback v0.0.0
	edit v0.0.0
	editrow v0.0.0-00010101000000-000000000000
	github.com/codegangsta/negroni v1.0.0
	github.com/gorilla/mux v1.8.0
	golang.org/x/text v0.3.7
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	home v0.0.0
	login v0.0.0
	logout v0.0.0
	middlewares v0.0.0
	unauthorized v0.0.0
	zingdata v0.0.0
)

require (
	auth v0.0.0 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.1 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	rbac v0.0.0 // indirect
	templates v0.0.0 // indirect
)

replace admin => ./adminportal/routes/admin

replace app => ./adminportal/app

replace auth => ./adminportal/auth

replace callback => ./adminportal/routes/callback

replace edit => ./adminportal/routes/edit

replace editrow => ./adminportal/routes/editrow

replace home => ./adminportal/routes/home

replace login => ./adminportal/routes/login

replace logout => ./adminportal/routes/logout

replace middlewares => ./adminportal/routes/middlewares

replace rbac => ./adminportal/rbac

replace templates => ./adminportal/routes/templates

replace unauthorized => ./adminportal/routes/unauthorized

replace zingdata => ./adminportal/routes/zingdata
