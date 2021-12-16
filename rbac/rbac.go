package rbac

import (
	"strings"
)

var policy = map[string]map[string]bool{
	// Tom <tal@whatexit.org>
	"https://dev-2vzqnqjr.us.auth0.com/ google-oauth2|108963384323763341815": {
		"editor": true,
	},

	// Michael Litzky <wondroustales@gmail.com>
	"https://dev-2vzqnqjr.us.auth0.com/ auth0|61baad71fa2cd10069eb2baf": {
		"editor": true,
	},

	// BestHabit3 <besthabit3@gmail.com> (intentionally has no access. For testing.)
	"https://dev-2vzqnqjr.us.auth0.com/ google-oauth2|101744589201358810643": {},
}

func Can(who, verb string) bool {
	if perms, ok := policy[who]; ok {
		if can, ok := perms[verb]; ok {
			return can
		}
	}
	return false
}

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
