package gofencer_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/fencerio/gofencer"
)

// Provide an API variable for greater test coverage
var apiKey = os.Getenv("API_KEY")

func TestGetGeofences(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		response, err := f.GetGeofences()
		if err != nil {
			t.Error(err)
			return
		}
		if response.Data[0].Status != "Public" && response.Data[0].Status != "Private" {
			t.Error("Got", response.Data[0].Status, "expected Public or Private")
		}
	}
}

func TestGetGeofencesNoApiKey(t *testing.T) {
	// Only run test if we have a key
	f := new(gofencer.API)
	_, err := f.GetGeofences()
	if err == nil {
		t.Error("Expected failure due to no API Key, did not fail")
	}
}

func TestNavigateIn(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.6750352651078, -2.4879334942895883)
		response, err := f.NavigateIn()
		if err != nil {
			t.Error(err)
			return
		}

		if response.Data.Distance != 5.875 {
			t.Error("Got", response.Data.Distance, "expected distance of 5.875")
		}

		if response.Data.Bearing != 359 {
			t.Error("Got", response.Data.Bearing, "expected distance of 359")
		}
	}
}

func TestNavigateInFromInside(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.72874258598, -2.48936793915987)
		_, err := f.NavigateIn()
		if err == nil {
			t.Error("This point is inside goefence, should not be able to in")
		}
	}
}

func TestNavigateOrigin(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.6750352651078, -2.4879334942895883)
		response, err := f.NavigateOrigin()
		if err != nil {
			t.Error(err)
			return
		}

		if response.Data.Distance != 5.968 {
			t.Error("Got", response.Data.Distance, "expected distance of 5.968")
		}

		if response.Data.Bearing != 359 {
			t.Error("Got", response.Data.Bearing, "expected distance of 359")
		}
	}
}

func TestNavigateOut(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.72874258598, -2.48936793915987)
		response, err := f.NavigateOut()
		if err != nil {
			t.Error(err)
			return
		}

		if response.Data.Distance != 0.082 {
			t.Error("Got", response.Data.Distance, "expected distance of 0.082")
		}

		if response.Data.Bearing != 287 {
			t.Error("Got", response.Data.Bearing, "expected distance of 287")
		}
	}
}

func TestNavigateOutFromOutside(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.6750352651078, -2.4879334942895883)
		_, err := f.NavigateOut()
		if err == nil {
			t.Error("This point is outside goefence, should not be able to navigate out")
		}
	}
}

func TestPositionInside(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.72874258598, -2.48936793915987)
		response, err := f.PositionInside()
		if err != nil {
			t.Error(err)
			return
		}

		if response.Data.Inside != true {
			t.Error("Got", response.Data.Inside, "expected true")
		}
	}
}

func TestPositionInside2(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.6750352651078, -2.4879334942895883)
		response, err := f.PositionInside()
		if err != nil {
			t.Error(err)
			return
		}

		if response.Data.Inside != false {
			t.Error("Got", response.Data.Inside, "expected false")
		}
	}
}

func TestPositionStatus(t *testing.T) {
	if apiKey != "" {
		// Only run test if we have a key
		f := new(gofencer.API)
		f.SetAPIKey(apiKey)
		f.SetAccessKey("4c89693a-02c4-4dd5-be7a-cdf801237e3d")
		f.SetLatLng(53.6750352651078, -2.4879334942895883)
		response, err := f.PositionStatus()
		if err != nil {
			t.Error(err)
			return
		}

		if response.Data.Inside != false {
			t.Error("Got", response.Data.Inside, "expected distance of true")
		}

		if response.Data.Enter.Distance != 5.875 {
			t.Error("Got", response.Data.Enter.Distance, "expected distance of 5.875")
		}

		if response.Data.Enter.Bearing != 359 {
			t.Error("Got", response.Data.Enter.Bearing, "expected bearing of 359")
		}

		if response.Data.Origin.Distance != 5.968 {
			t.Error("Got", response.Data.Origin.Distance, "expected distance of 5.968")
		}

		if response.Data.Origin.Bearing != 359 {
			t.Error("Got", response.Data.Origin.Bearing, "expected bearing of 359")
		}
	}
}

func TestSetAPIKey(t *testing.T) {
	f := new(gofencer.API)
	f.SetAPIKey("123456789")

	v := reflect.ValueOf(*f)
	apiKey := v.FieldByName("apiKey")

	// Check API key set
	if apiKey.String() != "123456789" {
		t.Error("API Key not correctly set")
	}

	// Check API base URI set
	baseURI := v.FieldByName("baseURI")
	if baseURI.String() != "https://api.fencer.io/" {
		t.Error("Base URI not correctly initialized")
	}

	// Check version set
	version := v.FieldByName("version")
	if version.String() != "v1.0" {
		t.Error("Version not correctly initialized")
	}
}

func TestSetAccessKey(t *testing.T) {
	f := new(gofencer.API)
	f.SetAccessKey("123456789")

	v := reflect.ValueOf(*f)
	accessKey := v.FieldByName("accessKey")

	// Check Access key set
	if accessKey.String() != "123456789" {
		t.Error("Geofence Access Key not correctly set")
	}
}

func TestSetVersion(t *testing.T) {
	f := new(gofencer.API)
	f.SetVersion("v1.1")

	v := reflect.ValueOf(*f)
	version := v.FieldByName("version")

	// Check Version set
	if version.String() != "v1.1" {
		t.Error("Version was not correctly set")
	}
}

func TestSetLat(t *testing.T) {
	f := new(gofencer.API)
	f.SetLat(53.12312321)

	v := reflect.ValueOf(*f)
	lat := v.FieldByName("lat")

	// Check Latitude set
	if lat.Float() != 53.12312321 {
		t.Error("Latitude was not correctly set")
	}
}

func TestSetLng(t *testing.T) {
	f := new(gofencer.API)
	f.SetLng(-2.354545654)

	v := reflect.ValueOf(*f)
	lng := v.FieldByName("lng")

	// Check Longitude set
	if lng.Float() != -2.354545654 {
		t.Error("Longitude was not correctly set")
	}
}

func TestSetLatLng(t *testing.T) {
	f := new(gofencer.API)
	f.SetLatLng(53.12312321, -2.354545654)

	v := reflect.ValueOf(*f)

	lat := v.FieldByName("lat")

	// Check Latitude set
	if lat.Float() != 53.12312321 {
		t.Error("Latitude was not correctly set in setLatLng")
	}

	lng := v.FieldByName("lng")

	// Check Longitude set
	if lng.Float() != -2.354545654 {
		t.Error("Longitude was not correctly set in setLatLng")
	}
}
