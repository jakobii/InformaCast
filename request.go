package InformaCast

import (
	"bytes"
	"crypto/tls"
	"net/http"
)


// request is the https request handlers for the InformaCast api
// request handler allows for unsecure https connections
func request (restMethod string, url string, body []byte , username string , password string) ( response *http.Response ,err error) {

	// this allow for unsecure connections
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// the client will meke the https request.
	// transport is set to unsecure.
	client := &http.Client{Transport: tr}

	//request
	request, err := http.NewRequest( restMethod , url , bytes.NewReader(body))
	if err != nil {
		return nil , err
	}

	// basic Authorization
	request.SetBasicAuth( username, password )

	response, err = client.Do(request)
	if err != nil {
		return nil , err
	}
	defer response.Body.Close()

	return response, nil

}
