package bleemeo

import (
	"context"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type authenticationProvider struct {
	tokenSource oauth2.TokenSource
}

func newAuthProvider(endpointURL, username, password, clientID string) authenticationProvider {
	cfg := clientcredentials.Config{
		ClientID: clientID,
		TokenURL: endpointURL + "/o/token/",
		EndpointParams: url.Values{
			"grant_type": {"password"},
			"username":   {username},
			"password":   {password},
		},
		AuthStyle: oauth2.AuthStyleInParams,
	}

	return authenticationProvider{
		tokenSource: cfg.TokenSource(context.Background()),
	}
}

func (ap authenticationProvider) injectHeader(req *http.Request) error {
	tk, err := ap.tokenSource.Token()
	if err != nil {
		return err
	}

	tk.SetAuthHeader(req)

	return nil
}
