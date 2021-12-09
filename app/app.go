package app

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

	parts := strings.SplitN(os.Getenv("ADMINPORTAL_SESSIONSTORE"), ":", 2)
	fmt.Printf("SESSION: type=%q\n", parts[0])
	switch parts[0] {

	case "file":
		fmt.Printf("SESSION: file\n")
		p := strings.SplitN(parts[1], ",", 2)
		Store = sessions.NewFilesystemStore(p[0], []byte(p[1]))

		//	case "redistore":
		//		fmt.Printf("SESSION: redistore\n")
		//		err, size, network, address, password, keyPairs := parseRedistore(parts[1])
		//		if err != nil {
		//			return fmt.Errorf("Init: %w", err)
		//		}
		//		fmt.Printf("DEBUG: redistore.NewRediStore(%d, %q, %q, %q, %q)\n", size, network, address, password, keyPairs)
		//		store, err := redistore.NewRediStore(size, network, address, password, keyPairs)
		//		if err != nil {
		//			fmt.Printf("DEBUG: failed redistore.NewRediStore: %w\n", err)
		//			panic(err)
		//		}
		//		//defer Store.Close()
		//		fmt.Printf("SESSION: redistore CONNECTED\n")

	default:
		fmt.Printf("SESSION: default\n")
		Store = sessions.NewFilesystemStore("", []byte("a secret only i know"))
	}

	gob.Register(map[string]interface{}{})
	return nil
}

func parseRedistore(s string) (err error, size int, network, address, password string, key []byte) {

	p := strings.Split(s, ",")

	size, err = strconv.Atoi(p[0])
	if err != nil {
		return fmt.Errorf("parseRedistore: invalid size %q: %w", p[0], err), 0, "", "", "", []byte{}
	}

	network = p[1]
	address = p[2]
	password = p[3]
	key = []byte(p[4])
	// TODO(tlim): handle multiple keyPairs
	return

}
