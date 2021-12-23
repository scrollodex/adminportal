package logout

import (
	"net/http"
	"net/url"
	"os"
)

// Handler renders the page.
func Handler(w http.ResponseWriter, r *http.Request) {

	domain := os.Getenv("AUTH0_DOMAIN")

	logoutURL, err := url.Parse("https://" + domain)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logoutURL.Path += "/v2/logout"
	parameters := url.Values{}

	var scheme string
	if r.TLS == nil {
		scheme = "http"
	} else {
		scheme = "https"
	}

	rhost := os.Getenv("ADMINPORTAL_REDIRECT_BASE")
	if rhost != "" {
		parameters.Add("returnTo", rhost)
	} else {
		returnTo, err := url.Parse(scheme + "://" + r.Host)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		parameters.Add("returnTo", returnTo.String())
	}

	parameters.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutURL.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}
