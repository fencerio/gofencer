package gofencer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// API provides methods for accessing Fencer.io API
type API struct {
	apiKey    string
	accessKey string
	baseURI   string
	version   string
	lat       float64
	lng       float64
}

type course struct {
	Distance float64 `json:"distance"`
	Bearing  int     `json:"bearing"`
}

type geofence struct {
	ID     string   `json:"id,omitempty"`
	Alias  string   `json:"alias,omitempty"`
	Status string   `json:"status,omitempty"`
	Tags   []string `json:"tags,omitempty"`
}

type data struct {
	Inside bool   `json:"inside,omitempty"`
	Enter  course `json:"enter,omitempty"`
	Exit   course `json:"exit,omitempty"`
	Origin course `json:"origin,omitempty"`
}

// PositionResponse is a result struct containing unmarshaled JSON
type PositionResponse struct {
	Data  data   `json:"data"`
	Error string `json:"error"`
}

// NavigateResponse is a result struct containing unmarshaled JSON
type NavigateResponse struct {
	Data  course `json:"data"`
	Error string `json:"error"`
}

// GeofenceResponse is a result struct containing unmarshaled JSON
type GeofenceResponse struct {
	Data  []geofence `json:"data"`
	Error string     `json:"error"`
}

// GetGeofences returns public and private geofences associated with an API key
func (a *API) GetGeofences() (GeofenceResponse, error) {
	var response GeofenceResponse

	action := "geofence"
	endPoint := a.baseURI + a.version + "/" + action
	buf, err := a.makeRequest(endPoint)
	if err != nil {
		return response, err
	}

	// Unmarshal data to struct
	err = json.Unmarshal(buf, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

/*
// GetPublicGeofences returns public and private geofences associated with an API key
func (a *API) GetPublicGeofences() (Response, error) {
	// Not implemented currently so return empty struct
	var response Response
	return response, nil
}s
*/

// NavigateIn returns bearing and distance into geofence specified by accessKey
func (a *API) NavigateIn() (NavigateResponse, error) {
	var response NavigateResponse
	if a.accessKey != "" {
		action := "navigation/in"
		endPoint := a.baseURI + a.version + "/" + action + "/" + a.accessKey

		if a.lat != 0 && a.lng != 0 {
			buf, err := a.makeRequest(endPoint)
			if err != nil {
				return response, err
			}

			// Unmarshal data to struct
			err = json.Unmarshal(buf, &response)
			if err != nil {
				return response, err
			}
			return response, nil
		}

		return response, errors.New("Error: Lat or Lng not set")
	}
	return response, errors.New("Error: Access Key not set")
}

// NavigateOrigin returns bearing and distance to geofence origin/center specified by accessKey
func (a *API) NavigateOrigin() (NavigateResponse, error) {
	var response NavigateResponse
	if a.accessKey != "" {
		action := "navigation/origin"
		endPoint := a.baseURI + a.version + "/" + action + "/" + a.accessKey

		if a.lat != 0 && a.lng != 0 {
			buf, err := a.makeRequest(endPoint)
			if err != nil {
				return response, err
			}

			// Unmarshal data to struct
			err = json.Unmarshal(buf, &response)
			if err != nil {
				return response, err
			}
			return response, nil
		}
		return response, errors.New("Error: Lat or Lng not set")
	}
	return response, errors.New("Error: Access Key not set")
}

// NavigateOut returns bearing and distance to geofence origin/center specified by accessKey
func (a *API) NavigateOut() (NavigateResponse, error) {
	var response NavigateResponse
	if a.accessKey != "" {
		action := "navigation/out"
		endPoint := a.baseURI + a.version + "/" + action + "/" + a.accessKey
		if a.lat != 0 && a.lng != 0 {
			buf, err := a.makeRequest(endPoint)

			if err != nil {
				return response, err
			}

			// Unmarshal data to struct
			err = json.Unmarshal(buf, &response)
			if err != nil {
				return response, err
			}
			return response, nil
		}
		return response, errors.New("Error: Lat or Lng not set")
	}
	return response, errors.New("Error: Access Key not set")
}

// PositionInside tests if inside geofence specified by accessKey
func (a *API) PositionInside() (PositionResponse, error) {
	var response PositionResponse
	if a.accessKey != "" {
		action := "position/inside"
		endPoint := a.baseURI + a.version + "/" + action + "/" + a.accessKey
		if a.lat != 0 && a.lng != 0 {
			buf, err := a.makeRequest(endPoint)

			if err != nil {
				return response, err
			}

			// Unmarshal data to struct
			err = json.Unmarshal(buf, &response)
			if err != nil {
				return response, err
			}
			return response, nil
		}
		return response, errors.New("Error: Lat or Lng not set")
	}
	return response, errors.New("Error: Access Key not set")
}

// PositionStatus tests if inside geofence specified by accessKey and additionally provides
// navigation information in or out and to origin
func (a *API) PositionStatus() (PositionResponse, error) {
	var response PositionResponse
	if a.accessKey != "" {
		action := "position/status"
		endPoint := a.baseURI + a.version + "/" + action + "/" + a.accessKey
		if a.lat != 0 && a.lng != 0 {
			buf, err := a.makeRequest(endPoint)

			if err != nil {
				return response, err
			}

			// Unmarshal data to struct
			err = json.Unmarshal(buf, &response)
			if err != nil {
				return response, err
			}
			return response, nil
		}
		return response, errors.New("Error: Lat or Lng not set")
	}
	return response, errors.New("Error: Access Key not set")
}

// SetAPIKey sets API key for request
func (a *API) SetAPIKey(apiKey string) {
	// Set some defaults
	a.baseURI = "https://api.fencer.io/"
	a.version = "v1.0"
	a.apiKey = apiKey
}

// SetAccessKey sets Access Key for request
func (a *API) SetAccessKey(accessKey string) {
	a.accessKey = accessKey
}

// SetVersion sets version of API to use
func (a *API) SetVersion(version string) {
	a.version = version
}

// SetLat sets the latitude coordinate position
func (a *API) SetLat(lat float64) {
	a.lat = lat
}

// SetLng sets the longitude coordinate position
func (a *API) SetLng(lng float64) {
	a.lng = lng
}

// SetLatLng sets both the latitude and longitude positions
func (a *API) SetLatLng(lat float64, lng float64) {
	a.lat = lat
	a.lng = lng
}

// Private method to make the request and process the response
func (a *API) makeRequest(endPoint string) ([]byte, error) {
	// Custom client for timeouts
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	request, _ := http.NewRequest("GET", endPoint, nil)

	if a.apiKey != "" {
		request.Header.Set("Authorization", a.apiKey)
		if a.lat != 0 && a.lng != 0 {
			lat := fmt.Sprintf("%f", a.lat)
			lng := fmt.Sprintf("%f", a.lng)
			request.Header.Set("Lat-Pos", lat)
			request.Header.Set("Lng-Pos", lng)
		}

		response, err := netClient.Do(request)

		defer response.Body.Close()
		if err != nil {
			return nil, err
		}

		if response.StatusCode != 200 {
			return nil, errors.New("Error: " + string(response.StatusCode))
		}

		// Handle response
		buf, _ := ioutil.ReadAll(response.Body)

		if err != nil {
			return nil, err
		}

		return buf, nil

	} else {
		return nil, errors.New("Error: No API key set")
	}
}
