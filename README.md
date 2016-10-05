# Gofencer
## Go library for accessing the Fencer.io API

Download with ```go get github.com/fencer.io/gofencer```

Consult the package and/or the API docs to establish which arguments are required. https://fencer.io/developers

Typical set up would be something like this.

```
package main 

import (
	"log"
	"github.com/fencer.io/gofencer"
)

func main() {
	apiKey := "f7c04eaf-5510-5878-xxxxx-acad3c60f"
	accessKey := "4c89693a-02c4-xxxx-xxxx-cdf801237e3d"
	lat := 53.6750352651078
	lng := -2.4879334942895883

	f := new (gofencer.API)
	f.SetAPIKey(apiKey)
	f.SetAccessKey(accessKey)
	f.SetLatLng(lat, lng)

	// Example call to the API
	// Determine if passed lat/lng is inside geofence with
	response, err := f.PositionInside()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response.Data.Inside) // true or false

}
```

## License

GPL license
