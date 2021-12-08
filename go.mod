module github.com/scrollodex/adminportal

go 1.17

require (
	github.com/codegangsta/negroni v1.0.0
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/sessions v1.2.1
	github.com/joho/godotenv v1.4.0
	github.com/scrollodex/dex v0.0.0-20211208115128-b34845446c27
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1
)

replace github.com/scrollodex/dex/reslist => ./dex/reslist

require (
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
