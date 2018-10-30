package InformaCast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// RestDialCastDialingConfig represents the JSON the API expects to receive.
// The bad naming convention is provided to you by the official InformaCast REST API documentation.
type RestDialCastDialingConfig struct {
	id                  int64
	dialingPatternRegEx string
	authType            int64
	messageId           int64
	messageDescription  string
	recipientGroups     []int64
	recipientGroupName  string
	dialcodes           []string
}

// DialingConfigurationResponses is the json object returned during get requests.
type DialingConfigurationResponses struct {
	Index               int64
	Id                  int64
	dialingPatternRegEx string
	message             string
	messageDescription  string
	link                string
}

type DialingConfigrations struct {
	QueryParameters
	HttpParameters
}

func (d DialingConfigrations) Get() ( Data []DialingConfigurationResponses, err error )  {

	url := "https://" + d.Server + ":8444/InformaCast/RESTServices/V1/Admin/DialCast/dialingConfigurations" + d.QueryParameters.Compile()

	// make the get Request
	response, err := request("Get", url, nil, d.Username, d.Password)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//parse response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	// we are only interested in the data part of this object.
	// the api allows batching the get requests, but for now this modules just gets everything.
	var fullJsonResponse struct {
		Total uint64
		Previous uint64
		Next string
		Data []DialingConfigurationResponses
	}

	json.Unmarshal(body, &fullJsonResponse)

	return fullJsonResponse.Data , nil
}

// Find will retrieve a single DialCast record
func (d DialingConfigrations) GetOne(c RestDialCastDialingConfig) ( Data DialingConfigurationResponses, err error )   {

	url := "https://" + d.Server + ":8444/InformaCast/RESTServices/V1/Admin/DialCast/dialingConfigurations/" + fmt.Sprint(c.id) + d.QueryParameters.Compile()

	// create json byte array
	Json, err := json.Marshal(c)
	if err != nil {
		return Data, err
	}

	// make the get Request
	response, err := request("Get", url, Json , d.Username, d.Password)
	if err != nil {
		return Data, err
	}
	defer response.Body.Close()

	//parse response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	// we are only interested in the data part of this object.
	// the api allows batching the get requests, but for now this modules just gets everything.
	var fullJsonResponse struct {
		Total uint64
		Previous uint64
		Next string
		Data []DialingConfigurationResponses
	}

	json.Unmarshal(body, &fullJsonResponse)

	return fullJsonResponse.Data[0] , nil
}


// New Creates a new DialCast record
func (d DialingConfigrations) New(c RestDialCastDialingConfig) ( Data DialingConfigurationResponses, err error )   {

	url := "https://" + d.Server + ":8444/InformaCast/RESTServices/V1/Admin/DialCast/dialingConfigurations" + d.QueryParameters.Compile()

	// create json byte array
	Json, err := json.Marshal(c)
	if err != nil {
		return Data, err
	}

	// make the get Request
	response, err := request("POST", url,Json, d.Username, d.Password)
	if err != nil {
		return Data, err
	}
	defer response.Body.Close()

	//parse response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	// we are only interested in the data part of this object.
	// the api allows batching the get requests, but for now this modules just gets everything.
	var fullJsonResponse struct {
		Total uint64
		Previous uint64
		Next string
		Data []DialingConfigurationResponses
	}

	json.Unmarshal(body, &fullJsonResponse)

	return fullJsonResponse.Data[0] , nil
}


// Update updates a single DialCast record.
func (d DialingConfigrations) Update(c RestDialCastDialingConfig) ( Data DialingConfigurationResponses, err error ) {

	url := "https://" + d.Server + ":8444/InformaCast/RESTServices/V1/Admin/DialCast/dialingConfigurations/" + fmt.Sprint(c.id) + d.QueryParameters.Compile()

	// create json byte array
	Json, err := json.Marshal(c)
	if err != nil {
		return Data, err
	}

	// make the get Request
	response, err := request("Get", url, Json , d.Username, d.Password)
	if err != nil {
		return Data, err
	}
	defer response.Body.Close()

	//parse response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	// we are only interested in the data part of this object.
	// the api allows batching the get requests, but for now this modules just gets everything.
	var fullJsonResponse struct {
		Total uint64
		Previous uint64
		Next string
		Data []DialingConfigurationResponses
	}

	json.Unmarshal(body, &fullJsonResponse)

	return fullJsonResponse.Data[0] , nil
}

func (d DialingConfigrations) Delete(c RestDialCastDialingConfig) (err error) {

	// find the record
	data, err := d.GetOne(c)

	if data.Id == 0 {
		return errors.New("Could not find DialCast record.")
	}

	url := "https://" + d.Server + ":8444/InformaCast/RESTServices/V1/Admin/DialCast/dialingConfigurations/" + fmt.Sprint(c.id)

	// create json byte array
	Json, err := json.Marshal(c)
	if err != nil {
		return err
	}

	// make the get Request
	response, err := request("Get", url, Json , d.Username, d.Password)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}
