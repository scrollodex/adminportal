package rbac

import (
	"strings"
)

const UserTom = "https://dev-2vzqnqjr.us.auth0.com/ google-oauth2|108963384323763341815"
const UserBh = "https://dev-2vzqnqjr.us.auth0.com/ google-oauth2|101744589201358810643"

var policy = map[string]map[string]bool{
	UserTom: {
		"editor": true,
	},
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
