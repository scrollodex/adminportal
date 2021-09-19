package app

import (
	"encoding/gob"
	"log"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Print(err.Error())
		return err
	}
	// FIXME(Tom): Get secret from ENV.
	Store = sessions.NewFilesystemStore("", []byte("a secret only i know"))
	gob.Register(map[string]interface{}{})
	return nil
}
