package rbac

import (
	"strings"
)

type userInfo = struct {
	access map[string]bool
	email  string
}

var standardNone = map[string]bool{}
var standardEditor = map[string]bool{"editor": true}

var policy = map[string]userInfo{
	// BestHabit3, added Dec 2021
	// (intentionally has no access. For testing.)
	"https://dev-2vzqnqjr.us.auth0.com/ google-oauth2|101744589201358810643": userInfo{
		email:  "besthabit3@gmail.com",
		access: standardNone,
	},
	// Tom Limoncelli, Added Dec 2021
	"https://dev-2vzqnqjr.us.auth0.com/ google-oauth2|108963384323763341815": userInfo{
		email:  "tal@whatexit.org",
		access: standardEditor,
	},
	// Michael Litzky, Added Dec 2021
	"https://dev-2vzqnqjr.us.auth0.com/ auth0|61baad71fa2cd10069eb2baf": userInfo{
		email:  "wondroustales@gmail.com",
		access: standardEditor,
	},
	// Geri Weitzman, Added 2021-12-17
	"https://dev-2vzqnqjr.us.auth0.com/ auth0|61bbca7dfa2cd10069eb7780": userInfo{
		email:  "geriweitzman@gmail.com",
		access: standardEditor,
	},

	// INSERT NEW PEOPLE HERE.

}

// EmailOf returns the email address of a user specified by iss+sub.
func EmailOf(iss, sub string) string {
	return policy[iss+" "+sub].email
}

// Can returns true if who is entitled to do verb.
func Can(who, verb string) bool {
	if u, ok := policy[who]; ok {
		if can, ok := u.access[verb]; ok {
			return can
		}
	}
	return false
}

// MakeUsername returns a stable name for m.
func MakeUsername(m interface{}) string {
	//fmt.Printf("MakeUsername(%v)\n", m)
	mp, ok := m.(map[string]interface{})
	if !ok {
		panic("An entire interface changed type. I just can't.")
	}

	iss := strings.ReplaceAll(mp["iss"].(string), " ", "%20")
	sub := mp["sub"].(string)
	return iss + " " + sub
}
